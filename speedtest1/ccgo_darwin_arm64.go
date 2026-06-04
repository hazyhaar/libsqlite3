// Code generated for darwin/arm64 by 'generator -DNDEBUG -DSQLITE_DISABLE_INTRINSIC -DSQLITE_ENABLE_DBPAGE_VTAB -DSQLITE_ENABLE_RTREE -ignore-unsupported-alignment -ignore-link-errors -o speedtest1/ccgo_darwin_arm64.go -I /var/folders/4f/mc8mts295pqf7gmnfcwh6g8w0000gn/T/libsqlite3/sqlite-src-3530200 /var/folders/4f/mc8mts295pqf7gmnfcwh6g8w0000gn/T/libsqlite3/sqlite-src-3530200/test/speedtest1.c -lsqlite3', DO NOT EDIT.

//go:build darwin && arm64

package main

import (
	"math"
	"reflect"
	"unsafe"

	"modernc.org/libc"
	"modernc.org/libsqlite3"
)

var _ = math.Pi
var _ reflect.Type
var _ unsafe.Pointer

const NAMEWIDTH = 60
const SQLITE_BLOB = 4
const SQLITE_CONFIG_HEAP = 8
const SQLITE_CONFIG_LOOKASIDE = 13
const SQLITE_CONFIG_MEMSTATUS = 9
const SQLITE_CONFIG_MULTITHREAD = 2
const SQLITE_CONFIG_PAGECACHE = 7
const SQLITE_CONFIG_SERIALIZED = 3
const SQLITE_CONFIG_SINGLETHREAD = 1
const SQLITE_DBCONFIG_LOOKASIDE = 1001
const SQLITE_DBCONFIG_STMT_SCANSTATUS = 1018
const SQLITE_DBSTATUS_CACHE_HIT = 7
const SQLITE_DBSTATUS_CACHE_MISS = 8
const SQLITE_DBSTATUS_CACHE_USED = 1
const SQLITE_DBSTATUS_CACHE_WRITE = 9
const SQLITE_DBSTATUS_LOOKASIDE_HIT = 4
const SQLITE_DBSTATUS_LOOKASIDE_MISS_FULL = 6
const SQLITE_DBSTATUS_LOOKASIDE_MISS_SIZE = 5
const SQLITE_DBSTATUS_LOOKASIDE_USED = 0
const SQLITE_DBSTATUS_SCHEMA_USED = 2
const SQLITE_DBSTATUS_STMT_USED = 3
const SQLITE_FCNTL_RESERVE_BYTES = 38
const SQLITE_FLOAT = 2
const SQLITE_OK = 0
const SQLITE_OPEN_CREATE = 4
const SQLITE_OPEN_NOMUTEX = 32768
const SQLITE_OPEN_READWRITE = 2
const SQLITE_ROW = 100
const SQLITE_STATUS_MALLOC_COUNT = 9
const SQLITE_STATUS_MALLOC_SIZE = 5
const SQLITE_STATUS_MEMORY_USED = 0
const SQLITE_STATUS_PAGECACHE_OVERFLOW = 2
const SQLITE_STATUS_PAGECACHE_SIZE = 7
const SQLITE_TESTCTRL_PRNG_SEED = 28
const SQLITE_UTF8 = 1

type __predefined_size_t = uint64

// C documentation
//
//	/*
//	** A program for performance testing.
//	**
//	** To build this program against an historical version of SQLite for comparison
//	** testing:
//	**
//	**    Unix:
//	**
//	**        ./configure --all
//	**        make clean speedtest1
//	**        mv speedtest1 speedtest1-current
//	**        cp $HISTORICAL_SQLITE3_C_H .
//	**        touch sqlite3.c sqlite3.h .target_source
//	**        make speedtest1
//	**        mv speedtest1 speedtest1-baseline
//	**
//	**    Windows:
//	**
//	**        nmake /f Makefile.msc clean speedtest1.exe
//	**        mv speedtest1.exe speedtest1-current.exe
//	**        cp $HISTORICAL_SQLITE_C_H .
//	**        touch sqlite3.c sqlite3.h .target_source
//	**        nmake /f Makefile.msc speedtest1.exe
//	**        mv speedtest1.exe speedtest1-baseline.exe
//	**
//	** The available command-line options are described below:
//	*/
var zHelp = [2897]int8{'U', 's', 'a', 'g', 'e', ':', ' ', '%', 's', ' ', '[', '-', '-', 'o', 'p', 't', 'i', 'o', 'n', 's', ']', ' ', 'D', 'A', 'T', 'A', 'B', 'A', 'S', 'E', 10, 'O', 'p', 't', 'i', 'o', 'n', 's', ':', 10, ' ', ' ', '-', '-', 'a', 'u', 't', 'o', 'v', 'a', 'c', 'u', 'u', 'm', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'E', 'n', 'a', 'b', 'l', 'e', ' ', 'A', 'U', 'T', 'O', 'V', 'A', 'C', 'U', 'U', 'M', ' ', 'm', 'o', 'd', 'e', 10, ' ', ' ', '-', '-', 'b', 'i', 'g', '-', 't', 'r', 'a', 'n', 's', 'a', 'c', 't', 'i', 'o', 'n', 's', ' ', ' ', 'A', 'd', 'd', ' ', 'B', 'E', 'G', 'I', 'N', '/', 'E', 'N', 'D', ' ', 'a', 'r', 'o', 'u', 'n', 'd', ' ', 'a', 'l', 'l', ' ', 'l', 'a', 'r', 'g', 'e', ' ', 't', 'e', 's', 't', 's', 10, ' ', ' ', '-', '-', 'c', 'a', 'c', 'h', 'e', 's', 'i', 'z', 'e', ' ', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 'P', 'R', 'A', 'G', 'M', 'A', ' ', 'c', 'a', 'c', 'h', 'e', '_', 's', 'i', 'z', 'e', '=', 'N', '.', ' ', 'N', 'o', 't', 'e', ':', ' ', 'N', ' ', 'i', 's', ' ', 'p', 'a', 'g', 'e', 's', ',', ' ', 'n', 'o', 't', ' ', 'b', 'y', 't', 'e', 's', 10, ' ', ' ', '-', '-', 'c', 'h', 'e', 'c', 'k', 'p', 'o', 'i', 'n', 't', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'R', 'u', 'n', ' ', 'P', 'R', 'A', 'G', 'M', 'A', ' ', 'w', 'a', 'l', '_', 'c', 'h', 'e', 'c', 'k', 'p', 'o', 'i', 'n', 't', ' ', 'a', 'f', 't', 'e', 'r', ' ', 'e', 'a', 'c', 'h', ' ', 't', 'e', 's', 't', ' ', 'c', 'a', 's', 'e', 10, ' ', ' ', '-', '-', 'e', 'x', 'c', 'l', 'u', 's', 'i', 'v', 'e', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'E', 'n', 'a', 'b', 'l', 'e', ' ', 'l', 'o', 'c', 'k', 'i', 'n', 'g', '_', 'm', 'o', 'd', 'e', '=', 'E', 'X', 'C', 'L', 'U', 'S', 'I', 'V', 'E', 10, ' ', ' ', '-', '-', 'e', 'x', 'p', 'l', 'a', 'i', 'n', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'L', 'i', 'k', 'e', ' ', '-', '-', 's', 'q', 'l', 'o', 'n', 'l', 'y', ' ', 'b', 'u', 't', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'd', 'd', 'e', 'd', ' ', 'E', 'X', 'P', 'L', 'A', 'I', 'N', ' ', 'k', 'e', 'y', 'w', 'o', 'r', 'd', 's', 10, ' ', ' ', '-', '-', 'f', 'u', 'l', 'l', 'f', 's', 'y', 'n', 'c', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'E', 'n', 'a', 'b', 'l', 'e', ' ', 'f', 'u', 'l', 'l', 'f', 's', 'y', 'n', 'c', '=', 'T', 'R', 'U', 'E', 10, ' ', ' ', '-', '-', 'h', 'a', 'r', 'd', '-', 'h', 'e', 'a', 'p', '-', 'l', 'i', 'm', 'i', 't', ' ', 'N', ' ', 'T', 'h', 'e', ' ', 'h', 'a', 'r', 'd', ' ', 'l', 'i', 'm', 'i', 't', ' ', 'o', 'n', ' ', 't', 'h', 'e', ' ', 'm', 'a', 'x', 'i', 'm', 'u', 'm', ' ', 'h', 'e', 'a', 'p', ' ', 's', 'i', 'z', 'e', 10, ' ', ' ', '-', '-', 'h', 'e', 'a', 'p', ' ', 'S', 'Z', ' ', 'M', 'I', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'M', 'e', 'm', 'o', 'r', 'y', ' ', 'a', 'l', 'l', 'o', 'c', 'a', 't', 'o', 'r', ' ', 'u', 's', 'e', 's', ' ', 'S', 'Z', ' ', 'b', 'y', 't', 'e', 's', ' ', '&', ' ', 'm', 'i', 'n', ' ', 'a', 'l', 'l', 'o', 'c', 'a', 't', 'i', 'o', 'n', ' ', 'M', 'I', 'N', 10, ' ', ' ', '-', '-', 'i', 'n', 'c', 'r', 'v', 'a', 'c', 'u', 'u', 'm', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'E', 'n', 'a', 'b', 'l', 'e', ' ', 'i', 'n', 'c', 'r', 'e', 'm', 'e', 'n', 'a', 't', 'a', 'l', ' ', 'v', 'a', 'c', 'u', 'u', 'm', ' ', 'm', 'o', 'd', 'e', 10, ' ', ' ', '-', '-', 'j', 'o', 'u', 'r', 'n', 'a', 'l', ' ', 'M', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 't', 'h', 'e', ' ', 'j', 'o', 'u', 'r', 'n', 'a', 'l', '_', 'm', 'o', 'd', 'e', ' ', 't', 'o', ' ', 'M', 10, ' ', ' ', '-', '-', 'k', 'e', 'y', ' ', 'K', 'E', 'Y', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 't', 'h', 'e', ' ', 'e', 'n', 'c', 'r', 'y', 'p', 't', 'i', 'o', 'n', ' ', 'k', 'e', 'y', ' ', 't', 'o', ' ', 'K', 'E', 'Y', 10, ' ', ' ', '-', '-', 'l', 'o', 'o', 'k', 'a', 's', 'i', 'd', 'e', ' ', 'N', ' ', 'S', 'Z', ' ', ' ', ' ', ' ', 'C', 'o', 'n', 'f', 'i', 'g', 'u', 'r', 'e', ' ', 'l', 'o', 'o', 'k', 'a', 's', 'i', 'd', 'e', ' ', 'f', 'o', 'r', ' ', 'N', ' ', 's', 'l', 'o', 't', 's', ' ', 'o', 'f', ' ', 'S', 'Z', ' ', 'b', 'y', 't', 'e', 's', ' ', 'e', 'a', 'c', 'h', 10, ' ', ' ', '-', '-', 'm', 'e', 'm', 'd', 'b', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'U', 's', 'e', ' ', 'a', 'n', ' ', 'i', 'n', '-', 'm', 'e', 'm', 'o', 'r', 'y', ' ', 'd', 'a', 't', 'a', 'b', 'a', 's', 'e', 10, ' ', ' ', '-', '-', 'm', 'm', 'a', 'p', ' ', 'S', 'Z', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'M', 'M', 'A', 'P', ' ', 't', 'h', 'e', ' ', 'f', 'i', 'r', 's', 't', ' ', 'S', 'Z', ' ', 'b', 'y', 't', 'e', 's', ' ', 'o', 'f', ' ', 't', 'h', 'e', ' ', 'd', 'a', 't', 'a', 'b', 'a', 's', 'e', ' ', 'f', 'i', 'l', 'e', 10, ' ', ' ', '-', '-', 'm', 'u', 'l', 't', 'i', 't', 'h', 'r', 'e', 'a', 'd', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 'm', 'u', 'l', 't', 'i', 't', 'h', 'r', 'e', 'a', 'd', 'e', 'd', ' ', 'm', 'o', 'd', 'e', 10, ' ', ' ', '-', '-', 'n', 'o', 'm', 'e', 'm', 's', 't', 'a', 't', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'D', 'i', 's', 'a', 'b', 'l', 'e', ' ', 'm', 'e', 'm', 'o', 'r', 'y', ' ', 's', 't', 'a', 't', 'i', 's', 't', 'i', 'c', 's', 10, ' ', ' ', '-', '-', 'n', 'o', 'm', 'u', 't', 'e', 'x', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'O', 'p', 'e', 'n', ' ', 'd', 'b', ' ', 'w', 'i', 't', 'h', ' ', 'S', 'Q', 'L', 'I', 'T', 'E', '_', 'O', 'P', 'E', 'N', '_', 'N', 'O', 'M', 'U', 'T', 'E', 'X', 10, ' ', ' ', '-', '-', 'n', 'o', 's', 'y', 'n', 'c', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 'P', 'R', 'A', 'G', 'M', 'A', ' ', 's', 'y', 'n', 'c', 'h', 'r', 'o', 'n', 'o', 'u', 's', '=', 'O', 'F', 'F', 10, ' ', ' ', '-', '-', 'n', 'o', 't', 'n', 'u', 'l', 'l', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'A', 'd', 'd', ' ', 'N', 'O', 'T', ' ', 'N', 'U', 'L', 'L', ' ', 'c', 'o', 'n', 's', 't', 'r', 'a', 'i', 'n', 't', 's', ' ', 't', 'o', ' ', 't', 'a', 'b', 'l', 'e', ' ', 'c', 'o', 'l', 'u', 'm', 'n', 's', 10, ' ', ' ', '-', '-', 'o', 'u', 't', 'p', 'u', 't', ' ', 'F', 'I', 'L', 'E', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 't', 'o', 'r', 'e', ' ', 'S', 'Q', 'L', ' ', 'o', 'u', 't', 'p', 'u', 't', ' ', 'i', 'n', ' ', 'F', 'I', 'L', 'E', 10, ' ', ' ', '-', '-', 'p', 'a', 'g', 'e', 's', 'i', 'z', 'e', ' ', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 't', 'h', 'e', ' ', 'p', 'a', 'g', 'e', ' ', 's', 'i', 'z', 'e', ' ', 't', 'o', ' ', 'N', 10, ' ', ' ', '-', '-', 'p', 'c', 'a', 'c', 'h', 'e', ' ', 'N', ' ', 'S', 'Z', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'C', 'o', 'n', 'f', 'i', 'g', 'u', 'r', 'e', ' ', 'N', ' ', 'p', 'a', 'g', 'e', 's', ' ', 'o', 'f', ' ', 'p', 'a', 'g', 'e', 'c', 'a', 'c', 'h', 'e', ' ', 'e', 'a', 'c', 'h', ' ', 'o', 'f', ' ', 's', 'i', 'z', 'e', ' ', 'S', 'Z', ' ', 'b', 'y', 't', 'e', 's', 10, ' ', ' ', '-', '-', 'p', 'r', 'i', 'm', 'a', 'r', 'y', 'k', 'e', 'y', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'U', 's', 'e', ' ', 'P', 'R', 'I', 'M', 'A', 'R', 'Y', ' ', 'K', 'E', 'Y', ' ', 'i', 'n', 's', 't', 'e', 'a', 'd', ' ', 'o', 'f', ' ', 'U', 'N', 'I', 'Q', 'U', 'E', ' ', 'w', 'h', 'e', 'r', 'e', ' ', 'a', 'p', 'p', 'r', 'o', 'p', 'r', 'i', 'a', 't', 'e', 10, ' ', ' ', '-', '-', 'r', 'e', 'p', 'e', 'a', 't', ' ', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'R', 'e', 'p', 'e', 'a', 't', ' ', 'e', 'a', 'c', 'h', ' ', 'S', 'E', 'L', 'E', 'C', 'T', ' ', 'N', ' ', 't', 'i', 'm', 'e', 's', ' ', '(', 'd', 'e', 'f', 'a', 'u', 'l', 't', ':', ' ', '1', ')', 10, ' ', ' ', '-', '-', 'r', 'e', 'p', 'r', 'e', 'p', 'a', 'r', 'e', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'R', 'e', 'p', 'r', 'e', 'p', 'a', 'r', 'e', ' ', 'e', 'a', 'c', 'h', ' ', 's', 't', 'a', 't', 'e', 'm', 'e', 'n', 't', ' ', 'u', 'p', 'o', 'n', ' ', 'e', 'v', 'e', 'r', 'y', ' ', 'i', 'n', 'v', 'o', 'c', 'a', 't', 'i', 'o', 'n', 10, ' ', ' ', '-', '-', 'r', 'e', 's', 'e', 'r', 'v', 'e', ' ', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'R', 'e', 's', 'e', 'r', 'v', 'e', ' ', 'N', ' ', 'b', 'y', 't', 'e', 's', ' ', 'o', 'n', ' ', 'e', 'a', 'c', 'h', ' ', 'd', 'a', 't', 'a', 'b', 'a', 's', 'e', ' ', 'p', 'a', 'g', 'e', 10, ' ', ' ', '-', '-', 's', 'c', 'r', 'i', 'p', 't', ' ', 'F', 'I', 'L', 'E', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'W', 'r', 'i', 't', 'e', ' ', 'a', 'n', ' ', 'S', 'Q', 'L', ' ', 's', 'c', 'r', 'i', 'p', 't', ' ', 'f', 'o', 'r', ' ', 't', 'h', 'e', ' ', 't', 'e', 's', 't', ' ', 'i', 'n', 't', 'o', ' ', 'F', 'I', 'L', 'E', 10, ' ', ' ', '-', '-', 's', 'e', 'r', 'i', 'a', 'l', 'i', 'z', 'e', 'd', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 's', 'e', 'r', 'i', 'a', 'l', 'i', 'z', 'e', 'd', ' ', 't', 'h', 'r', 'e', 'a', 'd', 'i', 'n', 'g', ' ', 'm', 'o', 'd', 'e', 10, ' ', ' ', '-', '-', 's', 'i', 'n', 'g', 'l', 'e', 't', 'h', 'r', 'e', 'a', 'd', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 's', 'i', 'n', 'g', 'l', 'e', '-', 't', 'h', 'r', 'e', 'a', 'd', 'e', 'd', ' ', 'm', 'o', 'd', 'e', ' ', '-', ' ', 'd', 'i', 's', 'a', 'b', 'l', 'e', 's', ' ', 'a', 'l', 'l', ' ', 'm', 'u', 't', 'e', 'x', 'i', 'n', 'g', 10, ' ', ' ', '-', '-', 's', 'q', 'l', 'o', 'n', 'l', 'y', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'N', 'o', '-', 'o', 'p', '.', ' ', ' ', 'O', 'n', 'l', 'y', ' ', 's', 'h', 'o', 'w', ' ', 't', 'h', 'e', ' ', 'S', 'Q', 'L', ' ', 't', 'h', 'a', 't', ' ', 'w', 'o', 'u', 'l', 'd', ' ', 'h', 'a', 'v', 'e', ' ', 'b', 'e', 'e', 'n', ' ', 'r', 'u', 'n', '.', 10, ' ', ' ', '-', '-', 's', 'h', 'r', 'i', 'n', 'k', '-', 'm', 'e', 'm', 'o', 'r', 'y', ' ', ' ', ' ', ' ', ' ', 'I', 'n', 'v', 'o', 'k', 'e', ' ', 's', 'q', 'l', 'i', 't', 'e', '3', '_', 'd', 'b', '_', 'r', 'e', 'l', 'e', 'a', 's', 'e', '_', 'm', 'e', 'm', 'o', 'r', 'y', '(', ')', ' ', 'f', 'r', 'e', 'q', 'u', 'e', 'n', 't', 'l', 'y', '.', 10, ' ', ' ', '-', '-', 's', 'i', 'z', 'e', ' ', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'R', 'e', 'l', 'a', 't', 'i', 'v', 'e', ' ', 't', 'e', 's', 't', ' ', 's', 'i', 'z', 'e', '.', ' ', ' ', 'D', 'e', 'f', 'a', 'u', 'l', 't', '=', '1', '0', '0', 10, ' ', ' ', '-', '-', 's', 'o', 'f', 't', '-', 'h', 'e', 'a', 'p', '-', 'l', 'i', 'm', 'i', 't', ' ', 'N', ' ', 'T', 'h', 'e', ' ', 's', 'o', 'f', 't', ' ', 'l', 'i', 'm', 'i', 't', ' ', 'o', 'n', ' ', 't', 'h', 'e', ' ', 'm', 'a', 'x', 'i', 'm', 'u', 'm', ' ', 'h', 'e', 'a', 'p', ' ', 's', 'i', 'z', 'e', 10, ' ', ' ', '-', '-', 's', 't', 'r', 'i', 'c', 't', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'U', 's', 'e', ' ', 'S', 'T', 'R', 'I', 'C', 'T', ' ', 't', 'a', 'b', 'l', 'e', ' ', 'w', 'h', 'e', 'r', 'e', ' ', 'a', 'p', 'p', 'r', 'o', 'p', 'r', 'i', 'a', 't', 'e', 10, ' ', ' ', '-', '-', 's', 't', 'a', 't', 's', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'h', 'o', 'w', ' ', 's', 't', 'a', 't', 'i', 's', 't', 'i', 'c', 's', ' ', 'a', 't', ' ', 't', 'h', 'e', ' ', 'e', 'n', 'd', 10, ' ', ' ', '-', '-', 's', 't', 'm', 't', 's', 'c', 'a', 'n', 's', 't', 'a', 't', 'u', 's', ' ', ' ', ' ', ' ', 'A', 'c', 't', 'i', 'v', 'a', 't', 'e', ' ', 'S', 'Q', 'L', 'I', 'T', 'E', '_', 'D', 'B', 'C', 'O', 'N', 'F', 'I', 'G', '_', 'S', 'T', 'M', 'T', '_', 'S', 'C', 'A', 'N', 'S', 'T', 'A', 'T', 'U', 'S', 10, ' ', ' ', '-', '-', 't', 'e', 'm', 'p', ' ', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'N', ' ', 'f', 'r', 'o', 'm', ' ', '0', ' ', 't', 'o', ' ', '9', '.', ' ', ' ', '0', ':', ' ', 'n', 'o', ' ', 't', 'e', 'm', 'p', ' ', 't', 'a', 'b', 'l', 'e', '.', ' ', '9', ':', ' ', 'a', 'l', 'l', ' ', 't', 'e', 'm', 'p', ' ', 't', 'a', 'b', 'l', 'e', 's', 10, ' ', ' ', '-', '-', 't', 'e', 's', 't', 's', 'e', 't', ' ', 'T', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'R', 'u', 'n', ' ', 't', 'e', 's', 't', '-', 's', 'e', 't', ' ', 'T', ' ', '(', 'm', 'a', 'i', 'n', ',', ' ', 'c', 't', 'e', ',', ' ', 'r', 't', 'r', 'e', 'e', ',', ' ', 'o', 'r', 'm', ',', ' ', 'f', 'p', ',', ' ', 'j', 's', 'o', 'n', ',', 10, ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 's', 't', 'a', 'r', ',', ' ', 'a', 'p', 'p', ',', ' ', 'd', 'e', 'b', 'u', 'g', ')', '.', ' ', ' ', 'C', 'a', 'n', ' ', 'b', 'e', ' ', 'a', ' ', 'c', 'o', 'm', 'm', 'a', '-', 's', 'e', 'p', 'a', 'r', 'a', 't', 'e', 'd', ' ', 'l', 'i', 's', 't', 10, ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'o', 'f', ' ', 'v', 'a', 'l', 'u', 'e', 's', ',', ' ', 'w', 'i', 't', 'h', ' ', '/', 'S', 'C', 'A', 'L', 'E', ' ', 's', 'u', 'f', 'f', 'i', 'x', 'e', 's', ' ', 'o', 'r', ' ', 'm', 'a', 'c', 'r', 'o', ' ', '"', 'm', 'i', 'x', '1', '"', 10, ' ', ' ', '-', '-', 't', 'r', 'a', 'c', 'e', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'T', 'u', 'r', 'n', ' ', 'o', 'n', ' ', 'S', 'Q', 'L', ' ', 't', 'r', 'a', 'c', 'i', 'n', 'g', 10, ' ', ' ', '-', '-', 't', 'h', 'r', 'e', 'a', 'd', 's', ' ', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'U', 's', 'e', ' ', 'u', 'p', ' ', 't', 'o', ' ', 'N', ' ', 't', 'h', 'r', 'e', 'a', 'd', 's', ' ', 'f', 'o', 'r', ' ', 's', 'o', 'r', 't', 'i', 'n', 'g', 10, ' ', ' ', '-', '-', 'u', 't', 'f', '1', '6', 'b', 'e', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 't', 'e', 'x', 't', ' ', 'e', 'n', 'c', 'o', 'd', 'i', 'n', 'g', ' ', 't', 'o', ' ', 'U', 'T', 'F', '-', '1', '6', 'B', 'E', 10, ' ', ' ', '-', '-', 'u', 't', 'f', '1', '6', 'l', 'e', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 't', 'e', 'x', 't', ' ', 'e', 'n', 'c', 'o', 'd', 'i', 'n', 'g', ' ', 't', 'o', ' ', 'U', 'T', 'F', '-', '1', '6', 'L', 'E', 10, ' ', ' ', '-', '-', 'v', 'e', 'r', 'i', 'f', 'y', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'R', 'u', 'n', ' ', 'a', 'd', 'd', 'i', 't', 'i', 'o', 'n', 'a', 'l', ' ', 'v', 'e', 'r', 'i', 'f', 'i', 'c', 'a', 't', 'i', 'o', 'n', ' ', 's', 't', 'e', 'p', 's', 10, ' ', ' ', '-', '-', 'v', 'f', 's', ' ', 'N', 'A', 'M', 'E', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'U', 's', 'e', ' ', 't', 'h', 'e', ' ', 'g', 'i', 'v', 'e', 'n', ' ', '(', 'p', 'r', 'e', 'i', 'n', 's', 't', 'a', 'l', 'l', 'e', 'd', ')', ' ', 'V', 'F', 'S', 10, ' ', ' ', '-', '-', 'w', 'i', 't', 'h', 'o', 'u', 't', '-', 'r', 'o', 'w', 'i', 'd', ' ', ' ', ' ', ' ', ' ', 'U', 's', 'e', ' ', 'W', 'I', 'T', 'H', 'O', 'U', 'T', ' ', 'R', 'O', 'W', 'I', 'D', ' ', 'w', 'h', 'e', 'r', 'e', ' ', 'a', 'p', 'p', 'r', 'o', 'p', 'r', 'i', 'a', 't', 'e', 10}

type va_list = uintptr

type sqlite3_int64 = int64

type sqlite3_vfs = struct {
	FiVersion          int32
	FszOsFile          int32
	FmxPathname        int32
	FpNext             uintptr
	FzName             uintptr
	FpAppData          uintptr
	FxOpen             uintptr
	FxDelete           uintptr
	FxAccess           uintptr
	FxFullPathname     uintptr
	FxDlOpen           uintptr
	FxDlError          uintptr
	FxDlSym            uintptr
	FxDlClose          uintptr
	FxRandomness       uintptr
	FxSleep            uintptr
	FxCurrentTime      uintptr
	FxGetLastError     uintptr
	FxCurrentTimeInt64 uintptr
	FxSetSystemCall    uintptr
	FxGetSystemCall    uintptr
	FxNextSystemCall   uintptr
}

type sqlite3_rtree_geometry = struct {
	FpContext uintptr
	FnParam   int32
	FaParam   uintptr
	FpUser    uintptr
	FxDelUser uintptr
}

type sqlite3_rtree_dbl = float64

type __uint32_t = uint32

type __darwin_ct_rune_t = int32

type u64 = uint64

// C documentation
//
//	/*
//	** State structure for a Hash hash in progress
//	*/
type HashContext = struct {
	FisInit uint8
	Fi      uint8
	Fj      uint8
	Fs      [256]uint8
	Fr      [32]uint8
}

