// Copyright (C) 2022 The go-sqlparser Authors All rights reserved.
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

// AndExpr represents an AND expression.
type AndExpr struct {
	left  Expr
	right Expr
}

// NewAndExpr returns a new AndExpr.
func NewAndExpr(left Expr, right Expr) *AndExpr {
	return &AndExpr{left: left, right: right}
}

// Left returns the left expression.
func (expr *AndExpr) Left() Expr {
	return expr.left
}

// Right returns the right expression.
func (expr *AndExpr) Right() Expr {
	return expr.right
}

// String returns the index string representation.
func (expr *AndExpr) String() string {
	return expr.left.String() + " AND " + expr.right.String()
}
