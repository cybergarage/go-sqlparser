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
PKG_COVER=${PKG_NAME}-cover
PKG_SRC_ROOT=${PKG_NAME}
PKG_ROOT=${GIT_ROOT}/${PRODUCT_NAME}/${PKG_NAME}
PKG=${PKG_ROOT}
               
TEST_PKG_NAME=${PKG_NAME}test
TEST_PKG_SRC_ROOT=${TEST_PKG_NAME}
TEST_PKG_ROOT=${GIT_ROOT}/${PRODUCT_NAME}/${TEST_PKG_NAME}
TEST_PKG=${TEST_PKG_ROOT}

.PHONY: version antlr clean

all: test

VERSION_GO=${PKG_SRC_ROOT}/version.go

${VERSION_GO}: ${PKG_SRC_ROOT}/version.gen
	$< > $@

version: ${VERSION_GO}
	-git commit ${PKG_SRC_DIR}/version.go -m "Update version"

antlr:
	-pushd ${PKG_SRC_ROOT}/antlr && make && popd

format: version
	gofmt -w ${PKG_SRC_ROOT} ${TEST_PKG_SRC_ROOT}

vet: format
	go vet ${PKG} ${TEST_PKG}

lint: vet
	golangci-lint run ${PKG_SRC_ROOT}/... ${TEST_PKG_SRC_ROOT}/...

build: lint
	go build -v ${PKG}

test: lint
	go test -v -p 1 -timeout 10m -cover -coverpkg=${PKG}/... -coverprofile=${PKG_COVER}.out ${PKG}/... ${TEST_PKG}/...

watchvet:
	fswatch -o . -e ".*" -i "\\.go$$" | xargs -n1 -I{} make vet

watchtest:
	fswatch -o . -e ".*" -i "\\.go$$" | xargs -n1 -I{} make test

watchlint:
	fswatch -o . -e ".*" -i "\\.go$$" | xargs -n1 -I{} make lint

clean:
	go clean -i ${PKG} ${TEST_PKG}

#
# Document
#

%.md : %.adoc
	asciidoctor -b docbook -a leveloffset=+1 -o - $< | pandoc -t markdown_strict --wrap=none -f docbook > $@
csvs := $(wildcard doc/*/*.csv)
docs := $(patsubst %.adoc,%.md,$(wildcard doc/*.adoc))
doc_touch: $(csvs)
	touch doc/*.adoc

doc: doc_touch $(docs)
