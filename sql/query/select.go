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

package query

import "github.com/cybergarage/go-sqlparser/sql/util/strings"

// SelectOption represents a select option function.
type SelectOption = func(*Select)

// Select is a "SELECT" statement.
type Select struct {
	TableList
	SelectorList
	*Condition
	orderBy *OrderBy
	limit   *Limit
}

// NewSelectWith returns a new Select statement instance with the specified parameters.
func NewSelectWith(selectors SelectorList, tbls TableList, w *Condition, opts ...SelectOption) *Select {
	stmt := &Select{
		SelectorList: selectors,
		TableList:    tbls,
		Condition:    w,
		orderBy:      NewOrderBy(),
		limit:        NewLimit(),
	}
	for _, opt := range opts {
		opt(stmt)
	}
	return stmt
}

// WithSelectOrderBy sets order by options.
func WithSelectOrderBy(name string, order Order) func(*Select) {
	return func(stmt *Select) {
		stmt.orderBy = NewOrderByWith(name, order)
	}
}

// WithSelectLimit sets order by options.
func WithSelectLimit(offset int, limit int) func(*Select) {
	return func(stmt *Select) {
		stmt.limit = NewLimitWith(offset, limit)
	}
}

// StatementType returns the statement type.
func (stmt *Select) StatementType() StatementType {
	return SelectStatement
}

// From returns the source table list.
func (stmt *Select) From() TableList {
	return stmt.Tables()
}

// Where returns the condition.
func (stmt *Select) Where() *Condition {
	return stmt.Condition
}

// OrderBy returns the order by clause.
func (stmt *Select) OrderBy() *OrderBy {
	return stmt.orderBy
}

// Limit returns the limit clause.
func (stmt *Select) Limit() *Limit {
	return stmt.limit
}

// String returns the statement string representation.
func (stmt *Select) String() string {
	selectorStr := "*"
	if 0 < len(stmt.SelectorList) {
		selectorStr = stmt.SelectorList.SelectorString()
	}
	strs := []string{
		"SELECT",
		selectorStr,
		"FROM",
		stmt.TableList.String(),
	}
	if stmt.Condition != nil {
		strs = append(strs, "WHERE", stmt.Condition.String())
	}
	strs = append(strs, stmt.orderBy.String())
	strs = append(strs, stmt.limit.String())
	return strings.JoinWithSpace(strs)
}
