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

func TestParallelParse(t *testing.T) {
	queries := []string{
		"SELECT * FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = 'foo' AND TABLE_NAME = 'bar'",
	}

	parseQuery := func(query string) {
		parser := sql.NewParser()
		_, err := parser.ParseString(query)
		if err != nil {
			t.Error(err)
			return
		}
	}

	for _, query := range queries {
		t.Run(query, func(t *testing.T) {
			t.Parallel()
			done := make(chan struct{}, 100)
			for range 100 {
				go func() {
					defer func() { done <- struct{}{} }()
					parseQuery(query)
				}()
			}
			for range 100 {
				<-done
			}
		})
	}
}
