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

package sql

import (
	"github.com/cybergarage/go-sqlparser/sql/query"
)

const (
	EQ  = query.EQ
	NEQ = query.NEQ
	LT  = query.LT
	LE  = query.LE
	GT  = query.GT
	GE  = query.GE
	IN  = query.IN
	NIN = query.NIN
)

// Function represents a function.
type Function = query.Function

// FunctionExecutor represents a function executor.
type FunctionExecutor = query.FunctionExecutor

// AggregateFunction represents an aggregate function.
type AggregateFunction = query.AggregateFunction

// AggregateResultSet represents an aggregate result set.
type AggregateResultSet = query.AggregateResultSet

// Schema represents a schema.
type Schema = query.Schema

// Statement represents a common query interface.
type Statement = query.Statement

// CreateDatabase represents a "CREATE DATABASE" statement interface.
type CreateDatabase interface {
	Statement
	DatabaseName() string
	IfNotExists() bool
}

// CreateTable represents a "CREATE TABLE" statement interface.
type CreateTable interface {
	Statement
	TableName() string
	Schema() *Schema
	IfNotExists() bool
}

// AlterDatabase represents a "ALTER DATABASE" statement interface.
type AlterDatabase interface {
	Statement
	DatabaseName() string
	RenameTo() *query.Database
}

// AlterTable represents a "ALTER TABLE" statement interface.
type AlterTable interface {
	Statement
	TableName() string
	AddColumn() (*query.Column, bool)
	AddIndex() (*query.Index, bool)
	DropColumn() (*query.Column, bool)
	RenameColumns() (*query.Column, *query.Column, bool)
	RenameTo() (*query.Table, bool)
}

// DropTable represents a "DROP TABLE" statement interface.
type DropDatabase interface {
	Statement
	DatabaseName() string
	IfExists() bool
}

// DropTable represents a "DROP TABLE" statement interface.
type DropTable interface {
	Statement
	Tables() query.TableList
	IfExists() bool
}

// Insert represents a "INSERT" statement interface.
type Insert interface {
	Statement
	TableName() string
	Columns() query.ColumnList
	IsSelectAll() bool
}

// Select represents a "SELECT" statement interface.
type Select interface {
	Statement
	IsSelectAll() bool
	Selectors() query.SelectorList
	From() query.TableList
	Limit() *query.Limit
	GroupBy() *query.GroupBy
	OrderBy() *query.OrderBy
	Where() *query.Condition
}

// Update represents a "UPDATE" statement interface.
type Update interface {
	Statement
	TableName() string
	Columns() query.ColumnList
	Where() *query.Condition
}

// Delete represents a "DELETE" statement interface.
type Delete interface {
	Statement
	TableName() string
	Where() *query.Condition
}

// Begin represents a "BEGIN" statement interface.
type Begin interface {
	Statement
}

// Commit represents a "COMMIT" statement interface.
type Commit interface {
	Statement
}

// Rollback represents a "ROLLBACK" statement interface.
type Rollback interface {
	Statement
}

// Copy represents a "COPY" statement interface.
type Copy interface {
	Statement
	Columns() query.ColumnList
	TableName() string
}

// Vacuum represents a "VACUUM" statement interface.
type Vacuum interface {
	Statement
	TableName() string
}

// Truncate represents a "TRUNCATE" statement interface.
type Truncate interface {
	Statement
	Tables() query.TableList
}

// Use represents a "USE" statement interface.
type Use interface {
	Statement
	DatabaseName() string
}
