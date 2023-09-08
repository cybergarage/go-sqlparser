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
type AlterTableOption = func(*AlterTable)

// AlterTable is a "ALTER TABLE" statement.
type AlterTable struct {
	*Schema
	*Table
	renameTableTo    *Table
	renameColumnFrom *Column
	renameColumnTo   *Column
	addColumn        *Column
	dropColumn       *Column
}

// NewAlterTableWith returns a new AlterTable statement instance with the specified options.
func NewAlterTableWith(tblName string, opts ...AlterTableOption) *AlterTable {
	stmt := &AlterTable{
		Schema:           nil,
		Table:            NewTableWith(tblName),
		renameTableTo:    nil,
		renameColumnFrom: nil,
		renameColumnTo:   nil,
		addColumn:        nil,
		dropColumn:       nil,
	}
	for _, opt := range opts {
		opt(stmt)
	}
	return stmt
}

// WithAlterTableSchema sets a schema.
func WithAlterTableSchema(name string) func(*AlterTable) {
	return func(stmt *AlterTable) {
		stmt.Schema = NewSchemaWith(name)
	}
}

// WithAlterTableRenameTo sets a rename table.
func WithAlterTableRenameTo(tbl *Table) func(*AlterTable) {
	return func(stmt *AlterTable) {
		stmt.renameTableTo = tbl
	}
}

// WithAlterTableRenameColumn sets a rename column.
func WithAlterTableRenameColumn(from *Column, to *Column) func(*AlterTable) {
	return func(stmt *AlterTable) {
		stmt.renameColumnFrom = from
		stmt.renameColumnTo = to
	}
}

// WithAlterTableAddColumn sets an add column.
func WithAlterTableAddColumn(column *Column) func(*AlterTable) {
	return func(stmt *AlterTable) {
		stmt.addColumn = column
	}
}

// WithAlterTableDropColumn sets a drop column.
func WithAlterTableDropColumn(column *Column) func(*AlterTable) {
	return func(stmt *AlterTable) {
		stmt.dropColumn = column
	}
}

// StatementType returns the statement type.
func (stmt *AlterTable) StatementType() StatementType {
	return AlterTableStatement
}

// RenameTable returns the rename table.
func (stmt *AlterTable) RenameTable() (*Table, bool) {
	if stmt.renameTableTo == nil {
		return nil, false
	}
	return stmt.renameTableTo, true
}

// RenameColums returns the rename columns.
func (stmt *AlterTable) RenameColums() (*Column, *Column, bool) {
	if stmt.renameColumnFrom == nil || stmt.renameColumnTo == nil {
		return nil, nil, false
	}
	return stmt.renameColumnFrom, stmt.renameColumnTo, true
}

// AddColum returns the add column.
func (stmt *AlterTable) AddColum() (*Column, bool) {
	if stmt.addColumn == nil {
		return nil, false
	}
	return stmt.addColumn, true
}

// DropColum returns the drop column.
func (stmt *AlterTable) DropColum() (*Column, bool) {
	if stmt.dropColumn == nil {
		return nil, false
	}
	return stmt.dropColumn, true
}

// String returns the statement string representation.
func (stmt *AlterTable) String() string {
	elems := []string{
		"ALTER",
		"TABLE",
		stmt.Table.String(),
	}
	if tbl, ok := stmt.RenameTable(); ok {
		elems = append(elems,
			[]string{
				"RENAME",
				"TO",
				tbl.String(),
			}...)
	}
	if f, t, ok := stmt.RenameColums(); ok {
		elems = append(elems,
			[]string{
				"RENAME",
				"COLUMN",
				f.Name(),
				"TO",
				t.Name(),
			}...)
	}
	if c, ok := stmt.AddColum(); ok {
		elems = append(elems, "ADD")
		if c.Constrains() != ColumnConstraintNone {
			elems = append(elems, c.Constrains().String())
		}
		elems = append(elems, c.DefinitionString())

	}
	if c, ok := stmt.DropColum(); ok {
		elems = append(elems,
			[]string{
				"DROP",
				c.Name(),
			}...)
	}
	return strings.JoinWithSpace(elems)
}
