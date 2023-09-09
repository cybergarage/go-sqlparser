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

// CastFunc returns the latest aggregated value.
type CastFunc = func(any) (any, error)

// BaseCastFunction represents a base cast function.
type BaseCastFunction struct {
	name string
	cast CastFunc
}

// NewBaseCastFunctionWith returns a new base cast function with the specified name and cast.
func NewBaseCastFunctionWith(name string, cast CastFunc) *BaseCastFunction {
	fn := &BaseCastFunction{
		name: strings.ToUpper(name),
		cast: cast,
	}
	return fn
}

// Name returns the name of the function.
func (fn *BaseCastFunction) Name() string {
	return fn.name
}

// Type returns the type of the function.
func (fn *BaseCastFunction) Type() FunctionType {
	return FunctionTypeCast
}

// Execute returns the executed value with the specified arguments.
func (fn *BaseCastFunction) Execute(args ...any) (any, error) {
	if len(args) != 1 {
		return nil, newErrInvalidArguments(fn.name, args)
	}
	return fn.cast(args[0])
}
