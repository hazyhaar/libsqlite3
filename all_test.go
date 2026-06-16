// Copyright 2023 The libsqlite3-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libsqlite3 // import "modernc.org/libsqlite3"

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"

	_ "modernc.org/ccgo/v4/lib"
	util "modernc.org/fileutil/ccgo"
	"modernc.org/libtcl8.6/library"
)

var (
	oInner    = flag.Bool("inner", false, "internal use")
	oMatch    = flag.String("match", "", "pattern match for tests")
	oMaxError = flag.Uint("maxerror", 0, "stop after <uint> errors")
	oQuiet    = flag.Bool("q", true, "reduce output")
	oStrace   = flag.Bool("strace", false, "strace TclTest")
	oStart    = flag.String("start", "", "-start=[$permutation:]$testfile")
	oSuite    = flag.String("suite", "full", "suite [test-file] to run")
	oVerbose  = flag.String("verbose", "0", `"0", "1" or "file"`)
	oXTags    = flag.String("xtags", "", "passed as -tags to go build of testfixture")

	expectedFailures = map[string]struct{}{
		// Exact memory accounting of the transpilled-to-Go code differs from the
		// original C code because we additionally allocate space for escaped local
		// variables while using the same heap. We currently keep the TLS stacks,
		// because of tcl8.6 implementation of closures, and do not free them until
		// TLS.Close. See libc.TLS.Alloc/Free for more details.
		"dbstatus-4.0.1": {},
		"dbstatus-4.1.1": {},
		"dbstatus-4.2.1": {},
		"dbstatus-4.2.2": {},
		"dbstatus-4.2.3": {},
		"dbstatus-4.2.4": {},
		"malloc5-6.2.2":  {},
		"malloc5-6.2.3":  {},

		// 2024-05-25: Reported at https://sqlite.org/forum/forumpost/8caab936c9
		"values-11.0": {},
		"values-11.1": {},
	}
	knownCFailures = map[string]struct{}{}

	goos   = runtime.GOOS
	goarch = runtime.GOARCH
	target = fmt.Sprintf("%s/%s", goos, goarch)
)

func init() {
	if goos == "windows" {
		for _, v := range []string{
			"dbstatus-2.0-1.a",
			"dbstatus-2.0-2.a",
			"dbstatus-2.0-3.a",
			"dbstatus-2.0-4.a",
			"dbstatus-2.0-5.a",
			"dbstatus-2.0-7.a",
			"dbstatus-2.64-1.a",
			"dbstatus-2.64-2.a",
			"dbstatus-2.64-3.a",
			"dbstatus-2.64-4.a",
			"dbstatus-2.64-5.a",
			"dbstatus-2.64-7.a",
			"dbstatus-2.120-1.a",
			"dbstatus-2.120-2.a",
			"dbstatus-2.120-3.a",
			"dbstatus-2.120-4.a",
			"dbstatus-2.120-5.a",
			"dbstatus-2.120-7.a",
		} {
			expectedFailures[v] = struct{}{}
		}
	}
}

func TestMain(m *testing.M) {
	rc := m.Run()
	os.Exit(rc)
}

