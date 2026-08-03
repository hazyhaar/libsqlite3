# Changelog

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
