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

// Query represents a common query interface.
type Query interface {
	StatementType() query.StatementType
	String() string
}

// CreateDatabase is a "CREATE DATABASE" statement.
type CreateDatabase interface {
	Query
	DatabaseName() string
	IfNotExists() bool
}

// CreateTable is a "CREATE TABLE" statement.
type CreateTable interface {
	Query
	TableName() string
	Schema() *query.Schema
	IfNotExists() bool
}

// AlterDatabase is a "ALTER DATABASE" statement.
type AlterDatabase interface {
	DatabaseName() string
	RenameTo() *query.Database
}

// AlterTable is a "ALTER TABLE" statement.
type AlterTable interface {
	TableName() string
	AddColumn() (*query.Column, bool)
	AddIndex() (*query.Index, bool)
	DropColumn() (*query.Column, bool)
	RenameColumns() (*query.Column, *query.Column, bool)
	RenameTo() (*query.Table, bool)
}

// DropTable is a "DROP TABLE" statement.
type DropDatabase interface {
	Query
	DatabaseName() string
	IfExists() bool
}

// DropTable is a "DROP TABLE" statement.
type DropTable interface {
	Tables() query.TableList
	IfExists() bool
}
