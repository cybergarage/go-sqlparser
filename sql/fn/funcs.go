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

const (
	// Math.
	AbsFunctionName   = "ABS"
	CeilFunctionName  = "CEIL"
	FloorFunctionName = "FLOOR"
	// Aggregator.
	MaxFunctionName   = "MAX"
	MinFunctionName   = "MIN"
	SumFunctionName   = "SUM"
	AvgFunctionName   = "AVG"
	CountFunctionName = "COUNT"
)

// NewFunctionExecutorForName returns a function executor with the specified name.
func NewFunctionExecutorForName(name string) (FunctionExecutor, error) {
	switch name {
	case MaxFunctionName:
		return NewMaxFunction(), nil
	case MinFunctionName:
		return NewMinFunction(), nil
	case SumFunctionName:
		return NewSumFunction(), nil
	case AvgFunctionName:
		return NewAvgFunction(), nil
	case CountFunctionName:
		return NewCountFunction(), nil
	case AbsFunctionName:
		return NewAbsFunction(), nil
	case FloorFunctionName:
		return NewFloorFunction(), nil
	case CeilFunctionName:
		return NewCeilFunction(), nil
	}
	return nil, newErrNotSupportedFunction(name)
}

// ExecuteFunction returns the executed value with the specified arguments.
func ExecuteFunction(name string, args ...any) (any, error) {
	fn, err := NewFunctionExecutorForName(name)
	if err != nil {
		return nil, err
	}
	return fn.Execute(args...)
}
