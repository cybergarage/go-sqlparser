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

// CmpExprOperator is an enum for CompExpr.Operator
type CmpExprOperator uint8

// Constants for Enum Type - CompExprOperator
const (
	EQ CmpExprOperator = iota
	NEQ
	LT
	GT
	LE
	GE
	IN
	NIN
)

var cmpExprOpsStrings = map[CmpExprOperator]string{
	EQ:  "=",
	NEQ: "!=", // "<>",
	LT:  "<",
	GT:  ">",
	LE:  "<=",
	GE:  ">=",
	IN:  "IN",
	NIN: "NOT IN",
}

// String returns the string representation.
func (t CmpExprOperator) String() string {
	s, ok := cmpExprOpsStrings[t]
	if !ok {
		return ""
	}
	return s
}

// CmpExpr represents an comparsion expression.
type CmpExpr struct {
	op    CmpExprOperator
	left  *Column
	right *Literal
}

// NewCmpExprWith returns a new CompExpr instance with the specified parameters.
func NewCmpExprWith(op CmpExprOperator, left *Column, right *Literal) *CmpExpr {
	return &CmpExpr{
		op:    op,
		left:  left,
		right: right,
	}
}

// NewEQ returns a new CompExpr instance with the specified parameters.
func NewEQWith(left *Column, right *Literal) *CmpExpr {
	return NewCmpExprWith(EQ, left, right)
}

// NewNEQ returns a new CompExpr instance with the specified parameters.
func NewNEQWith(left *Column, right *Literal) *CmpExpr {
	return NewCmpExprWith(NEQ, left, right)
}

// NewLT returns a new CompExpr instance with the specified parameters.
func NewLTWith(left *Column, right *Literal) *CmpExpr {
	return NewCmpExprWith(LT, left, right)
}

// NewGT returns a new CompExpr instance with the specified parameters.
func NewGTWith(left *Column, right *Literal) *CmpExpr {
	return NewCmpExprWith(GT, left, right)
}

// NewLE returns a new CompExpr instance with the specified parameters.
func NewLEWith(left *Column, right *Literal) *CmpExpr {
	return NewCmpExprWith(LE, left, right)
}

// NewGE returns a new CompExpr instance with the specified parameters.
func NewGEWith(left *Column, right *Literal) *CmpExpr {
	return NewCmpExprWith(GE, left, right)
}

// Operator returns the operator.
func (expr *CmpExpr) Operator() CmpExprOperator {
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
