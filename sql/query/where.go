// Copyright (C) 2019 Satoshi Konno. All rights reserved.
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

// Condition represents a where condition.
type Condition struct {
	expr Expr
}

// NewConditionWith returns a new where condition instance with the specified parameters.
func NewConditionWith(expr Expr) *Condition {
	return &Condition{
		expr: expr,
	}
}

// Where returns the table.
func (w *Condition) Where() *Condition {
	return w
}

// Expr returns the top expression.
func (w *Condition) Expr() Expr {
	return w.expr
}

// String returns the string representation.
func (w *Condition) String() string {
	return w.expr.String()
}

// IsEmpty returns true if the where condition is empty.
func (w *Condition) IsEmpty() bool {
	if w == nil {
		return true
	}
	return w.expr == nil
}
