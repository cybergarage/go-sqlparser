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
	"github.com/cybergarage/go-sqlparser/sql/fn"
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

// DataType represents a data type.
type DataType = query.DataType // DataType represents a data type.

// Constraint represents a constraint.
type Constraint = query.Constraint

// Function represents a function.
type Function = query.Function

// FunctionExecutor represents a function executor.
type FunctionExecutor = query.FunctionExecutor

// AggregateFunction represents an aggregate function.
type AggregateFunction = fn.AggregateFunction

// AggregateResultSet represents an aggregate result set.
type AggregateResultSet = fn.AggregateResultSet

// Schema represents a schema.
type Schema = query.Schema

// Statement represents a common query interface.
type Statement = query.Statement

// CreateDatabase represents a "CREATE DATABASE" statement interface.
type CreateDatabase = query.CreateDatabase

// CreateTable represents a "CREATE TABLE" statement interface.
type CreateTable = query.CreateTable

// AlterDatabase represents a "ALTER DATABASE" statement interface.
type AlterDatabase = query.AlterDatabase

// AlterTable represents a "ALTER TABLE" statement interface.
type AlterTable = query.AlterTable

// DropDatabase represents a "DROP DATABASE" statement interface.
type DropDatabase = query.DropDatabase

// DropTable represents a "DROP TABLE" statement interface.
type DropTable = query.DropTable

// Insert represents a "INSERT" statement interface.
type Insert = query.Insert

// Select represents a "SELECT" statement interface.
type Select = query.Select

// Update represents a "UPDATE" statement interface.
type Update = query.Update

// Delete represents a "DELETE" statement interface.
type Delete = query.Delete

// Begin represents a "BEGIN" statement interface.
type Begin = query.Begin

// Commit represents a "COMMIT" statement interface.
type Commit = query.Commit

// Rollback represents a "ROLLBACK" statement interface.
type Rollback = query.Rollback

// Copy represents a "COPY" statement interface.
type Copy = query.Copy

// Vacuum represents a "VACUUM" statement interface.
type Vacuum = query.Vacuum

// Truncate represents a "TRUNCATE" statement interface.
type Truncate = query.Truncate

// Use represents a "USE" statement interface.
type Use = query.Use

// CreateIndex represents a "CREATE INDEX" statement interface.
type CreateIndex = query.CreateIndex

// DropIndex represents a "DROP INDEX" statement interface.
type DropIndex = query.DropIndex
