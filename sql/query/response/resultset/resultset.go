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

// ResultSet represents a response resultset interface.
type ResultSet interface {
	// Schema returns the schema.
	Schema() Schema
	// Next returns the next row.
	Next() bool
	// Row returns the current row.
	Row() (Row, error)
	// RowsAffected returns the number of rows affected.
	RowsAffected() uint64
	// Close closes the resultset.
	Close() error
}