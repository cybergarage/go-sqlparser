###################################################################
#
# usql-go
#
# Copyright (C) The go-sqlparser Authors 2018
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

PACKAGE_ROOT=${GITHUB}
PACKAGES=\
	${PACKAGE_ROOT} \
	${PACKAGE_ROOT}/parser \
	${PACKAGE_ROOT}/parser/antlr

SOURCE_ROOT=src/${PACKAGE_ROOT}

.PHONY: version antlr clean

all: test

VERSION_GO=${SOURCE_ROOT}/version.go

${VERSION_GO}: ${SOURCE_ROOT}/version.gen
	$< > $@

version: ${VERSION_GO}

antlr:
	- pushd ${SOURCE_ROOT}/parser/antlr && antlr4 -package antlr -Dlanguage=Go SQL.g4 && popd

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
