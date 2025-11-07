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
	"strings"

	"github.com/cybergarage/go-safecast/safecast"
)

const (
	groupByDelimiter = ","
)

// GroupBy represents a group by clause in a SQL query.
type GroupBy = string

// GroupBySet represents a set of group by clauses.
type GroupBySet = string

// NewGroupBy creates a new GroupBy instance from a string.
func NewGroupBy(s string) GroupBy {
	return GroupBy(s)
}

// NewGroupBySet creates a new GroupBySet instance from a list of GroupBy instances.
func NewGroupBySet(groupBys ...GroupBy) GroupBySet {
	if len(groupBys) == 0 {
		return ""
	}
	return GroupBySet(strings.Join(groupBys, groupByDelimiter))
}

// NewGroupBySetKey creates a GroupBySet from a slice of GroupBy and a slice of group keys.
func NewGroupBySetKey(groupBys []GroupBy, groupKeys []any) (GroupBySet, error) {
	if len(groupBys) != len(groupKeys) {
		return "", fmt.Errorf("groupBys and groupKeys must have the same length: %d != %d", len(groupBys), len(groupKeys))
	}
	groupBySet := []GroupBy{}
	for _, groupKey := range groupKeys {
		var groupKeyStr string
		err := safecast.ToString(groupKey, &groupKeyStr)
		if err != nil {
			return "", err
		}
		groupBySet = append(groupBySet, NewGroupBy(groupKeyStr))
	}
	return NewGroupBySet(groupBySet...), nil
}

// NewGroupBysFromGroupBySet converts a GroupBySet to a slice of GroupBy instances.
func NewGroupBysFromGroupBySet(groupBySet GroupBySet) []GroupBy {
	if len(groupBySet) == 0 {
		return []GroupBy{}
	}
	return strings.Split(string(groupBySet), groupByDelimiter)
}
