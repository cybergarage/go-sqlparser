###################################################################
#
# usql-go
#
# Copyright (C) The go-sqlparser Authors 2017
#
# This is licensed under BSD-style license, see file COPYING.
#
###################################################################

SHELL := bash

PREFIX?=$(shell pwd)
GOPATH:=$(shell pwd)
export GOPATH

PACKAGE_NAME=usql

GITHUB_ROOT=github.com/cybergarage/usql-go
GITHUB=${GITHUB_ROOT}/${PACKAGE_NAME}

PACKAGE_ID=${GITHUB}
PACKAGES=\
	${PACKAGE_ID}

SOURCE_ROOT_DIR=src/${PACKAGE_ROOT}

.PHONY: version clean

all: test

VERSION_GO=${SOURCE_ROOT_DIR}/version.go

${VERSION_GO}: ${SOURCE_ROOT_DIR}/version.gen
	$< > $@

version: ${VERSION_GO}

format:
	gofmt -w src/${GITHUB}

vet: format
	go vet ${PACKAGES}

build: vet
	go build -v ${PACKAGES}

test: vet
	go test -v -cover ${PACKAGES}

clean:
	-rm ${PREFIX}/bin/*
	rm -rf _obj
	go clean -i ${PACKAGES}
