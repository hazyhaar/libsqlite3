# Reply handoff: the mptest "--timeout" issue is NOT a ccgo miscompilation

**From:** the `modernc.org/ccgo/v4` maintainer/agent.
**Re:** `HANDOFF-ccgo-mptest-timeout.md` ("unsound constant-folding of a runtime-assigned global-struct field").
**Status:** investigated end-to-end. **ccgo is compiling correctly — there is no fold.** The real
problem is on the libsqlite3 generation side, and a one-line `generator.go` fix (already applied,
see below) resolves it. A separate, minor ccgo bug that produced the confusing stray
`a_linux_amd64.go` has been fixed in ccgo.

---

## TL;DR

The committed `mptest/ccgo_*.go` was generated from the **upstream** SQLite source
`sqlite-src-3530300/mptest/mptest.c`, which genuinely:

- has **no** `defaultTimeout` field,
- literally writes `g.iTimeout = DEFAULT_TIMEOUT;` (a bare constant), and
- builds the client spawn command with `"...--trace %d"` — **no `--timeout %d`**.

The handoff compared that generated output against a **different** file,
`internal/overlay/mptest/mptest.c`, which *you* edited (commit f7a48d58, "raise timeouts to survive
slow builders") to add the `defaultTimeout` field, the `(iTmout>0)?iTmout:DEFAULT_TIMEOUT` runtime
assignment, and the extra `--timeout %d` argument. Seeing the field "missing" and the `?:` "gone"
in the generated code looked exactly like an aggressive constant-fold — but ccgo never saw that
source. **The generator compiles the upstream harness; the overlay was never wired into
generation**, so your edit has been inert.

ccgo handles **both** sources correctly:

| source fed to ccgo | generated result |
|---|---|
| upstream `mptest.c` (bare `DEFAULT_TIMEOUT`) | constant — matches the committed file **byte-for-byte** (except the embedded `SQLITE_SOURCE_ID` string) |
| overlay `mptest.c` (runtime `g.defaultTimeout`) | field kept, `g.FdefaultTimeout = (iTmout>0 ? iTmout : 10000)` at runtime, `mprintf` keeps `--timeout %d` + the 5th vararg — i.e. exactly what you want |

## Proof you can re-run

From the repo root, with a ccgo built from current `modernc.org/ccgo/v4`:

```sh
# 1. Extract the upstream harness and the amalgamation header.
mkdir -p /tmp/mp && cd /tmp/mp
unzip -oq /path/to/libsqlite3/sqlite-src-3530300.zip 'sqlite-src-3530300/mptest/mptest.c'
unzip -oj /path/to/libsqlite3/sqlite-amalgamation-3530300.zip 'sqlite3.h'

# 2. Apply the same sed patch the generator applies, then transpile the UPSTREAM source
#    (run from the libsqlite3 module root so -lsqlite3 resolves to modernc.org/libsqlite3).
sed 's/strcmp(sqlite3_sourceid()/0 \&\& strcmp(sqlite3_sourceid()/' \
    sqlite-src-3530300/mptest/mptest.c > /tmp/mp/mptest_upstream.c
cd /path/to/libsqlite3
ccgo -DNDEBUG -DSQLITE_DISABLE_INTRINSIC -I /tmp/mp \
     -ignore-unsupported-alignment -ignore-link-errors \
     -o /tmp/mp/upstream.go /tmp/mp/mptest_upstream.c -lsqlite3

# 3. Diff against the committed file — identical except the SQLITE_SOURCE_ID string literal.
diff <(tail -n +5 /tmp/mp/upstream.go) <(tail -n +5 mptest/ccgo_linux_amd64.go)
#   -> one differing line: "...d78alt1" (src-tree header) vs "...d782c62" (amalgamation header)

# 4. Now transpile the OVERLAY source the same way -> correct, --timeout-honoring code.
ccgo -DNDEBUG -DSQLITE_DISABLE_INTRINSIC -I /tmp/mp \
     -ignore-unsupported-alignment -ignore-link-errors \
     -o /tmp/mp/overlay.go internal/overlay/mptest/mptest.c -lsqlite3
grep -n 'FdefaultTimeout\|timeout %d' /tmp/mp/overlay.go   # field kept, --timeout %d kept
```

(For the record: the handoff's three minimal repros "all transpile correctly" for the simple reason
that there is no bug to trigger. And the source-id difference in step 3 is purely the embedded
version string — the committed file was generated with the src-tree's `sqlite3.h`, I used the
amalgamation's; it has nothing to do with the timeout logic.)

## The fix (already applied to `generator.go`)

Wire the overlay into generation — copy it over the upstream harness right before the existing
`sqlite3_sourceid` sed patch, so the sed still applies to the overlaid file:

```go
	os.Mkdir("mptest", 0770)
	// Overlay our patched mptest harness over the upstream one before transpiling.
	// ... (see the committed comment)
	mustCopyFile(filepath.Join(makeRoot, "mptest", "mptest.c"), filepath.Join("internal", "overlay", "mptest", "mptest.c"), nil)
	util.MustShell(true, nil, sed, "-i", `s/strcmp(sqlite3_sourceid()/0 \&\& strcmp(sqlite3_sourceid()/`, filepath.Join(makeRoot, "mptest", "mptest.c"))
```

`mustCopyFile(dst, src, nil)` overwrites the existing upstream copy (nil `canOverwrite` ⇒ overwrite).
`go build -o /dev/null generator*.go` passes with this change.

### What you still need to do

1. **Regenerate** so the shipped `mptest/ccgo_*.go` are rebuilt from the overlay:
   `make generate` (host), and the per-target builds for the other GOOS/GOARCH.
2. **Verify** each regenerated `mptest/ccgo_*.go` now has, in `main`:
   - an `FdefaultTimeout int32` field in `type Global`;
   - `g.FdefaultTimeout = <iTmout>0 ? iTmout : DEFAULT_TIMEOUT>` then `g.FiTimeout = g.FdefaultTimeout`;
   - the startClient `sqlite3_mprintf` keeping `--timeout %d` with `g.FdefaultTimeout` as the 5th `VaList` arg.
3. **Optional cleanup:** because ccgo honors the runtime assignment, you can drop the `120000`
   workaround and set `#define DEFAULT_TIMEOUT 10000` again in `internal/overlay/mptest/mptest.c`
   (keep the runtime `g.defaultTimeout = (iTmout>0)?iTmout:DEFAULT_TIMEOUT`), and point `--wait`'s
   default at `g.defaultTimeout`. `--timeout N` will then work for any N; 10000 is only the default
   when `--timeout` is absent.
4. You can retire `HANDOFF-ccgo-mptest-timeout.md` (resolved: not a ccgo bug).

## The stray `a_linux_amd64.go` (separate, minor ccgo bug — fixed in ccgo)

There was an untracked `a_linux_amd64.go` (`package main`, header `'ccgo --version'`) sitting in the
repo root. It breaks `go build ./...` (mixed `package main`/`libsqlite3` in one dir) and it also
briefly derailed my investigation (it corrupted `packages.Load("modernc.org/libsqlite3")`). Root
cause: `ccgo --version` used to fall through into link-with-no-inputs and write the default output
file `a_<goos>_<goarch>.go` instead of printing a version. During `./configure`, `cc --version`
(ccgo posing as cc) could drop this file.

Fixed in ccgo (`v4/lib/ccgo.go`): `--version` now prints `ccgo version <module-version>` and returns
without building. Once you bump the pinned ccgo past this fix, the stray file can no longer appear.
I moved the pre-existing one out of the tree (it was untracked, so not mine to delete); if a copy is
still present, just `rm a_linux_amd64.go`.

## Summary

- **No ccgo change was needed for the timeout behavior** — ccgo was already correct.
- **libsqlite3 fix:** `generator.go` now overlays `internal/overlay/mptest/mptest.c` (applied). Regenerate + verify.
- **ccgo fix (incidental):** `ccgo --version` no longer emits a stray `a_*.go`; release + repin at your leisure.
