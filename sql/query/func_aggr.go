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
	"strings"

	"github.com/cybergarage/go-safecast/safecast"
)

// AggregateFunc represents an aggregator function.
type AggregateFunc func(int, float64, float64) float64

// AggregateResultSet represents a result set of an aggregator function.
type AggregateResultSet map[any]float64

// AggregateFunction represents a base aggregator function.
type AggregateFunction struct {
	name       string
	aggregator AggregateFunc
	count      int
	values     AggregateResultSet
}

// NewAggregateFunctionWith returns a new base aggregator function with the specified name and aggregator.
func NewAggregateFunctionWith(name string, fn AggregateFunc) *AggregateFunction {
	return &AggregateFunction{
		name:       strings.ToUpper(name),
		aggregator: fn,
		count:      0,
		values:     AggregateResultSet{},
	}
}

// Name returns the name of the function.
func (fn *AggregateFunction) Name() string {
	return fn.name
}

// Type returns the type of the function.
func (fn *AggregateFunction) Type() FunctionType {
	return AggregateFunctionType
}

// Execute returns the executed value with the specified arguments.
func (fn *AggregateFunction) Execute(args ...any) (any, error) {
	if len(args) != 2 {
		return nil, newErrInvalidArguments(fn.name, args)
	}
	groupKey := args[0]

	var argValue float64
	err := safecast.ToFloat64(args[1], &argValue)
	if err != nil {
		return nil, err
	}

	var lastValue float64
	currValue, ok := fn.values[groupKey]
	if ok {
		lastValue = fn.aggregator(fn.count, currValue, argValue)
	} else {
		lastValue = fn.aggregator(fn.count, 0.0, argValue)
	}
	fn.values[groupKey] = lastValue
	fn.count++

	return lastValue, nil
}

// Aggregate returns the latest aggregated value.
func (fn *AggregateFunction) ResultSet() AggregateResultSet {
	switch fn.name {
	case AvgFunctionName:
		for groupKey, value := range fn.values {
			fn.values[groupKey] = value / float64(fn.count)
		}
	}
	return fn.values
}

// NewCountFunction returns a new count function.
func NewCountFunction() *AggregateFunction {
	return NewAggregateFunctionWith(
		CountFunctionName,
		func(count int, currValue float64, argValue float64) float64 {
			return float64(count + 1)
		},
	)
}

// NewMaxFunction returns a new max function.
func NewMaxFunction() *AggregateFunction {
	return NewAggregateFunctionWith(
		MaxFunctionName,
		func(count int, currValue float64, argValue float64) float64 {
			if count == 0 {
				return argValue
			}
			if currValue < argValue {
				return argValue
			}
			return currValue
		},
	)
}

// NewMinFunction returns a new min function.
func NewMinFunction() *AggregateFunction {
	return NewAggregateFunctionWith(
		MinFunctionName,
		func(count int, currValue float64, argValue float64) float64 {
			if count == 0 {
				return argValue
			}
			if argValue < currValue {
				return argValue
			}
			return currValue
		},
	)
}

// NewSumFunction returns a new sum function.
func NewSumFunction() *AggregateFunction {
	return NewAggregateFunctionWith(
		SumFunctionName,
		func(count int, currValue float64, argValue float64) float64 {
			return currValue + argValue
		},
	)
}

// NewAvgFunction returns a new avg function.
func NewAvgFunction() *AggregateFunction {
	return NewAggregateFunctionWith(
		AvgFunctionName,
		func(count int, currValue float64, argValue float64) float64 {
			return currValue + argValue
		},
	)
}
