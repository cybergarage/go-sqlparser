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

type rows struct {
	rows []Row
}

// RowsOption represents a functional option for rows.
type RowsOption func(*rows) error

// WithRows sets the rows for the rows.
func WithRows(v []Row) RowsOption {
	return func(r *rows) error {
		r.rows = v
		return nil
	}
}

// NewRows creates a new rows instance with the opt
func NewRows(opts ...RowsOption) ([]Row, error) {
	r := &rows{
		rows: nil,
	}
	for _, opt := range opts {
		if err := opt(r); err != nil {
			return nil, err
		}
	}
	return r.rows, nil
}
