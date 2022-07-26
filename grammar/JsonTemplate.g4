grammar JsonTemplate;

Less: '<';
LessOrEqual: '<=';
Equal: '==';
Greater: '>';
GreaterOrEqual: '>=';
NotEqual: '!=';
And: '&&';
Or: '||';
Not: '!';

//Math operators
Add: '+';
Subtract: '-';
Multiply: '*';
Divide: '/';

LeftParen: '(';
RightParen: ')';
LeftBracket: '[';
RightBracket: ']';

//Actions
Iteration: '#';
Question: '?';
Literal: '=';

NullCoalescing: '??';

Range: '..';
As: 'as';
Comma: ',';
Arrow: '=>';

LeftMustache: '{{';
RightMustache: '}}';

Null: 'null';
False: 'false';
True: 'true';

expression
    : LeftMustache? Iteration field (As name)? RightMustache?
    | LeftMustache? Question field RightMustache?
    | LeftMustache? Literal field RightMustache?
    | LeftMustache? field RightMustache?
    ;

lambda
    : name Arrow field
    | LeftParen (name (Comma name)*)* RightParen Arrow field
    ;

function_param
    : field
    | lambda
    ;

field
   : LeftParen field RightParen
   | True
   | False
   | Null
   | NUMBER
   | ESCAPED_STRING
   | array
   | name
   | field (Question? '.' name)
   | field (Question? LeftBracket index RightBracket)
   | field LeftParen (function_param (Comma function_param)*)? RightParen
   | Subtract field
   | field (Divide | Multiply) field
   | field (Add | Subtract) field
   | field Range field
   | field NullCoalescing field
   | field Equal field
   | field Less field
   | field LessOrEqual field
   | field Greater field
   | field GreaterOrEqual field
   | field NotEqual field
   | field And field
   | field Or field
   | Not field
   | field Question field (':' field)?
   ;

array
    : LeftBracket (field (Comma field)*)? RightBracket
    ;

name
   : STRING
   ;

index
   : field
   | NUMBER
   | ESCAPED_STRING
   ;

ESCAPED_STRING : ('"' ('\\' . | ~["\\])* '"') | ('\'' ('\\' . | ~['\\])* '\'');

STRING
   : [a-zA-Z_][a-zA-Z0-9_]*
   ;

NUMBER
   : [0-9]+('.'[0-9]+)?
   ;

WS
   : [ \r\n\t] -> channel(HIDDEN)
   ;