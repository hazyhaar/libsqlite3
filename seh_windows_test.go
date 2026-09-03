// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libsqlite3

import (
	"syscall"
	"testing"
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
