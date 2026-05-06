// Code generated for linux/s390x by 'generator -DNDEBUG -I /tmp/libsqlite3/sqlite-src-3530100 -ignore-unsupported-alignment -ignore-link-errors -o mptest/ccgo_linux_s390x.go /tmp/libsqlite3/sqlite-src-3530100/mptest/mptest.c -lsqlite3', DO NOT EDIT.

//go:build linux && s390x

package main

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
	"modernc.org/libsqlite3"
)

var _ reflect.Type
var _ unsafe.Pointer

const DEFAULT_TIMEOUT = 10000
const ENOENT = 2
const MX_ARG = 2
const SQLITE_BUSY = 5
const SQLITE_CONFIG_LOG = 16
const SQLITE_DONE = 101
const SQLITE_ERROR = 1
const SQLITE_FCNTL_VFSNAME = 12
const SQLITE_NOTICE = 27
const SQLITE_OK = 0
const SQLITE_OPEN_CREATE = 4
const SQLITE_OPEN_READWRITE = 2
const SQLITE_ROW = 100
const SQLITE_SCHEMA = 17
const SQLITE_UTF8 = 1

type va_list = uintptr

/* The suffix to append to the child command lines, if any */

/* The directory separator character(s) */

/* Mark a parameter as unused to suppress compiler warnings */

// C documentation
//
//	/* Global data
//	*/
type Global = struct {
	Fargv0            uintptr
	FzVfs             uintptr
	FzDbFile          uintptr
	Fdb               uintptr
	FzErrLog          uintptr
	FpErrLog          uintptr
	FzLog             uintptr
	FpLog             uintptr
	FzName            [32]uint8
	FtaskId           int32
	FiTrace           int32
	FbSqlTrace        int32
	FbIgnoreSqlErrors int32
	FnError           int32
	FnTest            int32
	FiTimeout         int32
	FbSync            int32
}

/* The suffix to append to the child command lines, if any */

/* The directory separator character(s) */

/* Mark a parameter as unused to suppress compiler warnings */

// C documentation
//
//	/* Global data
//	*/
var g Global

/* Default timeout */

// C documentation
//
//	/*
//	** Print a message adding zPrefix[] to the beginning of every line.
//	*/
func printWithPrefix(tls *libc.TLS, pOut uintptr, zPrefix uintptr, zMsg uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i int32
	_ = i
	for zMsg != 0 && *(*uint8)(unsafe.Pointer(zMsg)) != 0 {
		i = 0
		for {
			if !(*(*uint8)(unsafe.Pointer(zMsg + uintptr(i))) != 0 && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(zMsg + uintptr(i)))) != int32('\n') && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(zMsg + uintptr(i)))) != int32('\r')) {
				break
			}
			goto _1
		_1:
			;
			i = i + 1
		}
		libc.Xfprintf(tls, pOut, __ccgo_ts, libc.VaList(bp+8, zPrefix, i, zMsg))
		zMsg = zMsg + uintptr(i)
		for libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(zMsg))) == int32('\n') || libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(zMsg))) == int32('\r') {
			zMsg = zMsg + 1
		}
	}
}

// C documentation
//
//	/*
//	** Compare two pointers to strings, where the pointers might be NULL.
//	*/
func safe_strcmp(tls *libc.TLS, a uintptr, b uintptr) (r int32) {
	if a == b {
		return 0
	}
	if a == uintptr(0) {
		return -int32(1)
	}
	if b == uintptr(0) {
		return int32(1)
	}
	return libc.Xstrcmp(tls, a, b)
}

// C documentation
//
//	/*
//	** Return TRUE if string z[] matches glob pattern zGlob[].
//	** Return FALSE if the pattern does not match.
//	**
//	** Globbing rules:
//	**
//	**      '*'       Matches any sequence of zero or more characters.
//	**
//	**      '?'       Matches exactly one character.
//	**
//	**     [...]      Matches one character from the enclosed list of
//	**                characters.
//	**
//	**     [^...]     Matches one character not in the enclosed list.
//	**
//	**      '#'       Matches any sequence of one or more digits with an
//	**                optional + or - sign in front
//	*/
func strglob(tls *libc.TLS, zGlob uintptr, z uintptr) (r int32) {
	var c, c2, invert, prior_c, seen, v1, v3 int32
	var v2, v4, v5 uintptr
	var v6 bool
	_, _, _, _, _, _, _, _, _, _, _ = c, c2, invert, prior_c, seen, v1, v2, v3, v4, v5, v6
	for {
		v2 = zGlob
		zGlob = zGlob + 1
		v1 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(v2)))
		c = v1
		if !(v1 != 0) {
			break
		}
		if c == int32('*') {
			for {
				v4 = zGlob
				zGlob = zGlob + 1
				v3 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(v4)))
				c = v3
				if !(v3 == int32('*') || c == int32('?')) {
					break
				}
				if v6 = c == int32('?'); v6 {
					v5 = z
					z = z + 1
				}
				if v6 && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(v5))) == 0 {
					return 0
				}
			}
			if c == 0 {
				return int32(1)
			} else {
				if c == int32('[') {
					for *(*uint8)(unsafe.Pointer(z)) != 0 && strglob(tls, zGlob-uintptr(1), z) != 0 {
						z = z + 1
					}
					return libc.BoolInt32(libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z))) != 0)
				}
			}
			for {
				v2 = z
				z = z + 1
				v1 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(v2)))
				c2 = v1
				if !(v1 != 0) {
					break
				}
				for c2 != c {
					v4 = z
					z = z + 1
					c2 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(v4)))
					if c2 == 0 {
						return 0
					}
				}
				if strglob(tls, zGlob, z) != 0 {
					return int32(1)
				}
			}
			return 0
		} else {
			if c == int32('?') {
				v2 = z
				z = z + 1
				if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(v2))) == 0 {
					return 0
				}
			} else {
				if c == int32('[') {
					prior_c = 0
					seen = 0
					invert = 0
					v2 = z
					z = z + 1
					c = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(v2)))
					if c == 0 {
						return 0
					}
					v2 = zGlob
					zGlob = zGlob + 1
					c2 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(v2)))
					if c2 == int32('^') {
						invert = int32(1)
						v2 = zGlob
						zGlob = zGlob + 1
						c2 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(v2)))
					}
					if c2 == int32(']') {
						if c == int32(']') {
							seen = int32(1)
						}
						v2 = zGlob
						zGlob = zGlob + 1
						c2 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(v2)))
					}
					for c2 != 0 && c2 != int32(']') {
						if c2 == int32('-') && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(zGlob))) != int32(']') && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(zGlob))) != 0 && prior_c > 0 {
							v2 = zGlob
							zGlob = zGlob + 1
							c2 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(v2)))
							if c >= prior_c && c <= c2 {
								seen = int32(1)
							}
							prior_c = 0
						} else {
							if c == c2 {
								seen = int32(1)
							}
							prior_c = c2
						}
						v2 = zGlob
						zGlob = zGlob + 1
						c2 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(v2)))
					}
					if c2 == 0 || seen^invert == 0 {
						return 0
					}
				} else {
					if c == int32('#') {
						if (libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z))) == int32('-') || libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z))) == int32('+')) && libc.BoolInt32(uint32(*(*uint8)(unsafe.Pointer(z + 1)))-uint32('0') < uint32(10)) != 0 {
							z = z + 1
						}
						if !(libc.BoolInt32(uint32(*(*uint8)(unsafe.Pointer(z)))-libc.Uint32FromUint8('0') < libc.Uint32FromInt32(10)) != 0) {
							return 0
						}
						z = z + 1
						for libc.BoolInt32(uint32(*(*uint8)(unsafe.Pointer(z)))-uint32('0') < uint32(10)) != 0 {
							z = z + 1
						}
					} else {
						v2 = z
						z = z + 1
						if c != libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(v2))) {
							return 0
						}
					}
				}
			}
		}
	}
	return libc.BoolInt32(libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z))) == 0)
}

// C documentation
//
//	/*
//	** Close output stream pOut if it is not stdout or stderr
//	*/
func maybeClose(tls *libc.TLS, pOut uintptr) {
	if pOut != libc.Xstdout && pOut != libc.Xstderr {
		libc.Xfclose(tls, pOut)
	}
}

// C documentation
//
//	/*
//	** Print an error message
//	*/
func errorMessage(tls *libc.TLS, zFormat uintptr, va uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var ap va_list
	var zMsg uintptr
	var _ /* zPrefix at bp+0 */ [30]uint8
	_, _ = ap, zMsg
	ap = va
	zMsg = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	libsqlite3.Xsqlite3_snprintf(tls, int32(30), bp, __ccgo_ts+8, libc.VaList(bp+40, uintptr(unsafe.Pointer(&g))+64))
	if g.FpLog != 0 {
		printWithPrefix(tls, g.FpLog, bp, zMsg)
		libc.Xfflush(tls, g.FpLog)
	}
	if g.FpErrLog != 0 && safe_strcmp(tls, g.FzErrLog, g.FzLog) != 0 {
		printWithPrefix(tls, g.FpErrLog, bp, zMsg)
		libc.Xfflush(tls, g.FpErrLog)
	}
	libsqlite3.Xsqlite3_free(tls, zMsg)
	g.FnError = g.FnError + 1
}

// C documentation
//
//	/*
//	** Print an error message and then quit.
//	*/
func fatalError(tls *libc.TLS, zFormat uintptr, va uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var ap va_list
	var nTry, v1 int32
	var zMsg uintptr
	var v2 bool
	var _ /* zPrefix at bp+0 */ [30]uint8
	_, _, _, _, _ = ap, nTry, zMsg, v1, v2
	ap = va
	zMsg = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	libsqlite3.Xsqlite3_snprintf(tls, int32(30), bp, __ccgo_ts+19, libc.VaList(bp+40, uintptr(unsafe.Pointer(&g))+64))
	if g.FpLog != 0 {
		printWithPrefix(tls, g.FpLog, bp, zMsg)
		libc.Xfflush(tls, g.FpLog)
		maybeClose(tls, g.FpLog)
	}
	if g.FpErrLog != 0 && safe_strcmp(tls, g.FzErrLog, g.FzLog) != 0 {
		printWithPrefix(tls, g.FpErrLog, bp, zMsg)
		libc.Xfflush(tls, g.FpErrLog)
		maybeClose(tls, g.FpErrLog)
	}
	libsqlite3.Xsqlite3_free(tls, zMsg)
	if g.Fdb != 0 {
		nTry = 0
		g.FiTimeout = 0
		for {
			if v2 = trySql(tls, __ccgo_ts+30, 0) == int32(SQLITE_BUSY); v2 {
				v1 = nTry
				nTry = nTry + 1
			}
			if !(v2 && v1 < int32(100)) {
				break
			}
			libsqlite3.Xsqlite3_sleep(tls, int32(10))
		}
	}
	libsqlite3.Xsqlite3_close(tls, g.Fdb)
	libc.Xexit(tls, int32(1))
}

