# Copyright (C) 2019 The go-sqlparser Authors. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

SHELL := bash

PREFIX?=$(shell pwd)

GIT_ROOT=github.com/cybergarage
PRODUCT_NAME=go-sqlparser

PKG_NAME=sql
PKG_SRC_ROOT=${PKG_NAME}
PKG_ROOT=${GIT_ROOT}/${PRODUCT_NAME}/${PKG_NAME}
PKGS=${PKG_ROOT}/...
               
TEST_PKG_NAME=${PKG_NAME}test
TEST_PKG_SRC_ROOT=${TEST_PKG_NAME}
TEST_PKG_ROOT=${GIT_ROOT}/${PRODUCT_NAME}/${TEST_PKG_NAME}
TEST_PKGS=${TEST_PKG_ROOT}/...

.PHONY: version antlr clean

all: test

VERSION_GO=${PKG_SRC_ROOT}/version.go

${VERSION_GO}: ${PKG_SRC_ROOT}/version.gen
	$< > $@

version: ${VERSION_GO}

antlr:
	- pushd ${PKG_SRC_ROOT}/antlr && make && popd

format:
	gofmt -w ${PKG_SRC_ROOT} ${TEST_PKG_SRC_ROOT}

vet: format
	go vet ${PKGS} ${TEST_PKGS}

lint:
	golangci-lint run ${PKGS} ${TEST_PKGS}

build: lint
	go build -v ${PKGS}

test: vet
	go test -v -cover ${PKGS} ${TEST_PKGS}

watchvet:
	fswatch -o . -e ".*" -i "\\.go$$" | xargs -n1 -I{} make vet

watchtest:
	fswatch -o . -e ".*" -i "\\.go$$" | xargs -n1 -I{} make test

watchlint:
	fswatch -o . -e ".*" -i "\\.go$$" | xargs -n1 -I{} make lint

clean:
	go clean -i ${PKGS} ${TEST_PKGS}