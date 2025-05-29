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
	"github.com/cybergarage/go-sqlparser/sql/errors"
)

type resultset struct {
	schema       Schema
	rows         []Row
	rowsAffected uint
	rowCursor    uint
	offset       uint
	limit        uint
}

// ResultSet represents a response resultset interface.
type ResultSetOption func(*resultset) error

// WithResultSetRowsAffected returns a resultset option to set the rows affected.
func WithResultSetRowsAffected(rowsAffected uint) ResultSetOption {
	return func(r *resultset) error {
		r.rowsAffected = rowsAffected
		return nil
	}
}

// WithResultSetSchema returns a resultset option to set the schema.
func WithResultSetSchema(schema Schema) ResultSetOption {
	return func(r *resultset) error {
		r.schema = schema
		return nil
	}
}

// WithResultSetRows returns a resultset option to set the rows.
func WithResultSetRows(rows []Row) ResultSetOption {
	return func(r *resultset) error {
		r.rows = rows
		return nil
	}
}

// WithResultSetOffset returns a resultset option to set the offset.
func WithResultSetOffset(offset uint) ResultSetOption {
	return func(r *resultset) error {
		r.offset = offset
		return nil
	}
}

// WithResultSetLimit returns a resultset option to set the limit.
func WithResultSetLimit(limit uint) ResultSetOption {
	return func(r *resultset) error {
		r.limit = limit
		return nil
	}
}

// NewResultSet returns a new ResultSet.
func NewResultSet(opts ...ResultSetOption) ResultSet {
	rs := newResultSet()
	for _, opt := range opts {
		opt(rs)
	}
	return rs
}

func newResultSet() *resultset {
	return &resultset{
		schema:       nil,
		rows:         []Row{},
		offset:       0,
		limit:        0,
		rowsAffected: 0,
		rowCursor:    0,
	}
}

// NewResultSetFrom creates a new ResultSet from the given options.
func NewResultSetFrom(opts ...ResultSetOption) (ResultSet, error) {
	rs := newResultSet()
	for _, opt := range opts {
		if err := opt(rs); err != nil {
			return nil, err
		}
	}
	return rs, nil
}

// RowsAffected returns the number of rows affected.
func (r *resultset) RowsAffected() uint {
	return r.rowsAffected
}

// Next returns the next row.
func (r *resultset) Next() bool {
	if r.rowCursor < r.offset {
		r.rowCursor = r.offset
	}
	r.rowCursor++
	if 0 < r.limit {
		if (r.offset + r.limit) < r.rowCursor {
			return false
		}
	}
	return r.rowCursor <= uint(len(r.rows))
}

// Row returns the current row.
func (r *resultset) Row() (Row, error) {
	if (r.rowCursor == 0) || (uint(len(r.rows)) < r.rowCursor) {
		return nil, errors.ErrNoRows
	}
	return r.rows[r.rowCursor-1], nil
}

// Schema returns the schema.
func (r *resultset) Schema() Schema {
	return r.schema
}

// Close closes the resultset.
func (r *resultset) Close() error {
	return nil
}
