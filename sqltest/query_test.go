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
	"errors"
	"testing"

	"github.com/cybergarage/go-sqlparser/sql"
)

func TestQueries(t *testing.T) {
	queries := []string{
		"SELECT * FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = 'foo' AND TABLE_NAME = 'bar'",
	}
	for _, query := range queries {
		t.Run(query, func(t *testing.T) {
			TestQueryString(t, query)
		})
	}
}

func TestEmptyQueries(t *testing.T) {
	queries := []string{
		"",
		";",
	}

	for _, query := range queries {
		t.Run(query, func(t *testing.T) {
			parser := sql.NewParser()
			_, err := parser.ParseString(query)
			if err == nil {
				t.Errorf("Failed to parse the empty query (%s)", query)
				return
			}
			if !errors.Is(err, sql.ErrEmptyQuery) {
				t.Errorf("Failed to parse the empty query (%s)", query)
				return
			}
		})
	}
}