// C documentation
//
//	/* All global state is held in this structure */
type Global = struct {
	Fdb                uintptr
	FzDbName           uintptr
	FzVfs              uintptr
	FpStmt             uintptr
	FiStart            sqlite3_int64
	FiTotal            sqlite3_int64
	FbWithoutRowid     int32
	FbReprepare        int32
	FbSqlOnly          int32
	FbExplain          int32
	FbVerify           int32
	FbMemShrink        int32
	FeTemp             int32
	FszTest            int32
	FszBase            int32
	FnRepeat           int32
	FdoCheckpoint      int32
	FnReserve          int32
	FstmtScanStatus    int32
	FdoBigTransactions int32
	FzWR               uintptr
	FzNN               uintptr
	FzPK               uintptr
	Fx                 uint32
	Fy                 uint32
	FnResByte          u64
	FnResult           int32
	FzResult           [3000]int8
	FpScript           uintptr
	FhashFile          uintptr
	Fhash              HashContext
}

// C documentation
//
//	/* All global state is held in this structure */
var g Global

// C documentation
//
//	/* Return " TEMP" or "", as appropriate for creating a table.
//	*/
func isTemp(tls *libc.TLS, N int32) (r uintptr) {
	var v1 uintptr
	_ = v1
	if g.FeTemp >= N {
		v1 = __ccgo_ts
	} else {
		v1 = __ccgo_ts + 6
	}
	return v1
}

// C documentation
//
//	/* Print an error message and exit */
func fatal_error(tls *libc.TLS, zMsg uintptr, va uintptr) {
	var ap va_list
	_ = ap
	ap = va
	libc.Xvfprintf(tls, libc.X__stderrp, zMsg, ap)
	_ = ap
	libc.Xexit(tls, int32(1))
}

/****************************************************************************
** Hash algorithm used to verify that compilation is not miscompiled
** in such a was as to generate an incorrect result.
 */

// C documentation
//
//	/*
//	** Initialize a new hash.  iSize determines the size of the hash
//	** in bits and should be one of 224, 256, 384, or 512.  Or iSize
//	** can be zero to use the default hash size of 256 bits.
//	*/
func HashInit(tls *libc.TLS) {
	var k uint32
	_ = k
	g.Fhash.Fi = uint8(0)
	g.Fhash.Fj = uint8(0)
	k = uint32(0)
	for {
		if !(k < uint32(256)) {
			break
		}
		**(**uint8)(__ccgo_up(uintptr(unsafe.Pointer(&g)) + 3168 + 3 + uintptr(k))) = uint8(k)
		goto _1
	_1:
		;
		k = k + 1
	}
}

// C documentation
//
//	/*
//	** Make consecutive calls to the HashUpdate function to add new content
//	** to the hash
//	*/
func HashUpdate(tls *libc.TLS, aData uintptr, nData uint32) {
	var i, j, t uint8
	var k uint32
	_, _, _, _ = i, j, k, t
	i = g.Fhash.Fi
	j = g.Fhash.Fj
	if g.FhashFile != 0 {
		libc.Xfwrite(tls, aData, uint64(1), uint64(nData), g.FhashFile)
	}
	k = uint32(0)
	for {
		if !(k < nData) {
			break
		}
		j = libc.Uint8FromInt32(int32(j) + (libc.Int32FromUint8(**(**uint8)(__ccgo_up(uintptr(unsafe.Pointer(&g)) + 3168 + 3 + uintptr(i)))) + libc.Int32FromUint8(**(**uint8)(__ccgo_up(aData + uintptr(k))))))
		t = **(**uint8)(__ccgo_up(uintptr(unsafe.Pointer(&g)) + 3168 + 3 + uintptr(j)))
		**(**uint8)(__ccgo_up(uintptr(unsafe.Pointer(&g)) + 3168 + 3 + uintptr(j))) = **(**uint8)(__ccgo_up(uintptr(unsafe.Pointer(&g)) + 3168 + 3 + uintptr(i)))
		**(**uint8)(__ccgo_up(uintptr(unsafe.Pointer(&g)) + 3168 + 3 + uintptr(i))) = t
		i = i + 1
		goto _1
	_1:
		;
		k = k + 1
	}
	g.Fhash.Fi = i
	g.Fhash.Fj = j
}

// C documentation
//
//	/*
//	** After all content has been added, invoke HashFinal() to compute
//	** the final hash.  The hash result is stored in g.hash.r[].
//	*/
func HashFinal(tls *libc.TLS) {
	var i, j, t uint8
	var k uint32
	_, _, _, _ = i, j, k, t
	i = g.Fhash.Fi
	j = g.Fhash.Fj
	k = uint32(0)
	for {
		if !(k < uint32(32)) {
			break
		}
		i = i + 1
		t = **(**uint8)(__ccgo_up(uintptr(unsafe.Pointer(&g)) + 3168 + 3 + uintptr(i)))
		j = libc.Uint8FromInt32(int32(j) + libc.Int32FromUint8(t))
		**(**uint8)(__ccgo_up(uintptr(unsafe.Pointer(&g)) + 3168 + 3 + uintptr(i))) = **(**uint8)(__ccgo_up(uintptr(unsafe.Pointer(&g)) + 3168 + 3 + uintptr(j)))
		**(**uint8)(__ccgo_up(uintptr(unsafe.Pointer(&g)) + 3168 + 3 + uintptr(j))) = t
		t = libc.Uint8FromInt32(int32(t) + libc.Int32FromUint8(**(**uint8)(__ccgo_up(uintptr(unsafe.Pointer(&g)) + 3168 + 3 + uintptr(i)))))
		**(**uint8)(__ccgo_up(uintptr(unsafe.Pointer(&g)) + 3168 + 259 + uintptr(k))) = **(**uint8)(__ccgo_up(uintptr(unsafe.Pointer(&g)) + 3168 + 3 + uintptr(t)))
		goto _1
	_1:
		;
		k = k + 1
	}
}

/* End of the Hash hashing logic
*****************************************************************************/

// C documentation
//
//	/*
//	** Return the value of a hexadecimal digit.  Return -1 if the input
//	** is not a hex digit.
//	*/
func hexDigitValue(tls *libc.TLS, c int8) (r int32) {
	if int32(c) >= int32('0') && int32(c) <= int32('9') {
		return int32(c) - int32('0')
	}
	if int32(c) >= int32('a') && int32(c) <= int32('f') {
		return int32(c) - int32('a') + int32(10)
	}
	if int32(c) >= int32('A') && int32(c) <= int32('F') {
		return int32(c) - int32('A') + int32(10)
	}
	return -int32(1)
}

/* Provide an alternative to sqlite3_stricmp() in older versions of
** SQLite */

// C documentation
//
//	/*
//	** Interpret zArg as an integer value, possibly with suffixes.
//	*/
func integerValue(tls *libc.TLS, zArg uintptr) (r int32) {
	var i, isNeg, x, v1, v2 int32
	var v sqlite3_int64
	var v4, v5 __darwin_ct_rune_t
	var v9 int64
	_, _, _, _, _, _, _, _, _ = i, isNeg, v, x, v1, v2, v4, v5, v9
	v = 0
	isNeg = 0
	if int32(**(**int8)(__ccgo_up(zArg))) == int32('-') {
		isNeg = int32(1)
		zArg = zArg + 1
	} else {
		if int32(**(**int8)(__ccgo_up(zArg))) == int32('+') {
			zArg = zArg + 1
		}
	}
	if int32(**(**int8)(__ccgo_up(zArg))) == int32('0') && int32(**(**int8)(__ccgo_up(zArg + 1))) == int32('x') {
		zArg = zArg + uintptr(2)
		for {
			v1 = hexDigitValue(tls, **(**int8)(__ccgo_up(zArg)))
			x = v1
			if !(v1 >= 0) {
				break
			}
			v = v<<libc.Int32FromInt32(4) + int64(x)
			zArg = zArg + 1
		}
	} else {
		for {
			v4 = int32(**(**int8)(__ccgo_up(zArg)))
			if v4 < 0 || v4 >= libc.Int32FromInt32(1)<<libc.Int32FromInt32(8) {
				v2 = 0
			} else {
				v2 = libc.BoolInt32(!!(uint64(**(**__uint32_t)(__ccgo_up(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v4)*4)))&uint64(0x00000400) != 0))
			}
			v5 = v2
			goto _6
		_6:
			v1 = v5
			goto _3
		_3:
			if !(v1 != 0) {
				break
			}
			v = v*int64(10) + int64(**(**int8)(__ccgo_up(zArg))) - int64('0')
			zArg = zArg + 1
		}
	}
	i = 0
	for {
		if !(libc.Uint64FromInt32(i) < libc.Uint64FromInt64(144)/libc.Uint64FromInt64(16)) {
			break
		}
		if libsqlite3.Xsqlite3_stricmp(tls, aMult[i].FzSuffix, zArg) == 0 {
			v = v * int64(aMult[i].FiMult)
			break
		}
		goto _8
	_8:
		;
		i = i + 1
	}
	if v > int64(0x7fffffff) {
		fatal_error(tls, __ccgo_ts+34, 0)
	}
	if isNeg != 0 {
		v9 = -v
	} else {
		v9 = v
	}
	return int32(v9)
}

var aMult = [9]struct {
	FzSuffix uintptr
	FiMult   int32
}{
	0: {
		FzSuffix: __ccgo_ts + 7,
		FiMult:   int32(1024),
	},
	1: {
		FzSuffix: __ccgo_ts + 11,
		FiMult:   libc.Int32FromInt32(1024) * libc.Int32FromInt32(1024),
	},
	2: {
		FzSuffix: __ccgo_ts + 15,
		FiMult:   libc.Int32FromInt32(1024) * libc.Int32FromInt32(1024) * libc.Int32FromInt32(1024),
	},
	3: {
		FzSuffix: __ccgo_ts + 19,
		FiMult:   int32(1000),
	},
	4: {
		FzSuffix: __ccgo_ts + 22,
		FiMult:   int32(1000000),
	},
	5: {
		FzSuffix: __ccgo_ts + 25,
		FiMult:   int32(1000000000),
	},
	6: {
		FzSuffix: __ccgo_ts + 28,
		FiMult:   int32(1000),
	},
	7: {
		FzSuffix: __ccgo_ts + 30,
		FiMult:   int32(1000000),
	},
	8: {
		FzSuffix: __ccgo_ts + 32,
		FiMult:   int32(1000000000),
	},
}

// C documentation
//
//	/* Return the current wall-clock time, in milliseconds */
func speedtest1_timestamp(tls *libc.TLS) (r sqlite3_int64) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* r at bp+8 */ float64
	var _ /* t at bp+0 */ sqlite3_int64
	if clockVfs == uintptr(0) {
		clockVfs = libsqlite3.Xsqlite3_vfs_find(tls, uintptr(0))
	}
	if (*sqlite3_vfs)(unsafe.Pointer(clockVfs)).FiVersion >= int32(2) && (*sqlite3_vfs)(unsafe.Pointer(clockVfs)).FxCurrentTimeInt64 != uintptr(0) {
		(*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*sqlite3_vfs)(unsafe.Pointer(clockVfs)).FxCurrentTimeInt64})))(tls, clockVfs, bp)
	} else {
		(*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*sqlite3_vfs)(unsafe.Pointer(clockVfs)).FxCurrentTime})))(tls, clockVfs, bp+8)
		**(**sqlite3_int64)(__ccgo_up(bp)) = int64(float64(**(**float64)(__ccgo_up(bp + 8)) * libc.Float64FromFloat64(8.64e+07)))
	}
	return **(**sqlite3_int64)(__ccgo_up(bp))
}

var clockVfs = uintptr(0)

// C documentation
//
//	/* Return a pseudo-random unsigned integer */
func speedtest1_random(tls *libc.TLS) (r uint32) {
	g.Fx = g.Fx>>int32(1) ^ (uint32(1)+^(g.Fx&libc.Uint32FromInt32(1)))&uint32(0xd0000001)
	g.Fy = g.Fy*uint32(1103515245) + uint32(12345)
	return g.Fx ^ g.Fy
}

// C documentation
//
//	/* Map the value in within the range of 1...limit into another
//	** number in a way that is chatic and invertable.
//	*/
func swizzle(tls *libc.TLS, in uint32, limit uint32) (r uint32) {
	var out uint32
	_ = out
	out = uint32(0)
	for limit != 0 {
		out = out<<int32(1) | in&uint32(1)
		in = in >> uint32(1)
		limit = limit >> uint32(1)
	}
	return out
}

// C documentation
//
//	/* Round up a number so that it is a power of two minus one
//	*/
func roundup_allones(tls *libc.TLS, limit uint32) (r uint32) {
	var m uint32
	_ = m
	m = uint32(1)
	for m < limit {
		m = m<<int32(1) + uint32(1)
	}
	return m
}

// C documentation
//
//	/* The speedtest1_numbername procedure below converts its argment (an integer)
//	** into a string which is the English-language name for that number.
//	** The returned string should be freed with sqlite3_free().
//	**
//	** Example:
//	**
//	**     speedtest1_numbername(123)   ->  "one hundred twenty three"
//	*/
func speedtest1_numbername(tls *libc.TLS, n uint32, zOut uintptr, nOut int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i, v1 int32
	_, _ = i, v1
	i = 0
	if n >= uint32(1000000000) {
		i = i + speedtest1_numbername(tls, n/uint32(1000000000), zOut+uintptr(i), nOut-i)
		libsqlite3.Xsqlite3_snprintf(tls, nOut-i, zOut+uintptr(i), __ccgo_ts+255, 0)
		i = i + libc.Int32FromUint64(libc.Xstrlen(tls, zOut+uintptr(i)))
		n = n % uint32(1000000000)
	}
	if n >= uint32(1000000) {
		if i != 0 && i < nOut-int32(1) {
			v1 = i
			i = i + 1
			**(**int8)(__ccgo_up(zOut + uintptr(v1))) = int8(' ')
		}
		i = i + speedtest1_numbername(tls, n/uint32(1000000), zOut+uintptr(i), nOut-i)
		libsqlite3.Xsqlite3_snprintf(tls, nOut-i, zOut+uintptr(i), __ccgo_ts+264, 0)
		i = i + libc.Int32FromUint64(libc.Xstrlen(tls, zOut+uintptr(i)))
		n = n % uint32(1000000)
	}
	if n >= uint32(1000) {
		if i != 0 && i < nOut-int32(1) {
			v1 = i
			i = i + 1
			**(**int8)(__ccgo_up(zOut + uintptr(v1))) = int8(' ')
		}
		i = i + speedtest1_numbername(tls, n/uint32(1000), zOut+uintptr(i), nOut-i)
		libsqlite3.Xsqlite3_snprintf(tls, nOut-i, zOut+uintptr(i), __ccgo_ts+273, 0)
		i = i + libc.Int32FromUint64(libc.Xstrlen(tls, zOut+uintptr(i)))
		n = n % uint32(1000)
	}
	if n >= uint32(100) {
		if i != 0 && i < nOut-int32(1) {
			v1 = i
			i = i + 1
			**(**int8)(__ccgo_up(zOut + uintptr(v1))) = int8(' ')
		}
		libsqlite3.Xsqlite3_snprintf(tls, nOut-i, zOut+uintptr(i), __ccgo_ts+283, libc.VaList(bp+8, ones[n/uint32(100)]))
		i = i + libc.Int32FromUint64(libc.Xstrlen(tls, zOut+uintptr(i)))
		n = n % uint32(100)
	}
	if n >= uint32(20) {
		if i != 0 && i < nOut-int32(1) {
			v1 = i
			i = i + 1
			**(**int8)(__ccgo_up(zOut + uintptr(v1))) = int8(' ')
		}
		libsqlite3.Xsqlite3_snprintf(tls, nOut-i, zOut+uintptr(i), __ccgo_ts+294, libc.VaList(bp+8, tens[n/uint32(10)]))
		i = i + libc.Int32FromUint64(libc.Xstrlen(tls, zOut+uintptr(i)))
		n = n % uint32(10)
	}
	if n > uint32(0) {
		if i != 0 && i < nOut-int32(1) {
			v1 = i
			i = i + 1
			**(**int8)(__ccgo_up(zOut + uintptr(v1))) = int8(' ')
		}
		libsqlite3.Xsqlite3_snprintf(tls, nOut-i, zOut+uintptr(i), __ccgo_ts+294, libc.VaList(bp+8, ones[n]))
		i = i + libc.Int32FromUint64(libc.Xstrlen(tls, zOut+uintptr(i)))
	}
	if i == 0 {
		libsqlite3.Xsqlite3_snprintf(tls, nOut-i, zOut+uintptr(i), __ccgo_ts+71, 0)
		i = i + libc.Int32FromUint64(libc.Xstrlen(tls, zOut+uintptr(i)))
	}
	return i
}

var ones = [20]uintptr{
	0:  __ccgo_ts + 71,
	1:  __ccgo_ts + 76,
	2:  __ccgo_ts + 80,
	3:  __ccgo_ts + 84,
	4:  __ccgo_ts + 90,
	5:  __ccgo_ts + 95,
	6:  __ccgo_ts + 100,
	7:  __ccgo_ts + 104,
	8:  __ccgo_ts + 110,
	9:  __ccgo_ts + 116,
	10: __ccgo_ts + 121,
	11: __ccgo_ts + 125,
	12: __ccgo_ts + 132,
	13: __ccgo_ts + 139,
	14: __ccgo_ts + 148,
	15: __ccgo_ts + 157,
	16: __ccgo_ts + 165,
	17: __ccgo_ts + 173,
	18: __ccgo_ts + 183,
	19: __ccgo_ts + 192,
}

var tens = [10]uintptr{
	0: __ccgo_ts + 6,
	1: __ccgo_ts + 121,
	2: __ccgo_ts + 201,
	3: __ccgo_ts + 208,
	4: __ccgo_ts + 215,
	5: __ccgo_ts + 221,
	6: __ccgo_ts + 227,
	7: __ccgo_ts + 233,
	8: __ccgo_ts + 241,
	9: __ccgo_ts + 248,
}

// C documentation
//
//	/* Start a new test case */
var zDots = [72]int8{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'}
var iTestNumber = int32(0) /* Current test # for begin/end_test(). */
func speedtest1_begin_test(tls *libc.TLS, iTestNum int32, zTestName uintptr, va uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var ap va_list
	var n int32
	var zName uintptr
	_, _, _ = ap, n, zName
	n = libc.Int32FromUint64(libc.Xstrlen(tls, zTestName))
	iTestNumber = iTestNum
	ap = va
	zName = libsqlite3.Xsqlite3_vmprintf(tls, zTestName, ap)
	_ = ap
	n = libc.Int32FromUint64(libc.Xstrlen(tls, zName))
	if n > int32(NAMEWIDTH) {
		**(**int8)(__ccgo_up(zName + 60)) = 0
		n = int32(NAMEWIDTH)
	}
	if g.FpScript != 0 {
		libc.Xfprintf(tls, g.FpScript, __ccgo_ts+297, libc.VaList(bp+8, iTestNumber, n, zName))
	}
	if g.FbSqlOnly != 0 {
		libc.Xprintf(tls, __ccgo_ts+320, libc.VaList(bp+8, iTestNum, zName, int32(NAMEWIDTH)-n, uintptr(unsafe.Pointer(&zDots))))
	} else {
		libc.Xprintf(tls, __ccgo_ts+340, libc.VaList(bp+8, iTestNum, zName, int32(NAMEWIDTH)-n, uintptr(unsafe.Pointer(&zDots))))
		libc.Xfflush(tls, libc.X__stdoutp)
	}
	libsqlite3.Xsqlite3_free(tls, zName)
	g.FnResult = 0
	g.FiStart = speedtest1_timestamp(tls)
	g.Fx = uint32(0xad131d0b)
	g.Fy = uint32(0x44f9eac8)
}

// C documentation
//
//	/* Complete a test case */
func speedtest1_end_test(tls *libc.TLS) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var iElapseTime sqlite3_int64
	_ = iElapseTime
	iElapseTime = speedtest1_timestamp(tls) - g.FiStart
	if g.FdoCheckpoint != 0 {
		speedtest1_exec(tls, __ccgo_ts+354, 0)
	}
	if g.FpScript != 0 {
		libc.Xfprintf(tls, g.FpScript, __ccgo_ts+377, libc.VaList(bp+8, iTestNumber))
	}
	if !(g.FbSqlOnly != 0) {
		g.FiTotal += iElapseTime
		libc.Xprintf(tls, __ccgo_ts+393, libc.VaList(bp+8, int32(iElapseTime/libc.Int64FromInt32(1000)), int32(iElapseTime%libc.Int64FromInt32(1000))))
	}
	if g.FpStmt != 0 {
		libsqlite3.Xsqlite3_finalize(tls, g.FpStmt)
		g.FpStmt = uintptr(0)
	}
	iTestNumber = 0
}

// C documentation
//
//	/* Report end of testing */
func speedtest1_final(tls *libc.TLS) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var i int32
	_ = i
	if !(g.FbSqlOnly != 0) {
		libc.Xprintf(tls, __ccgo_ts+404, libc.VaList(bp+8, libc.Int32FromInt32(NAMEWIDTH)-libc.Int32FromInt32(5), uintptr(unsafe.Pointer(&zDots)), int32(g.FiTotal/libc.Int64FromInt32(1000)), int32(g.FiTotal%libc.Int64FromInt32(1000))))
	}
	if g.FbVerify != 0 {
		libc.Xprintf(tls, __ccgo_ts+432, libc.VaList(bp+8, g.FnResByte))
		HashUpdate(tls, __ccgo_ts+457, uint32(1))
		HashFinal(tls)
		i = 0
		for {
			if !(i < int32(24)) {
				break
			}
			libc.Xprintf(tls, __ccgo_ts+459, libc.VaList(bp+8, libc.Int32FromUint8(**(**uint8)(__ccgo_up(uintptr(unsafe.Pointer(&g)) + 3168 + 259 + uintptr(i))))))
			goto _1
		_1:
			;
			i = i + 1
		}
		if g.FhashFile != 0 && g.FhashFile != libc.X__stdoutp {
			libc.Xfclose(tls, g.FhashFile)
		}
		libc.Xprintf(tls, __ccgo_ts+457, 0)
	}
}

// C documentation
//
//	/* Print an SQL statement to standard output */
func printSql(tls *libc.TLS, zSql uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var n, v1, v5, v7, v8 int32
	var v10, v11 bool
	var v3 __darwin_ct_rune_t
	var v4 uint64
	_, _, _, _, _, _, _, _, _ = n, v1, v10, v11, v3, v4, v5, v7, v8
	n = libc.Int32FromUint64(libc.Xstrlen(tls, zSql))
	for {
		if v11 = n > 0; v11 {
			if v10 = int32(**(**int8)(__ccgo_up(zSql + uintptr(n-int32(1))))) == int32(';'); !v10 {
				v3 = libc.Int32FromUint8(libc.Uint8FromInt8(**(**int8)(__ccgo_up(zSql + uintptr(n-int32(1))))))
				v4 = uint64(0x00004000)
				v8 = libc.BoolInt32(v3 & ^libc.Int32FromInt32(0x7F) == 0)
				goto _9
			_9:
				if v8 != 0 {
					v7 = libc.BoolInt32(!!(uint64(**(**__uint32_t)(__ccgo_up(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v3)*4)))&v4 != 0))
				} else {
					v7 = libc.BoolInt32(!!(libc.X__maskrune(tls, v3, v4) != 0))
				}
				v5 = v7
				goto _6
			_6:
				v1 = v5
				goto _2
			_2:
			}
		}
		if !(v11 && (v10 || v1 != 0)) {
			break
		}
		n = n - 1
	}
	if g.FbExplain != 0 {
		libc.Xprintf(tls, __ccgo_ts+464, 0)
	}
	libc.Xprintf(tls, __ccgo_ts+473, libc.VaList(bp+8, n, zSql))
	if g.FbExplain != 0 && (libsqlite3.Xsqlite3_strglob(tls, __ccgo_ts+480, zSql) == 0 || libsqlite3.Xsqlite3_strglob(tls, __ccgo_ts+489, zSql) == 0 || libsqlite3.Xsqlite3_strglob(tls, __ccgo_ts+496, zSql) == 0) {
		libc.Xprintf(tls, __ccgo_ts+473, libc.VaList(bp+8, n, zSql))
	}
}

// C documentation
//
//	/* Shrink memory used, if appropriate and if the SQLite version is capable
//	** of doing so.
//	*/
func speedtest1_shrink_memory(tls *libc.TLS) {
	if g.FbMemShrink != 0 {
		libsqlite3.Xsqlite3_db_release_memory(tls, g.Fdb)
	}
}

// C documentation
//
//	/* Run SQL */
func speedtest1_exec(tls *libc.TLS, zFormat uintptr, va uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var ap va_list
	var rc int32
	var zSql uintptr
	var _ /* zErrMsg at bp+0 */ uintptr
	_, _, _ = ap, rc, zSql
	ap = va
	zSql = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	if g.FbSqlOnly != 0 {
		printSql(tls, zSql)
	} else {
		**(**uintptr)(__ccgo_up(bp)) = uintptr(0)
		if g.FpScript != 0 {
			libc.Xfprintf(tls, g.FpScript, __ccgo_ts+504, libc.VaList(bp+16, zSql))
		}
		rc = libsqlite3.Xsqlite3_exec(tls, g.Fdb, zSql, uintptr(0), uintptr(0), bp)
		if **(**uintptr)(__ccgo_up(bp)) != 0 {
			fatal_error(tls, __ccgo_ts+509, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(bp)), zSql))
		}
		if rc != SQLITE_OK {
			fatal_error(tls, __ccgo_ts+527, libc.VaList(bp+16, libsqlite3.Xsqlite3_errmsg(tls, g.Fdb)))
		}
	}
	libsqlite3.Xsqlite3_free(tls, zSql)
	speedtest1_shrink_memory(tls)
}

// C documentation
//
//	/* Run SQL and return the first column of the first row as a string.  The
//	** returned string is obtained from sqlite_malloc() and must be freed by
//	** the caller.
//	*/
func speedtest1_once(tls *libc.TLS, zFormat uintptr, va uintptr) (r uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var ap va_list
	var rc1 int32
	var z, z1, zResult, zSql uintptr
	var _ /* pStmt at bp+0 */ uintptr
	_, _, _, _, _, _ = ap, rc1, z, z1, zResult, zSql
	zResult = uintptr(0)
	ap = va
	zSql = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	if g.FbSqlOnly != 0 {
		printSql(tls, zSql)
	} else {
		rc1 = libsqlite3.Xsqlite3_prepare_v2(tls, g.Fdb, zSql, -int32(1), bp, uintptr(0))
		if rc1 != 0 {
			fatal_error(tls, __ccgo_ts+543, libc.VaList(bp+16, libsqlite3.Xsqlite3_errmsg(tls, g.Fdb)))
		}
		if g.FpScript != 0 {
			z = libsqlite3.Xsqlite3_expanded_sql(tls, **(**uintptr)(__ccgo_up(bp)))
			libc.Xfprintf(tls, g.FpScript, __ccgo_ts+558, libc.VaList(bp+16, z))
			libsqlite3.Xsqlite3_free(tls, z)
		}
		if libsqlite3.Xsqlite3_step(tls, **(**uintptr)(__ccgo_up(bp))) == int32(SQLITE_ROW) {
			z1 = libsqlite3.Xsqlite3_column_text(tls, **(**uintptr)(__ccgo_up(bp)), 0)
			if z1 != 0 {
				zResult = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+294, libc.VaList(bp+16, z1))
			}
		}
		rc1 = libsqlite3.Xsqlite3_reset(tls, **(**uintptr)(__ccgo_up(bp)))
		if rc1 != SQLITE_OK {
			fatal_error(tls, __ccgo_ts+562, libc.VaList(bp+16, libsqlite3.Xsqlite3_sql(tls, **(**uintptr)(__ccgo_up(bp))), rc1, libsqlite3.Xsqlite3_errmsg(tls, g.Fdb)))
		}
		libsqlite3.Xsqlite3_finalize(tls, **(**uintptr)(__ccgo_up(bp)))
	}
	libsqlite3.Xsqlite3_free(tls, zSql)
	speedtest1_shrink_memory(tls)
	return zResult
}

