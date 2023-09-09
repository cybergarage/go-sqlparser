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
)

// AggregatorFunc returns the latest aggregated value.
type AggregatorFunc = func(string, any) (any, error)

// BaseAggregatorFunction represents a base aggregator function.
type BaseAggregatorFunction struct {
	name      string
	aggregate AggregatorFunc
}

// NewBaseAggregatorFunctionWith returns a new base aggregator function with the specified name and aggregator.
func NewBaseAggregatorFunctionWith(name string, aggregator AggregatorFunc) *BaseAggregatorFunction {
	fn := &BaseAggregatorFunction{
		name:      strings.ToUpper(name),
		aggregate: aggregator,
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
	return fn.aggregate(groupKey, args[1])
}