// C documentation
//
//	/*
//	** Print a log message
//	*/
func logMessage(tls *libc.TLS, zFormat uintptr, va uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var ap va_list
	var zMsg uintptr
	var _ /* zPrefix at bp+0 */ [30]uint8
	_, _ = ap, zMsg
	ap = va
	zMsg = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	libsqlite3.Xsqlite3_snprintf(tls, int32(30), bp, __ccgo_ts+60, libc.VaList(bp+40, uintptr(unsafe.Pointer(&g))+64))
	if g.FpLog != 0 {
		printWithPrefix(tls, g.FpLog, bp, zMsg)
		libc.Xfflush(tls, g.FpLog)
	}
	libsqlite3.Xsqlite3_free(tls, zMsg)
}

// C documentation
//
//	/*
//	** Return the length of a string omitting trailing whitespace
//	*/
func clipLength(tls *libc.TLS, z uintptr) (r int32) {
	var n, v1, v2 int32
	var v4 bool
	_, _, _, _ = n, v1, v2, v4
	n = libc.Int32FromUint64(libc.Xstrlen(tls, z))
	for {
		if v4 = n > 0; v4 {
			v1 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z + uintptr(n-int32(1)))))
			v2 = libc.BoolInt32(v1 == int32(' ') || libc.Uint32FromInt32(v1)-uint32('\t') < uint32(5))
			goto _3
		_3:
		}
		if !(v4 && v2 != 0) {
			break
		}
		n = n - 1
	}
	return n
}

// C documentation
//
//	/*
//	** Auxiliary SQL function to return the name of the VFS
//	*/
func vfsNameFunc(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var db uintptr
	var _ /* zVfs at bp+0 */ uintptr
	_ = db
	db = libsqlite3.Xsqlite3_context_db_handle(tls, context)
	*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
	_ = argc
	_ = argv
	libsqlite3.Xsqlite3_file_control(tls, db, __ccgo_ts+65, int32(SQLITE_FCNTL_VFSNAME), bp)
	if *(*uintptr)(unsafe.Pointer(bp)) != 0 {
		libsqlite3.Xsqlite3_result_text(tls, context, *(*uintptr)(unsafe.Pointer(bp)), -int32(1), __ccgo_fp(libsqlite3.Xsqlite3_free))
	}
}

// C documentation
//
//	/*
//	** Busy handler with a g.iTimeout-millisecond timeout
//	*/
func busyHandler(tls *libc.TLS, pCD uintptr, count int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	_ = pCD
	if count*int32(10) > g.FiTimeout {
		if g.FiTimeout > 0 {
			errorMessage(tls, __ccgo_ts+70, libc.VaList(bp+8, g.FiTimeout))
		}
		return 0
	}
	libsqlite3.Xsqlite3_sleep(tls, int32(10))
	return int32(1)
}

// C documentation
//
//	/*
//	** SQL Trace callback
//	*/
func sqlTraceCallback(tls *libc.TLS, NotUsed1 uintptr, zSql uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	_ = NotUsed1
	logMessage(tls, __ccgo_ts+89, libc.VaList(bp+8, clipLength(tls, zSql), zSql))
}

// C documentation
//
//	/*
//	** SQL error log callback
//	*/
func sqlErrorCallback(tls *libc.TLS, pArg uintptr, iErrCode int32, zMsg uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	_ = pArg
	if iErrCode == int32(SQLITE_ERROR) && g.FbIgnoreSqlErrors != 0 {
		return
	}
	if iErrCode&int32(0xff) == int32(SQLITE_SCHEMA) && g.FiTrace < int32(3) {
		return
	}
	if g.FiTimeout == 0 && iErrCode&int32(0xff) == int32(SQLITE_BUSY) && g.FiTrace < int32(3) {
		return
	}
	if iErrCode&int32(0xff) == int32(SQLITE_NOTICE) {
		logMessage(tls, __ccgo_ts+96, libc.VaList(bp+8, zMsg))
	} else {
		errorMessage(tls, __ccgo_ts+106, libc.VaList(bp+8, iErrCode, zMsg))
	}
}

// C documentation
//
//	/*
//	** Prepare an SQL statement.  Issue a fatal error if unable.
//	*/
func prepareSql(tls *libc.TLS, zFormat uintptr, va uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var ap va_list
	var rc int32
	var zSql uintptr
	var _ /* pStmt at bp+0 */ uintptr
	_, _, _ = ap, rc, zSql
	*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
	ap = va
	zSql = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	rc = libsqlite3.Xsqlite3_prepare_v2(tls, g.Fdb, zSql, -int32(1), bp, uintptr(0))
	if rc != SQLITE_OK {
		libsqlite3.Xsqlite3_finalize(tls, *(*uintptr)(unsafe.Pointer(bp)))
		fatalError(tls, __ccgo_ts+122, libc.VaList(bp+16, libsqlite3.Xsqlite3_errmsg(tls, g.Fdb), zSql))
	}
	libsqlite3.Xsqlite3_free(tls, zSql)
	return *(*uintptr)(unsafe.Pointer(bp))
}

// C documentation
//
//	/*
//	** Run arbitrary SQL.  Issue a fatal error on failure.
//	*/
func runSql(tls *libc.TLS, zFormat uintptr, va uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var ap va_list
	var rc int32
	var zSql uintptr
	_, _, _ = ap, rc, zSql
	ap = va
	zSql = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	rc = libsqlite3.Xsqlite3_exec(tls, g.Fdb, zSql, uintptr(0), uintptr(0), uintptr(0))
	if rc != SQLITE_OK {
		fatalError(tls, __ccgo_ts+122, libc.VaList(bp+8, libsqlite3.Xsqlite3_errmsg(tls, g.Fdb), zSql))
	}
	libsqlite3.Xsqlite3_free(tls, zSql)
}

// C documentation
//
//	/*
//	** Try to run arbitrary SQL.  Return success code.
//	*/
func trySql(tls *libc.TLS, zFormat uintptr, va uintptr) (r int32) {
	var ap va_list
	var rc int32
	var zSql uintptr
	_, _, _ = ap, rc, zSql
	ap = va
	zSql = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	rc = libsqlite3.Xsqlite3_exec(tls, g.Fdb, zSql, uintptr(0), uintptr(0), uintptr(0))
	libsqlite3.Xsqlite3_free(tls, zSql)
	return rc
}

// C documentation
//
//	/* Structure for holding an arbitrary length string
//	*/
type String = struct {
	Fz      uintptr
	Fn      int32
	FnAlloc int32
}

// C documentation
//
//	/* Free a string */
func stringFree(tls *libc.TLS, p uintptr) {
	if (*String)(unsafe.Pointer(p)).Fz != 0 {
		libsqlite3.Xsqlite3_free(tls, (*String)(unsafe.Pointer(p)).Fz)
	}
	libc.Xmemset(tls, p, 0, uint64(16))
}

// C documentation
//
//	/* Append n bytes of text to a string.  If n<0 append the entire string. */
func stringAppend(tls *libc.TLS, p uintptr, z uintptr, n int32) {
	var nAlloc int32
	var zNew uintptr
	_, _ = nAlloc, zNew
	if n < 0 {
		n = libc.Int32FromUint64(libc.Xstrlen(tls, z))
	}
	if (*String)(unsafe.Pointer(p)).Fn+n >= (*String)(unsafe.Pointer(p)).FnAlloc {
		nAlloc = (*String)(unsafe.Pointer(p)).FnAlloc*int32(2) + n + int32(100)
		zNew = libsqlite3.Xsqlite3_realloc(tls, (*String)(unsafe.Pointer(p)).Fz, nAlloc)
		if zNew == uintptr(0) {
			fatalError(tls, __ccgo_ts+129, 0)
		}
		(*String)(unsafe.Pointer(p)).Fz = zNew
		(*String)(unsafe.Pointer(p)).FnAlloc = nAlloc
	}
	libc.Xmemcpy(tls, (*String)(unsafe.Pointer(p)).Fz+uintptr((*String)(unsafe.Pointer(p)).Fn), z, libc.Uint64FromInt32(n))
	*(*int32)(unsafe.Pointer(p + 8)) += n
	*(*uint8)(unsafe.Pointer((*String)(unsafe.Pointer(p)).Fz + uintptr((*String)(unsafe.Pointer(p)).Fn))) = uint8(0)
}

// C documentation
//
//	/* Reset a string to an empty string */
func stringReset(tls *libc.TLS, p uintptr) {
	if (*String)(unsafe.Pointer(p)).Fz == uintptr(0) {
		stringAppend(tls, p, __ccgo_ts+143, int32(1))
	}
	(*String)(unsafe.Pointer(p)).Fn = 0
	*(*uint8)(unsafe.Pointer((*String)(unsafe.Pointer(p)).Fz)) = uint8(0)
}

// C documentation
//
//	/* Append a new token onto the end of the string */
func stringAppendTerm(tls *libc.TLS, p uintptr, z uintptr) {
	var i, v2, v3 int32
	var v5 bool
	_, _, _, _ = i, v2, v3, v5
	if (*String)(unsafe.Pointer(p)).Fn != 0 {
		stringAppend(tls, p, __ccgo_ts+143, int32(1))
	}
	if z == uintptr(0) {
		stringAppend(tls, p, __ccgo_ts+145, int32(3))
		return
	}
	i = 0
	for {
		if v5 = *(*uint8)(unsafe.Pointer(z + uintptr(i))) != 0; v5 {
			v2 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z + uintptr(i))))
			v3 = libc.BoolInt32(v2 == int32(' ') || libc.Uint32FromInt32(v2)-uint32('\t') < uint32(5))
			goto _4
		_4:
		}
		if !(v5 && !(v3 != 0)) {
			break
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	if i > 0 && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z + uintptr(i)))) == 0 {
		stringAppend(tls, p, z, i)
		return
	}
	stringAppend(tls, p, __ccgo_ts+149, int32(1))
	for *(*uint8)(unsafe.Pointer(z)) != 0 {
		i = 0
		for {
			if !(*(*uint8)(unsafe.Pointer(z + uintptr(i))) != 0 && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z + uintptr(i)))) != int32('\'')) {
				break
			}
			goto _6
		_6:
			;
			i = i + 1
		}
		if *(*uint8)(unsafe.Pointer(z + uintptr(i))) != 0 {
			stringAppend(tls, p, z, i+int32(1))
			stringAppend(tls, p, __ccgo_ts+149, int32(1))
			z = z + uintptr(i+int32(1))
		} else {
			stringAppend(tls, p, z, i)
			break
		}
	}
	stringAppend(tls, p, __ccgo_ts+149, int32(1))
}

