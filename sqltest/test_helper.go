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
	"testing"

	"github.com/cybergarage/go-sqlparser/sql"
)

func TestQueryString(t *testing.T, queryStr string) {
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
			if parsedQueryStr != reParsedQueryStr {
				t.Errorf("[P] %s\n", parsedQueryStr)
				return
			}
		} else {
			t.Errorf("[P] %s\n", parsedQueryStr)
			return
		}
	}
	t.Logf("[P] %s\n", parsedQueryStr)
}
