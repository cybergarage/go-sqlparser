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

func TestAggregatorFunctions(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15}

	tests := []struct {
		function query.FunctionExecutor
		result   int
	}{
		{query.NewAvgFunction(), 6},
		{query.NewCountFunction(), len(values)},
		{query.NewMaxFunction(), 15},
		{query.NewMinFunction(), 1},
		{query.NewSumFunction(), 60},
	}

	for _, test := range tests {
		var lastValue float64
		for _, value := range values {
			lv, err := test.function.Execute("", value)
			if err != nil {
				t.Error(err)
				return
			}
			lastValue = lv.(float64)
		}
		if int(lastValue) != test.result {
			t.Errorf("The %s value (%d) is not (%d)", test.function.Name(), int(lastValue), test.result)
		}
	}
}
