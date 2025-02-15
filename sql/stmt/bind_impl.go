// Copyright (C) 2025 The go-sqlparser Authors. All rights reserved.
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
// limitations under the License..

package stmt

type bindStmt struct {
	query  string
	params BindParams
}

// BindStatementOption is a bind statement option.
type BindStatementOption func(*bindStmt)

// NewBindStatement creates a new bind statement.
func NewBindStatement(opts ...BindStatementOption) BindStatement {
	stmt := &bindStmt{
		query:  "",
		params: []BindParam{},
	}
	for _, opt := range opts {
		opt(stmt)
	}
	return stmt
}

// Statement returns the statement.
func (stmt *bindStmt) Statement() (Statement, error) {
	return nil, nil
}
