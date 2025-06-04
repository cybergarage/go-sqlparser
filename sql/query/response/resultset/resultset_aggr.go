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
	"fmt"

	"github.com/cybergarage/go-sqlparser/sql/fn"
	"github.com/cybergarage/go-sqlparser/sql/query"
)

// GroupBy represents a function type for grouping rows in a result set.
type GroupBy = query.GroupBy

// AggregatedResultSetOptions represents a functional option for configuring an aggregated result set.
type AggregatedResultSetOptions func(*aggrResultSet) error

type aggrResultSet struct {
	tblSchema query.Schema
	selectors query.Selectors
	srcRs     ResultSet
	groupBy   GroupBy
}

// NewAggregatedResultSet creates a new ResultSet with aggregated rows based on the provided options.
func WithAggregatedResultSetTableSchema(tblSchema query.Schema) AggregatedResultSetOptions {
	return func(rs *aggrResultSet) error {
		rs.tblSchema = tblSchema
		return nil
	}
}

// WithAggregatedResultSetSelectors sets the selectors for the aggregated result set.
func WithAggregatedResultSetSelectors(selectors query.Selectors) AggregatedResultSetOptions {
	return func(rs *aggrResultSet) error {
		rs.selectors = selectors
		return nil
	}
}

// WithAggregatedResultSetSource sets the source ResultSet for the aggregated result set.
func WithAggregatedResultSetSource(srcRs ResultSet) AggregatedResultSetOptions {
	return func(rs *aggrResultSet) error {
		rs.srcRs = srcRs
		return nil
	}
}

// WithAggregatedResultSetGroupBy sets the group by function for the aggregated result set.
func WithAggregatedResultSetGroupBy(groupBy GroupBy) AggregatedResultSetOptions {
	return func(rs *aggrResultSet) error {
		rs.groupBy = groupBy
		return nil
	}
}

// NewAggregatedResultSetFrom creates a new ResultSet with aggregated rows from the given ResultSet.
func NewAggregatedResultSetFrom(opts ...AggregatedResultSetOptions) (ResultSet, error) {
	// Set the specified options

	aggrOpts := &aggrResultSet{
		tblSchema: nil,
		selectors: nil,
		srcRs:     nil,
		groupBy:   nil,
	}

	for _, opt := range opts {
		if err := opt(aggrOpts); err != nil {
			return nil, err
		}
	}

	// Validate the options

	tblSchema := aggrOpts.tblSchema
	if tblSchema == nil {
		return nil, fmt.Errorf("table schema is required for aggregated result set")
	}

	selectors := aggrOpts.selectors
	if selectors == nil {
		return nil, fmt.Errorf("selectors are required for aggregated result set")
	}

	srcRs := aggrOpts.srcRs
	if srcRs == nil {
		return nil, fmt.Errorf("source result set is required for aggregated result set")
	}

	groupBy := aggrOpts.groupBy

	// Check if the selectors have aggregators

	if !selectors.HasAggregator() {
		return srcRs, nil
	}

	// Generate a new schema for the aggregated result set

	rsSchema := srcRs.Schema()
	aggrSchema := NewSchema(
		WithSchemaDatabaseName(rsSchema.DatabaseName()),
		WithSchemaTableSchema(tblSchema),
		WithSchemaSelectors(selectors),
	)

	// Generate a new aggregated rows

	aggrSet, err := selectors.Aggregators()
	if err != nil {
		return nil, err
	}

	resetOpts := []any{}
	if groupBy != nil {
		resetOpts = append(resetOpts, fn.GroupBy(groupBy.ColumnName()))
	}

	err = aggrSet.Reset(resetOpts...)
	if err != nil {
		return nil, err
	}

	for srcRs.Next() {
		row, err := srcRs.Row()
		if err != nil {
			return nil, err
		}
		err = aggrSet.Aggregate(row.Object())
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
				WithRowSchema(aggrSchema),
				WithRowObject(rowMap)))
	}

	// Return an aggregated result set

	return NewResultSet(
		WithResultSetSchema(aggrSchema),
		WithResultSetRows(aggrRows),
	), nil
}
