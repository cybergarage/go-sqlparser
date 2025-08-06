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

func TestGroupBySet(t *testing.T) {
	tests := []struct {
		groupBys []fn.GroupBy
	}{
		{groupBys: []fn.GroupBy{}},
		{groupBys: []fn.GroupBy{"column1"}},
		{groupBys: []fn.GroupBy{"column1", "column2"}},
		{groupBys: []fn.GroupBy{"column1", "column2", "column3"}},
	}

	for _, tt := range tests {
		t.Run("GroupBySet", func(t *testing.T) {
			groupBySet := fn.NewGroupBySet(tt.groupBys...)
			groupBys := fn.NewGroupBysFromGroupBySet(groupBySet)
			if len(groupBys) != len(tt.groupBys) {
				t.Errorf("Expected %d group bys, got %d", len(tt.groupBys), len(groupBys))
			}
		})
	}
}
