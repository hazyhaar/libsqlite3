// Copyright 2023 The libsqlite3-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"modernc.org/cc/v4"
	ccgo "modernc.org/ccgo/v4/lib"
	util "modernc.org/fileutil/ccgo"
)

const (
	archivePath  = "sqlite-amalgamation-" + versionTag + ".zip"
	archive2Path = "sqlite-src-" + versionTag + ".zip"
	versionTag   = "3530100"
)

var (
	goarch    = env("TARGET_GOARCH", env("GOARCH", runtime.GOARCH))
	goos      = env("TARGET_GOOS", env("GOOS", runtime.GOOS))
	target    = fmt.Sprintf("%s/%s", goos, goarch)
	sed       = "sed"
	j         = fmt.Sprint(runtime.GOMAXPROCS(-1))
	win       = os.Getenv("GO_GENERATE_WIN") == "1"
	win32     = os.Getenv("GO_GENERATE_WIN32") == "1"
	xgcc      string
	withTclsh = os.Getenv("MODERNC_ORG_SQLITE_WITH_TCLSH")
)

func fail(rc int, msg string, args ...any) {
	fmt.Fprintln(os.Stderr, strings.TrimSpace(fmt.Sprintf(msg, args...)))
	os.Exit(rc)
}

func main() {
	if ccgo.IsExecEnv() {
		if err := ccgo.NewTask(goos, goarch, os.Args, os.Stdout, os.Stderr, nil).Main(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		return
	}

	if !win && !win32 && target == "linux/amd64" && os.Getenv("GO_GENERATE_NOWIN") == "" {
		defer func() {
			util.MustShell(true, nil, "make", "windows", "windows_386")
			util.MustCopyFile(true, "internal/autogen/windows_386.mod", "go.mod", nil)
			util.MustCopyFile(true, "internal/autogen/windows_amd64.mod", "go.mod", nil)
			util.MustCopyFile(true, "internal/autogen/windows_arm64.mod", "go.mod", nil)
		}()
	}

	switch {
	case win:
		if target != "linux/amd64" {
			fail(1, "cross compiling for windows is supported only on linux/amd64")
		}

		goos = "windows"
		goarch = "amd64"
		xgcc = strings.TrimSpace(string(util.MustShell(true, nil, "which", "x86_64-w64-mingw32-gcc")))
	case win32:
		if target != "linux/amd64" {
			fail(1, "cross compiling for windows is supported only on linux/amd64")
		}

		goos = "windows"
		goarch = "386"
		xgcc = strings.TrimSpace(string(util.MustShell(true, nil, "which", "i686-w64-mingw32-gcc")))
	}

	switch target {
	case "freebsd/amd64", "freebsd/arm64", "openbsd/amd64", "openbsd/arm64", "darwin/amd64", "darwin/arm64", "netbsd/amd64":
		sed = "gsed"
	}

	f, err := os.Open(archivePath)
	if err != nil {
		fail(1, "cannot open zip file: %v\n", err)
	}

	f.Close()

	if f, err = os.Open(archive2Path); err != nil {
		fail(1, "cannot open zip file: %v\n", err)
	}

	f.Close()

	_, extractedArchivePath := filepath.Split(archivePath)
	extractedArchivePath = extractedArchivePath[:len(extractedArchivePath)-len(".zip")]
	tempDir := os.Getenv("GO_GENERATE_DIR")
	dev := os.Getenv("GO_GENERATE_DEV") != ""
	switch {
	case tempDir != "":
		os.RemoveAll(filepath.Join(tempDir, extractedArchivePath))
	default:
		var err error
		if tempDir, err = os.MkdirTemp("", "libsqlite3-generate"); err != nil {
			fail(1, "creating temp dir: %v\n", err)
		}

		defer func() {
			switch os.Getenv("GO_GENERATE_KEEP") {
			case "":
				os.RemoveAll(tempDir)
			default:
				fmt.Printf("%s: temporary directory kept\n", tempDir)
			}
		}()
	}
	libRoot := filepath.Join(tempDir, extractedArchivePath)
	makeRoot := libRoot
	fmt.Fprintf(os.Stderr, "archivePath %s\n", archivePath)
	fmt.Fprintf(os.Stderr, "extractedArchivePath %s\n", extractedArchivePath)
	fmt.Fprintf(os.Stderr, "tempDir %s\n", tempDir)
	fmt.Fprintf(os.Stderr, "libRoot %s\n", libRoot)
	fmt.Fprintf(os.Stderr, "makeRoot %s\n", makeRoot)
	util.MustShell(true, nil, "unzip", "-q", archivePath, "-d", tempDir)

	// https://gitlab.com/cznic/sqlite/-/issues/173
	util.MustShell(true, nil, "patch", filepath.Join(libRoot, "sqlite3.c"), filepath.Join("internal", "sqlite_issue173.patch"))
	// https://gitlab.com/cznic/libsqlite3/-/issues/1
	util.MustShell(true, nil, "patch", filepath.Join(libRoot, "sqlite3.c"), filepath.Join("internal", "issue1.patch"))

	// https://gitlab.com/cznic/sqlite/-/issues/180
	// We do not have long double, the field is already zero, skip the C-racy test altogether.
	util.MustShell(true, nil, sed, "-i", `s/sqlite3Config.bUseLongDouble = hasHighPrecisionDouble(rc);/\/\/ disabled/`, filepath.Join(libRoot, "sqlite3.c"))
	// Another C-race, enforce atomic access.
	util.MustShell(true, nil, sed, "-i", `0,/int isInit;*True after/{s/int isInit/volatile int isInit/}`, filepath.Join(libRoot, "sqlite3.c"))

	util.MustShell(true, nil, sed, "-i", `s/#if (defined(__GNUC__) || defined(__clang__)) \\/#if 0 \&\& (defined(__GNUC__) || defined(__clang__)) \\/`, filepath.Join(libRoot, "sqlite3.c"))

	fixWin(tempDir)
	cwd := util.MustAbsCwd(true)
	ilibc := filepath.Join(cwd, "..", "libc", "include", goos, goarch)
	ilibz := filepath.Join(cwd, "..", "libz", "include", goos, goarch)
	ilibtcl := filepath.Join(cwd, "..", "libtcl8.6", "include", goos, goarch)
	result := "sqlite3.go"
	util.MustInDir(true, makeRoot, func() (err error) {
		util.MustShell(true, nil, "sh", "-c", "go mod init example.com/libsqlite3 ; go get modernc.org/libc@v1.72.2 modernc.org/libz@v0.17.3 modernc.org/libtcl8.6@v0.17.4")
		config := []string{os.Args[0]}
		if dev {
			util.MustShell(true, nil, "sh", "-c", "go work init ; go work use $GOPATH/src/modernc.org/libc $GOPATH/src/modernc.org/libz $GOPATH/src/modernc.org/libtcl8.6")
			config = append(config,
				"-absolute-paths",
				"-keep-object-files",
				"-positions",
			)
		}
		if m64Double := cc.LongDouble64Flag(goos, goarch); m64Double != "" {
			config = append(config, m64Double)
		}
		config = append(config,
			"--package-name", "libsqlite3",
			"--prefix-enumerator=_",
			"--prefix-external=x_",
			"--prefix-field=F",
			"--prefix-static-internal=_",
			"--prefix-static-none=_",
			"--prefix-tagged-enum=_",
			"--prefix-tagged-struct=T",
			"--prefix-tagged-union=T",
			"--prefix-typename=T",
			"--prefix-undefined=_",
			"-ignore-unsupported-alignment",
			"-ignore-link-errors",
			"-import=sync",

			"-DHAVE_USLEEP",
			"-DLONGDOUBLE_TYPE=double",
			"-DNDEBUG",
			"-DSQLITE_DEFAULT_MEMSTATUS=0",
			"-DSQLITE_DISABLE_INTRINSIC",
			"-DSQLITE_ENABLE_COLUMN_METADATA",
			"-DSQLITE_ENABLE_DBPAGE_VTAB",
			"-DSQLITE_ENABLE_DBSTAT_VTAB",
			"-DSQLITE_ENABLE_FTS5",
			"-DSQLITE_ENABLE_GEOPOLY",
			"-DSQLITE_ENABLE_JSON1",
			"-DSQLITE_ENABLE_MATH_FUNCTIONS",
			"-DSQLITE_ENABLE_MEMORY_MANAGEMENT",
			"-DSQLITE_ENABLE_OFFSET_SQL_FUNC",
			"-DSQLITE_ENABLE_PREUPDATE_HOOK",
			"-DSQLITE_ENABLE_RBU",
			"-DSQLITE_ENABLE_RTREE",
			"-DSQLITE_ENABLE_SESSION",
			"-DSQLITE_ENABLE_SNAPSHOT",
			"-DSQLITE_ENABLE_STAT4",
			"-DSQLITE_ENABLE_UNLOCK_NOTIFY",
			"-DSQLITE_HAVE_ZLIB=1",
			"-DSQLITE_LIKE_DOESNT_MATCH_BLOBS",
			"-DSQLITE_SOUNDEX",
			"-DSQLITE_THREADSAFE=1",
			"-DSQLITE_WITHOUT_ZONEMALLOC",
			"-D_LARGEFILE64_SOURCE",
			"-I", ilibc,
			"-I", ilibz,
			"-I", ilibtcl,

			"-extended-errors",
			"-o", result,
			"sqlite3.c",
		)
		switch goos {
		case "linux":
			// nop
		default:
			config = append(config, "-DSQLITE_MUTEX_NOOP")
		}
		switch {
		case win:
			config = append(config,
				"--cpp", xgcc,
				"--goarch", goarch,
				"--goos", goos,
				"-DSQLITE_HAVE_C99_MATH_FUNCS=(1)",
				"-DSQLITE_OS_WIN=1",
				"-DSQLITE_OMIT_SEH",
				"-build-lines", "//go:build windows && (amd64 || arm64)\n",
				"-map", "gcc=x86_64-w64-mingw32-gcc",
			)
		case win32:
			config = append(config,
				"--cpp", xgcc,
				"--goarch", goarch,
				"--goos", goos,
				"-DSQLITE_HAVE_C99_MATH_FUNCS=(1)",
				"-DSQLITE_OS_WIN=1",
				"-DSQLITE_OMIT_SEH",
				"-map", "gcc=i686-w64-mingw32-gcc",
			)
		default:
			config = append(config, "-DSQLITE_OS_UNIX=1")
		}
		switch target {
		case "freebsd/amd64", "freebsd/arm64", "openbsd/amd64", "openbsd/arm64", "netbsd/amd64":
			config = append(config, "-ltcl8.6")
		}
		config = append(config, "-eval-all-macros")
		if err := ccgo.NewTask(goos, goarch, config, os.Stdout, os.Stderr, nil).Main(); err != nil {
			return err
		}

		util.MustShell(true, nil, sed, "-i", `s/\<T__\([a-zA-Z0-9][a-zA-Z0-9_]\+\)/t__\1/g`, result)
		util.MustShell(true, nil, sed, "-i", `s/\<x_\([a-zA-Z0-9][a-zA-Z0-9_]\+\)/X\1/g`, result)
		util.MustShell(true, nil, sed, "-i", `s/func _sqlite3MutexInit(tls \*libc\.TLS) (r int32) {/var mu sync.Mutex; func _sqlite3MutexInit(tls \*libc\.TLS) (r int32) { mu\.Lock(); defer mu\.Unlock();/`, result)
		switch target {
		case "openbsd/amd64", "openbsd/arm64":
			fmt.Printf("sqlite3.c: result=%q", result)
			util.MustShell(true, nil, sed, "-i.bak", `s/\&___stdin\>/libc.Xstdin/g`, result)
			util.MustShell(true, nil, sed, "-i.bak", `s/\&___stdout\>/libc.Xstdout/g`, result)
			util.MustShell(true, nil, sed, "-i.bak", `s/\&___stderr\>/libc.Xstderr/g`, result)
		}
		return nil
	})

	fn := fmt.Sprintf("ccgo_%s_%s.go", goos, goarch)
	if win {
		fn = fmt.Sprintf("ccgo_%s.go", goos)
	}
	mustCopyFile(fn, filepath.Join(makeRoot, result), nil)
	mustCopyFile(filepath.Join("include", "sqlite3.h"), filepath.Join(makeRoot, "sqlite3.h"), nil)
	mustCopyFile(filepath.Join("include", "sqlite3ext.h"), filepath.Join(makeRoot, "sqlite3ext.h"), nil)
	_, extractedArchivePath = filepath.Split(archive2Path)
	extractedArchivePath = extractedArchivePath[:len(extractedArchivePath)-len(".zip")]
	tempDir = os.Getenv("GO_GENERATE_DIR")
	switch {
	case tempDir != "":
		os.RemoveAll(filepath.Join(tempDir, extractedArchivePath))
	default:
		var err error
		if tempDir, err = os.MkdirTemp("", "libsqlite3-generate"); err != nil {
			fail(1, "creating temp dir: %v\n", err)
		}

		defer func() {
			switch os.Getenv("GO_GENERATE_KEEP") {
			case "":
				os.RemoveAll(tempDir)
			default:
				fmt.Printf("%s: temporary directory kept\n", tempDir)
			}
		}()
	}
	libRoot = filepath.Join(tempDir, extractedArchivePath)
	makeRoot = libRoot
	fmt.Fprintf(os.Stderr, "archivePath %s\n", archive2Path)
	fmt.Fprintf(os.Stderr, "extractedArchivePath %s\n", extractedArchivePath)
	fmt.Fprintf(os.Stderr, "tempDir %s\n", tempDir)
	fmt.Fprintf(os.Stderr, "libRoot %s\n", libRoot)
	fmt.Fprintf(os.Stderr, "makeRoot %s\n", makeRoot)
	util.MustShell(true, nil, "unzip", "-q", archive2Path, "-d", tempDir)
	// https://gitlab.com/cznic/sqlite/-/issues/173
	util.MustShell(true, nil, "patch", filepath.Join(makeRoot, "src", "os_unix.c"), filepath.Join("internal", "sqlite_issue173.patch2"))
	// https://gitlab.com/cznic/libsqlite3/-/issues/1
	util.MustShell(true, nil, "patch", filepath.Join(makeRoot, "src", "pcache1.c"), filepath.Join("internal", "issue1.patch2"))
	mustCopyDir(makeRoot, filepath.Join("internal", "overlay", "generator"), nil, false)
	fixWin(tempDir)
	util.MustShell(true, nil, sed, "-i", `s/#if (defined(__GNUC__) || defined(__clang__)) \\/#if 0 \&\& (defined(__GNUC__) || defined(__clang__)) \\/`, filepath.Join(libRoot, "src", "util.c"))
	mustCopyFile("LICENSE-SQLITE.md", filepath.Join(libRoot, "LICENSE.md"), nil)
	util.MustInDir(true, makeRoot, func() (err error) {
		util.MustShell(true, nil, "sh", "-c", `
go mod init example.com/libsqlite3
go get \
	modernc.org/libc@v1.72.2 \
	modernc.org/libtcl8.6@v0.17.4 \
	modernc.org/libz@v0.17.3 \
`)
		if dev {
			util.MustShell(true, nil, "sh", "-c", `
go work init
go work use \
	$GOPATH/src/modernc.org/libc \
	$GOPATH/src/modernc.org/libtcl8.6 \
	$GOPATH/src/modernc.org/libz \
`)
		}
		var config []string
		if m64Double := cc.LongDouble64Flag(goos, goarch); m64Double != "" {
			config = append(config, m64Double)
		}
		config = append(config,
			"-DHAVE_USLEEP",
			"-DLONGDOUBLE_TYPE=double",
			"-DNDEBUG",
			"-DSQLITE_DEFAULT_MEMSTATUS=1",
			"-DSQLITE_DISABLE_INTRINSIC",
			"-DSQLITE_ENABLE_COLUMN_METADATA",
			"-DSQLITE_ENABLE_DBPAGE_VTAB",
			"-DSQLITE_ENABLE_DBSTAT_VTAB",
			"-DSQLITE_ENABLE_EXPLAIN_COMMENTS",
			"-DSQLITE_ENABLE_FTS5",
			"-DSQLITE_ENABLE_GEOPOLY",
			"-DSQLITE_ENABLE_JSON1",
			"-DSQLITE_ENABLE_MATH_FUNCTIONS",
			"-DSQLITE_ENABLE_MEMORY_MANAGEMENT",
			"-DSQLITE_ENABLE_OFFSET_SQL_FUNC",
			"-DSQLITE_ENABLE_PREUPDATE_HOOK",
			"-DSQLITE_ENABLE_RBU",
			"-DSQLITE_ENABLE_RTREE",
			"-DSQLITE_ENABLE_SESSION",
			"-DSQLITE_ENABLE_SNAPSHOT",
			"-DSQLITE_ENABLE_STAT4",
			"-DSQLITE_ENABLE_UNLOCK_NOTIFY",
			"-DSQLITE_LIKE_DOESNT_MATCH_BLOBS",
			"-DSQLITE_SOUNDEX",
			"-DSQLITE_WITHOUT_ZONEMALLOC",
			"-D_LARGEFILE64_SOURCE",
		)
		switch target {
		case
			"linux/amd64":
			config = append(config, "-DSQLITE_THREADSAFE=1")
		default:
			config = append(config, "-DSQLITE_MUTEX_NOOP")
		}
		switch {
		case win:
			config = append(config,
				"-DSQLITE_OS_WIN=1",
			)
			util.MustShell(true, nil, "sh", "-c", fmt.Sprintf("CFLAGS='%s' ./configure --build=x86-64_gnu-linux --host=x86_64-w64-mingw32 --disable-shared --disable-load-extension", strings.Join(config, " ")))
		case win32:
			config = append(config,
				"-DSQLITE_OS_WIN=1",
			)
			util.MustShell(true, nil, "sh", "-c", fmt.Sprintf("CFLAGS='%s' ./configure --build=x86-64_gnu-linux --host=i686-w64-mingw32 --disable-shared --disable-load-extension", strings.Join(config, " ")))
		default:
			config = append(config, "-DSQLITE_OS_UNIX=1", "-DHAVE_MALLOC_USABLE_SIZE=1")
			config = append(config, "-I"+ilibc)
			config = append(config, "-I"+ilibz)
			config = append(config, "-I"+ilibtcl)
			s := ""
			if withTclsh != "" {
				s = fmt.Sprintf("--with-tclsh=%s", withTclsh)
			}
			switch target {
			case "openbsd/amd64", "openbsd/arm64":
				s += "--with-tcl=/usr/local/lib/tcl/tcl8.6"
			}
			util.MustShell(true, nil, "sh", "-c", fmt.Sprintf("CFLAGS='%s' ./configure %s --disable-shared --disable-load-extension", strings.Join(config, " "), s))
			util.MustShell(true, nil, "sh", "-c", "echo '#define HAVE_MALLOC_USABLE_SIZE 1' >> config.h")
			util.MustShell(true, nil, "sh", "-c", "echo '#define HAVE_MEMORY_H 1' >> config.h")
		}
		args := []string{os.Args[0]}
		if dev {
			args = append(args,
				"-absolute-paths",
				"-keep-object-files",
				"-positions",
			)
		}
		args = append(args,
			"--prefix-enumerator=_",
			"--prefix-external=x_",
			"--prefix-field=F",
			"--prefix-macro=m_",
			"--prefix-static-internal=_",
			"--prefix-static-none=_",
			"--prefix-tagged-enum=_",
			"--prefix-tagged-struct=T",
			"--prefix-tagged-union=T",
			"--prefix-typename=T",
			"--prefix-undefined=_",
			"-extended-errors",
			"-ignore-unsupported-alignment",
			"-ignore-link-errors",
		)
		switch target {
		case "freebsd/amd64", "freebsd/arm64", "openbsd/amd64", "openbsd/arm64", "netbsd/amd64":
			args = append(args, "-ltcl8.6")
		}
		switch {
		case win:
			util.MustShell(true, nil, sed, "-i", `s/#if !defined(_WIN32) || defined(__MSVCRT__)/#if !defined(__CCGO__) \&\& (!defined(_WIN32) || defined(__MSVCRT__))/`, "src/test_fs.c")
			ccgo.NewTask(
				goos, goarch,
				append(args,
					"--cpp", xgcc,
					"--goarch", goarch,
					"--goos", goos,
					"-DSQLITE_HAVE_C99_MATH_FUNCS=(1)",
					"-Dmalloc_usable_size(x)=( (int) ( unsigned long long ) ( malloc_usable_size ( x ) ) )",
					"-build-lines", "//go:build windows && (amd64 || arm64)\n// +build windows\n// +build amd64 arm64",
					"-map", "gcc=x86_64-w64-mingw32-gcc",
					"-exec", "make", "-j", j,
					"BEXE=",

					"CFLAGS=-mlong-double-64 -DLONGDOUBLE_TYPE=double -DSQLITE_WITHOUT_ZONEMALLOC -DNDEBUG -DSQLITE_DISABLE_INTRINSIC -DSQLITE_OS_WIN=1 -DSQLITE_OS_UNIX=0 -D_MSC_VER=1 -DSQLITE_OMIT_SEH",
					"TOP=../sqlite-src-"+versionTag,
					"TEXE=.exe",
					"testfixture.exe",
				),
				os.Stdout, os.Stderr, nil,
			).Exec()
			return nil
		case win32:
			util.MustShell(true, nil, sed, "-i", `s/#if !defined(_WIN32) || defined(__MSVCRT__)/#if !defined(__CCGO__) \&\& (!defined(_WIN32) || defined(__MSVCRT__))/`, "src/test_fs.c")
			ccgo.NewTask(
				goos, goarch,
				append(args,
					"--cpp", xgcc,
					"--goarch", goarch,
					"--goos", goos,
					"-DSQLITE_HAVE_C99_MATH_FUNCS=(1)",
					"-Dmalloc_usable_size(x)=( (int) ( unsigned long long ) ( malloc_usable_size ( x ) ) )",
					"-map", "gcc=i686-w64-mingw32-gcc",
					"-exec", "make", "-j", j,
					"BEXE=",

					"CFLAGS=-mlong-double-64 -DLONGDOUBLE_TYPE=double -DNDEBUG -DSQLITE_DISABLE_INTRINSIC -DSQLITE_OS_WIN=1 -DSQLITE_OS_UNIX=0 -D_MSC_VER=1 -DSQLITE_OMIT_SEH",
					"TOP=../sqlite-src-"+versionTag,
					"TEXE=.exe",
					"testfixture.exe",
				),
				os.Stdout, os.Stderr, nil,
			).Exec()
			return nil
		default:
			switch {
			case os.Getenv("GO_GENERATE_TEST") != "":
				args = append(args, "-exec", "make", "-j", j, "fulltestonly")
			default:
				args = append(args, "-exec", "make", "-j", j, "testfixture")
			}
			err := ccgo.NewTask(goos, goarch, args, os.Stdout, os.Stderr, nil).Exec()
			switch target {
			case "darwin/amd64", "darwin/arm64":
				util.Shell(nil, sed, "-i", `/func _guess_number_of_cores(/,/^}/d`, "testfixture.go")
			}
			return err
		}
	})

	os.Mkdir("speedtest1", 0770)
	switch {
	case win:
		if err := ccgo.NewTask(
			goos, goarch,
			[]string{
				os.Args[0],
				"--cpp", xgcc,
				"--goarch", goarch,
				"--goos", goos,
				"-DNDEBUG",
				"-DSQLITE_DISABLE_INTRINSIC",
				"-DSQLITE_ENABLE_DBPAGE_VTAB",
				"-DSQLITE_OMIT_SEH",
				"-DSQLITE_OS_WIN=1",
				"-DSQLITE_ENABLE_RTREE",
				"-I", makeRoot,
				"-build-lines", "//go:build windows && (amd64 || arm64)\n// +build windows\n// +build amd64 arm64",
				"-map", "gcc=x86_64-w64-mingw32-gcc",
				"-o", filepath.Join("speedtest1", fn),
				filepath.Join(makeRoot, "test", "speedtest1.c"),
				"-lsqlite3",
			},
			os.Stdout, os.Stderr,
			nil,
		).Main(); err != nil {
			fail(1, "%s\n", err)
		}
	case win32:
		if err := ccgo.NewTask(
			goos, goarch,
			[]string{
				os.Args[0],
				"--cpp", xgcc,
				"--goarch", goarch,
				"--goos", goos,
				"-DNDEBUG",
				"-DSQLITE_DISABLE_INTRINSIC",
				"-DSQLITE_ENABLE_DBPAGE_VTAB",
				"-DSQLITE_OMIT_SEH",
				"-DSQLITE_OS_WIN=1",
				"-DSQLITE_ENABLE_RTREE",
				"-I", makeRoot,
				"-map", "gcc=i686-w64-mingw32-gcc",
				"-o", filepath.Join("speedtest1", fn),
				filepath.Join(makeRoot, "test", "speedtest1.c"),
				"-lsqlite3",
			},
			os.Stdout, os.Stderr,
			nil,
		).Main(); err != nil {
			fail(1, "%s\n", err)
		}
	default:
		if err := ccgo.NewTask(
			goos, goarch,
			[]string{
				os.Args[0],
				"-DNDEBUG",
				"-DSQLITE_DISABLE_INTRINSIC",
				"-DSQLITE_ENABLE_DBPAGE_VTAB",
				"-DSQLITE_ENABLE_RTREE",
				"-ignore-unsupported-alignment",
				"-ignore-link-errors",
				"-o", filepath.Join("speedtest1", fn),
				"-I", makeRoot,
				filepath.Join(makeRoot, "test", "speedtest1.c"),
				"-lsqlite3",
			},
			os.Stdout, os.Stderr,
			nil,
		).Main(); err != nil {
			fail(1, "%s\n", err)
		}
	}
	switch target {
	case "openbsd/amd64", "openbsd/arm64":
		result := filepath.Join("speedtest1", fn)
		fmt.Printf("speedtest1.c: result=%q", result)
		util.MustShell(true, nil, sed, "-i.bak", `s/\&_stdin\>/libc.Xstdin/g`, result)
		util.MustShell(true, nil, sed, "-i.bak", `s/\&_stdout\>/libc.Xstdout/g`, result)
		util.MustShell(true, nil, sed, "-i.bak", `s/\&_stderr\>/libc.Xstderr/g`, result)
	}
	os.Mkdir("mptest", 0770)
	util.MustShell(true, nil, sed, "-i", `s/strcmp(sqlite3_sourceid()/0 \&\& strcmp(sqlite3_sourceid()/`, filepath.Join(makeRoot, "mptest", "mptest.c"))
	switch {
	case win:
		if err := ccgo.NewTask(
			goos, goarch,
			[]string{
				os.Args[0],
				"--cpp", xgcc,
				"--goarch", goarch,
				"--goos", goos,
				"-DNDEBUG",
				"-DSQLITE_DISABLE_INTRINSIC",
				"-DSQLITE_ENABLE_DBPAGE_VTAB",
				"-DSQLITE_OMIT_SEH",
				"-DSQLITE_OS_WIN=1",
				"-I", makeRoot,
				"-build-lines", "//go:build windows && (amd64 || arm64)\n// +build windows\n// +build amd64 arm64",
				"-map", "gcc=x86_64-w64-mingw32-gcc",
				"-o", filepath.Join("mptest", fn),
				filepath.Join(makeRoot, "mptest", "mptest.c"),
				"-lsqlite3",
			},
			os.Stdout, os.Stderr,
			nil,
		).Main(); err != nil {
			fail(1, "%s\n", err)
		}
	case win32:
		if err := ccgo.NewTask(
			goos, goarch,
			[]string{
				os.Args[0],
				"--cpp", xgcc,
				"--goarch", goarch,
				"--goos", goos,
				"-DNDEBUG",
				"-DSQLITE_DISABLE_INTRINSIC",
				"-DSQLITE_ENABLE_DBPAGE_VTAB",
				"-DSQLITE_OMIT_SEH",
				"-DSQLITE_OS_WIN=1",
				"-I", makeRoot,
				"-map", "gcc=i686-w64-mingw32-gcc",
				"-o", filepath.Join("mptest", fn),
				filepath.Join(makeRoot, "mptest", "mptest.c"),
				"-lsqlite3",
			},
			os.Stdout, os.Stderr,
			nil,
		).Main(); err != nil {
			fail(1, "%s\n", err)
		}
	default:
		if err := ccgo.NewTask(
			goos, goarch,
			[]string{
				os.Args[0],
				"-DNDEBUG",
				"-DSQLITE_DISABLE_INTRINSIC",
				"-I", makeRoot,
				"-ignore-unsupported-alignment",
				"-ignore-link-errors",
				"-o", filepath.Join("mptest", fn),
				filepath.Join(makeRoot, "mptest", "mptest.c"),
				"-lsqlite3",
			},
			os.Stdout, os.Stderr,
			nil,
		).Main(); err != nil {
			fail(1, "%s\n", err)
		}
	}
	switch target {
	case "openbsd/amd64", "openbsd/arm64":
		result := filepath.Join("mptest", fn)
		fmt.Printf("mptest.c: result=%q", result)
		util.MustShell(true, nil, sed, "-i.bak", `s/\&_stdin\>/libc.Xstdin/g`, result)
		util.MustShell(true, nil, sed, "-i.bak", `s/\&_stdout\>/libc.Xstdout/g`, result)
		util.MustShell(true, nil, sed, "-i.bak", `s/\&_stderr\>/libc.Xstderr/g`, result)
	}
	util.MustShell(true, nil, "sh", "-c", fmt.Sprintf("cp %s %s", filepath.Join(makeRoot, "mptest", "*.*test"), "mptest/"))

	os.RemoveAll(filepath.Join("internal", "test"))
	util.MustMkdirs(true, "internal/testfixture", "internal/test")
	switch {
	case win || win32:
		mustCopyFile(filepath.Join("internal", "testfixture", fn), filepath.Join(makeRoot, "testfixture.exe.go"), nil)
	default:
		mustCopyFile(filepath.Join("internal", "testfixture", fn), filepath.Join(makeRoot, "testfixture.go"), nil)
	}
	switch target {
	case "openbsd/amd64", "openbsd/arm64":
		result := filepath.Join("internal", "testfixture", fn)
		fmt.Printf("testfixture: result=%q", result)
		util.MustShell(true, nil, sed, "-i.bak", `s/\&___stdin\>/libc.Xstdin/g`, result)
		util.MustShell(true, nil, sed, "-i.bak", `s/\&___stdout\>/libc.Xstdout/g`, result)
		util.MustShell(true, nil, sed, "-i.bak", `s/\&___stderr\>/libc.Xstderr/g`, result)
	}
	util.MustShell(true, nil, "find", ".", "-name", "*.bak", "-delete")
	util.Shell(nil, "rm", "-rf", filepath.Join("internal", "test"))
	mustCopyDir(filepath.Join("internal", "test"), filepath.Join(makeRoot, "test"), nil, false)
	mustCopyDir("internal/test", "internal/overlay/test", nil, true)
	util.MustShell(true, nil, "go", "test", "-vet=off", "-run", "@")
	util.Shell(nil, "git", "status")
}

func fixWin(dir string) {
	pat := []byte("<Windows.h>")
	repl := []byte("<windows.h>")
	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(path, ".c") {
			return nil
		}

		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		if !bytes.Contains(b, pat) {
			return nil
		}

		return os.WriteFile(path, bytes.ReplaceAll(b, pat, repl), 0660)
	}); err != nil {
		fail(1, "%s\n", err)
	}
}

