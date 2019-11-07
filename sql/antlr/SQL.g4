/******************************************************************
* Copyright (C) 2019 The go-sqlparser Authors. All rights reserved.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
******************************************************************/

grammar SQL;

/******************************************************************
 *
 * PARSER RULES
 *
******************************************************************/

queries
	: statement (SEMICOLON statement)* ';'*
	;	

statement
	: showStmt
	| useStmt
	| createStmt
	| create_indexStmt
	| drop_collectionStmt
	| drop_indexStmt
	| selectStmt
	| insertStmt
	| updateStmt
	| deleteStmt
	;

/******************************************************************
*
* SHOW
*
******************************************************************/

showStmt
	: SHOW collection_section
	;

/******************************************************************
*
* USE
*
******************************************************************/

useStmt
	: USE collection_section
	;
	
/******************************************************************
*
* CREATE
*
******************************************************************/

createStmt
	: createDatabase
	| createCollection
	;

createDatabase
	: CREATE DATABASE database
	;

createCollection
	: CREATE COLLECTION collection_section (OPTIONS expression)?
	;

/******************************************************************
*
* DROP
*
******************************************************************/

drop_collectionStmt
	: DROP COLLECTION collection_section
	;

/******************************************************************
*
* CREATE INDEX
*
******************************************************************/

create_indexStmt
	: CREATE COLLECTION_INDEX index_section
	;

/******************************************************************
*
* DROP INDEX
*
******************************************************************/

drop_indexStmt
	: DROP COLLECTION_INDEX index_section
	;

/******************************************************************
*
* SELECT
*
******************************************************************/

selectStmt
	: SELECT (DISTINCT | ALL)? columns from_section (where_section)? (grouping_section)? (having_section)? (sorting_section)? (limit_section)? (offset_section)? 
	;

columns
	: ASTERISK
	| column (',' column)*
	;

from_section
	: (FROM table_name) (COMMA table_name)*
	;

table_name
	: data_source
	;

data_source
	: collection_name
	;

grouping_section
	: GROUP BY expression (COMMA expression)*
	;
	
having_section
	: HAVING expression
	;

sorting_section
	: ORDER BY sorting_item (COMMA sorting_item)*
	;
		
sorting_item
	: property (sorting_specification)?
	;

sorting_specification
	: ASC
	| DESC
	;

limit_section
	: LIMIT (expression_literal COMMA)? expression_literal
	;

offset_section
	: OFFSET
	;

/******************************************************************
*
* INSERT
*
******************************************************************/

insertStmt
	: (sync_operator)? INSERT INTO collection_section (insert_columns_section)? insert_values_section
	;

insert_columns_section
	: '(' column (',' column)* ')'
	;

insert_values_section
	: VALUE expression
	| VALUES expression
	;

/******************************************************************
*
* UPDATE
*
******************************************************************/

updateStmt
	: (sync_operator)? UPDATE collection_section SET property_section (COMMA property_section)* (where_section)?
	;

property_section
	: property SINGLE_EQ expression_literal
	;

/******************************************************************
*
* DELETE
*
******************************************************************/

deleteStmt
	: (sync_operator)? DELETE FROM collection_section (where_section)?
	;

/******************************************************************
*
* Expression
*
******************************************************************/

expression
 	: expression_list
	;

expression_list
 	: expression_literal
	| expression_function
	| expression_binary_operator (expression_logic_operator expression_binary_operator)*
	| '(' expression_array (COMMA expression_array)* ')'
	| '{' (expression_dictionary) (COMMA expression_dictionary)* '}'
	| '[' expression_array (COMMA expression_array)* ']'
	;

expression_literal
 	: expression_literal_value
 	;

expression_literal_value
 	: property_literal
	| integer_literal
	| real_literal
	| string_literal
	| true_literal
	| false_literal
	| NIL
	| CURRENT_TIME
	| CURRENT_DATE
	| CURRENT_TIMESTAMP
	;

expression_dictionary
	: name ':' expression_literal
	;

dictionary_literal
	: name ':' expression_literal
	;

expression_array
	: expression_literal
	;

array_literal
	: expression_literal
	;

expression_logic_operator
	: logical_operator
	;

expression_binary_operator
	: expression_operator
	;

expression_function
	: IDENTIFIER '(' (function_value)? ')' 
	;

function_name
	: IDENTIFIER
	;

function_value
	: expression (COMMA expression)*
	| ASTERISK
	;

expression_operator
	: expression_literal binary_operator expression_literal
	;

binary_operator
	: SINGLE_EQ
	| DOUBLE_EQ
	| OP_LT
	| LE
	| GT
	| GE
	| NOTEQ
	;

logical_operator
	: AND
	| OR
	;

property_literal
	: IDENTIFIER 
	;

integer_literal
	: NUMBER
	;

real_literal
	: FLOAT
	;

string_literal
	: STRING
	;

true_literal
	: TRUE
	;
	
false_literal
	: FALSE
	;

/******************************************************************
*
* COMMON
*
******************************************************************/

sync_operator
	: SYNC
	| ASYNC
	;

compound_operator
	: UNION (ALL)?
	| INTERSECT
	| EXCEPT
	;

condition_operator
	: SINGLE_EQ
	| DOUBLE_EQ
	| OP_LT
	| LE
	| GT
	| GE
	| NOTEQ
	;

property
	: IDENTIFIER 
	;

