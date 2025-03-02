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
)

// ColumnDef represents a data definition interface.
type ColumnDef interface {
	// Constraint returns the column constrains.
	Constraint() Constraint
	// DataType returns the column data type.
	DataType() DataType
	// DataTypeSize returns the column data type size.
	DataTypeSize() int
	// String returns the string representation.
	String() string
}

// dataDef represents a data definition.
type dataDef struct {
	consts Constraint
	Type   DataType
	Length int
}

// NewDataDef returns a new DataType instance with the specified type and length.
func NewDataDef(t DataType, l int) ColumnDef {
	return &dataDef{
		consts: ConstraintNone,
		Type:   t,
		Length: l,
	}
}

// NewUnknownDataDef returns a new unknown data type instance.
func NewUnknownDataDef() ColumnDef {
	return &dataDef{
		consts: ConstraintNone,
		Type:   UnknownData,
		Length: 0,
	}
}

// NewDataDefWith returns the data type of the specified string.
func NewDataDefWith(s string, l int) (ColumnDef, error) {
	dt, err := NewDataTypeFrom(s)
	if err != nil {
		return nil, err
	}
	return NewDataDef(dt, l), nil
}

// Constraint returns the column constrains.
func (da *dataDef) Constraint() Constraint {
	return da.consts
}

// DataType returns the column data type.
func (da *dataDef) DataType() DataType {
	return da.Type
}

// DataTypeSize returns the column data type length.
func (da *dataDef) DataTypeSize() int {
	return da.Length
}

// String returns the string representation.
func (da *dataDef) String() string {
	s := da.Type.String()
	if 0 < da.Length {
		s += fmt.Sprintf("(%d)", da.Length)
	}
	return s
}