func mustCopyDir(dst, src string, canOverwrite func(fn string, fi os.FileInfo) bool, srcNotExistsOk bool) (files int, bytes int64) {
	file, bytes, err := copyDir(dst, src, canOverwrite, srcNotExistsOk)
	if err != nil {
		fail(1, "%s\n", err)
	}

	return file, bytes
}

func copyDir(dst, src string, canOverwrite func(fn string, fi os.FileInfo) bool, srcNotExistsOk bool) (files int, bytes int64, rerr error) {
	dst = filepath.FromSlash(dst)
	src = filepath.FromSlash(src)
	si, err := os.Stat(src)
	if err != nil {
		if os.IsNotExist(err) && srcNotExistsOk {
			err = nil
		}
		return 0, 0, err
	}

	if !si.IsDir() {
		return 0, 0, fmt.Errorf("cannot copy a file: %s", src)
	}

	return files, bytes, filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Mode()&os.ModeSymlink != 0 {
			target, err := filepath.EvalSymlinks(path)
			if err != nil {
				return fmt.Errorf("cannot evaluate symlink %s: %v", path, err)
			}

			if info, err = os.Stat(target); err != nil {
				return fmt.Errorf("cannot stat %s: %v", target, err)
			}

			if info.IsDir() {
				rel, err := filepath.Rel(src, path)
				if err != nil {
					return err
				}

				dst2 := filepath.Join(dst, rel)
				if err := os.MkdirAll(dst2, 0770); err != nil {
					return err
				}

				f, b, err := copyDir(dst2, target, canOverwrite, srcNotExistsOk)
				files += f
				bytes += b
				return err
			}

			path = target
		}

		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return os.MkdirAll(filepath.Join(dst, rel), 0770)
		}

		n, err := copyFile(filepath.Join(dst, rel), path, canOverwrite)
		if err != nil {
			return err
		}

		files++
		bytes += n
		return nil
	})
}

