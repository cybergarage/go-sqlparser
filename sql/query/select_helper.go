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
	"github.com/cybergarage/go-sqlparser/sql/fn"
)

// SelectHelper is an interface for "SELECT" statement helpers.
type SelectHelper interface {
	// HasAggregator returns true if the statement has an aggregate function.
	HasAggregator() bool
}

// HasAggregator returns true if the statement has an aggregate function.
func (stmt *selectStmt) HasAggregator() bool {
	return stmt.selectors.HasAggregator()
}

// Aggregators returns the set of aggregators for the select statement.
func (stmt *selectStmt) Aggregators() (AggregatorSet, error) {
	aggrSet, err := stmt.selectors.Aggregators()
	if err != nil {
		return nil, err
	}
	resetOpts := []any{}
	groupBy := stmt.GroupBy()
	if groupBy != nil {
		resetOpts = append(resetOpts, fn.GroupBy(groupBy.ColumnName()))
	}
	return aggrSet, aggrSet.Reset(resetOpts...)
}