func TestConcurrentProcesses(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	switch target {
	case "linux/s390x":
		t.Skip("TODO") // VM too slow.
	}

	dir, err := os.MkdirTemp("", "sqlite-test-")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		os.RemoveAll(dir)
	}()

	m, err := filepath.Glob(filepath.FromSlash("mptest/*"))
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range m {
		if s := filepath.Ext(v); s != ".test" && s != ".subtest" {
			continue
		}

		b, err := os.ReadFile(v)
		if err != nil {
			t.Fatal(err)
		}

		if runtime.GOOS == "windows" {
			// reference tests are in *nix format -- but git on windows does line-ending
			// xlation by default if someone has it 'off' this has no impact.
			//
			// '\r\n'  -->  '\n'
			b = bytes.ReplaceAll(b, []byte("\r\n"), []byte("\n"))
		}

		if err := os.WriteFile(filepath.Join(dir, filepath.Base(v)), b, 0666); err != nil {
			t.Fatal(err)
		}
	}

	bin := "./mptest"
	if runtime.GOOS == "windows" {
		bin += "mptest.exe"
	}
	args := []string{"build", "-o", filepath.Join(dir, bin)}
	if s := *oXTags; s != "" {
		args = append(args, "-tags", s)
	}
	args = append(args, "modernc.org/libsqlite3/mptest")
	out, err := exec.Command("go", args...).CombinedOutput()
	if err != nil {
		t.Fatalf("%s\n%v", out, err)
	}

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	defer os.Chdir(wd)

	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}

	//	mptester$(TEXE):	sqlite3.lo $(TOP)/mptest/mptest.c
	//		$(LTLINK) -o $@ -I. $(TOP)/mptest/mptest.c sqlite3.lo \
	//			$(TLIBS) -rpath "$(libdir)"
	//
	//	MPTEST1=./mptester$(TEXE) mptest.db $(TOP)/mptest/crash01.test --repeat 20
	//	MPTEST2=./mptester$(TEXE) mptest.db $(TOP)/mptest/multiwrite01.test --repeat 20
	//	mptest:	mptester$(TEXE)
	//		rm -f mptest.db
	//		$(MPTEST1) --journalmode DELETE
	//		$(MPTEST2) --journalmode WAL
	//		$(MPTEST1) --journalmode WAL
	//		$(MPTEST2) --journalmode PERSIST
	//		$(MPTEST1) --journalmode PERSIST
	//		$(MPTEST2) --journalmode TRUNCATE
	//		$(MPTEST1) --journalmode TRUNCATE
	//		$(MPTEST2) --journalmode DELETE

	mptest1 := []string{bin, "mptest.db", "crash01.test", "--repeat", "20", "--timeout", "120000"}
	mptest2 := []string{bin, "mptest.db", "multiwrite01.test", "--repeat", "20", "--timeout", "120000"}
	mptest1 = mptest1[:len(mptest1):len(mptest1)]
	mptest2 = mptest2[:len(mptest2):len(mptest2)]
	os.Remove("mptest.db")
outer:
	for _, args := range [][]string{
		append(mptest1, "--journalmode", "DELETE"),
		append(mptest2, "--journalmode", "WAL"),
		append(mptest1, "--journalmode", "WAL"),
		append(mptest2, "--journalmode", "PERSIST"),
		append(mptest1, "--journalmode", "PERSIST"),
		append(mptest2, "--journalmode", "TRUNCATE"),
		append(mptest1, "--journalmode", "TRUNCATE"),
		append(mptest2, "--journalmode", "DELETE"),
	} {
		fmt.Printf("execute %v\n", args)
		out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
		if err != nil {
			t.Fatalf("FAIL err=%v out=%s", err, out)
		}

		// just remove it so we don't get a file busy race-condition when we spin up
		// the next script
		if runtime.GOOS == "windows" {
			os.Remove("mptest.db")
		}

		a := strings.Split(string(out), "\n")
		for _, v := range a {
			if strings.HasPrefix(v, "Summary:") {
				b := strings.Fields(v)
				if len(b) < 2 {
					t.Fatalf("unexpected format of %q", v)
				}

				n, err := strconv.Atoi(b[1])
				if err != nil {
					t.Fatalf("unexpected format of %q", v)
				}

				if n != 0 {
					t.Errorf("%s", out)
				}

				t.Logf("%v", v)
				continue outer
			}

		}
		t.Fatalf("%s\nerror: summary line not found", out)
	}
}

var blacklist = []string{}

