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

// MathFunc represents an math function.
type MathFunc func([]float64) (any, error)

// MathResultSet represents a result set of an math function.
type MathResultSet map[any]float64

// mathFunction represents a base math function.
type mathFunction struct {
	*execImpl

	executor MathFunc
}

// NewMathFunctionWith returns a new base math function with the specified name and math.
func NewMathFunctionWith(name string, mathFn MathFunc, opts ...ExecutorOption) Executor {
	fn := &mathFunction{
		execImpl: nil,
		executor: mathFn,
	}
	fn.execImpl = NewExecutor(
		WithExecutorName(name),
		WithExecutorType(MathFunction),
		WithExecutorFunc(fn.execute),
	)
	for _, opt := range opts {
		opt(fn.execImpl)
	}
	return fn
}

// execute returns the executed value with the specified arguments.
func (fn *mathFunction) execute(args ...any) (any, error) {
	fargs := make([]float64, 0, len(args))
	for _, arg := range args {
		var fv float64
		err := safecast.ToFloat64(arg, &fv)
		if err != nil {
			return nil, newErrInvalidArguments(fn.name, args...)
		}
		fargs = append(fargs, fv)
	}
	return fn.executor(fargs)
}
