// Copyright (C) 2018 The go-sqlparser Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"fmt"

	go_antlr "github.com/antlr/antlr4/runtime/Go/antlr"

	antlr "github.com/cybergarage/usql-go/usql/parser/antlr"
)

// antlrParser represents a FQL parser based on ANTLR.
type antlrParser struct {
	Parser
}

// NewParser returns a new parse.
func NewParser() Parser {
	parser := &antlrParser{}
	return parser
}

// ParseString parses a specified FQL string.
func (parser *antlrParser) ParseString(queryString string) ([]Query, error) {
	if len(queryString) <= 0 {
		return nil, fmt.Errorf(errorEmptyQuery)
	}

	input := go_antlr.NewInputStream(queryString)
	lexer := antlr.NewSQLLexer(input)
	stream := go_antlr.NewCommonTokenStream(lexer, 0)
	p := antlr.NewSQLParser(stream)
	el := newANTLRParserErrorListener()
	p.AddErrorListener(el)
	p.BuildParseTrees = true
	tree := p.Queries()
	pl := newANTLRParserListener()
	go_antlr.ParseTreeWalkerDefault.Walk(pl, tree)
	if !el.IsSuccess() {
		return nil, fmt.Errorf("%s (%s)", queryString, el.GetError().Error())
	}

	return nil, nil
}
