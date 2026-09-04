// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libsqlite3

import (
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"testing"
	"unsafe"

	"modernc.org/libc"
	"modernc.org/libc/sys/types"
)

// sehConn is a raw sqlite3* connection for the SEH tests.
type sehConn struct {
	tls *libc.TLS
	db  uintptr
}

func sehOpen(t *testing.T, path string) *sehConn {
	t.Helper()
	tls := libc.NewTLS()
	zPath, err := libc.CString(path)
	if err != nil {
		t.Fatal(err)
	}

	defer libc.Xfree(tls, zPath)

	ppDb := libc.Xcalloc(tls, 1, types.Size_t(unsafe.Sizeof(uintptr(0))))
	defer libc.Xfree(tls, ppDb)

	if rc := Xsqlite3_open_v2(tls, zPath, ppDb, SQLITE_OPEN_READWRITE|SQLITE_OPEN_CREATE, 0); rc != SQLITE_OK {
		t.Fatalf("sqlite3_open_v2: rc=%d", rc)
	}

	return &sehConn{tls: tls, db: *(*uintptr)(unsafe.Pointer(ppDb))}
}

func (c *sehConn) close() {
	Xsqlite3_close_v2(c.tls, c.db)
	c.tls.Close()
}

// exec runs sql and returns the extended result code and the error message.
func (c *sehConn) exec(sql string) (rc int32, msg string) {
	zSql, err := libc.CString(sql)
	if err != nil {
		return -1, err.Error()
	}

	defer libc.Xfree(c.tls, zSql)

	if rc = Xsqlite3_exec(c.tls, c.db, zSql, 0, 0, 0); rc != SQLITE_OK {
		return Xsqlite3_extended_errcode(c.tls, c.db), libc.GoString(Xsqlite3_errmsg(c.tls, c.db))
	}

	return SQLITE_OK, ""
}

func (c *sehConn) mustExec(t *testing.T, sql string) {
	t.Helper()
	if rc, msg := c.exec(sql); rc != SQLITE_OK {
		t.Fatalf("%s: rc=%d %s", sql, rc, msg)
	}
}

// queryInt runs a query that returns one integer.
func (c *sehConn) queryInt(t *testing.T, sql string) int32 {
	t.Helper()
	zSql, err := libc.CString(sql)
	if err != nil {
		t.Fatal(err)
	}

	defer libc.Xfree(c.tls, zSql)

	ppStmt := libc.Xcalloc(c.tls, 1, types.Size_t(unsafe.Sizeof(uintptr(0))))
	defer libc.Xfree(c.tls, ppStmt)

	if rc := Xsqlite3_prepare_v2(c.tls, c.db, zSql, -1, ppStmt, 0); rc != SQLITE_OK {
		t.Fatalf("%s: prepare rc=%d %s", sql, rc, libc.GoString(Xsqlite3_errmsg(c.tls, c.db)))
	}

	pStmt := *(*uintptr)(unsafe.Pointer(ppStmt))
	defer Xsqlite3_finalize(c.tls, pStmt)

	if rc := Xsqlite3_step(c.tls, pStmt); rc != SQLITE_ROW {
		t.Fatalf("%s: step rc=%d %s", sql, rc, libc.GoString(Xsqlite3_errmsg(c.tls, c.db)))
	}

	return Xsqlite3_column_int(c.tls, pStmt, 0)
}

// sehSetup opens a WAL mode database with some content, so that the -shm file
// exists and is mapped.
func sehSetup(t *testing.T) (c *sehConn, path string) {
	t.Helper()
	path = filepath.Join(t.TempDir(), "seh.db")
	c = sehOpen(t, path)
	c.mustExec(t, "PRAGMA journal_mode=WAL")
	c.mustExec(t, "CREATE TABLE t(x)")
	c.mustExec(t, "INSERT INTO t VALUES(1),(2),(3)")
	if fi, err := os.Stat(path + "-shm"); err != nil || fi.Size() < walIndexPgsz {
		t.Fatalf("-shm: %v %v", fi, err)
	}

	return c, path
}

