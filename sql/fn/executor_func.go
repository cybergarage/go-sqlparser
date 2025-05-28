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
	"strings"
)

// ExecutorOption is an option for the function executor.
type ExecutorOption = execOption

// WithExecutorName returns an option to set the name of the executor.
func WithExecutorName(name string) ExecutorOption {
	return withExecName(name)
}

// WithExecutorArguments returns an option to set the arguments for the executor.
func WithExecutorArguments(args []string) ExecutorOption {
	return withExecArguments(args)
}

// NewExecutorForName returns a function executor with the specified name.
func NewExecutorForName(name string, opts ...ExecutorOption) (Executor, error) {
	upperName := strings.ToUpper(name)
	var ex Executor
	switch {
	case strings.HasPrefix(upperName, AbsFunctionName):
		ex = NewAbsFunction(opts...)
	case strings.HasPrefix(upperName, FloorFunctionName):
		ex = NewFloorFunction(opts...)
	case strings.HasPrefix(upperName, CeilFunctionName):
		ex = NewCeilFunction(opts...)
	case strings.HasPrefix(upperName, RoundFunctionName):
		ex = NewRoundFunction(opts...)
	case strings.HasPrefix(upperName, SqrtFunctionName):
		ex = NewSqrtFunction(opts...)
	case strings.HasPrefix(upperName, LogFunctionName):
		ex = NewLogFunction(opts...)
	case strings.HasPrefix(upperName, Log10FunctionName):
		ex = NewLog10Function(opts...)
	case strings.HasPrefix(upperName, ExpFunctionName):
		ex = NewExpFunction(opts...)
	case strings.HasPrefix(upperName, PowerFunctionName):
		ex = NewPowerFunction(opts...)
	case strings.HasPrefix(upperName, ModFunctionName):
		ex = NewModFunction(opts...)
	case strings.HasPrefix(upperName, SinFunctionName):
		ex = NewSinFunction(opts...)
	case strings.HasPrefix(upperName, CosFunctionName):
		ex = NewCosFunction(opts...)
	case strings.HasPrefix(upperName, TanFunctionName):
		ex = NewTanFunction(opts...)
	case strings.HasPrefix(upperName, RandFunctionName):
		ex = NewRandFunction(opts...)
	case strings.HasPrefix(upperName, PiFunctionName):
		ex = NewPiFunction(opts...)
	case strings.HasPrefix(upperName, CurrentTimestampFunctionName):
		ex = NewCurrentTimestampFunction(opts...)
	case strings.HasPrefix(upperName, NowFunctionName):
		ex = NewNowFunction(opts...)
	default:
		return nil, newErrNotSupportedFunction(name)
	}
	return ex, nil
}

// ExecuteFunction returns the executed value with the specified arguments.
func ExecuteFunction(name string, v any) (any, error) {
	fn, err := NewExecutorForName(name)
	if err != nil {
		return nil, err
	}
	return fn.Execute(v)
}
