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
	std_strings "strings"

	"github.com/cybergarage/go-sqlparser/sql/util/strings"
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
	switch std_strings.ToUpper(order) {
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

// Order represents an ORDER BY clause.
type Order struct {
	column string
	order  OrderType
}

// NewOrder returns a new order instance.
func NewOrder() *Order {
	return &Order{
		column: "",
		order:  OrderNone,
	}
}

func NewOrderWith(column string, order OrderType) *Order {
	return &Order{
		column: column,
		order:  order,
	}
}

// Column returns the column name.
func (order *Order) Column() string {
	return order.column
}

// Order returns the order.
func (order *Order) Order() OrderType {
	return order.order
}

// IsNone returns true whether the order is none.
func (order *Order) IsNone() bool {
	return order.order.IsNone()
}

// IsAsc returns true whether the order is asc.
func (order *Order) IsAsc() bool {
	return order.order.IsAsc()
}

// IsDesc returns true whether the order is desc.
func (order *Order) IsDesc() bool {
	return order.order.IsDesc()
}

// String returns the string representation.
func (order *Order) String() string {
	if order.IsNone() {
		return ""
	}
	return strings.JoinWithSpace(
		[]string{
			order.column,
			order.order.String(),
		})
}

// OrderBy represents an ORDER BY clause.
type OrderBy struct {
	column string
	order  OrderType
}

// NewOrderBy returns a new OrderBy instance.
func NewOrderBy() *OrderBy {
	return &OrderBy{
		column: "",
		order:  OrderNone,
	}
}

func NewOrderByWith(column string, order OrderType) *OrderBy {
	return &OrderBy{
		column: column,
		order:  order,
	}
}

// SetColumn sets the column name.
func (orderBy *OrderBy) SetColumn(name string) *OrderBy {
	orderBy.column = name
	return orderBy
}

// SetOrder sets the order.
func (orderBy *OrderBy) SetOrder(order OrderType) *OrderBy {
	orderBy.order = order
	return orderBy
}

// Column returns the column name.
func (orderBy *OrderBy) Column() string {
	return orderBy.column
}

// Order returns the order.
func (orderBy *OrderBy) Order() OrderType {
	return orderBy.order
}

// IsNone returns true whether the order is none.
func (orderBy *OrderBy) IsNone() bool {
	return orderBy.order.IsNone()
}

// IsAsc returns true whether the order is asc.
func (orderBy *OrderBy) IsAsc() bool {
	return orderBy.order.IsAsc()
}

// IsDesc returns true whether the order is desc.
func (orderBy *OrderBy) IsDesc() bool {
	return orderBy.order.IsDesc()
}

// String returns the string representation.
func (orderBy *OrderBy) String() string {
	if orderBy.IsNone() {
		return ""
	}
	return strings.JoinWithSpace(
		[]string{
			"ORDER BY",
			orderBy.column,
			orderBy.order.String(),
		})
}
