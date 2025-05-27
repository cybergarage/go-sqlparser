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

package fn_test

import (
	"testing"

	"github.com/cybergarage/go-sqlparser/sql/fn"
)

func TestRegisteredFunctionTypes(t *testing.T) {
	names := fn.RegisteredFunctionNames()
	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			_, err := fn.NewFunctionTypeForName(name)
			if err != nil {
				t.Errorf("Failed to create function type for %s: %v", name, err)
			}
		})
	}
}
