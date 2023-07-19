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

// CompExprOperator is an enum for CompExpr.Operator
type CompExprOperator uint8

// Constants for Enum Type - CompExprOperator
const (
	EqualOp CompExprOperator = iota
	NotEqualOp
	LessThanOp
	GreaterThanOp
	LessEqualOp
	GreaterEqualOp
	InOp
	NotInOp
)

var compExprOpsStrings = map[CompExprOperator]string{
	EqualOp:        "=",
	NotEqualOp:     "!=", // "<>",
	LessThanOp:     "<",
	GreaterThanOp:  ">",
	LessEqualOp:    "<=",
	GreaterEqualOp: ">=",
	InOp:           "IN",
	NotInOp:        "NOT IN",
}

// String returns the string representation.
func (t CompExprOperator) String() string {
	s, ok := compExprOpsStrings[t]
	if !ok {
		return ""
	}
	return s
}

// CmpExpr represents an comparsion expression.
type CmpExpr struct {
	op    CompExprOperator
	left  *Column
	right *Literal
}

// NewCmpExpr returns a new CompExpr instance with the specified parameters.
func NewCmpExpr(op CompExprOperator, left *Column, right *Literal) *CmpExpr {
	return &CmpExpr{
		op:    op,
		left:  left,
		right: right,
	}
}

// Operator returns the operator.
func (expr *CmpExpr) Operator() CompExprOperator {
	return expr.op
}

// Left returns the left column.
func (expr *CmpExpr) Left() *Column {
	return expr.left
}

// Right returns the right literal.
func (expr *CmpExpr) Right() *Literal {
	return expr.right
}

// String returns the index string representation.
func (expr *CmpExpr) String() string {
	return expr.left.String() + " " + expr.op.String() + " " + expr.right.String()
}
