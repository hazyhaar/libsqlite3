# Handoff to ccgo: unsound constant-folding of a runtime-assigned global-struct field

**Audience:** the `modernc.org/ccgo/v4` maintainer/agent.
**Reporter:** found while root-causing an `openbsd/arm64` builder failure in `modernc.org/libsqlite3`.
**Status:** confirmed, reproducible, present on **all** targets. Not yet fixed in ccgo. `libsqlite3`
carries a source-level workaround (see the end) so it is not blocking, but the workaround only
papers over a genuine ccgo miscompilation that likely affects other programs.

---

## One-sentence summary

ccgo replaces every read of the global-struct field `g.defaultTimeout` with the compile-time
constant `DEFAULT_TIMEOUT` (and drops the field from the struct **and** from a variadic
`sqlite3_mprintf` call), even though `g.defaultTimeout` is assigned a **runtime** value
(`iTmout`, parsed from `argv`). The `?:` that selects the runtime value is folded away to its
constant arm.

## The input (C)

`internal/overlay/mptest/mptest.c` — the upstream SQLite mptest harness, copied in as an overlay.
Relevant lines (grep these; line numbers may drift):

```c
/* struct field */
int defaultTimeout;                 /* mptest.c:89, inside the global "struct Global g" */
#define DEFAULT_TIMEOUT 10000       /* mptest.c:94 */

/* main(): --timeout is parsed into a RUNTIME int, then stored into the field */
int iTmout = 0;                                         /* mptest.c:1297 */
zTmout = findOption(argv+2, &n, "timeout", 1);          /* mptest.c:1325 */
if( zTmout ) iTmout = atoi(zTmout);                     /* mptest.c:1326 */
...
if( iTmout>0 ) sqlite3_busy_timeout(g.db, iTmout);      /* mptest.c:1373  <-- iTmout is live here */
...
g.defaultTimeout = (iTmout > 0) ? iTmout : DEFAULT_TIMEOUT;  /* mptest.c:1393  <-- runtime assign */
g.iTimeout       = g.defaultTimeout;                        /* mptest.c:1394 */

/* the field is later read in several places, incl. a variadic call that
   passes it as the 5th argument for the "--timeout %d" specifier: */
zSys = sqlite3_mprintf("%s \"%s\" --client %d --trace %d --timeout %d",
             g.argv0, g.zDbFile, iClient, g.iTrace, g.defaultTimeout);   /* mptest.c:644-645 */
```

`iTmout` is unambiguously a runtime value: it is `atoi()` of a command-line option, and it is used
at runtime one statement earlier (`sqlite3_busy_timeout(g.db, iTmout)`, which **is** emitted
correctly — see below). So `g.defaultTimeout` is NOT a compile-time constant.

## The output (generated Go) — what's wrong

Reference file: `mptest/ccgo_linux_amd64.go` (host target; same defect in every `mptest/ccgo_*.go`).

**1. The struct field is gone.** `type Global = struct { ... }` has `FiTimeout int32` but **no
`FdefaultTimeout`**:

```
$ awk '/^type Global = struct/,/^}/' mptest/ccgo_linux_amd64.go | grep -i timeout
	FiTimeout         int32          # only this; FdefaultTimeout is absent
```

**2. Every `g.defaultTimeout` read became the constant.** The C `g.iTimeout = g.defaultTimeout;`
(5 occurrences) all transpiled to:

```go
g.FiTimeout = int32(DEFAULT_TIMEOUT)     // const DEFAULT_TIMEOUT = 10000
```

**3. The `?:` was folded to its constant arm — the smoking gun.** In generated `main`, `iTmout` is
parsed and used correctly for `busy_timeout`, but the following field assignment lost the runtime
arm entirely:

```go
if iTmout > 0 {
    libsqlite3.Xsqlite3_busy_timeout(tls, g.Fdb, iTmout)   // iTmout honored HERE (correct)
}
...
g.FiTimeout = int32(DEFAULT_TIMEOUT)   // C was: g.defaultTimeout=(iTmout>0)?iTmout:DEFAULT_TIMEOUT; g.iTimeout=g.defaultTimeout;
                                       // the (iTmout>0)?iTmout:... arm was discarded
```

**4. The variadic `sqlite3_mprintf` lost the specifier AND the argument.** The format-string
literal was truncated from `...--trace %d --timeout %d` to `...--trace %d`, and the 5th vararg
(`g.defaultTimeout`) was dropped — 4 varargs where the C has 5:

```go
// startClient(): C format has 5 conversions + 5 args; generated has 4 + 4
zSys = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+743,
        libc.VaList(bp+8, g.Fargv0, g.FzDbFile, iClient, g.FiTrace))   // no g.defaultTimeout
// and __ccgo_ts+743 is the literal:  %s "%s" --client %d --trace %d   (no --timeout %d)
```

## Why it matters

