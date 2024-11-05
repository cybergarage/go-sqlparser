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

// StatementType is a statement type.
type StatementType uint8

const (
	CreateDatabaseStatement = iota
	CreateTableStatement
	CreateIndexStatement
	InsertStatement
	SelectStatement
	UpdateStatement
	DeleteStatement
	DropDatabaseStatement
	DropTableStatement
	DropIndexStatement
	AlterDatabaseStatement
	AlterTableStatement
	AlterIndexStatement
	CopyStatement
	CommitStatement
	VacuumStatement
	TruncateStatement
	BeginStatement
	RollbackStatement
	UseStatement
)

// StatementType returns the statement type.
func (stmtType StatementType) String() string {
	switch stmtType {
	case CreateDatabaseStatement:
		return "CREATE DATABASE"
	case CreateTableStatement:
		return "CREATE TABLE"
	case CreateIndexStatement:
		return "CREATE INDEX"
	case InsertStatement:
		return "INSERT"
	case SelectStatement:
		return "SELECT"
	case UpdateStatement:
		return "UPDATE"
	case DeleteStatement:
		return "DELETE"
	case DropDatabaseStatement:
		return "DROP DATABASE"
	case DropIndexStatement:
		return "DROP INDEX"
	case DropTableStatement:
		return "DROP TABLE"
	case AlterDatabaseStatement:
		return "ALTER DATABASE"
	case AlterTableStatement:
		return "ALTER TABLE"
	case AlterIndexStatement:
		return "ALTER INDEX"
	case CopyStatement:
		return "COPY"
	case CommitStatement:
		return "COMMIT"
	case VacuumStatement:
		return "VACUUM"
	case TruncateStatement:
		return "TRUNCATE"
	case BeginStatement:
		return "BEGIN"
	case RollbackStatement:
		return "ROLLBACK"
	case UseStatement:
		return "USE"
	default:
		return ""
	}
}
