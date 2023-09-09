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

// BaseAggregatorFunction represents a base aggregator function.
type BaseAggregatorFunction struct {
	name   string
	values map[string]any
}

// NewAggregatorFunctionWith returns a new base aggregator function with the specified name and aggregator.
func NewAggregatorFunctionWith(name string) FunctionExecutor {
	fn := &BaseAggregatorFunction{
		name:   strings.ToUpper(name),
		values: make(map[string]any),
	}
	return fn
}

// Name returns the name of the function.
func (fn *BaseAggregatorFunction) Name() string {
	return fn.name
}

// Type returns the type of the function.
func (fn *BaseAggregatorFunction) Type() FunctionType {
	return FunctionTypeAggregator
}

// Execute returns the executed value with the specified arguments.
func (fn *BaseAggregatorFunction) Execute(args ...any) (any, error) {
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
			lastValue = 1
		default:
			switch argValue.(type) {
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
				var v int64
				err := safecast.ToInt64(argValue, &v)
				if err != nil {
					return nil, err
				}
				lastValue = v
			case float32, float64:
				var v float64
				err := safecast.ToFloat64(argValue, &v)
				if err != nil {
					return nil, err
				}
				lastValue = v
			default:
				return nil, newErrInvalidArguments(fn.name, argValue)
			}
		}
	} else {
		switch fn.name {
		case CountFunctionName:
			lv, ok := lastValue.(int64)
			if ok {
				lastValue = lv + 1
			}
		default:
			switch lv := lastValue.(type) {
			case int64:
				var v int64
				err := safecast.ToInt64(argValue, &v)
				if err != nil {
					return nil, err
				}
				switch fn.name {
				case MinFunctionName:
					if v < lv {
						lastValue = v
					}
				case MaxFunctionName:
					if lv < v {
						lastValue = v
					}
				case AvgFunctionName:
					lastValue = (lv + v) / 2
				case SumFunctionName:
					lastValue = lv + v
				default:
					return nil, newErrInvalidArguments(fn.name, argValue)
				}
			case float64:
				var v float64
				err := safecast.ToFloat64(argValue, &v)
				if err != nil {
					return nil, err
				}
				lastValue = lv + v
			default:
				return nil, newErrInvalidArguments(fn.name, argValue)
			}
		}
	}

	fn.values[groupKey] = lastValue

	return lastValue, nil
}

// func NewMaxFunction() FunctionExecutor {
// 	return NewBaseAggregatorFunctionWith(MaxFunctionName)
// }
