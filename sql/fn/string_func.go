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

// NewUpperFunctionWith returns a new string function with the specified name and executor.
func NewUpper(opts ...ExecutorOption) Executor {
	return NewStringFunctionWith(
		UpperFunctionName,
		func(args []string) (any, error) {
			if len(args) != 1 {
				return nil, newErrInvalidArguments(UpperFunctionName, args)
			}
			return strings.ToUpper(args[0]), nil
		},
		opts...,
	)
}

// NewLowerFunctionWith returns a new string function with the specified name and executor.
func NewLower(opts ...ExecutorOption) Executor {
	return NewStringFunctionWith(
		LowerFunctionName,
		func(args []string) (any, error) {
			if len(args) != 1 {
				return nil, newErrInvalidArguments(LowerFunctionName, args)
			}
			return strings.ToLower(args[0]), nil
		},
		opts...,
	)
}

// NewTrimFunctionWith returns a new string function with the specified name and executor.
func NewTrim(opts ...ExecutorOption) Executor {
	return NewStringFunctionWith(
		TrimFunctionName,
		func(args []string) (any, error) {
			if len(args) != 1 {
				return nil, newErrInvalidArguments(TrimFunctionName, args)
			}
			return strings.TrimSpace(args[0]), nil
		},
		opts...,
	)
}
