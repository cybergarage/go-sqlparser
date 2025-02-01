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

import "strings"

// DataType represents a data type.
type DataType uint8

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

// NewDataTypeFrom create a data type from the specified value.
func NewDataTypeFrom(a any) (DataType, error) {
	var s string
	switch v := a.(type) {
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		return UnknownData, newErrInvalidDataType(a)
	}
	us := strings.ToUpper(s)
	for dataType, dataTypeString := range dataTypeStrings {
		if dataTypeString == us {
			return dataType, nil
		}
		dataTypeString = strings.ReplaceAll(dataTypeString, " ", "")
		if dataTypeString == us {
			return dataType, nil
		}
	}
	return UnknownData, newErrInvalidDataType(a)
}

// String returns the string representation.
func (t DataType) String() string {
	s, ok := dataTypeStrings[t]
	if !ok {
		return ""
	}
	return s
}
