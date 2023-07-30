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
)

func TestCreateQueries(t *testing.T) {
	re := ".*create.*\\.sql"
	testQueryDirectoryWithRegex(t, sqlTestResourceQueriesDirectory, re)
}

func TestInsertQueries(t *testing.T) {
	re := ".*insert.*\\.sql"
	testQueryDirectoryWithRegex(t, sqlTestResourceQueriesDirectory, re)
}

func TestUpdateQueries(t *testing.T) {
	re := ".*update.*\\.sql"
	testQueryDirectoryWithRegex(t, sqlTestResourceQueriesDirectory, re)
}

func TestSelectQueries(t *testing.T) {
	re := ".*select.*\\.sql"
	testQueryDirectoryWithRegex(t, sqlTestResourceQueriesDirectory, re)
}

func TestDeleteQueries(t *testing.T) {
	re := ".*delete.*\\.sql"
	testQueryDirectoryWithRegex(t, sqlTestResourceQueriesDirectory, re)
}

func TestAlterQueries(t *testing.T) {
	re := ".*alter.*\\.sql"
	testQueryDirectoryWithRegex(t, sqlTestResourceQueriesDirectory, re)
}

func TestDropQueries(t *testing.T) {
	re := ".*drop.*\\.sql"
	testQueryDirectoryWithRegex(t, sqlTestResourceQueriesDirectory, re)
}

func TestExtraQueries(t *testing.T) {
	res := []string{
		".*copy.*\\.sql",
	}
	for _, re := range res {
		testQueryDirectoryWithRegex(t, sqlTestResourceQueriesDirectory, re)
	}
}
