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
)

// Literal represents a constant value.
type Literal struct {
	v any
}

// NullLiteral represents a null value.
var NullLiteral = NewLiteralWith(nil)

// NewLiteralWith returns a new Literal instance with the specified value.
func NewLiteralWith(v any) *Literal {
	return &Literal{
		v: v,
	}
}

// Literal returns the literal.
func (lit *Literal) Literal() *Literal {
	return lit
}

// Value returns the literal value.
func (lit *Literal) Value() any {
	return lit.v
}

// String returns the string representation.
func (lit *Literal) String() string {
	return fmt.Sprintf("%v", lit.v)
}
