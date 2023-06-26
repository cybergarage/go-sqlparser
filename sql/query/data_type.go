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

import "fmt"

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

// NewDataTypeFrom returns the data type of the specified string.
func NewDataTypeFrom(s string, l int) (*DataType, error) {
	switch s {
	case "BIGINT":
		return &DataType{BigIntData, l}, nil
	case "BINARY":
		return &DataType{BinaryData, l}, nil
	case "BIT":
		return &DataType{BitData, l}, nil
	case "BLOB":
		return &DataType{BlobData, l}, nil
	case "BOOLEAN":
		return &DataType{BooleanData, l}, nil
	case "CHAR":
		return &DataType{CharData, l}, nil
	case "CHARACTER":
		return &DataType{CharacterData, l}, nil
	case "CLOB":
		return &DataType{ClobData, l}, nil
	case "DATE":
		return &DataType{DateData, l}, nil
	case "DECIMAL":
		return &DataType{DecimalData, l}, nil
	case "DOUBLE":
		return &DataType{DoubleData, l}, nil
	case "FLOAT":
		return &DataType{FloatData, l}, nil
	case "INT":
		return &DataType{IntData, l}, nil
	case "INTEGER":
		return &DataType{IntegerData, l}, nil
	case "LONGBLOB":
		return &DataType{LongBlobData, l}, nil
	case "LONGTEXT":
		return &DataType{LongTextData, l}, nil
	case "MEDIUMBLOB":
		return &DataType{MediumBlobData, l}, nil
	case "MEDIUMINT":
		return &DataType{MediumIntData, l}, nil
	case "MEDIUMTEXT":
		return &DataType{MediumTextData, l}, nil
	case "NUMERIC":
		return &DataType{NumericData, l}, nil
	case "REAL":
		return &DataType{RealData, l}, nil
	case "SET":
		return &DataType{SetData, l}, nil
	case "SMALLINT":
		return &DataType{SmallIntData, l}, nil
	case "TEXT":
		return &DataType{TextData, l}, nil
	case "TIME":
		return &DataType{TimeData, l}, nil
	case "TIMESTAMP":
		return &DataType{TimeStampData, l}, nil
	case "TINYBLOB":
		return &DataType{TinyBlobData, l}, nil
	case "TINYINT":
		return &DataType{TinyIntData, l}, nil
	case "TINYTEXT":
		return &DataType{TinyTextData, l}, nil
	case "VARBINARY":
		return &DataType{VarBinaryData, l}, nil
	case "VARCHAR":
		return &DataType{VarCharData, l}, nil
	case "VARCHARACTER":
		return &DataType{VarCharacterData, l}, nil
	case "YEAR":
		return &DataType{YearData, l}, nil
	default:
		return nil, fmt.Errorf("%w: %s", ErrInvalidDataType, s)
	}
}

// DataType returns the column data type.
func (da *DataType) DataType() int {
	return da.Type
}

// DataLength returns the column data length.
func (da *DataType) DataLength() int {
	return da.Length
}
