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
	"fmt"

	"github.com/cybergarage/go-safecast/safecast"
	"github.com/cybergarage/go-sqlparser/sql/util/strings"
)

// LiteralType represents a literal type.
type LiteralType uint8

const (
	UnknownLiteral = iota
	StringLiteral
)

// LiteralOption represents a literal option function.
type LiteralOption = func(*Literal)

// Literal represents a constant value.
type Literal struct {
	v any
	t LiteralType
}

// NullLiteral represents a null value.
var NullLiteral = NewLiteralWith(nil)

// NewLiteral returns a new Literal instance.
func NewLiteral() *Literal {
	return &Literal{
		v: nil,
		t: UnknownLiteral,
	}
}

// NewLiteralWith returns a new Literal instance with the specified value.
func NewLiteralWith(v any, opts ...LiteralOption) *Literal {
	l := &Literal{
		v: v,
		t: UnknownLiteral,
	}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

// WithSchemaColumns returns a schema option to set the columns.
func WithLiteralType(t LiteralType) func(*Literal) {
	return func(lit *Literal) {
		lit.t = t
	}
}

// HasValue returns true whether the literal has a value.
func (lit *Literal) HasValue() bool {
	return lit != nil
}

// Equal returns true whether the literal value is equal to the specified value.
func (lit *Literal) Equal(v any) bool {
	if lit == nil || lit.v == nil {
		return false
	}
	return safecast.Equal(lit.v, v)
}

// IsAsterisk returns true whether the literal is an asterisk.
func (lit *Literal) IsAsterisk() bool {
	return lit.Equal("*")
}

// IsPlaceHolder returns true whether the literal is a place holder.
func (lit *Literal) IsPlaceHolder() bool {
	return lit.Equal("?")
}

// SetValue sets a value.
func (lit *Literal) SetValue(v any) *Literal {
	lit.v = v
	return lit
}

// Value returns the literal value.
func (lit *Literal) Value() any {
	return lit.v
}

// SetValueType sets a literal type.
func (lit *Literal) SetValueType(t LiteralType) *Literal {
	lit.t = t
	return lit
}

// ValueType returns the literal type.
func (lit *Literal) ValueType() LiteralType {
	return lit.t
}

// ValueString returns the string representation.
func (lit *Literal) ValueString() string {
	if lit == nil {
		return ""
	}
	switch lit.t {
	case StringLiteral:
		return strings.EscapeString(fmt.Sprintf("%v", lit.v))
	}
	return fmt.Sprintf("%v", lit.v)
}

// String returns the string representation.
func (lit *Literal) String() string {
	return lit.ValueString()
}
