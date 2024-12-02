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

import (
	"github.com/cybergarage/go-sqlparser/sql/errors"
)

// NewAlterTableFrom convert the specified statement to an AlterTable.
func NewAlterTableFrom(stmt Statement) (AlterTable, error) {
	switch stmt := stmt.(type) {
	case CreateIndex:
		alterTbl := NewAlterTableWith(
			stmt.TableName(),
			WithAlterTableAddIndex(stmt.Index()),
		)
		return alterTbl, nil
	case DropIndex:
		alterTbl := NewAlterTableWith(
			stmt.TableName(),
			WithAlterTableDropIndex(stmt.Index()),
		)
		return alterTbl, nil
	default:
		return nil, errors.NewErrNotSupported(stmt.StatementType().String())
	}
}
