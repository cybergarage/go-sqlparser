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
	"github.com/cybergarage/go-safecast/safecast"
	"github.com/cybergarage/go-sqlparser/sql/errors"
)

type row struct {
	values []any
}

// RowOptions represents a functional option for resultsetRow.
type RowOptions func(*row)

// WithRowValues returns a functional option for resultsetRow.
func WithRowValues(values []any) RowOptions {
	return func(row *row) {
		row.values = values
	}
}

// NewRow returns a new resultsetRow.
func NewRow(opts ...RowOptions) Row {
	row := &row{
		values: []any{},
	}
	for _, opt := range opts {
		opt(row)
	}
	return row
}

// Values returns the values.
func (row *row) Values() []any {
	return row.values
}

// ValueAt returns the value at the specified index.
func (row *row) ValueAt(index int) (any, error) {
	if len(row.values) <= index {
		return nil, errors.ErrNotExist
	}
	return row.values[index], nil
}

// Scan scans the values.
func (row *row) Scan(tos ...any) error {
	for i, to := range tos {
		if len(row.values) <= i {
			return errors.ErrNotExist
		}
		v := row.values[i]
		err := safecast.To(v, to)
		if err != nil {
			return err
		}
	}
	return nil
}

// ScanAt scans the value at the specified index.
func (row *row) ScanAt(index int, to any) error {
	v, err := row.ValueAt(index)
	if err != nil {
		return err
	}
	return safecast.To(v, to)
}
