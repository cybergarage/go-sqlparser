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

// Selector represents a selector in a select query.
type Selector interface {
	// Name returns the name of the selector.
	Name() string
	// IsFunction returns true whether the column is a function.
	IsFunction() bool
	// Function returns the function if the column is a function.
	Function() (Function, bool)
	// Arguments returns the executor arguments.
	Arguments() Arguments
	// String returns the string representation of the selector.
	String() string
	// SelectorHelper provides additional methods for selectors.
	SelectorHelper
}
