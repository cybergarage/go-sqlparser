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
	strconv "strings"

	"github.com/cybergarage/go-sqlparser/sql/util/strings"
)

// CopyFormat represents a copy format.
type CopyFormat uint8

const (
	// CopyFormatText represents a text format.
	CopyFormatText CopyFormat = iota
	// CopyFormatCSV represents a CSV format.
	CopyFormatCSV
	// CopyFormatBinary represents a binary format.
	CopyFormatBinary
)

// String returns the string representation of the format.
func (fmt CopyFormat) String() string {
	switch fmt {
	case CopyFormatText:
		return "TEXT"
	case CopyFormatCSV:
		return "CSV"
	case CopyFormatBinary:
		return "BINARY"
	}
	return ""
}

// CopyOption represents a copy option.
type CopyOption = func(*Copy)

// WithCopyColumns returns a copy option to set the columns.
func WithCopyColumns(columns ...*Column) func(*Copy) {
	return func(stmt *Copy) {
		stmt.ColumnList = NewColumnsWith(columns...)
	}
}

// WithCopyFormat returns a copy option to set the format.
func WithCopyFormat(fmt string) func(*Copy) {
	return func(stmt *Copy) {
		switch strconv.ToUpper(fmt) {
		case "TEXT":
			stmt.format = CopyFormatText
		case "CSV":
			stmt.format = CopyFormatCSV
		case "BINARY":
			stmt.format = CopyFormatBinary
		}
	}
}

// Copy is a "COPY" statement.
type Copy struct {
	*Table
	ColumnList
	source string
	format CopyFormat
}

// NewCopyWith returns a new Copy statement instance with the specified parameters.
func NewCopyWith(tblName string, src string, opts ...CopyOption) *Copy {
	copy := &Copy{
		Table:      NewTableWith(tblName),
		ColumnList: NewColumns(),
		source:     src,
		format:     CopyFormatText,
	}
	for _, opt := range opts {
		opt(copy)
	}
	return copy
}

// StatementType returns the statement type.
func (stmt *Copy) StatementType() StatementType {
	return CopyStatement
}

// Source returns the source resource.
func (stmt *Copy) Source() string {
	return stmt.source
}

func (stmt *Copy) Format() CopyFormat {
	return stmt.format
}

// String returns the statement string representation.
func (stmt *Copy) String() string {
	strs := []string{
		"COPY",
		stmt.TableName(),
	}
	if 0 < len(stmt.ColumnList) {
		strs = append(strs, "("+strings.JoinWithSpace(stmt.ColumnList.Names())+")")
	}
	strs = append(strs,
		"FROM",
		stmt.Source(),
		stmt.format.String(),
	)
	return strings.JoinWithSpace(strs)
}
