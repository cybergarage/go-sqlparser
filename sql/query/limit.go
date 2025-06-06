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

	"github.com/cybergarage/go-sqlparser/sql/util/strings"
)

// Limit represents a LIMIT interface.
type Limit interface {
	// Offset returns the offset.
	Offset() uint
	// Limit returns the limit.
	Limit() uint
	// String returns the string representation.
	String() string
}

// limitParam represents a LIMIT clause.
type limitParam struct {
	offset uint
	limit  uint
}

// NewLimit returns a new limit instance.
func NewLimit() Limit {
	return &limitParam{
		offset: 0,
		limit:  0,
	}
}

// NewLimitWith returns a new limitParam instance with the specified offset and limit.
func NewLimitWith(offset uint, limit uint) *limitParam {
	return &limitParam{
		offset: offset,
		limit:  limit,
	}
}

// SetLimit sets the limit.
func (limit *limitParam) SetLimit(n uint) *limitParam {
	limit.limit = n
	return limit
}

// SetOffset sets the offset.
func (limit *limitParam) SetOffset(n uint) *limitParam {
	limit.offset = n
	return limit
}

// Offset returns the offset.
func (limit *limitParam) Offset() uint {
	return limit.offset
}

// limitParam returns the limit.
func (limit *limitParam) Limit() uint {
	return uint(limit.limit)
}

// String returns the string representation.
func (limit *limitParam) String() string {
	strs := []string{}
	if 0 < limit.limit {
		strs = append(strs, "LIMIT", fmt.Sprintf("%d", limit.limit))
	}
	if 0 < limit.offset {
		strs = append(strs, "OFFSET", fmt.Sprintf("%d", limit.offset))
	}
	return strings.JoinWithSpace(strs)
}
