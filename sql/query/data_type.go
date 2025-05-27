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
	"strings"

	"github.com/cybergarage/go-sqlparser/sql/fn"
)

// DataType represents a data type.
type DataType uint8

const (
	UnknownData DataType = iota
	BigIntType
	BinaryType
	BitType
	BlobType
	BooleanType
	CharType
	CharacterType
	ClobType
	DateType
	DateTimeType
	DecimalType
	DoubleType
	DoublePrecisionType
	FloatType
	IntType
	IntegerType
	LongBlobType
	LongTextType
	MediumBlobType
	MediumIntType
	MediumTextType
	NumericType
	RealType
	SetType
	SmallIntType
	TextType
	TimeType
	TimeStampType
	TinyBlobType
	TinyIntType
	TinyTextType
	VarBinaryType
	VarCharType
	VarCharacterType
	YearType
	// PostgreSQL
	// https://www.postgresql.org/docs/current/datatype.html
	SerialType
	BigSerialType
	SmallSerialType
)

var dataTypeStrings = map[DataType]string{
	BigIntType:          "BIGINT",
	BinaryType:          "BINARY",
	BitType:             "BIT",
	BlobType:            "BLOB",
	BooleanType:         "BOOLEAN",
	CharType:            "CHAR",
	CharacterType:       "CHARACTER",
	ClobType:            "CLOB",
	DateType:            "DATE",
	DateTimeType:        "DATETIME",
	DecimalType:         "DECIMAL",
	DoubleType:          "DOUBLE",
	DoublePrecisionType: "DOUBLE PRECISION",
	FloatType:           "FLOAT",
	IntType:             "INT",
	IntegerType:         "INTEGER",
	LongBlobType:        "LONGBLOB",
	LongTextType:        "LONGTEXT",
	MediumBlobType:      "MEDIUMBLOB",
	MediumIntType:       "MEDIUMINT",
	MediumTextType:      "MEDIUMTEXT",
	NumericType:         "NUMERIC",
	RealType:            "REAL",
	SetType:             "SET",
	SmallIntType:        "SMALLINT",
	TextType:            "TEXT",
	TimeType:            "TIME",
	TimeStampType:       "TIMESTAMP",
	TinyBlobType:        "TINYBLOB",
	TinyIntType:         "TINYINT",
	TinyTextType:        "TINYTEXT",
	VarBinaryType:       "VARBINARY",
	VarCharType:         "VARCHAR",
	VarCharacterType:    "VARCHARACTER",
	YearType:            "YEAR",
	SerialType:          "SERIAL",
	BigSerialType:       "BIGSERIAL",
	SmallSerialType:     "SMALLSERIAL",
}

// NewDataTypeFrom create a data type from the specified value.
func NewDataTypeFrom(a any) (DataType, error) {
	var s string
	switch v := a.(type) {
	case string:
		s = v
	case []byte:
		s = string(v)
	case fn.Function:
		return NewDataTypeForFunction(v)
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

// NewDataTypeForFunction creates a data type for the specified function.
func NewDataTypeForFunction(fx Function) (DataType, error) {
	if fx == nil {
		return UnknownData, newErrInvalidDataType(fx)
	}
	switch fx.Type() {
	case fn.MathFunction:
		return FloatType, nil
	case fn.AggregateFunction:
		return FloatType, nil
	case fn.CastFunction:
		return FloatType, nil
	case fn.ArithFunction:
		return FloatType, nil
	default:
		return UnknownData, newErrInvalidDataType(fx)
	}
}

// String returns the string representation.
func (t DataType) String() string {
	s, ok := dataTypeStrings[t]
	if !ok {
		return ""
	}
	return s
}
