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

import (
	"fmt"
)

// WithResultSetRowsOf returns a resultset option to set the rows from a given value.
func WithResultSetRowsOf(v any) ResultSetOption {
	return func(r *resultset) error {
		var rows []Row
		switch v := v.(type) {
		case []Row:
			rows = v
		case []map[string]any:
			rows = make([]Row, len(v))
			for i, row := range v {
				rows[i] = NewRow(
					WithRowObject(row),
					WithRowSchema(r.schema),
				)
			}
		default:
			return fmt.Errorf("unsupported type for rows: %T", v)
		}
		r.rows = rows
		return nil
	}
}
