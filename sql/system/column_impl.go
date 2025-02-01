// Copyright (C) 2025 The go-sqlparser Authors. All rights reserved.
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

package system

import (
	"fmt"

	"github.com/cybergarage/go-sqlparser/sql/query"
)

type column struct {
	name       string
	dataType   DataType
	constraint Constraint
}

type ColumnOption func(*column) error

// WithColumnName returns a column option to set the column name.
func WithColumnName(a any) ColumnOption {
	return func(c *column) error {
		var name string
		switch v := a.(type) {
		case string:
			name = v
		case []byte:
			name = string(v)
		default:
			return fmt.Errorf("%w column name %v", ErrInvalid, a)
		}
		c.name = name
		return nil
	}
}

// WithColumnDataType returns a column option to set the data type.
func WithColumnDataType(v any) ColumnOption {
	return func(c *column) error {
		dt, err := query.NewDataTypeFrom(v)
		if err != nil {
			return err
		}
		c.dataType = dt
		return nil
	}
}

// WithColumnConstraint returns a column option to set the column constraint.
func WithColumnConstraint(v any) ColumnOption {
	return func(c *column) error {
		// c.constraint = constraint
		return nil
	}
}

// NewColumn returns a new column.
func NewColumn(opts ...ColumnOption) (Column, error) {
	c := &column{
		name:       "",
		dataType:   query.UnknownData,
		constraint: query.ConstraintNone,
	}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

// Name returns the column name.
func (c *column) Name() string {
	return c.name
}

// DataType returns the data type.
func (c *column) DataType() DataType {
	return c.dataType
}

// Constraint returns the column constraint.
func (c *column) Constraint() Constraint {
	return c.constraint
}
