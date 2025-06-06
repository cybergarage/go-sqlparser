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

package fn

import (
	"fmt"

	"github.com/cybergarage/go-safecast/safecast"
)

// AggrResetFunc is a function that resets the aggregation state.
type AggrResetFunc func(aggr *aggrImpl) (float64, error)

// AggrAggregateFunc is a function that performs aggregation on the given values.
type AggrAggregateFunc func(aggr *aggrImpl, accumulatedValue float64, inputValue float64) (float64, error)

// AggrFinalizeFunc is a function that finalizes the aggregation and returns the result.
type AggrFinalizeFunc func(aggr *aggrImpl, accumulatedValue float64, accumulatedCount int) (any, error)

type aggrImpl struct {
	name        string
	args        []string
	columns     []string
	aggrs       []float64
	counts      int
	groupBy     string
	groupAggrs  map[any][]float64
	groupCounts map[any]int
	resetFunc   AggrResetFunc
	aggFunc     AggrAggregateFunc
	finalFunc   AggrFinalizeFunc
}

// aggrOption is a function that configures the Aggr aggregator.
type aggrOption func(*aggrImpl) error

// newAggr creates a new Aggr aggregator with the given options.
func newAggr() *aggrImpl {
	aggr := &aggrImpl{
		name:        "",
		args:        make([]string, 0),
		columns:     make([]string, 0),
		groupBy:     "",
		aggrs:       make([]float64, 0),
		counts:      0,
		groupAggrs:  make(map[any][]float64),
		groupCounts: make(map[any]int),
		resetFunc:   nil,
		aggFunc:     nil,
		finalFunc:   nil,
	}
	return aggr
}

// withAggrName sets the name of the Aggr aggregator.
func withAggrName(name string) aggrOption {
	return func(aggr *aggrImpl) error {
		aggr.name = name
		return nil
	}
}

// withAggrArguments sets the arguments for the Aggr aggregator.
func withAggrArguments(args []string) aggrOption {
	return func(aggr *aggrImpl) error {
		if 1 < len(aggr.args) {
			return fmt.Errorf("multiple argument %w : %v", ErrNotSupported, aggr.args)
		}
		aggr.args = args
		return nil
	}
}

// withAggrResetFunc sets the reset function for the Aggr aggregator.
func withAggrResetFunc(resetFunc AggrResetFunc) aggrOption {
	return func(aggr *aggrImpl) error {
		aggr.resetFunc = resetFunc
		return nil
	}
}

// withAggrAggreateFunc sets the aggregation function for the Aggr aggregator.
func withAggrAggreateFunc(aggFunc AggrAggregateFunc) aggrOption {
	return func(aggr *aggrImpl) error {
		aggr.aggFunc = aggFunc
		return nil
	}
}

// withAggrFinalizeFunc sets the finalization function for the Aggr aggregator.
func withAggrFinalizeFunc(finalFunc AggrFinalizeFunc) aggrOption {
	return func(aggr *aggrImpl) error {
		aggr.finalFunc = finalFunc
		return nil
	}
}

// withAggrGroupBy sets the group by column for the Aggr aggregator.
func withAggrGroupBy(group string) aggrOption {
	return func(aggr *aggrImpl) error {
		if len(group) == 0 {
			return nil
		}
		aggr.groupBy = group
		return nil
	}
}

// Name returns the name of the aggregator.
func (aggr *aggrImpl) Name() string {
	return aggr.name
}

// Type returns the type of the aggregator.
func (aggrImpl *aggrImpl) Type() FunctionType {
	return AggregateFunction
}

// Arguments returns the arguments of the aggregator.
func (aggr *aggrImpl) Arguments() []string {
	return aggr.args
}

// GroupBy returns the group by column name and a boolean indicating if it is set.
func (aggr *aggrImpl) GroupBy() (string, bool) {
	if len(aggr.groupBy) == 0 {
		return "", false
	}
	return aggr.groupBy, true
}

// Reset resets the aggregator to its initial state.
func (aggr *aggrImpl) Reset(opts ...any) error {
	// Apply options to the aggregator
	for _, opt := range opts {
		switch opt := opt.(type) {
		case aggrOption:
			if err := opt(aggr); err != nil {
				return fmt.Errorf("failed to apply option: %w", err)
			}
		case GroupBy:
			aggr.groupBy = string(opt)
		default:
			return fmt.Errorf("%w option type %T is not supported", ErrInvalid, opt)
		}
	}

	// Set the column names
	aggr.columns = []string{}
	if groupBy, ok := aggr.GroupBy(); ok {
		aggr.columns = append(aggr.columns, groupBy)
	}
	for _, arg := range aggr.args {
		aggr.columns = append(aggr.columns, fmt.Sprintf("%s(%s)", aggr.Name(), arg))
	}

	// Validate the arguments

	if len(aggr.columns) == 0 {
		return fmt.Errorf("no argument %w", ErrNotSupported)
	}

	// Reset aggregator variables

	aggr.aggrs = make([]float64, len(aggr.columns))
	for n := range aggr.aggrs {
		nv, err := aggr.resetFunc(aggr)
		if err != nil {
			return err
		}
		aggr.aggrs[n] = nv
	}

	aggr.counts = 0

	aggr.groupAggrs = make(map[any][]float64)
	aggr.groupCounts = make(map[any]int)

	return nil
}