func mustCopyFile(dst, src string, canOverwrite func(fn string, fi os.FileInfo) bool) int64 {
	n, err := copyFile(dst, src, canOverwrite)
	if err != nil {
		fail(1, "%s\n", err)
	}

	return n
}

func copyFile(dst, src string, canOverwrite func(fn string, fi os.FileInfo) bool) (n int64, rerr error) {
	src = filepath.FromSlash(src)
	si, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if si.IsDir() {
		return 0, fmt.Errorf("cannot copy a directory: %s", src)
	}

	dst = filepath.FromSlash(dst)
	if si.Size() == 0 {
		return 0, os.Remove(dst)
	}

	dstDir := filepath.Dir(dst)
	di, err := os.Stat(dstDir)
	switch {
	case err != nil:
		if !os.IsNotExist(err) {
			return 0, err
		}

		if err := os.MkdirAll(dstDir, 0770); err != nil {
			return 0, err
		}
	case err == nil:
		if !di.IsDir() {
			return 0, fmt.Errorf("cannot create directory, file exists: %s", dst)
		}
	}

	di, err = os.Stat(dst)
	switch {
	case err != nil && !os.IsNotExist(err):
		return 0, err
	case err == nil:
		if di.IsDir() {
			return 0, fmt.Errorf("cannot overwite a directory: %s", dst)
		}

		if canOverwrite != nil && !canOverwrite(dst, di) {
			return 0, fmt.Errorf("cannot overwite: %s", dst)
		}
	}

	s, err := os.Open(src)
	if err != nil {
		return 0, err
	}

	defer s.Close()
	r := bufio.NewReader(s)

	d, err := os.Create(dst)

	defer func() {
		if err := d.Close(); err != nil && rerr == nil {
			rerr = err
			return
		}

		if err := os.Chmod(dst, si.Mode()); err != nil && rerr == nil {
			rerr = err
			return
		}

		if err := os.Chtimes(dst, si.ModTime(), si.ModTime()); err != nil && rerr == nil {
			rerr = err
			return
		}
	}()

	w := bufio.NewWriter(d)

	defer func() {
		if err := w.Flush(); err != nil && rerr == nil {
			rerr = err
		}
	}()

	return io.Copy(w, r)
}

func env(name, deflt string) (r string) {
	r = deflt
	if s := os.Getenv(name); s != "" {
		r = s
	}
	return r
}
