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
	"strings"
)

// executorFunc is a function type that takes any number of arguments and returns a value and an error.
type executorFunc func(...any) (any, error)

// execImpl represents a base math function.
type execImpl struct {
	name    string
	columns []string
	t       FunctionType
	fn      executorFunc
}

// NewMathFunctionWith returns a new base math function with the specified name and math.
func newExecWith(name string, t FunctionType) *execImpl {
	return &execImpl{
		name: strings.ToUpper(name),
		t:    t,
		fn: func(args ...any) (any, error) {
			return nil, fmt.Errorf("function %s not implemented", name)
		},
	}
}

// Name returns the name of the function.
func (ex *execImpl) Name() string {
	return ex.name
}

// Type returns the type of the function.
func (ex *execImpl) Type() FunctionType {
	return ex.t
}

func (ex *execImpl) ExecuteRow(args ...any) (any, error) {
	return ex.fn(args...)
}

func (ex *execImpl) ExecuteMap(m map[string]any) (any, error) {
	row := make([]any, 0, len(ex.columns))
	for _, colum := range ex.columns {
		value, ok := m[colum]
		if !ok {
			return nil, fmt.Errorf("%w column %s not found in map", ErrNotFound, colum)
		}
		row = append(row, value)
	}
	return ex.ExecuteRow(row...)
}

func (ex *execImpl) Execute(v any) (any, error) {
	switch v := v.(type) {
	case []any:
		return ex.ExecuteRow(v)
	case map[string]any:
		return ex.ExecuteMap(v)
	default:
		return ex.ExecuteRow(v)
	}
}