// C documentation
//
//	/* Prepare an SQL statement */
func speedtest1_prepare(tls *libc.TLS, zFormat uintptr, va uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var ap va_list
	var rc int32
	var zSql uintptr
	_, _, _ = ap, rc, zSql
	ap = va
	zSql = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	if g.FbSqlOnly != 0 {
		printSql(tls, zSql)
	} else {
		if g.FpStmt != 0 {
			libsqlite3.Xsqlite3_finalize(tls, g.FpStmt)
		}
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, g.Fdb, zSql, -int32(1), uintptr(unsafe.Pointer(&g))+24, uintptr(0))
		if rc != 0 {
			fatal_error(tls, __ccgo_ts+543, libc.VaList(bp+8, libsqlite3.Xsqlite3_errmsg(tls, g.Fdb)))
		}
	}
	libsqlite3.Xsqlite3_free(tls, zSql)
}

// C documentation
//
//	/* Run an SQL statement previously prepared */
func speedtest1_run(tls *libc.TLS) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var aBlob, z, z1, v4 uintptr
	var eType, i, iBlob, len1, n, nBlob, rc, v3 int32
	var _ /* pNew at bp+8 */ uintptr
	var _ /* zChar at bp+2 */ [2]uint8
	var _ /* zPrefix at bp+0 */ [2]uint8
	_, _, _, _, _, _, _, _, _, _, _, _ = aBlob, eType, i, iBlob, len1, n, nBlob, rc, z, z1, v3, v4
	if g.FbSqlOnly != 0 {
		return
	}
	g.FnResult = 0
	if g.FpScript != 0 {
		z = libsqlite3.Xsqlite3_expanded_sql(tls, g.FpStmt)
		libc.Xfprintf(tls, g.FpScript, __ccgo_ts+558, libc.VaList(bp+24, z))
		libsqlite3.Xsqlite3_free(tls, z)
	}
	for libsqlite3.Xsqlite3_step(tls, g.FpStmt) == int32(SQLITE_ROW) {
		n = libsqlite3.Xsqlite3_column_count(tls, g.FpStmt)
		i = 0
		for {
			if !(i < n) {
				break
			}
			z1 = libsqlite3.Xsqlite3_column_text(tls, g.FpStmt, i)
			if z1 == uintptr(0) {
				z1 = __ccgo_ts + 584
			}
			len1 = libc.Int32FromUint64(libc.Xstrlen(tls, z1))
			if g.FbVerify != 0 {
				eType = libsqlite3.Xsqlite3_column_type(tls, g.FpStmt, i)
				(**(**[2]uint8)(__ccgo_up(bp)))[0] = uint8('\n')
				(**(**[2]uint8)(__ccgo_up(bp)))[int32(1)] = libc.Uint8FromInt8(**(**int8)(__ccgo_up(__ccgo_ts + 588 + uintptr(eType))))
				if g.FnResByte != 0 {
					HashUpdate(tls, bp, uint32(2))
				} else {
					HashUpdate(tls, bp+uintptr(1), uint32(1))
				}
				if eType == int32(SQLITE_FLOAT) {
					/* Omit the value of floating-point results from the verification
					 ** hash.  The only thing we record is the fact that the result was
					 ** a floating-point value. */
					g.FnResByte += uint64(2)
				} else {
					if eType == int32(SQLITE_BLOB) {
						nBlob = libsqlite3.Xsqlite3_column_bytes(tls, g.FpStmt, i)
						aBlob = libsqlite3.Xsqlite3_column_blob(tls, g.FpStmt, i)
						iBlob = 0
						for {
							if !(iBlob < nBlob) {
								break
							}
							(**(**[2]uint8)(__ccgo_up(bp + 2)))[0] = libc.Uint8FromInt8(**(**int8)(__ccgo_up(__ccgo_ts + 595 + uintptr(libc.Int32FromUint8(**(**uint8)(__ccgo_up(aBlob + uintptr(iBlob))))>>int32(4)))))
							(**(**[2]uint8)(__ccgo_up(bp + 2)))[int32(1)] = libc.Uint8FromInt8(**(**int8)(__ccgo_up(__ccgo_ts + 595 + uintptr(libc.Int32FromUint8(**(**uint8)(__ccgo_up(aBlob + uintptr(iBlob))))&int32(15)))))
							HashUpdate(tls, bp+2, uint32(2))
							goto _2
						_2:
							;
							iBlob = iBlob + 1
						}
						g.FnResByte += libc.Uint64FromInt32(nBlob*int32(2) + int32(2))
					} else {
						HashUpdate(tls, z1, libc.Uint32FromInt32(len1))
						g.FnResByte += libc.Uint64FromInt32(len1 + int32(2))
					}
				}
			}
			if libc.Uint64FromInt32(g.FnResult+len1) < libc.Uint64FromInt64(3000)-libc.Uint64FromInt32(2) {
				if g.FnResult > 0 {
					v4 = uintptr(unsafe.Pointer(&g)) + 144
					v3 = *(*int32)(unsafe.Pointer(v4))
					*(*int32)(unsafe.Pointer(v4)) = *(*int32)(unsafe.Pointer(v4)) + 1
					**(**int8)(__ccgo_up(uintptr(unsafe.Pointer(&g)) + 148 + uintptr(v3))) = int8(' ')
				}
				libc.X__builtin___memcpy_chk(tls, uintptr(unsafe.Pointer(&g))+148+uintptr(g.FnResult), z1, libc.Uint64FromInt32(len1+int32(1)), ^__predefined_size_t(0))
				g.FnResult += len1
			}
			goto _1
		_1:
			;
			i = i + 1
		}
	}
	if g.FbReprepare != 0 {
		libsqlite3.Xsqlite3_prepare_v2(tls, g.Fdb, libsqlite3.Xsqlite3_sql(tls, g.FpStmt), -int32(1), bp+8, uintptr(0))
		rc = libsqlite3.Xsqlite3_finalize(tls, g.FpStmt)
		if rc != SQLITE_OK {
			fatal_error(tls, __ccgo_ts+562, libc.VaList(bp+24, libsqlite3.Xsqlite3_sql(tls, **(**uintptr)(__ccgo_up(bp + 8))), rc, libsqlite3.Xsqlite3_errmsg(tls, g.Fdb)))
		}
		g.FpStmt = **(**uintptr)(__ccgo_up(bp + 8))
	} else {
		rc = libsqlite3.Xsqlite3_reset(tls, g.FpStmt)
		if rc != SQLITE_OK {
			fatal_error(tls, __ccgo_ts+562, libc.VaList(bp+24, libsqlite3.Xsqlite3_sql(tls, g.FpStmt), rc, libsqlite3.Xsqlite3_errmsg(tls, g.Fdb)))
		}
	}
	speedtest1_shrink_memory(tls)
}

// C documentation
//
//	/* The sqlite3_trace() callback function */
func traceCallback(tls *libc.TLS, NotUsed uintptr, zSql uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var n, v1, v5, v7, v8 int32
	var v10, v11 bool
	var v3 __darwin_ct_rune_t
	var v4 uint64
	_, _, _, _, _, _, _, _, _ = n, v1, v10, v11, v3, v4, v5, v7, v8
	n = libc.Int32FromUint64(libc.Xstrlen(tls, zSql))
	for {
		if v11 = n > 0; v11 {
			if v10 = int32(**(**int8)(__ccgo_up(zSql + uintptr(n-int32(1))))) == int32(';'); !v10 {
				v3 = libc.Int32FromUint8(libc.Uint8FromInt8(**(**int8)(__ccgo_up(zSql + uintptr(n-int32(1))))))
				v4 = uint64(0x00004000)
				v8 = libc.BoolInt32(v3 & ^libc.Int32FromInt32(0x7F) == 0)
				goto _9
			_9:
				if v8 != 0 {
					v7 = libc.BoolInt32(!!(uint64(**(**__uint32_t)(__ccgo_up(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v3)*4)))&v4 != 0))
				} else {
					v7 = libc.BoolInt32(!!(libc.X__maskrune(tls, v3, v4) != 0))
				}
				v5 = v7
				goto _6
			_6:
				v1 = v5
				goto _2
			_2:
			}
		}
		if !(v11 && (v10 || v1 != 0)) {
			break
		}
		n = n - 1
	}
	libc.Xfprintf(tls, libc.X__stderrp, __ccgo_ts+473, libc.VaList(bp+8, n, zSql))
}

// C documentation
//
//	/* Substitute random() function that gives the same random
//	** sequence on each run, for repeatability. */
func randomFunc(tls *libc.TLS, context uintptr, NotUsed int32, NotUsed2 uintptr) {
	libsqlite3.Xsqlite3_result_int64(tls, context, libc.Int64FromUint32(speedtest1_random(tls)))
}

// C documentation
//
//	/* Estimate the square root of an integer */
func est_square_root(tls *libc.TLS, x int32) (r int32) {
	var n, y0, y1 int32
	_, _, _ = n, y0, y1
	y0 = x / int32(2)
	n = 0
	for {
		if !(y0 > 0 && n < int32(10)) {
			break
		}
		y1 = (y0 + x/y0) / int32(2)
		if y1 == y0 {
			break
		}
		y0 = y1
		goto _1
	_1:
		;
		n = n + 1
	}
	return y0
}

// C documentation
//
//	/*
//	** The main and default testset
//	*/
func testset_main(tls *libc.TLS) {
	bp := tls.Alloc(2064)
	defer tls.Free(2064)
	var i, len1, maxb, n, sz, v1 int32
	var x1, x2 uint32
	var v21 uintptr
	var _ /* zNum at bp+0 */ [2000]int8
	_, _, _, _, _, _, _, _, _ = i, len1, maxb, n, sz, x1, x2, v1, v21 /* Maximum swizzled value */
	x1 = uint32(0)
	x2 = uint32(0) /* Parameters */
	len1 = 0       /* A number name */
	v1 = g.FszTest * libc.Int32FromInt32(500)
	n = v1
	sz = v1
	(**(**[2000]int8)(__ccgo_up(bp)))[0] = 0
	maxb = libc.Int32FromUint32(roundup_allones(tls, libc.Uint32FromInt32(sz)))
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+612, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+648, 0)
	speedtest1_exec(tls, __ccgo_ts+654, libc.VaList(bp+2008, isTemp(tls, int32(9)), g.FzNN, g.FzNN, g.FzNN))
	speedtest1_prepare(tls, __ccgo_ts+712, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = swizzle(tls, libc.Uint32FromInt32(i), libc.Uint32FromInt32(maxb))
		speedtest1_numbername(tls, x1, bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int64(tls, g.FpStmt, int32(1), libc.Int64FromUint32(x1))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), i)
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(3), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _2
	_2:
		;
		i = i + 1
	}
	speedtest1_exec(tls, __ccgo_ts+758, 0)
	speedtest1_end_test(tls)
	n = sz
	speedtest1_begin_test(tls, int32(110), __ccgo_ts+765, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+648, 0)
	speedtest1_exec(tls, __ccgo_ts+802, libc.VaList(bp+2008, isTemp(tls, int32(5)), g.FzNN, g.FzPK, g.FzNN, g.FzNN, g.FzWR))
	speedtest1_prepare(tls, __ccgo_ts+865, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = swizzle(tls, libc.Uint32FromInt32(i), libc.Uint32FromInt32(maxb))
		speedtest1_numbername(tls, x1, bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), i)
		libsqlite3.Xsqlite3_bind_int64(tls, g.FpStmt, int32(2), libc.Int64FromUint32(x1))
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(3), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _3
	_3:
		;
		i = i + 1
	}
	speedtest1_exec(tls, __ccgo_ts+758, 0)
	speedtest1_end_test(tls)
	n = sz
	speedtest1_begin_test(tls, int32(120), __ccgo_ts+910, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+648, 0)
	speedtest1_exec(tls, __ccgo_ts+949, libc.VaList(bp+2008, isTemp(tls, int32(3)), g.FzNN, g.FzPK, g.FzNN, g.FzNN, g.FzWR))
	speedtest1_prepare(tls, __ccgo_ts+1012, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = swizzle(tls, libc.Uint32FromInt32(i), libc.Uint32FromInt32(maxb))
		speedtest1_numbername(tls, x1, bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), i)
		libsqlite3.Xsqlite3_bind_int64(tls, g.FpStmt, int32(1), libc.Int64FromUint32(x1))
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(3), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _4
	_4:
		;
		i = i + 1
	}
	speedtest1_exec(tls, __ccgo_ts+758, 0)
	speedtest1_end_test(tls)
	n = int32(25)
	speedtest1_begin_test(tls, int32(130), __ccgo_ts+1057, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+648, 0)
	speedtest1_prepare(tls, __ccgo_ts+1096, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		if (i-int32(1))%g.FnRepeat == 0 {
			x1 = speedtest1_random(tls) % libc.Uint32FromInt32(maxb)
			x2 = speedtest1_random(tls)%uint32(10) + libc.Uint32FromInt32(sz/int32(5000)) + x1
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), libc.Int32FromUint32(x1))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), libc.Int32FromUint32(x2))
		speedtest1_run(tls)
		goto _5
	_5:
		;
		i = i + 1
	}
	speedtest1_exec(tls, __ccgo_ts+758, 0)
	speedtest1_end_test(tls)
	n = int32(10)
	speedtest1_begin_test(tls, int32(140), __ccgo_ts+1201, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+648, 0)
	speedtest1_prepare(tls, __ccgo_ts+1229, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		if (i-int32(1))%g.FnRepeat == 0 {
			x1 = speedtest1_random(tls) % libc.Uint32FromInt32(maxb)
			(**(**[2000]int8)(__ccgo_up(bp)))[0] = int8('%')
			len1 = speedtest1_numbername(tls, libc.Uint32FromInt32(i), bp+uintptr(1), libc.Int32FromUint64(libc.Uint64FromInt64(2000)-libc.Uint64FromInt32(2)))
			(**(**[2000]int8)(__ccgo_up(bp)))[len1] = int8('%')
			(**(**[2000]int8)(__ccgo_up(bp)))[len1+int32(1)] = 0
		}
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(1), bp, len1+int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _6
	_6:
		;
		i = i + 1
	}
	speedtest1_exec(tls, __ccgo_ts+758, 0)
	speedtest1_end_test(tls)
	n = int32(10)
	speedtest1_begin_test(tls, int32(142), __ccgo_ts+1324, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+648, 0)
	speedtest1_prepare(tls, __ccgo_ts+1357, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		if (i-int32(1))%g.FnRepeat == 0 {
			x1 = speedtest1_random(tls) % libc.Uint32FromInt32(maxb)
			(**(**[2000]int8)(__ccgo_up(bp)))[0] = int8('%')
			len1 = speedtest1_numbername(tls, libc.Uint32FromInt32(i), bp+uintptr(1), libc.Int32FromUint64(libc.Uint64FromInt64(2000)-libc.Uint64FromInt32(2)))
			(**(**[2000]int8)(__ccgo_up(bp)))[len1] = int8('%')
			(**(**[2000]int8)(__ccgo_up(bp)))[len1+int32(1)] = 0
		}
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(1), bp, len1+int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _7
	_7:
		;
		i = i + 1
	}
	speedtest1_exec(tls, __ccgo_ts+758, 0)
	speedtest1_end_test(tls)
	n = int32(10) /* g.szTest/5; */
	speedtest1_begin_test(tls, int32(145), __ccgo_ts+1421, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+648, 0)
	speedtest1_prepare(tls, __ccgo_ts+1464, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		if (i-int32(1))%g.FnRepeat == 0 {
			x1 = speedtest1_random(tls) % libc.Uint32FromInt32(maxb)
			(**(**[2000]int8)(__ccgo_up(bp)))[0] = int8('%')
			len1 = speedtest1_numbername(tls, libc.Uint32FromInt32(i), bp+uintptr(1), libc.Int32FromUint64(libc.Uint64FromInt64(2000)-libc.Uint64FromInt32(2)))
			(**(**[2000]int8)(__ccgo_up(bp)))[len1] = int8('%')
			(**(**[2000]int8)(__ccgo_up(bp)))[len1+int32(1)] = 0
		}
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(1), bp, len1+int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _8
	_8:
		;
		i = i + 1
	}
	speedtest1_exec(tls, __ccgo_ts+758, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(150), __ccgo_ts+1537, 0)
	speedtest1_exec(tls, __ccgo_ts+1561, 0)
	speedtest1_exec(tls, __ccgo_ts+1568, 0)
	speedtest1_exec(tls, __ccgo_ts+1602, 0)
	speedtest1_exec(tls, __ccgo_ts+1629, 0)
	speedtest1_exec(tls, __ccgo_ts+1663, 0)
	speedtest1_exec(tls, __ccgo_ts+1695, 0)
	speedtest1_exec(tls, __ccgo_ts+1725, 0)
	speedtest1_end_test(tls)
	n = sz / int32(5)
	speedtest1_begin_test(tls, int32(160), __ccgo_ts+1733, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+648, 0)
	speedtest1_prepare(tls, __ccgo_ts+1770, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		if (i-int32(1))%g.FnRepeat == 0 {
			x1 = speedtest1_random(tls) % libc.Uint32FromInt32(maxb)
			x2 = speedtest1_random(tls)%uint32(10) + libc.Uint32FromInt32(sz/int32(5000)) + x1
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), libc.Int32FromUint32(x1))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), libc.Int32FromUint32(x2))
		speedtest1_run(tls)
		goto _9
	_9:
		;
		i = i + 1
	}
	speedtest1_exec(tls, __ccgo_ts+758, 0)
	speedtest1_end_test(tls)
	n = sz / int32(5)
	speedtest1_begin_test(tls, int32(161), __ccgo_ts+1875, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+648, 0)
	speedtest1_prepare(tls, __ccgo_ts+1907, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		if (i-int32(1))%g.FnRepeat == 0 {
			x1 = speedtest1_random(tls) % libc.Uint32FromInt32(maxb)
			x2 = speedtest1_random(tls)%uint32(10) + libc.Uint32FromInt32(sz/int32(5000)) + x1
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), libc.Int32FromUint32(x1))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), libc.Int32FromUint32(x2))
		speedtest1_run(tls)
		goto _10
	_10:
		;
		i = i + 1
	}
	speedtest1_exec(tls, __ccgo_ts+758, 0)
	speedtest1_end_test(tls)
	n = sz / int32(5)
	speedtest1_begin_test(tls, int32(170), __ccgo_ts+2012, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+648, 0)
	speedtest1_prepare(tls, __ccgo_ts+2046, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		if (i-int32(1))%g.FnRepeat == 0 {
			x1 = swizzle(tls, libc.Uint32FromInt32(i), libc.Uint32FromInt32(maxb))
			len1 = speedtest1_numbername(tls, x1, bp, libc.Int32FromUint64(libc.Uint64FromInt64(2000)-libc.Uint64FromInt32(1)))
		}
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(1), bp, len1, libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _11
	_11:
		;
		i = i + 1
	}
	speedtest1_exec(tls, __ccgo_ts+758, 0)
	speedtest1_end_test(tls)
	n = sz
	speedtest1_begin_test(tls, int32(180), __ccgo_ts+2158, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+648, 0)
	speedtest1_exec(tls, __ccgo_ts+2188, libc.VaList(bp+2008, isTemp(tls, int32(1)), g.FzNN, g.FzPK, g.FzNN, g.FzNN, g.FzWR))
	speedtest1_exec(tls, __ccgo_ts+2259, 0)
	speedtest1_exec(tls, __ccgo_ts+2285, 0)
	speedtest1_exec(tls, __ccgo_ts+2311, 0)
	speedtest1_exec(tls, __ccgo_ts+758, 0)
	speedtest1_end_test(tls)
	n = sz
	speedtest1_begin_test(tls, int32(190), __ccgo_ts+2343, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+2371, 0)
	speedtest1_exec(tls, __ccgo_ts+2387, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(200), __ccgo_ts+2420, 0)
	speedtest1_exec(tls, __ccgo_ts+2420, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(210), __ccgo_ts+2427, 0)
	speedtest1_exec(tls, __ccgo_ts+2461, 0)
	speedtest1_exec(tls, __ccgo_ts+2505, 0)
	speedtest1_end_test(tls)
	n = sz / int32(5)
	speedtest1_begin_test(tls, int32(230), __ccgo_ts+2527, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+648, 0)
	speedtest1_prepare(tls, __ccgo_ts+2564, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = speedtest1_random(tls) % libc.Uint32FromInt32(maxb)
		x2 = speedtest1_random(tls)%uint32(10) + libc.Uint32FromInt32(sz/int32(5000)) + x1
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), libc.Int32FromUint32(x1))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), libc.Int32FromUint32(x2))
		speedtest1_run(tls)
		goto _12
	_12:
		;
		i = i + 1
	}
	speedtest1_exec(tls, __ccgo_ts+758, 0)
	speedtest1_end_test(tls)
	n = sz
	speedtest1_begin_test(tls, int32(240), __ccgo_ts+2623, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+648, 0)
	speedtest1_prepare(tls, __ccgo_ts+2653, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = speedtest1_random(tls)%libc.Uint32FromInt32(sz) + uint32(1)
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), libc.Int32FromUint32(x1))
		speedtest1_run(tls)
		goto _13
	_13:
		;
		i = i + 1
	}
	speedtest1_exec(tls, __ccgo_ts+758, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(250), __ccgo_ts+2697, libc.VaList(bp+2008, sz))
	speedtest1_exec(tls, __ccgo_ts+2738, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(260), __ccgo_ts+2758, 0)
	speedtest1_exec(tls, __ccgo_ts+2505, 0)
	speedtest1_end_test(tls)
	n = sz / int32(5)
	speedtest1_begin_test(tls, int32(270), __ccgo_ts+2791, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+648, 0)
	speedtest1_prepare(tls, __ccgo_ts+2828, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = speedtest1_random(tls)%libc.Uint32FromInt32(maxb) + uint32(1)
		x2 = speedtest1_random(tls)%uint32(10) + libc.Uint32FromInt32(sz/int32(5000)) + x1
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), libc.Int32FromUint32(x1))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), libc.Int32FromUint32(x2))
		speedtest1_run(tls)
		goto _14
	_14:
		;
		i = i + 1
	}
	speedtest1_exec(tls, __ccgo_ts+758, 0)
	speedtest1_end_test(tls)
	n = sz
	speedtest1_begin_test(tls, int32(280), __ccgo_ts+2882, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+648, 0)
	speedtest1_prepare(tls, __ccgo_ts+2912, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = speedtest1_random(tls)%libc.Uint32FromInt32(sz) + uint32(1)
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), libc.Int32FromUint32(x1))
		speedtest1_run(tls)
		goto _15
	_15:
		;
		i = i + 1
	}
	speedtest1_exec(tls, __ccgo_ts+758, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(290), __ccgo_ts+2951, libc.VaList(bp+2008, sz))
	speedtest1_exec(tls, __ccgo_ts+2990, 0)
	speedtest1_exec(tls, __ccgo_ts+3034, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(300), __ccgo_ts+3078, libc.VaList(bp+2008, sz))
	speedtest1_exec(tls, __ccgo_ts+2371, 0)
	speedtest1_exec(tls, __ccgo_ts+3119, 0)
	speedtest1_exec(tls, __ccgo_ts+3184, 0)
	speedtest1_end_test(tls)
	n = sz / int32(5)
	speedtest1_begin_test(tls, int32(310), __ccgo_ts+3249, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+648, 0)
	speedtest1_prepare(tls, __ccgo_ts+3268, 0)
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = speedtest1_random(tls)%libc.Uint32FromInt32(sz) + uint32(1)
		x2 = speedtest1_random(tls)%uint32(10) + x1 + uint32(4)
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), libc.Int32FromUint32(x1))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), libc.Int32FromUint32(x2))
		speedtest1_run(tls)
		goto _16
	_16:
		;
		i = i + 1
	}
	speedtest1_exec(tls, __ccgo_ts+758, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(320), __ccgo_ts+3382, libc.VaList(bp+2008, n))
	speedtest1_prepare(tls, __ccgo_ts+3405, 0)
	libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), est_square_root(tls, g.FszTest)*int32(50))
	speedtest1_run(tls)
	speedtest1_end_test(tls)
	v1 = g.FszTest * libc.Int32FromInt32(700)
	n = v1
	sz = v1
	(**(**[2000]int8)(__ccgo_up(bp)))[0] = 0
	maxb = libc.Int32FromUint32(roundup_allones(tls, libc.Uint32FromInt32(sz/int32(3))))
	speedtest1_begin_test(tls, int32(400), __ccgo_ts+3523, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+648, 0)
	speedtest1_exec(tls, __ccgo_ts+3548, libc.VaList(bp+2008, isTemp(tls, int32(9)), g.FzNN))
	speedtest1_prepare(tls, __ccgo_ts+3596, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = swizzle(tls, libc.Uint32FromInt32(i), libc.Uint32FromInt32(maxb))
		speedtest1_numbername(tls, libc.Uint32FromInt32(i), bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), int32(libc.Int64FromUint32(x1)))
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(2), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _18
	_18:
		;
		i = i + 1
	}
	speedtest1_exec(tls, __ccgo_ts+758, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(410), __ccgo_ts+3640, libc.VaList(bp+2008, n))
	if g.FdoBigTransactions != 0 {
		/* Historical note: tests 410 and 510 have historically not used
		 ** explicit transactions. The --big-transactions flag was added
		 ** 2022-09-08 to support the WASM/OPFS build, as the run-times
		 ** approach 1 minute for each of these tests if they're not in an
		 ** explicit transaction. The run-time effect of --big-transaciions
		 ** on native builds is negligible. */
		speedtest1_exec(tls, __ccgo_ts+648, 0)
	}
	speedtest1_prepare(tls, __ccgo_ts+3661, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = swizzle(tls, libc.Uint32FromInt32(i), libc.Uint32FromInt32(maxb))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), int32(libc.Int64FromUint32(x1)))
		speedtest1_run(tls)
		goto _19
	_19:
		;
		i = i + 1
	}
	if g.FdoBigTransactions != 0 {
		speedtest1_exec(tls, __ccgo_ts+758, 0)
	}
	speedtest1_end_test(tls)
	v1 = g.FszTest * libc.Int32FromInt32(700)
	n = v1
	sz = v1
	(**(**[2000]int8)(__ccgo_up(bp)))[0] = 0
	maxb = libc.Int32FromUint32(roundup_allones(tls, libc.Uint32FromInt32(sz/int32(3))))
	speedtest1_begin_test(tls, int32(500), __ccgo_ts+3703, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+648, 0)
	if libsqlite3.Xsqlite3_libversion_number(tls) >= int32(3008002) {
		v21 = __ccgo_ts + 3725
	} else {
		v21 = __ccgo_ts + 6
	}
	speedtest1_exec(tls, __ccgo_ts+3739, libc.VaList(bp+2008, isTemp(tls, int32(9)), g.FzNN, v21))
	speedtest1_prepare(tls, __ccgo_ts+3786, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = swizzle(tls, libc.Uint32FromInt32(i), libc.Uint32FromInt32(maxb))
		speedtest1_numbername(tls, x1, bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), i)
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(1), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _22
	_22:
		;
		i = i + 1
	}
	speedtest1_exec(tls, __ccgo_ts+758, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(510), __ccgo_ts+3830, libc.VaList(bp+2008, n))
	if g.FdoBigTransactions != 0 {
		/* See notes for test 410. */
		speedtest1_exec(tls, __ccgo_ts+648, 0)
	}
	speedtest1_prepare(tls, __ccgo_ts+3854, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = swizzle(tls, libc.Uint32FromInt32(i), libc.Uint32FromInt32(maxb))
		speedtest1_numbername(tls, x1, bp, int32(2000))
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(1), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _23
	_23:
		;
		i = i + 1
	}
	if g.FdoBigTransactions != 0 {
		speedtest1_exec(tls, __ccgo_ts+758, 0)
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(520), __ccgo_ts+3896, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+3915, 0)
	speedtest1_exec(tls, __ccgo_ts+3942, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(980), __ccgo_ts+3969, 0)
	speedtest1_exec(tls, __ccgo_ts+3969, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(990), __ccgo_ts+3992, 0)
	speedtest1_exec(tls, __ccgo_ts+3992, 0)
	speedtest1_end_test(tls)
}

