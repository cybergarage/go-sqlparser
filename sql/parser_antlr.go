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
	"fmt"

	go_antlr "github.com/antlr/antlr4/runtime/Go/antlr"

	antlr "github.com/cybergarage/go-sql/sql/antlr"
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
