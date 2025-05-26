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
	"fmt"
)

// NewArithFunctionWith returns a new arithmetic function with the given operator and implementation.
func NewArithFunctionFor(ope string) (Executor, error) {
	switch ope {
	case AddFunctionID:
		return NewAddFunction(), nil
	case SubFunctionID:
		return NewSubFunction(), nil
	case MulFunctionID:
		return NewMulFunction(), nil
	case DivFunctionID:
		return NewDivFunction(), nil
	case ModFunctionID:
		return NewModFunction(), nil
	default:
		return nil, fmt.Errorf("unknown arithmetic operator: %s", ope)
	}
}

// NewAddFunction returns a new add function.
func NewAddFunction() Executor {
	return NewArithFunctionWith(
		AddFunctionID,
		func(v1, v2 float64) (float64, error) {
			return (v1 + v2), nil
		},
	)
}

// NewSubFunction returns a new sub function.
func NewSubFunction() Executor {
	return NewArithFunctionWith(
		SubFunctionID,
		func(v1, v2 float64) (float64, error) {
			return (v1 - v2), nil
		},
	)
}

// NewMulFunction returns a new multiple function.
func NewMulFunction() Executor {
	return NewArithFunctionWith(
		MulFunctionID,
		func(v1, v2 float64) (float64, error) {
			return (v1 * v2), nil
		},
	)
}

// NewDivFunction returns a new division function.
func NewDivFunction() Executor {
	return NewArithFunctionWith(
		DivFunctionID,
		func(v1, v2 float64) (float64, error) {
			return (v1 / v2), nil
		},
	)
}

// NewModFunction returns a new mod function.
func NewModFunction() Executor {
	return NewArithFunctionWith(
		ModFunctionID,
		func(v1, v2 float64) (float64, error) {
			return (float64(int(v1) % int(v2))), nil
		},
	)
}
