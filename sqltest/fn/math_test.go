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

package fn_test

import (
	"math"
	"testing"

	"github.com/cybergarage/go-sqlparser/sql/fn"
)

func TestMathExecutors(t *testing.T) {
	tests := []struct {
		fn     fn.Executor
		arg    []any
		result any
	}{
		{fn.NewAbsFunction(), []any{float64(-1)}, float64(1)},
		{fn.NewFloorFunction(), []any{float64(5.95)}, int(5)},
		{fn.NewCeilFunction(), []any{float64(5.95)}, int(6)},
		{fn.NewRoundFunction(), []any{float64(5.95)}, int(6)},
		{fn.NewLogFunction(), []any{float64(100)}, math.Log(float64(100))},
		{fn.NewLog10Function(), []any{float64(100)}, math.Log10(float64(100))},
		{fn.NewSqrtFunction(), []any{float64(16)}, math.Sqrt(float64(16))},
		{fn.NewExpFunction(), []any{float64(1)}, math.Exp(float64(1))},
		{fn.NewPowerFunction(), []any{float64(2), float64(3)}, float64(8)},
		{fn.NewModFunction(), []any{float64(10), float64(3)}, float64(1)},
		{fn.NewSinFunction(), []any{math.Pi}, math.Sin(math.Pi)},
		{fn.NewCosFunction(), []any{math.Pi}, math.Cos(math.Pi)},
		{fn.NewTanFunction(), []any{math.Pi}, math.Tan(math.Pi)},
		{fn.NewRandFunction(), nil, nil}, // Random function does not have a fixed result
	}

	for _, test := range tests {
		r, err := test.fn.Execute(test.arg)
		if err != nil {
			t.Error(err)
			return
		}
		if test.result == nil {
			continue // Skip comparison for random function
		}
		if r != test.result {
			t.Errorf("The %s value (%v) is not (%v)", test.fn.Name(), r, test.result)
		}
	}
}
