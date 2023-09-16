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
	"github.com/cybergarage/go-sqlparser/sql/util/strings"
)

// Begin represents a begin statement.
type Begin struct {
}

// NewBegin returns a new begin statement.
func NewBegin() *Begin {
	return &Begin{}
}

// StatementType returns the statement type.
func (stmt *Begin) StatementType() StatementType {
	return BeginStatement
}

// String returns the statement string representation.
func (stmt *Begin) String() string {
	strs := []string{
		"BEGIN",
	}
	return strings.JoinWithSpace(strs)
}