// C documentation
//
//	/*
//	** Callback function for evalSql()
//	*/
func evalCallback(tls *libc.TLS, pCData uintptr, argc int32, argv uintptr, azCol uintptr) (r int32) {
	var i int32
	var p uintptr
	_, _ = i, p
	p = pCData
	_ = azCol
	i = 0
	for {
		if !(i < argc) {
			break
		}
		stringAppendTerm(tls, p, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)))
		goto _1
	_1:
		;
		i = i + 1
	}
	return 0
}

// C documentation
//
//	/*
//	** Run arbitrary SQL and record the results in an output string
//	** given by the first parameter.
//	*/
func evalSql(tls *libc.TLS, p uintptr, zFormat uintptr, va uintptr) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var ap va_list
	var rc int32
	var zSql uintptr
	var _ /* zErr at bp+8 */ [30]uint8
	var _ /* zErrMsg at bp+0 */ uintptr
	_, _, _ = ap, rc, zSql
	*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
	ap = va
	zSql = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	rc = libsqlite3.Xsqlite3_exec(tls, g.Fdb, zSql, __ccgo_fp(evalCallback), p, bp)
	libsqlite3.Xsqlite3_free(tls, zSql)
	if rc != 0 {
		libsqlite3.Xsqlite3_snprintf(tls, int32(30), bp+8, __ccgo_ts+151, libc.VaList(bp+48, rc))
		stringAppendTerm(tls, p, bp+8)
		if *(*uintptr)(unsafe.Pointer(bp)) != 0 {
			stringAppendTerm(tls, p, *(*uintptr)(unsafe.Pointer(bp)))
			libsqlite3.Xsqlite3_free(tls, *(*uintptr)(unsafe.Pointer(bp)))
		}
	}
	return rc
}

// C documentation
//
//	/*
//	** Auxiliary SQL function to recursively evaluate SQL.
//	*/
func evalFunc(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var db, zSql uintptr
	var rc int32
	var _ /* res at bp+0 */ String
	var _ /* zErrMsg at bp+16 */ uintptr
	_, _, _ = db, rc, zSql
	db = libsqlite3.Xsqlite3_context_db_handle(tls, context)
	zSql = libsqlite3.Xsqlite3_value_text(tls, *(*uintptr)(unsafe.Pointer(argv)))
	*(*uintptr)(unsafe.Pointer(bp + 16)) = uintptr(0)
	_ = argc
	libc.Xmemset(tls, bp, 0, uint64(16))
	rc = libsqlite3.Xsqlite3_exec(tls, db, zSql, __ccgo_fp(evalCallback), bp, bp+16)
	if *(*uintptr)(unsafe.Pointer(bp + 16)) != 0 {
		libsqlite3.Xsqlite3_result_error(tls, context, *(*uintptr)(unsafe.Pointer(bp + 16)), -int32(1))
		libsqlite3.Xsqlite3_free(tls, *(*uintptr)(unsafe.Pointer(bp + 16)))
	} else {
		if rc != 0 {
			libsqlite3.Xsqlite3_result_error_code(tls, context, rc)
		} else {
			libsqlite3.Xsqlite3_result_text(tls, context, (*(*String)(unsafe.Pointer(bp))).Fz, -int32(1), uintptr(-libc.Int32FromInt32(1)))
		}
	}
	stringFree(tls, bp)
}

// C documentation
//
//	/*
//	** Look up the next task for client iClient in the database.
//	** Return the task script and the task number and mark that
//	** task as being under way.
//	*/
func startScript(tls *libc.TLS, iClient int32, pzScript uintptr, pTaskId uintptr, pzTaskName uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var n, rc, taskId, totalTime, v1 int32
	var pStmt uintptr
	_, _, _, _, _, _ = n, pStmt, rc, taskId, totalTime, v1
	pStmt = uintptr(0)
	totalTime = 0
	*(*uintptr)(unsafe.Pointer(pzScript)) = uintptr(0)
	g.FiTimeout = 0
	for int32(1) != 0 {
		rc = trySql(tls, __ccgo_ts+161, 0)
		if rc == int32(SQLITE_BUSY) {
			libsqlite3.Xsqlite3_sleep(tls, int32(10))
			totalTime = totalTime + int32(10)
			continue
		}
		if rc != SQLITE_OK {
			fatalError(tls, __ccgo_ts+177, libc.VaList(bp+8, libsqlite3.Xsqlite3_errmsg(tls, g.Fdb)))
		}
		if g.FnError != 0 || g.FnTest != 0 {
			runSql(tls, __ccgo_ts+196, libc.VaList(bp+8, g.FnError, g.FnTest))
			g.FnError = 0
			g.FnTest = 0
		}
		pStmt = prepareSql(tls, __ccgo_ts+249, libc.VaList(bp+8, iClient))
		rc = libsqlite3.Xsqlite3_step(tls, pStmt)
		libsqlite3.Xsqlite3_finalize(tls, pStmt)
		if rc == int32(SQLITE_ROW) {
			runSql(tls, __ccgo_ts+295, libc.VaList(bp+8, iClient))
			g.FiTimeout = int32(DEFAULT_TIMEOUT)
			runSql(tls, __ccgo_ts+326, 0)
			return int32(SQLITE_DONE)
		}
		pStmt = prepareSql(tls, __ccgo_ts+346, libc.VaList(bp+8, iClient))
		rc = libsqlite3.Xsqlite3_step(tls, pStmt)
		if rc == int32(SQLITE_ROW) {
			n = libsqlite3.Xsqlite3_column_bytes(tls, pStmt, 0)
			*(*uintptr)(unsafe.Pointer(pzScript)) = libsqlite3.Xsqlite3_malloc(tls, n+int32(1))
			libc.Xstrcpy(tls, *(*uintptr)(unsafe.Pointer(pzScript)), libsqlite3.Xsqlite3_column_text(tls, pStmt, 0))
			v1 = libsqlite3.Xsqlite3_column_int(tls, pStmt, int32(1))
			taskId = v1
			*(*int32)(unsafe.Pointer(pTaskId)) = v1
			*(*uintptr)(unsafe.Pointer(pzTaskName)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+438, libc.VaList(bp+8, libsqlite3.Xsqlite3_column_text(tls, pStmt, int32(2))))
			libsqlite3.Xsqlite3_finalize(tls, pStmt)
			runSql(tls, __ccgo_ts+441, libc.VaList(bp+8, taskId))
			g.FiTimeout = int32(DEFAULT_TIMEOUT)
			runSql(tls, __ccgo_ts+326, 0)
			return SQLITE_OK
		}
		libsqlite3.Xsqlite3_finalize(tls, pStmt)
		if rc == int32(SQLITE_DONE) {
			if totalTime > int32(30000) {
				errorMessage(tls, __ccgo_ts+524, 0)
				runSql(tls, __ccgo_ts+573, libc.VaList(bp+8, iClient))
				libsqlite3.Xsqlite3_close(tls, g.Fdb)
				libc.Xexit(tls, int32(1))
			}
			for trySql(tls, __ccgo_ts+613, 0) == int32(SQLITE_BUSY) {
				libsqlite3.Xsqlite3_sleep(tls, int32(10))
				totalTime = totalTime + int32(10)
			}
			libsqlite3.Xsqlite3_sleep(tls, int32(100))
			totalTime = totalTime + int32(100)
			continue
		}
		fatalError(tls, __ccgo_ts+438, libc.VaList(bp+8, libsqlite3.Xsqlite3_errmsg(tls, g.Fdb)))
	}
	g.FiTimeout = int32(DEFAULT_TIMEOUT)
	return r
}

// C documentation
//
//	/*
//	** Mark a script as having finished.   Remove the CLIENT table entry
//	** if bShutdown is true.
//	*/
func finishScript(tls *libc.TLS, iClient int32, taskId int32, bShutdown int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	runSql(tls, __ccgo_ts+620, libc.VaList(bp+8, taskId))
	if bShutdown != 0 {
		runSql(tls, __ccgo_ts+295, libc.VaList(bp+8, iClient))
	}
	return SQLITE_OK
}

// C documentation
//
//	/*
//	** Start up a client process for iClient, if it is not already
//	** running.  If the client is already running, then this routine
//	** is a no-op.
//	*/
func startClient(tls *libc.TLS, iClient int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var rc int32
	var zSys uintptr
	_, _ = rc, zSys
	runSql(tls, __ccgo_ts+701, libc.VaList(bp+8, iClient))
	if libsqlite3.Xsqlite3_changes(tls, g.Fdb) != 0 {
		zSys = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+743, libc.VaList(bp+8, g.Fargv0, g.FzDbFile, iClient, g.FiTrace))
		if g.FbSqlTrace != 0 {
			zSys = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+774, libc.VaList(bp+8, zSys))
		}
		if g.FbSync != 0 {
			zSys = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+788, libc.VaList(bp+8, zSys))
		}
		if g.FzVfs != 0 {
			zSys = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+798, libc.VaList(bp+8, zSys, g.FzVfs))
		}
		if g.FiTrace >= int32(2) {
			logMessage(tls, __ccgo_ts+812, libc.VaList(bp+8, zSys))
		}
		zSys = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+825, libc.VaList(bp+8, zSys))
		rc = libc.Xsystem(tls, zSys)
		if rc != 0 {
			errorMessage(tls, __ccgo_ts+830, libc.VaList(bp+8, rc))
		}
		libsqlite3.Xsqlite3_free(tls, zSys)
	}
}