value 	
	: IDENTIFIER	
	;

name
	: IDENTIFIER
	;

collection_section
	: collection_name
	;

collection_name
	: IDENTIFIER
	;

database
	: IDENTIFIER
	;

column
	: ((expression) (AS name)?)
	;

index_section
	: index_name
	;

index_name
	: IDENTIFIER
	| string_literal
	;

where_section
	: WHERE expression
	;

/*------------------------------------------------------------------
 * LEXER RULES
 *------------------------------------------------------------------*/


ASTERISK
	: '*'
	;
	
SINGLE_EQ
	: '='
	;

DOUBLE_EQ
	: '=='
	;

OP_LT
	: '<'
	;
	
LE
	: '<='
	;
	
GT
	: '>'
	;

GE	
	: '>='
	;
	
NOTEQ
	: '!='
	;

AND
	: A N D
	| '&'
	;

OR
	: O R
	| '|'
	;

COMMA
	: ','
	;

SEMICOLON
	: ';'
	;
	
fragment 
A	: 'A'
	| 'a'
	;

fragment 
B	: 'B'
	| 'b'
	;

fragment 
C	: 'C'
	| 'c'
	;

fragment 
D	: 'D'
	| 'd'
	;

fragment 
E	: 'E'
	| 'e'
	;

fragment 
F	: 'F'
	| 'f'
	;

fragment 
G	: 'G'
	| 'g'
	;

fragment 
H	: 'H'
	| 'h'
	;

fragment 
I	: 'I'
	| 'i'
	;

fragment 
J	: 'J'
	| 'j'
	;

fragment 
K	: 'K'
	| 'k'
	;

fragment 
L	: 'L'
	| 'l'
	;

fragment 
M	: 'M'
	| 'm'
	;

fragment 
N	: 'N'
	| 'n'
	;

fragment 
O	: 'O'
	| 'o'
	;

fragment 
P	: 'P'
	| 'p'
	;

fragment 
Q	: 'Q'
	| 'q'
	;

fragment 
R	: 'R'
	| 'r'
	;

fragment 
S	: 'S'
	| 's'
	;

fragment 
T	: 'T'
	| 't'
	;

fragment 
U	: 'U'
	| 'u'
	;

fragment 
V	: 'V'
	| 'v'
	;

fragment 
W	: 'W'
	| 'w'
	;

fragment 
X	: 'X'
	| 'x'
	;

fragment 
Y	: 'Y'
	| 'y'
	;

fragment 
Z	: 'Z'
	| 'z'
	;

ALL
	: A L L
	;
	
ANCESTOR
	: A N C E S T O R
	;

AS
	: A S
	;

ASC
	: A S C
	;
	
ASYNC
	: A S Y N C
	;
	
BY
	: B Y
	;

CREATE
	: C R E A T E
	;

COLLECTION
	: C O L L E C T I O N
	;

CURRENT_DATE
	: C U R R E N T '_' D A T E
	;

CURRENT_TIME
	: C U R R E N T '_' T I M E
	;

CURRENT_TIMESTAMP
	: C U R R E N T '_' T I M E S T A M P
	;

DATABASE
	: D A T A B A S E
	;

DESC
	: D E S C
	;

DELETE
	: D E L E T E
	;

DISTINCT
	: D I S T I N C T
	;

DROP
	: D R O P
	;

EACH
	: E A C H
	;
		
EXCEPT
	: E X C E P T
	;

FLATTEN
	: F L A T T E N
	;
		
FROM
	: F R O M
	;

GROUP
	: G R O U P
	;
		
HAVING
	: H A V I N G
	;
	
IN
	: I N
	;

COLLECTION_INDEX
	: I N D E X
	;

INSERT
	: I N S E R T
	;

INTERSECT
	: I N T E R S E C T
	;

INTO
	: I N T O
	;

IS
	: I S
	;

LIMIT
	: L I M I T
	;

NIL
	: N U L L
	;

OFFSET
	: O F F S E T
	;

OPTIONS
	: O P T I O N S
	;

ORDER
	: O R D E R
	;
	
SELECT
	: S E L E C T
	;

SET
	: S E T
	;

SHOW
	: S H O W
	;

SYNC
	: S Y N C
	;

USE
	: U S E
	;

UNION
	: U N I O N
	;

UPDATE
	: U P D A T E
	;

WHERE
	: W H E R E
	;

VALUE
	: V A L U E
	;

VALUES
	: V A L U E S
	;

TRUE
	: T R U E
	;
	
FALSE
	: F A L S E
	;

/******************************************************************
*
* COMMON
*
******************************************************************/

WS  :   ( ' '
        | '\t'
        | '\r'
        | '\n'
        ) -> skip
    ;
	
IDENTIFIER  
	: [a-zA-Z_] [a-zA-Z_0-9]* 
    ;

NUMBER 
	: '0'..'9'+
	;

FLOAT
	:   ('0'..'9')+ '.' ('0'..'9')* EXPONENT?
	|   '.' ('0'..'9')+ EXPONENT?
	|   ('0'..'9')+ EXPONENT
	;

STRING
	: '\'' ( ~'\'' | '\'\'' )* '\''
	;

fragment
EXPONENT : ('e'|'E') ('+'|'-')? ('0'..'9')+ ;

fragment
HEX_DIGIT : ('0'..'9'|'a'..'f'|'A'..'F') ;
