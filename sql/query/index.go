// Copyright (C) 2022 The go-sqlparser Authors All rights reserved.
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
)

const (
	indexNameSep = "_"
)

// Index represents a index interface.
type Index interface {
	// Name returns the index name.
	Name() string
	// Type returns the index type.
	Type() IndexType
	// String returns the index string representation.
	Columns() Columns
	// DefinitionString returns the index definition string representation.
	DefinitionString() string
}

// index represents a index.
type index struct {
	name    string
	typ     IndexType
	columns Columns
}

// NewIndexWith returns a new index instance.
func NewIndexWith(name string, t IndexType, columns Columns) Index {
	idx := &index{
		name:    name,
		typ:     t,
		columns: columns,
	}
	return idx
}

// NewPrimaryIndexWith returns a new primary index instance.
func NewPrimaryIndexWith(columns Columns) Index {
	idxName := strings.Join(columns.Names(), indexNameSep)
	return NewIndexWith(idxName, PrimaryIndex, columns)
}

// NewSecondaryIndexWith returns a new secondary index instance.
func NewSecondaryIndexWith(name string, columns Columns) Index {
	return NewIndexWith(name, SecondaryIndex, columns)
}

// Name returns the index name.
func (idx *index) Name() string {
	return idx.name
}

// Type returns the index type.
func (idx *index) Type() IndexType {
	return idx.typ
}

// Columns returns the index columns.
func (idx *index) Columns() Columns {
	return idx.columns
}

// String returns the index string representation.
func (idx *index) String() string {
	return idx.name
}

// DefinitionString returns the index definition string representation.
func (idx *index) DefinitionString() string {
	s := idx.typ.String()
	s += " ("
	s += strings.Join(idx.Columns().Names(), ", ")
	s += ")"
	return s
}
