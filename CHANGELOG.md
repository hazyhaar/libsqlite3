# Changelog

 * 2026-07-21: mptest: raise the busy/barrier timeouts so slow builders (openbsd/arm64 under
   QEMU) stop failing TestConcurrentProcesses with "database is locked" / "timeout waiting for
   all clients". The transpiled mptest ignores --timeout because ccgo constant-folds
   g.defaultTimeout; bump DEFAULT_TIMEOUT and the --wait default to match the intended 120s. See
   internal/overlay/mptest/mptest.c and HANDOFF-ccgo-mptest-timeout.md. Takes effect on the next
   per-target regeneration.

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