// TestSEHTruncatedShm truncates the -shm file under the live mapping: the next
// access to the wal-index faults, which the emulation must turn into
// SQLITE_IOERR_IN_PAGE; once the file is restored the connection recovers by
// rebuilding the wal-index from the -wal file.
func TestSEHTruncatedShm(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("a mapped file cannot be truncated on windows")
	}

	c, path := sehSetup(t)
	defer c.close()

	shm := path + "-shm"
	if err := os.Truncate(shm, 0); err != nil {
		t.Fatal(err)
	}

	if rc, msg := c.exec("SELECT count(*) FROM t"); rc != SQLITE_IOERR_IN_PAGE {
		t.Fatalf("got rc=%d (%s), want SQLITE_IOERR_IN_PAGE", rc, msg)
	}

	if err := os.Truncate(shm, walIndexPgsz); err != nil {
		t.Fatal(err)
	}

	c.mustExec(t, "SELECT count(*) FROM t")
	c.mustExec(t, "INSERT INTO t VALUES(4)")
	c.mustExec(t, "PRAGMA wal_checkpoint(TRUNCATE)")
	c.mustExec(t, "SELECT count(*) FROM t")
}

// TestSEHInjectedFault raises a simulated fault at every SEH_INJECT_FAULT site
// that a read, a write and a checkpoint reach, one at a time, and checks that
// each turns into SQLITE_IOERR_IN_PAGE and that the connection works normally
// afterwards.
func TestSEHInjectedFault(t *testing.T) {
	c, _ := sehSetup(t)
	defer c.close()

	for _, sql := range []string{
		"SELECT count(*) FROM t",
		"INSERT INTO t VALUES(4)",
		"PRAGMA wal_checkpoint(PASSIVE)",
		"PRAGMA wal_checkpoint(TRUNCATE)",
		"BEGIN; INSERT INTO t VALUES(5); SAVEPOINT s; INSERT INTO t VALUES(6); ROLLBACK TO s; COMMIT",
	} {
		sites := 0
		for n := int32(1); ; n++ {
			SehInject(n)
			rc, msg := c.exec(sql)
			pending := SehPending()
			SehInject(0)
			if pending > 0 { // fewer than n sites reached: the statement ran to completion
				if rc != SQLITE_OK {
					t.Fatalf("%s with %d pending sites: rc=%d %s", sql, pending, rc, msg)
				}

				break
			}

			sites++
			if rc != SQLITE_IOERR_IN_PAGE {
				t.Fatalf("%s, simulated fault at site %d: got rc=%d (%s), want SQLITE_IOERR_IN_PAGE", sql, n, rc, msg)
			}

			c.exec("ROLLBACK") // in case the statement list left a transaction open
			c.mustExec(t, "SELECT count(*) FROM t")
		}
		if sites == 0 {
			t.Fatalf("%s: no SEH_INJECT_FAULT site reached", sql)
		}

		t.Logf("%-90s %d sites", sql, sites)
	}
}

// TestSEHCheckpointInvalidatesCache guards the invariant the checkpoint thunk
// must preserve: sqlite3WalCheckpoint() learns through isChanged that
// walIndexReadHdr() loaded a new wal-index header inside the protected block
// and zeroes pWal->hdr afterwards, so that the connection's next read
// transaction sees the header as changed and resets its page cache. A thunk
// that keeps isChanged to itself leaves the checkpointing connection serving
// pages from before the other connection's write, in every checkpoint mode.
// The test needs no fault and passes with or without the emulation.
func TestSEHCheckpointInvalidatesCache(t *testing.T) {
	for _, mode := range []string{"PASSIVE", "FULL", "RESTART", "TRUNCATE"} {
		t.Run(mode, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "seh.db")
			a := sehOpen(t, path)
			defer a.close()

			a.mustExec(t, "PRAGMA journal_mode=WAL")
			a.mustExec(t, "CREATE TABLE t(x)")
			a.mustExec(t, "INSERT INTO t VALUES(1)")
			if got := a.queryInt(t, "SELECT x FROM t"); got != 1 { // a's page cache now holds the row
				t.Fatalf("got x=%d, want 1", got)
			}

			b := sehOpen(t, path)
			defer b.close()

			b.mustExec(t, "UPDATE t SET x=2")
			a.mustExec(t, "PRAGMA wal_checkpoint("+mode+")") // a loads b's header inside the protected block
			if got := a.queryInt(t, "SELECT x FROM t"); got != 2 {
				t.Errorf("stale page served from a's cache after wal_checkpoint(%s): got x=%d, want 2", mode, got)
			}
		})
	}
}