// C documentation
//
//	/*
//	** A testset for common table expressions.  This exercises code
//	** for views, subqueries, co-routines, etc.
//	*/
func testset_cte(tls *libc.TLS) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var nElem int32
	var rSpacing float64
	var zPuz uintptr
	_, _, _ = nElem, rSpacing, zPuz
	if g.FszTest < int32(25) {
		zPuz = azPuzzle[0]
	} else {
		if g.FszTest < int32(70) {
			zPuz = azPuzzle[int32(1)]
		} else {
			zPuz = azPuzzle[int32(2)]
		}
	}
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+4246, 0)
	speedtest1_prepare(tls, __ccgo_ts+4277, 0)
	libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(1), zPuz, -int32(1), libc.UintptrFromInt32(0))
	speedtest1_run(tls)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(200), __ccgo_ts+5081, 0)
	speedtest1_prepare(tls, __ccgo_ts+5109, 0)
	libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(1), zPuz, -int32(1), libc.UintptrFromInt32(0))
	speedtest1_run(tls)
	speedtest1_end_test(tls)
	rSpacing = float64(5) / float64(g.FszTest)
	speedtest1_begin_test(tls, int32(300), __ccgo_ts+5921, libc.VaList(bp+8, rSpacing))
	speedtest1_prepare(tls, __ccgo_ts+5952, 0)
	libsqlite3.Xsqlite3_bind_double(tls, g.FpStmt, int32(1), float64(rSpacing*float64(0.05)))
	libsqlite3.Xsqlite3_bind_double(tls, g.FpStmt, int32(2), rSpacing)
	speedtest1_run(tls)
	speedtest1_end_test(tls)
	nElem = int32(10000) * g.FszTest
	speedtest1_begin_test(tls, int32(400), __ccgo_ts+6548, libc.VaList(bp+8, nElem))
	speedtest1_prepare(tls, __ccgo_ts+6585, libc.VaList(bp+8, nElem, nElem))
	speedtest1_run(tls)
	speedtest1_end_test(tls)
}

var azPuzzle = [3]uintptr{
	0: __ccgo_ts + 4000,
	1: __ccgo_ts + 4082,
	2: __ccgo_ts + 4164,
}

// C documentation
//
//	/*
//	** Compute a pseudo-random floating point ascii number.
//	*/
func speedtest1_random_ascii_fp(tls *libc.TLS, zFP uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var x, y, z int32
	_, _, _ = x, y, z
	x = libc.Int32FromUint32(speedtest1_random(tls))
	y = libc.Int32FromUint32(speedtest1_random(tls))
	z = y % int32(10)
	if z < 0 {
		z = -z
	}
	y = y / int32(10)
	libsqlite3.Xsqlite3_snprintf(tls, int32(100), zFP, __ccgo_ts+6816, libc.VaList(bp+8, y, z, x%int32(200)))
}

// C documentation
//
//	/*
//	** A testset for floating-point numbers.
//	*/
func testset_fp(tls *libc.TLS) {
	bp := tls.Alloc(240)
	defer tls.Free(240)
	var i, n int32
	var _ /* zFP1 at bp+0 */ [100]int8
	var _ /* zFP2 at bp+100 */ [100]int8
	_, _ = i, n
	n = g.FszTest * int32(5000)
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+6825, libc.VaList(bp+208, n*int32(2)))
	speedtest1_exec(tls, __ccgo_ts+648, 0)
	speedtest1_exec(tls, __ccgo_ts+6856, libc.VaList(bp+208, isTemp(tls, int32(1)), g.FzNN, g.FzNN))
	speedtest1_prepare(tls, __ccgo_ts+6897, libc.VaList(bp+208, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		speedtest1_random_ascii_fp(tls, bp)
		speedtest1_random_ascii_fp(tls, bp+100)
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(1), bp, -int32(1), libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(2), bp+100, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _1
	_1:
		;
		i = i + 1
	}
	speedtest1_exec(tls, __ccgo_ts+758, 0)
	speedtest1_end_test(tls)
	n = g.FszTest/int32(25) + int32(2)
	speedtest1_begin_test(tls, int32(110), __ccgo_ts+6939, libc.VaList(bp+208, n))
	speedtest1_prepare(tls, __ccgo_ts+6956, 0)
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		speedtest1_random_ascii_fp(tls, bp)
		speedtest1_random_ascii_fp(tls, bp+100)
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(1), bp, -int32(1), libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(2), bp+100, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _2
	_2:
		;
		i = i + 1
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(120), __ccgo_ts+7004, 0)
	speedtest1_exec(tls, __ccgo_ts+1561, 0)
	speedtest1_exec(tls, __ccgo_ts+7029, 0)
	speedtest1_exec(tls, __ccgo_ts+7056, 0)
	speedtest1_exec(tls, __ccgo_ts+7083, 0)
	speedtest1_exec(tls, __ccgo_ts+1725, 0)
	speedtest1_end_test(tls)
	n = g.FszTest/int32(3) + int32(2)
	speedtest1_begin_test(tls, int32(130), __ccgo_ts+7113, libc.VaList(bp+208, n))
	speedtest1_prepare(tls, __ccgo_ts+6956, 0)
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		speedtest1_random_ascii_fp(tls, bp)
		speedtest1_random_ascii_fp(tls, bp+100)
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(1), bp, -int32(1), libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(2), bp+100, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _3
	_3:
		;
		i = i + 1
	}
	speedtest1_end_test(tls)
	n = g.FszTest * int32(5000)
	speedtest1_begin_test(tls, int32(140), __ccgo_ts+7138, libc.VaList(bp+208, n))
	speedtest1_exec(tls, __ccgo_ts+7158, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(150), __ccgo_ts+7201, libc.VaList(bp+208, n*int32(4)))
	speedtest1_exec(tls, __ccgo_ts+7219, 0)
	speedtest1_end_test(tls)
}

// C documentation
//
//	/*
//	** A testset for star-schema queries.
//	*/
func testset_star(tls *libc.TLS) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var i, n int32
	_, _ = i, n
	n = g.FszTest * int32(50)
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+7314, libc.VaList(bp+8, n))
	speedtest1_exec(tls, __ccgo_ts+7350, 0)
	speedtest1_exec(tls, __ccgo_ts+7523, libc.VaList(bp+8, n))
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(110), __ccgo_ts+7878, 0)
	i = int32(1)
	for {
		if !(i <= int32(8)) {
			break
		}
		speedtest1_exec(tls, __ccgo_ts+7919, libc.VaList(bp+8, i, i))
		goto _1
	_1:
		;
		i = i + 1
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(120), __ccgo_ts+7967, 0)
	i = int32(1)
	for {
		if !(i <= int32(8)) {
			break
		}
		speedtest1_exec(tls, __ccgo_ts+7991, libc.VaList(bp+8, i, i, i, i))
		speedtest1_exec(tls, __ccgo_ts+8065, libc.VaList(bp+8, int32(4)*(i+int32(1)), i, int32(2)*(i+int32(1)), i))
		if i&int32(2) != 0 {
			speedtest1_exec(tls, __ccgo_ts+8253, libc.VaList(bp+8, i, i, i))
		} else {
			speedtest1_exec(tls, __ccgo_ts+8302, libc.VaList(bp+8, i, i, i, i))
		}
		goto _2
	_2:
		;
		i = i + 1
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(130), __ccgo_ts+8363, 0)
	speedtest1_exec(tls, __ccgo_ts+8401, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(130), __ccgo_ts+8739, 0)
	speedtest1_exec(tls, __ccgo_ts+8766, 0)
	speedtest1_end_test(tls)
}

// C documentation
//
//	/*
//	** Tests that simulate an application opening and closing an SQLite database
//	** frequently.  Fossil is used as the model.  The focus here is on rapidly
//	** parsing the database schema and rapidly generating prepared statements,
//	** in other words, rapid start-up of Fossil-like applications.
//	**
//	** The same database has no data, so the performance of sqlite3_step() is
//	** not significant to this testset.
//	*/
func testset_app(tls *libc.TLS) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var dbMain uintptr
	var i, n int32
	var _ /* dbAux at bp+0 */ uintptr
	_, _, _ = dbMain, i, n
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+9178, 0)
	speedtest1_exec(tls, __ccgo_ts+9217, 0)
	if libsqlite3.Xsqlite3_compileoption_used(tls, __ccgo_ts+15766) != 0 {
		speedtest1_exec(tls, __ccgo_ts+15778, 0)
	} else {
		speedtest1_exec(tls, __ccgo_ts+15958, 0)
	}
	speedtest1_exec(tls, __ccgo_ts+16507, 0)
	speedtest1_end_test(tls)
	n = g.FszTest * int32(3)
	speedtest1_begin_test(tls, int32(110), __ccgo_ts+18673, libc.VaList(bp+16, n))
	i = 0
	for {
		if !(i < n) {
			break
		}
		dbMain = g.Fdb
		**(**uintptr)(__ccgo_up(bp)) = uintptr(0)
		if g.FzDbName != 0 && **(**int8)(__ccgo_up(g.FzDbName)) != 0 {
			if libsqlite3.Xsqlite3_open_v2(tls, g.FzDbName, bp, int32(SQLITE_OPEN_READWRITE), g.FzVfs) != 0 {
				fatal_error(tls, __ccgo_ts+18708, libc.VaList(bp+16, g.FzDbName))
			}
			g.Fdb = **(**uintptr)(__ccgo_up(bp))
		}
		speedtest1_exec(tls, __ccgo_ts+18739, 0)
		speedtest1_exec(tls, __ccgo_ts+19180, 0)
		speedtest1_exec(tls, __ccgo_ts+19383, 0)
		speedtest1_exec(tls, __ccgo_ts+19711, 0)
		speedtest1_exec(tls, __ccgo_ts+19773, 0)
		libsqlite3.Xsqlite3_close(tls, **(**uintptr)(__ccgo_up(bp)))
		g.Fdb = dbMain
		goto _1
	_1:
		;
		i = i + 1
	}
	speedtest1_end_test(tls)
}

// C documentation
//
//	/* Generate two numbers between 1 and mx.  The first number is less than
//	** the second.  Usually the numbers are near each other but can sometimes
//	** be far apart.
//	*/
func twoCoords(tls *libc.TLS, p1 int32, p2 int32, mx uint32, pX0 uintptr, pX1 uintptr) {
	var d, span, x0, x1 uint32
	_, _, _, _ = d, span, x0, x1
	span = mx/uint32(100) + uint32(1)
	if speedtest1_random(tls)%uint32(3) == uint32(0) {
		span = span * libc.Uint32FromInt32(p1)
	}
	if speedtest1_random(tls)%libc.Uint32FromInt32(p2) == uint32(0) {
		span = mx / uint32(2)
	}
	d = speedtest1_random(tls)%span + uint32(1)
	x0 = speedtest1_random(tls)%(mx-d) + uint32(1)
	x1 = x0 + d
	**(**uint32)(__ccgo_up(pX0)) = x0
	**(**uint32)(__ccgo_up(pX1)) = x1
}

// C documentation
//
//	/* The following routine is an R-Tree geometry callback.  It returns
//	** true if the object overlaps a slice on the Y coordinate between the
//	** two values given as arguments.  In other words
//	**
//	**     SELECT count(*) FROM rt1 WHERE id MATCH xslice(10,20);
//	**
//	** Is the same as saying:
//	**
//	**     SELECT count(*) FROM rt1 WHERE y1>=10 AND y0<=20;
//	*/
func xsliceGeometryCallback(tls *libc.TLS, p uintptr, nCoord int32, aCoord uintptr, pRes uintptr) (r int32) {
	**(**int32)(__ccgo_up(pRes)) = libc.BoolInt32(**(**float64)(__ccgo_up(aCoord + 3*8)) >= **(**sqlite3_rtree_dbl)(__ccgo_up((*sqlite3_rtree_geometry)(unsafe.Pointer(p)).FaParam)) && **(**float64)(__ccgo_up(aCoord + 2*8)) <= **(**sqlite3_rtree_dbl)(__ccgo_up((*sqlite3_rtree_geometry)(unsafe.Pointer(p)).FaParam + 1*8)))
	return SQLITE_OK
}

// C documentation
//
//	/*
//	** A testset for the R-Tree virtual table
//	*/
func testset_rtree(tls *libc.TLS, p1 int32, p2 int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var aCheck uintptr
	var i, iStep, mxCoord, mxRowid, n, v1 uint32
	var _ /* x0 at bp+0 */ uint32
	var _ /* x1 at bp+4 */ uint32
	var _ /* y0 at bp+8 */ uint32
	var _ /* y1 at bp+12 */ uint32
	var _ /* z0 at bp+16 */ uint32
	var _ /* z1 at bp+20 */ uint32
	_, _, _, _, _, _, _ = aCheck, i, iStep, mxCoord, mxRowid, n, v1
	aCheck = libsqlite3.Xsqlite3_malloc(tls, libc.Int32FromUint64(uint64(4)*libc.Uint64FromInt32(g.FszTest)*uint64(500)))
	mxCoord = uint32(15000)
	v1 = libc.Uint32FromInt32(g.FszTest * libc.Int32FromInt32(500))
	n = v1
	mxRowid = v1
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+20381, libc.VaList(bp+32, n))
	speedtest1_exec(tls, __ccgo_ts+648, 0)
	speedtest1_exec(tls, __ccgo_ts+20407, 0)
	speedtest1_prepare(tls, __ccgo_ts+20466, 0)
	i = uint32(1)
	for {
		if !(i <= n) {
			break
		}
		twoCoords(tls, p1, p2, mxCoord, bp, bp+4)
		twoCoords(tls, p1, p2, mxCoord, bp+8, bp+12)
		twoCoords(tls, p1, p2, mxCoord, bp+16, bp+20)
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), libc.Int32FromUint32(i))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), libc.Int32FromUint32(**(**uint32)(__ccgo_up(bp))))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(3), libc.Int32FromUint32(**(**uint32)(__ccgo_up(bp + 4))))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(4), libc.Int32FromUint32(**(**uint32)(__ccgo_up(bp + 8))))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(5), libc.Int32FromUint32(**(**uint32)(__ccgo_up(bp + 12))))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(6), libc.Int32FromUint32(**(**uint32)(__ccgo_up(bp + 16))))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(7), libc.Int32FromUint32(**(**uint32)(__ccgo_up(bp + 20))))
		speedtest1_run(tls)
		goto _2
	_2:
		;
		i = i + 1
	}
	speedtest1_exec(tls, __ccgo_ts+758, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(101), __ccgo_ts+20532, 0)
	speedtest1_exec(tls, __ccgo_ts+20567, 0)
	speedtest1_exec(tls, __ccgo_ts+20625, 0)
	speedtest1_end_test(tls)
	n = libc.Uint32FromInt32(g.FszTest * int32(200))
	speedtest1_begin_test(tls, int32(110), __ccgo_ts+20658, libc.VaList(bp+32, n))
	speedtest1_prepare(tls, __ccgo_ts+20701, 0)
	iStep = mxCoord / n
	i = uint32(0)
	for {
		if !(i < n) {
			break
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), libc.Int32FromUint32(i*iStep))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), libc.Int32FromUint32((i+uint32(1))*iStep))
		speedtest1_run(tls)
		**(**int32)(__ccgo_up(aCheck + uintptr(i)*4)) = libc.Xatoi(tls, uintptr(unsafe.Pointer(&g))+148)
		goto _3
	_3:
		;
		i = i + 1
	}
	speedtest1_end_test(tls)
	if g.FbVerify != 0 {
		n = libc.Uint32FromInt32(g.FszTest * int32(200))
		speedtest1_begin_test(tls, int32(111), __ccgo_ts+20750, 0)
		speedtest1_prepare(tls, __ccgo_ts+20797, 0)
		iStep = mxCoord / n
		i = uint32(0)
		for {
			if !(i < n) {
				break
			}
			libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), libc.Int32FromUint32(i*iStep))
			libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), libc.Int32FromUint32((i+uint32(1))*iStep))
			speedtest1_run(tls)
			if **(**int32)(__ccgo_up(aCheck + uintptr(i)*4)) != libc.Xatoi(tls, uintptr(unsafe.Pointer(&g))+148) {
				fatal_error(tls, __ccgo_ts+20845, libc.VaList(bp+32, i, i*iStep, (i+uint32(1))*iStep, **(**int32)(__ccgo_up(aCheck + uintptr(i)*4)), libc.Xatoi(tls, uintptr(unsafe.Pointer(&g))+148)))
			}
			goto _4
		_4:
			;
			i = i + 1
		}
		speedtest1_end_test(tls)
	}
	n = libc.Uint32FromInt32(g.FszTest * int32(200))
	speedtest1_begin_test(tls, int32(120), __ccgo_ts+20887, libc.VaList(bp+32, n))
	speedtest1_prepare(tls, __ccgo_ts+20928, 0)
	iStep = mxCoord / n
	i = uint32(0)
	for {
		if !(i < n) {
			break
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), libc.Int32FromUint32(i*iStep))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), libc.Int32FromUint32((i+uint32(1))*iStep))
		speedtest1_run(tls)
		**(**int32)(__ccgo_up(aCheck + uintptr(i)*4)) = libc.Xatoi(tls, uintptr(unsafe.Pointer(&g))+148)
		goto _5
	_5:
		;
		i = i + 1
	}
	speedtest1_end_test(tls)
	if g.FbVerify != 0 {
		n = libc.Uint32FromInt32(g.FszTest * int32(200))
		speedtest1_begin_test(tls, int32(121), __ccgo_ts+20977, 0)
		speedtest1_prepare(tls, __ccgo_ts+21022, 0)
		iStep = mxCoord / n
		i = uint32(0)
		for {
			if !(i < n) {
				break
			}
			libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), libc.Int32FromUint32(i*iStep))
			libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), libc.Int32FromUint32((i+uint32(1))*iStep))
			speedtest1_run(tls)
			if **(**int32)(__ccgo_up(aCheck + uintptr(i)*4)) != libc.Xatoi(tls, uintptr(unsafe.Pointer(&g))+148) {
				fatal_error(tls, __ccgo_ts+20845, libc.VaList(bp+32, i, i*iStep, (i+uint32(1))*iStep, **(**int32)(__ccgo_up(aCheck + uintptr(i)*4)), libc.Xatoi(tls, uintptr(unsafe.Pointer(&g))+148)))
			}
			goto _6
		_6:
			;
			i = i + 1
		}
		speedtest1_end_test(tls)
	}
	n = libc.Uint32FromInt32(g.FszTest * int32(200))
	speedtest1_begin_test(tls, int32(125), __ccgo_ts+21070, libc.VaList(bp+32, n))
	libsqlite3.Xsqlite3_rtree_geometry_callback(tls, g.Fdb, __ccgo_ts+21106, __ccgo_fp(xsliceGeometryCallback), uintptr(0))
	speedtest1_prepare(tls, __ccgo_ts+21113, 0)
	iStep = mxCoord / n
	i = uint32(0)
	for {
		if !(i < n) {
			break
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), libc.Int32FromUint32(i*iStep))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), libc.Int32FromUint32((i+uint32(1))*iStep))
		speedtest1_run(tls)
		if **(**int32)(__ccgo_up(aCheck + uintptr(i)*4)) != libc.Xatoi(tls, uintptr(unsafe.Pointer(&g))+148) {
			fatal_error(tls, __ccgo_ts+20845, libc.VaList(bp+32, i, i*iStep, (i+uint32(1))*iStep, **(**int32)(__ccgo_up(aCheck + uintptr(i)*4)), libc.Xatoi(tls, uintptr(unsafe.Pointer(&g))+148)))
		}
		goto _7
	_7:
		;
		i = i + 1
	}
	speedtest1_end_test(tls)
	n = libc.Uint32FromInt32(g.FszTest * int32(400))
	speedtest1_begin_test(tls, int32(130), __ccgo_ts+21167, libc.VaList(bp+32, n))
	speedtest1_prepare(tls, __ccgo_ts+21210, 0)
	iStep = mxCoord / n
	i = uint32(0)
	for {
		if !(i < n) {
			break
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), libc.Int32FromUint32(i*iStep))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), libc.Int32FromUint32((i+uint32(1))*iStep))
		speedtest1_run(tls)
		**(**int32)(__ccgo_up(aCheck + uintptr(i)*4)) = libc.Xatoi(tls, uintptr(unsafe.Pointer(&g))+148)
		goto _8
	_8:
		;
		i = i + 1
	}
	speedtest1_end_test(tls)
	n = libc.Uint32FromInt32(g.FszTest * int32(500))
	speedtest1_begin_test(tls, int32(140), __ccgo_ts+21303, libc.VaList(bp+32, n))
	speedtest1_prepare(tls, __ccgo_ts+21320, 0)
	i = uint32(1)
	for {
		if !(i <= n) {
			break
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), libc.Int32FromUint32(i))
		speedtest1_run(tls)
		goto _9
	_9:
		;
		i = i + 1
	}
	speedtest1_end_test(tls)
	n = libc.Uint32FromInt32(g.FszTest * int32(50))
	speedtest1_begin_test(tls, int32(150), __ccgo_ts+21350, libc.VaList(bp+32, n))
	speedtest1_prepare(tls, __ccgo_ts+21373, 0)
	i = uint32(1)
	for {
		if !(i <= n) {
			break
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), libc.Int32FromUint32(i*uint32(251)%mxRowid+uint32(1)))
		speedtest1_run(tls)
		goto _10
	_10:
		;
		i = i + 1
	}
	speedtest1_end_test(tls)
	n = libc.Uint32FromInt32(g.FszTest * int32(5))
	speedtest1_begin_test(tls, int32(155), __ccgo_ts+21421, libc.VaList(bp+32, n))
	speedtest1_prepare(tls, __ccgo_ts+21462, 0)
	iStep = mxCoord / n
	i = uint32(0)
	for {
		if !(i < n) {
			break
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), libc.Int32FromUint32(i*iStep))
		speedtest1_run(tls)
		**(**int32)(__ccgo_up(aCheck + uintptr(i)*4)) = libc.Xatoi(tls, uintptr(unsafe.Pointer(&g))+148)
		goto _11
	_11:
		;
		i = i + 1
	}
	speedtest1_end_test(tls)
	n = libc.Uint32FromInt32(g.FszTest * int32(50))
	speedtest1_begin_test(tls, int32(160), __ccgo_ts+21524, libc.VaList(bp+32, n))
	speedtest1_prepare(tls, __ccgo_ts+21547, 0)
	i = uint32(1)
	for {
		if !(i <= n) {
			break
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), libc.Int32FromUint32(i*uint32(257)%mxRowid+uint32(1)))
		speedtest1_run(tls)
		goto _12
	_12:
		;
		i = i + 1
	}
	speedtest1_end_test(tls)
	n = libc.Uint32FromInt32(g.FszTest * int32(5))
	speedtest1_begin_test(tls, int32(165), __ccgo_ts+21575, libc.VaList(bp+32, n))
	speedtest1_prepare(tls, __ccgo_ts+21616, 0)
	iStep = mxCoord / n
	i = uint32(0)
	for {
		if !(i < n) {
			break
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), libc.Int32FromUint32(i*iStep))
		speedtest1_run(tls)
		**(**int32)(__ccgo_up(aCheck + uintptr(i)*4)) = libc.Xatoi(tls, uintptr(unsafe.Pointer(&g))+148)
		goto _13
	_13:
		;
		i = i + 1
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(170), __ccgo_ts+21658, 0)
	speedtest1_exec(tls, __ccgo_ts+21705, 0)
	speedtest1_end_test(tls)
}

// C documentation
//
//	/*
//	** A testset that does key/value storage on tables with many columns.
//	** This is the kind of workload generated by ORMs such as CoreData.
//	*/
func testset_orm(tls *libc.TLS) {
	bp := tls.Alloc(2016)
	defer tls.Free(2016)
	var i, j, len1, n, nRow, x1, v1 uint32
	var _ /* zNum at bp+0 */ [2000]int8
	_, _, _, _, _, _, _ = i, j, len1, n, nRow, x1, v1 /* A number name */
	v1 = libc.Uint32FromInt32(g.FszTest * libc.Int32FromInt32(250))
	n = v1
	nRow = v1
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+21748, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+21761, 0)
	speedtest1_prepare(tls, __ccgo_ts+25661, 0)
	i = uint32(0)
	for {
		if !(i < n) {
			break
		}
		x1 = speedtest1_random(tls)
		speedtest1_numbername(tls, x1%uint32(1000), bp, int32(2000))
		len1 = libc.Uint32FromInt32(libc.Int32FromUint64(libc.Xstrlen(tls, bp)))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), libc.Int32FromUint32(i^uint32(0xf)))
		j = uint32(0)
		for {
			if !(zType[j] != 0) {
				break
			}
			switch int32(zType[j]) {
			case int32('I'):
				fallthrough
			case int32('T'):
				libsqlite3.Xsqlite3_bind_int64(tls, g.FpStmt, libc.Int32FromUint32(j+uint32(2)), libc.Int64FromUint32(x1))
			case int32('F'):
				libsqlite3.Xsqlite3_bind_double(tls, g.FpStmt, libc.Int32FromUint32(j+uint32(2)), float64(x1))
			case int32('V'):
				fallthrough
			case int32('B'):
				libsqlite3.Xsqlite3_bind_text64(tls, g.FpStmt, libc.Int32FromUint32(j+uint32(2)), bp, uint64(len1), libc.UintptrFromInt32(0), uint8(SQLITE_UTF8))
				break
			}
			goto _3
		_3:
			;
			j = j + 1
		}
		speedtest1_run(tls)
		goto _2
	_2:
		;
		i = i + 1
	}
	speedtest1_exec(tls, __ccgo_ts+1725, 0)
	speedtest1_end_test(tls)
	n = libc.Uint32FromInt32(g.FszTest * int32(250))
	speedtest1_begin_test(tls, int32(110), __ccgo_ts+28905, libc.VaList(bp+2008, n))
	speedtest1_prepare(tls, __ccgo_ts+28928, 0)
	i = uint32(0)
	for {
		if !(i < n) {
			break
		}
		x1 = speedtest1_random(tls) % nRow
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), libc.Int32FromUint32(x1))
		speedtest1_run(tls)
		goto _4
	_4:
		;
		i = i + 1
	}
	speedtest1_end_test(tls)
}

var zType = [120]int8{'I', 'B', 'B', 'I', 'I', 'I', 'T', 'I', 'V', 'V', 'I', 'T', 'B', 'T', 'B', 'F', 'B', 'F', 'I', 'T', 'T', 'F', 'B', 'T', 'B', 'V', 'B', 'V', 'I', 'F', 'T', 'B', 'B', 'F', 'I', 'T', 'F', 'F', 'V', 'B', 'I', 'F', 'I', 'V', 'B', 'V', 'V', 'V', 'B', 'T', 'V', 'T', 'I', 'B', 'B', 'F', 'F', 'I', 'V', 'I', 'B', 'T', 'B', 'T', 'V', 'T', 'T', 'F', 'T', 'V', 'T', 'V', 'F', 'F', 'I', 'I', 'T', 'I', 'F', 'B', 'I', 'T', 'F', 'T', 'T', 'F', 'F', 'F', 'V', 'B', 'I', 'I', 'B', 'T', 'T', 'I', 'T', 'F', 'T', 'F', 'F', 'V', 'V', 'V', 'F', 'I', 'I', 'I', 'T', 'V', 'B', 'B', 'V', 'F', 'F', 'T', 'V', 'V', 'B'}

