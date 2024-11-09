// Copyright (C) 2019 The go-sqlparser Authors. All rights reserved.
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

package query

import (
	"github.com/cybergarage/go-sqlparser/sql/util/strings"
)

// CreateIndex is a "CREATE INDEX" statement interface.
type CreateIndex interface {
	Statement
	// TableName returns the table name.
	TableName() string
	// Index returns the index.
	Index() Index
	// IfNotExists returns true if the "IF NOT EXISTS" option is set.
	IfNotExists() bool
}

// createIndex is a "CREATE INDEX" statement.
type createIndex struct {
	schema Schema
	*IfNotExistsOpt
}

// NewCreateIndexWith returns a new createIndex statement instance with the specified parameters.
func NewCreateIndexWith(schema Schema, ifne *IfNotExistsOpt) CreateIndex {
	return &createIndex{
		schema:         schema,
		IfNotExistsOpt: ifne,
	}
}

// StatementType returns the statement type.
func (stmt *createIndex) StatementType() StatementType {
	return CreateIndexStatement
}

// TableName returns the table name.
func (stmt *createIndex) TableName() string {
	return stmt.schema.FullTableName()
}

// Index returns the index.
func (stmt *createIndex) Index() Index {
	idxes := stmt.schema.Indexes()
	if len(idxes) == 0 {
		return nil
	}
	return idxes[0]
}

// String returns the statement string representation.
func (stmt *createIndex) String() string {
	strs := []string{
		"CREATE",
		"INDEX",
	}
	if stmt.IfNotExists() {
		strs = append(strs, stmt.IfNotExistsOpt.String())
	}
	idx := stmt.Index()
	if idx != nil {
		strs = append(strs, idx.Name())
		strs = append(strs, "ON")
		strs = append(strs, stmt.TableName()+"("+strings.JoinWithComma(idx.Columns().ColumnNames())+")")
	}
	return strings.JoinWithSpace(strs)
}
