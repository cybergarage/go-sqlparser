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

import (
	"github.com/cybergarage/go-sqlparser/sql/util/strings"
)

// GroupByNone represents an empty GROUP BY clause.
const GroupByNone = ""

// GroupBy represents an GROUP interface.
type GroupBy interface {
	// ColumnName returns the column name.
	ColumnName() string
	// String returns the string representation.
	String() string
}

// groupBy represents an GROUP BY clause.
type groupBy struct {
	column string
}

// NewGroupBy returns a new groupBy instance.
func NewGroupBy() *groupBy {
	return &groupBy{
		column: "",
	}
}

// NewGroupByWith returns a new groupBy instance with the specified column name.
func NewGroupByWith(column string) *groupBy {
	return &groupBy{
		column: column,
	}
}

// SetColumn sets the column name.
func (groupBy *groupBy) SetColumn(name string) *groupBy {
	groupBy.column = name
	return groupBy
}

// ColumnName returns the column name.
func (groupBy *groupBy) ColumnName() string {
	return groupBy.column
}

// String returns the string representation.
func (groupBy *groupBy) String() string {
	if len(groupBy.column) == 0 {
		return ""
	}
	return strings.JoinWithSpace(
		[]string{
			"GROUP BY",
			groupBy.column,
		})
}
