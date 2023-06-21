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

// IndexType represents a index type.
type IndexType int

const (
	UnknownIndex   IndexType = 0
	PrimaryIndex   IndexType = 1
	SecondaryIndex IndexType = 2
)

// Index represents a index.
type Index struct {
	name string
	typ  IndexType
	Columns
}

// NewIndex returns a new Index instance.
func NewIndexWith(name string, t IndexType, colums Columns) *Index {
	idx := &Index{
		name:    name,
		typ:     UnknownIndex,
		Columns: colums,
	}
	return idx
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
