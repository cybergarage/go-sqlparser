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
	"sort"
	"strconv"

	"github.com/cybergarage/go-safecast/safecast"
	"github.com/cybergarage/go-sqlparser/sql/errors"
)

type row struct {
	object map[string]any
	schema Schema
	values []any
}

// RowOption represents a functional option for Row.
type RowOption func(*row)

// WithRowObject returns a functional option for row object.
func WithRowObject(object map[string]any) RowOption {
	return func(row *row) {
		row.object = object
	}
}

// WithRowSchema returns a functional option for row schema.
func WithRowSchema(schema Schema) RowOption {
	return func(row *row) {
		row.schema = schema
	}
}

// WithRowValues returns a functional option for row values.
func WithRowValues(values []any) RowOption {
	return func(row *row) {
		row.values = values
	}
}

// NewRow returns a new resultsetRow.
func NewRow(opts ...RowOption) Row {
	row := &row{
		object: nil,
		schema: nil,
		values: nil,
	}
	for _, opt := range opts {
		opt(row)
	}
	return row
}

// Objects returns the row objects.
func (row *row) Object() map[string]any {
	if row.object != nil {
		return row.object
	}
	if row.values == nil {
		return map[string]any{}
	}
	// Generate object from the values.
	names := []string{}
	if row.schema != nil {
		for _, column := range row.schema.Columns() {
			names = append(names, column.Name())
		}
	} else {
		for n := range len(row.values) {
			names = append(names, strconv.Itoa((n + 1)))
		}
	}
	row.object = make(map[string]any)
	for n, name := range names {
		var v any
		if n < len(row.values) {
			v = row.values[n]
		} else {
			v = nil
		}
		row.object[name] = v
	}
	return row.object
}

// Values returns the row values.
func (row *row) Values() []any {
	if row.values != nil {
		return row.values
	}
	if row.object == nil {
		return []any{}
	}
	// Generate values from the object.
	var names []string
	if row.schema != nil {
		for _, column := range row.schema.Columns() {
			names = append(names, column.Name())
		}
	} else {
		var keys []string
		for key := range row.object {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		names = append(names, keys...)
	}
	row.values = []any{}
	for _, name := range names {
		v, ok := row.object[name]
		if ok {
			row.values = append(row.values, v)
		} else {
			row.values = append(row.values, nil)
		}
	}
	return row.values
}

// ValueAt returns the row value at the specified index.
func (row *row) ValueAt(index int) (any, error) {
	values := row.Values()
	if len(values) <= index {
		return nil, errors.ErrNotExist
	}
	return values[index], nil
}

// ValueBy returns the row value by the specified name.
func (row *row) ValueBy(name string) (any, error) {
	obj := row.Object()
	if obj == nil {
		return nil, errors.ErrNotExist
	}
	v, ok := obj[name]
	if !ok {
		return nil, errors.ErrNotExist
	}
	return v, nil
}

// Scan scans the values.
func (row *row) Scan(tos ...any) error {
	values := row.Values()
	for n, to := range tos {
		if len(values) <= n {
			return errors.ErrNotExist
		}
		v := values[n]
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

// ScanBy scans the value by the specified name.
func (row *row) ScanBy(name string, to any) error {
	v, err := row.ValueBy(name)
	if err != nil {
		return err
	}
	return safecast.To(v, to)
}
