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

package sqlite

import (
	"fmt"

	go_antlr "github.com/antlr4-go/antlr/v4"
	errors "github.com/cybergarage/go-sqlparser/sql/parser"
	antlr "github.com/cybergarage/go-sqlparser/sql/parser/sqlite/antlr"
	"github.com/cybergarage/go-sqlparser/sql/query"
)

// Parser represents a FQL parser based on ANTLR.
type Parser struct {
}

// NewParser returns a new parser.
func NewParser() *Parser {
	parser := &Parser{}
	return parser
}

// ParseString parses a specified FQL string.
func (parser *Parser) ParseString(queryString string) ([]query.Statement, error) {
	if len(queryString) <= 0 {
		return nil, errors.ErrEmptyQuery
	}

	input := go_antlr.NewInputStream(queryString)
	lexer := antlr.NewSQLiteLexer(input)
	stream := go_antlr.NewCommonTokenStream(lexer, 0)
	p := antlr.NewSQLiteParser(stream)
	el := newANTLRParserErrorListener()
	p.AddErrorListener(el)
	p.BuildParseTrees = true
	tree := p.Parse()

	v := newANTLRParserVisitor()
	tree.Accept(v)
	if !el.IsSuccess() {
		return nil, fmt.Errorf("%s (%s)", queryString, el.GetError().Error())
	}

	pl := newANTLRParserListener()
	go_antlr.ParseTreeWalkerDefault.Walk(pl, tree)

	if !el.IsSuccess() {
		return nil, fmt.Errorf("%s (%s)", queryString, el.GetError().Error())
	}

	return nil, nil
}
