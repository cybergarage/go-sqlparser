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

// OrderType represents an ordertype.
type OrderType uint8

const (
	// OrderNone represents none order.
	OrderNone OrderType = iota
	// OrderAsc represents asc order.
	OrderAsc
	// OrderDesc represents desc order.
	OrderDesc
)

const (
	orderNoneString = "NONE"
	orderAscString  = "ASC"
	orderDescString = "DESC"
)

// NewOrderTypeWith returns a new OrderType instance.
func NewOrderTypeWith(order string) OrderType {
	switch strings.ToUpper(order) {
	case orderAscString:
		return OrderAsc
	case orderDescString:
		return OrderDesc
	}
	return OrderNone
}

// IsNone returns true whether the order is none.
func (t OrderType) IsNone() bool {
	return t == OrderNone
}

// IsAsc returns true whether the order is asc.
func (t OrderType) IsAsc() bool {
	return t == OrderAsc
}

// IsDesc returns true whether the order is desc.
func (t OrderType) IsDesc() bool {
	return t == OrderDesc
}

// String returns the string representation.
func (t OrderType) String() string {
	switch t {
	case OrderAsc:
		return orderAscString
	case OrderDesc:
		return orderDescString
	default:
		return orderNoneString
	}
}
