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

// Condition represents a where condition interface.
type Condition interface {
	// HasConditions returns true if the condition has conditions.
	HasConditions() bool
	// Expr returns the top expression.
	Expr() Expr
	// String returns the string representation.
	String() string
}

// condition represents a where condition.
type condition struct {
	expr Expr
}

// NewCondition returns a new where condition instance.
func NewCondition() Condition {
	return &condition{
		expr: nil,
	}
}

// NewConditionWith returns a new where condition instance with the specified parameters.
func NewConditionWith(expr Expr) Condition {
	return &condition{
		expr: expr,
	}
}

// Where returns the table.
func (w *condition) Where() *condition {
	return w
}

// Expr returns the top expression.
func (w *condition) Expr() Expr {
	return w.expr
}

// String returns the string representation.
func (w *condition) String() string {
	if w.expr == nil {
		return ""
	}
	return w.expr.String()
}

// HasConditions returns true if the condition has conditions.
func (w *condition) HasConditions() bool {
	return w.expr != nil
}
