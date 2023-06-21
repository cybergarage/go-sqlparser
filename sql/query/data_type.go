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
type DataType int

const (
	BigIntData = iota
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

// GetDataTypeString returns the data type of the specified string.
func GetDataTypeFrom(s string) (DataType, error) {
	switch s {
	case "BIGINT":
		return BigIntData, nil
	case "BINARY":
		return BinaryData, nil
	case "BIT":
		return BitData, nil
	case "BLOB":
		return BlobData, nil
	case "BOOLEAN":
		return BooleanData, nil
	case "CHAR":
		return CharData, nil
	case "CHARACTER":
		return CharacterData, nil
	case "CLOB":
		return ClobData, nil
	case "DATE":
		return DateData, nil
	case "DECIMAL":
		return DecimalData, nil
	case "DOUBLE":
		return DoubleData, nil
	case "FLOAT":
		return FloatData, nil
	case "INT":
		return IntData, nil
	case "INTEGER":
		return IntegerData, nil
	case "LONGBLOB":
		return LongBlobData, nil
	case "LONGTEXT":
		return LongTextData, nil
	case "MEDIUMBLOB":
		return MediumBlobData, nil
	case "MEDIUMINT":
		return MediumIntData, nil
	case "MEDIUMTEXT":
		return MediumTextData, nil
	case "NUMERIC":
		return NumericData, nil
	case "REAL":
		return RealData, nil
	case "SET":
		return SetData, nil
	case "SMALLINT":
		return SmallIntData, nil
	case "TEXT":
		return TextData, nil
	case "TIME":
		return TimeData, nil
	case "TIMESTAMP":
		return TimeStampData, nil
	case "TINYBLOB":
		return TinyBlobData, nil
	case "TINYINT":
		return TinyIntData, nil
	case "TINYTEXT":
		return TinyTextData, nil
	case "VARBINARY":
		return VarBinaryData, nil
	case "VARCHAR":
		return VarCharData, nil
	case "VARCHARACTER":
		return VarCharacterData, nil
	case "YEAR":
		return YearData, nil
	default:
		return -1, fmt.Errorf("invalid data type: %s", s)
	}
}
