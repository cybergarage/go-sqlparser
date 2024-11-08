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
	"github.com/cybergarage/go-sqlparser/sql/net"
)

// Conn represents a connection.
type Conn = net.Conn

// DDOExecutor defines a executor interface for DDO (Data Definition Operations).
type DDOExecutor interface {
	// CreateDatabase handles a CREATE DATABASE query.
	CreateDatabase(Conn, CreateDatabase) error
	// CreateTable handles a CREATE TABLE query.
	CreateTable(Conn, CreateTable) error
	// AlterDatabase handles a ALTER DATABASE query.
	AlterDatabase(Conn, AlterDatabase) error
	// AlterTable handles a ALTER TABLE query.
	AlterTable(Conn, AlterTable) error
	// DropDatabase handles a DROP DATABASE query.
	DropDatabase(Conn, DropDatabase) error
	// DropIndex handles a DROP INDEX query.
	DropTable(Conn, DropTable) error
}

// DMOExecutor defines a executor interface for DMO (Data Manipulation Operations).
type DMOExecutor interface {
	// Insert handles a INSERT query.
	Insert(Conn, Insert) error
	// Select handles a SELECT query.
	Select(Conn, Select) (ResultSet, error)
	// Update handles a UPDATE query.
	Update(Conn, Update) (ResultSet, error)
	// Delete handles a DELETE query.
	Delete(Conn, Delete) (ResultSet, error)
}

// DCOExecutor defines a executor interface for DCO (Data Control Operations).
type DCOExecutor interface {
	// Use handles a USE query.
	Use(Conn, Use) error
}

// TCOExecutor defines a executor interface for TCO (Transaction Control Operations).
type TCOExecutor interface {
	// Begin handles a BEGIN query.
	Begin(Conn, Begin) error
	// Commit handles a COMMIT query.
	Commit(Conn, Commit) error
	// Rollback handles a ROLLBACK query.
	Rollback(Conn, Rollback) error
}

// QueryExecutor represents a user query message executor.
type QueryExecutor interface {
	DDOExecutor
	DMOExecutor
	DCOExecutor
	TCOExecutor
}

// ErrorHandler represents a user error handler.
type ErrorHandler interface {
	ParserError(Conn, string, error) error
}

// Executor represents a frontend message executor.
type Executor interface {
	QueryExecutor
	ErrorHandler
}
