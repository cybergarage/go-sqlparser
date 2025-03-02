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
	TestQueryDirectoryWithRegex(t, sqlTestResourceQueriesDirectory, re)
}

func TestInsertQueries(t *testing.T) {
	re := ".*insert.*\\.sql"
	TestQueryDirectoryWithRegex(t, sqlTestResourceQueriesDirectory, re)
}

func TestUpdateQueries(t *testing.T) {
	re := ".*update.*\\.sql"
	TestQueryDirectoryWithRegex(t, sqlTestResourceQueriesDirectory, re)
}

func TestSelectQueries(t *testing.T) {
	re := ".*select.*\\.sql"
	TestQueryDirectoryWithRegex(t, sqlTestResourceQueriesDirectory, re)
}

func TestDeleteQueries(t *testing.T) {
	re := ".*delete.*\\.sql"
	TestQueryDirectoryWithRegex(t, sqlTestResourceQueriesDirectory, re)
}

func TestAlterQueries(t *testing.T) {
	re := ".*alter.*\\.sql"
	TestQueryDirectoryWithRegex(t, sqlTestResourceQueriesDirectory, re)
}

func TestDropQueries(t *testing.T) {
	re := ".*drop.*\\.sql"
	TestQueryDirectoryWithRegex(t, sqlTestResourceQueriesDirectory, re)
}

func TestTransactionQueries(t *testing.T) {
	re := ".*transaction.*\\.sql"
	TestQueryDirectoryWithRegex(t, sqlTestResourceQueriesDirectory, re)
}

func TestSystemQueries(t *testing.T) {
	re := "system.*\\.sql"
	TestQueryDirectoryWithRegex(t, sqlTestResourceQueriesDirectory, re)
}

func TestExtraQueries(t *testing.T) {
	res := []string{
		".*copy.*\\.sql",
		".*commit.*\\.sql",
		".*vacuum.*\\.sql",
		".*truncate.*\\.sql",
		".*ping.*\\.sql",
		".*use.*\\.sql",
	}
	for _, re := range res {
		TestQueryDirectoryWithRegex(t, sqlTestResourceQueriesDirectory, re)
	}
}

func TestYCSBQueries(t *testing.T) {
	res := []string{
		"ycsb.*\\.sql",
	}
	for _, re := range res {
		TestQueryDirectoryWithRegex(t, sqlTestResourceQueriesDirectory, re)
	}
}

func TestSysbenchQueries(t *testing.T) {
	res := []string{
		"sysbench.*\\.sql",
	}
	for _, re := range res {
		TestQueryDirectoryWithRegex(t, sqlTestResourceQueriesDirectory, re, WithConfigValidationMode(ValidationModeStrict))
	}
}
