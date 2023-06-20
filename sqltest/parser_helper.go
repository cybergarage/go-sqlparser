// Copyright (C) 2019 The go-sqlparser Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sqltest

import (
	"io/ioutil"
	"regexp"
	"strings"
	"testing"

	"github.com/cybergarage/go-sqlparser/sql"
	"github.com/cybergarage/go-sqlparser/sqltest/util"
)

const (
	sqlTestResourceQueriesDirectory = "resources/sql/"
)

func testQueryString(t *testing.T, queryStr string) {
	parser := sql.NewParser()

	queries, err := parser.ParseString(queryStr)
	if err != nil {
		t.Error(err)
		return
	}

	repStrings := []string{"\n", "\t", "  "}
	for _, repString := range repStrings {
		queryStr = strings.ReplaceAll(queryStr, repString, " ")
	}

	if err != nil {
		t.Errorf("%s\n", queryStr)
		for _, query := range queries {
			t.Errorf("%s\n", query.String())
		}
		return
	}

	t.Logf("[S] %s\n", queryStr)
	for _, query := range queries {
		t.Logf("[P] %s\n", query.String())
	}
}

func testQueryFile(t *testing.T, file *util.File) {
	queryBytes, err := ioutil.ReadFile(file.Path)
	if err != nil {
		t.Error(err)
		return
	}

	queries := strings.Split(string(queryBytes), "\n\n")
	for _, query := range queries {
		query = strings.TrimSpace(query)
		if len(query) <= 0 {
			continue
		}
		testQueryString(t, query)
	}
}

func testQueryDirectoryWithRegex(t *testing.T, dir string, fileRegex string) {
	re, err := regexp.Compile(fileRegex)
	if err != nil {
		t.Error(err)
		return
	}

	searchPath := util.NewFileWithPath(dir)
	files, err := searchPath.ListFilesWithRegexp(re)
	if err != nil {
		t.Error(err)
	}

	for _, file := range files {
		testQueryFile(t, file)
	}
}
