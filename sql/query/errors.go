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

// ErrInvalidDataType is returned when the data type is invalid.
var ErrInvalidDataType = errors.New("invalid data type")

// ErrNotFound is returned when the data type is invalid.
var ErrNotFound = errors.New("not found")

func newErrColumnNotFound(name string) error {
	return fmt.Errorf("column (%s) %w", name, ErrNotFound)
}

func newErrColumnIndexOutOfRange(n int) error {
	return fmt.Errorf("column index (%d) %w", n, ErrNotFound)
}
