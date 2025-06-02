// Copyright (C) 2024 The go-mysql Authors. All rights reserved.
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

package resultset

import (
	"github.com/cybergarage/go-sqlparser/sql/fn"
)

// GroupBy represents a function type for grouping rows in a result set.
type GroupBy = fn.GroupBy

// NewAggregatedResultSetFrom creates a new ResultSet with aggregated rows from the given ResultSet.
func NewAggregatedResultSetFrom(rs ResultSet, opts ...any) (ResultSet, error) {
	groupBy := ""
	for _, opt := range opts {
		if g, ok := opt.(GroupBy); ok {
			groupBy = string(g)
		}
	}

	schema := rs.Schema()
	selectors := schema.Selectors()

	if !selectors.HasAggregator() {
		return rs, nil
	}

	aggrSet, err := selectors.Aggregators()
	if err != nil {
		return nil, err
	}

	resetOpts := []any{}
	if 0 < len(groupBy) {
		resetOpts = append(resetOpts, fn.GroupBy(groupBy))
	}

	err = aggrSet.Reset(resetOpts...)
	if err != nil {
		return nil, err
	}

	for rs.Next() {
		row, err := rs.Row()
		if err != nil {
			return nil, err
		}
		err = aggrSet.Aggregate(row)
		if err != nil {
			return nil, err
		}
	}

	aggrResultSet, err := aggrSet.Finalize()
	if err != nil {
		return nil, err
	}

	aggrRows := []Row{}
	for aggrResultSet.Next() {
		rowMap, err := aggrResultSet.Map()
		if err != nil {
			return nil, err
		}
		aggrRows = append(aggrRows,
			NewRow(
				WithRowSchema(schema),
				WithRowObject(rowMap)))
	}

	return NewResultSet(
		WithResultSetSchema(schema),
		WithResultSetRows(aggrRows),
	), nil
}