func TestTclTest(t *testing.T) {
	blacklist := map[string]struct{}{}
	switch runtime.GOOS {
	case "windows":
		// See https://gitlab.com/cznic/sqlite/-/issues/23#note_599920077 for details.
		blacklist["symlink2.test"] = struct{}{}
		// This test should not nun on windows but the platform test fails.
		blacklist["readonly.test"] = struct{}{}
		//TODO
		blacklist["snapshot3.test"] = struct{}{}
	}
	switch target {
	case
		"linux/arm64",   // OOM killed on rpi5
		"linux/loong64", // OOM killed on loong64b
		"linux/riscv64": // hangs

		// # This test causes thrashing on machines with smaller amounts of
		// # memory.  Make sure the host has at least 8GB available before running
		// # this test.
		blacklist["bigsort.test"] = struct{}{}
	case "linux/ppc64le":
		blacklist["bigsort.test"] = struct{}{} // OOM killed on ppc64le
		knownCFailures["snapshot_fault-4.1.1"] = struct{}{}
	case "freebsd/arm":
		// like-14.2 asserts a LIKE-optimization query completes in under 1s.
		// On the emulated 32-bit arm builder the query is correct but slow
		// (~1096ms > 1000ms threshold), so the timing assertion fails. Not a
		// correctness issue.
		knownCFailures["like-14.2"] = struct{}{}
	}

	if err := setMaxOpenFiles(1024); err != nil { // Avoid misc7.test hanging for a long time.
		t.Fatal(err)
	}

	tmpDir := t.TempDir()
	fmt.Printf("tmpDir=%v\n", tmpDir)
	tclLibrary := filepath.Join(tmpDir, "tcllib")
	if err := os.MkdirAll(tclLibrary, 0770); err != nil {
		t.Fatal(err)
	}

	if err := cp(library.FS, "assets", tclLibrary); err != nil {
		t.Fatal(err)
	}

	os.Setenv("TCL_LIBRARY", tclLibrary)
	os.Setenv("PATH", fmt.Sprintf("%s%c%s", tmpDir, os.PathListSeparator, os.Getenv("PATH")))
	bin := filepath.Join(tmpDir, "testfixture")
	if goos == "windows" {
		bin += ".exe"
	}
	if out, err := util.Shell(nil, "go", "build", "-o", bin, "-tags="+*oXTags, fmt.Sprintf(".%c%s", os.PathSeparator, filepath.Join("internal", "testfixture"))); err != nil {
		t.Fatalf("%s\nFAIL: %v", out, err)
	}

	testsSrc := filepath.Join("internal", "test")
	tests := filepath.Join(tmpDir, "test")
	if err := os.Mkdir(tests, 0770); err != nil {
		t.Fatal(err)
	}

	if _, _, err := util.CopyDir(tests, testsSrc, nil); err != nil {
		t.Fatal(err)
	}

	for k := range blacklist {
		if err := os.Remove(filepath.Join(tests, k)); err != nil {
			t.Fatal(err)
		}
	}

	tests, err := filepath.Abs(tests)
	if err != nil {
		t.Fatal(err)
	}

	var args []string
	switch s := *oSuite; s {
	case "":
		args = []string{filepath.Join(tests, "all.test")}
		trc("%q", args)
	default:
		a := strings.Split(s, " ")
		args = append([]string{filepath.Join(tests, "permutations.test")}, a...)
		trc("%q", args)
	}
	if *oStart != "" {
		args = append(args, fmt.Sprintf("-start=%s", *oStart))
	}
	if *oMatch != "" {
		args = append(args, fmt.Sprintf("-match=%s", *oMatch))
	}
	if *oVerbose != "" {
		args = append(args, fmt.Sprintf("-verbose=%s", *oVerbose))
	}
	if *oVerbose == "" {
		if *oQuiet {
			args = append(args, "-q")
		}
	}
	if *oMaxError != 0 {
		args = append(args, fmt.Sprintf("-maxerror=%v", *oMaxError))
	}
	var out []byte
	util.InDir(tmpDir, func() error {
		bin := filepath.Base(bin)
		switch {
		case *oStrace:
			out, err = util.Shell(nil, "strace", append([]string{"-f", "-r", bin}, args...)...)
		default:
			out, err = util.Shell(nil, bin, args...)
		}
		if err != nil {
			switch err.Error() {
			case "exit status 1":
				t.Logf("fail: %v\n%s", err, out)
			default:
				t.Errorf("fail: %v\n%s", err, out)
			}
		}
		return nil
	})
	s := string(out)
	const (
		tagFailures    = "!Failures on these tests: "
		tagMemoryUsage = "Current memory usage: "
		tagSIGABRT     = "SIGABRT"
	)
	if strings.Contains(s, tagSIGABRT) {
		t.Errorf("SIGABRT: test crashed")
		return
	}

	if !strings.Contains(s, tagMemoryUsage) {
		t.Errorf("final summary not detected (test crashed?)")
		return
	}

	x := strings.Index(s, tagFailures)
	if x < 0 {
		return
	}

	s = s[x+len(tagFailures):]
	s = strings.TrimSpace(s[:strings.IndexByte(s, '\n')])
	a := strings.Fields(s)
	for _, v := range a {
		if _, ok := expectedFailures[v]; ok {
			t.Logf("%s: expected fail", v)
			continue
		}

		if _, ok := knownCFailures[v]; ok {
			t.Logf("%s: fails in C", v)
			continue
		}

		t.Errorf("%s FAIL", v)
	}
}

func cp(fsys embed.FS, rootDir, destDir string) (err error) {
	return fs.WalkDir(fsys, rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			if path = path[len(rootDir):]; path != "" {
				path = filepath.Join(destDir, path)
				err = os.MkdirAll(path, 0770)
			}
			return err
		}

		b, err := fs.ReadFile(fsys, path)
		if err != nil {
			return err
		}

		path = filepath.Join(destDir, path[len(rootDir)+1:])
		err = os.WriteFile(path, b, 0660)
		return err
	})
}
