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

// Copy is a "COPY" statement.
type Copy struct {
	*Table
	source string
}

// NewCopyWith returns a new Copy statement instance with the specified parameters.
func NewCopyWith(tblName string, src string) *Copy {
	return &Copy{
		Table:  NewTableWith(tblName),
		source: src,
	}
}

// StatementType returns the statement type.
func (stmt *Copy) StatementType() StatementType {
	return CopyStatement
}

// Source returns the source resource.
func (stmt *Copy) Source() string {
	return stmt.source
}

// String returns the statement string representation.
func (stmt *Copy) String() string {
	strs := []string{
		"COPY",
		stmt.TableName(),
		"FROM",
		stmt.Source(),
	}
	return strings.JoinWithSpace(strs)
}
