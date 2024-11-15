// Copyright (C) 2024 The go-mysql Authors. All rights reserved.
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

package resultset

// Row represents a row interface.
type Row interface {
	// Object returns the row object.
	Object() map[string]any
	// Values returns the row values.
	Values() []any
	// ValueAt returns the row value at the specified index.
	ValueAt(int) (any, error)
	// ValueBy returns the row value by the specified name.
	ValueBy(string) (any, error)
	// Scan scans the values.
	Scan(...any) error
	// ScanAt scans the value at the specified index.
	ScanAt(int, any) error
	// ScanBy scans the value by the specified name.
	ScanBy(string, any) error
}
