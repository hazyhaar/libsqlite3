// Copyright 2026 The libsqlite3-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libsqlite3 // import "modernc.org/libsqlite3"

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"unsafe"

	"modernc.org/libc"
)

var oInnerOFD = flag.Bool("inner-ofd", false, "internal use")

// TestOFDLocking is the positive control for TestTclTestOFD and
// TestConcurrentProcessesOFD, which only set MODERNC_SQLITE_OFD_LOCK for their
// child processes and would stay green if the variable were misspelled, no
// longer read before the first lock or no longer inherited: it proves that
// the variable selects the kind of kernel lock the library places, and that
// OFD locks survive what the switch exists for, a close() of an unrelated
// descriptor of the database file. The library settles the mode once per
// process, so every value is observed in a child process: the test binary
// re-executes itself with -inner-ofd, the child opens a database, takes
// RESERVED (BEGIN IMMEDIATE), reports the /proc/locks entries of the database
// inode, closes and reopens a stray descriptor of the file, and reports again.
func TestOFDLocking(t *testing.T) {
	if *oInnerOFD {
		innerOFDLocking(t)
		return
	}

	self, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		env      string // "" = variable unset
		kind     string // /proc/locks lock type expected before the stray close
		survives bool   // locks still there after the stray close
	}{
		{"", "POSIX", false},
		{"0", "POSIX", false},
		{"1", "OFDLCK", true},
		{"yes", "OFDLCK", true},
	} {
		name := "unset"
		if tc.env != "" {
			name = tc.env
		}
		t.Run(name, func(t *testing.T) {
			cmd := exec.Command(self, "-test.run=^TestOFDLocking$", "-test.v", "-inner-ofd")
			for _, v := range os.Environ() {
				if !strings.HasPrefix(v, ofdEnv+"=") {
					cmd.Env = append(cmd.Env, v)
				}
			}
			if tc.env != "" {
				cmd.Env = append(cmd.Env, ofdEnv+"="+tc.env)
			}
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("%s\n%v", out, err)
			}

			var before, after []string
			for _, line := range strings.Split(string(out), "\n") {
				switch {
				case strings.HasPrefix(line, "before: "):
					before = append(before, line[len("before: "):])
				case strings.HasPrefix(line, "after: "):
					after = append(after, line[len("after: "):])
				}
			}
			t.Logf("%s=%q: before the stray close %q, after %q", ofdEnv, tc.env, before, after)
			if len(before) == 0 {
				t.Fatalf("no kernel lock on the database inode inside BEGIN IMMEDIATE:\n%s", out)
			}

			for _, v := range before {
				if !strings.HasPrefix(v, tc.kind+" ") {
					t.Errorf("expected %s locks, got %q", tc.kind, v)
				}
			}
			switch {
			case tc.survives && len(after) != len(before):
				t.Errorf("OFD locks must survive a close() of another descriptor: before %d, after %d", len(before), len(after))
			case !tc.survives && len(after) != 0:
				t.Errorf("POSIX locks are dropped by a close() of another descriptor (upstream behaviour), but %d survived", len(after))
			}
		})
	}
}

// innerOFDLocking is the child side of TestOFDLocking.
func innerOFDLocking(t *testing.T) {
	path := filepath.Join(t.TempDir(), "ofd.db")
	tls := libc.NewTLS()
	defer tls.Close()

	zPath, err := libc.CString(path)
	if err != nil {
		t.Fatal(err)
	}

	defer libc.Xfree(tls, zPath)
	ppDb := libc.Xmalloc(tls, 8)
	defer libc.Xfree(tls, ppDb)
	if rc := Xsqlite3_open_v2(tls, zPath, ppDb, SQLITE_OPEN_READWRITE|SQLITE_OPEN_CREATE, 0); rc != SQLITE_OK {
		t.Fatalf("sqlite3_open_v2: %d", rc)
	}

	db := *(*uintptr)(unsafe.Pointer(ppDb))
	defer Xsqlite3_close(tls, db)
	exec1 := func(sql string) {
		zSQL, err := libc.CString(sql)
		if err != nil {
			t.Fatal(err)
		}

		defer libc.Xfree(tls, zSQL)
		if rc := Xsqlite3_exec(tls, db, zSQL, 0, 0, 0); rc != SQLITE_OK {
			t.Fatalf("%q: %d: %s", sql, rc, libc.GoString(Xsqlite3_errmsg(tls, db)))
		}
	}
	exec1("CREATE TABLE t(x); BEGIN IMMEDIATE; INSERT INTO t VALUES(1);")
	for _, v := range procLocks(t, path) {
		fmt.Printf("before: %s\n", v)
	}

	// The stray descriptor: what an application's own os.Open+Close of the
	// database file does to the process's POSIX locks, and not to OFD locks.
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}

	f.Close()
	for _, v := range procLocks(t, path) {
		fmt.Printf("after: %s\n", v)
	}
	exec1("COMMIT")
}

// procLocks returns the /proc/locks entries for the inode of path, as
// "<type> <access> <mode> ..." (the leading "N: " stripped).
func procLocks(t *testing.T, path string) (r []string) {
	fi, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}

	ino := fi.Sys().(*syscall.Stat_t).Ino
	b, err := os.ReadFile("/proc/locks")
	if err != nil {
		t.Fatal(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		if !strings.Contains(line, fmt.Sprintf(":%d ", ino)) {
			continue
		}

		if x := strings.Index(line, ": "); x >= 0 {
			line = line[x+2:]
		}
		r = append(r, strings.Join(strings.Fields(line), " "))
	}
	return r
}
