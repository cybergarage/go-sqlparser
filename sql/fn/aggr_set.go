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

// AggregatorSet is a collection of Aggregators.
// It provides methods to manage and operate on multiple aggregators.
type AggregatorSet []Aggregator

// NewAggregatorSet creates a new AggregatorSet with the given aggregators.
func NewAggregatorSetWith(aggrs []Aggregator) AggregatorSet {
	aggrSet := make(AggregatorSet, len(aggrs))
	copy(aggrSet, aggrs)
	return aggrSet
}

// NewAggregatorSetForNames creates a new AggregatorSet for the given names.
func NewAggregatorSetForNames(names []string, opts ...aggrOption) (AggregatorSet, error) {
	aggrSet := make(AggregatorSet, len(names))
	for i, name := range names {
		aggr, err := NewAggregatorForName(name, opts...)
		if err != nil {
			return nil, err
		}
		aggrSet[i] = aggr
	}
	return aggrSet, nil
}

// Reset resets all aggregators in the set.
func (aggrSet *AggregatorSet) Reset(opts ...any) error {
	for _, aggr := range *aggrSet {
		if err := aggr.Reset(opts...); err != nil {
			return err
		}
	}
	return nil
}

// Aggregate aggregates a row of data using all aggregators in the set.
func (aggrSet *AggregatorSet) Aggregate(row Row) error {
	for _, aggr := range *aggrSet {
		if err := aggr.Aggregate(row); err != nil {
			return err
		}
	}
	return nil
}

// Finalize finalizes the aggregation and returns the result set.
func (aggrSet *AggregatorSet) Finalize() ([]ResultSet, error) {
	resultSet := make([]ResultSet, len(*aggrSet))
	for i, aggr := range *aggrSet {
		result, err := aggr.Finalize()
		if err != nil {
			return nil, err
		}
		resultSet[i] = result
	}
	return resultSet, nil
}
