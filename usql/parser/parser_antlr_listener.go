// Copyright (C) 2018 The go-sqlparser Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	antlr "github.com/cybergarage/usql-go/usql/parser/antlr"
)

type antlrParserListener struct {
	*antlr.BaseSQLListener
	*ParserStack
}

func newANTLRParserListener() *antlrParserListener {
	l := &antlrParserListener{
		BaseSQLListener: &antlr.BaseSQLListener{},
		ParserStack:     NewParserStack(),
	}
	return l
}
