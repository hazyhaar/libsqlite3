# Handoff: emulate SQLite's Windows SEH in the ccgo build (cznic/sqlite#221)

**From:** the `modernc.org/sqlite` side, 2026-09-03, in reply to GitHub PR modernc-org/sqlite#7.
**Re:** branch `seh-emulation` of this repo: `internal/sqlite_issue221.patch`, `generator.go`,
`seh.go`, `seh_test.go`, `seh_unix_test.go`, `seh_windows_test.go`, `CHANGELOG.md`, plus regenerated
`ccgo_linux_amd64.go`, `ccgo_windows.go` and `ccgo_windows_386.go` for testing (a separate commit
the farm's autogen supersedes). Its counterpart in `modernc.org/sqlite` is
the branch `seh-emulation` there: `lib/seh.go`, `seh_test.go`, `CHANGELOG.md`.
**Status of the farm at the time of writing:** `origin/master` = `7f08a480` (the libc v1.75.7 bump,
picking up cznic/libc!34 "optimize for performance") is being tested by the builders for a release.
Nothing on either branch is merged or pushed, and nothing on them touches `go.mod`, so nothing
here can start a farm-wide regeneration by accident.

**Sequence agreed with cznic:** wait for the libsqlite3 release, re-vendor and tag `modernc.org/sqlite`,
*then* merge this branch (§5). The branches exist so that nothing is lost in between.

---

## 0. TL;DR

- Since SQLite 3.43 (default since 3.44) MSVC builds wrap nine WAL entry points in `__try/__except`
  (`SQLITE_USE_SEH`): a Windows *in-page error* while touching the memory-mapped `-shm` file is
  turned into `SQLITE_IOERR_IN_PAGE` (8714) after SQLite's own cleanup, and the connection stays
  usable. Every non-MSVC build, ours included, dies instead. That is #221: `unexpected fault address
  0x20e69690000 [signal 0xc0000006 ...]` in `libc.Xmemcpy` under `_walIndexTryHdr`.
- Go cannot express `__try`, but the Go runtime already converts exactly that fault
  (`EXCEPTION_IN_PAGE_ERROR`, and `SIGBUS`/`SIGSEGV` on unix) into a recoverable panic carrying the
  fault address whenever `debug.SetPanicOnFault(true)` is in effect for the goroutine. The message in
  #221 is the runtime's *crash* branch of the same code path (`runtime.sigpanic`), taken because the
  flag was off.
- So the fix is a C patch plus one hand-written Go file, not a ccgo feature: the patch enables
  `SQLITE_USE_SEH` under `__CCGO__` and rewrites the nine `SEH_TRY{...}SEH_EXCEPT(...)` blocks as calls
  through `modernc_seh_try()`; `seh.go` implements it with `SetPanicOnFault` + `recover`, runs SQLite's
  `walHandleException()` on a caught fault and returns `SQLITE_IOERR_IN_PAGE`.
- Verified on this machine (linux/amd64, §4): the patch applies and transpiles; a real `SIGBUS` from a
  truncated `-shm` mapping becomes `SQLITE_IOERR_IN_PAGE` and the connection recovers; 64 injected
  faults across read, write, checkpoint and savepoint paths are all handled; the trampoline's
  address filter and guard restoration work; the guard costs ~14 ns per protected call.
- Not a behaviour change unless a fault happens. No API change. No opt-in needed; an opt-out is a
  one-liner if wanted (§3.5).

## 1. The problem and what SEH buys

`wal.c` reads and writes the wal-index through `pWal->apWiData[]`, pointers into the `xShmMap`
mapping of the `-shm` file (32 KiB pages, `WALINDEX_PGSZ`). When Windows cannot page a mapped file
in - filter drivers, AV/EDR, network or deduplicated volumes, a snapshot in progress - the access
raises `STATUS_IN_PAGE_ERROR` (0xC0000006). The mapping is not optional in WAL mode, unlike
`PRAGMA mmap_size` for the database file, which is why upstream added SEH for it and nothing else:

- 3.43.0 (2023-08-24): `SQLITE_USE_SEH` compile-time option.
- 3.44.0 (2023-11-01): enabled by default when built with MSVC.

`sqliteInt.h` hard-codes the condition (`#if defined(_MSC_VER) && !defined(SQLITE_OMIT_SEH)`), so a
mingw build - ours is preprocessed by `x86_64-w64-mingw32-gcc`, and so is mattn/go-sqlite3 on
Windows - never had it; our `-DSQLITE_OMIT_SEH` in `generator.go` was belt and braces (still needed
for the testfixture, which is built with `-D_MSC_VER=1`).

What the MSVC build does on a fault, and what we now do too:

1. `sehExceptionFilter()` accepts `EXCEPTION_IN_PAGE_ERROR` only.
2. The handler runs `walHandleException()` (or just sets `rc = SQLITE_IOERR_IN_PAGE` at four sites):
   it releases the transient locks recorded in `pWal->lockMask` (keeping the ones `readLock`,
   `writeLock`, `ckptLock` say are owned), frees `pWal->pFree`, restores `pWal->apWiData[iWiPg]` from
   `pWal->pWiValue`, and returns `SQLITE_IOERR_IN_PAGE`.
3. The statement fails with "disk I/O error"; the pager goes through its normal error path
   (`pager_error` → rollback → `pager_unlock` clears `errCode`) and the connection stays usable.
   `sqlite3_system_errno()` reports the NTSTATUS via `pWal->iSysErrno` in the MSVC build; we leave
   it 0, Go does not hand the NTSTATUS out.

## 2. Why PR #7 was declined

Full text: https://github.com/modernc-org/sqlite/pull/7#issuecomment-5523699915. In short: the PR is a
`lib/`-only re-vendor from a private libsqlite3 tree (no C source, unreproducible, would be reverted
by the next `make vendor`); it silently adds `-DSQLITE_MAX_MMAP_SIZE=0` for Windows (kills
`PRAGMA mmap_size`, a backward-compat break); and its design - per-process heap copies of the
`-shm` synced by `ReadFile`/`WriteFile` at lock transitions - breaks the shared-memory assumptions
of wal.c across processes (a writer's commit flush clobbers other processes' read marks, the header
is published before the hash pages it describes, a reset header stays unflushed until the write
unlock) and makes every read transaction re-read the whole wal-index. Its non-Windows targets were
content-identical to master, and the Windows targets built. The contributor may resubmit along the
lines of this branch.

## 3. Design

### 3.1 Mechanism (Go runtime facts, go1.27 sources)

- `runtime/signal_windows.go`, `isgoexception()`: `EXCEPTION_ACCESS_VIOLATION` and
  `EXCEPTION_IN_PAGE_ERROR` raised while the PC is inside Go text are handled by Go and delivered to
  the goroutine as a call to `sigpanic()`.
- `sigpanic()`: for those two codes, `sigcode1 < 0x1000` → `panicmem()` (nil dereference);
  otherwise `if gp.paniconfault { panicmemAddr(gp.sigcode1) }` else
  `print("unexpected fault address ", ...); throw("fault")`. The latter is the exact text in #221.
- `panicmemAddr` panics with `runtime.errorAddressString`, a `runtime.Error` with `Addr() uintptr`.
- `debug.SetPanicOnFault(bool) bool` sets/returns the per-goroutine flag. SQLite runs on the
  caller's goroutine, so arming it around the protected body is sufficient.
- unix (`signal_unix.go`) does the same for `SIGBUS`/`SIGSEGV` with a non-nil address; the classic
  producer is a `MAP_SHARED` mapping whose file was truncated, which is what the tests use.
- Unwinding: the panic runs deferred calls in the frames it unwinds. Transpiled frames only
  `defer tls.Free(n)` (verified: no non-deferred `tls.Free` in the generated code), so the libc TLS
  stack stays balanced. The frames between the fault and the trampoline hold no mutexes: the
  protected regions only take `xShmLock` locks, which `walHandleException()` accounts for.
- `recover()` only works in a deferred function of a frame that is still on the stack, i.e. the
  guard must be a Go function *called from* the protected region's caller. Hence the C-side thunks.
- libc's `Xsetjmp` returns 0 and `Xlongjmp` panics `TODO`: C alone cannot express a non-local exit
  under ccgo, which rules out any pure-C emulation.

### 3.2 The C side: `internal/sqlite_issue221.patch`

Applied by `generator.go` after `sqlite_issue173.patch`, `issue1.patch` and `sqlite_issue255.patch`
(the patch was diffed against that state; `patch` applies it without fuzz). 17 hunks:

1. `sqliteInt.h`: `#if (defined(_MSC_VER) || defined(__CCGO__)) && !defined(SQLITE_OMIT_SEH)`.
   ccgo predefines `__CCGO__` (`ccgo/v4/lib/ccgo.go`). Everything else that `SQLITE_USE_SEH` guards -
   the `Wal` fields `lockMask`, `pFree`, `pWiValue`, `iWiPg`, `iSysErrno`, the `lockMask`
   bookkeeping in `walLockShared/Exclusive` and `walUnlock*`, `walHandleException()`,
   `walAssertLockmask()`, `sqlite3WalSystemErrno()`, `sqlite3PagerWalSystemErrno()`,
   `sqlite3SystemError()` in main.c - is plain C and compiles unchanged.
2. The SEH macro block of wal.c gets an `#ifdef __CCGO__` branch ahead of the MSVC one:
   - declarations of `int modernc_seh_try(Wal*, int(*)(Wal*,void*), void*, int(*)(Wal*))` and
     `void modernc_seh_inject(Wal*)`, both provided by `seh.go`;
   - `SEH_INJECT_FAULT` → `modernc_seh_inject(pWal)`;
   - `static int walSehTry(...)`: the `assert(walAssertLockmask(pWal) && nSehTry==0)` /
     `VVA_ONLY(nSehTry++)` bookkeeping of `SEH_TRY`/`SEH_EXCEPT` around `modernc_seh_try()`;
   - **`SEH_TRY` and `SEH_EXCEPT` are deliberately left undefined under `__CCGO__`**, so that a
     block added by a future SQLite upgrade fails to transpile instead of silently running
     unguarded. `<windows.h>`, `sehExceptionFilter()` and `sehInjectFault()` (uses `RaiseException`)
     stay MSVC-only.
3. The nine sites, each as `#ifdef __CCGO__` / `#else` (original) / `#endif`:

   | wal.c function | body | thunk | on fault |
   |---|---|---|---|
   | `sqlite3WalSnapshotRecover` (SNAPSHOT) | `walSnapshotRecover(pWal,pBuf1,pBuf2)` | `walSehSnapshotRecover`, `void*[2]` | `SQLITE_IOERR_IN_PAGE` |
   | `sqlite3WalBeginReadTransaction` | `walBeginReadTransaction(pWal,pChanged)` | `walSehBeginReadTransaction`, arg = `pChanged` | `walHandleException` |
   | `sqlite3WalFindFrame` | `walFindFrame(pWal,pgno,piRead)` | `walSehFindFrame`, `struct WalSehFindFrame` | `SQLITE_IOERR_IN_PAGE` |
   | `sqlite3WalBeginWriteTransaction` | inline `memcmp` of the header | `walSehBeginWriteCheck` | `SQLITE_IOERR_IN_PAGE` |
   | `sqlite3WalUndo` | inline header restore + `xUndo` loop + `walCleanupHash` | `walSehUndo`, `struct WalSehUndo` | `SQLITE_IOERR_IN_PAGE` |
   | `sqlite3WalSavepointUndo` | `walCleanupHash(pWal)` | `walSehCleanupHash` | `SQLITE_IOERR_IN_PAGE` |
   | `sqlite3WalFrames` | `walFrames(...)` | `walSehFrames`, `struct WalSehFrames` | `walHandleException` |
   | `sqlite3WalCheckpoint` | inline: read header, `walCheckpoint`, outputs | `walSehCheckpoint`, `struct WalSehCheckpoint` | `walHandleException` |
   | `sqlite3WalSnapshotCheck` (SNAPSHOT) | inline: CKPT shared lock + salt check | `walSehSnapshotCheck`, arg = `pSnapshot` | `walHandleException` |

   The inline bodies were moved verbatim; the only edits are `a->` accesses and local `rc`
   variables that were function locals before. Where the original block assigned `rc` on top of an
   earlier `rc`, that earlier value was provably `SQLITE_OK` (checked at each site), so
   `rc = walSehTry(...)` is equivalent.

How the patch was produced, for the next SQLite upgrade: a Python script doing exact-string
replacements on the pre-patched amalgamation (asserting exactly one match each), then
`diff -u --label sqlite3.c --label sqlite3.c sqlite3.c.orig sqlite3.c`. Upstream's SEH code has
been stable since 3.44, so refreshing is expected to be a matter of context lines.

`generator.go`: applies the patch; drops `-DSQLITE_OMIT_SEH` from the two *library* configurations
for Windows (the testfixture, speedtest1 and mptest configurations keep it: testfixture also sets
`-D_MSC_VER=1`, and the other two do not compile the library); adds `GO_GENERATE_LIBONLY=1`, which
returns right after the library has been generated and copied - no testfixture/speedtest1/mptest -
for iterating on `internal/*.patch` (a full library-only run for one target takes a couple of
minutes here).

### 3.3 The Go side: `seh.go`

Hand-written, lives next to the transpiled code here (`package libsqlite3`) and in
`modernc.org/sqlite/lib` (`package sqlite3`); the two copies must stay identical but for the package
clause (precedent: `libsqlite3_windows.go`'s `___umulh`). Symbols, as ccgo emits them for undefined
externs with `--prefix-undefined=_` under `-ignore-link-errors` (verified with a 10-line transpile):

- `func _modernc_seh_try(tls, pWal, xBody, pArg, xOnFault uintptr) int32`: arms the guard, calls
  `xBody(pWal, pArg)` through the ccgo function-pointer convention, restores the guard on the way
  out. On a caught fault: `sehLog()` → `xOnFault(pWal)` if non-zero, else `SQLITE_IOERR_IN_PAGE`.
  Any other panic is re-raised.
- `sehCaught()`: accepts a simulated fault, or a `runtime.Error` with `Addr()` whose address lies in
  one of `pWal->apWiData[0..nWiData)` (32 KiB each). This is *stricter* than SQLite's filter (any
  `EXCEPTION_IN_PAGE_ERROR` inside the block): a fault anywhere else is a bug and keeps crashing,
  which also keeps libc-level bugs like the netbsd mmap PAD bug (#246) or the armv6 SIGBUS (#197)
  from being masked - well, not entirely: a wrong mapping *inside* the wal-index would be reported
  as an I/O error, with the address in the log line below.
- `sehLog()`: `sqlite3_log(SQLITE_IOERR_IN_PAGE, "SEH emulation: memory fault at %p in the
  wal-index mapping, returning SQLITE_IOERR_IN_PAGE")`, so it lands wherever the application routes
  `SQLITE_CONFIG_LOG`.
- `func _modernc_seh_inject(tls, pWal uintptr)`: `SEH_INJECT_FAULT`. Countdown in an atomic; the
  n-th site reached after `SehInject(n)` panics with a private value the trampoline accepts.
  `SehPending()` reads the countdown. This replaces `sqlite3FaultSim(650)`, which only exists in
  `SQLITE_TEST` builds; a call per site costs ~2 ns.

### 3.4 What changes for users

- Without a fault: nothing observable. Cost: one `SetPanicOnFault` pair, one `defer`/`recover` and
  one indirect call per guarded entry point (~14 ns measured, 0.5 ns for the unguarded call), the
  `lockMask` bookkeeping on lock operations, ~17 cheap `_modernc_seh_inject` calls per transaction.
  `sqlite3WalFindFrame` is per page lookup, so a WAL-mode read of N pages pays N × ~20 ns.
- With a fault inside the wal-index mapping: the statement fails with `disk I/O error`, extended
  code 8714 (`SQLITE_IOERR_IN_PAGE = SQLITE_IOERR | 34<<8`), the connection stays usable, and one
  `sqlite3_log` line is emitted. In `modernc.org/sqlite` the error is a `*sqlite.Error` with
  `Code()` 10 or 8714 depending on whether the connection enabled extended result codes.
- No API surface beyond `lib.SehInject`/`lib.SehPending` for tests. Nothing to migrate.

### 3.5 Decisions taken, and the alternatives

- **All targets, not Windows-only.** Upstream is Windows-only because SEH is. Our mechanism is
  portable, a `SIGBUS` on a truncated `-shm` is the unix twin of the Windows fault, an error beats
  a crash, and - decisive for a maintainer without Windows - it makes the whole thing testable on
  linux/amd64 (`TestSEHTruncatedShm`). To go Windows-only instead: add `-DSQLITE_OMIT_SEH` to the
  `default:` (unix) library configuration in `generator.go`; `seh.go` stays as is (unused symbols
  are fine) and the tests skip themselves. The historical unix `-shm` faults in this project were
  our own bugs (#144 old OpenBSD, #197 a libc version mismatch), which the address filter and the
  log line keep diagnosable.
- **No opt-out switch yet.** If wanted, a package-level atomic checked at the top of
  `_modernc_seh_try` (skip `SetPanicOnFault`, call the body directly) restores today's crash for
  debugging; or a build tag. Not needed for compatibility: no user code can depend on a crash.
- **Not ccgo.** A general `__try/__except` would need new keywords in cc/v4, the
  `GetExceptionCode`/`GetExceptionInformation` intrinsics, lowering of `return`/`goto`/`__leave` out
  of the protected block, fault classification, and `windows.h` types for the filter - weeks of
  compiler work for a single consumer whose upstream already factored the protected code into
  wrapper + body functions. The patch route reuses SQLite's own cleanup code untouched.
- **Address filter on** (§3.3). Turn it off by returning `addr, true` from `sehCaught` for any
  `sehFaultAddr` if it ever proves too strict (e.g. a platform reporting page-rounded addresses,
  which none of ours do: Windows hands out `ExceptionInformation[1]`, unix `si_addr`).
- **`iSysErrno` stays 0.** Go does not expose the NTSTATUS behind the exception.

## 4. Verification done (linux/amd64, this machine, 2026-09-03)

- `internal/sqlite_issue221.patch` (453 lines) applies after the three existing patches with no
  offset or fuzz. `gcc -fsyntax-only` of the result passes with `-D__CCGO__` under `NDEBUG` and
  under `SQLITE_DEBUG` (asserts on), without `__CCGO__` (SEH stays off: no `SEH_TRY` reaches the
  compiler), and with `x86_64-w64-mingw32-gcc -DSQLITE_OS_WIN=1 -D__CCGO__`.
- `GO_GENERATE_NOWIN=1 GO_GENERATE_LIBONLY=1 GO_GENERATE_DIR=... go run generator*.go` on the
  branch: `ccgo_linux_amd64.go` regenerated; it contains one `_modernc_seh_try(` call (inside
  `_walSehTry`), 17 `_modernc_seh_inject(` calls, `_walSehCheckpoint`, `_walHandleException`, and
  the `FlockMask`/`FpWiValue` fields; `go build ./...` passes.
- `go test -count=1 -run TestSEH -v .`:
  - `TestSEHTruncatedShm`: `os.Truncate(-shm, 0)` under the live mapping; `SELECT` →
    `SQLITE_IOERR_IN_PAGE`; `os.Truncate(-shm, 32768)`; `SELECT`, `INSERT`,
    `PRAGMA wal_checkpoint(TRUNCATE)`, `SELECT` all succeed (the wal-index is rebuilt from the -wal).
  - `TestSEHInjectedFault`: one simulated fault per reachable `SEH_INJECT_FAULT` site, connection
    checked after each: `SELECT` 10 sites, `INSERT` 14, `wal_checkpoint(PASSIVE)` 14,
    `wal_checkpoint(TRUNCATE)` 8, `BEGIN; INSERT; SAVEPOINT; INSERT; ROLLBACK TO; COMMIT` 18.
  - `TestSEHTrampoline`: real `SIGSEGV` on a `PROT_NONE` page registered as the only wal-index page
    of a fake `Wal`: body result passes through when no fault; the on-fault callback runs; the
    `xOnFault == 0` form yields 8714; a simulated fault is handled; `SetPanicOnFault` is restored;
    with the page not registered the fault propagates.
- Standalone mechanism proof (before the patch existed): fault in a truncated `MAP_SHARED`
  mapping under five frames with deferred cleanups, recovered with the right address, deferred
  functions ran, process healthy afterwards; guard overhead 14.4 ns/op vs 0.5 ns/op direct.
- windows/amd64 transpile with the patch: `GO_GENERATE_WIN=1 GO_GENERATE_LIBONLY=1 ... go run generator*.go`
  regenerated `ccgo_windows.go` (banner without `-DSQLITE_OMIT_SEH`; same symbol counts as the
  linux transpile); `GOOS=windows GOARCH=amd64` and `GOARCH=arm64` `go build ./...` and
  `go test -c` pass, so `seh_windows_test.go` compiles. Not run: no Windows machine here; the
  builders will run `TestSEHTrampoline` (real access violation through the trampoline) and
  `TestSEHInjectedFault`. windows/386 was regenerated the same way (`GO_GENERATE_WIN32=1`), same
  symbol counts, `GOOS=windows GOARCH=386 go build ./...` and `go test -c` pass.
- Multi-process WAL sanity (`TestConcurrentProcesses`, mptest, with the emulation active):
  `go test -count=1 -run 'TestConcurrentProcesses$' .` on this branch, i.e. the committed
  `mptest/ccgo_linux_amd64.go` linked against the patched library: PASS in 411 s, `crash01.test`
  and `multiwrite01.test` × 4 journal modes, `--repeat 20`, 0 errors out of 1880/1620 tests each
  (8 runs). The guard sits on every WAL entry point those processes hammer, so this is the
  multi-process regression check for "nothing changes without a fault".
- `modernc.org/sqlite` driver-level tests (`seh_test.go` on its `seh-emulation` branch), run
  against `lib/` temporarily re-vendored from this branch: with `../libsqlite3` pointed at this branch, `make vendor`'s
  steps (expand, vendor, undup, gofmt) ran clean - the shared `_walSehTry` landed in a
  `lib/sqlite_g_*.go` file - and with `go.mod` bumped to libc v1.75.7 (what this branch's
  transpiles link against) `go test -run TestSEH -v .` passed: `TestSEHTruncatedShm` (real fault →
  `*sqlite.Error` "disk I/O error", then recovery) and `TestSEHInjectedFault` (SELECT 10, INSERT 14,
  `wal_checkpoint(PASSIVE)` 14, `wal_checkpoint(TRUNCATE)` 8 sites). A WAL-related subset of the
  driver suite (`-run 'WAL|Wal|Checkpoint|TestConcurrentGoroutines|TestIssue25[0-9]|TestBackup|TestSnapshot'`)
  passed in 47 s. None of the vendored files were committed: the sqlite branch carries only the
  hand-written pieces, the real re-vendor happens after the libsqlite3 release.

## 5. Landing it (after the libsqlite3 release and the sqlite tag)

1. Here, in one commit on master: `git merge seh-emulation` - keep the "regenerate for testing"
   commit in it, so that the three Windows builders, which never autogen, test the patched
   transpiles from their first run - then fill in the CHANGELOG date and **blank all 20
   `internal/autogen/*.mod` snapshots** (`for f in internal/autogen/*.mod; do echo > $f; done`;
   nuc64 rewrites the three Windows ones itself). The farm regenerates a target only when its
   snapshot differs from `go.mod` (`builder_test.go:2220`), and nothing on the branch touches
   `go.mod`, so a plain merge would regenerate nothing: the tree would carry the patched
   linux/amd64 and Windows transpiles next to 16 unpatched ones and `TestSEHInjectedFault` would
   fail on those 16 ("no SEH_INJECT_FAULT site reached"). Do not push while another sweep is still
   running, a blanked set restarts it. The farm's autogen then produces the 19 transpiles with
   the patch; the Windows builders run `TestSEHTrampoline` and `TestSEHInjectedFault`
   (`TestSEHTruncatedShm` skips on windows). (Corrected 2026-09-04, see §8.)
2. `modernc.org/sqlite`: merge its `seh-emulation` branch (`lib/seh.go`, `seh_test.go`, CHANGELOG
   with the version/date filled in), then `make vendor` from the regenerated libsqlite3, `make
   editor`, `go test -run TestSEH -v`, the usual suite, tag once the dashboard is green.
   `lib/seh.go` is a verbatim copy of `seh.go` with `package sqlite3`; `vendor_libs` does not copy
   hand-written files, so keep the two in sync by hand (as for `libsqlite3_windows.go`).
3. Close #221 with a note pointing at the CHANGELOG entry; tell hazyhaar on PR #7 what landed.
4. Optional follow-ups: a `doc.go` paragraph on `SQLITE_IOERR_IN_PAGE`; an opt-out switch (§3.5);
   a `.patch2` for `src/wal.c`+`src/sqliteInt.h` if the Tcl testfixture should ever run upstream's
   `walseh*.test` (it cannot today: `-D_MSC_VER=1 -DSQLITE_OMIT_SEH`).

## 6. Things to watch

- **SQLite upgrades.** If a new `SEH_TRY` appears, the ccgo transpile fails with an undefined
  `SEH_TRY` - by design. Rewrite the new site like the nine here. If the patch's context drifts,
  regenerate it (§3.2).
- **`go vet`** reports "possible misuse of unsafe.Pointer" for `seh.go`'s `(*TWal)(unsafe.Pointer(pWal))`
  and the `apWiData` walk - the same class as thousands of lines of generated code; `make all` runs
  golint and staticcheck, not vet.
- **Panics from other sources inside a guarded body** (a nil dereference, a libc `todo`) are re-raised
  unchanged; the re-panic keeps the original frames in the trace.
- **Concurrency.** `SehInject` is process-global test scaffolding; tests using it must not run in
  parallel with each other.
- **`SetPanicOnFault` is per goroutine** and restored on every exit path, so it cannot leak into user
  code; nested guards (there are none today) would be handled by the save/restore.

## 7. Facts checked along the way (so nobody has to re-check them)

- `SQLITE_IOERR_IN_PAGE = 8714`, not 6410 (the PR comment was corrected).
- Only wal.c uses `SEH_TRY`; the other `SQLITE_USE_SEH` uses (pager.h/pager.c/main.c) are plain C.
- 9 `SEH_TRY` sites in 3.53.4 (two under `SQLITE_ENABLE_SNAPSHOT`, which we enable), 17
  `SEH_INJECT_FAULT` sites reachable in our build.
- ccgo emits an undefined extern as `_name(tls, ...)` and takes function addresses with
  `__ccgo_fp(_fn)`; without `-ignore-link-errors` it fails with `undefined: "name" external`.
- `libc.NewVaList(uintptr)` is supported; free the result with `libc.Xfree`.
- The three Windows targets share one transpile for amd64/arm64 (`ccgo_windows.go`) plus
  `ccgo_windows_386.go`; there are 20 targets and 19 transpiles, not 24.

## 8. Addenda from the MR !4 round (2026-09-04)

hazyhaar's second round - libsqlite3!4 plus modernc-org/sqlite#7 rebased - is a from-scratch build
of this design; `HANDOFF-seh-emulation-mr4.md` (untracked, repo root) is the review and
`HANDOFF-seh-emulation-mr4-REPLY.md` the reply from this side. Not merged: its trampoline is
windows-only while the patch enables SEH on every target (nine `undefined: _modernc_seh_try` on a
regenerated linux/amd64), its `walCheckpointThunk` keeps `isChanged` local so the checkpointing
connection serves stale pages afterwards (x=1 after another connection wrote 2, all four
checkpoint modes), and its trampoline treats any panic without `Addr()` as an in-page fault. Two
things changed on this branch because of it:

- `TestSEHCheckpointInvalidatesCache` in `seh_test.go`: two connections, A reads, B writes, A
  checkpoints in each of the four modes, A must then read B's value. It needs no fault, passes with
  or without the emulation, and guards the one thunk that writes to a function local.
- `TestSEHTrampoline` no longer uses Go closures. `__ccgo_fp` of a non-escaping closure is the
  address of a funcval on the goroutine stack; the panic path between the fault and the callback
  grows the stack inside `sehLog()` (`libc.NewVaList` allocates), and the `uintptr` the trampoline
  holds then points into the old copy of the stack - the callback ran with a stale context and
  incremented a dead `faultCalls` (4 runs in 5 on this machine with go1.27.0; the 2026-09-03 pass
  in §4 was luck: the goroutine happened to sit on an already-grown stack). The library is not
  affected: its `xBody`/`xOnFault` are static funcvals and `pArg` is on the libc TLS stack. The
  test now uses top-level functions and package-level counters, exactly like the transpiled
  thunks. Proof recorded in the mr4 reply.

Also confirmed here with ccgo v4.35.0: the `-eval-all-macros` finding behind §3.3's hard-coded
32768 (`const WALINDEX_PGSZ = 0`, use sites right). Fixed the same day in cc/v4, where the
evaluation lives (branch `eval-all-macros` in `../cc`, credit to hazyhaar in the commit): the
`#if` evaluator that treated every identifier as 0 is replaced by a parse and type check of the
expanded replacement list in file scope. It corrects 235 exported constants of the linux/amd64
transpile, not just the two; the categories and the downstream consequences (sqlite's
`vendor_libs` must filter `SQLITE_STATIC` like `SQLITE_TRANSIENT`, libc's `MAP_FAILED` changes) are
in `HANDOFF-seh-emulation-mr4-REPLY.md` §5. It reaches this repo through a cc tag, a ccgo tag
and a `go.mod` bump, i.e. a farm sweep of its own or one shared with the SEH landing.