Because `g.defaultTimeout` is pinned to `10000`, the transpiled mptest ignores `--timeout N`
entirely: the custom busy handler's threshold (`g.iTimeout`) is stuck at 10 s, and child worker
processes are spawned with **no** `--timeout` at all. On a slow/emulated builder (`openbsd/arm64`
under QEMU) operations legitimately exceed 10 s, so the coordinator fatally aborts
(`ERROR: timeout after 10000ms` / `FATAL: database is locked`) even though the harness passed
`--timeout 120000`. A natively-compiled mptest (clang) built from the same source honors
`--timeout` and passes cleanly. See `internal/overlay/mptest/mptest.c` and the
`openbsd-arm64-mptest-timeout` analysis; the builder log is at
`gitlab.com/cznic/builder .../logs/modernc.org/libsqlite3/openbsd_arm64`.

## Likely root cause (for you to confirm)

An interprocedural / whole-program constant-propagation pass appears to prove `g.defaultTimeout ==
DEFAULT_TIMEOUT` and then eliminate the field, its runtime assignment, and its uses. The proof is
unsound: the field is conditionally assigned `iTmout`, a value not known at compile time (and that
ccgo itself treats as runtime one statement earlier, in the emitted `busy_timeout(g.db, iTmout)`).

### Minimal isolation — attempted, NOT yet reproduced (these transpile CORRECTLY)

I tried to reduce this to a standalone repro with ccgo v4.34.6 and **failed** — the following all
generate correct code (field kept, `?:` runtime arm preserved, variadic arg + `--timeout %d`
retained). Recording them so you don't re-tread the same ground; the trigger must involve something
these omit (candidates below):

```c
/* (1) single function — CORRECT */
struct G { int f; } g;
int main(int c,char**v){ int r=c>1?atoi(v[1]):0; g.f=(r>0)?r:42; printf("%d\n",g.f); }

/* (2) cross-function field-to-field copy (mimics g.iTimeout = g.defaultTimeout) — CORRECT */
struct G { int iTimeout, defaultTimeout; } g;
static void useit(void){ g.iTimeout = g.defaultTimeout; }
int main(int c,char**v){ int t=c>1?atoi(v[1]):0; g.defaultTimeout=(t>0)?t:10000; useit();
                         printf("%d\n",g.iTimeout); }

/* (3) variadic-arg read (mimics the sqlite3_mprintf call), built with the SAME flags the
       generator uses (-DNDEBUG -DSQLITE_DISABLE_INTRINSIC -ignore-unsupported-alignment
       -ignore-link-errors) — CORRECT: FdefaultTimeout kept, VaList(... g.FdefaultTimeout),
       format string keeps "--timeout %d" */
extern char *sqlite3_mprintf_stub(const char*,...);
struct G { int iTrace, defaultTimeout; char*argv0; } g;
static char *build(int cl){ return sqlite3_mprintf_stub("--client %d --trace %d --timeout %d",
                                                        cl, g.iTrace, g.defaultTimeout); }
int main(int c,char**v){ int t=c>1?atoi(v[1]):0; g.defaultTimeout=(t>0)?t:10000;
                         printf("%s\n", build(2)); }
```

So the trigger is **not** simply "a `?:` into a global-struct field," nor cross-function reads, nor
the variadic pass. Things the real case has that these don't — please bisect these:
- `mptest` is transpiled with **`-lsqlite3`** (it imports the generated `libsqlite3` package), i.e.
  a larger whole-program/link context;
- the `Global` struct is large (17 fields) and `g.defaultTimeout` is read from **many** functions;
- full ccgo flag set from `generator.go` (the `--prefix-*` family, `-eval-all-macros` upstream,
  etc.);
- some optimizer threshold (number of uses / inlining / SSA over the whole module).

The reliable reproduction remains regenerating mptest itself (next section).

## How to reproduce in this repo

1. On linux/amd64: `make generate` (regenerates the host target incl. `mptest/ccgo_linux_amd64.go`),
   or just inspect the committed file.
2. Confirm the four observations above (the `awk`/`grep` snippets are copy-pasteable).
3. Cross-check any other `mptest/ccgo_*.go` — the defect is identical everywhere.

## How to verify a fix

After a ccgo fix, regenerating must yield, in `mptest/ccgo_*.go`:
- an `FdefaultTimeout int32` field in `type Global`;
- generated `main` computing it at runtime, e.g. `g.FdefaultTimeout = <iTmout>0 ? iTmout : 10000>`
  and `g.FiTimeout = g.FdefaultTimeout` (no bare `int32(DEFAULT_TIMEOUT)` where the C reads the
  field);
- the startClient `sqlite3_mprintf` format retaining `--timeout %d` with `g.FdefaultTimeout` as the
  5th `VaList` argument.

## Workaround currently proposed in libsqlite3 (context, not a ccgo fix)

Because a source-level `: g.defaultTimeout` would just re-fold, `libsqlite3` plans to bump the
constant instead — `#define DEFAULT_TIMEOUT 120000` and point the `--wait` default at it — so the
folded value is at least the intended 120 s. That is a hack around this bug; once ccgo honors the
runtime assignment, the clean fix (`--wait` default `= g.defaultTimeout`, `DEFAULT_TIMEOUT` back to
10000) becomes possible and `--timeout N` works for any N.
