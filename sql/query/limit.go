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

	"github.com/cybergarage/go-sqlparser/sql/util/strings"
)

// Limit represents a LIMIT interface.
type Limit interface {
	// Offset returns the offset.
	Offset() int
	// Limit returns the limit.
	Limit() int
	// String returns the string representation.
	String() string
}

// limitParam represents a LIMIT clause.
type limitParam struct {
	offset int
	limit  int
}

// NewLimit returns a new limit instance.
func NewLimit() Limit {
	return &limitParam{
		offset: 0,
		limit:  0,
	}
}

// NewLimitWith returns a new limitParam instance with the specified offset and limit.
func NewLimitWith(offset int, limit int) *limitParam {
	return &limitParam{
		offset: offset,
		limit:  limit,
	}
}

// SetLimit sets the limit.
func (limit *limitParam) SetLimit(n int) *limitParam {
	limit.limit = n
	return limit
}

// SetOffset sets the offset.
func (limit *limitParam) SetOffset(n int) *limitParam {
	limit.offset = n
	return limit
}

// Offset returns the offset.
func (limit *limitParam) Offset() int {
	return limit.offset
}

// limitParam returns the limit.
func (limit *limitParam) Limit() int {
	return limit.limit
}

// String returns the string representation.
func (limit *limitParam) String() string {
	strs := []string{}
	if 0 < limit.limit {
		strs = append(strs, "LIMIT", strconv.Itoa(limit.limit))
	}
	if 0 < limit.offset {
		strs = append(strs, "OFFSET", strconv.Itoa(limit.offset))
	}
	return strings.JoinWithSpace(strs)
}
