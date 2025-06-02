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
	groupBy      string
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

// WithResultSetGroupBy returns a resultset option to set the group by clause.
func WithResultSetGroupBy(groupBy string) ResultSetOption {
	return func(r *resultset) error {
		r.groupBy = groupBy
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
		groupBy:      "",
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
	if !rs.schema.Selectors().HasAggregator() {
		return rs, nil
	}
	return rs.Aggregate()
}

// RowsAffected returns the number of rows affected.
func (rs *resultset) RowsAffected() uint {
	return rs.rowsAffected
}

// Next returns the next row.
func (rs *resultset) Next() bool {
	if rs.rowCursor < rs.offset {
		rs.rowCursor = rs.offset
	}
	rs.rowCursor++
	if 0 < rs.limit {
		if (rs.offset + rs.limit) < rs.rowCursor {
			return false
		}
	}
	return rs.rowCursor <= uint(len(rs.rows))
}

// Row returns the current row.
func (rs *resultset) Row() (Row, error) {
	if (rs.rowCursor == 0) || (uint(len(rs.rows)) < rs.rowCursor) {
		return nil, errors.ErrNoRows
	}
	return rs.rows[rs.rowCursor-1], nil
}

// Schema returns the schema.
func (rs *resultset) Schema() Schema {
	return rs.schema
}

// Offset returns the offset.
func (rs *resultset) Offset() uint {
	return rs.offset
}

// Limit returns the limit.
func (rs *resultset) Limit() uint {
	return rs.limit
}

// GroupBy returns the group by clause.
func (rs *resultset) GroupBy() string {
	return rs.groupBy
}

// Close closes the resultset.
func (rs *resultset) Close() error {
	return nil
}
