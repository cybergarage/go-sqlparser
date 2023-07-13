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

// Database represents a database.
type Database struct {
	name string
}

// NewDatabaseWith returns a new Database instance with the specified name.
func NewDatabaseWith(name string) *Database {
	return &Database{
		name: name,
	}
}

// Database returns the database.
func (db *Database) Database() *Database {
	return db
}

// DatabaseName returns the database name.
func (db *Database) DatabaseName() string {
	return db.name
}

// String returns the string representation.
func (db *Database) String() string {
	return db.name
}
