// Copyright (C) 2018 The go-sqlparser Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"fmt"

	antlr "github.com/antlr/antlr4/runtime/Go/antlr"
)

type antlrParserErrorListener struct {
	antlr.ErrorListener
	done bool
	msg  string
}

func newANTLRParserErrorListener() *antlrParserErrorListener {
	l := &antlrParserErrorListener{
		done: true,
		msg:  "",
	}
	return l
}

func (l *antlrParserErrorListener) IsSuccess() bool {
	return l.done
}

func (l *antlrParserErrorListener) GetError() error {
	return fmt.Errorf("%s", l.msg)
}

func (l *antlrParserErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.done = false
	l.msg = msg
}

func (l *antlrParserErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {

}

func (l *antlrParserErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {

}

func (l *antlrParserErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {

}
