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
	"github.com/cybergarage/go-safecast/safecast"
)

// ArithFunc represents an arithmetic function.
type ArithFunc func(float64, float64) (float64, error)

// arithFunction represents a base arithmetic function.
type arithFunction struct {
	*execImpl
	operator ArithFunc
}

// NewArithFunctionWith returns a new base arithmetic function with the specified name and arithmetic.
func NewArithFunctionWith(name string, arithFn ArithFunc, opts ...ExecutorOption) Executor {
	fn := &arithFunction{
		execImpl: newExecWith(name, ArithOperator),
		operator: arithFn,
	}
	fn.fn = fn.execute
	for _, opt := range opts {
		opt(fn.execImpl)
	}
	return fn
}

// execute returns the executed value with the specified arguments.
func (fn *arithFunction) execute(args ...any) (any, error) {
	if len(args) != 2 {
		return nil, newErrInvalidArguments(fn.name, args)
	}
	var fv1 float64
	err := safecast.ToFloat64(args[0], &fv1)
	if err != nil {
		return nil, err
	}
	var fv2 float64
	err = safecast.ToFloat64(args[1], &fv2)
	if err != nil {
		return nil, err
	}
	return fn.operator(fv1, fv2)
}
