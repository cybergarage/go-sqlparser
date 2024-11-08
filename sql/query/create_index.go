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

// CreateIndex is a "CREATE INDEX" statement.
type CreateIndex struct {
	*Schema
	*IfNotExistsOpt
}

// NewCreateIndexWith returns a new CreateIndex statement instance with the specified parameters.
func NewCreateIndexWith(schema *Schema, ifne *IfNotExistsOpt) *CreateIndex {
	return &CreateIndex{
		Schema:         schema,
		IfNotExistsOpt: ifne,
	}
}

// StatementType returns the statement type.
func (stmt *CreateIndex) StatementType() StatementType {
	return CreateIndexStatement
}

// String returns the statement string representation.
func (stmt *CreateIndex) String() string {
	// TODO: implement me
	return ""
}