// The body, the injecting body and the on-fault callback of TestSEHTrampoline
// are top-level functions with package-level state, like the transpiled thunks
// the trampoline is called with in the library: __ccgo_fp of a top-level
// function is the address of its static funcval. A Go closure would not do: a
// closure that does not escape lives on the goroutine stack, the panic path
// between the fault and the callback can grow the stack (sehLog allocates),
// and after such a move the uintptr the trampoline holds points into the old
// copy of the stack, so the callback would run with a stale context and
// update a dead copy of its captured variables. (Seen on linux/amd64 with
// go1.27: faultCalls stayed 0 in four runs out of five.)
var sehTramp struct {
	bodyCalls, faultCalls int
}

func sehTrampBody(tls *libc.TLS, pWal, pArg uintptr) int32 {
	sehTramp.bodyCalls++
	return int32(*(*byte)(unsafe.Pointer(pArg)))
}

func sehTrampInject(tls *libc.TLS, pWal, pArg uintptr) int32 {
	_modernc_seh_inject(tls, pWal)
	return 1
}

func sehTrampOnFault(tls *libc.TLS, pWal uintptr) int32 {
	sehTramp.faultCalls++
	return 4242
}

// TestSEHTrampoline drives modernc_seh_try directly with a body that reads a
// page that cannot be accessed, standing in for an unreadable -shm mapping.
func TestSEHTrampoline(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	page := sehNoAccessPage(t)

	// A fake Wal whose only wal-index page is the inaccessible one.
	pWal := libc.Xcalloc(tls, 1, types.Size_t(unsafe.Sizeof(TWal{})))
	defer libc.Xfree(tls, pWal)

	apWiData := libc.Xcalloc(tls, 1, types.Size_t(unsafe.Sizeof(uintptr(0))))
	defer libc.Xfree(tls, apWiData)

	*(*uintptr)(unsafe.Pointer(apWiData)) = page
	w := (*TWal)(unsafe.Pointer(pWal))
	w.FapWiData = apWiData
	w.FnWiData = 1

	okByte := libc.Xcalloc(tls, 1, 1)
	defer libc.Xfree(tls, okByte)

	*(*byte)(unsafe.Pointer(okByte)) = 7

	sehTramp.bodyCalls, sehTramp.faultCalls = 0, 0
	xBody, xInject, xOnFault := __ccgo_fp(sehTrampBody), __ccgo_fp(sehTrampInject), __ccgo_fp(sehTrampOnFault)

	// No fault: the body's result comes through.
	if rc := _modernc_seh_try(tls, pWal, xBody, okByte, xOnFault); rc != 7 || sehTramp.bodyCalls != 1 || sehTramp.faultCalls != 0 {
		t.Fatalf("rc=%d bodyCalls=%d faultCalls=%d", rc, sehTramp.bodyCalls, sehTramp.faultCalls)
	}

	// A fault in a wal-index page runs the on-fault callback.
	if rc := _modernc_seh_try(tls, pWal, xBody, page, xOnFault); rc != 4242 || sehTramp.bodyCalls != 2 || sehTramp.faultCalls != 1 {
		t.Fatalf("rc=%d bodyCalls=%d faultCalls=%d", rc, sehTramp.bodyCalls, sehTramp.faultCalls)
	}

	// ... or yields SQLITE_IOERR_IN_PAGE without one.
	if rc := _modernc_seh_try(tls, pWal, xBody, page, 0); rc != SQLITE_IOERR_IN_PAGE {
		t.Fatalf("rc=%d", rc)
	}

	// A simulated fault is handled the same way.
	SehInject(1)
	rc := _modernc_seh_try(tls, pWal, xInject, 0, xOnFault)
	SehInject(0)
	if rc != 4242 || sehTramp.faultCalls != 2 {
		t.Fatalf("rc=%d faultCalls=%d", rc, sehTramp.faultCalls)
	}

	// The guard is restored afterwards.
	if debug.SetPanicOnFault(false) {
		t.Fatal("SetPanicOnFault left enabled")
	}

	// A fault outside the wal-index pages is not the emulation's business and
	// propagates as a crash would, here as a panic the test recovers itself.
	w.FnWiData = 0
	func() {
		defer func() {
			if e := recover(); e == nil {
				t.Fatal("expected the fault to propagate")
			}
		}()
		debug.SetPanicOnFault(true)
		defer debug.SetPanicOnFault(false)
		_modernc_seh_try(tls, pWal, xBody, page, xOnFault)
	}()
	if sehTramp.faultCalls != 2 {
		t.Fatalf("faultCalls=%d", sehTramp.faultCalls)
	}
}
