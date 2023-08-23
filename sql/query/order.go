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

import "github.com/cybergarage/go-sqlparser/sql/util/strings"

// Order represents an ordertype.
type Order uint8

const (
	// OrderNone represents none order.
	OrderNone Order = iota
	// OrderAsc represents asc order.
	OrderAsc
	// OrderDesc represents desc order.
	OrderDesc
)

// String returns the string representation.
func (t Order) String() string {
	switch t {
	case OrderAsc:
		return "ASC"
	case OrderDesc:
		return "DESC"
	default:
		return "NONE"
	}
}

// OrderBy represents an ORDER BY clause.
type OrderBy struct {
	column string
	order  Order
}

// NewOrderBy returns a new OrderBy instance.
func NewOrderBy() *OrderBy {
	return &OrderBy{
		column: "",
		order:  OrderNone,
	}
}

func NewOrderByWith(column string, order Order) *OrderBy {
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
func (orderBy *OrderBy) SetOrder(order Order) *OrderBy {
	orderBy.order = order
	return orderBy
}

// Column returns the column name.
func (orderBy *OrderBy) Column() string {
	return orderBy.column
}

// Order returns the order.
func (orderBy *OrderBy) Order() Order {
	return orderBy.order
}

// IsNone returns true whether the order is none.
func (orderBy *OrderBy) IsNone() bool {
	return orderBy.order == OrderNone
}

// IsAsc returns true whether the order is asc.
func (orderBy *OrderBy) IsAsc() bool {
	return orderBy.order == OrderAsc
}

// IsDesc returns true whether the order is desc.
func (orderBy *OrderBy) IsDesc() bool {
	return orderBy.order == OrderDesc
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
