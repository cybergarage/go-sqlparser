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

// ColumnHelper provides additional methods for columns in a query.
type ColumnHelper interface {
	// ValueString returns the column value string.
	ValueString() string
	// IsAsterisk returns true if the column is an asterisk.
	IsAsterisk() bool
	// IsPlaceHolder returns true whether the column is a place holder.
	IsPlaceHolder() bool
	// IsFunction returns true whether the column is a function.
	IsFunction() bool
	// Function returns the function if the column is a function.
	Function() (Function, bool)
	// Arguments returns the executor arguments.
	Arguments() []any
}
