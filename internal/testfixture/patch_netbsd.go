// Copyright 2025 The libsqlite3-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main // import "modernc.org/libsqlite3"

import (
	"modernc.org/libc"
	"modernc.org/libc/pthread"
)

func ___libc_cond_destroy(t *libc.TLS, pCond uintptr) int32 {
	return libc.Xpthread_cond_destroy(t, pCond)
}

func ___libc_cond_init(t *libc.TLS, pCond, pAttr uintptr) int32 {
	return libc.Xpthread_cond_init(t, pCond, pAttr)
}

func ___libc_cond_signal(t *libc.TLS, pCond uintptr) int32 {
	return libc.Xpthread_cond_signal(t, pCond)
}

func ___libc_cond_wait(t *libc.TLS, pCond, pMutex uintptr) int32 {
	return libc.Xpthread_cond_wait(t, pCond, pMutex)
}

func ___libc_mutex_destroy(t *libc.TLS, pMutex uintptr) int32 {
	return libc.Xpthread_mutex_destroy(t, pMutex)
}

func ___libc_mutex_init(t *libc.TLS, pMutex, pAttr uintptr) int32 {
	return libc.Xpthread_mutex_init(t, pMutex, pAttr)
}

func ___libc_mutex_lock(t *libc.TLS, pMutex uintptr) int32 {
	return libc.Xpthread_mutex_lock(t, pMutex)
}

func ___libc_mutex_unlock(t *libc.TLS, pMutex uintptr) int32 {
	return libc.Xpthread_mutex_unlock(t, pMutex)
}

func ___libc_thr_detach(t *libc.TLS, thread pthread.Pthread_t) int32 {
	return libc.Xpthread_detach(t, thread)
}

func ___libc_thr_yield(t *libc.TLS) {
	libc.Xsched_yield(t)
}

func ___isnand(t *libc.TLS, n float64) int32 {
	return libc.Xisnan(t, n)
}
