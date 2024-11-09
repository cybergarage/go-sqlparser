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
type SelectOption = func(*selectStmt)

// Select represents a "SELECT" statement interface.
type Select interface {
	Statement
	// IsSelectAll returns true if the statement is a "SELECT *".
	IsSelectAll() bool
	// Selectors returns the selectors.
	Selectors() SelectorList
	// From returns the source table list.
	From() TableList
	// Limit returns the limit clause.
	Limit() Limit
	// GroupBy returns the group by clause.
	GroupBy() GroupBy
	// OrderBy returns the order by clause.
	OrderBy() OrderBy
	// Where returns the condition.
	Where() Condition
}

// selectStmt is a "SELECT" statement.
type selectStmt struct {
	TableList
	SelectorList
	Condition
	orderBy OrderBy
	limit   Limit
	groupBy GroupBy
}

// NewSelectWith returns a new selectStmt statement instance with the specified parameters.
func NewSelectWith(selectors SelectorList, tbls TableList, w Condition, opts ...SelectOption) *selectStmt {
	stmt := &selectStmt{
		SelectorList: selectors,
		TableList:    tbls,
		Condition:    w,
		orderBy:      NewOrderBy(),
		limit:        NewLimit(),
		groupBy:      NewGroupBy(),
	}
	for _, opt := range opts {
		opt(stmt)
	}
	return stmt
}

// WithSelectOrderBy sets order by options.
func WithSelectOrderBy(orderBy OrderBy) func(*selectStmt) {
	return func(stmt *selectStmt) {
		stmt.orderBy = orderBy
	}
}

// WithSelectLimit sets order by options.
func WithSelectLimit(offset int, limit int) func(*selectStmt) {
	return func(stmt *selectStmt) {
		stmt.limit = NewLimitWith(offset, limit)
	}
}

// WithSelectOrderBy sets order by options.
func WithSelectGroupBy(name string) func(*selectStmt) {
	return func(stmt *selectStmt) {
		stmt.groupBy = NewGroupByWith(name)
	}
}

// StatementType returns the statement type.
func (stmt *selectStmt) StatementType() StatementType {
	return SelectStatement
}

// From returns the source table list.
func (stmt *selectStmt) From() TableList {
	return stmt.Tables()
}

// Where returns the condition.
func (stmt *selectStmt) Where() Condition {
	return stmt.Condition
}

// OrderBy returns the order by clause.
func (stmt *selectStmt) OrderBy() OrderBy {
	return stmt.orderBy
}

// Limit returns the limit clause.
func (stmt *selectStmt) Limit() Limit {
	return stmt.limit
}

// GroupBy returns the group by clause.
func (stmt *selectStmt) GroupBy() GroupBy {
	return stmt.groupBy
}

// String returns the statement string representation.
func (stmt *selectStmt) String() string {
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
	if !stmt.Condition.IsEmpty() {
		strs = append(strs, "WHERE", stmt.Condition.String())
	}
	strs = append(strs, stmt.orderBy.String())
	strs = append(strs, stmt.limit.String())
	strs = append(strs, stmt.groupBy.String())
	return strings.JoinWithSpace(strs)
}
