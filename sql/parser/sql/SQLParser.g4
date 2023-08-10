select_statement
    : SELECT select_list FROM table_reference
    ;

select_list
    : select_list_item (',' select_list_item)*
    ;

select_list_item
    : column_reference
    | function_call
    ;

table_reference
    : table_name
    | join_clause
    ;

join_clause
    : table_reference JOIN table_reference ON join_condition
    ;

join_condition
    : boolean_expression
    ;

boolean_expression
    : comparison_expression
    | boolean_expression AND boolean_expression
    | boolean_expression OR boolean_expression
    ;

comparison_expression
    : arithmetic_expression '=' arithmetic_expression
    | arithmetic_expression '<>' arithmetic_expression
    | arithmetic_expression '<' arithmetic_expression
    | arithmetic_expression '>' arithmetic_expression
    | arithmetic_expression '<=' arithmetic_expression
    | arithmetic_expression '>=' arithmetic_expression
    ;

arithmetic_expression
    : term (PLUS term | MINUS term)*
    ;

term
    : factor (MULT factor | DIV factor)*
    ;

factor
    : column_reference
    | function_call
    | literal
    | '(' arithmetic_expression ')'
    ;

column_reference
    : table_name '.' column_name
    | column_name
    ;

function_call
    : function_name '(' function_arguments ')'
    ;

function_name
    : 'COUNT'
    | 'SUM'
    | 'AVG'
    | 'MAX'
    | 'MIN'
    ;

function_arguments
    : arithmetic_expression (',' arithmetic_expression)*
    ;

literal
    : INTEGER
    | FLOAT
    | STRING
    ;

table_name
    : IDENTIFIER
    ;

column_name
    : IDENTIFIER
    ;

IDENTIFIER
    : [a-zA-Z][a-zA-Z0-9_]*
    ;

INTEGER
    : [0-9]+
    ;

FLOAT
    : [0-9]+ '.' [0-9]+
    ;

STRING
    : '\'' (~'\'' | '\'\'' | '\\\'')* '\''
    ;

PLUS
    : '+'
    ;

MINUS
    : '-'
    ;

MULT
    : '*'
    ;

DIV
    : '/'
    ;

AND
    : 'AND'
    ;

OR
    : 'OR'
    ;

SELECT
    : 'SELECT'
    ;

FROM
    : 'FROM'
    ;

JOIN
    : 'JOIN'
    ;

ON
    : 'ON'
    ;

WS
    : [ \t\r\n]+ -> skip
    ;
