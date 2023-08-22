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
	"strconv"
)

// Limit represents a LIMIT clause.
type Limit struct {
	offset int
	limit  int
}

// NewLimit returns a new Limit instance.
func NewLimit() *Limit {
	return &Limit{
		offset: 0,
		limit:  0,
	}
}

// NewLimitWith returns a new Limit instance with the specified offset and limit.
func NewLimitWith(offset int, limit int) *Limit {
	return &Limit{
		offset: offset,
		limit:  limit,
	}
}

// SetLimit sets the limit.
func (limit *Limit) SetLimit(n int) *Limit {
	limit.limit = n
	return limit
}

// SetOffset sets the offset.
func (limit *Limit) SetOffset(n int) *Limit {
	limit.offset = n
	return limit
}

// Offset returns the offset.
func (limit *Limit) Offset() int {
	return limit.offset
}

// Limit returns the limit.
func (limit *Limit) Limit() int {
	return limit.limit
}

// String returns the string representation.
func (limit *Limit) String() string {
	if limit.offset <= 0 && limit.limit <= 0 {
		return ""
	}
	if limit.offset <= 0 {
		return "LIMIT " + strconv.Itoa(limit.limit)
	}
	return "LIMIT " + strconv.Itoa(limit.limit) + "OFFSET " + strconv.Itoa(limit.offset)
}
