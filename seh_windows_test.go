// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libsqlite3

import (
	"syscall"
	"testing"
	"unsafe"
)

// sehNoAccessPage returns the address of a page that faults when read: a
// reserved but never committed region.
func sehNoAccessPage(t *testing.T) uintptr {
	t.Helper()
	const (
		memReserve   = 0x2000
		memRelease   = 0x8000
		pageNoAccess = 0x01
	)
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	r, _, err := kernel32.NewProc("VirtualAlloc").Call(0, 65536, memReserve, pageNoAccess)
	if r == 0 {
		t.Fatal(err)
	}

	t.Cleanup(func() { kernel32.NewProc("VirtualFree").Call(r, 0, memRelease) })
	return r
}

// TestSEHWindowsPageNoAccess verifies that a live WAL-index page 0 protected
// with PAGE_NOACCESS is intercepted by SEH during sqlite3_exec, returning
// SQLITE_IOERR_IN_PAGE without crashing, and leaving the connection healthy.
func TestSEHWindowsPageNoAccess(t *testing.T) {
	c, _ := sehSetup(t)
	defer c.close()

	// Walk ccgo structures: Tsqlite3 -> FaDb[0].FpBt -> FpPager -> FpWal
	db := (*Tsqlite3)(unsafe.Pointer(c.db))
	pDb0 := (*TDb)(unsafe.Pointer(db.FaDb))
	pBtree := (*TBtree)(unsafe.Pointer(pDb0.FpBt))
	pBtShared := (*TBtShared)(unsafe.Pointer(pBtree.FpBt))
	pPager := (*TPager)(unsafe.Pointer(pBtShared.FpPager))
	pWal := (*TWal)(unsafe.Pointer(pPager.FpWal))

	if pWal.FnWiData < 1 || pWal.FapWiData == 0 {
		t.Fatalf("wal structures not initialized: nWiData=%d, apWiData=%x", pWal.FnWiData, pWal.FapWiData)
	}

	page0 := *(*uintptr)(unsafe.Pointer(pWal.FapWiData))
	if page0 == 0 {
		t.Fatalf("page0 is null")
	}

	// Release read lock to force walIndexReadHdr on next statement
	pWal.FreadLock = -1

	const pageNoAccess = 0x01
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	virtualProtect := kernel32.NewProc("VirtualProtect")

	var oldProtect uint32
	r, _, err := virtualProtect.Call(page0, 4096, pageNoAccess, uintptr(unsafe.Pointer(&oldProtect)))
	if r == 0 {
		t.Fatalf("VirtualProtect PAGE_NOACCESS failed: %v", err)
	}

	// Execute query: must be caught by SEH and return SQLITE_IOERR_IN_PAGE (8714)
	rc, msg := c.exec("SELECT count(*) FROM t")

	// Restore page protection immediately
	virtualProtect.Call(page0, 4096, uintptr(oldProtect), uintptr(unsafe.Pointer(&oldProtect)))

	if rc != SQLITE_IOERR_IN_PAGE {
		t.Fatalf("expected SQLITE_IOERR_IN_PAGE (%d), got rc=%d (msg: %s)", SQLITE_IOERR_IN_PAGE, rc, msg)
	}
	t.Logf("Fault successfully caught by engine SEH: rc=%d, msg=%s", rc, msg)

	// Verify reuse of the SAME connection (upstream MSVC recovery contract)
	n := c.queryInt(t, "SELECT count(*) FROM t")
	if n != 3 {
		t.Fatalf("expected count=3 after recovery, got %d", n)
	}
	t.Logf("Recovery contract verified: the same connection remains usable after fault recovery (count=%d).", n)
}
