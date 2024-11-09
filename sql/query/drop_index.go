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

// DropIndex is a "DROP INDEX" statement interface.
type DropIndex interface {
	Statement
	// TableName returns the table name.
	TableName() string
	// IndexName returns the index name.
	IndexName() string
	// IfExists returns true if the "IF EXISTS" option is set.
	IfExists() bool
}

// dropIndex is a "DROP INDEX" statement.
type dropIndex struct {
	Index
	Schema
	*IfExistsOpt
}

// NewDropIndexWith returns a new dropIndex statement instance with the specified parameters.
func NewDropIndexWith(schemaName string, idxName string, ife *IfExistsOpt) DropIndex {
	return &dropIndex{
		Schema:      NewSchemaWith(schemaName),
		Index:       NewIndexWith(idxName, UnknownIndex, NewColumns()),
		IfExistsOpt: ife,
	}
}

// TableName returns the table name.
func (stmt *dropIndex) TableName() string {
	return stmt.Schema.FullTableName()
}

// IndexName returns the index name.
func (stmt *dropIndex) IndexName() string {
	return stmt.Index.Name()
}

// StatementType returns the statement type.
func (stmt *dropIndex) StatementType() StatementType {
	return DropIndexStatement
}

// String returns the statement string representation.
func (stmt *dropIndex) String() string {
	strs := []string{
		"DROP",
		"INDEX",
	}
	if stmt.IfExists() {
		strs = append(strs, stmt.IfExistsOpt.String())
	}
	strs = append(strs, stmt.Index.Name())
	tableName := stmt.Schema.FullTableName()
	if 0 < len(tableName) {
		strs = append(strs, "ON", tableName)
	}
	return strings.JoinWithSpace(strs)
}
