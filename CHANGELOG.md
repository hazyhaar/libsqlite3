# Changelog

 * 2026-08-27: Linux OFD locking is now opt-in and off by default - cznic/sqlite#255, where Gani
   Georgiev asked for it to ship opt-in for a couple of releases before it becomes the default.
   Nothing had been released in between, so the default behaviour of every released version and of
   this one is the same: upstream's POSIX record locks on every unix target. Enable OFD locks per
   process, before the first database file is opened, either with the environment variable
   MODERNC_SQLITE_OFD_LOCK=1 - the library reads it itself, once, in sqlite3_os_init(), so
   testfixture, mptest and anything else linking the library pick it up without any Go plumbing -
   or with the new Xmodernc_ofd_locking(tls, 1), which overrides the environment and returns the
   previous setting: -1 means OFD locks are not available (every unix target but linux, or a linux
   kernel that rejected F_OFD_*), -2 that the setting is frozen because the process has already
   attempted a lock - the kind of lock in use cannot change under a held one, a POSIX F_UNLCK does
   not release an OFD lock and vice versa, so a late flip would strand every lock held; the freeze
   lifts at sqlite3_shutdown() - and a negative argument only queries. The function exists on every
   unix target and not on windows, so the Go side needs one unix file and one windows stub. The
   switch is process-wide on purpose, not a per-database option: POSIX and OFD locks are different
   owners even inside one process, so all connections to one file must use the same kind. Two
   fixes ride along. The patch is now guarded by defined(__linux__) as well as defined(F_OFD_SETLK):
   macOS has had F_OFD_* since 10.13 and darwin transpiles against the host SDK, so the 2026-08-26
   change compiled the whole OFD path into darwin/{amd64,arm64} - only the fact that libc's darwin
   Xfcntl64 panics on unknown commands instead of passing them through stopped that from being a
   silent switch, and darwin-m1 went red on TestConcurrentProcesses ("Xfcntl64: TODO ... 4 90");
   with the guard darwin and the BSDs get upstream's locking and export only the -1 returning
   setter. And the EINVAL fallback is latched: it can only happen on the very first OFD fcntl();
   once one has succeeded a later EINVAL is returned as the I/O error it is instead of switching
   the process to POSIX mode under OFD locks that a POSIX F_UNLCK cannot release. Database-file
   locks go through the OFD wrapper and no longer through osSetPosixAdvisoryLock(), which only the
   -shm locks use; that is the same call today, but with SQLITE_ENABLE_SETLK_TIMEOUT it would drop
   the blocking-lock timeout for the database file, so the patch refuses to compile with that
   option until it is revisited. New Makefile
   targets: `make locktest` runs the lock/WAL subset of the Tcl suite in both modes, `make
   tcltest_ofd` and `make mptest_ofd` the Tcl suite resp. mptest with OFD locks. On the builders
   TestTclTestOFD and TestConcurrentProcessesOFD (linux only) run the lock/WAL subset - -suite=locks,
   45 files, the only part of the suite that can tell the two locking modes apart - resp. mptest with
   MODERNC_SQLITE_OFD_LOCK=1, so every linux target covers both modes at the cost of a few minutes
   (walthread.test's cases run for a fixed 20 s each) plus one more mptest; TestOFDLocking is their
   positive control: it observes in /proc/locks
   that the variable selects OFD resp. POSIX kernel locks and that only the former survive a
   close() of a stray descriptor. linux/amd64 regenerated here; every internal/autogen/*.mod snapshot
   blanked again so the farm regenerates the ten targets that had already picked up the 2026-08-26
   inputs.

 * 2026-08-26: Linux: database file locks are now Open File Description locks (F_OFD_SETLK*)
   (opt-in and off by default since 2026-08-27, see above):
   Nathan Herring's MR !3 (cznic/sqlite#255) plus a fixup commit. POSIX record locks belong to the
   process, so any close() of an unrelated descriptor of the database file - e.g. an os.File the
   application opened and closed - silently dropped every lock SQLite held on it; OFD locks belong to
   the open file description, so that can no longer happen. SQLite's unixInodeInfo refcounting assumes
   a single lock owner per inode and process, so every kernel lock on an inode goes through one
   designated descriptor (pInode->hLock, the first locker's; a read-write connection joining at SHARED
   takes over from a read-only one, because F_WRLCK needs a writable descriptor). Where F_OFD_* is
   unsupported (EINVAL, kernels before 3.15) and on the other unix targets everything stays upstream's
   POSIX locking. The fixup on top of the MR: the take-over locked through the old read-only descriptor
   and then released it, leaving the process without any kernel lock for the rest of the read
   transaction (an external writer could commit under two open read transactions); the designated
   descriptor was recorded before the lock was taken, so a SHARED attempt failing after its PENDING
   lock had succeeded left a stale fd behind (SQLITE_IOERR_LOCK, or a PENDING-byte lock leaked onto
   whatever file reused the fd number); and none of that machinery was gated on OFD locks being in
   effect, so the take-over also ran on darwin/*BSD, where a whole-file F_UNLCK through any descriptor
   releases all of the process's locks. generator.go keeps libc at v1.75.5 (= go.mod; the tag that has
   F_OFD_* for every linux target) and every internal/autogen/*.mod snapshot is blanked, so the farm
   regenerates all 20 targets rather than only the eight linux ones: the code above is compiled on
   every unix target and any target-specific mismatch should show up now rather than at the next
   unrelated dependency bump. Regression tests for the take-over: cznic/sqlite!136.

 * 2026-08-23: testfixture: build with -DCONFIG_SLOWDOWN_FACTOR=10.0 (generator.go), upstream's
   knob for slow test builds. like-14.{1,2} time one GLOB resp. LIKE query with many wildcards and
   fail above 1000*$sqlite_options(configslower) microseconds - the test prints "ms", but Tcl's
   [time] reports microseconds - and on the emulated builders the transpiled testfixture needs
   300-1500 us for that single round trip (linux/loong64 QEMU VM, measured in isolation, vs.
   135-728 us for native C on the same VM; the failing builder run saw 597 us for the GLOB and
   1500 us for the LIKE variant), so the stock 1 ms limit is a coin toss that already had to be
   waived on freebsd/arm and linux/s390x. 10x keeps the test meaningful: it guards against the
   exponential pattern matcher of SQLite < 3.16.0, which needs ~75 ms per query natively on an
   amd64 desktop (3.15.2) where the fixed one needs ~20 us. Takes effect on the next per-target
   regeneration: linux/amd64 regenerated here, and the all_test.go waivers for freebsd/arm and
   linux/s390x are dropped in favour of regenerating those two (their internal/autogen/*.mod
   snapshots are blanked).

 * 2026-08-03: Upgrade to SQLite 3.53.4. It carries upstream's own fix for the super-journal
   rollback corruption reported from here on 2026-07-20
   (https://sqlite.org/forum/info/2026-07-20T18:27:00Z): check-in bf70dadc2d455844 applies the
   same one-line `zOut[0]==0` test we had been carrying, and `src/pager.c` has no other change
   between 3.53.3 and 3.53.4. internal/sqlite_superjournal.patch{,2} is therefore dropped.
   Upstream also added test/crash9.test as the regression test for it, which runs as part of the
   default `full` suite.

 * 2026-07-21: mptest: fix TestConcurrentProcesses on slow builders (openbsd/arm64 under QEMU),
   which failed with "database is locked" / "timeout waiting for all clients". Root cause: the
   generator transpiled the *upstream* mptest.c, which hardwires the busy timeout to the
   DEFAULT_TIMEOUT constant and never forwards --timeout to spawned clients, so --timeout 120000
   was ignored and 10s was too short. (Not a ccgo bug — ccgo transpiles both sources faithfully.)
   Wire internal/overlay/mptest/mptest.c — which makes the timeout the runtime g.defaultTimeout and
   forwards --timeout — into generation (generator.go), and point the --wait barrier default at
   g.defaultTimeout. Takes effect on the next per-target regeneration.

 * 2026-07-20: Patch an upstream SQLite 3.53.3 regression where a hot rollback journal was
   deleted without being played back, leaving the database corrupted, after a crash during a
   multi-database (ATTACH) transaction. See internal/sqlite_superjournal.patch{,2}.

 * 2026-06-13: Add netbsd/amd64 support.

 * 2026-05-28: Add freebsd/{386,arm} support (not yet functional).

 * 2026-05-05: Upgrade to SQLite 3.53.1

 * 2026-04-10: Upgrade to SQLite 3.53.0

 * 2026-01-16: Upgrade to SQLite 3.51.3
  * linux/s390x is failing one test, see https://sqlite.org/forum/forumpost/cdeb669113

 * 2026-01-16: Upgrade to SQLite 3.51.1

 * 2026-01-06: v1.11.0 - Add tier 2 openbsd/{amd,arm64}.

 * 2024-07-22: v1.5.2 - Add windows/386 support.

 * 2024-03-12: v1.2.0 - Add linux/loong64 support.

 * 2024-02-13: v1.0.0
