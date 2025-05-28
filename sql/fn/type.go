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

import "fmt"

// FunctionType represents a function type.
type FunctionType int

const (
	UnknownFunctionType FunctionType = iota
	MathFunction
	AggregateFunction
	CastFunction
	ArithOperator
	TimeFunction
)

// NewFunctionTypeForName returns a FunctionType based on the function name.
func NewFunctionTypeForName(name string) (FunctionType, error) {
	switch name {
	case AbsFunctionName, FloorFunctionName, CeilFunctionName, RoundFunctionName, SqrtFunctionName, LogFunctionName, Log10FunctionName, ExpFunctionName, PowerFunctionName, ModFunctionName, SinFunctionName, CosFunctionName, TanFunctionName, RandFunctionName, PiFunctionName:
		return MathFunction, nil
	case SumFunctionName, AvgFunctionName, CountFunctionName, MaxFunctionName, MinFunctionName:
		return AggregateFunction, nil
	case AddOperatorID, SubOperatorID, MulOperatorID, DivOperatorID, ModOperatorID:
		return ArithOperator, nil
	case CurrentTimestampFunctionName, NowFunctionName:
		return TimeFunction, nil
	default:
		return UnknownFunctionType, fmt.Errorf("unknown function type: %s", name)
	}
}
