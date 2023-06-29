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
	"io"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/cybergarage/go-sqlparser/sql"
	"github.com/cybergarage/go-sqlparser/sqltest/util"
)

const (
	sqlTestResourceQueriesDirectory = "resources/sql/"
)

func readQueryFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func formalizeQuery(query string) string {
	query = strings.TrimSpace(query)
	trimStrings := []string{";"}
	for _, trimString := range trimStrings {
		query = strings.Trim(query, trimString)
	}
	regExps := []struct {
		From string
		To   string
	}{
		{`\s{2,}`, " "},
	}
	for _, regExp := range regExps {
		re := regexp.MustCompile(regExp.From)
		query = re.ReplaceAllString(query, regExp.To)
	}
	repStrings := []struct {
		From string
		To   string
	}{
		{"\n", " "},
		{"\t", " "},
		{"( ", "("},
		{" )", ")"},
	}
	for _, repString := range repStrings {
		query = strings.ReplaceAll(query, repString.From, repString.To)
	}
	return query
}

func testQueryString(t *testing.T, queryStr string) {
	parser := sql.NewParser()

	parsedQueries, err := parser.ParseString(queryStr)
	if err != nil {
		t.Error(err)
		return
	}

	if err != nil || len(parsedQueries) != 1 {
		t.Errorf("%s\n", queryStr)
		for _, query := range parsedQueries {
			t.Errorf("%s\n", query.String())
		}
		return
	}

	parsedQuery := parsedQueries[0]
	queryStr = formalizeQuery(queryStr)
	t.Logf("[S] %s\n", queryStr)
	parsedQueryStr := formalizeQuery(parsedQuery.String())
	// STEP1: Compare the parsed query with the original query for an exact match.
	if queryStr != parsedQueryStr {
		// STEP2: Compare the parsed query with the original query for semantic match.
		reParsedQueries, err := parser.ParseString(parsedQueryStr)
		if err == nil && len(reParsedQueries) == 1 {
			reParsedQuery := reParsedQueries[0]
			reParsedQueryStr := formalizeQuery(reParsedQuery.String())
			if queryStr != reParsedQueryStr {
				t.Errorf("[P] %s\n", parsedQueryStr)
				return
			}
		}
		t.Errorf("[P] %s\n", parsedQueryStr)
		return
	}
	t.Logf("[P] %s\n", parsedQueryStr)
}

func testQueryFile(t *testing.T, file *util.File) {
	queryBytes, err := readQueryFile(file.Path)
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
