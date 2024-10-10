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

// Response represents a response message interface.
type Response interface {
	// Bytes returns the message bytes.
	Bytes() ([]byte, error)
}

// Responses represents a list of response.
type Responses []Response

// DDOExecutor defines a executor interface for DDO (Data Definition Operations).
type DDOExecutor interface {
	// CreateDatabase handles a CREATE DATABASE query.
	CreateDatabase(Conn, *query.CreateDatabase) (Responses, error)
	// CreateTable handles a CREATE TABLE query.
	CreateTable(Conn, *query.CreateTable) (Responses, error)
	// AlterDatabase handles a ALTER DATABASE query.
	AlterDatabase(Conn, *query.AlterDatabase) (Responses, error)
	// AlterTable handles a ALTER TABLE query.
	AlterTable(Conn, *query.AlterTable) (Responses, error)
	// DropDatabase handles a DROP DATABASE query.
	DropDatabase(Conn, *query.DropDatabase) (Responses, error)
	// DropIndex handles a DROP INDEX query.
	DropTable(Conn, *query.DropTable) (Responses, error)
}

// DMOExecutor defines a executor interface for DMO (Data Manipulation Operations).
type DMOExecutor interface {
	// Insert handles a INSERT query.
	Insert(Conn, *query.Insert) (Responses, error)
	// Select handles a SELECT query.
	Select(Conn, *query.Select) (Responses, error)
	// Update handles a UPDATE query.
	Update(Conn, *query.Update) (Responses, error)
	// Delete handles a DELETE query.
	Delete(Conn, *query.Delete) (Responses, error)
}

// TCLExecutor defines a executor interface for TCL (Transaction Control Language).
type TCLExecutor interface {
	// Begin handles a BEGIN query.
	Begin(*Conn, *query.Begin) (Responses, error)
	// Commit handles a COMMIT query.
	Commit(*Conn, *query.Commit) (Responses, error)
	// Rollback handles a ROLLBACK query.
	Rollback(*Conn, *query.Rollback) (Responses, error)
}

// QueryExecutor represents a user query message executor.
type QueryExecutor interface {
	DDOExecutor
	DMOExecutor
}

// SystemQueryExecutor represents a system query message executor.
type SystemQueryExecutor interface {
	// SystemSelect handles a SELECT query for system tables.
	SystemSelect(Conn, *query.Select) (Responses, error)
}

// ErrorHandler represents a user error handler.
type ErrorHandler interface {
	ParserError(Conn, string, error) (Responses, error)
}

// Executor represents a frontend message executor.
type Executor interface { // nolint: interfacebloat
	TCLExecutor
	QueryExecutor
	SystemQueryExecutor
	ErrorHandler
}
