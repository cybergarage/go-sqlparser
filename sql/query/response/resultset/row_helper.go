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

// NewRowFromObject creates a new row from the specified object.
func NewRowFromObject(schema Schema, objMap map[string]any) (Row, error) {
	selectors := schema.Selectors()
	rowObj := make(map[string]any, 0)
	for _, selector := range selectors {
		name := selector.Name()
		rowObj[name] = nil
		if fx, ok := selector.Function(); ok {
			if executor, err := fx.Executor(); err == nil {
				value, err := executor.Execute(objMap)
				if err != nil {
					return nil, err
				}
				rowObj[name] = value
				continue
			}
		}
		value, ok := objMap[name]
		if ok {
			rowObj[name] = value
		}
	}
	return NewRow(
		WithRowSchema(schema),
		WithRowObject(objMap),
	), nil
}