// C documentation
//
//	/*
//	** Read the entire content of a file into memory
//	*/
func readFile(tls *libc.TLS, zFilename uintptr) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var in, z uintptr
	var sz int64
	_, _, _ = in, sz, z
	in = libc.Xfopen(tls, zFilename, __ccgo_ts+864)
	if in == uintptr(0) {
		fatalError(tls, __ccgo_ts+867, libc.VaList(bp+8, zFilename))
	}
	libc.Xfseek(tls, in, 0, int32(2))
	sz = libc.Xftell(tls, in)
	libc.Xrewind(tls, in)
	z = libsqlite3.Xsqlite3_malloc(tls, int32(sz+int64(1)))
	sz = libc.Int64FromUint64(libc.Xfread(tls, z, uint64(1), libc.Uint64FromInt64(sz), in))
	*(*uint8)(unsafe.Pointer(z + uintptr(sz))) = uint8(0)
	libc.Xfclose(tls, in)
	return z
}

// C documentation
//
//	/*
//	** Return the length of the next token.
//	*/
func tokenLength(tls *libc.TLS, z uintptr, pnLine uintptr) (r int32) {
	var c, c1, delim, inC, n, v1, v2, v4 int32
	var v16 bool
	_, _, _, _, _, _, _, _, _ = c, c1, delim, inC, n, v1, v16, v2, v4
	n = 0
	v1 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z)))
	v2 = libc.BoolInt32(v1 == int32(' ') || libc.Uint32FromInt32(v1)-uint32('\t') < uint32(5))
	goto _3
_3:
	;
	if v2 != 0 || libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z))) == int32('/') && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z + 1))) == int32('*') {
		inC = 0
		if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z))) == int32('/') {
			inC = int32(1)
			n = int32(2)
		}
		for {
			v2 = n
			n = n + 1
			v1 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z + uintptr(v2))))
			c = v1
			if !(v1 != 0) {
				break
			}
			if c == int32('\n') {
				*(*int32)(unsafe.Pointer(pnLine)) = *(*int32)(unsafe.Pointer(pnLine)) + 1
			}
			v1 = libc.Int32FromUint8(libc.Uint8FromInt32(c))
			v2 = libc.BoolInt32(v1 == int32(' ') || libc.Uint32FromInt32(v1)-uint32('\t') < uint32(5))
			goto _8
		_8:
			if v2 != 0 {
				continue
			}
			if inC != 0 && c == int32('*') && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z + uintptr(n)))) == int32('/') {
				n = n + 1
				inC = 0
			} else {
				if !(inC != 0) && c == int32('/') && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z + uintptr(n)))) == int32('*') {
					n = n + 1
					inC = int32(1)
				} else {
					if !(inC != 0) {
						break
					}
				}
			}
		}
		n = n - 1
	} else {
		if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z))) == int32('-') && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z + 1))) == int32('-') {
			n = int32(2)
			for {
				if !(*(*uint8)(unsafe.Pointer(z + uintptr(n))) != 0 && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z + uintptr(n)))) != int32('\n')) {
					break
				}
				goto _9
			_9:
				;
				n = n + 1
			}
			if *(*uint8)(unsafe.Pointer(z + uintptr(n))) != 0 {
				*(*int32)(unsafe.Pointer(pnLine)) = *(*int32)(unsafe.Pointer(pnLine)) + 1
				n = n + 1
			}
		} else {
			if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z))) == int32('"') || libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z))) == int32('\'') {
				delim = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z)))
				n = int32(1)
				for {
					if !(*(*uint8)(unsafe.Pointer(z + uintptr(n))) != 0) {
						break
					}
					if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z + uintptr(n)))) == int32('\n') {
						*(*int32)(unsafe.Pointer(pnLine)) = *(*int32)(unsafe.Pointer(pnLine)) + 1
					}
					if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z + uintptr(n)))) == delim {
						n = n + 1
						if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z + uintptr(n+int32(1))))) != delim {
							break
						}
					}
					goto _10
				_10:
					;
					n = n + 1
				}
			} else {
				n = int32(1)
				for {
					v1 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z + uintptr(n))))
					c1 = v1
					if v16 = v1 != 0; v16 {
						v2 = libc.Int32FromUint8(libc.Uint8FromInt32(c1))
						v4 = libc.BoolInt32(v2 == int32(' ') || libc.Uint32FromInt32(v2)-uint32('\t') < uint32(5))
						goto _15
					_15:
					}
					if !(v16 && !(v4 != 0) && c1 != int32('"') && c1 != int32('\'') && c1 != int32(';')) {
						break
					}
					goto _11
				_11:
					;
					n = n + 1
				}
			}
		}
	}
	return n
}

// C documentation
//
//	/*
//	** Copy a single token into a string buffer.
//	*/
func extractToken(tls *libc.TLS, zIn uintptr, nIn int32, zOut uintptr, nOut int32) (r int32) {
	var i, v2, v3 int32
	var v5 bool
	_, _, _, _ = i, v2, v3, v5
	if nIn <= 0 {
		*(*uint8)(unsafe.Pointer(zOut)) = uint8(0)
		return 0
	}
	i = 0
	for {
		if v5 = i < nIn && i < nOut-int32(1); v5 {
			v2 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(zIn + uintptr(i))))
			v3 = libc.BoolInt32(v2 == int32(' ') || libc.Uint32FromInt32(v2)-uint32('\t') < uint32(5))
			goto _4
		_4:
		}
		if !(v5 && !(v3 != 0)) {
			break
		}
		*(*uint8)(unsafe.Pointer(zOut + uintptr(i))) = *(*uint8)(unsafe.Pointer(zIn + uintptr(i)))
		goto _1
	_1:
		;
		i = i + 1
	}
	*(*uint8)(unsafe.Pointer(zOut + uintptr(i))) = uint8(0)
	return i
}

// C documentation
//
//	/*
//	** Find the number of characters up to the start of the next "--end" token.
//	*/
func findEnd(tls *libc.TLS, z uintptr, pnLine uintptr) (r int32) {
	var n, v1, v2 int32
	var v4, v5 bool
	_, _, _, _, _ = n, v1, v2, v4, v5
	n = 0
	for {
		if v5 = *(*uint8)(unsafe.Pointer(z + uintptr(n))) != 0; v5 {
			if v4 = libc.Xstrncmp(tls, z+uintptr(n), __ccgo_ts+896, uint64(5)) != 0; !v4 {
				v1 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z + uintptr(n+int32(5)))))
				v2 = libc.BoolInt32(v1 == int32(' ') || libc.Uint32FromInt32(v1)-uint32('\t') < uint32(5))
				goto _3
			_3:
			}
		}
		if !(v5 && (v4 || !(v2 != 0))) {
			break
		}
		n = n + tokenLength(tls, z+uintptr(n), pnLine)
	}
	return n
}

// C documentation
//
//	/*
//	** Find the number of characters up to the first character past the
//	** of the next "--endif"  or "--else" token. Nested --if commands are
//	** also skipped.
//	*/
func findEndif(tls *libc.TLS, z uintptr, stopAtElse int32, pnLine uintptr) (r int32) {
	var len1, n, skip, v1, v2, v5, v6 int32
	var v4, v8, v9 bool
	_, _, _, _, _, _, _, _, _, _ = len1, n, skip, v1, v2, v4, v5, v6, v8, v9
	n = 0
	for *(*uint8)(unsafe.Pointer(z + uintptr(n))) != 0 {
		len1 = tokenLength(tls, z+uintptr(n), pnLine)
		if v4 = libc.Xstrncmp(tls, z+uintptr(n), __ccgo_ts+902, uint64(7)) == 0; v4 {
			v1 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z + uintptr(n+int32(7)))))
			v2 = libc.BoolInt32(v1 == int32(' ') || libc.Uint32FromInt32(v1)-uint32('\t') < uint32(5))
			goto _3
		_3:
		}
		if v9 = v4 && v2 != 0; !v9 {
			if v8 = stopAtElse != 0 && libc.Xstrncmp(tls, z+uintptr(n), __ccgo_ts+910, uint64(6)) == 0; v8 {
				v5 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z + uintptr(n+int32(6)))))
				v6 = libc.BoolInt32(v5 == int32(' ') || libc.Uint32FromInt32(v5)-uint32('\t') < uint32(5))
				goto _7
			_7:
			}
		}
		if v9 || v8 && v6 != 0 {
			return n + len1
		}
		if v4 = libc.Xstrncmp(tls, z+uintptr(n), __ccgo_ts+917, uint64(4)) == 0; v4 {
			v1 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z + uintptr(n+int32(4)))))
			v2 = libc.BoolInt32(v1 == int32(' ') || libc.Uint32FromInt32(v1)-uint32('\t') < uint32(5))
			goto _12
		_12:
		}
		if v4 && v2 != 0 {
			skip = findEndif(tls, z+uintptr(n)+uintptr(len1), 0, pnLine)
			n = n + (skip + len1)
		} else {
			n = n + len1
		}
	}
	return n
}

// C documentation
//
//	/*
//	** Wait for a client process to complete all its tasks
//	*/
func waitForClient(tls *libc.TLS, iClient int32, iTimeout int32, zErrPrefix uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var pStmt uintptr
	var rc, v1 int32
	_, _, _ = pStmt, rc, v1
	if iClient > 0 {
		pStmt = prepareSql(tls, __ccgo_ts+922, libc.VaList(bp+8, iClient))
	} else {
		pStmt = prepareSql(tls, __ccgo_ts+1018, 0)
	}
	g.FiTimeout = 0
	for {
		v1 = libsqlite3.Xsqlite3_step(tls, pStmt)
		rc = v1
		if !((v1 == int32(SQLITE_BUSY) || rc == int32(SQLITE_ROW)) && iTimeout > 0) {
			break
		}
		libsqlite3.Xsqlite3_reset(tls, pStmt)
		libsqlite3.Xsqlite3_sleep(tls, int32(50))
		iTimeout = iTimeout - int32(50)
	}
	libsqlite3.Xsqlite3_finalize(tls, pStmt)
	g.FiTimeout = int32(DEFAULT_TIMEOUT)
	if rc != int32(SQLITE_DONE) {
		if zErrPrefix == uintptr(0) {
			zErrPrefix = __ccgo_ts + 1099
		}
		if iClient > 0 {
			errorMessage(tls, __ccgo_ts+1100, libc.VaList(bp+8, zErrPrefix, iClient))
		} else {
			errorMessage(tls, __ccgo_ts+1132, libc.VaList(bp+8, zErrPrefix))
		}
	}
}

