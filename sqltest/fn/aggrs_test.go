// Copyright (C) 2019 The go-postgresql Authors. All rights reserved.
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

package fn_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/cybergarage/go-safecast/safecast"
	"github.com/cybergarage/go-sqlparser/sql/fn"

	"github.com/cybergarage/go-logger/log"
)

func TestAggregators(t *testing.T) {
	log.EnableStdoutDebug(true)

	tests := []struct {
		groupBys          []string
		args              []string
		rows              [][]any
		expectedSumRows   [][]float64
		expectedAvgRows   [][]float64
		expectedMinRows   [][]float64
		expectedMaxRows   [][]float64
		expectedCountRows [][]float64
		expectedRowCount  int
	}{
		{
			groupBys: []string{},
			args:     []string{"foo"},
			rows: [][]any{
				{1},
				{2},
				{3},
			},
			expectedSumRows:   [][]float64{{6}},
			expectedAvgRows:   [][]float64{{2}},
			expectedMinRows:   [][]float64{{1}},
			expectedMaxRows:   [][]float64{{3}},
			expectedCountRows: [][]float64{{3}},
			expectedRowCount:  1,
		},
		{
			groupBys: []string{},
			args:     []string{"foo"},
			rows: [][]any{
				{1},
				{2},
				{3},
				{4},
			},
			expectedSumRows:   [][]float64{{10}},
			expectedAvgRows:   [][]float64{{2.5}},
			expectedMinRows:   [][]float64{{1}},
			expectedMaxRows:   [][]float64{{4}},
			expectedCountRows: [][]float64{{4}},
			expectedRowCount:  1,
		},
		{
			groupBys: []string{"bar"},
			args:     []string{"foo"},
			rows: [][]any{
				{1, 1},
				{2, 2},
				{3, 3},
				{4, 4},
			},
			expectedSumRows:   [][]float64{{1, 1}, {2, 2}, {3, 3}, {4, 4}},
			expectedAvgRows:   [][]float64{{1, 1}, {2, 2}, {3, 3}, {4, 4}},
			expectedMinRows:   [][]float64{{1, 1}, {2, 2}, {3, 3}, {4, 4}},
			expectedMaxRows:   [][]float64{{1, 1}, {2, 2}, {3, 3}, {4, 4}},
			expectedCountRows: [][]float64{{1, 1}, {2, 1}, {3, 1}, {4, 1}},
			expectedRowCount:  4,
		},
		{
			groupBys: []string{"bar"},
			args:     []string{"foo"},
			rows: [][]any{
				{1, 1},
				{2, 2},
				{3, 3},
				{4, 4},
				{1, 2},
				{2, 4},
				{3, 6},
				{4, 8},
			},
			expectedSumRows:   [][]float64{{1, 3}, {2, 6}, {3, 9}, {4, 12}},
			expectedAvgRows:   [][]float64{{1, 1.5}, {2, 3}, {3, 4.5}, {4, 6}},
			expectedMinRows:   [][]float64{{1, 1}, {2, 2}, {3, 3}, {4, 4}},
			expectedMaxRows:   [][]float64{{1, 2}, {2, 4}, {3, 6}, {4, 8}},
			expectedCountRows: [][]float64{{1, 2}, {2, 2}, {3, 2}, {4, 2}},
			expectedRowCount:  4,
		},
		{
			groupBys: []string{"foo", "bar"},
			args:     []string{"baz"},
			rows: [][]any{
				{1, 1, 1},
				{1, 2, 2},
				{2, 1, 3},
				{2, 2, 4},
				{3, 1, 5},
				{3, 2, 6},
				{4, 1, 7},
				{4, 2, 8},
				{1, 1, 9},
				{2, 2, 10},
				{3, 1, 11},
				{4, 2, 12},
			},
			expectedSumRows:   [][]float64{{1, 1, 10}, {1, 2, 2}, {2, 1, 3}, {2, 2, 14}, {3, 1, 16}, {3, 2, 6}, {4, 1, 7}, {4, 2, 20}},
			expectedAvgRows:   [][]float64{{1, 1, 5}, {1, 2, 2}, {2, 1, 3}, {2, 2, 7}, {3, 1, 8}, {3, 2, 6}, {4, 1, 7}, {4, 2, 10}},
			expectedMinRows:   [][]float64{{1, 1, 1}, {1, 2, 2}, {2, 1, 3}, {2, 2, 4}, {3, 1, 5}, {3, 2, 6}, {4, 1, 7}, {4, 2, 8}},
			expectedMaxRows:   [][]float64{{1, 1, 9}, {1, 2, 2}, {2, 1, 3}, {2, 2, 10}, {3, 1, 11}, {3, 2, 6}, {4, 1, 7}, {4, 2, 12}},
			expectedCountRows: [][]float64{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}, {2, 2, 2}, {3, 1, 2}, {3, 2, 1}, {4, 1, 1}, {4, 2, 2}},
			expectedRowCount:  8,
		},
	}

	for n, test := range tests {

		t.Run(fmt.Sprintf("%02d", n), func(t *testing.T) {

			aggrFuncs := []func() (fn.Aggregator, error){
				func() (fn.Aggregator, error) {
					return fn.NewSum(
						fn.WithSumGroupBys(test.groupBys...),
						fn.WithSumArguments(test.args),
					)
				},
				func() (fn.Aggregator, error) {
					return fn.NewAvg(
						fn.WithAvgGroupBys(test.groupBys...),
						fn.WithAvgArguments(test.args),
					)
				},
				func() (fn.Aggregator, error) {
					return fn.NewMin(
						fn.WithMinGroupBys(test.groupBys...),
						fn.WithMinArguments(test.args),
					)
				},
				func() (fn.Aggregator, error) {
					return fn.NewMax(
						fn.WithMaxGroupBys(test.groupBys...),
						fn.WithMaxArguments(test.args),
					)
				},
				func() (fn.Aggregator, error) {
					return fn.NewCount(
						fn.WithCountGroupBys(test.groupBys...),
						fn.WithCountArguments(test.args),
					)
				},
			}

			for _, aggrFunc := range aggrFuncs {

				// Aggregate

				testAggr, err := aggrFunc()
				if err != nil {
					t.Error(err)
					continue
				}

				t.Run(testAggr.Name(), func(t *testing.T) {

					err := testAggr.Reset(test.groupBys)
					if err != nil {
						t.Errorf("Error resetting %s: %v", testAggr.Name(), err)
						return
					}

					for _, row := range test.rows {
						if err := testAggr.Aggregate(row); err != nil {
							t.Errorf("Error adding row: %v", err)
							continue
						}
					}

					rs, err := testAggr.Finalize(fn.OrderByAsc)
					if err != nil {
						t.Errorf("Error finalizing Sum: %v", err)
						return
					}

					rsRows := [][]any{}
					for rs.Next() {
						row, err := rs.Row()
						if err != nil {
							t.Errorf("Error getting row: %v", err)
							continue
						}
						rsRows = append(rsRows, row)
					}

					sort.Slice(rsRows, func(i, j int) bool {
						var ii, ij int
						if err := safecast.ToInt(rsRows[i][0], &ii); err != nil {
							return false
						}
						if err := safecast.ToInt(rsRows[j][0], &ij); err != nil {
							return false
						}
						return ii < ij
					})

					if len(rsRows) != test.expectedRowCount {
						t.Errorf("Expected %d rows, got %d", test.expectedRowCount, len(rsRows))
						return
					}

					// Compare the result set with the expected rows

					var expectedRows [][]float64

					switch testAggr.(type) {
					case *fn.Sum:
						expectedRows = test.expectedSumRows
					case *fn.Avg:
						expectedRows = test.expectedAvgRows
					case *fn.Min:
						expectedRows = test.expectedMinRows
					case *fn.Max:
						expectedRows = test.expectedMaxRows
					case *fn.Count:
						expectedRows = test.expectedCountRows
					default:
						t.Errorf("Unexpected fn.type: %T", testAggr)
						return
					}

					for n, expectedRow := range expectedRows {
						rsRow := rsRows[n]
						if len(rsRow) != len(expectedRow) {
							t.Errorf("%s(%v): Expected %d columns, got %d", testAggr.Name(), test.rows, len(expectedRow), len(rsRow))
							continue
						}
						for i, expectedValue := range expectedRow {
							var rsRowValue float64
							if err := safecast.ToFloat64(rsRow[i], &rsRowValue); err != nil {
								t.Errorf("Error converting row value to float64: %v", err)
								continue
							}
							if rsRowValue != expectedValue {
								t.Errorf("%s(%v): Expected %v, got %v", testAggr.Name(), test.rows, expectedRow, rsRow)
								continue
							}
						}
					}
				})
			}
		})
	}
}
