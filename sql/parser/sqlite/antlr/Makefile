# Copyright (C} 2019  The go-sqlparser Authors. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"};
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#	http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

SHELL := bash

.PHONY: ${SQL_LEXER_G4} ${SQL_GRAMMER_G4}

GIT_RAW_URL=https://raw.githubusercontent.com/antlr/grammars-v4/master/sql/sqlite/
PKG_NAME=go-sql
SQL_GRAMMER_G4 = SQLiteParser.g4
SQL_LEXER_G4 = SQLiteLexer.g4
SQL_GRAMMER_FILES = ${SQL_LEXER_G4} ${SQL_GRAMMER_G4}
ANTLR=antlr

all: antlr

${SQL_LEXER_G4}:
	wget -N ${GIT_RAW_URL}${SQL_LEXER_G4}

${SQL_GRAMMER_G4}:
	wget -N ${GIT_RAW_URL}${SQL_GRAMMER_G4}
	patch ${SQL_GRAMMER_G4} < ${SQL_GRAMMER_G4}.patch

antlr:${SQL_LEXER_G4} ${SQL_GRAMMER_G4}
	@${ANTLR} -package antlr -Dlanguage=Go -visitor ${SQL_GRAMMER_FILES}