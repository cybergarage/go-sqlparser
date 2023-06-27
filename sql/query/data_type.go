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
	"fmt"
	"strings"
)

// DataType represents a data type.
type DataType struct {
	Type   int
	Length int
}

const (
	UnknownData = iota
	BigIntData  = iota
	BinaryData
	BitData
	BlobData
	BooleanData
	CharData
	CharacterData
	ClobData
	DateData
	DecimalData
	DoubleData
	FloatData
	IntData
	IntegerData
	LongBlobData
	LongTextData
	MediumBlobData
	MediumIntData
	MediumTextData
	NumericData
	RealData
	SetData
	SmallIntData
	TextData
	TimeData
	TimeStampData
	TinyBlobData
	TinyIntData
	TinyTextData
	VarBinaryData
	VarCharData
	VarCharacterData
	YearData
)

var dataTypeStrings = map[int]string{
	BigIntData:       "BIGINT",
	BinaryData:       "BINARY",
	BitData:          "BIT",
	BlobData:         "BLOB",
	BooleanData:      "BOOLEAN",
	CharData:         "CHAR",
	CharacterData:    "CHARACTER",
	ClobData:         "CLOB",
	DateData:         "DATE",
	DecimalData:      "DECIMAL",
	DoubleData:       "DOUBLE",
	FloatData:        "FLOAT",
	IntData:          "INT",
	IntegerData:      "INTEGER",
	LongBlobData:     "LONGBLOB",
	LongTextData:     "LONGTEXT",
	MediumBlobData:   "MEDIUMBLOB",
	MediumIntData:    "MEDIUMINT",
	MediumTextData:   "MEDIUMTEXT",
	NumericData:      "NUMERIC",
	RealData:         "REAL",
	SetData:          "SET",
	SmallIntData:     "SMALLINT",
	TextData:         "TEXT",
	TimeData:         "TIME",
	TimeStampData:    "TIMESTAMP",
	TinyBlobData:     "TINYBLOB",
	TinyIntData:      "TINYINT",
	TinyTextData:     "TINYTEXT",
	VarBinaryData:    "VARBINARY",
	VarCharData:      "VARCHAR",
	VarCharacterData: "VARCHARACTER",
	YearData:         "YEAR",
}

// NewDataTypeFrom returns the data type of the specified string.
func NewDataTypeFrom(s string, l int) (*DataType, error) {
	for dataType, dataTypeString := range dataTypeStrings {
		if dataTypeString == strings.ToUpper(s) {
			return &DataType{
				Type:   dataType,
				Length: l,
			}, nil
		}
	}
	return nil, fmt.Errorf("%w: %s", ErrInvalidDataType, s)
}

// DataType returns the column data type.
func (da *DataType) DataType() int {
	return da.Type
}

// DataLength returns the column data length.
func (da *DataType) DataLength() int {
	return da.Length
}

// String returns the string representation.
func (da *DataType) String() string {
	s, ok := dataTypeStrings[da.Type]
	if ok {
		return s
	}
	return ""
}
