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

// FunctionType represents a function type.
type FunctionType int

const (
	MathFunctionType FunctionType = iota
	AggregatorFunctionType
	CastFunctionType
)

// FunctionExecutor represents a function executor interface.
type FunctionExecutor interface {
	// Name returns the name of the function.
	Name() string
	// Type returns the type of the function.
	Type() FunctionType
	// Execute returns the executed value with the specified arguments.
	Execute(...any) (any, error)
}

// Function represents a column.
type Function struct {
	name string
	ArgumentList
}

// NewFunctionWith returns a column instance.
func NewFunctionWith(name string, args ...Argument) *Function {
	fn := &Function{
		name:         strings.ToUpper(name),
		ArgumentList: NewArgumentsWith(args...),
	}
	return fn
}

// Name returns the column name.
func (fn *Function) Name() string {
	return fn.name
}

// IsName returns true whether the column name is the specified one.
func (fn *Function) IsName(name string) bool {
	return fn.name == name
}

// SelectorString returns the selector string representation.
func (fn *Function) SelectorString() string {
	return fn.name + "(" + fn.ArgumentList.String() + ")"
}
