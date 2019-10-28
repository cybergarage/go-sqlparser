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

package util

import (
	"testing"
)

const (
	errorInvalidParserObject = "Invalid parser object (%d != %d)"
	errorInvalidStackSize    = "Invalid parser stack size (%d != %d)"
)

func TestNewStack(t *testing.T) {
	s := NewStack()

	if s.Size() != 0 {
		t.Errorf(errorInvalidStackSize, s.Size(), 0)
	}

	for n := 1; n <= 10; n++ {
		var value int
		value = n
		s.PushObject(value)
		if s.Size() != n {
			t.Errorf(errorInvalidStackSize, s.Size(), n)
		}
	}

	for n := 10; n >= 1; n-- {
		obj := s.PopObject()
		value, _ := obj.(int)
		if value != n {
			t.Errorf(errorInvalidParserObject, value, n)
		}
		if s.Size() != (n - 1) {
			t.Errorf(errorInvalidStackSize, s.Size(), (n - 1))
		}
	}

	if s.Size() != 0 {
		t.Errorf(errorInvalidStackSize, s.Size(), 0)
	}
}
