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

	"github.com/cybergarage/go-sqlparser/sqltest/util"
)

const (
	sqlTestResourceQueriesDirectory = "res/sql/"
	querySeparator                  = ";"
	ignorePrefix                    = "#"
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

func trimQuery(query string) string {
	trimStrings := []string{
		querySeparator,
		"\n",
		" "}
	for _, trimString := range trimStrings {
		query = strings.Trim(query, trimString)
	}
	return query
}

func formalizeQuery(query string) string {
	query = trimQuery(query)
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

func TestQueryFile(t *testing.T, file *util.File) {
	queryBytes, err := readQueryFile(file.Path())
	if err != nil {
		t.Error(err)
		return
	}

	queries := strings.Split(string(queryBytes), querySeparator)
	for _, query := range queries {
		query = trimQuery(query)
		if len(query) == 0 || strings.HasPrefix(query, ignorePrefix) {
			continue
		}
		t.Run(formalizeQuery(query), func(t *testing.T) {
			TestQueryString(t, query)
		})
	}
}

func TestQueryDirectoryWithRegex(t *testing.T, dir string, fileRegex string) {
	t.Run(fileRegex, func(t *testing.T) {
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
			t.Run(file.Name(), func(t *testing.T) {
				TestQueryFile(t, file)
			})
		}
	})
}
