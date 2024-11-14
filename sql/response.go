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

package sql

import (
	"github.com/cybergarage/go-sqlparser/sql/query/response/resultset"
)

// ResultSetColumn represents a column interface in a resultset.
type ResultSetColumn = resultset.Column

// ResultSetSchema represents a schema interface in a resultset.
type ResultSetSchema = resultset.Schema

// Row represents a row interface.
type ResultSetRow = resultset.Row

// ResultSet represents a response resultset interface.
type ResultSet = resultset.ResultSet
