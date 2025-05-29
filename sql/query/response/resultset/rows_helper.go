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

type rows struct {
	mapRows   []map[string]any
	selectors query.Selectors
	scheama   Schema
	groupBy   string
}

// RowsOption represents a functional option for rows.
type RowsOption func(*rows) error

// WithRowsMapRows sets the rows from a slice of map[string]any.
func WithRowsMapRows(v []map[string]any) RowsOption {
	return func(r *rows) error {
		r.mapRows = v
		return nil
	}
}

// WithRowsSelector sets the selectors for the rows.
func WithRowsSelectors(selectors query.Selectors) RowsOption {
	return func(r *rows) error {
		r.selectors = selectors
		return nil
	}
}

// WithRowsGroupBy sets the group by clause for the rows.
func WithRowsGroupBy(groupBy string) RowsOption {
	return func(r *rows) error {
		r.groupBy = groupBy
		return nil
	}
}

// WithRowsSchema sets the schema for the rows.
func WithRowsSchema(schema Schema) RowsOption {
	return func(r *rows) error {
		r.scheama = schema
		return nil
	}
}

// NewRows creates a new rows instance with the opt
func NewRows(opts ...RowsOption) ([]Row, error) {
	r := &rows{
		mapRows:   []map[string]any{},
		selectors: nil,
	}
	for _, opt := range opts {
		if err := opt(r); err != nil {
			return nil, err
		}
	}

	// Validate required fields

	if r.scheama == nil {
		return nil, fmt.Errorf("schema %w", query.ErrNotSet)
	}

	if r.selectors == nil {
		return nil, fmt.Errorf("selectors %w", query.ErrNotSet)
	}

	// Set seed map rows

	mapRows := r.mapRows

	// Aggregate new row maps if aggregators are present

	if r.selectors.HasAggregator() {
		aggrSet, err := r.selectors.Aggregators()
		if err != nil {
			return nil, err
		}

		resetOpts := []any{}
		if r.groupBy != "" {
			resetOpts = append(resetOpts, fn.GroupBy(r.groupBy))
		}

		err = aggrSet.Reset(resetOpts...)
		if err != nil {
			return nil, err
		}

		for _, row := range mapRows {
			err := aggrSet.Aggregate(row)
			if err != nil {
				return nil, err
			}
		}

		resultSet, err := aggrSet.Finalize()
		if err != nil {
			return nil, err
		}

		mapRows = []map[string]any{}
		for resultSet.Next() {
			rowMap, err := resultSet.Map()
			if err != nil {
				return nil, err
			}
			mapRows = append(mapRows, rowMap)
		}
	}

	// Generate rows from map rows

	rsRows := []Row{}
	for _, rowMap := range mapRows {
		rowValues := []any{}
		for _, selector := range r.selectors {
			var rowValue any
			rowValue = nil
			if fx, ok := selector.Function(); ok {
				if executor, err := fx.Executor(); err == nil {
					rowValue, err = executor.Execute(fn.NewMapWithMap(rowMap))
					if err != nil {
						return nil, err
					}
				}
			}
			if rowValue == nil {
				var ok bool
				selectorName := selector.String()
				rowValue, ok = rowMap[selectorName]
				if !ok {
					return nil, fmt.Errorf("selector %s not found in row map", selectorName)
				}
			}
			rowValues = append(rowValues, rowValue)
		}
		rsRow := NewRow(
			WithRowSchema(r.scheama),
			WithRowValues(rowValues),
		)
		rsRows = append(rsRows, rsRow)
	}

	return rsRows, nil
}
