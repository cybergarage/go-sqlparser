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
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type antlrParserErrorListener struct {
	*antlr.DiagnosticErrorListener
	done bool
	msg  string
}

func newANTLRParserErrorListener() *antlrParserErrorListener {
	l := &antlrParserErrorListener{
		DiagnosticErrorListener: antlr.NewDiagnosticErrorListener(true),
		done:                    true,
		msg:                     "",
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
	l.msg = "line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg
	l.done = false
}

func (l *antlrParserErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	// l.DiagnosticErrorListener.ReportAmbiguity(recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
}

func (l *antlrParserErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	// l.DiagnosticErrorListener.ReportAttemptingFullContext(recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
}

func (l *antlrParserErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	// l.DiagnosticErrorListener.ReportContextSensitivity(recognizer, dfa, startIndex, stopIndex, prediction, configs)
}
