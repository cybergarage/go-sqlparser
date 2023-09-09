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

// AggregatorFunction represents a base aggregator function.
type AggregatorFunction struct {
	name   string
	values map[string]float64
}

// NewAggregatorFunctionWith returns a new base aggregator function with the specified name and aggregator.
func NewAggregatorFunctionWith(name string) FunctionExecutor {
	fn := &AggregatorFunction{
		name:   strings.ToUpper(name),
		values: map[string]float64{},
	}
	return fn
}

// Name returns the name of the function.
func (fn *AggregatorFunction) Name() string {
	return fn.name
}

// Type returns the type of the function.
func (fn *AggregatorFunction) Type() FunctionType {
	return FunctionTypeAggregator
}

// Execute returns the executed value with the specified arguments.
func (fn *AggregatorFunction) Execute(args ...any) (any, error) {
	if len(args) != 2 {
		return nil, newErrInvalidArguments(fn.name, args)
	}
	groupKey, ok := args[0].(string)
	if !ok {
		return nil, newErrInvalidArguments(fn.name, args)
	}

	argValue := args[1]
	lastValue, ok := fn.values[groupKey]
	if ok {
		switch fn.name {
		case CountFunctionName:
			lastValue = lastValue + 1
		default:
			var v float64
			err := safecast.ToFloat64(argValue, &v)
			if err != nil {
				return nil, err
			}
			switch fn.name {
			case MinFunctionName:
				if v < lastValue {
					lastValue = v
				}
			case MaxFunctionName:
				if lastValue < v {
					lastValue = v
				}
			case AvgFunctionName:
				lastValue = (lastValue + v) / 2
			case SumFunctionName:
				lastValue = lastValue + v
			default:
				return nil, newErrInvalidArguments(fn.name, argValue)
			}
		}
	} else {
		switch fn.name {
		case CountFunctionName:
			lastValue = 1
		default:
			var v float64
			err := safecast.ToFloat64(argValue, &v)
			if err != nil {
				return nil, err
			}
			lastValue = v
		}
	}

	fn.values[groupKey] = lastValue

	return lastValue, nil
}

// NewCountFunction returns a new count function.
func NewCountFunction() FunctionExecutor {
	return NewAggregatorFunctionWith(CountFunctionName)
}

// NewMaxFunction returns a new max function.
func NewMaxFunction() FunctionExecutor {
	return NewAggregatorFunctionWith(MaxFunctionName)
}

// NewMinFunction returns a new min function.
func NewMinFunction() FunctionExecutor {
	return NewAggregatorFunctionWith(MinFunctionName)
}

// NewSumFunction returns a new sum function.
func NewSumFunction() FunctionExecutor {
	return NewAggregatorFunctionWith(SumFunctionName)
}

// NewAvgFunction returns a new avg function.
func NewAvgFunction() FunctionExecutor {
	return NewAggregatorFunctionWith(AvgFunctionName)
}
