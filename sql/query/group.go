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

// GroupBy represents an ORDER BY clause.
type GroupBy struct {
	column string
}

// NewGroupBy returns a new GroupBy instance.
func NewGroupBy() *GroupBy {
	return &GroupBy{
		column: "",
	}
}

// NewGroupByWith returns a new GroupBy instance with the specified column name.
func NewGroupByWith(column string) *GroupBy {
	return &GroupBy{
		column: column,
	}
}

// SetColumn sets the column name.
func (groupBy *GroupBy) SetColumn(name string) *GroupBy {
	groupBy.column = name
	return groupBy
}

// Column returns the column name.
func (groupBy *GroupBy) Column() string {
	return groupBy.column
}

// String returns the string representation.
func (groupBy *GroupBy) String() string {
	if len(groupBy.column) <= 0 {
		return ""
	}
	return strings.JoinWithSpace(
		[]string{
			"ORDER BY",
			groupBy.column,
		})
}
