# Copyright 2023 The libsqlite3-go Authors. All rights reserved.
# Use of the source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all clean dev download edit editor extraquick generate mptest test work xtest speedtest1 tcltest

SHELL=/bin/bash -o pipefail

GREP = 'TRC\|TODO\|ERRORF\|FAIL\|undefined:'

DIR = /tmp/libsqlite3
ZIP = sqlite-amalgamation-3530200.zip
ZIP2 = sqlite-src-3530200.zip
URL = https://www.sqlite.org/2026/$(ZIP)
URL2 = https://www.sqlite.org/2026/$(ZIP2)

all: editor
	golint 2>&1
	staticcheck 2>&1

build_all_targets:
	./build_all_targets.sh
	echo done

clean:
	rm -f log-* cpu.test mem.test *.out go.work*
	go clean

clean-dev:
	rm -f ccgo_linux_amd64.go internal/testfixture/ccgo_linux_amd64.go
	rm -f ccgo_windows.go internal/testfixture/ccgo_windows.go
	rm -f internal/autogen/linux_amd64.mod internal/autogen/windows*.mod

edit:
	@if [ -f "Session.vim" ]; then gvim -S & else gvim -p Makefile go.mod builder.json all_test.go generator.go libsqlite3.go CHANGELOG.md & fi

editor:
	gofmt -l -s -w .
	go test -c -o /dev/null
	go build -v  -o /dev/null ./...
	go build -o /dev/null generator*.go

download:
	@if [ ! -f $(ZIP) ]; then wget $(URL) ; fi
	@if [ ! -f $(ZIP2) ]; then wget $(URL2) ; fi

generate: download
	mkdir -p $(DIR) || true
	rm -rf $(DIR)/*
	GO_GENERATE_DIR=$(DIR) go run generator*.go
	go build -v ./...
	go build -v ./...
	git status

dev: download
	mkdir -p $(DIR) || true
	rm -rf $(DIR)/*
	echo -n > /tmp/ccgo.log
	GO_GENERATE_DIR=$(DIR) GO_GENERATE_DEV=1 go run -tags=ccgo.dmesg,ccgo.assert generator*.go
	go build -v ./...
	git status
	grep $(GREP) /tmp/ccgo.log || true

extraquick:
	go test -v -timeout 24h -run Tcl -verbose=1 -suite=extraquick

mptest:
	go test -v -timeout 24h -run TestConcurrentProcesses

speedtest1:
	go run ./speedtest1

tcltest:
	go test -v -timeout 24h -run TestTcl

test:
	go test -v -timeout 24h

windows: download
	mkdir -p $(DIR) || true
	rm -rf $(DIR)/*
	echo -n > /tmp/ccgo.log
	GO_GENERATE_WIN=1 GO_GENERATE_DIR=$(DIR) go run generator*.go
	GOOS=windows GOARCH=amd64 go build -v ./...
	GOOS=windows GOARCH=amd64 go test -v -c -o /dev/null
	GOOS=windows GOARCH=arm64 go build -v ./...
	GOOS=windows GOARCH=arm64 go test -v -c -o /dev/null
	git status

windows-dev: download
	mkdir -p $(DIR) || true
	rm -rf $(DIR)/*
	echo -n > /tmp/ccgo.log
	GO_GENERATE_WIN=1 GO_GENERATE_DIR=$(DIR) GO_GENERATE_DEV=1 go run -tags=ccgo.dmesg,ccgo.assert generator*.go
	GOOS=windows GOARCH=amd64 go build -v ./...
	GOOS=windows GOARCH=amd64 go test -v -c -o /dev/null
	GOOS=windows GOARCH=arm64 go build -v ./...
	GOOS=windows GOARCH=arm64 go test -v -c -o /dev/null
	git status
	grep $(GREP) /tmp/ccgo.log || true

windows_386: download
	mkdir -p $(DIR) || true
	rm -rf $(DIR)/*
	echo -n > /tmp/ccgo.log
	GO_GENERATE_WIN32=1 GO_GENERATE_DIR=$(DIR) go run generator*.go
	GOOS=windows GOARCH=386 go build -v ./...
	GOOS=windows GOARCH=386 go test -v -c -o /dev/null
	git status

windows_386-dev: download
	mkdir -p $(DIR) || true
	rm -rf $(DIR)/*
	echo -n > /tmp/ccgo.log
	GO_GENERATE_WIN32=1 GO_GENERATE_DIR=$(DIR) GO_GENERATE_DEV=1 go run -tags=ccgo.dmesg,ccgo.assert generator*.go
	GOOS=windows GOARCH=386 go build -v ./...
	GOOS=windows GOARCH=386 go test -v -c -o /dev/null
	git status
	grep $(GREP) /tmp/ccgo.log || true

work:
	rm -f go.work*
	go work init
	go work use .
	go work use ../cc/v4
	go work use ../ccgo/v3
	go work use ../ccgo/v4
	go work use ../libc
	go work use ../libtcl8.6
	go work use ../libz

xtest:
	GO_GENERATE_NOWIN=1 GO_GENERATE_TEST=1 GOMAXPROCS=1 make dev ; go test -v -run Options
