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

// Where represents a where condition.
type Where struct {
	expr Expr
}

// NewWhereWith returns a new where condition instance with the specified parameters.
func NewWhereWith(expr Expr) *Where {
	return &Where{
		expr: expr,
	}
}

// Where returns the table.
func (w *Where) Where() *Where {
	return w
}

// Expr returns the top expression.
func (w *Where) Expr() Expr {
	return w.expr
}

// String returns the string representation.
func (w *Where) String() string {
	return ""
}
