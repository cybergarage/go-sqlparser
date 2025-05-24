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
	name string
	t    FunctionType
	fn   executorFunc
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

// Execute returns the executed value with the specified arguments.
func (ex *execImpl) Execute(args ...any) (any, error) {
	return ex.fn(args...)
}
