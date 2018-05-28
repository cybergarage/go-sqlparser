// Copyright (C) 2018 The go-sqlparser Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

// Parser represents a parser interface for SQL.
type Parser interface {
	String() string
}
