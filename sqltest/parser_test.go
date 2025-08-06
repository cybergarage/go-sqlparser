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
	"math/rand"
	"testing"
	"time"

	"github.com/cybergarage/go-sqlparser/sql"
)

func TestParallelParsing(t *testing.T) {
	queries := []string{
		"SELECT id, name FROM users WHERE age > 30",
		"INSERT INTO orders (id, user_id, total) VALUES (1, 42, 99.99)",
		"UPDATE products SET price = price * 1.1 WHERE category = 'electronics'",
		"CREATE TABLE logs (id INT PRIMARY KEY, message TEXT, created_at TIMESTAMP)",
		"DROP TABLE IF EXISTS temp_data",
		"ALTER TABLE users ADD COLUMN last_login TIMESTAMP",
	}

	parseQuery := func(query string) {
		parser := sql.NewParser()
		stmts, err := parser.ParseString(query)
		if err != nil {
			t.Error(err)
			return
		}
		time.Sleep(time.Millisecond * 100)
		if len(stmts) != 1 {
			t.Errorf("Expected 1 statement, got %d", len(stmts))
			return
		}
	}

	done := make(chan struct{}, 100)
	for range 100 {
		go func() {
			defer func() { done <- struct{}{} }()
			query := queries[rand.Intn(len(queries))]
			parseQuery(query)
		}()
	}
	for range 100 {
		<-done
	}
}

func TestNestedParsing(t *testing.T) {
	queries := []string{
		"SELECT id, name FROM users WHERE age > 30",
		"INSERT INTO orders (id, user_id, total) VALUES (1, 42, 99.99)",
	}

	parser01 := sql.NewParser()
	stmts01, err := parser01.ParseString(queries[0])
	if err != nil {
		t.Error(err)
		return
	}

	parser02 := sql.NewParser()
	stmts02, err := parser02.ParseString(queries[1])
	if err != nil {
		t.Error(err)
		return
	}
	if len(stmts02) != 1 {
		t.Errorf("Expected 1 statement, got %d", len(stmts02))
		return
	}

	if len(stmts01) != 1 {
		t.Errorf("Expected 1 statement, got %d", len(stmts01))
		return
	}
}
