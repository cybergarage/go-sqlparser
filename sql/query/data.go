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
type DataType uint8

// DataDef represents a data definition.
type DataDef struct {
	Type   DataType
	Length int
}

const (
	UnknownData = iota
	BigIntData
	BinaryData
	BitData
	BlobData
	BooleanData
	CharData
	CharacterData
	ClobData
	DateData
	DateTimeData
	DecimalData
	DoubleData
	DoublePrecision
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

var dataTypeStrings = map[DataType]string{
	BigIntData:       "BIGINT",
	BinaryData:       "BINARY",
	BitData:          "BIT",
	BlobData:         "BLOB",
	BooleanData:      "BOOLEAN",
	CharData:         "CHAR",
	CharacterData:    "CHARACTER",
	ClobData:         "CLOB",
	DateData:         "DATE",
	DateTimeData:     "DATETIME",
	DecimalData:      "DECIMAL",
	DoubleData:       "DOUBLE",
	DoublePrecision:  "DOUBLE PRECISION",
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

// String returns the string representation.
func (t DataType) String() string {
	s, ok := dataTypeStrings[t]
	if !ok {
		return ""
	}
	return s
}

// NewDataWith returns a new DataType instance with the specified type and length.
func NewDataWith(t DataType, l int) *DataDef {
	return &DataDef{
		Type:   t,
		Length: l,
	}
}

// NewDataFrom returns the data type of the specified string.
func NewDataFrom(s string, l int) (*DataDef, error) {
	us := strings.ToUpper(s)
	for dataType, dataTypeString := range dataTypeStrings {
		if dataTypeString == us {
			return NewDataWith(dataType, l), nil
		}
	}
	for dataType, dataTypeString := range dataTypeStrings {
		dataTypeString = strings.ReplaceAll(dataTypeString, " ", "")
		if dataTypeString == us {
			return NewDataWith(dataType, l), nil
		}
	}
	return nil, fmt.Errorf("%w data type : %s", ErrInvalid, s)
}

// DataType returns the column data type.
func (da *DataDef) DataType() DataType {
	return da.Type
}

// DataTypeLen returns the column data type length.
func (da *DataDef) DataTypeLen() int {
	return da.Length
}

// String returns the string representation.
func (da *DataDef) String() string {
	s, ok := dataTypeStrings[da.Type]
	if !ok {
		return ""
	}
	if 0 < da.Length {
		s += fmt.Sprintf("(%d)", da.Length)
	}
	return s
}
