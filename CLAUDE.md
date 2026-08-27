# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

`modernc.org/libsqlite3` is the **SQLite C amalgamation (currently 3.53.4) mechanically
transpiled to pure Go** with [`ccgo/v4`](https://pkg.go.dev/modernc.org/ccgo/v4), running on
`modernc.org/libc`. It exposes the raw C ABI only — there is no idiomatic Go API here.

There is almost no hand-written logic. The real "source" of this package is the pinned SQLite
zips plus the transpiler configuration in `generator.go`.

Downstream consumers (this repo is a source donor as much as a Go package):
- `modernc.org/sqlite` — its `make vendor` (`vendor_libs/main.go`) reads `../libsqlite3/ccgo_*.go`
  directly from a sibling checkout, rewrites `package libsqlite3` → `package sqlite3`, and writes
  `lib/sqlite_<goos>_<goarch>.go`. Changes here reach users only after that vendoring step.
- `modernc.org/libsqlite_vec` — compiles against `../libsqlite3/include` and imports this module.

## Do not hand-edit the generated files

Machine output, each headed `DO NOT EDIT`; a regeneration silently discards any manual change:

| path | what | size |
|---|---|---|
| `ccgo_<goos>_<goarch>.go` | the library, `package libsqlite3` | ~226k lines / ~9 MB each |
| `internal/testfixture/ccgo_*.go` | the Tcl test harness, `package main` | ~312k lines each |
| `mptest/ccgo_*.go`, `speedtest1/ccgo_*.go` | upstream tools, `package main` | |
| `internal/test/**` | ~1300 files copied from upstream `test/`, wiped and re-copied every generate | |
| `include/sqlite3.h`, `include/sqlite3ext.h` | copied out of the amalgamation | |

Never open a generated file whole — grep for the symbol instead.

To change generated behavior, change an **input**, not the output:
- `generator.go` — ccgo flags, the `sed` post-passes, `versionTag`.
- `internal/*.patch` — applied to the amalgamation's `sqlite3.c`; `internal/*.patch2` — applied to
  the `sqlite-src` tree (`src/os_unix.c`, `src/pcache1.c`).
- `internal/overlay/{generator,test,mptest}/` — files copied over the extracted upstream tree
  (`config.guess`/`config.sub`; four replacement `.test` scripts; `mptest.c`).
- or override a symbol from a **hand-written** Go file.

The hand-maintained files are: `libsqlite3.go` (package doc + the Tier 1/Tier 2 platform table),
`libsqlite3_freebsd.go` / `libsqlite3_windows.go` (libc shims for symbols ccgo can't resolve:
`__inline_isnan*`, `__umulh`), `etc.go` (`origin`/`todo`/`trc` debug helpers),
`rlimit.go` / `rulimit.go` / `norlimit.go` (per-platform `setMaxOpenFiles`), `all_test.go`,
`race_test.go`, `generator.go` (`//go:build ignore`), and
`internal/testfixture/patch_{darwin,freebsd,netbsd,windows}.go`. The darwin one is the canonical
override example: `generator.go` `sed`-deletes `_guess_number_of_cores` from `testfixture.go` and
`patch_darwin.go` supplies a Go replacement using `runtime.GOMAXPROCS`.

## Commands

```sh
make editor              # the fast loop: gofmt -l -s -w . + go test -c + go build ./... + build generator
make all                 # editor + golint + staticcheck (no config, defaults)
make test                # go test -v -timeout 24h — everything; takes hours
make tcltest             # TestTclTest (full permutation) + TestTclTestOFD (lock subset, OFD mode)
make tcltest_ofd         # the full permutation with MODERNC_SQLITE_OFD_LOCK=1 (opt-in OFD locks, linux)
make extraquick          # TestTclTest with the "extraquick" permutation
make locktest            # the 42 lock/WAL Tcl files in both locking modes, ~30 s each
make mptest              # only TestConcurrentProcesses
make mptest_ofd          # only TestConcurrentProcessesOFD (mptest with the OFD locks on)
make speedtest1          # go run ./speedtest1
make build_all_targets   # cross build + test-compile every target, with -tags=none and -tags=dmesg
make work                # go.work over sibling cc/v4, ccgo/v3, ccgo/v4, libc, libtcl8.6, libz
make clean               # log-*, *.test, *.out, go.work*
```

`build_all_targets.sh` sweeps exactly the 20 targets in `builder.json`'s `test` matrix — keep the
two in sync when adding or dropping a platform.

Narrowing the test run (flags are defined in `all_test.go`):

```sh
go test -v -timeout 24h -run TestTclTest -suite=extraquick   # permutation from internal/test/permutations.test
go test -v -run TestTclTest -suite="veryquick fts5*"         # extra words are passed through to permutations.test
go test -v -run 'TestTclTest$' -suite=locks                  # "full" + the 42-file lock/WAL subset (lockTests in all_test.go)
go test -v -run TestTclTest -start=walrestart.test -maxerror=1
go test -v -run TestConcurrentProcesses                      # builds ./mptest, runs crash01/multiwrite01 × journal modes
go test -v -run TestIssueSqlite173                           # re-execs itself with -race -inner
go test -run @                                               # compile-only check (matches nothing)
```

`-suite` defaults to `full`; `-suite=""` runs `all.test` instead of `permutations.test`. Other
flags: `-match=<glob>`, `-verbose=0|1|file`, `-q`, `-strace`, `-xtags=<build tags>` (passed to the
`go build` of testfixture/mptest).

`TestTclTest` builds `./internal/testfixture` at run time and copies the Tcl library out of
`modernc.org/libtcl8.6/library`'s `embed.FS` into a temp `TCL_LIBRARY`.

## Regeneration

Prerequisites: sibling checkouts `../libc`, `../libz`, `../libtcl8.6` (the generator passes their
`include/<goos>/<goarch>` dirs with `-I`); `unzip`, `patch`, `tclsh`, a C toolchain and autotools;
GNU sed installed as **`gsed`** on darwin/freebsd/netbsd/openbsd; mingw
(`x86_64-w64-mingw32-gcc`, `i686-w64-mingw32-gcc`) for the Windows cross-generation.

```sh
make generate   # host target only
make dev        # same + GO_GENERATE_DEV=1, -tags=ccgo.dmesg,ccgo.assert, -absolute-paths
                # -keep-object-files -positions, then greps /tmp/ccgo.log for TRC/TODO/ERRORF/FAIL/undefined
make windows    # cross-generate windows/{amd64,arm64} from linux/amd64 (ccgo_windows.go)
make windows_386
```

`generator.go` runs two phases:
1. **amalgamation zip** → apply `internal/*.patch` + the `sed` fixes to `sqlite3.c` → ccgo →
   `sed` identifier renames → `ccgo_<goos>_<goarch>.go`, `include/sqlite3.h`, `include/sqlite3ext.h`.
2. **sqlite-src zip** → apply `internal/*.patch2` + overlay → `./configure` →
   `ccgo -exec make testfixture` → `internal/testfixture/ccgo_*.go`; then transpile `speedtest1.c`
   and `mptest.c`, copy `mptest/*.test`, and refresh `internal/test` from upstream `test/` plus
   `internal/overlay/test`.

On linux/amd64 `make generate` **also** cross-generates the three Windows targets afterwards
(deferred `make windows windows_386`) unless `GO_GENERATE_NOWIN=1`, and writes the
`internal/autogen/windows_*.mod` snapshots.

Environment: `GO_GENERATE_DIR` (work dir; the Makefile uses `/tmp/libsqlite3`), `GO_GENERATE_DEV`,
`GO_GENERATE_WIN`, `GO_GENERATE_WIN32`, `GO_GENERATE_NOWIN`, `GO_GENERATE_TEST` (builds
`make fulltestonly` instead of `testfixture`), `GO_GENERATE_KEEP`, `TARGET_GOOS`/`TARGET_GOARCH`,
`MODERNC_ORG_SQLITE_WITH_TCLSH`.

### Version bumps — what must stay in sync

- **SQLite**: `versionTag` in `generator.go`, `ZIP`/`ZIP2`/`URL`/`URL2` in the `Makefile`, the
  `download` URLs in `builder.json`, and the version column of the platform table in
  `libsqlite3.go`'s doc comment.
- **`internal/issue1.patch{,2}` and `internal/sqlite_issue173.patch{,2}` are ed-style diffs with
  absolute line numbers** (`55620c55620`), not context diffs. They break on any upstream line-number
  shift, so every SQLite bump means regenerating them against the new sources.
  `internal/sqlite_issue255.patch{,2}` (OFD locking, below) are unified diffs made with `diff -u`
  from a before/after tree - "before" = pristine `sqlite3.c` with `sqlite_issue173.patch` and
  `issue1.patch` applied, resp. pristine `src/os_unix.c` with `sqlite_issue173.patch2` - and
  apply at zero offset; regenerate them the same way rather than editing hunks by hand.
- **libc / libz / libtcl8.6**: hard-coded in the two `go get` lines inside `generator.go` and must
  match `go.mod` (cf. the `generator.go: update libc version` commits).

## Per-target generation and the builder farm

Every target except Windows must be generated **on a matching machine**, which is why the history
is a stream of `<host> auto generate` commits (`nuc64`, `pi64`, `pi32`, `darwin-m1`, `s390x`,
`riscv64`, `loong64b`, …). That fleet is driven by `builder.json` — the `autogen` / `test` /
`autotag` target regexes and the zips to download — and reports to
<https://modern-c.appspot.com/-/builder/?importpath=modernc.org%2flibsqlite3>.

`internal/autogen/<goos>_<goarch>.mod` is a snapshot of `go.mod` as of that target's last
successful generation. The farm regenerates a target when the snapshot differs from `go.mod`; so
bumping a dependency triggers a sweep implicitly, and
`for f in internal/autogen/*.mod; do echo > $f; done` forces one explicitly.

Tier 1 vs Tier 2 is documented in `libsqlite3.go`: openbsd/{amd64,arm64} is Tier 2, which
guarantees only that the package builds and that some tests pass — bugs there don't block a
release, and the package doc warns against production use.

## Generated API conventions

Set by ccgo's `--prefix-*` flags plus the `sed` renames in `generator.go`:

- **Functions**: ccgo emits `x_sqlite3_open`, a `sed` pass rewrites `x_` → `X`, giving
  `Xsqlite3_open`, `Xsqlite3_prepare_v2`, … (~367 exported).
- **Macros keep their C names** as plain Go constants: `SQLITE_OK`, `SQLITE_OPEN_READWRITE`
  (`-eval-all-macros`, no `--prefix-macro` in phase 1). This is the API `modernc.org/sqlite`'s
  `lib` package re-exports.
- **Types** `T`-prefixed (`Tsqlite3`, `Tsqlite3_stmt`), **struct fields** `F`-prefixed,
  statics/internals/enums `_`-prefixed (`_sqlite3MutexInit`).
- Every function takes `tls *libc.TLS` first; all pointers are `uintptr`.
- `internal/testfixture` compiles `sqlite3.c` into itself and therefore keeps the raw `x_`
  prefixes and `m_`-prefixed macros — the phase-1 renames don't apply to it. `mptest` and
  `speedtest1` are `package main` too but *import* `modernc.org/libsqlite3` (`m_` macros, `X`
  library calls), so they exercise the library file, not a private copy — which is why the
  darwin OFD panic of 2026-08-27 showed up in `TestConcurrentProcesses`.
- Build constraints are `//go:build <goos> && <goarch>`, except `ccgo_windows.go`, which is
  `windows && (amd64 || arm64)`.

## Upstream fixes baked into generation

None at present. Historical note, because it explains older builder logs: SQLite 3.53.3 had a
super-journal regression — `readSuperJournal()` returned a non-NULL pointer to an *empty* name
when a crash zeroed the name (its byte-sum checksum still validates) and `pager_playback()`
tested the pointer rather than `zSuper[0]`, so the hot journal was deleted without being rolled
back and the database was left corrupted. It made `crash.test` fail ~2% of runs on every
platform ("file is not a database" / "database disk image is malformed"). Reported from here
(<https://sqlite.org/forum/info/2026-07-20T18:27:00Z>) and carried as
`internal/sqlite_superjournal.patch{,2}` until **3.53.4 shipped the identical one-line fix**
upstream (check-in `bf70dadc2d455844`), which is when the patch was dropped. Upstream's
regression test for it is `test/crash9.test`, new in 3.53.4 and part of the `full` suite.

## OFD locking (opt-in) baked into generation

`internal/sqlite_issue255.patch{,2}` (cznic/sqlite#255, MR !3 + fixups): on linux the `unix` VFS can
lock database files with Open File Description locks (`F_OFD_SETLK*`) instead of POSIX record locks,
which any `close()` of an unrelated descriptor of the file silently drops. Facts to keep straight:

- **Off by default.** `unixOfdActive()` = `ofdEnabled && ofdSupported`. `ofdEnabled` is settled once,
  by `ofdInitOnce()` from `sqlite3_os_init()` reading `MODERNC_SQLITE_OFD_LOCK` (non-empty, not
  starting with `0`), or by `modernc_ofd_locking(int)` → `Xmodernc_ofd_locking`, which overrides the
  environment, returns the previous value (`-1` = unavailable; a negative argument only queries) and
  must be called before the first database file is opened. Process-wide by kernel fact: POSIX and
  OFD locks from the same process are different owners and conflict, so a per-DSN switch would mean
  mixed-mode inodes.
- **Guard is `defined(F_OFD_SETLK) && defined(__linux__)`, both halves needed.** macOS has had
  `F_OFD_*` since 10.13 and darwin transpiles against the host SDK, so `F_OFD_SETLK` alone pulled the
  whole path into `ccgo_darwin_*.go`, where libc's `Xfcntl64` panics on cmd 90 (2026-08-27 darwin-m1
  red). The `#else` branch is upstream's locking plus the `-1` returning setter, on every unix
  target; windows has no `os_unix.c` and no setter. Check after a sweep: `_ofdSupported` appears in
  the eight `ccgo_linux_*.go` only; `Xmodernc_ofd_locking` in every non-windows `ccgo_*.go`.
- **One designated descriptor per inode** (`pInode->hLock`, first locker's; a read-write connection
  joining at SHARED takes over from a read-only one) because `unixInodeInfo` assumes one lock owner
  per inode and process. The `hLock` bookkeeping runs on every unix target but is consulted only
  through `unixLockFd()`, which ignores it unless `unixOfdActive()`.
- **EINVAL fallback is latched** (`ofdConfirmed`): only the very first OFD `fcntl()` may flip the
  process to POSIX mode; a later EINVAL is returned as an I/O error, since a POSIX `F_UNLCK` cannot
  release an OFD lock already held.
- **Acceptance gate:** `make locktest` (42 lock/WAL Tcl files, both modes, ~30 s each on a desktop),
  `make mptest_ofd`, and the scenario program from the MR !3 review (`escape`, `exclro`, `stale`,
  `stale2`, `rorw`, `interleave`, `rogue`, `excl`, `pending`) in default, env-var, setter and
  forced-POSIX mode. On the farm, `TestTclTestOFD` (the subset) and `TestConcurrentProcessesOFD`
  (mptest) run with the switch on, linux only, `t.Setenv` — the second mode costs ~1.3% of the Tcl
  run plus one more mptest, which is why the full permutation is *not* run twice there
  (`make tcltest_ofd` does that by hand). Only the subset can tell the modes apart: everything
  else opens one connection, takes SHARED and never sees a difference.

## Race/threading fixes baked into generation

Deliberate and load-bearing — don't "clean them up" out of `generator.go` or the patches:

- `_sqlite3MutexInit` is wrapped in a package-level `sync.Mutex` (hence ccgo's `-import=sync`).
- **Every** `int isInit` becomes `volatile int isInit` — ccgo turns a `volatile` read into
  `libc.AtomicLoadPInt32`, so this is real codegen, not a cosmetic marker. The `sed` is written
  `0,/int isInit;*True after/{…}` as if to hit only the first declaration, but that address regex
  matches nothing (`;*` is zero-or-more semicolons, and the real line has ` /* True after…` in
  between), so the range runs to EOF and all four declarations are covered. Narrowing it to the
  first would *remove* atomic loads from the other three — a behaviour change needing a full
  regeneration sweep, so leave it unless you mean it. `bUnderPressure` and `randomnessPid` get the
  same `volatile` treatment via `internal/issue1.patch{,2}` and `internal/sqlite_issue173.patch{,2}`.
- The `#if (defined(__GNUC__) || defined(__clang__))` intrinsics block is disabled (`#if 0 && …`) in
  both `sqlite3.c` and `src/util.c`.
- `-DSQLITE_THREADSAFE=1` only on linux; every other target gets `-DSQLITE_MUTEX_NOOP`.
- `-DLONGDOUBLE_TYPE=double` plus ccgo's `-mlong-double-64`.
- `race_test.go`'s `TestIssueSqlite173` is the regression guard for the `unixRandomness` race; it
  re-runs itself under `-race` and tolerates "`-race` is not supported" / "unsupported VMA range".

## Known-failure bookkeeping

Before chasing a Tcl failure, check the tables at the top of `all_test.go`:

- `expectedFailures` — `dbstatus-4.*`, `malloc5-6.2.*`, `values-11.*`; memory accounting differs
  from C because escaped locals share the heap and TLS stacks live until `TLS.Close`. Plus 18
  `dbstatus-2.*` entries added for windows in `init()`.
- `knownCFailures` — fails in upstream C too: `snapshot_fault-4.1.1` (linux/ppc64le).
- `like-14.{1,2}` is **not** in either table any more, don't put it back. It is a **1000 µs**
  single-shot timing assertion (the test prints "ms", but Tcl's `[time]` reports microseconds) on
  one GLOB/LIKE round trip through the Tcl binding, prepare included — a coin toss on the emulated
  builders (freebsd/arm ~1096 µs; linux/loong64 300–1500 µs run to run; s390x 564 µs alone but over
  the limit inside the ~20 h full run), which is why it used to be waived for freebsd/arm and
  linux/s390x. Since 2026-08-23 `generator.go` builds testfixture with
  `-DCONFIG_SLOWDOWN_FACTOR=10.0` (upstream's knob for slow builds, → `$sqlite_options(configslower)`
  → a 10 ms limit). It takes effect per target as its `internal/testfixture/ccgo_*.go` is
  regenerated, so a `like-14` red means that target's testfixture predates it: blank its
  `internal/autogen/<goos>_<goarch>.mod` to regenerate it.
- per-target blacklists in `TestTclTest` — `bigsort.test`
  (linux/{arm64,loong64,riscv64,ppc64le,s390x}), `symlink2.test`/`readonly.test`/`snapshot3.test`
  on windows; `TestConcurrentProcesses` is skipped on linux/s390x (VM too slow).
- `bigsort.test` on linux/s390x is **not** a RAM shortage, whatever the builder log says.
  `modernc.org/memory` mmaps one VMA per 64 KB page, so the ~6.7 GB of C heap the test needs
  (`PRAGMA cache_size = 4194304`, ie. a 4 GiB page cache) blows past the default `vm.max_map_count`
  of 65530 at ~4.1 GB. Past the ceiling the kernel still extends existing VMAs, so libc keeps
  allocating, but the next `mmap` needing a *fresh* one fails — either the Go runtime's heap-arena
  metadata (`fatal error: out of memory allocating heap arena metadata`, process dies, builder
  reports `final summary not detected`) or SQLite's allocator (`SQLITE_NOMEM`, clean
  `bigsort-1.1 FAIL`), depending on who asks first. Both signatures were reproduced 2026-08-18;
  death came at 8.1 GB RSS with 23.8 GB free on the 32 GB box. `sysctl -w vm.max_map_count=262144`
  on the host is the alternative to the blacklist entry.
- `setMaxOpenFiles(1024)` runs before the Tcl suite to keep `misc7.test` from hanging.
- The copied test corpus is re-stamped with the current time right after `util.CopyDir` — **do not
  drop that walk**. `CopyDir` preserves the checkout's mtime *and* atime, so the ~1285 files landed
  under `$TMPDIR` looking weeks stale, and the OS temp reapers deleted the not-yet-sourced ones
  mid-run: OpenBSD `/etc/daily` (`find -x /tmp -type f … -atime +7 -delete`) and Windows
  `SilentCleanup` (`VolumeCaches\Temporary Files`, `LastAccess` 7 days, fired by low free disk).
  The symptom was a *different* `couldn't read file ".../<X>.test": no such file or directory`
  every run — win32, pi400 (windows/arm64) and openbsd/arm64, because a run has to straddle a
  cleanup pass; the fast builders finish first.
