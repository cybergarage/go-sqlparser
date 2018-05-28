// Copyright (C) 2018 The go-sqlparser Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

// Query represents a query interface.
type Query interface {
	String() string
}
