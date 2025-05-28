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

// NewArithOperatorFor returns a new arithmetic function with the given operator and implementation.
func NewArithOperatorFor(ope string) (Executor, error) {
	switch ope {
	case AddOperatorID:
		return NewAddOperator(), nil
	case SubOperatorID:
		return NewSubOperator(), nil
	case MulOperatorID:
		return NewMulOperator(), nil
	case DivOperatorID:
		return NewDivOperator(), nil
	case ModOperatorID:
		return NewModOperator(), nil
	default:
		return nil, fmt.Errorf("unknown arithmetic operator: %s", ope)
	}
}

// NewAddOperator returns a new add function.
func NewAddOperator() Executor {
	return NewArithFunctionWith(
		AddOperatorID,
		func(v1, v2 float64) (float64, error) {
			return (v1 + v2), nil
		},
	)
}

// NewSubOperator returns a new sub function.
func NewSubOperator() Executor {
	return NewArithFunctionWith(
		SubOperatorID,
		func(v1, v2 float64) (float64, error) {
			return (v1 - v2), nil
		},
	)
}

// NewMulOperator returns a new multiple function.
func NewMulOperator() Executor {
	return NewArithFunctionWith(
		MulOperatorID,
		func(v1, v2 float64) (float64, error) {
			return (v1 * v2), nil
		},
	)
}

// NewDivOperator returns a new division function.
func NewDivOperator() Executor {
	return NewArithFunctionWith(
		DivOperatorID,
		func(v1, v2 float64) (float64, error) {
			return (v1 / v2), nil
		},
	)
}

// NewModOperator returns a new mod function.
func NewModOperator() Executor {
	return NewArithFunctionWith(
		ModOperatorID,
		func(v1, v2 float64) (float64, error) {
			return (float64(int(v1) % int(v2))), nil
		},
	)
}
