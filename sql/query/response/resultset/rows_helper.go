// Copyright (C) 2024 The go-mysql Authors. All rights reserved.
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

package resultset

// NewRowsFromMapRows creates a new rows instance with the given map rows.
func NewRowsFromMapRows(mapRows MapRows) ([]Row, error) {
	rsRows := []Row{}
	for _, rowMap := range mapRows {
		rsRow := NewRow(
			WithRowObject(rowMap),
		)
		rsRows = append(rsRows, rsRow)
	}

	return rsRows, nil
}
