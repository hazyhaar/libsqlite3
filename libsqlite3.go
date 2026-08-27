// Copyright 2023 The libsqlite3-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generator.go

// Package libsqlite3 is a ccgo/v4 version of libsqlite3 (SQLite, http://sqlite.org)
//
// # Supported platforms and architectures - Tier 1
//
// Tier 1 platforms are the primary, officially supported targets.  When a new
// version is released, any critical bugs found on Tier 1 platforms are treated
// as release blockers. The release will be postponed until such issues are
// resolved.
//
//	OS      Arch    SQLite version
//	------------------------------
//	darwin     amd64   3.53.4
//	darwin     arm64   3.53.4
//	freebsd    386     3.53.4
//	freebsd    amd64   3.53.4
//	freebsd    amd64   3.53.4
//	freebsd    arm     3.53.4
//	linux      386     3.53.4
//	linux      amd64   3.53.4
//	linux      arm     3.53.4
//	linux      arm64   3.53.4
//	linux      loong64 3.53.4
//	linux      ppc64le 3.53.4
//	linux      riscv64 3.53.4
//	linux      s390x   3.53.4
//	netbsd     amd64   3.53.4
//	openbsd7.8 amd64   3.53.4
//	openbsd7.8 arm64   3.53.4
//	windows    386     3.53.4
//	windows    amd64   3.53.4
//	windows    arm64   3.53.4
//
// # Supported platforms and architectures - Tier 2
//
// Tier 2 platforms are supported by on a best-effort basis. Critical bugs on
// Tier 2 platforms do not block new releases. However, fixes contributed by
// external contributors are very welcome and encouraged. Tier 2 support
// guarantees only that the package will build and that at least some tests are
// passing.
//
// # There are no Tier 2 platforms at the moment
//
// # File locking on Linux
//
// Database files are locked with POSIX record locks, as upstream SQLite does.
// On linux the library can use Open File Description locks (F_OFD_SETLK)
// instead, which survive a close() of an unrelated descriptor of the same
// file. This is opt-in and process-wide: set the environment variable
// MODERNC_SQLITE_OFD_LOCK=1 before the process opens its first database file,
// or call Xmodernc_ofd_locking(tls, 1) before that point. The function returns
// the previous setting; -1 where OFD locks are not available; -2 when the
// setting can no longer be changed because the process has already attempted
// a lock - the kind of lock cannot change under a held one, so it is frozen
// from the first lock until Xsqlite3_shutdown. A negative argument only
// queries. It exists on every unix target and not on windows. See
// CHANGELOG.md (2026-08-27) and https://gitlab.com/cznic/sqlite/-/issues/255.
//
// # Builders
//
// Builder results available at:
//
// https://modern-c.appspot.com/-/builder/?importpath=modernc.org%2flibsqlite3
package libsqlite3 // import "modernc.org/libsqlite3"
