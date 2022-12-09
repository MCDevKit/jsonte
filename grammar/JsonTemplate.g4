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

LeftBrace: '{';
RightBrace: '}';

Null: 'null';
False: 'false';
True: 'true';

expression
    : (LeftBrace LeftBrace)? Iteration field (As name (Comma name)?)? (RightBrace RightBrace)?
    | (LeftBrace LeftBrace)? Question field (RightBrace RightBrace)?
    | (LeftBrace LeftBrace)? Literal field (RightBrace RightBrace)?
    | (LeftBrace LeftBrace)? field (RightBrace RightBrace)?
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
   | object
   | name
   | field (Question? '.' name)
   | field (Question? LeftBracket index RightBracket)
   | field LeftParen (function_param (Comma function_param)*)? RightParen
   | Subtract field
   | Not field
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
   | field Question field (':' field)?
   ;

array
    : LeftBracket (field (Comma field)*)? RightBracket
    ;

object
    : '{' (object_field (Comma object_field)*)? '}'
    ;

object_field
    : name ':' field
    | ESCAPED_STRING ':' field
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
   : [a-zA-Z_$][a-zA-Z0-9_$]*
   ;

NUMBER
   : [0-9]+('.'[0-9]+)?
   ;

WS
   : [ \r\n\t] -> channel(HIDDEN)
   ;