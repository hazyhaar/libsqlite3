// Copyright 2026 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

package libsqlite3

import (
	"syscall"
	"testing"
	"unsafe"
)

// sehNoAccessPage returns the address of a page that faults when read.
func sehNoAccessPage(t *testing.T) uintptr {
	t.Helper()
	b, err := syscall.Mmap(-1, 0, 65536, syscall.PROT_NONE, syscall.MAP_PRIVATE|syscall.MAP_ANON)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { syscall.Munmap(b) })
	return uintptr(unsafe.Pointer(unsafe.SliceData(b)))
}
