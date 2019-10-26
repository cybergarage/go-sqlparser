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
PRODUCT_NAME=go-sql
PACKAGE_NAME=sql


PACKAGE_ROOT=${GIT_ROOT}/${PRODUCT_NAME}/${PACKAGE_NAME}
PACKAGES=\
	${PACKAGE_ROOT} \
	${PACKAGE_ROOT}/antlr
               
SOURCE_ROOT=${PACKAGE_NAME}

.PHONY: version antlr clean

all: test

VERSION_GO=${SOURCE_ROOT}/version.go

${VERSION_GO}: ${SOURCE_ROOT}/version.gen
	$< > $@

version: ${VERSION_GO}

antlr:
	- pushd ${SOURCE_ROOT}/antlr && make && popd

format:
	gofmt -w ${SOURCE_ROOT}

vet: format
	go vet ${PACKAGES}

build: vet
	go build -v ${PACKAGES}

test: vet
	go test -v -cover ${PACKAGES}

clean:
	go clean -i ${PACKAGES}