// C documentation
//
//	/*
//	*/
func testset_trigger(tls *libc.TLS) {
	bp := tls.Alloc(2016)
	defer tls.Free(2016)
	var NROW, NROW2, ii, jj, x1 int32
	var _ /* zNum at bp+0 */ [2000]int8
	_, _, _, _, _ = NROW, NROW2, ii, jj, x1 /* A number name */
	NROW = int32(500) * g.FszTest
	NROW2 = int32(100) * g.FszTest
	speedtest1_exec(tls, __ccgo_ts+31681, 0)
	jj = int32(1)
	for {
		if !(jj <= int32(3)) {
			break
		}
		speedtest1_prepare(tls, __ccgo_ts+32009, libc.VaList(bp+2008, jj))
		ii = 0
		for {
			if !(ii < NROW) {
				break
			}
			x1 = libc.Int32FromUint32(speedtest1_random(tls) % libc.Uint32FromInt32(NROW))
			speedtest1_numbername(tls, libc.Uint32FromInt32(x1), bp, int32(2000))
			libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), x1)
			libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(2), bp, -int32(1), libc.UintptrFromInt32(0))
			speedtest1_run(tls)
			goto _2
		_2:
			;
			ii = ii + 1
		}
		goto _1
	_1:
		;
		jj = jj + 1
	}
	speedtest1_exec(tls, __ccgo_ts+32044, 0)
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+32127, 0)
	speedtest1_prepare(tls, __ccgo_ts+32141, 0)
	speedtest1_run(tls)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(110), __ccgo_ts+32208, 0)
	speedtest1_prepare(tls, __ccgo_ts+32222, 0)
	speedtest1_run(tls)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(120), __ccgo_ts+32281, 0)
	jj = int32(1)
	for {
		if !(jj <= int32(3)) {
			break
		}
		speedtest1_prepare(tls, __ccgo_ts+32295, libc.VaList(bp+2008, jj))
		ii = 0
		for {
			if !(ii < NROW2) {
				break
			}
			libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), ii*int32(3))
			speedtest1_run(tls)
			goto _4
		_4:
			;
			ii = ii + int32(3)
		}
		goto _3
	_3:
		;
		jj = jj + 1
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(130), __ccgo_ts+32329, 0)
	jj = int32(1)
	for {
		if !(jj <= int32(3)) {
			break
		}
		speedtest1_prepare(tls, __ccgo_ts+32344, libc.VaList(bp+2008, jj))
		ii = 0
		for {
			if !(ii < NROW2) {
				break
			}
			libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), ii*int32(3))
			speedtest1_run(tls)
			goto _6
		_6:
			;
			ii = ii + int32(3)
		}
		goto _5
	_5:
		;
		jj = jj + 1
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(140), __ccgo_ts+32329, 0)
	jj = int32(1)
	for {
		if !(jj <= int32(3)) {
			break
		}
		speedtest1_prepare(tls, __ccgo_ts+32344, libc.VaList(bp+2008, jj))
		ii = 0
		for {
			if !(ii < NROW2) {
				break
			}
			libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), ii*int32(3))
			speedtest1_run(tls)
			goto _8
		_8:
			;
			ii = ii + int32(3)
		}
		goto _7
	_7:
		;
		jj = jj + 1
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(150), __ccgo_ts+32378, 0)
	speedtest1_prepare(tls, __ccgo_ts+32397, 0)
	jj = 0
	for {
		if !(jj < NROW2) {
			break
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), jj*int32(3))
		speedtest1_run(tls)
		goto _9
	_9:
		;
		jj = jj + 1
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(160), __ccgo_ts+32512, 0)
	speedtest1_exec(tls, __ccgo_ts+648, 0)
	speedtest1_prepare(tls, __ccgo_ts+32533, 0)
	jj = 0
	for {
		if !(jj < NROW2) {
			break
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), jj)
		speedtest1_run(tls)
		goto _10
	_10:
		;
		jj = jj + 1
	}
	speedtest1_exec(tls, __ccgo_ts+758, 0)
	speedtest1_end_test(tls)
	speedtest1_exec(tls, __ccgo_ts+32568, 0)
	speedtest1_begin_test(tls, int32(170), __ccgo_ts+32616, 0)
	speedtest1_exec(tls, __ccgo_ts+32638, 0)
	speedtest1_end_test(tls)
	speedtest1_exec(tls, __ccgo_ts+32683, 0)
	speedtest1_begin_test(tls, int32(180), __ccgo_ts+33173, 0)
	speedtest1_prepare(tls, __ccgo_ts+33190, 0)
	jj = 0
	for {
		if !(jj < NROW2) {
			break
		}
		speedtest1_numbername(tls, libc.Uint32FromInt32(jj), bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), jj)
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(2), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _11
	_11:
		;
		jj = jj + 1
	}
	speedtest1_end_test(tls)
	/*
	 ** Note: Of the queries, only half actually update a row. This property
	 ** was copied over from speed4p.test, where it was probably introduced
	 ** inadvertantly.
	 */
	speedtest1_begin_test(tls, int32(190), __ccgo_ts+33226, 0)
	speedtest1_prepare(tls, __ccgo_ts+33243, 0)
	jj = int32(1)
	for {
		if !(jj <= NROW2*int32(2)) {
			break
		}
		speedtest1_numbername(tls, libc.Uint32FromInt32(jj*int32(2)), bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), jj*int32(2))
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(2), bp, -int32(1), libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(3), jj)
		speedtest1_run(tls)
		goto _12
	_12:
		;
		jj = jj + int32(2)
	}
	speedtest1_end_test(tls)
	/*
	 ** Note: Same again.
	 */
	speedtest1_begin_test(tls, int32(200), __ccgo_ts+33289, 0)
	speedtest1_prepare(tls, __ccgo_ts+33306, 0)
	jj = int32(1)
	for {
		if !(jj <= NROW2*int32(2)) {
			break
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), jj*int32(2))
		speedtest1_run(tls)
		goto _13
	_13:
		;
		jj = jj + int32(2)
	}
	speedtest1_end_test(tls)
	speedtest1_exec(tls, __ccgo_ts+758, 0)
	/*
	 ** The following block contains the same tests as the above block that
	 ** tests triggers, with one crucial difference: no triggers are defined.
	 ** So the difference in speed between these tests and the preceding ones
	 ** is the amount of time taken to compile and execute the trigger programs.
	 */
	speedtest1_exec(tls, __ccgo_ts+33338, 0)
	speedtest1_begin_test(tls, int32(210), __ccgo_ts+33443, 0)
	speedtest1_prepare(tls, __ccgo_ts+33190, 0)
	jj = 0
	for {
		if !(jj < NROW2) {
			break
		}
		speedtest1_numbername(tls, libc.Uint32FromInt32(jj), bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), jj)
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(2), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _14
	_14:
		;
		jj = jj + 1
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(210), __ccgo_ts+33462, 0)
	speedtest1_prepare(tls, __ccgo_ts+33243, 0)
	jj = int32(1)
	for {
		if !(jj <= NROW2*int32(2)) {
			break
		}
		speedtest1_numbername(tls, libc.Uint32FromInt32(jj*int32(2)), bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), jj*int32(2))
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(2), bp, -int32(1), libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(3), jj)
		speedtest1_run(tls)
		goto _15
	_15:
		;
		jj = jj + int32(2)
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(220), __ccgo_ts+33481, 0)
	speedtest1_prepare(tls, __ccgo_ts+33306, 0)
	jj = int32(1)
	for {
		if !(jj <= NROW2*int32(2)) {
			break
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), jj*int32(2))
		speedtest1_run(tls)
		goto _16
	_16:
		;
		jj = jj + int32(2)
	}
	speedtest1_end_test(tls)
	speedtest1_exec(tls, __ccgo_ts+758, 0)
}

// C documentation
//
//	/*
//	** A testset used for debugging speedtest1 itself.
//	*/
func testset_debug1(tls *libc.TLS) {
	bp := tls.Alloc(2048)
	defer tls.Free(2048)
	var i, n, x1, x2 uint32
	var _ /* zNum at bp+0 */ [2000]int8
	_, _, _, _ = i, n, x1, x2 /* A number name */
	n = libc.Uint32FromInt32(g.FszTest)
	i = uint32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = swizzle(tls, i, n)
		x2 = swizzle(tls, x1, n)
		speedtest1_numbername(tls, x1, bp, int32(2000))
		libc.Xprintf(tls, __ccgo_ts+33500, libc.VaList(bp+2008, i, x1, x2, bp))
		goto _1
	_1:
		;
		i = i + 1
	}
}

// C documentation
//
//	/*
//	** Performance tests for JSON.
//	*/
func testset_json(tls *libc.TLS) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var r uint32
	_ = r
	r = uint32(0x12345678)
	libsqlite3.Xsqlite3_test_control(tls, int32(SQLITE_TESTCTRL_PRNG_SEED), libc.VaList(bp+8, r, g.Fdb))
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+33516, libc.VaList(bp+8, g.FszTest*int32(5)))
	speedtest1_exec(tls, __ccgo_ts+33545, libc.VaList(bp+8, g.FszTest*int32(5)))
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(110), __ccgo_ts+34505, libc.VaList(bp+8, g.FszTest))
	speedtest1_exec(tls, __ccgo_ts+34551, libc.VaList(bp+8, g.FszTest))
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(120), __ccgo_ts+34631, 0)
	speedtest1_exec(tls, __ccgo_ts+34672, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(130), __ccgo_ts+34790, 0)
	speedtest1_exec(tls, __ccgo_ts+34831, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(140), __ccgo_ts+34949, 0)
	speedtest1_exec(tls, __ccgo_ts+34968, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(141), __ccgo_ts+35329, 0)
	speedtest1_exec(tls, __ccgo_ts+35359, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(150), __ccgo_ts+35541, 0)
	speedtest1_exec(tls, __ccgo_ts+35589, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(160), __ccgo_ts+35839, 0)
	speedtest1_exec(tls, __ccgo_ts+35887, 0)
	speedtest1_end_test(tls)
}

// C documentation
//
//	/*
//	** This testset focuses on the speed of parsing numeric literals (integers
//	** and real numbers). This was added to test the impact of allowing "_"
//	** characters to appear in numeric SQL literals to make them easier to read.
//	** For example, "SELECT 1_000_000;" instead of "SELECT 1000000;".
//	*/
func testset_parsenumber(tls *libc.TLS) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var NROW, ii int32
	var zSql1, zSql2, zSql3, zSql4 uintptr
	_, _, _, _, _, _ = NROW, ii, zSql1, zSql2, zSql3, zSql4
	zSql1 = __ccgo_ts + 36038
	zSql2 = __ccgo_ts + 36077
	zSql3 = __ccgo_ts + 36209
	zSql4 = __ccgo_ts + 36255
	NROW = int32(100) * g.FszTest
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+36393, libc.VaList(bp+8, NROW))
	ii = 0
	for {
		if !(ii < NROW) {
			break
		}
		libsqlite3.Xsqlite3_exec(tls, g.Fdb, zSql1, uintptr(0), uintptr(0), uintptr(0))
		goto _1
	_1:
		;
		ii = ii + 1
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(110), __ccgo_ts+36419, libc.VaList(bp+8, NROW))
	ii = 0
	for {
		if !(ii < NROW) {
			break
		}
		libsqlite3.Xsqlite3_exec(tls, g.Fdb, zSql2, uintptr(0), uintptr(0), uintptr(0))
		goto _2
	_2:
		;
		ii = ii + 1
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(200), __ccgo_ts+36445, libc.VaList(bp+8, NROW))
	ii = 0
	for {
		if !(ii < NROW) {
			break
		}
		libsqlite3.Xsqlite3_exec(tls, g.Fdb, zSql3, uintptr(0), uintptr(0), uintptr(0))
		goto _3
	_3:
		;
		ii = ii + 1
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(210), __ccgo_ts+36468, libc.VaList(bp+8, NROW))
	ii = 0
	for {
		if !(ii < NROW) {
			break
		}
		libsqlite3.Xsqlite3_exec(tls, g.Fdb, zSql4, uintptr(0), uintptr(0), uintptr(0))
		goto _4
	_4:
		;
		ii = ii + 1
	}
	speedtest1_end_test(tls)
}

func xCompileOptions(tls *libc.TLS, pCtx uintptr, nVal int32, azVal uintptr, azCol uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	libc.Xprintf(tls, __ccgo_ts+36491, libc.VaList(bp+8, **(**uintptr)(__ccgo_up(azVal))))
	return SQLITE_OK
}