// AggregateRow aggregates a row of data using the aggregator.
// The row is expected to be an array where the first element is the group value (if grouping is enabled),
func (aggr *aggrImpl) AggregateRow(row []any) error {
	if len(aggr.columns) != len(row) {
		return fmt.Errorf("%w column count (%d != %d)", ErrInvalid, len(aggr.columns), len(row))
	}

	toFloat64 := func(v any) (float64, error) {
		switch v := v.(type) {
		case nil:
			return 0, nil
		default:
			var fv float64
			if err := safecast.ToFloat64(v, &fv); err != nil {
				return 0, fmt.Errorf("%w value %v : %s", ErrInvalid, v, err)
			}
			return fv, nil
		}
	}

	if _, ok := aggr.GroupBy(); ok {
		group := row[0]
		if _, ok := aggr.groupAggrs[group]; !ok {
			aggr.groupAggrs[group] = make([]float64, (len(aggr.columns) - 1))
			for n := range aggr.groupAggrs[group] {
				nv, err := aggr.resetFunc(aggr)
				if err != nil {
					return err
				}
				aggr.groupAggrs[group][n] = nv
			}
			aggr.groupCounts[group] = 0
		}
		for n, rv := range row[1:] {
			fv, err := toFloat64(rv)
			if err != nil {
				return fmt.Errorf("[%d] %w row : %s", n+1, ErrInvalid, err)
			}
			nv, err := aggr.aggFunc(aggr, aggr.groupAggrs[group][n], fv)
			if err != nil {
				return err
			}
			aggr.groupAggrs[group][n] = nv
		}
		aggr.groupCounts[group]++
	} else {
		for n, rv := range row {
			fv, err := toFloat64(rv)
			if err != nil {
				return fmt.Errorf("[%d] %w row : %s", n, ErrInvalid, err)
			}
			nv, err := aggr.aggFunc(aggr, aggr.aggrs[n], fv)
			if err != nil {
				return err
			}
			aggr.aggrs[n] = nv
		}
		aggr.counts++
	}

	return nil
}

// AggregateMap aggregates a map of data using the aggregator.
func (aggr *aggrImpl) AggregateMap(m map[string]any) error {
	row := make([]any, 0, len(aggr.columns))
	for _, arg := range aggr.args {
		var v any
		ok := false
		switch arg {
		case "*":
			ok = true
			v = nil
		default:
			v, ok = m[arg]
		}
		if !ok {
			return fmt.Errorf("column (%s) %w in map", arg, ErrNotFound)
		}
		row = append(row, v)
	}
	return aggr.AggregateRow(row)
}

// Aggregate aggregates a map or an array. The map represents a row of data, and the array is a
// list of rows. If grouping is enabled, the array row must have a group value as the first element.
func (aggr *aggrImpl) Aggregate(v any) error {
	switch v := v.(type) {
	case []any:
		return aggr.AggregateRow(v)
	case map[string]any:
		return aggr.AggregateMap(v)
	case Row:
		return aggr.AggregateRow(v)
	case Map:
		return aggr.AggregateMap(v)
	default:
		return fmt.Errorf("%w type %T is not supported for aggregation", ErrInvalid, v)
	}
}

// Finalize finalizes the aggregation and returns the result.
func (aggr *aggrImpl) Finalize() (ResultSet, error) {
	rows := make([]Row, 0)
	if _, ok := aggr.GroupBy(); ok {
		for group, values := range aggr.groupAggrs {
			groupCnt, ok := aggr.groupCounts[group]
			if !ok {
				return nil, fmt.Errorf("group count %w for group %v", ErrNotFound, group)
			}
			row := make([]any, 0)
			row = append(row, group)
			for _, value := range values {
				fv, err := aggr.finalFunc(aggr, value, groupCnt)
				if err != nil {
					return nil, err
				}
				row = append(row, fv)
			}
			rows = append(rows, row)
		}
	} else {
		row := make([]any, 0)
		for _, value := range aggr.aggrs {
			fv, err := aggr.finalFunc(aggr, value, aggr.counts)
			if err != nil {
				return nil, err
			}
			row = append(row, fv)
		}
		rows = append(rows, row)
	}
	return NewResultSet(
		WithRows(rows),
		WithColumns(aggr.columns),
	), nil
}
