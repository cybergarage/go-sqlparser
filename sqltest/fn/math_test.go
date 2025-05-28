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
	"testing"

	"github.com/cybergarage/go-sqlparser/sql/fn"
)

func TestMathExecutors(t *testing.T) {
	tests := []struct {
		fn     fn.Executor
		arg    any
		result any
	}{
		{fn.NewAbsFunction(), float64(-1), float64(1)},
		{fn.NewFloorFunction(), float64(5.95), int(5)},
		{fn.NewCeilFunction(), float64(5.95), int(6)},
		{fn.NewRoundFunction(), float64(5.95), int(6)},
		{fn.NewLogFunction(), float64(100), float64(4.605170185988092)},
		{fn.NewLog10Function(), float64(100), float64(2)},
		{fn.NewSqrtFunction(), float64(16), float64(4)},
		{fn.NewExpFunction(), float64(1), float64(2.718281828459045)},
	}

	for _, test := range tests {
		r, err := test.fn.Execute(test.arg)
		if err != nil {
			t.Error(err)
			return
		}
		if r != test.result {
			t.Errorf("The %s value (%v) is not (%v)", test.fn.Name(), r, test.result)
		}
	}
}
