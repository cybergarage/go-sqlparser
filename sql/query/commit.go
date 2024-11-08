// Commitright (C) 2019 The go-sqlparser Authors. All rights reserved.
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

// Commit represents a "COMMIT" statement interface.
type Commit interface {
	Statement
}

// commitStmt is a "COMMIT" statement.
type commitStmt struct {
}

// NewCommitWith returns a new commitStmt statement instance with the specified parameters.
func NewCommit() *commitStmt {
	return &commitStmt{}
}

// StatementType returns the statement type.
func (stmt *commitStmt) StatementType() StatementType {
	return CommitStatement
}

// String returns the statement string representation.
func (stmt *commitStmt) String() string {
	strs := []string{
		"COMMIT",
	}
	return strings.JoinWithSpace(strs)
}