func main1(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var cacheSize, doAutovac, doExclusive, doFullFSync, doIncrvac, doPCache, doTrace, i, kk, memDb, mmapSize, mnHeap, nHardHeapLmt, nHeap, nLook, nPCache, nSoftHeapLmt, nThread, noSync, openFlags, pageSize, rc, showStats, szLook, szPCache, v2, v3 int32
	var pHeap, pLook, pPCache, pVfs, z, zComma, zEncoding, zJMode, zKey, zObj, zSep, zSql, zTSet, zThisTest, v15 uintptr
	var v19, v20 __darwin_ct_rune_t
	var v23 bool
	var _ /* iCur at bp+0 */ int32
	var _ /* iHi at bp+4 */ int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = cacheSize, doAutovac, doExclusive, doFullFSync, doIncrvac, doPCache, doTrace, i, kk, memDb, mmapSize, mnHeap, nHardHeapLmt, nHeap, nLook, nPCache, nSoftHeapLmt, nThread, noSync, openFlags, pHeap, pLook, pPCache, pVfs, pageSize, rc, showStats, szLook, szPCache, z, zComma, zEncoding, zJMode, zKey, zObj, zSep, zSql, zTSet, zThisTest, v15, v19, v2, v20, v23, v3
	doAutovac = 0   /* True for --autovacuum */
	cacheSize = 0   /* Desired cache size.  0 means default */
	doExclusive = 0 /* True for --exclusive */
	doFullFSync = 0 /* True for --fullfsync */
	nHeap = 0
	mnHeap = 0          /* Heap size from --heap */
	doIncrvac = 0       /* True for --incrvacuum */
	zJMode = uintptr(0) /* Journal mode */
	zKey = uintptr(0)   /* Encryption key */
	nHardHeapLmt = 0    /* The hard heap limit */
	nSoftHeapLmt = 0    /* The soft heap limit */
	nLook = -int32(1)
	szLook = 0   /* --lookaside configuration */
	noSync = 0   /* True for --nosync */
	pageSize = 0 /* Desired page size.  0 means default */
	nPCache = 0
	szPCache = 0                                                                                     /* --pcache configuration */
	doPCache = 0                                                                                     /* True if --pcache is seen */
	showStats = 0                                                                                    /* True for --stats */
	nThread = 0                                                                                      /* --threads value */
	mmapSize = 0                                                                                     /* How big of a memory map to use */
	memDb = 0                                                                                        /* --memdb.  Use an in-memory database */
	openFlags = libc.Int32FromInt32(SQLITE_OPEN_READWRITE) | libc.Int32FromInt32(SQLITE_OPEN_CREATE) /* SQLITE_OPEN_xxx flags. */
	zTSet = __ccgo_ts + 36514                                                                        /* Which --testset torun */
	doTrace = 0                                                                                      /* True for --trace */
	zEncoding = uintptr(0)                                                                           /* --utf16be or --utf16le */
	pHeap = uintptr(0)                                                                               /* Allocated heap space */
	pLook = uintptr(0)                                                                               /* Allocated lookaside space */
	pPCache = uintptr(0)                                                                             /* API return code */
	/*
	 ** Confirms that argc has at least N arguments following argv[i]. */
	/* Display the version of SQLite being tested */
	libc.Xprintf(tls, __ccgo_ts+36519, libc.VaList(bp+16, libsqlite3.Xsqlite3_libversion(tls), libsqlite3.Xsqlite3_sourceid(tls)))
	/* Process command-line arguments */
	g.FzDbName = uintptr(0)
	g.FzVfs = uintptr(0)
	g.FzWR = __ccgo_ts + 6
	g.FzNN = __ccgo_ts + 6
	g.FzPK = __ccgo_ts + 36554
	g.FszTest = int32(100)
	g.FszBase = int32(100)
	g.FnRepeat = int32(1)
	i = int32(1)
	for {
		if !(i < argc) {
			break
		}
		z = **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))
		if int32(**(**int8)(__ccgo_up(z))) == int32('-') {
			for cond := true; cond; cond = int32(**(**int8)(__ccgo_up(z))) == int32('-') {
				z = z + 1
			}
			if libc.Xstrcmp(tls, z, __ccgo_ts+36561) == 0 {
				doAutovac = int32(1)
			} else {
				if libc.Xstrcmp(tls, z, __ccgo_ts+36572) == 0 {
					g.FdoBigTransactions = int32(1)
				} else {
					if libc.Xstrcmp(tls, z, __ccgo_ts+36589) == 0 {
						if i >= argc-int32(1) {
							fatal_error(tls, __ccgo_ts+36599, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
						}
						i = i + 1
						v2 = i
						cacheSize = integerValue(tls, **(**uintptr)(__ccgo_up(argv + uintptr(v2)*8)))
					} else {
						if libc.Xstrcmp(tls, z, __ccgo_ts+36623) == 0 {
							doExclusive = int32(1)
						} else {
							if libc.Xstrcmp(tls, z, __ccgo_ts+36633) == 0 {
								doFullFSync = int32(1)
							} else {
								if libc.Xstrcmp(tls, z, __ccgo_ts+36643) == 0 {
									g.FdoCheckpoint = int32(1)
								} else {
									if libc.Xstrcmp(tls, z, __ccgo_ts+36654) == 0 {
										g.FbSqlOnly = int32(1)
										g.FbExplain = int32(1)
									} else {
										if libc.Xstrcmp(tls, z, __ccgo_ts+36662) == 0 {
											if i >= argc-int32(1) {
												fatal_error(tls, __ccgo_ts+36599, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
											}
											nHardHeapLmt = integerValue(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i+int32(1))*8)))
											i = i + int32(1)
										} else {
											if libc.Xstrcmp(tls, z, __ccgo_ts+36678) == 0 {
												if i >= argc-int32(2) {
													fatal_error(tls, __ccgo_ts+36599, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
												}
												nHeap = integerValue(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i+int32(1))*8)))
												mnHeap = integerValue(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i+int32(2))*8)))
												i = i + int32(2)
											} else {
												if libc.Xstrcmp(tls, z, __ccgo_ts+36683) == 0 {
													doIncrvac = int32(1)
												} else {
													if libc.Xstrcmp(tls, z, __ccgo_ts+36694) == 0 {
														if i >= argc-int32(1) {
															fatal_error(tls, __ccgo_ts+36599, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
														}
														i = i + 1
														v2 = i
														zJMode = **(**uintptr)(__ccgo_up(argv + uintptr(v2)*8))
													} else {
														if libc.Xstrcmp(tls, z, __ccgo_ts+36702) == 0 {
															if i >= argc-int32(1) {
																fatal_error(tls, __ccgo_ts+36599, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
															}
															i = i + 1
															v2 = i
															zKey = **(**uintptr)(__ccgo_up(argv + uintptr(v2)*8))
														} else {
															if libc.Xstrcmp(tls, z, __ccgo_ts+36706) == 0 {
																if i >= argc-int32(2) {
																	fatal_error(tls, __ccgo_ts+36599, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
																}
																nLook = integerValue(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i+int32(1))*8)))
																szLook = integerValue(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i+int32(2))*8)))
																i = i + int32(2)
															} else {
																if libc.Xstrcmp(tls, z, __ccgo_ts+36716) == 0 {
																	memDb = int32(1)
																} else {
																	if libc.Xstrcmp(tls, z, __ccgo_ts+36722) == 0 {
																		libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_MULTITHREAD), 0)
																	} else {
																		if libc.Xstrcmp(tls, z, __ccgo_ts+36734) == 0 {
																			libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_MEMSTATUS), libc.VaList(bp+16, 0))
																		} else {
																			if libc.Xstrcmp(tls, z, __ccgo_ts+36744) == 0 {
																				if i >= argc-int32(1) {
																					fatal_error(tls, __ccgo_ts+36599, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
																				}
																				i = i + 1
																				v2 = i
																				mmapSize = integerValue(tls, **(**uintptr)(__ccgo_up(argv + uintptr(v2)*8)))
																			} else {
																				if libc.Xstrcmp(tls, z, __ccgo_ts+36749) == 0 {
																					openFlags = openFlags | int32(SQLITE_OPEN_NOMUTEX)
																				} else {
																					if libc.Xstrcmp(tls, z, __ccgo_ts+36757) == 0 {
																						noSync = int32(1)
																					} else {
																						if libc.Xstrcmp(tls, z, __ccgo_ts+36764) == 0 {
																							g.FzNN = __ccgo_ts + 36772
																						} else {
																							if libc.Xstrcmp(tls, z, __ccgo_ts+36781) == 0 {
																								if i >= argc-int32(1) {
																									fatal_error(tls, __ccgo_ts+36599, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
																								}
																								i = i + 1
																								if libc.Xstrcmp(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)), __ccgo_ts+36788) == 0 {
																									g.FhashFile = libc.X__stdoutp
																								} else {
																									g.FhashFile = libc.Xfopen(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)), __ccgo_ts+36790)
																									if g.FhashFile == uintptr(0) {
																										fatal_error(tls, __ccgo_ts+36793, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
																									}
																								}
																							} else {
																								if libc.Xstrcmp(tls, z, __ccgo_ts+36823) == 0 {
																									if i >= argc-int32(1) {
																										fatal_error(tls, __ccgo_ts+36599, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
																									}
																									i = i + 1
																									v2 = i
																									pageSize = integerValue(tls, **(**uintptr)(__ccgo_up(argv + uintptr(v2)*8)))
																								} else {
																									if libc.Xstrcmp(tls, z, __ccgo_ts+36832) == 0 {
																										if i >= argc-int32(2) {
																											fatal_error(tls, __ccgo_ts+36599, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
																										}
																										nPCache = integerValue(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i+int32(1))*8)))
																										szPCache = integerValue(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i+int32(2))*8)))
																										doPCache = int32(1)
																										i = i + int32(2)
																									} else {
																										if libc.Xstrcmp(tls, z, __ccgo_ts+36839) == 0 {
																											g.FzPK = __ccgo_ts + 36850
																										} else {
																											if libc.Xstrcmp(tls, z, __ccgo_ts+36862) == 0 {
																												if i >= argc-int32(1) {
																													fatal_error(tls, __ccgo_ts+36599, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
																												}
																												i = i + 1
																												v2 = i
																												g.FnRepeat = integerValue(tls, **(**uintptr)(__ccgo_up(argv + uintptr(v2)*8)))
																											} else {
																												if libc.Xstrcmp(tls, z, __ccgo_ts+36869) == 0 {
																													g.FbReprepare = int32(1)
																												} else {
																													if libc.Xstrcmp(tls, z, __ccgo_ts+36879) == 0 {
																														libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_SERIALIZED), 0)
																													} else {
																														if libc.Xstrcmp(tls, z, __ccgo_ts+36890) == 0 {
																															libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_SINGLETHREAD), 0)
																														} else {
																															if libc.Xstrcmp(tls, z, __ccgo_ts+36903) == 0 {
																																if i >= argc-int32(1) {
																																	fatal_error(tls, __ccgo_ts+36599, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
																																}
																																if g.FpScript != 0 {
																																	libc.Xfclose(tls, g.FpScript)
																																}
																																i = i + 1
																																v2 = i
																																g.FpScript = libc.Xfopen(tls, **(**uintptr)(__ccgo_up(argv + uintptr(v2)*8)), __ccgo_ts+36790)
																																if g.FpScript == uintptr(0) {
																																	fatal_error(tls, __ccgo_ts+36910, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
																																}
																															} else {
																																if libc.Xstrcmp(tls, z, __ccgo_ts+36943) == 0 {
																																	g.FbSqlOnly = int32(1)
																																} else {
																																	if libc.Xstrcmp(tls, z, __ccgo_ts+36951) == 0 {
																																		g.FbMemShrink = int32(1)
																																	} else {
																																		if libc.Xstrcmp(tls, z, __ccgo_ts+36965) == 0 {
																																			if i >= argc-int32(1) {
																																				fatal_error(tls, __ccgo_ts+36599, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
																																			}
																																			i = i + 1
																																			v3 = i
																																			v2 = integerValue(tls, **(**uintptr)(__ccgo_up(argv + uintptr(v3)*8)))
																																			g.FszBase = v2
																																			g.FszTest = v2
																																		} else {
																																			if libc.Xstrcmp(tls, z, __ccgo_ts+36970) == 0 {
																																				if i >= argc-int32(1) {
																																					fatal_error(tls, __ccgo_ts+36599, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
																																				}
																																				nSoftHeapLmt = integerValue(tls, **(**uintptr)(__ccgo_up(argv + uintptr(i+int32(1))*8)))
																																				i = i + int32(1)
																																			} else {
																																				if libc.Xstrcmp(tls, z, __ccgo_ts+36986) == 0 {
																																					showStats = int32(1)
																																				} else {
																																					if libc.Xstrcmp(tls, z, __ccgo_ts+36992) == 0 {
																																						if i >= argc-int32(1) {
																																							fatal_error(tls, __ccgo_ts+36599, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
																																						}
																																						i = i + 1
																																						if int32(**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))) < int32('0') || int32(**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))) > int32('9') || int32(**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(argv + uintptr(i)*8)) + 1))) != 0 {
																																							fatal_error(tls, __ccgo_ts+36997, 0)
																																						}
																																						g.FeTemp = int32(**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))) - int32('0')
																																					} else {
																																						if libc.Xstrcmp(tls, z, __ccgo_ts+37050) == 0 {
																																							if i >= argc-int32(1) {
																																								fatal_error(tls, __ccgo_ts+36599, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
																																							}
																																							i = i + 1
																																							v2 = i
																																							zTSet = **(**uintptr)(__ccgo_up(argv + uintptr(v2)*8))
																																						} else {
																																							if libc.Xstrcmp(tls, z, __ccgo_ts+37058) == 0 {
																																								doTrace = int32(1)
																																							} else {
																																								if libc.Xstrcmp(tls, z, __ccgo_ts+37064) == 0 {
																																									if i >= argc-int32(1) {
																																										fatal_error(tls, __ccgo_ts+36599, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
																																									}
																																									i = i + 1
																																									v2 = i
																																									nThread = integerValue(tls, **(**uintptr)(__ccgo_up(argv + uintptr(v2)*8)))
																																								} else {
																																									if libc.Xstrcmp(tls, z, __ccgo_ts+37072) == 0 {
																																										zEncoding = __ccgo_ts + 37072
																																									} else {
																																										if libc.Xstrcmp(tls, z, __ccgo_ts+37080) == 0 {
																																											zEncoding = __ccgo_ts + 37080
																																										} else {
																																											if libc.Xstrcmp(tls, z, __ccgo_ts+37088) == 0 {
																																												g.FbVerify = int32(1)
																																												HashInit(tls)
																																											} else {
																																												if libc.Xstrcmp(tls, z, __ccgo_ts+37095) == 0 {
																																													if i >= argc-int32(1) {
																																														fatal_error(tls, __ccgo_ts+36599, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
																																													}
																																													i = i + 1
																																													v2 = i
																																													g.FzVfs = **(**uintptr)(__ccgo_up(argv + uintptr(v2)*8))
																																												} else {
																																													if libc.Xstrcmp(tls, z, __ccgo_ts+37099) == 0 {
																																														if i >= argc-int32(1) {
																																															fatal_error(tls, __ccgo_ts+36599, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))))
																																														}
																																														i = i + 1
																																														v2 = i
																																														g.FnReserve = libc.Xatoi(tls, **(**uintptr)(__ccgo_up(argv + uintptr(v2)*8)))
																																													} else {
																																														if libc.Xstrcmp(tls, z, __ccgo_ts+37107) == 0 {
																																															g.FstmtScanStatus = int32(1)
																																														} else {
																																															if libc.Xstrcmp(tls, z, __ccgo_ts+37122) == 0 {
																																																if libc.Xstrstr(tls, g.FzWR, __ccgo_ts+37136) != uintptr(0) {
																																																	/* no-op */
																																																} else {
																																																	if libc.Xstrstr(tls, g.FzWR, __ccgo_ts+37144) != uintptr(0) {
																																																		g.FzWR = __ccgo_ts + 37151
																																																	} else {
																																																		g.FzWR = __ccgo_ts + 3725
																																																	}
																																																}
																																																g.FzPK = __ccgo_ts + 36850
																																															} else {
																																																if libc.Xstrcmp(tls, z, __ccgo_ts+37172) == 0 {
																																																	if libc.Xstrstr(tls, g.FzWR, __ccgo_ts+37144) != uintptr(0) {
																																																		/* no-op */
																																																	} else {
																																																		if libc.Xstrstr(tls, g.FzWR, __ccgo_ts+37136) != uintptr(0) {
																																																			g.FzWR = __ccgo_ts + 37151
																																																		} else {
																																																			g.FzWR = __ccgo_ts + 37144
																																																		}
																																																	}
																																																} else {
																																																	if libc.Xstrcmp(tls, z, __ccgo_ts+37179) == 0 || libc.Xstrcmp(tls, z, __ccgo_ts+37184) == 0 {
																																																		libc.Xprintf(tls, uintptr(unsafe.Pointer(&zHelp)), libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv))))
																																																		libc.Xexit(tls, 0)
																																																	} else {
																																																		fatal_error(tls, __ccgo_ts+37186, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)), **(**uintptr)(__ccgo_up(argv))))
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
		} else {
			if g.FzDbName == uintptr(0) {
				g.FzDbName = **(**uintptr)(__ccgo_up(argv + uintptr(i)*8))
			} else {
				fatal_error(tls, __ccgo_ts+37227, libc.VaList(bp+16, **(**uintptr)(__ccgo_up(argv + uintptr(i)*8)), **(**uintptr)(__ccgo_up(argv))))
			}
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	if nHeap > 0 {
		pHeap = libc.Xmalloc(tls, libc.Uint64FromInt32(nHeap))
		if pHeap == uintptr(0) {
			fatal_error(tls, __ccgo_ts+37270, libc.VaList(bp+16, nHeap))
		}
		rc = libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_HEAP), libc.VaList(bp+16, pHeap, nHeap, mnHeap))
		if rc != 0 {
			fatal_error(tls, __ccgo_ts+37300, libc.VaList(bp+16, rc))
		}
	}
	if doPCache != 0 {
		if nPCache > 0 && szPCache > 0 {
			pPCache = libc.Xmalloc(tls, libc.Uint64FromInt64(int64(nPCache)*int64(szPCache)))
			if pPCache == uintptr(0) {
				fatal_error(tls, __ccgo_ts+37331, libc.VaList(bp+16, int64(nPCache)*int64(szPCache)))
			}
		}
		rc = libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_PAGECACHE), libc.VaList(bp+16, pPCache, szPCache, nPCache))
		if rc != 0 {
			fatal_error(tls, __ccgo_ts+37365, libc.VaList(bp+16, rc))
		}
	}
	if nLook >= 0 {
		libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_LOOKASIDE), libc.VaList(bp+16, 0, 0))
	}
	libsqlite3.Xsqlite3_initialize(tls)
	if g.FzDbName != uintptr(0) {
		pVfs = libsqlite3.Xsqlite3_vfs_find(tls, g.FzVfs)
		/* For some VFSes, e.g. opfs, unlink() is not sufficient. Use the
		 ** selected (or default) VFS's xDelete method to delete the
		 ** database. This is specifically important for the "opfs" VFS
		 ** when running from a WASM build of speedtest1, so that the db
		 ** can be cleaned up properly. For historical compatibility, we'll
		 ** also simply unlink(). */
		if pVfs != uintptr(0) {
			(*(*func(*libc.TLS, uintptr, uintptr, int32) int32)(unsafe.Pointer(&struct{ uintptr }{(*sqlite3_vfs)(unsafe.Pointer(pVfs)).FxDelete})))(tls, pVfs, g.FzDbName, int32(1))
		}
		libc.Xunlink(tls, g.FzDbName)
	}
	/* Open the database and the input file */
	if memDb != 0 {
		v15 = __ccgo_ts + 37398
	} else {
		v15 = g.FzDbName
	}
	if libsqlite3.Xsqlite3_open_v2(tls, v15, uintptr(unsafe.Pointer(&g)), openFlags, g.FzVfs) != 0 {
		fatal_error(tls, __ccgo_ts+18708, libc.VaList(bp+16, g.FzDbName))
	}
	if nLook > 0 && szLook > 0 {
		pLook = libc.Xmalloc(tls, libc.Uint64FromInt32(nLook*szLook))
		rc = libsqlite3.Xsqlite3_db_config(tls, g.Fdb, int32(SQLITE_DBCONFIG_LOOKASIDE), libc.VaList(bp+16, pLook, szLook, nLook))
		if rc != 0 {
			fatal_error(tls, __ccgo_ts+37407, libc.VaList(bp+16, rc))
		}
	}
	if g.FnReserve > 0 {
		libsqlite3.Xsqlite3_file_control(tls, g.Fdb, uintptr(0), int32(SQLITE_FCNTL_RESERVE_BYTES), uintptr(unsafe.Pointer(&g))+92)
	}
	if g.FstmtScanStatus != 0 {
		libsqlite3.Xsqlite3_db_config(tls, g.Fdb, int32(SQLITE_DBCONFIG_STMT_SCANSTATUS), libc.VaList(bp+16, int32(1), 0))
	}
	/* Set database connection options */
	libsqlite3.Xsqlite3_create_function(tls, g.Fdb, __ccgo_ts+37443, 0, int32(SQLITE_UTF8), uintptr(0), __ccgo_fp(randomFunc), uintptr(0), uintptr(0))
	if doTrace != 0 {
		libsqlite3.Xsqlite3_trace(tls, g.Fdb, __ccgo_fp(traceCallback), uintptr(0))
	}
	if memDb > 0 {
		speedtest1_exec(tls, __ccgo_ts+37450, 0)
	}
	if mmapSize > 0 {
		speedtest1_exec(tls, __ccgo_ts+37475, libc.VaList(bp+16, mmapSize))
	}
	speedtest1_exec(tls, __ccgo_ts+37495, libc.VaList(bp+16, nThread))
	if zKey != 0 {
		speedtest1_exec(tls, __ccgo_ts+37513, libc.VaList(bp+16, zKey))
	}
	if zEncoding != 0 {
		speedtest1_exec(tls, __ccgo_ts+37530, libc.VaList(bp+16, zEncoding))
	}
	if doAutovac != 0 {
		speedtest1_exec(tls, __ccgo_ts+37549, 0)
	} else {
		if doIncrvac != 0 {
			speedtest1_exec(tls, __ccgo_ts+37573, 0)
		}
	}
	if pageSize != 0 {
		speedtest1_exec(tls, __ccgo_ts+37604, libc.VaList(bp+16, pageSize))
	}
	if cacheSize != 0 {
		speedtest1_exec(tls, __ccgo_ts+37624, libc.VaList(bp+16, cacheSize))
	}
	if noSync != 0 {
		speedtest1_exec(tls, __ccgo_ts+37645, 0)
	} else {
		if doFullFSync != 0 {
			speedtest1_exec(tls, __ccgo_ts+37668, 0)
		}
	}
	if doExclusive != 0 {
		speedtest1_exec(tls, __ccgo_ts+37688, 0)
	}
	if zJMode != 0 {
		speedtest1_exec(tls, __ccgo_ts+37718, libc.VaList(bp+16, zJMode))
	}
	if nHardHeapLmt > 0 {
		speedtest1_exec(tls, __ccgo_ts+37741, libc.VaList(bp+16, nHardHeapLmt))
	}
	if nSoftHeapLmt > 0 {
		speedtest1_exec(tls, __ccgo_ts+37767, libc.VaList(bp+16, nSoftHeapLmt))
	}
	if zJMode != 0 {
		speedtest1_exec(tls, __ccgo_ts+37718, libc.VaList(bp+16, zJMode))
	}
	if g.FbExplain != 0 {
		libc.Xprintf(tls, __ccgo_ts+37793, 0)
	}
	if libc.Xstrcmp(tls, zTSet, __ccgo_ts+36514) == 0 {
		zTSet = uintptr(unsafe.Pointer(&zMix1Tests))
	}
	for cond := true; cond; cond = **(**int8)(__ccgo_up(zTSet)) != 0 {
		zThisTest = zTSet
		zComma = libc.Xstrchr(tls, zThisTest, int32(','))
		if zComma != 0 {
			**(**int8)(__ccgo_up(zComma)) = 0
			zTSet = zComma + uintptr(1)
		} else {
			zTSet = __ccgo_ts + 6
		}
		zSep = libc.Xstrchr(tls, zThisTest, int32('/'))
		if zSep != 0 {
			kk = int32(1)
			for {
				if v23 = **(**int8)(__ccgo_up(zSep + uintptr(kk))) != 0; v23 {
					v19 = libc.Int32FromUint8(libc.Uint8FromInt8(**(**int8)(__ccgo_up(zSep + uintptr(kk)))))
					if v19 < 0 || v19 >= libc.Int32FromInt32(1)<<libc.Int32FromInt32(8) {
						v3 = 0
					} else {
						v3 = libc.BoolInt32(!!(uint64(**(**__uint32_t)(__ccgo_up(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v19)*4)))&uint64(0x00000400) != 0))
					}
					v20 = v3
					goto _21
				_21:
					v2 = v20
					goto _18
				_18:
				}
				if !(v23 && v2 != 0) {
					break
				}
				goto _16
			_16:
				;
				kk = kk + 1
			}
			if kk == int32(1) || int32(**(**int8)(__ccgo_up(zSep + uintptr(kk)))) != 0 {
				fatal_error(tls, __ccgo_ts+37812, libc.VaList(bp+16, zThisTest))
			}
			g.FszTest = g.FszBase * integerValue(tls, zSep+uintptr(1)) / int32(100)
			if g.FszTest <= 0 {
				g.FszTest = int32(1)
			}
			**(**int8)(__ccgo_up(zSep)) = 0
		} else {
			g.FszTest = g.FszBase
		}
		if g.FiTotal > 0 || zComma == uintptr(0) {
			libc.Xprintf(tls, __ccgo_ts+37847, libc.VaList(bp+16, zThisTest))
		}
		if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+37874) == 0 {
			testset_main(tls)
		} else {
			if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+37879) == 0 {
				testset_debug1(tls)
			} else {
				if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+37886) == 0 {
					testset_orm(tls)
				} else {
					if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+37890) == 0 {
						testset_cte(tls)
					} else {
						if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+37894) == 0 {
							testset_star(tls)
						} else {
							if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+37899) == 0 {
								testset_app(tls)
							} else {
								if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+37903) == 0 {
									testset_fp(tls)
								} else {
									if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+37906) == 0 {
										testset_json(tls)
									} else {
										if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+37911) == 0 {
											testset_trigger(tls)
										} else {
											if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+37919) == 0 {
												testset_parsenumber(tls)
											} else {
												if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+37931) == 0 {
													testset_rtree(tls, int32(6), int32(147))
												} else {
													fatal_error(tls, __ccgo_ts+37937, libc.VaList(bp+16, zThisTest))
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
		if **(**int8)(__ccgo_up(zTSet)) != 0 {
			speedtest1_begin_test(tls, int32(999), __ccgo_ts+38006, 0)
			for int32(1) != 0 {
				zObj = speedtest1_once(tls, __ccgo_ts+38025, 0)
				if zObj == uintptr(0) {
					break
				}
				zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+38095, libc.VaList(bp+16, zObj))
				speedtest1_exec(tls, zSql, 0)
				libsqlite3.Xsqlite3_free(tls, zSql)
				libsqlite3.Xsqlite3_free(tls, zObj)
			}
			for int32(1) != 0 {
				zObj = speedtest1_once(tls, __ccgo_ts+38116, 0)
				if zObj == uintptr(0) {
					break
				}
				zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+38095, libc.VaList(bp+16, zObj))
				speedtest1_exec(tls, zSql, 0)
				libsqlite3.Xsqlite3_free(tls, zSql)
				libsqlite3.Xsqlite3_free(tls, zObj)
			}
			speedtest1_end_test(tls)
		}
	}
	speedtest1_final(tls)
	if showStats != 0 {
		libsqlite3.Xsqlite3_exec(tls, g.Fdb, __ccgo_ts+38186, __ccgo_fp(xCompileOptions), uintptr(0), uintptr(0))
	}
	/* Database connection statistics printed after both prepared statements
	 ** have been finalized */
	if showStats != 0 {
		libsqlite3.Xsqlite3_db_status(tls, g.Fdb, SQLITE_DBSTATUS_LOOKASIDE_USED, bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+38209, libc.VaList(bp+16, **(**int32)(__ccgo_up(bp)), **(**int32)(__ccgo_up(bp + 4))))
		libsqlite3.Xsqlite3_db_status(tls, g.Fdb, int32(SQLITE_DBSTATUS_LOOKASIDE_HIT), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+38254, libc.VaList(bp+16, **(**int32)(__ccgo_up(bp + 4))))
		libsqlite3.Xsqlite3_db_status(tls, g.Fdb, int32(SQLITE_DBSTATUS_LOOKASIDE_MISS_SIZE), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+38290, libc.VaList(bp+16, **(**int32)(__ccgo_up(bp + 4))))
		libsqlite3.Xsqlite3_db_status(tls, g.Fdb, int32(SQLITE_DBSTATUS_LOOKASIDE_MISS_FULL), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+38326, libc.VaList(bp+16, **(**int32)(__ccgo_up(bp + 4))))
		libsqlite3.Xsqlite3_db_status(tls, g.Fdb, int32(SQLITE_DBSTATUS_CACHE_USED), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+38362, libc.VaList(bp+16, **(**int32)(__ccgo_up(bp))))
		libsqlite3.Xsqlite3_db_status(tls, g.Fdb, int32(SQLITE_DBSTATUS_CACHE_HIT), bp, bp+4, int32(1))
		libc.Xprintf(tls, __ccgo_ts+38404, libc.VaList(bp+16, **(**int32)(__ccgo_up(bp))))
		libsqlite3.Xsqlite3_db_status(tls, g.Fdb, int32(SQLITE_DBSTATUS_CACHE_MISS), bp, bp+4, int32(1))
		libc.Xprintf(tls, __ccgo_ts+38440, libc.VaList(bp+16, **(**int32)(__ccgo_up(bp))))
		libsqlite3.Xsqlite3_db_status(tls, g.Fdb, int32(SQLITE_DBSTATUS_CACHE_WRITE), bp, bp+4, int32(1))
		libc.Xprintf(tls, __ccgo_ts+38476, libc.VaList(bp+16, **(**int32)(__ccgo_up(bp))))
		libsqlite3.Xsqlite3_db_status(tls, g.Fdb, int32(SQLITE_DBSTATUS_SCHEMA_USED), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+38512, libc.VaList(bp+16, **(**int32)(__ccgo_up(bp))))
		libsqlite3.Xsqlite3_db_status(tls, g.Fdb, int32(SQLITE_DBSTATUS_STMT_USED), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+38554, libc.VaList(bp+16, **(**int32)(__ccgo_up(bp))))
	}
	libsqlite3.Xsqlite3_close(tls, g.Fdb)
	/* Global memory usage statistics printed after the database connection
	 ** has closed.  Memory usage should be zero at this point. */
	if showStats != 0 {
		libsqlite3.Xsqlite3_status(tls, SQLITE_STATUS_MEMORY_USED, bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+38596, libc.VaList(bp+16, **(**int32)(__ccgo_up(bp)), **(**int32)(__ccgo_up(bp + 4))))
		libsqlite3.Xsqlite3_status(tls, int32(SQLITE_STATUS_MALLOC_COUNT), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+38641, libc.VaList(bp+16, **(**int32)(__ccgo_up(bp)), **(**int32)(__ccgo_up(bp + 4))))
		libsqlite3.Xsqlite3_status(tls, int32(SQLITE_STATUS_PAGECACHE_OVERFLOW), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+38686, libc.VaList(bp+16, **(**int32)(__ccgo_up(bp)), **(**int32)(__ccgo_up(bp + 4))))
		libsqlite3.Xsqlite3_status(tls, int32(SQLITE_STATUS_MALLOC_SIZE), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+38731, libc.VaList(bp+16, **(**int32)(__ccgo_up(bp + 4))))
		libsqlite3.Xsqlite3_status(tls, int32(SQLITE_STATUS_PAGECACHE_SIZE), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+38773, libc.VaList(bp+16, **(**int32)(__ccgo_up(bp + 4))))
	}
	if g.FpScript != 0 {
		libc.Xfclose(tls, g.FpScript)
	}
	/* Release memory */
	libc.Xfree(tls, pLook)
	libc.Xfree(tls, pPCache)
	libc.Xfree(tls, pHeap)
	return 0
}

/* "mix1" is a macro testset: */
var zMix1Tests = [62]int8{'m', 'a', 'i', 'n', ',', 'o', 'r', 'm', '/', '2', '5', ',', 'c', 't', 'e', '/', '2', '0', ',', 'j', 's', 'o', 'n', ',', 'f', 'p', '/', '3', ',', 'p', 'a', 'r', 's', 'e', 'n', 'u', 'm', 'b', 'e', 'r', '/', '2', '5', ',', 'r', 't', 'r', 'e', 'e', '/', '1', '0', ',', 's', 't', 'a', 'r', ',', 'a', 'p', 'p'}

func main() {
	libc.Start(main1)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

func __ccgo_up(n uintptr) unsafe.Pointer {
	return unsafe.Pointer(&n)
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = " TEMP\x00\x00KiB\x00MiB\x00GiB\x00KB\x00MB\x00GB\x00K\x00M\x00G\x00parameter too large - max 2147483648\x00zero\x00one\x00two\x00three\x00four\x00five\x00six\x00seven\x00eight\x00nine\x00ten\x00eleven\x00twelve\x00thirteen\x00fourteen\x00fifteen\x00sixteen\x00seventeen\x00eighteen\x00nineteen\x00twenty\x00thirty\x00forty\x00fifty\x00sixty\x00seventy\x00eighty\x00ninety\x00 billion\x00 million\x00 thousand\x00%s hundred\x00%s\x00-- begin test %d %.*s\n\x00/* %4d - %s%.*s */\n\x00%4d - %s%.*s \x00PRAGMA wal_checkpoint;\x00-- end test %d\n\x00%4d.%03ds\n\x00       TOTAL%.*s %4d.%03ds\n\x00Verification Hash: %llu \x00\n\x00%02x\x00EXPLAIN \x00%.*s;\n\x00CREATE *\x00DROP *\x00ALTER *\x00%s;\n\x00SQL error: %s\n%s\n\x00exec error: %s\n\x00SQL error: %s\n\x00%s\n\x00%s\nError code %d: %s\n\x00nil\x00-IFTBN\x000123456789abcdef\x00%d INSERTs into table with no index\x00BEGIN\x00CREATE%s TABLE z1(a INTEGER %s, b INTEGER %s, c TEXT %s);\x00INSERT INTO z1 VALUES(?1,?2,?3); --  %d times\x00COMMIT\x00%d ordered INSERTS with one index/PK\x00CREATE%s TABLE z2(a INTEGER %s %s, b INTEGER %s, c TEXT %s) %s\x00INSERT INTO z2 VALUES(?1,?2,?3); -- %d times\x00%d unordered INSERTS with one index/PK\x00CREATE%s TABLE t3(a INTEGER %s %s, b INTEGER %s, c TEXT %s) %s\x00INSERT INTO t3 VALUES(?1,?2,?3); -- %d times\x00%d SELECTS, numeric BETWEEN, unindexed\x00SELECT count(*), avg(b), sum(length(c)), group_concat(c) FROM z1\n WHERE b BETWEEN ?1 AND ?2; -- %d times\x00%d SELECTS, LIKE, unindexed\x00SELECT count(*), avg(b), sum(length(c)), group_concat(c) FROM z1\n WHERE c LIKE ?1; -- %d times\x00%d SELECTS w/ORDER BY, unindexed\x00SELECT a, b, c FROM z1 WHERE c LIKE ?1\n ORDER BY a; -- %d times\x00%d SELECTS w/ORDER BY and LIMIT, unindexed\x00SELECT a, b, c FROM z1 WHERE c LIKE ?1\n ORDER BY a LIMIT 10; -- %d times\x00CREATE INDEX five times\x00BEGIN;\x00CREATE UNIQUE INDEX t1b ON z1(b);\x00CREATE INDEX t1c ON z1(c);\x00CREATE UNIQUE INDEX t2b ON z2(b);\x00CREATE INDEX t2c ON z2(c DESC);\x00CREATE INDEX t3bc ON t3(b,c);\x00COMMIT;\x00%d SELECTS, numeric BETWEEN, indexed\x00SELECT count(*), avg(b), sum(length(c)), group_concat(a) FROM z1\n WHERE b BETWEEN ?1 AND ?2; -- %d times\x00%d SELECTS, numeric BETWEEN, PK\x00SELECT count(*), avg(b), sum(length(c)), group_concat(a) FROM z2\n WHERE a BETWEEN ?1 AND ?2; -- %d times\x00%d SELECTS, text BETWEEN, indexed\x00SELECT count(*), avg(b), sum(length(c)), group_concat(a) FROM z1\n WHERE c BETWEEN ?1 AND (?1||'~'); -- %d times\x00%d INSERTS with three indexes\x00CREATE%s TABLE t4(\n  a INTEGER %s %s,\n  b INTEGER %s,\n  c TEXT %s\n) %s\x00CREATE INDEX t4b ON t4(b)\x00CREATE INDEX t4c ON t4(c)\x00INSERT INTO t4 SELECT * FROM z1\x00DELETE and REFILL one table\x00DELETE FROM z2;\x00INSERT INTO z2 SELECT * FROM z1;\x00VACUUM\x00ALTER TABLE ADD COLUMN, and query\x00ALTER TABLE z2 ADD COLUMN d INT DEFAULT 123\x00SELECT sum(d) FROM z2\x00%d UPDATES, numeric BETWEEN, indexed\x00UPDATE z2 SET d=b*2 WHERE b BETWEEN ?1 AND ?2; -- %d times\x00%d UPDATES of individual rows\x00UPDATE z2 SET d=b*3 WHERE a=?1; -- %d times\x00One big UPDATE of the whole %d-row table\x00UPDATE z2 SET d=b*4\x00Query added column after filling\x00%d DELETEs, numeric BETWEEN, indexed\x00DELETE FROM z2 WHERE b BETWEEN ?1 AND ?2; -- %d times\x00%d DELETEs of individual rows\x00DELETE FROM t3 WHERE a=?1; -- %d times\x00Refill two %d-row tables using REPLACE\x00REPLACE INTO z2(a,b,c) SELECT a,b,c FROM z1\x00REPLACE INTO t3(a,b,c) SELECT a,b,c FROM z1\x00Refill a %d-row table using (b&1)==(a&1)\x00INSERT INTO z2(a,b,c)\n SELECT a,b,c FROM z1  WHERE (b&1)==(a&1);\x00INSERT INTO z2(a,b,c)\n SELECT a,b,c FROM z1  WHERE (b&1)<>(a&1);\x00%d four-ways joins\x00SELECT z1.c FROM z1, z2, t3, t4\n WHERE t4.a BETWEEN ?1 AND ?2\n   AND t3.a=t4.b\n   AND z2.a=t3.b\n   AND z1.c=z2.c;\x00subquery in result set\x00SELECT sum(a), max(c),\n       avg((SELECT a FROM z2 WHERE 5+z2.b=z1.b) AND rowid<?1), max(c)\n FROM z1 WHERE rowid<?1;\x00%d REPLACE ops on an IPK\x00CREATE%s TABLE t5(a INTEGER PRIMARY KEY, b %s);\x00REPLACE INTO t5 VALUES(?1,?2); --  %d times\x00%d SELECTS on an IPK\x00SELECT b FROM t5 WHERE a=?1; --  %d times\x00%d REPLACE on TEXT PK\x00WITHOUT ROWID\x00CREATE%s TABLE t6(a TEXT PRIMARY KEY, b %s)%s;\x00REPLACE INTO t6 VALUES(?1,?2); --  %d times\x00%d SELECTS on a TEXT PK\x00SELECT b FROM t6 WHERE a=?1; --  %d times\x00%d SELECT DISTINCT\x00SELECT DISTINCT b FROM t5;\x00SELECT DISTINCT b FROM t6;\x00PRAGMA integrity_check\x00ANALYZE\x00534...9..67.195....98....6.8...6...34..8.3..1....2...6.6....28....419..5...28..79\x0053....9..6..195....98....6.8...6...34..8.3..1....2...6.6....28....419..5....8..79\x0053.......6..195....98....6.8...6...34..8.3..1....2...6.6....28....419..5....8..79\x00Sudoku with recursive 'digits'\x00WITH RECURSIVE\n  input(sud) AS (VALUES(?1)),\n  digits(z,lp) AS (\n    VALUES('1', 1)\n    UNION ALL\n    SELECT CAST(lp+1 AS TEXT), lp+1 FROM digits WHERE lp<9\n  ),\n  x(s, ind) AS (\n    SELECT sud, instr(sud, '.') FROM input\n    UNION ALL\n    SELECT\n      substr(s, 1, ind-1) || z || substr(s, ind+1),\n      instr( substr(s, 1, ind-1) || z || substr(s, ind+1), '.' )\n     FROM x, digits AS z\n    WHERE ind>0\n      AND NOT EXISTS (\n            SELECT 1\n              FROM digits AS lp\n             WHERE z.z = substr(s, ((ind-1)/9)*9 + lp, 1)\n                OR z.z = substr(s, ((ind-1)%%9) + (lp-1)*9 + 1, 1)\n                OR z.z = substr(s, (((ind-1)/3) %% 3) * 3\n                        + ((ind-1)/27) * 27 + lp\n                        + ((lp-1) / 3) * 6, 1)\n         )\n  )\nSELECT s FROM x WHERE ind=0;\x00Sudoku with VALUES 'digits'\x00WITH RECURSIVE\n  input(sud) AS (VALUES(?1)),\n  digits(z,lp) AS (VALUES('1',1),('2',2),('3',3),('4',4),('5',5),\n                         ('6',6),('7',7),('8',8),('9',9)),\n  x(s, ind) AS (\n    SELECT sud, instr(sud, '.') FROM input\n    UNION ALL\n    SELECT\n      substr(s, 1, ind-1) || z || substr(s, ind+1),\n      instr( substr(s, 1, ind-1) || z || substr(s, ind+1), '.' )\n     FROM x, digits AS z\n    WHERE ind>0\n      AND NOT EXISTS (\n            SELECT 1\n              FROM digits AS lp\n             WHERE z.z = substr(s, ((ind-1)/9)*9 + lp, 1)\n                OR z.z = substr(s, ((ind-1)%%9) + (lp-1)*9 + 1, 1)\n                OR z.z = substr(s, (((ind-1)/3) %% 3) * 3\n                        + ((ind-1)/27) * 27 + lp\n                        + ((lp-1) / 3) * 6, 1)\n         )\n  )\nSELECT s FROM x WHERE ind=0;\x00Mandelbrot Set with spacing=%f\x00WITH RECURSIVE \n  xaxis(x) AS (VALUES(-2.0) UNION ALL SELECT x+?1 FROM xaxis WHERE x<1.2),\n  yaxis(y) AS (VALUES(-1.0) UNION ALL SELECT y+?2 FROM yaxis WHERE y<1.0),\n  m(iter, cx, cy, x, y) AS (\n    SELECT 0, x, y, 0.0, 0.0 FROM xaxis, yaxis\n    UNION ALL\n    SELECT iter+1, cx, cy, x*x-y*y + cx, 2.0*x*y + cy FROM m \n     WHERE (x*x + y*y) < 4.0 AND iter<28\n  ),\n  m2(iter, cx, cy) AS (\n    SELECT max(iter), cx, cy FROM m GROUP BY cx, cy\n  ),\n  a(t) AS (\n    SELECT group_concat( substr(' .+*#', 1+min(iter/7,4), 1), '') \n    FROM m2 GROUP BY cy\n  )\nSELECT group_concat(rtrim(t),x'0a') FROM a;\x00EXCEPT operator on %d-element tables\x00WITH RECURSIVE \n  z1(x) AS (VALUES(2) UNION ALL SELECT x+2 FROM z1 WHERE x<%d),\n  z2(y) AS (VALUES(3) UNION ALL SELECT y+3 FROM z2 WHERE y<%d)\nSELECT count(x), avg(x) FROM (\n  SELECT x FROM z1 EXCEPT SELECT y FROM z2 ORDER BY 1\n);\x00%d.%de%d\x00Fill a table with %d FP values\x00CREATE%s TABLE z1(a REAL %s, b REAL %s);\x00INSERT INTO z1 VALUES(?1,?2); -- %d times\x00%d range queries\x00SELECT sum(b) FROM z1 WHERE a BETWEEN ?1 AND ?2\x00CREATE INDEX three times\x00CREATE INDEX t1a ON z1(a);\x00CREATE INDEX t1b ON z1(b);\x00CREATE INDEX t1ab ON z1(a,b);\x00%d indexed range queries\x00%d calls to round()\x00SELECT sum(round(a,2)+round(b,4)) FROM z1;\x00%d printf() calls\x00WITH c(fmt) AS (VALUES('%%g'),('%%e'),('%%!g'),('%%.20f'))SELECT sum(printf(fmt,a)) FROM z1, c\x00Create a fact table with %d entries\x00CREATE TABLE facttab( attr01 INT, attr02 INT, attr03 INT, data01 TEXT, attr04 INT, attr05 INT, attr06 INT, attr07 INT, attr08 INT, factid INTEGER PRIMARY KEY, data02 TEXT);\x00WITH RECURSIVE counter(nnn) AS(VALUES(1) UNION ALL SELECT nnn+1 FROM counter WHERE nnn<%d)INSERT INTO facttab(attr01,attr02,attr03,attr04,attr05,attr06,attr07,attr08,data01,data02)SELECT random()%%12, random()%%13, random()%%14, random()%%15,random()%%16, random()%%17, random()%%18, random()%%19,concat('data-',nnn), format('%%x',random()) FROM counter;\x00Create indexes on all attributes columns\x00CREATE INDEX fact_attr%02d ON facttab(attr%02d)\x00Create dimension tables\x00CREATE TABLE dimension%02d(beta%02d INT, content%02d TEXT, rate%02d REAL)\x00WITH RECURSIVE ctr(nn) AS (VALUES(1) UNION ALL SELECT nn+1 FROM ctr WHERE nn<%d) INSERT INTO dimension%02d   SELECT nn%%(%d), concat('content-%02d-',nn), (random()%%10000)*0.125 FROM ctr;\x00CREATE INDEX dim%02d ON dimension%02d(beta%02d);\x00CREATE INDEX dim%02d ON dimension%02d(beta%02d,content%02d);\x00Star query over the entire fact table\x00SELECT count(*), max(content04), min(content03), sum(rate04), avg(rate05) FROM facttab, dimension01, dimension02, dimension03, dimension04, dimension05, dimension06, dimension07, dimension08 WHERE attr01=beta01 AND attr02=beta02 AND attr03=beta03 AND attr04=beta04 AND attr05=beta05 AND attr06=beta06 AND attr07=beta07 AND attr08=beta08;\x00Star query with LEFT JOINs\x00SELECT count(*), max(content04), min(content03), sum(rate04), avg(rate05) FROM facttab LEFT JOIN dimension01 ON attr01=beta01 LEFT JOIN dimension02 ON attr02=beta02 JOIN dimension03 ON attr03=beta03 JOIN dimension04 ON attr04=beta04 JOIN dimension05 ON attr05=beta05 LEFT JOIN dimension06 ON attr06=beta06 JOIN dimension07 ON attr07=beta07 JOIN dimension08 ON attr08=beta08 WHERE facttab.data01 LIKE 'data-9%%';\x00Generate a Fossil-like database schema\x00BEGIN;CREATE TABLE blob(\n  rid INTEGER PRIMARY KEY,\n  rcvid INTEGER,\n  size INTEGER,\n  uuid TEXT UNIQUE NOT NULL,\n  content BLOB,\n  CHECK( length(uuid)>=40 AND rid>0 )\n);\nCREATE TABLE delta(\n  rid INTEGER PRIMARY KEY,\n  srcid INTEGER NOT NULL REFERENCES blob\n);\nCREATE TABLE rcvfrom(\n  rcvid INTEGER PRIMARY KEY,\n  uid INTEGER REFERENCES user,\n  mtime DATETIME,\n  nonce TEXT UNIQUE,\n  ipaddr TEXT\n);\nCREATE TABLE private(rid INTEGER PRIMARY KEY);\nCREATE TABLE accesslog(\n  uname TEXT,\n  ipaddr TEXT,\n  success BOOLEAN,\n  mtime TIMESTAMP\n);\nCREATE TABLE user(\n  uid INTEGER PRIMARY KEY,\n  login TEXT UNIQUE,\n  pw TEXT,\n  cap TEXT,\n  cookie TEXT,\n  ipaddr TEXT,\n  cexpire DATETIME,\n  info TEXT,\n  mtime DATE,\n  photo BLOB\n, jx TEXT DEFAULT '{}');\nCREATE TABLE reportfmt(\n   rn INTEGER PRIMARY KEY,\n   owner TEXT,\n   title TEXT UNIQUE,\n   mtime INTEGER,\n   cols TEXT,\n   sqlcode TEXT\n, jx TEXT DEFAULT '{}');\nCREATE TABLE config(\n  name TEXT PRIMARY KEY NOT NULL,\n  value CLOB, mtime INTEGER,\n  CHECK( typeof(name)='text' AND length(name)>=1 )\n) WITHOUT ROWID;\nCREATE TABLE shun(uuid PRIMARY KEY, mtime INTEGER, scom TEXT)\n  WITHOUT ROWID;\nCREATE TABLE concealed(\n  hash TEXT PRIMARY KEY,\n  content TEXT\n, mtime INTEGER) WITHOUT ROWID;\nCREATE TABLE admin_log(\n id INTEGER PRIMARY KEY,\n time INTEGER, -- Seconds since 1970\n page TEXT,    -- path of page\n who TEXT,     -- User who made the change\n  what TEXT     -- What changed\n);\nCREATE TABLE unversioned(\n  name TEXT PRIMARY KEY,\n  rcvid INTEGER,\n  mtime DATETIME,\n  hash TEXT,\n  sz INTEGER,\n  encoding INT,\n  content BLOB\n) WITHOUT ROWID;\nCREATE TABLE subscriber(\n  subscriberId INTEGER PRIMARY KEY,\n  subscriberCode BLOB DEFAULT (randomblob(32)) UNIQUE,\n  semail TEXT UNIQUE COLLATE nocase,\n  suname TEXT,\n  sverified BOOLEAN DEFAULT true,\n  sdonotcall BOOLEAN,\n  sdigest BOOLEAN,\n  ssub TEXT,\n  sctime INTDATE,\n  mtime INTDATE,\n  smip TEXT\n, lastContact INT);\nCREATE TABLE pending_alert(\n  eventid TEXT PRIMARY KEY,\n  sentSep BOOLEAN DEFAULT false,\n  sentDigest BOOLEAN DEFAULT false\n, sentMod BOOLEAN DEFAULT false) WITHOUT ROWID;\nCREATE TABLE filename(\n  fnid INTEGER PRIMARY KEY,\n  name TEXT UNIQUE\n) STRICT;\nCREATE TABLE mlink(\n  mid INTEGER,\n  fid INTEGER,\n  pmid INTEGER,\n  pid INTEGER,\n  fnid INTEGER REFERENCES filename,\n  pfnid INTEGER,\n  mperm INTEGER,\n  isaux INT DEFAULT 0\n) STRICT;\nCREATE TABLE plink(\n  pid INTEGER REFERENCES blob,\n  cid INTEGER REFERENCES blob,\n  isprim INT,\n  mtime REAL,\n  baseid INTEGER REFERENCES blob,\n  UNIQUE(pid, cid)\n) STRICT;\nCREATE TABLE leaf(rid INTEGER PRIMARY KEY);\nCREATE TABLE event(\n  type TEXT,\n  mtime REAL,\n  objid INTEGER PRIMARY KEY,\n  tagid INTEGER,\n  uid INTEGER REFERENCES user,\n  bgcolor TEXT,\n  euser TEXT,\n  user TEXT,\n  ecomment TEXT,\n  comment TEXT,\n  brief TEXT,\n  omtime REAL\n) STRICT;\nCREATE TABLE phantom(\n  rid INTEGER PRIMARY KEY\n);\nCREATE TABLE orphan(\n  rid INTEGER PRIMARY KEY,\n  baseline INTEGER\n) STRICT;\nCREATE TABLE unclustered(\n  rid INTEGER PRIMARY KEY\n);\nCREATE TABLE unsent(\n  rid INTEGER PRIMARY KEY\n);\nCREATE TABLE tag(\n  tagid INTEGER PRIMARY KEY,\n  tagname TEXT UNIQUE\n) STRICT;\nCREATE TABLE tagxref(\n  tagid INTEGER REFERENCES tag,\n  tagtype INTEGER,\n  srcid INTEGER REFERENCES blob,\n  origid INTEGER REFERENCES blob,\n  value TEXT,\n  mtime REAL,\n  rid INTEGER REFERENCES blob,\n  UNIQUE(rid, tagid)\n) STRICT;\nCREATE TABLE backlink(\n  target TEXT,\n  srctype INT,\n  srcid INT,\n  mtime REAL,\n  UNIQUE(target, srctype, srcid)\n) STRICT;\nCREATE TABLE attachment(\n  attachid INTEGER PRIMARY KEY,\n  isLatest INT DEFAULT 0,\n  mtime REAL,\n  src TEXT,\n  target TEXT,\n  filename TEXT,\n  comment TEXT,\n  user TEXT\n) STRICT;\nCREATE TABLE cherrypick(\n  parentid INT,\n  childid INT,\n  isExclude INT DEFAULT false,\n  PRIMARY KEY(parentid, childid)\n) WITHOUT ROWID, STRICT;\nCREATE TABLE vcache(\n  vid INTEGER,         -- check-in ID\n  fname TEXT,          -- filename\n  rid INTEGER,         -- artifact ID\n  PRIMARY KEY(vid,fname)\n) WITHOUT ROWID;\nCREATE TABLE synclog(\n  sfrom TEXT,\n  sto TEXT,\n  stime INT NOT NULL,\n  stype TEXT,\n  PRIMARY KEY(sfrom,sto)\n) WITHOUT ROWID;\nCREATE TABLE chat(\n  msgid INTEGER PRIMARY KEY AUTOINCREMENT,\n  mtime JULIANDAY,\n  lmtime TEXT,\n  xfrom TEXT,\n  xmsg  TEXT,\n  fname TEXT,\n  fmime TEXT,\n  mdel INT,\n  file  BLOB\n);\nCREATE TABLE ftsdocs(\n  rowid INTEGER PRIMARY KEY,\n  type CHAR(1),\n  rid INTEGER,\n  name TEXT,\n  idxed BOOLEAN,\n  label TEXT,\n  url TEXT,\n  mtime DATE,\n  bx TEXT,\n  UNIQUE(type,rid)\n);\nCREATE TABLE ticket(\n  -- Do not change any column that begins with tkt_\n  tkt_id INTEGER PRIMARY KEY,\n  tkt_uuid TEXT UNIQUE,\n  tkt_mtime DATE,\n  tkt_ctime DATE,\n  -- Add as many fields as required below this line\n  type TEXT,\n  status TEXT,\n  subsystem TEXT,\n  priority TEXT,\n  severity TEXT,\n  foundin TEXT,\n  private_contact TEXT,\n  resolution TEXT,\n  title TEXT,\n  comment TEXT\n);\nCREATE TABLE ticketchng(\n  -- Do not change any column that begins with tkt_\n  tkt_id INTEGER REFERENCES ticket,\n  tkt_rid INTEGER REFERENCES blob,\n  tkt_mtime DATE,\n  tkt_user TEXT,\n  -- Add as many fields as required below this line\n  login TEXT,\n  username TEXT,\n  mimetype TEXT,\n  icomment TEXT\n);\nCREATE TABLE forumpost(\n  fpid INTEGER PRIMARY KEY,\n  froot INT,\n  fprev INT,\n  firt INT,\n  fmtime REAL\n);\nCREATE INDEX delta_i1 ON delta(srcid);\nCREATE INDEX blob_rcvid ON blob(rcvid);\nCREATE INDEX subscriberUname\n  ON subscriber(suname) WHERE suname IS NOT NULL;\nCREATE INDEX mlink_i1 ON mlink(mid);\nCREATE INDEX mlink_i2 ON mlink(fnid);\nCREATE INDEX mlink_i3 ON mlink(fid);\nCREATE INDEX mlink_i4 ON mlink(pid);\nCREATE INDEX plink_i2 ON plink(cid,pid);\nCREATE INDEX event_i1 ON event(mtime);\nCREATE INDEX orphan_baseline ON orphan(baseline);\nCREATE INDEX tagxref_i1 ON tagxref(tagid, mtime);\nCREATE INDEX backlink_src ON backlink(srcid, srctype);\nCREATE INDEX attachment_idx1 ON attachment(target, filename, mtime);\nCREATE INDEX attachment_idx2 ON attachment(src);\nCREATE INDEX cherrypick_cid ON cherrypick(childid);\nCREATE INDEX ftsdocIdxed ON ftsdocs(type,rid,name) WHERE idxed==0;\nCREATE INDEX ftsdocName ON ftsdocs(name) WHERE type='w';\nCREATE INDEX ticketchng_idx1 ON ticketchng(tkt_id, tkt_mtime);\nCREATE INDEX forumthread ON forumpost(froot,fmtime);\nCREATE VIEW artifact(rid,rcvid,size,atype,srcid,hash,content) AS\n  SELECT blob.rid,rcvid,size,1,srcid,uuid,content\n    FROM blob LEFT JOIN delta ON (blob.rid=delta.rid);\nCREATE VIEW ftscontent AS\n  SELECT rowid, type, rid, name, idxed, label, url, mtime,\n         title(type,rid,name) AS 'title', body(type,rid,name) AS 'body'\n    FROM ftsdocs;\n\x00ENABLE_FTS5\x00CREATE VIRTUAL TABLE ftsidx\n  USING fts5(content=\"ftscontent\", title, body);\nCREATE VIRTUAL TABLE chatfts1 USING fts5(\n  xmsg, content=chat, content_rowid=msgid,tokenize=porter);\n\x00CREATE TABLE ftsidx_data(id INTEGER PRIMARY KEY, block BLOB);\nCREATE TABLE ftsidx_idx(segid, term, pgno, PRIMARY KEY(segid, term))\n  WITHOUT ROWID;\nCREATE TABLE ftsidx_docsize(id INTEGER PRIMARY KEY, sz BLOB);\nCREATE TABLE ftsidx_config(k PRIMARY KEY, v) WITHOUT ROWID;\nCREATE TABLE chatfts1_data(id INTEGER PRIMARY KEY, block BLOB);\nCREATE TABLE chatfts1_idx(segid, term, pgno, PRIMARY KEY(segid, term))\n  WITHOUT ROWID;\nCREATE TABLE chatfts1_docsize(id INTEGER PRIMARY KEY, sz BLOB);\nCREATE TABLE chatfts1_config(k PRIMARY KEY, v) WITHOUT ROWID;\n\x00ANALYZE sqlite_schema;\nINSERT INTO sqlite_stat1(tbl,idx,stat) VALUES\n  ('ftsidx_config','ftsidx_config','1 1'),\n  ('ftsidx_idx','ftsidx_idx','4215 401 1'),\n  ('user','sqlite_autoindex_user_1','25 1'),\n  ('phantom',NULL,'26'),\n  ('reportfmt','sqlite_autoindex_reportfmt_1','9 1'),\n  ('rcvfrom','sqlite_autoindex_rcvfrom_1','18445 401'),\n  ('private',NULL,'99'),\n  ('mlink','mlink_i4','116678 401'),\n  ('mlink','mlink_i3','121212 2'),\n  ('mlink','mlink_i2','106372 401'),\n  ('mlink','mlink_i1','99298 5'),\n  ('ftsidx_data',NULL,'3795'),\n  ('leaf',NULL,'1559'),\n  ('delta','delta_i1','66340 1'),\n  ('unversioned','unversioned','3 1'),\n  ('pending_alert','pending_alert','3 1'),\n  ('cherrypick','cherrypick_cid','680 2'),\n  ('cherrypick','cherrypick','628 1 1'),\n  ('config','config','128 1'),\n  ('ftsidx_docsize',NULL,'33848'),\n  ('event','event_i1','36096 1'),\n  ('plink','plink_i2','38236 1 1'),\n  ('plink','sqlite_autoindex_plink_1','38357 1 1'),\n  ('shun','shun','10 1'),\n  ('concealed','concealed','110 1'),\n  ('vcache','vcache','1888 401 1'),\n  ('ftsdocs','ftsdocName','19 1'),\n  ('ftsdocs','ftsdocIdxed','168 84 1 1'),\n  ('ftsdocs','sqlite_autoindex_ftsdocs_1','37312 401 1'),\n  ('subscriber','subscriberUname','5 1'),\n  ('subscriber','sqlite_autoindex_subscriber_2','37 1'),\n  ('subscriber','sqlite_autoindex_subscriber_1','37 1'),\n  ('tag','sqlite_autoindex_tag_1','2990 1'),\n  ('filename','sqlite_autoindex_filename_1','3168 1'),\n  ('chat',NULL,'56124'),\n  ('tagxref','tagxref_i1','40992 401 2'),\n  ('tagxref','sqlite_autoindex_tagxref_1','79233 3 1'),\n  ('attachment','attachment_idx2','11 1'),\n  ('attachment','attachment_idx1','11 2 2 1'),\n  ('blob','blob_rcvid','128240 201'),\n  ('blob','sqlite_autoindex_blob_1','126480 1'),\n  ('synclog','synclog','12 3 1'),\n  ('backlink','backlink_src','2160 2 2'),\n  ('backlink','sqlite_autoindex_backlink_1','2340 2 2 1'),\n  ('accesslog',NULL,'38'),\n  ('chatfts1_config','chatfts1_config','1 1'),\n  ('chatfts1_idx','chatfts1_idx','688 230 1'),\n  ('ticket','sqlite_autoindex_ticket_1','794 1'),\n  ('ticketchng','ticketchng_idx1','2089 3 1'),\n  ('forumpost','forumthread','4 4 1'),\n  ('unclustered',NULL,'12');\nCOMMIT;\x00Open and use the database %d times\x00Cannot open database file: %s\n\x00SELECT name FROM pragma_table_list /*scan*/ WHERE schema='repository' AND type IN ('table','virtual') AND name NOT IN ('admin_log', 'blob','delta','rcvfrom','user','alias','config','shun','private','reportfmt','concealed','accesslog','modreq','purgeevent','purgeitem','unversioned','subscriber','pending_alert','chat') AND name NOT GLOB 'sqlite_*' AND name NOT GLOB 'fx_*';SELECT 1 FROM pragma_table_xinfo('ticket') WHERE name = 'mimetype';\x00SELECT name, value, unixepoch()/86400-value, date(value*86400,'unixepoch') FROM config WHERE name in ('email-renew-warning','email-renew-cutoff');SELECT count(*) FROM pending_alert WHERE NOT sentDigest;\x00WITH priors(rid,who) AS (  SELECT firt, coalesce(euser,user)    FROM forumpost LEFT JOIN event ON fpid=objid   WHERE fpid=12345  UNION ALL  SELECT firt, coalesce(euser,user)    FROM priors, forumpost LEFT JOIN event ON fpid=objid   WHERE fpid=rid)SELECT ','||group_concat(DISTINCT 'u'||who)||','||group_concat(rid) FROM priors;\x00CREATE TEMP TABLE IF NOT EXISTS ok(rid INTEGER PRIMARY KEY);\n\x00WITH RECURSIVE\n  parent(pid,cid,isCP) AS (\n    SELECT plink.pid, plink.cid, 0 AS xisCP FROM plink\n    UNION ALL\n    SELECT parentid, childid, 1 FROM cherrypick WHERE NOT isExclude\n  ),\n  ancestor(rid, mtime, isCP) AS (\n    SELECT 123, mtime, 0 FROM event WHERE objid=$object\n    UNION\n    SELECT parent.pid, event.mtime, parent.isCP\n      FROM ancestor, parent, event\n     WHERE parent.cid=ancestor.rid\n       AND event.objid=parent.pid\n       AND NOT ancestor.isCP\n       AND (event.mtime>=$date OR parent.pid=$pid)\n     ORDER BY mtime DESC LIMIT 10\n  )\n  INSERT OR IGNORE INTO ok SELECT rid FROM ancestor;\x00%d INSERTs into an r-tree\x00CREATE VIRTUAL TABLE rt1 USING rtree(id,x0,x1,y0,y1,z0,z1)\x00INSERT INTO rt1(id,x0,x1,y0,y1,z0,z1)VALUES(?1,?2,?3,?4,?5,?6,?7)\x00Copy from rtree to a regular table\x00CREATE TABLE z1(id INTEGER PRIMARY KEY,x0,x1,y0,y1,z0,z1)\x00INSERT INTO z1 SELECT * FROM rt1\x00%d one-dimensional intersect slice queries\x00SELECT count(*) FROM rt1 WHERE x0>=?1 AND x1<=?2\x00Verify result from 1-D intersect slice queries\x00SELECT count(*) FROM z1 WHERE x0>=?1 AND x1<=?2\x00Count disagree step %d: %d..%d.  %d vs %d\x00%d one-dimensional overlap slice queries\x00SELECT count(*) FROM rt1 WHERE y1>=?1 AND y0<=?2\x00Verify result from 1-D overlap slice queries\x00SELECT count(*) FROM z1 WHERE y1>=?1 AND y0<=?2\x00%d custom geometry callback queries\x00xslice\x00SELECT count(*) FROM rt1 WHERE id MATCH xslice(?1,?2)\x00%d three-dimensional intersect box queries\x00SELECT count(*) FROM rt1 WHERE x1>=?1 AND x0<=?2 AND y1>=?1 AND y0<=?2 AND z1>=?1 AND z0<=?2\x00%d rowid queries\x00SELECT * FROM rt1 WHERE id=?1\x00%d UPDATEs using rowid\x00UPDATE rt1 SET x0=x0+100, x1=x1+100 WHERE id=?1\x00%d UPDATEs using one-dimensional overlap\x00UPDATE rt1 SET x0=x0-100, x1=x1-100 WHERE y1>=?1 AND y0<=?1+5\x00%d DELETEs using rowid\x00DELETE FROM rt1 WHERE id=?1\x00%d DELETEs using one-dimensional overlap\x00DELETE FROM rt1 WHERE y1>=?1 AND y0<=?1+5\x00Restore deleted entries using INSERT OR IGNORE\x00INSERT OR IGNORE INTO rt1 SELECT * FROM z1\x00Fill %d rows\x00BEGIN;CREATE TABLE ZLOOKSLIKECOREDATA (  ZPK INTEGER PRIMARY KEY,  ZTERMFITTINGHOUSINGCOMMAND INTEGER,  ZBRIEFGOBYDODGERHEIGHT BLOB,  ZCAPABLETRIPDOORALMOND BLOB,  ZDEPOSITPAIRCOLLEGECOMET INTEGER,  ZFRAMEENTERSIMPLEMOUTH INTEGER,  ZHOPEFULGATEHOLECHALK INTEGER,  ZSLEEPYUSERGRANDBOWL TIMESTAMP,  ZDEWPEACHCAREERCELERY INTEGER,  ZHANGERLITHIUMDINNERMEET VARCHAR,  ZCLUBRELEASELIZARDADVICE VARCHAR,  ZCHARGECLICKHUMANEHIRE INTEGER,  ZFINGERDUEPIZZAOPTION TIMESTAMP,  ZFLYINGDOCTORTABLEMELODY BLOB,  ZLONGFINLEAVEIMAGEOIL TIMESTAMP,  ZFAMILYVISUALOWNERMATTER BLOB,  ZGOLDYOUNGINITIALNOSE FLOAT,  ZCAUSESALAMITERMCYAN BLOB,  ZSPREADMOTORBISCUITBACON FLOAT,  ZGIFTICEFISHGLUEHAIR INTEGER,  ZNOTICEPEARPOLICYJUICE TIMESTAMP,  ZBANKBUFFALORECOVERORBIT TIMESTAMP,  ZLONGDIETESSAYNATURE FLOAT,  ZACTIONRANGEELEGANTNEUTRON BLOB,  ZCADETBRIGHTPLANETBANK TIMESTAMP,  ZAIRFORGIVEHEADFROG BLOB,  ZSHARKJUSTFRUITMOVIE VARCHAR,  ZFARMERMORNINGMIRRORCONCERN BLOB,  ZWOODPOETRYCOBBLERBENCH VARCHAR,  ZHAFNIUMSCRIPTSALADMOTOR INTEGER,  ZPROBLEMCLUBPOPOVERJELLY FLOAT,  ZEIGHTLEADERWORKERMOST TIMESTAMP,  ZGLASSRESERVEBARIUMMEAL BLOB,  ZCLAMBITARUGULAFAJITA BLOB,  ZDECADEJOYOUSWAVEHABIT FLOAT,  ZCOMPANYSUMMERFIBERELF INTEGER,  ZTREATTESTQUILLCHARGE TIMESTAMP,  ZBROWBALANCEKEYCHOWDER FLOAT,  ZPEACHCOPPERDINNERLAKE FLOAT,  ZDRYWALLBEYONDBROWNBOWL VARCHAR,  ZBELLYCRASHITEMLACK BLOB,  ZTENNISCYCLEBILLOFFICER INTEGER,  ZMALLEQUIPTHANKSGLUE FLOAT,  ZMISSREPLYHUMANLIVING INTEGER,  ZKIWIVISUALPRIDEAPPLE VARCHAR,  ZWISHHITSKINMOTOR BLOB,  ZCALMRACCOONPROGRAMDEBIT VARCHAR,  ZSHINYASSISTLIVINGCRAB VARCHAR,  ZRESOLVEWRISTWRAPAPPLE VARCHAR,  ZAPPEALSIMPLESECONDHOUSING BLOB,  ZCORNERANCHORTAPEDIVER TIMESTAMP,  ZMEMORYREQUESTSOURCEBIG VARCHAR,  ZTRYFACTKEEPMILK TIMESTAMP,  ZDIVERPAINTLEATHEREASY INTEGER,  ZSORTMISTYQUOTECABBAGE BLOB,  ZTUNEGASBUFFALOCAPITAL BLOB,  ZFILLSTOPLAWJOYFUL FLOAT,  ZSTEELCAREFULPLATENUMBER FLOAT,  ZGIVEVIVIDDIVINEMEANING INTEGER,  ZTREATPACKFUTURECONVERT VARCHAR,  ZCALMLYGEMFINISHEFFECT INTEGER,  ZCABBAGESOCKEASEMINUTE BLOB,  ZPLANETFAMILYPUREMEMORY TIMESTAMP,  ZMERRYCRACKTRAINLEADER BLOB,  ZMINORWAYPAPERCLASSY TIMESTAMP,  ZEAGLELINEMINEMAIL VARCHAR,  ZRESORTYARDGREENLET TIMESTAMP,  ZYARDOREGANOVIVIDJEWEL TIMESTAMP,  ZPURECAKEVIVIDNEATLY FLOAT,  ZASKCONTACTMONITORFUN TIMESTAMP,  ZMOVEWHOGAMMAINCH VARCHAR,  ZLETTUCEBIRDMEETDEBATE TIMESTAMP,  ZGENENATURALHEARINGKITE VARCHAR,  ZMUFFINDRYERDRAWFORTUNE FLOAT,  ZGRAYSURVEYWIRELOVE FLOAT,  ZPLIERSPRINTASKOREGANO INTEGER,  ZTRAVELDRIVERCONTESTLILY INTEGER,  ZHUMORSPICESANDKIDNEY TIMESTAMP,  ZARSENICSAMPLEWAITMUON INTEGER,  ZLACEADDRESSGROUNDCAREFUL FLOAT,  ZBAMBOOMESSWASABIEVENING BLOB,  ZONERELEASEAVERAGENURSE INTEGER,  ZRADIANTWHENTRYCARD TIMESTAMP,  ZREWARDINSIDEMANGOINTENSE FLOAT,  ZNEATSTEWPARTIRON TIMESTAMP,  ZOUTSIDEPEAHENCOUNTICE TIMESTAMP,  ZCREAMEVENINGLIPBRANCH FLOAT,  ZWHALEMATHAVOCADOCOPPER FLOAT,  ZLIFEUSELEAFYBELL FLOAT,  ZWEALTHLINENGLEEFULDAY VARCHAR,  ZFACEINVITETALKGOLD BLOB,  ZWESTAMOUNTAFFECTHEARING INTEGER,  ZDELAYOUTCOMEHORNAGENCY INTEGER,  ZBIGTHINKCONVERTECONOMY BLOB,  ZBASEGOUDAREGULARFORGIVE TIMESTAMP,  ZPATTERNCLORINEGRANDCOLBY TIMESTAMP,  ZCYANBASEFEEDADROIT INTEGER,  ZCARRYFLOORMINNOWDRAGON TIMESTAMP,  ZIMAGEPENCILOTHERBOTTOM FLOAT,  ZXENONFLIGHTPALEAPPLE TIMESTAMP,  ZHERRINGJOKEFEATUREHOPEFUL FLOAT,  ZCAPYEARLYRIVETBRUSH FLOAT,  ZAGEREEDFROGBASKET VARCHAR,  ZUSUALBODYHALIBUTDIAMOND VARCHAR,  ZFOOTTAPWORDENTRY VARCHAR,  ZDISHKEEPBLESTMONITOR FLOAT,  ZBROADABLESOLIDCASUAL INTEGER,  ZSQUAREGLEEFULCHILDLIGHT INTEGER,  ZHOLIDAYHEADPONYDETAIL INTEGER,  ZGENERALRESORTSKYOPEN TIMESTAMP,  ZGLADSPRAYKIDNEYGUPPY VARCHAR,  ZSWIMHEAVYMENTIONKIND BLOB,  ZMESSYSULFURDREAMFESTIVE BLOB,  ZSKYSKYCLASSICBRIEF VARCHAR,  ZDILLASKHOKILEMON FLOAT,  ZJUNIORSHOWPRESSNOVA FLOAT,  ZSIZETOEAWARDFRESH TIMESTAMP,  ZKEYFAILAPRICOTMETAL VARCHAR,  ZHANDYREPAIRPROTONAIRPORT VARCHAR,  ZPOSTPROTEINHANDLEACTOR BLOB);\x00INSERT INTO ZLOOKSLIKECOREDATA(ZPK,ZAIRFORGIVEHEADFROG,ZGIFTICEFISHGLUEHAIR,ZDELAYOUTCOMEHORNAGENCY,ZSLEEPYUSERGRANDBOWL,ZGLASSRESERVEBARIUMMEAL,ZBRIEFGOBYDODGERHEIGHT,ZBAMBOOMESSWASABIEVENING,ZFARMERMORNINGMIRRORCONCERN,ZTREATPACKFUTURECONVERT,ZCAUSESALAMITERMCYAN,ZCALMRACCOONPROGRAMDEBIT,ZHOLIDAYHEADPONYDETAIL,ZWOODPOETRYCOBBLERBENCH,ZHAFNIUMSCRIPTSALADMOTOR,ZUSUALBODYHALIBUTDIAMOND,ZOUTSIDEPEAHENCOUNTICE,ZDIVERPAINTLEATHEREASY,ZWESTAMOUNTAFFECTHEARING,ZSIZETOEAWARDFRESH,ZDEWPEACHCAREERCELERY,ZSTEELCAREFULPLATENUMBER,ZCYANBASEFEEDADROIT,ZCALMLYGEMFINISHEFFECT,ZHANDYREPAIRPROTONAIRPORT,ZGENENATURALHEARINGKITE,ZBROADABLESOLIDCASUAL,ZPOSTPROTEINHANDLEACTOR,ZLACEADDRESSGROUNDCAREFUL,ZIMAGEPENCILOTHERBOTTOM,ZPROBLEMCLUBPOPOVERJELLY,ZPATTERNCLORINEGRANDCOLBY,ZNEATSTEWPARTIRON,ZAPPEALSIMPLESECONDHOUSING,ZMOVEWHOGAMMAINCH,ZTENNISCYCLEBILLOFFICER,ZSHARKJUSTFRUITMOVIE,ZKEYFAILAPRICOTMETAL,ZCOMPANYSUMMERFIBERELF,ZTERMFITTINGHOUSINGCOMMAND,ZRESORTYARDGREENLET,ZCABBAGESOCKEASEMINUTE,ZSQUAREGLEEFULCHILDLIGHT,ZONERELEASEAVERAGENURSE,ZBIGTHINKCONVERTECONOMY,ZPLIERSPRINTASKOREGANO,ZDECADEJOYOUSWAVEHABIT,ZDRYWALLBEYONDBROWNBOWL,ZCLUBRELEASELIZARDADVICE,ZWHALEMATHAVOCADOCOPPER,ZBELLYCRASHITEMLACK,ZLETTUCEBIRDMEETDEBATE,ZCAPABLETRIPDOORALMOND,ZRADIANTWHENTRYCARD,ZCAPYEARLYRIVETBRUSH,ZAGEREEDFROGBASKET,ZSWIMHEAVYMENTIONKIND,ZTRAVELDRIVERCONTESTLILY,ZGLADSPRAYKIDNEYGUPPY,ZBANKBUFFALORECOVERORBIT,ZFINGERDUEPIZZAOPTION,ZCLAMBITARUGULAFAJITA,ZLONGFINLEAVEIMAGEOIL,ZLONGDIETESSAYNATURE,ZJUNIORSHOWPRESSNOVA,ZHOPEFULGATEHOLECHALK,ZDEPOSITPAIRCOLLEGECOMET,ZWEALTHLINENGLEEFULDAY,ZFILLSTOPLAWJOYFUL,ZTUNEGASBUFFALOCAPITAL,ZGRAYSURVEYWIRELOVE,ZCORNERANCHORTAPEDIVER,ZREWARDINSIDEMANGOINTENSE,ZCADETBRIGHTPLANETBANK,ZPLANETFAMILYPUREMEMORY,ZTREATTESTQUILLCHARGE,ZCREAMEVENINGLIPBRANCH,ZSKYSKYCLASSICBRIEF,ZARSENICSAMPLEWAITMUON,ZBROWBALANCEKEYCHOWDER,ZFLYINGDOCTORTABLEMELODY,ZHANGERLITHIUMDINNERMEET,ZNOTICEPEARPOLICYJUICE,ZSHINYASSISTLIVINGCRAB,ZLIFEUSELEAFYBELL,ZFACEINVITETALKGOLD,ZGENERALRESORTSKYOPEN,ZPURECAKEVIVIDNEATLY,ZKIWIVISUALPRIDEAPPLE,ZMESSYSULFURDREAMFESTIVE,ZCHARGECLICKHUMANEHIRE,ZHERRINGJOKEFEATUREHOPEFUL,ZYARDOREGANOVIVIDJEWEL,ZFOOTTAPWORDENTRY,ZWISHHITSKINMOTOR,ZBASEGOUDAREGULARFORGIVE,ZMUFFINDRYERDRAWFORTUNE,ZACTIONRANGEELEGANTNEUTRON,ZTRYFACTKEEPMILK,ZPEACHCOPPERDINNERLAKE,ZFRAMEENTERSIMPLEMOUTH,ZMERRYCRACKTRAINLEADER,ZMEMORYREQUESTSOURCEBIG,ZCARRYFLOORMINNOWDRAGON,ZMINORWAYPAPERCLASSY,ZDILLASKHOKILEMON,ZRESOLVEWRISTWRAPAPPLE,ZASKCONTACTMONITORFUN,ZGIVEVIVIDDIVINEMEANING,ZEIGHTLEADERWORKERMOST,ZMISSREPLYHUMANLIVING,ZXENONFLIGHTPALEAPPLE,ZSORTMISTYQUOTECABBAGE,ZEAGLELINEMINEMAIL,ZFAMILYVISUALOWNERMATTER,ZSPREADMOTORBISCUITBACON,ZDISHKEEPBLESTMONITOR,ZMALLEQUIPTHANKSGLUE,ZGOLDYOUNGINITIALNOSE,ZHUMORSPICESANDKIDNEY)VALUES(?1,?26,?20,?93,?8,?33,?3,?81,?28,?60,?18,?47,?109,?29,?30,?104,?86,?54,?92,?117,?9,?58,?97,?61,?119,?73,?107,?120,?80,?99,?31,?96,?85,?50,?71,?42,?27,?118,?36,?2,?67,?62,?108,?82,?94,?76,?35,?40,?11,?88,?41,?72,?4,?83,?102,?103,?112,?77,?111,?22,?13,?34,?15,?23,?116,?7,?5,?90,?57,?56,?75,?51,?84,?25,?63,?37,?87,?114,?79,?38,?14,?10,?21,?48,?89,?91,?110,?69,?45,?113,?12,?101,?68,?105,?46,?95,?74,?24,?53,?39,?6,?64,?52,?98,?65,?115,?49,?70,?59,?32,?44,?100,?55,?66,?16,?19,?106,?43,?17,?78);\x00Query %d rows by rowid\x00SELECT ZCYANBASEFEEDADROIT,ZJUNIORSHOWPRESSNOVA,ZCAUSESALAMITERMCYAN,ZHOPEFULGATEHOLECHALK,ZHUMORSPICESANDKIDNEY,ZSWIMHEAVYMENTIONKIND,ZMOVEWHOGAMMAINCH,ZAPPEALSIMPLESECONDHOUSING,ZHAFNIUMSCRIPTSALADMOTOR,ZNEATSTEWPARTIRON,ZLONGFINLEAVEIMAGEOIL,ZDEWPEACHCAREERCELERY,ZXENONFLIGHTPALEAPPLE,ZCALMRACCOONPROGRAMDEBIT,ZUSUALBODYHALIBUTDIAMOND,ZTRYFACTKEEPMILK,ZWEALTHLINENGLEEFULDAY,ZLONGDIETESSAYNATURE,ZLIFEUSELEAFYBELL,ZTREATPACKFUTURECONVERT,ZMEMORYREQUESTSOURCEBIG,ZYARDOREGANOVIVIDJEWEL,ZDEPOSITPAIRCOLLEGECOMET,ZSLEEPYUSERGRANDBOWL,ZBRIEFGOBYDODGERHEIGHT,ZCLUBRELEASELIZARDADVICE,ZCAPABLETRIPDOORALMOND,ZDRYWALLBEYONDBROWNBOWL,ZASKCONTACTMONITORFUN,ZKIWIVISUALPRIDEAPPLE,ZNOTICEPEARPOLICYJUICE,ZPEACHCOPPERDINNERLAKE,ZSTEELCAREFULPLATENUMBER,ZGLADSPRAYKIDNEYGUPPY,ZCOMPANYSUMMERFIBERELF,ZTENNISCYCLEBILLOFFICER,ZIMAGEPENCILOTHERBOTTOM,ZWESTAMOUNTAFFECTHEARING,ZDIVERPAINTLEATHEREASY,ZSKYSKYCLASSICBRIEF,ZMESSYSULFURDREAMFESTIVE,ZMERRYCRACKTRAINLEADER,ZBROADABLESOLIDCASUAL,ZGLASSRESERVEBARIUMMEAL,ZTUNEGASBUFFALOCAPITAL,ZBANKBUFFALORECOVERORBIT,ZTREATTESTQUILLCHARGE,ZBAMBOOMESSWASABIEVENING,ZREWARDINSIDEMANGOINTENSE,ZEAGLELINEMINEMAIL,ZCALMLYGEMFINISHEFFECT,ZKEYFAILAPRICOTMETAL,ZFINGERDUEPIZZAOPTION,ZCADETBRIGHTPLANETBANK,ZGOLDYOUNGINITIALNOSE,ZMISSREPLYHUMANLIVING,ZEIGHTLEADERWORKERMOST,ZFRAMEENTERSIMPLEMOUTH,ZBIGTHINKCONVERTECONOMY,ZFACEINVITETALKGOLD,ZPOSTPROTEINHANDLEACTOR,ZHERRINGJOKEFEATUREHOPEFUL,ZCABBAGESOCKEASEMINUTE,ZMUFFINDRYERDRAWFORTUNE,ZPROBLEMCLUBPOPOVERJELLY,ZGIVEVIVIDDIVINEMEANING,ZGENENATURALHEARINGKITE,ZGENERALRESORTSKYOPEN,ZLETTUCEBIRDMEETDEBATE,ZBASEGOUDAREGULARFORGIVE,ZCHARGECLICKHUMANEHIRE,ZPLANETFAMILYPUREMEMORY,ZMINORWAYPAPERCLASSY,ZCAPYEARLYRIVETBRUSH,ZSIZETOEAWARDFRESH,ZARSENICSAMPLEWAITMUON,ZSQUAREGLEEFULCHILDLIGHT,ZSHINYASSISTLIVINGCRAB,ZCORNERANCHORTAPEDIVER,ZDECADEJOYOUSWAVEHABIT,ZTRAVELDRIVERCONTESTLILY,ZFLYINGDOCTORTABLEMELODY,ZSHARKJUSTFRUITMOVIE,ZFAMILYVISUALOWNERMATTER,ZFARMERMORNINGMIRRORCONCERN,ZGIFTICEFISHGLUEHAIR,ZOUTSIDEPEAHENCOUNTICE,ZSPREADMOTORBISCUITBACON,ZWISHHITSKINMOTOR,ZHOLIDAYHEADPONYDETAIL,ZWOODPOETRYCOBBLERBENCH,ZAIRFORGIVEHEADFROG,ZBROWBALANCEKEYCHOWDER,ZDISHKEEPBLESTMONITOR,ZCLAMBITARUGULAFAJITA,ZPLIERSPRINTASKOREGANO,ZRADIANTWHENTRYCARD,ZDELAYOUTCOMEHORNAGENCY,ZPURECAKEVIVIDNEATLY,ZPATTERNCLORINEGRANDCOLBY,ZHANDYREPAIRPROTONAIRPORT,ZAGEREEDFROGBASKET,ZSORTMISTYQUOTECABBAGE,ZFOOTTAPWORDENTRY,ZRESOLVEWRISTWRAPAPPLE,ZDILLASKHOKILEMON,ZFILLSTOPLAWJOYFUL,ZACTIONRANGEELEGANTNEUTRON,ZRESORTYARDGREENLET,ZCREAMEVENINGLIPBRANCH,ZWHALEMATHAVOCADOCOPPER,ZGRAYSURVEYWIRELOVE,ZBELLYCRASHITEMLACK,ZHANGERLITHIUMDINNERMEET,ZCARRYFLOORMINNOWDRAGON,ZMALLEQUIPTHANKSGLUE,ZTERMFITTINGHOUSINGCOMMAND,ZONERELEASEAVERAGENURSE,ZLACEADDRESSGROUNDCAREFUL FROM ZLOOKSLIKECOREDATA WHERE ZPK=?1;\x00BEGIN;CREATE TABLE z1(rowid INTEGER PRIMARY KEY, i INTEGER, t TEXT);CREATE TABLE z2(rowid INTEGER PRIMARY KEY, i INTEGER, t TEXT);CREATE TABLE t3(rowid INTEGER PRIMARY KEY, i INTEGER, t TEXT);CREATE VIEW v1 AS SELECT rowid, i, t FROM z1;CREATE VIEW v2 AS SELECT rowid, i, t FROM z2;CREATE VIEW v3 AS SELECT rowid, i, t FROM t3;\x00INSERT INTO t%d VALUES(NULL,?1,?2)\x00CREATE INDEX i1 ON z1(t);CREATE INDEX i2 ON z2(t);CREATE INDEX i3 ON t3(t);COMMIT;\x00speed4p-join1\x00SELECT * FROM z1, z2, t3 WHERE z1.oid = z2.oid AND z2.oid = t3.oid\x00speed4p-join2\x00SELECT * FROM z1, z2, t3 WHERE z1.t = z2.t AND z2.t = t3.t\x00speed4p-view1\x00SELECT * FROM v%d WHERE rowid = ?\x00speed4p-table1\x00SELECT * FROM t%d WHERE rowid = ?\x00speed4p-subselect1\x00SELECT (SELECT t FROM z1 WHERE rowid = ?1),(SELECT t FROM z2 WHERE rowid = ?1),(SELECT t FROM t3 WHERE rowid = ?1)\x00speed4p-rowid-update\x00UPDATE z1 SET i=i+1 WHERE rowid=?1\x00CREATE TABLE t5(t TEXT PRIMARY KEY, i INTEGER);\x00speed4p-insert-ignore\x00INSERT OR IGNORE INTO t5 SELECT t, i FROM z1\x00CREATE TABLE log(op TEXT, r INTEGER, i INTEGER, t TEXT);CREATE TABLE t4(rowid INTEGER PRIMARY KEY, i INTEGER, t TEXT);CREATE TRIGGER t4_trigger1 AFTER INSERT ON t4 BEGIN  INSERT INTO log VALUES('INSERT INTO t4', new.rowid, new.i, new.t);END;CREATE TRIGGER t4_trigger2 AFTER UPDATE ON t4 BEGIN  INSERT INTO log VALUES('UPDATE OF t4', new.rowid, new.i, new.t);END;CREATE TRIGGER t4_trigger3 AFTER DELETE ON t4 BEGIN  INSERT INTO log VALUES('DELETE OF t4', old.rowid, old.i, old.t);END;BEGIN;\x00speed4p-trigger1\x00INSERT INTO t4 VALUES(NULL, ?1, ?2)\x00speed4p-trigger2\x00UPDATE t4 SET i = ?1, t = ?2 WHERE rowid = ?3\x00speed4p-trigger3\x00DELETE FROM t4 WHERE rowid = ?1\x00DROP TABLE t4;DROP TABLE log;VACUUM;CREATE TABLE t4(rowid INTEGER PRIMARY KEY, i INTEGER, t TEXT);BEGIN;\x00speed4p-notrigger1\x00speed4p-notrigger2\x00speed4p-notrigger3\x00%5d %5d %5d %s\n\x00table J1 is %d rows of JSONB\x00CREATE TABLE j1(x JSONB);\nWITH RECURSIVE\n  jval(n,j) AS (\n    VALUES(0,'{}'),(1,'[]'),(2,'true'),(3,'false'),(4,'null'),\n          (5,'{x:1,y:2}'),(6,'0.0'),(7,'3.14159'),(8,'-99.9'),\n          (9,'[1,2,\"\\n\\u2192\\\"\\u2190\",4]')\n  ),\n  c(x) AS (VALUES(1) UNION ALL SELECT x+1 FROM c WHERE x<26*26-1),\n  array1(y) AS MATERIALIZED (\n    SELECT jsonb_group_array(\n      jsonb_object('x',x,\n                  'y',jsonb(coalesce(j,random()%%10000)),\n                  'z',hex(randomblob(50)))\n    )\n    FROM c LEFT JOIN jval ON (x%%20)=n\n  ),\n  object1(z) AS MATERIALIZED (\n    SELECT jsonb_group_object(char(0x61+x%%26,0x61+(x/26)%%26),\n                      jsonb( coalesce(j,random()%%10000)))\n      FROM c LEFT JOIN jval ON (x%%20)=n\n  ),\n  c2(n) AS (VALUES(1) UNION ALL SELECT n+1 FROM c2 WHERE n<%d)\nINSERT INTO j1(x)\n  SELECT jsonb_object('a',n,'b',n+10000,'c',jsonb(y),'d',jsonb(z),\n                     'e',n+20000,'f',n+30000)\n    FROM array1, object1, c2;\x00table J2 is %d rows from J1 converted to text\x00CREATE TABLE j2(x JSON TEXT);\nINSERT INTO j2(x) SELECT json(x) FROM j1 LIMIT %d\x00create indexes on JSON expressions on J1\x00BEGIN;\nCREATE INDEX j1x1 ON j1(x->>'a');\nCREATE INDEX j1x2 ON j1(x->>'b');\nCREATE INDEX j1x3 ON j1(x->>'f');\nCOMMIT;\n\x00create indexes on JSON expressions on J2\x00BEGIN;\nCREATE INDEX j2x1 ON j2(x->>'a');\nCREATE INDEX j2x2 ON j2(x->>'b');\nCREATE INDEX j2x3 ON j2(x->>'f');\nCOMMIT;\n\x00queries against J1\x00WITH c(n) AS (VALUES(0) UNION ALL SELECT n+1 FROM c WHERE n<7)\n  SELECT sum(x->>format('$.c[%%d].x',n)) FROM c, j1;\nWITH c(n) AS (VALUES(1) UNION ALL SELECT n+1 FROM c WHERE n<5)\n  SELECT sum(x->>format('$.\"c\"[#-%%d].y',n)) FROM c, j1;\nSELECT sum(x->>'$.d.ez' + x->>'$.d.\"xz\"' + x->>'a' + x->>'$.c[10].y') FROM j1;\nSELECT x->>'$.d.tz[2]', x->'$.d.tz' FROM j1;\n\x00queries involving json_type()\x00WITH c(n) AS (VALUES(1) UNION ALL SELECT n+1 FROM c WHERE n<20)\n  SELECT json_type(x,format('$.c[#-%%d].y',n)), count(*)\n    FROM c, j1\n   WHERE j1.rowid=1\n   GROUP BY 1 ORDER BY 2;\x00json_insert()/set()/remove() on every row of J1\x00BEGIN;\nUPDATE j1 SET x=jsonb_insert(x,'$.g',(x->>'f')+1,'$.h',3.14159,'$.i','hello',\n                               '$.j',json('{x:99}'),'$.k','{y:98}');\nUPDATE j1 SET x=jsonb_set(x,'$.e',(x->>'f')-1);\nUPDATE j1 SET x=jsonb_remove(x,'$.d');\nCOMMIT;\n\x00json_insert()/set()/remove() on every row of J2\x00BEGIN;\nUPDATE j2 SET x=json_insert(x,'$.g',(x->>'f')+1);\nUPDATE j2 SET x=json_set(x,'$.e',(x->>'f')-1);\nUPDATE j2 SET x=json_remove(x,'$.d');\nCOMMIT;\n\x00SELECT 1, 12, 123, 1234, 12345, 123456\x00SELECT 8227256643844975616, 7932208612563860480, 2010730661871032832, 9138463067404021760, 2557616153664746496, 2557616153664746496\x00SELECT 1.0, 1.2, 1.23, 123.4, 1.2345, 1.23456\x00SELECT 8.227256643844975616, 7.932208612563860480, 2.010730661871032832, 9.138463067404021760, 2.557616153664746496, 2.557616153664746496\x00parsing %d small integers\x00parsing %d large integers\x00parsing %d small reals\x00parsing %d large reals\x00-- Compile option: %s\n\x00mix1\x00-- Speedtest1 for SQLite %s %.48s\n\x00UNIQUE\x00autovacuum\x00big-transactions\x00cachesize\x00missing argument on %s\n\x00exclusive\x00fullfsync\x00checkpoint\x00explain\x00hard-heap-limit\x00heap\x00incrvacuum\x00journal\x00key\x00lookaside\x00memdb\x00multithread\x00nomemstat\x00mmap\x00nomutex\x00nosync\x00notnull\x00NOT NULL\x00output\x00-\x00wb\x00cannot open \"%s\" for writing\n\x00pagesize\x00pcache\x00primarykey\x00PRIMARY KEY\x00repeat\x00reprepare\x00serialized\x00singlethread\x00script\x00unable to open output file \"%s\"\n\x00sqlonly\x00shrink-memory\x00size\x00soft-heap-limit\x00stats\x00temp\x00argument to --temp should be integer between 0 and 9\x00testset\x00trace\x00threads\x00utf16le\x00utf16be\x00verify\x00vfs\x00reserve\x00stmtscanstatus\x00without-rowid\x00WITHOUT\x00STRICT\x00WITHOUT ROWID,STRICT\x00strict\x00help\x00?\x00unknown option: %s\nUse \"%s -?\" for help\n\x00surplus argument: %s\nUse \"%s -?\" for help\n\x00cannot allocate %d-byte heap\n\x00heap configuration failed: %d\n\x00cannot allocate %lld-byte pcache\n\x00pcache configuration failed: %d\n\x00:memory:\x00lookaside configuration failed: %d\n\x00random\x00PRAGMA temp_store=memory\x00PRAGMA mmap_size=%d\x00PRAGMA threads=%d\x00PRAGMA key('%s')\x00PRAGMA encoding=%s\x00PRAGMA auto_vacuum=FULL\x00PRAGMA auto_vacuum=INCREMENTAL\x00PRAGMA page_size=%d\x00PRAGMA cache_size=%d\x00PRAGMA synchronous=OFF\x00PRAGMA fullfsync=ON\x00PRAGMA locking_mode=EXCLUSIVE\x00PRAGMA journal_mode=%s\x00PRAGMA hard_heap_limit=%d\x00PRAGMA soft_heap_limit=%d\x00.explain\n.echo on\n\x00bad modifier on testset name: \"%s\"\x00       Begin testset \"%s\"\n\x00main\x00debug1\x00orm\x00cte\x00star\x00app\x00fp\x00json\x00trigger\x00parsenumber\x00rtree\x00unknown testset: \"%s\"\nChoices: cte debug1 fp main orm rtree trigger\n\x00Reset the database\x00SELECT name FROM main.sqlite_master WHERE sql LIKE 'CREATE %%TABLE%%'\x00DROP TABLE main.\"%w\"\x00SELECT name FROM temp.sqlite_master WHERE sql LIKE 'CREATE %%TABLE%%'\x00PRAGMA compile_options\x00-- Lookaside Slots Used:        %d (max %d)\n\x00-- Successful lookasides:       %d\n\x00-- Lookaside size faults:       %d\n\x00-- Lookaside OOM faults:        %d\n\x00-- Pager Heap Usage:            %d bytes\n\x00-- Page cache hits:             %d\n\x00-- Page cache misses:           %d\n\x00-- Page cache writes:           %d\n\x00-- Schema Heap Usage:           %d bytes\n\x00-- Statement Heap Usage:        %d bytes\n\x00-- Memory Used (bytes):         %d (max %d)\n\x00-- Outstanding Allocations:     %d (max %d)\n\x00-- Pcache Overflow Bytes:       %d (max %d)\n\x00-- Largest Allocation:          %d bytes\n\x00-- Largest Pcache Allocation:   %d bytes\n\x00"
