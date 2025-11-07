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
	// Math.
	case strings.HasPrefix(upperName, AbsFunctionName):
		ex = NewAbs(opts...)
	case strings.HasPrefix(upperName, FloorFunctionName):
		ex = NewFloor(opts...)
	case strings.HasPrefix(upperName, CeilFunctionName):
		ex = NewCeil(opts...)
	case strings.HasPrefix(upperName, RoundFunctionName):
		ex = NewRound(opts...)
	case strings.HasPrefix(upperName, SqrtFunctionName):
		ex = NewSqrt(opts...)
	case strings.HasPrefix(upperName, LogFunctionName):
		ex = NewLog(opts...)
	case strings.HasPrefix(upperName, Log10FunctionName):
		ex = NewLog10(opts...)
	case strings.HasPrefix(upperName, ExpFunctionName):
		ex = NewExp(opts...)
	case strings.HasPrefix(upperName, PowerFunctionName):
		ex = NewPower(opts...)
	case strings.HasPrefix(upperName, ModFunctionName):
		ex = NewMod(opts...)
	case strings.HasPrefix(upperName, SinFunctionName):
		ex = NewSin(opts...)
	case strings.HasPrefix(upperName, CosFunctionName):
		ex = NewCos(opts...)
	case strings.HasPrefix(upperName, TanFunctionName):
		ex = NewTan(opts...)
	case strings.HasPrefix(upperName, RandFunctionName):
		ex = NewRand(opts...)
	case strings.HasPrefix(upperName, PiFunctionName):
		ex = NewPI(opts...)
	// String.
	case strings.HasPrefix(upperName, UpperFunctionName):
		ex = NewUpper(opts...)
	case strings.HasPrefix(upperName, LowerFunctionName):
		ex = NewLower(opts...)
	case strings.HasPrefix(upperName, TrimFunctionName):
		ex = NewTrim(opts...)
	// Time.
	case strings.HasPrefix(upperName, CurrentTimestampFunctionName):
		ex = NewCurrentTimestamp(opts...)
	case strings.HasPrefix(upperName, NowFunctionName):
		ex = NewNow(opts...)
	// Airthmetic.
	case strings.EqualFold(upperName, AddOperatorID):
		ex = NewAddOperator(opts...)
	case strings.EqualFold(upperName, SubOperatorID):
		ex = NewSubOperator(opts...)
	case strings.EqualFold(upperName, MulOperatorID):
		ex = NewMulOperator(opts...)
	case strings.EqualFold(upperName, DivOperatorID):
		ex = NewDivOperator(opts...)
	case strings.EqualFold(upperName, ModOperatorID):
		ex = NewModOperator(opts...)
	default:
		return nil, NewErrNotSupportedFunction(name)
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
