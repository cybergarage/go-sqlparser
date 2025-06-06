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
// limitations under the License..

package query

import (
	"errors"
	"fmt"
)

// ErrInvalid is returned when the value is invalid.
var ErrInvalid = errors.New("invalid")

// ErrNotFound is returned when the data type is invalid.
var ErrNotFound = errors.New("not found")

// ErrNotSupported is returned when the function is not supported.
var ErrNotSupported = errors.New("not supported")

// ErrNotSet is returned when the value is not set.
var ErrNotSet = errors.New("not set")

// NewErrStatementInvalid  returns a new error for invalid statements.
func NewErrStatementInvalid(stmt Statement) error {
	return fmt.Errorf("statement (%s) %w", stmt, ErrInvalid)
}

// NewErrStatementNotSupported returns a new error for unsupported statements.
func NewErrStatementNotSupported(stmt Statement) error {
	return fmt.Errorf("statement (%s) %w", stmt, ErrNotSupported)
}

func newErrColumnNotFound(name string) error {
	return fmt.Errorf("column (%s) %w", name, ErrNotFound)
}

func newErrIndexNotFound(name string) error {
	return fmt.Errorf("index (%s) %w", name, ErrNotFound)
}

func newErrColumnIndexOutOfRange(n int) error {
	return fmt.Errorf("column index (%d) %w", n, ErrNotFound)
}

func newErrNotFoundFunction(name string) error {
	return fmt.Errorf("function (%s) %w", name, ErrNotFound)
}

func newErrInvalidDataType(s any) error {
	return fmt.Errorf("data type (%s) %w", s, ErrInvalid)
}

func newErrInvaidDataDef(s any) error {
	return fmt.Errorf("data def (%s) %w", s, ErrInvalid)
}
