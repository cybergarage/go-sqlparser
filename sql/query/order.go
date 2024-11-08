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

// Order represents an ORDER BY interface.
type Order interface {
	Column() string
	Type() OrderType
	IsNone() bool
	IsAsc() bool
	IsDesc() bool
	String() string
}

// orderParam represents an ORDER BY clause.
type orderParam struct {
	column string
	order  OrderType
}

// NewOrder returns a new order instance.
func NewOrder() Order {
	return &orderParam{
		column: "",
		order:  OrderNone,
	}
}

// NewOrderWith returns a new order instance with the specified column and order.
func NewOrderWith(column string, order OrderType) Order {
	return &orderParam{
		column: column,
		order:  order,
	}
}

// Column returns the column name.
func (order *orderParam) Column() string {
	return order.column
}

// Type returns the order type.
func (order *orderParam) Type() OrderType {
	return order.order
}

// IsNone returns true whether the order is none.
func (order *orderParam) IsNone() bool {
	return order.order.IsNone()
}

// IsAsc returns true whether the order is asc.
func (order *orderParam) IsAsc() bool {
	return order.order.IsAsc()
}

// IsDesc returns true whether the order is desc.
func (order *orderParam) IsDesc() bool {
	return order.order.IsDesc()
}

// String returns the string representation.
func (order *orderParam) String() string {
	if order.IsNone() {
		return ""
	}
	return strings.JoinWithSpace(
		[]string{
			order.column,
			order.order.String(),
		})
}

// OrderBy represents an ORDER BY interface.
type OrderBy interface {
	Orders() []Order
	String() string
}

// orderByParam represents an ORDER BY clause.
type orderByParam struct {
	orders []Order
}

// NewOrderBy returns a new orderByParam instance.
func NewOrderBy() OrderBy {
	return &orderByParam{
		orders: []Order{},
	}
}

// NewOrderByWith returns a new orderByParam instance with the specified orders.
func NewOrderByWith(orders []Order) OrderBy {
	return &orderByParam{
		orders: orders,
	}
}

// Orders returns the orders.
func (orderBy *orderByParam) Orders() []Order {
	return orderBy.orders
}

// String returns the string representation.
func (orderBy *orderByParam) String() string {
	if len(orderBy.orders) == 0 {
		return ""
	}
	orderStrs := []string{}
	for _, order := range orderBy.orders {
		orderStrs = append(orderStrs, order.String())
	}
	orderStr := strings.JoinWithComma(orderStrs)
	return strings.JoinWithSpace(
		[]string{
			"ORDER BY",
			orderStr,
		})
}
