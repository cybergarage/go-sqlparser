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

type TimeFunc func([]string) (any, error)

type timeFunction struct {
	*execImpl

	executor TimeFunc
}

// NewTimeFunctionWith returns a new time function with the specified name and executor.
func NewTimeFunctionWith(name string, t TimeFunc, opts ...ExecutorOption) *timeFunction {
	fn := &timeFunction{
		execImpl: nil,
		executor: t,
	}
	fn.execImpl = NewExecutor(
		WithExecutorName(name),
		WithExecutorType(TimeFunction),
		WithExecutorFunc(fn.execute),
	)
	for _, opt := range opts {
		opt(fn.execImpl)
	}
	return fn
}

// execute returns the executed value with the specified arguments.
func (fn *timeFunction) execute(args ...any) (any, error) {
	fargs := make([]string, 0, len(args))
	for _, arg := range args {
		var str string
		err := safecast.ToString(arg, &str)
		if err != nil {
			return nil, newErrInvalidArguments(fn.name, args...)
		}
		fargs = append(fargs, str)
	}
	return fn.executor(fargs)
}
