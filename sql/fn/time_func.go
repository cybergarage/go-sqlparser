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
	"time"
)

// NewCurrentTimestampFunction returns a new current timestamp function.
func NewCurrentTimestampFunction(opts ...ExecutorOption) Executor {
	return NewTimeFunctionWith(
		CurrentTimestampFunctionName,
		func(args []string) (any, error) {
			return time.Now(), nil
		},
		opts...,
	)
}

// NewTimeFunctionWith returns a new time function with the specified name and executor.
func NewNowFunction(opts ...ExecutorOption) Executor {
	return NewTimeFunctionWith(
		NowFunctionName,
		func(args []string) (any, error) {
			return time.Now(), nil
		},
		opts...,
	)
}
