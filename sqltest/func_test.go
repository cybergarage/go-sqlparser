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

package sqltest

import (
	"testing"

	"github.com/cybergarage/go-sqlparser/sql/query"
)

func TestAggregateFunctions(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15}

	groupKey := ""
	tests := []struct {
		fn     *query.AggregateFunction
		result int
	}{
		{query.NewAvgFunction(), 6},
		{query.NewCountFunction(), len(values)},
		{query.NewMaxFunction(), 15},
		{query.NewMinFunction(), 1},
		{query.NewSumFunction(), 60},
	}

	for _, test := range tests {
		for _, value := range values {
			_, err := test.fn.Execute(groupKey, value)
			if err != nil {
				t.Error(err)
				return
			}
		}
		rs := test.fn.ResultSet()
		r, ok := rs[groupKey]
		if !ok {
			t.Errorf("The %s result is not found", test.fn.Name())
			return
		}
		if int(r) != test.result {
			t.Errorf("The %s value (%d) is not (%d)", test.fn.Name(), int(r), test.result)
		}
	}
}

func TestMathFunctions(t *testing.T) {
	tests := []struct {
		fn     *query.MathFunction
		arg    any
		result any
	}{
		{query.NewAbsFunction(), float64(-1), float64(1)},
		{query.NewFloorFunction(), float64(5.95), int64(5)},
		{query.NewCeilFunction(), float64(5.95), int64(6)},
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
