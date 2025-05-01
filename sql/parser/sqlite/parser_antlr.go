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
	inputStream   *go_antlr.InputStream
	lexer         *antlr.SQLiteLexer
	tokenStream   *go_antlr.CommonTokenStream
	parser        *antlr.SQLiteParser
	stmts         query.StatementList
	errorListener *antlrParserErrorListener
}

// NewParser returns a new parser.
func NewParser() *Parser {
	parser := &Parser{
		inputStream:   nil,
		lexer:         nil,
		tokenStream:   nil,
		parser:        nil,
		stmts:         nil,
		errorListener: newANTLRParserErrorListener(),
	}
	return parser
}

// ParseString parses a specified FQL string.
func (parser *Parser) ParseString(queryString string) ([]query.Statement, error) {
	queryStringLen := len(queryString)
	switch queryStringLen {
	case 0:
		return nil, errors.ErrEmptyQuery
	case 1:
		if queryString[0] == ';' {
			return nil, errors.ErrEmptyQuery
		}
	}

	if len(queryString) == 0 {
		return nil, errors.ErrEmptyQuery
	}

	parser.inputStream = go_antlr.NewInputStream(queryString)
	parser.lexer = antlr.NewSQLiteLexer(parser.inputStream)
	parser.tokenStream = go_antlr.NewCommonTokenStream(parser.lexer, 0)
	parser.parser = antlr.NewSQLiteParser(parser.tokenStream)
	parser.parser.AddErrorListener(parser.errorListener)
	parser.parser.BuildParseTrees = true
	tree := parser.parser.Parse()

	if !parser.errorListener.IsSuccess() {
		return nil, fmt.Errorf("%s (%w)", queryString, parser.errorListener.GetError())
	}

	v := tree.Accept(newANTLRVisitor())
	stmtList, ok := v.(query.StatementList)
	if !ok {
		return nil, fmt.Errorf("failed to parse query: %s", queryString)
	}
	parser.stmts = stmtList

	return parser.stmts, nil
}
