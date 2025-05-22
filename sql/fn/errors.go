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

package fn

import (
	"errors"
	"fmt"
)

// ErrInvalid is returned when the value is invalid.
var ErrInvalid = errors.New("invalid")

// ErrNotFound is returned when the data type is invalid.
var ErrNotFound = errors.New("not found")

// ErrNoData is returned when there is no data.
var ErrNoData = errors.New("no data")

// ErrNotSupported is returned when the function is not supported.
var ErrNotSupported = errors.New("not supported")

// ErrNotImplemented is returned when the function is not implemented.
var ErrNotImplemented = errors.New("not implemented")

func newErrInvalidArguments(name string, args ...any) error {
	return fmt.Errorf("%w arguments %s(%v)", ErrInvalid, name, args)
}

func newErrNotSupportedFunction(name string) error {
	return fmt.Errorf("function (%s) %w", name, ErrNotSupported)
}
