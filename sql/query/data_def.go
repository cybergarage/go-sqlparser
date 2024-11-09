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

// DataDef represents a data definition interface.
type DataDef interface {
	// DataType returns the column data type.
	DataType() DataType
	// DataTypeLen returns the column data type length.
	DataTypeLen() int
	// String returns the string representation.
	String() string
}

// dataDef represents a data definition.
type dataDef struct {
	Type   DataType
	Length int
}

// NewDataDef returns a new DataType instance with the specified type and length.
func NewDataDef(t DataType, l int) DataDef {
	return &dataDef{
		Type:   t,
		Length: l,
	}
}

// NewUnknownDataDef returns a new unknown data type instance.
func NewUnknownDataDef() DataDef {
	return &dataDef{
		Type:   UnknownData,
		Length: 0,
	}
}

// NewDataDefFrom returns the data type of the specified string.
func NewDataDefFrom(s string, l int) (DataDef, error) {
	us := strings.ToUpper(s)
	for dataType, dataTypeString := range dataTypeStrings {
		if dataTypeString == us {
			return NewDataDef(dataType, l), nil
		}
	}
	for dataType, dataTypeString := range dataTypeStrings {
		dataTypeString = strings.ReplaceAll(dataTypeString, " ", "")
		if dataTypeString == us {
			return NewDataDef(dataType, l), nil
		}
	}
	return nil, fmt.Errorf("%w data type : %s", ErrInvalid, s)
}

// DataType returns the column data type.
func (da *dataDef) DataType() DataType {
	return da.Type
}

// DataTypeLen returns the column data type length.
func (da *dataDef) DataTypeLen() int {
	return da.Length
}

// String returns the string representation.
func (da *dataDef) String() string {
	s, ok := dataTypeStrings[da.Type]
	if !ok {
		return ""
	}
	if 0 < da.Length {
		s += fmt.Sprintf("(%d)", da.Length)
	}
	return s
}