// C documentation
//
//	/* Return a pointer to the tail of a filename
//	*/
func filenameTail(tls *libc.TLS, z uintptr) (r uintptr) {
	var i, j, v2 int32
	_, _, _ = i, j, v2
	v2 = libc.Int32FromInt32(0)
	j = v2
	i = v2
	for {
		if !(*(*uint8)(unsafe.Pointer(z + uintptr(i))) != 0) {
			break
		}
		if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z + uintptr(i)))) == int32('/') {
			j = i + int32(1)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return z + uintptr(j)
}

// C documentation
//
//	/*
//	** Interpret zArg as a boolean value.  Return either 0 or 1.
//	*/
func booleanValue(tls *libc.TLS, zArg uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i int32
	_ = i
	if zArg == uintptr(0) {
		return 0
	}
	i = 0
	for {
		if !(libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(zArg + uintptr(i)))) >= int32('0') && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(zArg + uintptr(i)))) <= int32('9')) {
			break
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	if i > 0 && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(zArg + uintptr(i)))) == 0 {
		return libc.Xatoi(tls, zArg)
	}
	if libsqlite3.Xsqlite3_stricmp(tls, zArg, __ccgo_ts+1166) == 0 || libsqlite3.Xsqlite3_stricmp(tls, zArg, __ccgo_ts+1169) == 0 {
		return int32(1)
	}
	if libsqlite3.Xsqlite3_stricmp(tls, zArg, __ccgo_ts+1173) == 0 || libsqlite3.Xsqlite3_stricmp(tls, zArg, __ccgo_ts+1177) == 0 {
		return 0
	}
	errorMessage(tls, __ccgo_ts+1180, libc.VaList(bp+8, zArg))
	return 0
}

// C documentation
//
//	/* This routine exists as a convenient place to set a debugger
//	** breakpoint.
//	*/
func test_breakpoint(tls *libc.TLS) {
	libc.PostIncAtomicInt32P(uintptr(unsafe.Pointer(&cnt)), 1)
}

var cnt int32

/* Maximum number of arguments to a --command */

