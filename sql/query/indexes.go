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
)

// IndexList represents an index array.
type IndexList []*Index

// NewIndexList returns a new Indexes instance.
func NewIndexList() IndexList {
	return make(IndexList, 0)
}

// Indexes returns an index array.
func (indexes IndexList) Indexes() IndexList {
	return indexes
}

// DefString returns the index definition string representation.
func (indexes IndexList) DefString() string {
	elems := make([]string, len(indexes))
	for n, index := range indexes {
		elems[n] = index.DefString()
	}
	return strings.Join(elems, ",")
}
