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

import "github.com/cybergarage/go-sqlparser/sql/util/strings"

// AlterTableOption represents an alter table option function.
type AlterTableOption = func(*alterTableStmt)

// AlterTable represents a "ALTER TABLE" statement interface.
type AlterTable interface {
	Statement
	TableName() string
	AddColumn() (*Column, bool)
	AddIndex() (*Index, bool)
	DropColumn() (*Column, bool)
	RenameColumns() (*Column, *Column, bool)
	RenameTo() (*Table, bool)
}

// alterTableStmt is a "ALTER TABLE" statement.
type alterTableStmt struct {
	*Schema
	*Table
	renameTableTo    *Table
	renameColumnFrom *Column
	renameColumnTo   *Column
	addColumn        *Column
	addIndex         *Index
	dropColumn       *Column
	dropIndex        *Index
}

// NewAlterTableWith returns a new alterTable statement instance with the specified options.
func NewAlterTableWith(tblName string, opts ...AlterTableOption) *alterTableStmt {
	stmt := &alterTableStmt{
		Schema:           nil,
		Table:            NewTableWith(tblName),
		renameTableTo:    nil,
		renameColumnFrom: nil,
		renameColumnTo:   nil,
		addColumn:        nil,
		addIndex:         nil,
		dropColumn:       nil,
		dropIndex:        nil,
	}
	for _, opt := range opts {
		opt(stmt)
	}
	return stmt
}

// WithAlterTableSchema sets a schema.
func WithAlterTableSchema(name string) func(*alterTableStmt) {
	return func(stmt *alterTableStmt) {
		stmt.Schema = NewSchemaWith(name)
	}
}

// WithAlterTableRenameTo sets a rename table.
func WithAlterTableRenameTo(tbl *Table) func(*alterTableStmt) {
	return func(stmt *alterTableStmt) {
		stmt.renameTableTo = tbl
	}
}

// WithAlterTableRenameColumn sets a rename column.
func WithAlterTableRenameColumn(from *Column, to *Column) func(*alterTableStmt) {
	return func(stmt *alterTableStmt) {
		stmt.renameColumnFrom = from
		stmt.renameColumnTo = to
	}
}

// WithAlterTableAddColumn sets an add column.
func WithAlterTableAddColumn(column *Column) func(*alterTableStmt) {
	return func(stmt *alterTableStmt) {
		stmt.addColumn = column
	}
}

// WithAlterTableAddIndex sets an add index.
func WithAlterTableAddIndex(index *Index) func(*alterTableStmt) {
	return func(stmt *alterTableStmt) {
		stmt.addIndex = index
	}
}

// WithAlterTableDropColumn sets a drop column.
func WithAlterTableDropColumn(column *Column) func(*alterTableStmt) {
	return func(stmt *alterTableStmt) {
		stmt.dropColumn = column
	}
}

// WithAlterTableDropColumn sets a drop index.
func WithAlterTableDropIndex(index *Index) func(*alterTableStmt) {
	return func(stmt *alterTableStmt) {
		stmt.dropIndex = index
	}
}

// StatementType returns the statement type.
func (stmt *alterTableStmt) StatementType() StatementType {
	return AlterTableStatement
}

// RenameTo returns the rename table.
func (stmt *alterTableStmt) RenameTo() (*Table, bool) {
	if stmt.renameTableTo == nil {
		return nil, false
	}
	return stmt.renameTableTo, true
}

// RenameColumns returns the rename columns.
func (stmt *alterTableStmt) RenameColumns() (*Column, *Column, bool) {
	if stmt.renameColumnFrom == nil || stmt.renameColumnTo == nil {
		return nil, nil, false
	}
	return stmt.renameColumnFrom, stmt.renameColumnTo, true
}

// AddColumn returns the add column.
func (stmt *alterTableStmt) AddColumn() (*Column, bool) {
	if stmt.addColumn == nil {
		return nil, false
	}
	return stmt.addColumn, true
}

// AddIndex returns the add index.
func (stmt *alterTableStmt) AddIndex() (*Index, bool) {
	if stmt.addIndex == nil {
		return nil, false
	}
	return stmt.addIndex, true
}

// DropColumn returns the drop column.
func (stmt *alterTableStmt) DropColumn() (*Column, bool) {
	if stmt.dropColumn == nil {
		return nil, false
	}
	return stmt.dropColumn, true
}

// DropIndex returns the drop index.
func (stmt *alterTableStmt) DropIndex() (*Index, bool) {
	if stmt.dropIndex == nil {
		return nil, false
	}
	return stmt.dropIndex, true
}

// String returns the statement string representation.
func (stmt *alterTableStmt) String() string {
	elems := []string{
		"ALTER",
		"TABLE",
		stmt.Table.String(),
	}
	if tbl, ok := stmt.RenameTo(); ok {
		elems = append(elems,
			[]string{
				"RENAME",
				"TO",
				tbl.String(),
			}...)
	}
	if f, t, ok := stmt.RenameColumns(); ok {
		elems = append(elems,
			[]string{
				"RENAME",
				"COLUMN",
				f.Name(),
				"TO",
				t.Name(),
			}...)
	}
	if c, ok := stmt.AddColumn(); ok {
		elems = append(elems, "ADD")
		if c.Constrains() != ConstraintNone {
			elems = append(elems, c.Constrains().String())
		}
		elems = append(elems, c.DefinitionString())
	}
	if i, ok := stmt.AddIndex(); ok {
		elems = append(elems, "ADD")
		elems = append(elems, i.DefinitionString())
	}
	if c, ok := stmt.DropColumn(); ok {
		elems = append(elems,
			[]string{
				"DROP",
				c.Name(),
			}...)
	}
	return strings.JoinWithSpace(elems)
}
