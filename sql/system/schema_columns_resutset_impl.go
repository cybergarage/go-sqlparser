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
	"strings"

	"github.com/cybergarage/go-sqlparser/sql/query/response/resultset"
)

// ResultSet represents a response resultset interface.
type ResultSet = resultset.ResultSet

type rs struct {
	columns []SchemaColumn
}

// NewSchemaColumns returns a new SchemaColumns instance.
type SchemaColumnOption func(*rs)

// WithColumns returns a SchemaColumnOption that sets the columns.
func WithColumns(columns []SchemaColumn) SchemaColumnOption {
	return func(s *rs) {
		s.columns = columns
	}
}

// NewSchemaColumns returns a new SchemaColumns instance.
func NewSchemaColumns(opts ...SchemaColumnOption) SchemaColumnsResultSet {
	s := &rs{
		columns: []SchemaColumn{},
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Columns returns the columns.
func (s *rs) Columns() []SchemaColumn {
	return s.columns
}

// LookupColumn returns the column by name.
func (s *rs) LookupColumn(name string) (SchemaColumn, error) {
	for _, column := range s.columns {
		if strings.EqualFold(column.Name(), name) {
			return column, nil
		}
	}
	return nil, newErrColumnNotFound(name)
}
