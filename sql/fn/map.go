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

// Map represents a map with string keys and values of any type.
type Map map[string]any

// NewMap creates a new Map with the given key-value pairs.
func NewMapWithRow(columns []string, row Row) Map {
	m := make(Map, len(columns))
	for i, column := range columns {
		if i < len(row) {
			m[column] = row[i]
		} else {
			m[column] = nil
		}
	}
	return m
}