// C documentation
//
//	/*
//	** Run a script.
//	*/
func runScript(tls *libc.TLS, iClient int32, taskId int32, zScript uintptr, zFilename uintptr) {
	bp := tls.Alloc(1312)
	defer tls.Free(1312)
	var c, iBegin, iEnd, iNewClient, iTarget, iTimeout, ii, isGlob, j, jj, jj1, jj2, jj3, k, len1, n, nArg, prevLine, rc, rc1, v1, v2 int32
	var pStmt, zAns, zAns1, zCopy, zNewFile, zNewScript, zSql, zSql1, zTName, zTask, zToDel, v23 uintptr
	var v9 bool
	var _ /* azArg at bp+1054 */ [2][100]uint8
	var _ /* lineno at bp+0 */ int32
	var _ /* sResult at bp+8 */ String
	var _ /* zCmd at bp+24 */ [30]uint8
	var _ /* zError at bp+54 */ [1000]uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, iBegin, iEnd, iNewClient, iTarget, iTimeout, ii, isGlob, j, jj, jj1, jj2, jj3, k, len1, n, nArg, pStmt, prevLine, rc, rc1, zAns, zAns1, zCopy, zNewFile, zNewScript, zSql, zSql1, zTName, zTask, zToDel, v1, v2, v23, v9
	*(*int32)(unsafe.Pointer(bp)) = int32(1)
	prevLine = int32(1)
	ii = 0
	iBegin = 0
	libc.Xmemset(tls, bp+8, 0, uint64(16))
	stringReset(tls, bp+8)
	for {
		v1 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(zScript + uintptr(ii))))
		c = v1
		if !(v1 != 0) {
			break
		}
		prevLine = *(*int32)(unsafe.Pointer(bp))
		len1 = tokenLength(tls, zScript+uintptr(ii), bp)
		v1 = libc.Int32FromUint8(libc.Uint8FromInt32(c))
		v2 = libc.BoolInt32(v1 == int32(' ') || libc.Uint32FromInt32(v1)-uint32('\t') < uint32(5))
		goto _4
	_4:
		;
		if v2 != 0 || c == int32('/') && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(zScript + uintptr(ii+int32(1))))) == int32('*') {
			ii = ii + len1
			continue
		}
		if c != int32('-') || libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(zScript + uintptr(ii+int32(1))))) != int32('-') || !(libc.BoolInt32(uint32(*(*uint8)(unsafe.Pointer(zScript + uintptr(ii+int32(2)))))|libc.Uint32FromInt32(32)-libc.Uint32FromUint8('a') < libc.Uint32FromInt32(26)) != 0) {
			ii = ii + len1
			continue
		}
		/* Run any prior SQL before processing the new --command */
		if ii > iBegin {
			zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1202, libc.VaList(bp+1264, ii-iBegin, zScript+uintptr(iBegin)))
			evalSql(tls, bp+8, zSql, 0)
			libsqlite3.Xsqlite3_free(tls, zSql)
			iBegin = ii + len1
		}
		/* Parse the --command */
		if g.FiTrace >= int32(2) {
			logMessage(tls, __ccgo_ts+1202, libc.VaList(bp+1264, len1, zScript+uintptr(ii)))
		}
		n = extractToken(tls, zScript+uintptr(ii)+uintptr(2), len1-int32(2), bp+24, int32(30))
		nArg = 0
		for {
			if !(n < len1-int32(2) && nArg < int32(MX_ARG)) {
				break
			}
			for {
				if v9 = n < len1-int32(2); v9 {
					v1 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(zScript + uintptr(ii+int32(2)+n))))
					v2 = libc.BoolInt32(v1 == int32(' ') || libc.Uint32FromInt32(v1)-uint32('\t') < uint32(5))
					goto _8
				_8:
				}
				if !(v9 && v2 != 0) {
					break
				}
				n = n + 1
			}
			if n >= len1-int32(2) {
				break
			}
			n = n + extractToken(tls, zScript+uintptr(ii)+uintptr(2)+uintptr(n), len1-int32(2)-n, bp+1054+uintptr(nArg)*100, int32(100))
			goto _5
		_5:
			;
			nArg = nArg + 1
		}
		j = nArg
		for {
			if !(j < int32(MX_ARG)) {
				break
			}
			v1 = j
			j = j + 1
			*(*uint8)(unsafe.Pointer(bp + 1054 + uintptr(v1)*100)) = uint8(0)
			goto _10
		_10:
			;
			j = j + 1
		}
		/*
		 **  --sleep N
		 **
		 ** Pause for N milliseconds
		 */
		if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1207) == 0 {
			libsqlite3.Xsqlite3_sleep(tls, libc.Xatoi(tls, bp+1054))
		} else {
			/*
			 **   --exit N
			 **
			 ** Exit this process.  If N>0 then exit without shutting down
			 ** SQLite.  (In other words, simulate a crash.)
			 */
			if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1213) == 0 {
				rc = libc.Xatoi(tls, bp+1054)
				finishScript(tls, iClient, taskId, int32(1))
				if rc == 0 {
					libsqlite3.Xsqlite3_close(tls, g.Fdb)
				}
				libc.Xexit(tls, rc)
			} else {
				/*
				 **   --testcase NAME
				 **
				 ** Begin a new test case.  Announce in the log that the test case
				 ** has begun.
				 */
				if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1218) == 0 {
					if g.FiTrace == int32(1) {
						logMessage(tls, __ccgo_ts+1202, libc.VaList(bp+1264, len1-int32(1), zScript+uintptr(ii)))
					}
					stringReset(tls, bp+8)
				} else {
					/*
					 **   --finish
					 **
					 ** Mark the current task as having finished, even if it is not.
					 ** This can be used in conjunction with --exit to simulate a crash.
					 */
					if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1227) == 0 && iClient > 0 {
						finishScript(tls, iClient, taskId, int32(1))
					} else {
						/*
						 **  --reset
						 **
						 ** Reset accumulated results back to an empty string
						 */
						if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1234) == 0 {
							stringReset(tls, bp+8)
						} else {
							/*
							 **  --match ANSWER...
							 **
							 ** Check to see if output matches ANSWER.  Report an error if not.
							 */
							if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1240) == 0 {
								zAns = zScript + uintptr(ii)
								jj = int32(7)
								for {
									if v9 = jj < len1-int32(1); v9 {
										v1 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(zAns + uintptr(jj))))
										v2 = libc.BoolInt32(v1 == int32(' ') || libc.Uint32FromInt32(v1)-uint32('\t') < uint32(5))
										goto _15
									_15:
									}
									if !(v9 && v2 != 0) {
										break
									}
									goto _12
								_12:
									;
									jj = jj + 1
								}
								zAns = zAns + uintptr(jj)
								if len1-jj-int32(1) != (*(*String)(unsafe.Pointer(bp + 8))).Fn || libc.Xstrncmp(tls, (*(*String)(unsafe.Pointer(bp + 8))).Fz, zAns, libc.Uint64FromInt32(len1-jj-int32(1))) != 0 {
									errorMessage(tls, __ccgo_ts+1246, libc.VaList(bp+1264, prevLine, zFilename, len1-jj-int32(1), zAns, (*(*String)(unsafe.Pointer(bp + 8))).Fz))
								}
								g.FnTest = g.FnTest + 1
								stringReset(tls, bp+8)
							} else {
								/*
								 **  --glob ANSWER...
								 **  --notglob ANSWER....
								 **
								 ** Check to see if output does or does not match the glob pattern
								 ** ANSWER.
								 */
								if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1291) == 0 || libc.Xstrcmp(tls, bp+24, __ccgo_ts+1296) == 0 {
									zAns1 = zScript + uintptr(ii)
									isGlob = libc.BoolInt32(libc.Int32FromUint8((*(*[30]uint8)(unsafe.Pointer(bp + 24)))[0]) == int32('g'))
									jj1 = int32(9) - int32(3)*isGlob
									for {
										if v9 = jj1 < len1-int32(1); v9 {
											v1 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(zAns1 + uintptr(jj1))))
											v2 = libc.BoolInt32(v1 == int32(' ') || libc.Uint32FromInt32(v1)-uint32('\t') < uint32(5))
											goto _20
										_20:
										}
										if !(v9 && v2 != 0) {
											break
										}
										goto _17
									_17:
										;
										jj1 = jj1 + 1
									}
									zAns1 = zAns1 + uintptr(jj1)
									zCopy = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1202, libc.VaList(bp+1264, len1-jj1-int32(1), zAns1))
									if libc.BoolInt32(libsqlite3.Xsqlite3_strglob(tls, zCopy, (*(*String)(unsafe.Pointer(bp + 8))).Fz) == 0)^isGlob != 0 {
										errorMessage(tls, __ccgo_ts+1304, libc.VaList(bp+1264, prevLine, zFilename, zCopy, (*(*String)(unsafe.Pointer(bp + 8))).Fz))
									}
									libsqlite3.Xsqlite3_free(tls, zCopy)
									g.FnTest = g.FnTest + 1
									stringReset(tls, bp+8)
								} else {
									/*
									 **  --output
									 **
									 ** Output the result of the previous SQL.
									 */
									if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1347) == 0 {
										logMessage(tls, __ccgo_ts+438, libc.VaList(bp+1264, (*(*String)(unsafe.Pointer(bp + 8))).Fz))
									} else {
										/*
										 **  --source FILENAME
										 **
										 ** Run a subscript from a separate file.
										 */
										if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1354) == 0 {
											zToDel = uintptr(0)
											zNewFile = bp + 1054
											if !(libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(zNewFile))) == libc.Int32FromUint8('/')) {
												k = libc.Int32FromUint64(libc.Xstrlen(tls, zFilename)) - int32(1)
												for {
													if !(k >= 0 && !(libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(zFilename + uintptr(k)))) == libc.Int32FromUint8('/'))) {
														break
													}
													goto _22
												_22:
													;
													k = k - 1
												}
												if k > 0 {
													v23 = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1361, libc.VaList(bp+1264, k, zFilename, zNewFile))
													zToDel = v23
													zNewFile = v23
													if zNewFile == uintptr(0) {
														libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+1369, 0)
														libc.Xabort(tls)
													}
												}
											}
											zNewScript = readFile(tls, zNewFile)
											if g.FiTrace != 0 {
												logMessage(tls, __ccgo_ts+1384, libc.VaList(bp+1264, zNewFile))
											}
											runScript(tls, 0, 0, zNewScript, zNewFile)
											libsqlite3.Xsqlite3_free(tls, zNewScript)
											if g.FiTrace != 0 {
												logMessage(tls, __ccgo_ts+1403, libc.VaList(bp+1264, zNewFile))
											}
											libsqlite3.Xsqlite3_free(tls, zToDel)
										} else {
											/*
											 **  --print MESSAGE....
											 **
											 ** Output the remainder of the line to the log file
											 */
											if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1420) == 0 {
												jj2 = int32(7)
												for {
													if v9 = jj2 < len1; v9 {
														v1 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(zScript + uintptr(ii+jj2))))
														v2 = libc.BoolInt32(v1 == int32(' ') || libc.Uint32FromInt32(v1)-uint32('\t') < uint32(5))
														goto _27
													_27:
													}
													if !(v9 && v2 != 0) {
														break
													}
													goto _24
												_24:
													;
													jj2 = jj2 + 1
												}
												logMessage(tls, __ccgo_ts+1202, libc.VaList(bp+1264, len1-jj2, zScript+uintptr(ii)+uintptr(jj2)))
											} else {
												/*
												 **  --if EXPR
												 **
												 ** Skip forward to the next matching --endif or --else if EXPR is false.
												 */
												if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1426) == 0 {
													jj3 = int32(4)
													for {
														if v9 = jj3 < len1; v9 {
															v1 = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(zScript + uintptr(ii+jj3))))
															v2 = libc.BoolInt32(v1 == int32(' ') || libc.Uint32FromInt32(v1)-uint32('\t') < uint32(5))
															goto _32
														_32:
														}
														if !(v9 && v2 != 0) {
															break
														}
														goto _29
													_29:
														;
														jj3 = jj3 + 1
													}
													pStmt = prepareSql(tls, __ccgo_ts+1429, libc.VaList(bp+1264, len1-jj3, zScript+uintptr(ii)+uintptr(jj3)))
													rc1 = libsqlite3.Xsqlite3_step(tls, pStmt)
													if rc1 != int32(SQLITE_ROW) || libsqlite3.Xsqlite3_column_int(tls, pStmt, 0) == 0 {
														ii = ii + findEndif(tls, zScript+uintptr(ii)+uintptr(len1), int32(1), bp)
													}
													libsqlite3.Xsqlite3_finalize(tls, pStmt)
												} else {
													/*
													 **  --else
													 **
													 ** This command can only be encountered if currently inside an --if that
													 ** is true.  Skip forward to the next matching --endif.
													 */
													if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1441) == 0 {
														ii = ii + findEndif(tls, zScript+uintptr(ii)+uintptr(len1), 0, bp)
													} else {
														/*
														 **  --endif
														 **
														 ** This command can only be encountered if currently inside an --if that
														 ** is true or an --else of a false if.  This is a no-op.
														 */
														if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1446) == 0 {
															/* no-op */
														} else {
															/*
															 **  --start CLIENT
															 **
															 ** Start up the given client.
															 */
															if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1452) == 0 && iClient == 0 {
																iNewClient = libc.Xatoi(tls, bp+1054)
																if iNewClient > 0 {
																	startClient(tls, iNewClient)
																}
															} else {
																/*
																 **  --wait CLIENT TIMEOUT
																 **
																 ** Wait until all tasks complete for the given client.  If CLIENT is
																 ** "all" then wait for all clients to complete.  Wait no longer than
																 ** TIMEOUT milliseconds (default 10,000)
																 */
																if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1458) == 0 && iClient == 0 {
																	if nArg >= int32(2) {
																		v1 = libc.Xatoi(tls, bp+1054+1*100)
																	} else {
																		v1 = int32(10000)
																	}
																	iTimeout = v1
																	libsqlite3.Xsqlite3_snprintf(tls, int32(1000), bp+54, __ccgo_ts+1463, libc.VaList(bp+1264, prevLine, zFilename))
																	waitForClient(tls, libc.Xatoi(tls, bp+1054), iTimeout, bp+54)
																} else {
																	/*
																	 **  --task CLIENT
																	 **     <task-content-here>
																	 **  --end
																	 **
																	 ** Assign work to a client.  Start the client if it is not running
																	 ** already.
																	 */
																	if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1478) == 0 && iClient == 0 {
																		iTarget = libc.Xatoi(tls, bp+1054)
																		iEnd = findEnd(tls, zScript+uintptr(ii)+uintptr(len1), bp)
																		if iTarget < 0 {
																			errorMessage(tls, __ccgo_ts+1483, libc.VaList(bp+1264, prevLine, zFilename, iTarget))
																		} else {
																			zTask = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1202, libc.VaList(bp+1264, iEnd, zScript+uintptr(ii)+uintptr(len1)))
																			if nArg > int32(1) {
																				zTName = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+438, libc.VaList(bp+1264, bp+1054+1*100))
																			} else {
																				zTName = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1520, libc.VaList(bp+1264, filenameTail(tls, zFilename), prevLine))
																			}
																			startClient(tls, iTarget)
																			runSql(tls, __ccgo_ts+1526, libc.VaList(bp+1264, iTarget, zTask, zTName))
																			libsqlite3.Xsqlite3_free(tls, zTask)
																			libsqlite3.Xsqlite3_free(tls, zTName)
																		}
																		iEnd = iEnd + tokenLength(tls, zScript+uintptr(ii)+uintptr(len1)+uintptr(iEnd), bp)
																		len1 = len1 + iEnd
																		iBegin = ii + len1
																	} else {
																		/*
																		 **  --breakpoint
																		 **
																		 ** This command calls "test_breakpoint()" which is a routine provided
																		 ** as a convenient place to set a debugger breakpoint.
																		 */
																		if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1582) == 0 {
																			test_breakpoint(tls)
																		} else {
																			/*
																			 **  --show-sql-errors BOOLEAN
																			 **
																			 ** Turn display of SQL errors on and off.
																			 */
																			if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1593) == 0 {
																				if nArg >= int32(1) {
																					v1 = libc.BoolInt32(!(booleanValue(tls, bp+1054) != 0))
																				} else {
																					v1 = int32(1)
																				}
																				g.FbIgnoreSqlErrors = v1
																			} else {
																				/* error */
																				errorMessage(tls, __ccgo_ts+1609, libc.VaList(bp+1264, prevLine, zFilename, bp+24))
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		ii = ii + len1
	}
	if iBegin < ii {
		zSql1 = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1202, libc.VaList(bp+1264, ii-iBegin, zScript+uintptr(iBegin)))
		runSql(tls, zSql1, 0)
		libsqlite3.Xsqlite3_free(tls, zSql1)
	}
	stringFree(tls, bp+8)
}

// C documentation
//
//	/*
//	** Look for a command-line option.  If present, return a pointer.
//	** Return NULL if missing.
//	**
//	** hasArg==0 means the option is a flag.  It is either present or not.
//	** hasArg==1 means the option has an argument.  Return a pointer to the
//	** argument.
//	*/
func findOption(tls *libc.TLS, azArg uintptr, pnArg uintptr, zOption uintptr, hasArg int32) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i, j, nArg, v2, v3 int32
	var z, zReturn uintptr
	_, _, _, _, _, _, _ = i, j, nArg, z, zReturn, v2, v3
	zReturn = uintptr(0)
	nArg = *(*int32)(unsafe.Pointer(pnArg))
	i = 0
	for {
		if !(i < nArg) {
			break
		}
		if i+hasArg >= nArg {
			break
		}
		z = *(*uintptr)(unsafe.Pointer(azArg + uintptr(i)*8))
		if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z))) != int32('-') {
			goto _1
		}
		z = z + 1
		if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z))) == int32('-') {
			if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(z + 1))) == 0 {
				break
			}
			z = z + 1
		}
		if libc.Xstrcmp(tls, z, zOption) == 0 {
			if hasArg != 0 && i == nArg-int32(1) {
				fatalError(tls, __ccgo_ts+1645, libc.VaList(bp+8, z))
			}
			if hasArg != 0 {
				zReturn = *(*uintptr)(unsafe.Pointer(azArg + uintptr(i+int32(1))*8))
			} else {
				zReturn = *(*uintptr)(unsafe.Pointer(azArg + uintptr(i)*8))
			}
			j = i + int32(1) + libc.BoolInt32(hasArg != 0)
			for j < nArg {
				v2 = i
				i = i + 1
				v3 = j
				j = j + 1
				*(*uintptr)(unsafe.Pointer(azArg + uintptr(v2)*8)) = *(*uintptr)(unsafe.Pointer(azArg + uintptr(v3)*8))
			}
			*(*int32)(unsafe.Pointer(pnArg)) = i
			return zReturn
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return zReturn
}

