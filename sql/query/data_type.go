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
type DataType int

// Data represents a data type.
type Data struct {
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

// String returns the string representation.
func (t DataType) String() string {
	s, ok := dataTypeStrings[t]
	if !ok {
		return ""
	}
	return s
}

// NewDataTypeWith returns a new DataType instance with the specified type and length.
func NewDataTypeWith(t DataType, l int) *Data {
	return &Data{
		Type:   t,
		Length: l,
	}
}

// NewDataFrom returns the data type of the specified string.
func NewDataFrom(s string, l int) (*Data, error) {
	for dataType, dataTypeString := range dataTypeStrings {
		if dataTypeString == strings.ToUpper(s) {
			return NewDataTypeWith(dataType, l), nil
		}
	}
	return nil, fmt.Errorf("%w: %s", ErrInvalidDataType, s)
}

// DataType returns the column data type.
func (da *Data) DataType() int {
	return da.Type
}

// DataLength returns the column data length.
func (da *Data) DataLength() int {
	return da.Length
}

// String returns the string representation.
func (da *Data) String() string {
	s, ok := dataTypeStrings[da.Type]
	if !ok {
		return ""
	}
	if 0 < da.Length {
		s += fmt.Sprintf("(%d)", da.Length)
	}
	return s
}
