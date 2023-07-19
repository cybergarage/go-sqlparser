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

import "strings"

// IndexType represents a index type.
type IndexType int

const (
	UnknownIndex   IndexType = 0
	PrimaryIndex   IndexType = 1
	SecondaryIndex IndexType = 2
)

// String returns the string representation.
func (t IndexType) String() string {
	switch t {
	case PrimaryIndex:
		return "PRIMARY KEY"
	case SecondaryIndex:
		return "SECONDARY KEY"
	default:
		return "UNKNOWN KEY"
	}
}

// Index represents a index.
type Index struct {
	name string
	typ  IndexType
	ColumnList
}

// NewIndexWith returns a new index instance.
func NewIndexWith(name string, t IndexType, columns ColumnList) *Index {
	idx := &Index{
		name:       name,
		typ:        t,
		ColumnList: columns,
	}
	return idx
}

// NewPrimaryIndexWith returns a new primary index instance.
func NewPrimaryIndexWith(columns ColumnList) *Index {
	return NewIndexWith("", PrimaryIndex, columns)
}

// NewSecondaryIndexWith returns a new secondary index instance.
func NewSecondaryIndexWith(name string, columns ColumnList) *Index {
	return NewIndexWith(name, SecondaryIndex, columns)
}

// Name returns the index name.
func (idx *Index) Name() string {
	return idx.name
}

// Type returns the index type.
func (idx *Index) Type() IndexType {
	return idx.typ
}

// String returns the index string representation.
func (idx *Index) String() string {
	return idx.name
}

// DefString returns the index definition string representation.
func (idx *Index) DefString() string {
	s := idx.typ.String()
	s += " ("
	s += strings.Join(idx.ColumnList.Names(), ", ")
	s += ")"
	return s
}
