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

const (
	BoolObjectId            = 16
	ByteaObjectId           = 17
	CharObjectId            = 18
	NameObjectId            = 19
	Int8ObjectId            = 20
	Int2ObjectId            = 21
	Int2VectorObjectId      = 22
	Int4ObjectId            = 23
	RegProcObjectId         = 24
	TextObjectId            = 25
	OidObjectId             = 26
	TidObjectId             = 27
	XidObjectId             = 28
	CidObjectId             = 29
	XmlObjectId             = 142
	PointObjectId           = 600
	LsegObjectId            = 601
	PathObjectId            = 602
	BoxObjectId             = 603
	PolygonObjectId         = 604
	LineObjectId            = 628
	LineArrayObjectId       = 629
	CircleObjectId          = 718
	CircleArrayObjectId     = 719
	MacaddrObjectId         = 829
	Macaddr8ObjectId        = 774
	InetObjectId            = 869
	InetArrayObjectId       = 1040
	CidrObjectId            = 650
	CidrArrayObjectId       = 651
	Float4ObjectId          = 700
	Float8ObjectId          = 701
	UnknownObjectId         = 705
	AbstimeObjectId         = 702
	ReltimeObjectId         = 703
	TintervalObjectId       = 704
	PolygonArrayObjectId    = 628
	OidvectorObjectId       = 30
	BpcharObjectId          = 1042
	VarcharObjectId         = 1043
	DateObjectId            = 1082
	TimeObjectId            = 1083
	TimestampObjectId       = 1114
	TimestampTzObjectId     = 1184
	IntervalObjectId        = 1186
	TimeTzObjectId          = 1266
	BitObjectId             = 1560
	VarbitObjectId          = 1562
	NumericObjectId         = 1700
	RefcursorObjectId       = 1790
	RegprocedureObjectId    = 2202
	RegoperObjectId         = 2203
	RegoperatorObjectId     = 2204
	RegclassObjectId        = 2205
	RegtypeObjectId         = 2206
	RecordObjectId          = 2249
	CstringObjectId         = 2275
	AnyObjectId             = 2276
	AnyarrayObjectId        = 2277
	VoidObjectId            = 2278
	TriggerObjectId         = 2279
	LanguageHandlerObjectId = 2280
	InternalObjectId        = 2281
	OpaqueObjectId          = 2282
	AnyelementObjectId      = 2283
	AnynonarrayObjectId     = 2776
	GeometryObjectId        = 3614
	GeograpyObjectId        = 4326
)

// NewDataWith returns a new DataType instance with the specified type and length.
func NewDataWith(t DataType, l int) *Data {
	return &Data{
		Type:   t,
		Length: l,
	}
}

// NewDataFrom returns the data type of the specified string.
func NewDataFrom(s string, l int) (*Data, error) {
	for dataType, dataTypeString := range dataTypeStrings {
		if dataTypeString == strings.ToUpper(s) {
			return NewDataWith(dataType, l), nil
		}
	}
	return nil, fmt.Errorf("%w: %s", ErrInvalidDataType, s)
}

// DataType returns the column data type.
func (da *Data) DataType() DataType {
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
