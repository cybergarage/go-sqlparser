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
	"math"
)

// NewAbsFunction returns a new abs function.
func NewAbsFunction() Executor {
	return NewMathFunctionWith(
		AbsFunctionName,
		func(v float64) (any, error) {
			return math.Abs(v), nil
		},
	)
}

// NewFloorFunction returns a new floor function.
func NewFloorFunction() Executor {
	return NewMathFunctionWith(
		FloorFunctionName,
		func(v float64) (any, error) {
			return int(math.Floor(v)), nil
		},
	)
}

// NewCeilFunction returns a new ceil function.
func NewCeilFunction() Executor {
	return NewMathFunctionWith(
		CeilFunctionName,
		func(v float64) (any, error) {
			return int(math.Ceil(v)), nil
		},
	)
}