// C documentation
//
//	/* Print a usage message for the program and exit */
func usage(tls *libc.TLS, argv0 uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i int32
	var zTail uintptr
	_, _ = i, zTail
	zTail = argv0
	i = 0
	for {
		if !(*(*uint8)(unsafe.Pointer(argv0 + uintptr(i))) != 0) {
			break
		}
		if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(argv0 + uintptr(i)))) == int32('/') {
			zTail = argv0 + uintptr(i) + uintptr(1)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+1693, libc.VaList(bp+8, zTail))
	libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+1732, 0)
	libc.Xexit(tls, int32(1))
}

// C documentation
//
//	/* Report on unrecognized arguments */
func unrecognizedArguments(tls *libc.TLS, argv0 uintptr, nArg int32, azArg uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i int32
	_ = i
	libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+2304, libc.VaList(bp+8, argv0))
	i = 0
	for {
		if !(i < nArg) {
			break
		}
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+2332, libc.VaList(bp+8, *(*uintptr)(unsafe.Pointer(azArg + uintptr(i)*8))))
		goto _1
	_1:
		;
		i = i + 1
	}
	libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+2336, 0)
	libc.Xexit(tls, int32(1))
}

func main1(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var i, iClient, iRep, iTimeout, iTmout, nRep, nTry, openFlags, rc, v4 int32
	var pStmt, zCOption, zClient, zJMode, zNRep, zTmout, zTrace, v3 uintptr
	var v5 bool
	var _ /* n at bp+0 */ int32
	var _ /* taskId at bp+16 */ int32
	var _ /* zScript at bp+8 */ uintptr
	var _ /* zTaskName at bp+24 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = i, iClient, iRep, iTimeout, iTmout, nRep, nTry, openFlags, pStmt, rc, zCOption, zClient, zJMode, zNRep, zTmout, zTrace, v3, v4, v5
	openFlags = int32(SQLITE_OPEN_READWRITE)
	nRep = int32(1)
	iTmout = 0
	g.Fargv0 = *(*uintptr)(unsafe.Pointer(argv))
	g.FiTrace = int32(1)
	if argc < int32(2) {
		usage(tls, *(*uintptr)(unsafe.Pointer(argv)))
	}
	g.FzDbFile = *(*uintptr)(unsafe.Pointer(argv + 1*8))
	if strglob(tls, __ccgo_ts+2338, g.FzDbFile) != 0 {
		usage(tls, *(*uintptr)(unsafe.Pointer(argv)))
	}
	if libc.Bool(0 != 0) && libc.Xstrcmp(tls, libsqlite3.Xsqlite3_sourceid(tls), __ccgo_ts+2345) != 0 {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+2430, libc.VaList(bp+40, libsqlite3.Xsqlite3_sourceid(tls), __ccgo_ts+2345))
		libc.Xexit(tls, int32(1))
	}
	*(*int32)(unsafe.Pointer(bp)) = argc - int32(2)
	libsqlite3.Xsqlite3_snprintf(tls, int32(32), uintptr(unsafe.Pointer(&g))+64, __ccgo_ts+2490, libc.VaList(bp+40, libc.Xgetpid(tls)))
	zJMode = findOption(tls, argv+uintptr(2)*8, bp, __ccgo_ts+2502, int32(1))
	zNRep = findOption(tls, argv+uintptr(2)*8, bp, __ccgo_ts+2514, int32(1))
	if zNRep != 0 {
		nRep = libc.Xatoi(tls, zNRep)
	}
	if nRep < int32(1) {
		nRep = int32(1)
	}
	g.FzVfs = findOption(tls, argv+uintptr(2)*8, bp, __ccgo_ts+2521, int32(1))
	zClient = findOption(tls, argv+uintptr(2)*8, bp, __ccgo_ts+2525, int32(1))
	g.FzErrLog = findOption(tls, argv+uintptr(2)*8, bp, __ccgo_ts+2532, int32(1))
	g.FzLog = findOption(tls, argv+uintptr(2)*8, bp, __ccgo_ts+2539, int32(1))
	zTrace = findOption(tls, argv+uintptr(2)*8, bp, __ccgo_ts+2543, int32(1))
	if zTrace != 0 {
		g.FiTrace = libc.Xatoi(tls, zTrace)
	}
	if findOption(tls, argv+uintptr(2)*8, bp, __ccgo_ts+2549, 0) != uintptr(0) {
		g.FiTrace = 0
	}
	zTmout = findOption(tls, argv+uintptr(2)*8, bp, __ccgo_ts+2555, int32(1))
	if zTmout != 0 {
		iTmout = libc.Xatoi(tls, zTmout)
	}
	g.FbSqlTrace = libc.BoolInt32(findOption(tls, argv+uintptr(2)*8, bp, __ccgo_ts+2563, 0) != uintptr(0))
	g.FbSync = libc.BoolInt32(findOption(tls, argv+uintptr(2)*8, bp, __ccgo_ts+2572, 0) != uintptr(0))
	if g.FzErrLog != 0 {
		g.FpErrLog = libc.Xfopen(tls, g.FzErrLog, __ccgo_ts+2577)
	} else {
		g.FpErrLog = libc.Xstderr
	}
	if g.FzLog != 0 {
		g.FpLog = libc.Xfopen(tls, g.FzLog, __ccgo_ts+2577)
	} else {
		g.FpLog = libc.Xstdout
	}
	libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_LOG), libc.VaList(bp+40, __ccgo_fp(sqlErrorCallback), 0))
	if zClient != 0 {
		iClient = libc.Xatoi(tls, zClient)
		if iClient < int32(1) {
			fatalError(tls, __ccgo_ts+2579, libc.VaList(bp+40, iClient))
		}
		libsqlite3.Xsqlite3_snprintf(tls, int32(32), uintptr(unsafe.Pointer(&g))+64, __ccgo_ts+2606, libc.VaList(bp+40, libc.Xgetpid(tls), iClient))
	} else {
		nTry = 0
		if g.FiTrace > 0 {
			libc.Xprintf(tls, __ccgo_ts+2622, libc.VaList(bp+40, *(*uintptr)(unsafe.Pointer(argv))))
			i = int32(1)
			for {
				if !(i < argc) {
					break
				}
				libc.Xprintf(tls, __ccgo_ts+2332, libc.VaList(bp+40, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
				goto _1
			_1:
				;
				i = i + 1
			}
			libc.Xprintf(tls, __ccgo_ts+2336, 0)
			libc.Xprintf(tls, __ccgo_ts+2632, 0)
			i = 0
			for {
				v3 = libsqlite3.Xsqlite3_compileoption_get(tls, i)
				zCOption = v3
				if !(v3 != uintptr(0)) {
					break
				}
				libc.Xprintf(tls, __ccgo_ts+2737, libc.VaList(bp+40, zCOption))
				goto _2
			_2:
				;
				i = i + 1
			}
			libc.Xfflush(tls, libc.Xstdout)
		}
		iClient = 0
		for {
			if nTry%int32(5) == int32(4) {
				if nTry > int32(5) {
					v3 = __ccgo_ts + 2750
				} else {
					v3 = __ccgo_ts + 1099
				}
				libc.Xprintf(tls, __ccgo_ts+2757, libc.VaList(bp+40, v3, g.FzDbFile))
			}
			rc = libc.Xunlink(tls, g.FzDbFile)
			if rc != 0 && *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(ENOENT) {
				rc = 0
			}
			goto _6
		_6:
			;
			if v5 = rc != 0; v5 {
				nTry = nTry + 1
				v4 = nTry
			}
			if !(v5 && v4 < int32(60) && libsqlite3.Xsqlite3_sleep(tls, int32(1000)) > 0) {
				break
			}
		}
		if rc != 0 {
			fatalError(tls, __ccgo_ts+2786, libc.VaList(bp+40, g.FzDbFile, nTry))
		}
		openFlags = openFlags | int32(SQLITE_OPEN_CREATE)
	}
	rc = libsqlite3.Xsqlite3_open_v2(tls, g.FzDbFile, uintptr(unsafe.Pointer(&g))+24, openFlags, g.FzVfs)
	if rc != 0 {
		fatalError(tls, __ccgo_ts+2827, libc.VaList(bp+40, g.FzDbFile))
	}
	if iTmout > 0 {
		libsqlite3.Xsqlite3_busy_timeout(tls, g.Fdb, iTmout)
	}
	if zJMode != 0 {
		runSql(tls, __ccgo_ts+2844, libc.VaList(bp+40, zJMode))
	}
	if !(g.FbSync != 0) {
		trySql(tls, __ccgo_ts+2868, 0)
	}
	libsqlite3.Xsqlite3_enable_load_extension(tls, g.Fdb, int32(1))
	libsqlite3.Xsqlite3_busy_handler(tls, g.Fdb, __ccgo_fp(busyHandler), uintptr(0))
	libsqlite3.Xsqlite3_create_function(tls, g.Fdb, __ccgo_ts+2891, 0, int32(SQLITE_UTF8), uintptr(0), __ccgo_fp(vfsNameFunc), uintptr(0), uintptr(0))
	libsqlite3.Xsqlite3_create_function(tls, g.Fdb, __ccgo_ts+2899, int32(1), int32(SQLITE_UTF8), uintptr(0), __ccgo_fp(evalFunc), uintptr(0), uintptr(0))
	g.FiTimeout = int32(DEFAULT_TIMEOUT)
	if g.FbSqlTrace != 0 {
		libsqlite3.Xsqlite3_trace(tls, g.Fdb, __ccgo_fp(sqlTraceCallback), uintptr(0))
	}
	if iClient > 0 {
		if *(*int32)(unsafe.Pointer(bp)) > 0 {
			unrecognizedArguments(tls, *(*uintptr)(unsafe.Pointer(argv)), *(*int32)(unsafe.Pointer(bp)), argv+uintptr(2)*8)
		}
		if g.FiTrace != 0 {
			logMessage(tls, __ccgo_ts+2904, 0)
		}
		for int32(1) != 0 {
			*(*uintptr)(unsafe.Pointer(bp + 24)) = uintptr(0)
			rc = startScript(tls, iClient, bp+8, bp+16, bp+24)
			if rc == int32(SQLITE_DONE) {
				break
			}
			if g.FiTrace != 0 {
				logMessage(tls, __ccgo_ts+2917, libc.VaList(bp+40, *(*uintptr)(unsafe.Pointer(bp + 24)), *(*int32)(unsafe.Pointer(bp + 16))))
			}
			runScript(tls, iClient, *(*int32)(unsafe.Pointer(bp + 16)), *(*uintptr)(unsafe.Pointer(bp + 8)), *(*uintptr)(unsafe.Pointer(bp + 24)))
			if g.FiTrace != 0 {
				logMessage(tls, __ccgo_ts+2931, libc.VaList(bp+40, *(*uintptr)(unsafe.Pointer(bp + 24)), *(*int32)(unsafe.Pointer(bp + 16))))
			}
			finishScript(tls, iClient, *(*int32)(unsafe.Pointer(bp + 16)), 0)
			libsqlite3.Xsqlite3_free(tls, *(*uintptr)(unsafe.Pointer(bp + 24)))
			libsqlite3.Xsqlite3_sleep(tls, int32(10))
		}
		if g.FiTrace != 0 {
			logMessage(tls, __ccgo_ts+2943, 0)
		}
	} else {
		if *(*int32)(unsafe.Pointer(bp)) == 0 {
			fatalError(tls, __ccgo_ts+2954, 0)
		}
		if *(*int32)(unsafe.Pointer(bp)) > int32(1) {
			unrecognizedArguments(tls, *(*uintptr)(unsafe.Pointer(argv)), *(*int32)(unsafe.Pointer(bp)), argv+uintptr(2)*8)
		}
		runSql(tls, __ccgo_ts+2978, 0)
		*(*uintptr)(unsafe.Pointer(bp + 8)) = readFile(tls, *(*uintptr)(unsafe.Pointer(argv + 2*8)))
		iRep = int32(1)
		for {
			if !(iRep <= nRep) {
				break
			}
			if g.FiTrace != 0 {
				logMessage(tls, __ccgo_ts+3414, libc.VaList(bp+40, *(*uintptr)(unsafe.Pointer(argv + 2*8)), iRep))
			}
			runScript(tls, 0, 0, *(*uintptr)(unsafe.Pointer(bp + 8)), *(*uintptr)(unsafe.Pointer(argv + 2*8)))
			if g.FiTrace != 0 {
				logMessage(tls, __ccgo_ts+3442, libc.VaList(bp+40, *(*uintptr)(unsafe.Pointer(argv + 2*8)), iRep))
			}
			goto _8
		_8:
			;
			iRep = iRep + 1
		}
		libsqlite3.Xsqlite3_free(tls, *(*uintptr)(unsafe.Pointer(bp + 8)))
		waitForClient(tls, 0, int32(2000), __ccgo_ts+3468)
		trySql(tls, __ccgo_ts+3488, 0)
		libsqlite3.Xsqlite3_sleep(tls, int32(10))
		g.FiTimeout = 0
		iTimeout = int32(1000)
		for {
			v4 = trySql(tls, __ccgo_ts+3517, 0)
			rc = v4
			if !((v4 == int32(SQLITE_BUSY) || rc == int32(SQLITE_ROW)) && iTimeout > 0) {
				break
			}
			libsqlite3.Xsqlite3_sleep(tls, int32(10))
			iTimeout = iTimeout - int32(10)
		}
		libsqlite3.Xsqlite3_sleep(tls, int32(100))
		pStmt = prepareSql(tls, __ccgo_ts+3538, 0)
		iTimeout = int32(1000)
		for {
			v4 = libsqlite3.Xsqlite3_step(tls, pStmt)
			rc = v4
			if !(v4 == int32(SQLITE_BUSY) && iTimeout > 0) {
				break
			}
			libsqlite3.Xsqlite3_sleep(tls, int32(10))
			iTimeout = iTimeout - int32(10)
		}
		if rc == int32(SQLITE_ROW) {
			g.FnError += libsqlite3.Xsqlite3_column_int(tls, pStmt, 0)
			g.FnTest += libsqlite3.Xsqlite3_column_int(tls, pStmt, int32(1))
		}
		libsqlite3.Xsqlite3_finalize(tls, pStmt)
	}
	libsqlite3.Xsqlite3_close(tls, g.Fdb)
	maybeClose(tls, g.FpLog)
	maybeClose(tls, g.FpErrLog)
	if iClient == 0 {
		libc.Xprintf(tls, __ccgo_ts+3573, libc.VaList(bp+40, g.FnError, g.FnTest))
		libc.Xprintf(tls, __ccgo_ts+3609, libc.VaList(bp+40, *(*uintptr)(unsafe.Pointer(argv))))
		i = int32(1)
		for {
			if !(i < argc) {
				break
			}
			libc.Xprintf(tls, __ccgo_ts+2332, libc.VaList(bp+40, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
			goto _11
		_11:
			;
			i = i + 1
		}
		libc.Xprintf(tls, __ccgo_ts+2336, 0)
	}
	return libc.BoolInt32(g.FnError > 0)
}

func main() {
	libc.Start(main1)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "%s%.*s\n\x00%s:ERROR: \x00%s:FATAL: \x00UPDATE client SET wantHalt=1;\x00%s: \x00main\x00timeout after %dms\x00[%.*s]\x00(info) %s\x00(errcode=%d) %s\x00%s\n%s\n\x00out of memory\x00 \x00nil\x00'\x00error(%d)\x00BEGIN IMMEDIATE\x00in startScript: %s\x00UPDATE counters SET nError=nError+%d, nTest=nTest+%d\x00SELECT 1 FROM client WHERE id=%d AND wantHalt\x00DELETE FROM client WHERE id=%d\x00COMMIT TRANSACTION;\x00SELECT script, id, name FROM task WHERE client=%d AND starttime IS NULL ORDER BY id LIMIT 1\x00%s\x00UPDATE task   SET starttime=strftime('%%Y-%%m-%%d %%H:%%M:%%f','now') WHERE id=%d;\x00Waited over 30 seconds with no work.  Giving up.\x00DELETE FROM client WHERE id=%d; COMMIT;\x00COMMIT\x00UPDATE task   SET endtime=strftime('%%Y-%%m-%%d %%H:%%M:%%f','now') WHERE id=%d;\x00INSERT OR IGNORE INTO client VALUES(%d,0)\x00%s \"%s\" --client %d --trace %d\x00%z --sqltrace\x00%z --sync\x00%z --vfs \"%s\"\x00system('%q')\x00%z &\x00system() fails with error code %d\x00rb\x00cannot open \"%s\" for reading\x00--end\x00--endif\x00--else\x00--if\x00SELECT 1 FROM task WHERE client=%d   AND client IN (SELECT id FROM client)  AND endtime IS NULL\x00SELECT 1 FROM task WHERE client IN (SELECT id FROM client)   AND endtime IS NULL\x00\x00%stimeout waiting for client %d\x00%stimeout waiting for all clients\x00on\x00yes\x00off\x00no\x00unknown boolean: [%s]\x00%.*s\x00sleep\x00exit\x00testcase\x00finish\x00reset\x00match\x00line %d of %s:\nExpected [%.*s]\n     Got [%s]\x00glob\x00notglob\x00line %d of %s:\nExpected [%s]\n     Got [%s]\x00output\x00source\x00%.*s/%s\x00out of memory\n\x00begin script [%s]\n\x00end script [%s]\n\x00print\x00if\x00SELECT %.*s\x00else\x00endif\x00start\x00wait\x00line %d of %s\n\x00task\x00line %d of %s: bad client number: %d\x00%s:%d\x00INSERT INTO task(client,script,name) VALUES(%d,'%q',%Q)\x00breakpoint\x00show-sql-errors\x00line %d of %s: unknown command --%s\x00command-line option \"--%s\" requires an argument\x00Usage: %s DATABASE ?OPTIONS? ?SCRIPT?\n\x00Options:\n   --errlog FILENAME           Write errors to FILENAME\n   --journalmode MODE          Use MODE as the journal_mode\n   --log FILENAME              Log messages to FILENAME\n   --quiet                     Suppress unnecessary output\n   --vfs NAME                  Use NAME as the VFS\n   --repeat N                  Repeat the test N times\n   --sqltrace                  Enable SQL tracing\n   --sync                      Enable synchronous disk writes\n   --timeout MILLISEC          Busy timeout is MILLISEC\n   --trace BOOLEAN             Enable or disable tracing\n\x00%s: unrecognized arguments:\x00 %s\x00\n\x00*.test\x002026-05-05 10:34:17 c88b22011a54b4f6fbd149e9f8e4de77658ce58143a1af0e3785e4e64751alt1\x00SQLite library and header mismatch\nLibrary: %s\nHeader:  %s\n\x00%05d.mptest\x00journalmode\x00repeat\x00vfs\x00client\x00errlog\x00log\x00trace\x00quiet\x00timeout\x00sqltrace\x00sync\x00a\x00illegal client number: %d\n\x00%05d.client%02d\x00BEGIN: %s\x00With SQLite 3.53.1 2026-05-05 10:34:17 c88b22011a54b4f6fbd149e9f8e4de77658ce58143a1af0e3785e4e64751alt1\n\x00-DSQLITE_%s\n\x00still \x00... %strying to unlink '%s'\n\x00unable to unlink '%s' after %d attempts\n\x00cannot open [%s]\x00PRAGMA journal_mode=%Q;\x00PRAGMA synchronous=OFF\x00vfsname\x00eval\x00start-client\x00begin %s (%d)\x00end %s (%d)\x00end-client\x00missing script filename\x00DROP TABLE IF EXISTS task;\nDROP TABLE IF EXISTS counters;\nDROP TABLE IF EXISTS client;\nCREATE TABLE task(\n  id INTEGER PRIMARY KEY,\n  name TEXT,\n  client INTEGER,\n  starttime DATE,\n  endtime DATE,\n  script TEXT\n);CREATE INDEX task_i1 ON task(client, starttime);\nCREATE INDEX task_i2 ON task(client, endtime);\nCREATE TABLE counters(nError,nTest);\nINSERT INTO counters VALUES(0,0);\nCREATE TABLE client(id INTEGER PRIMARY KEY, wantHalt);\n\x00begin script [%s] cycle %d\n\x00end script [%s] cycle %d\n\x00during shutdown...\n\x00UPDATE client SET wantHalt=1\x00SELECT 1 FROM client\x00SELECT nError, nTest FROM counters\x00Summary: %d errors out of %d tests\n\x00END: %s\x00"
