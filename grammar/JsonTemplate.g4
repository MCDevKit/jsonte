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
Spread: '...';
As: 'as';
Comma: ',';
Arrow: '=>';
Colon: ':';
Semicolon: ';';
Dot: '.';

LeftBrace: '{';
RightBrace: '}';

Null: 'null';
False: 'false';
True: 'true';

Return: 'return';
For: 'for';
In: 'in';
If: 'if';
Else: 'else';
While: 'while';
Do: 'do';
Break: 'break';
Continue: 'continue';

script
    : statement*
    ;

statement
    : statements
    | field Semicolon
    | Return field? Semicolon
    | Break Semicolon
    | Continue Semicolon
    | For (name (Comma name)?) In field statements
    | If field statements (Else If field statements)* (Else statements)?
    | While field statements
    | Do statements While field Semicolon
    ;

statements
    : LeftBrace statement* RightBrace
    ;

expression
    : (LeftBrace LeftBrace)? Iteration field (As name (Comma name)?)? (RightBrace RightBrace)? EOF
    | (LeftBrace LeftBrace)? Question field (RightBrace RightBrace)? EOF
    | (LeftBrace LeftBrace)? Literal field (RightBrace RightBrace)? EOF
    | (LeftBrace LeftBrace)? field (RightBrace RightBrace)? EOF
    ;

lambda
    : name Arrow field
    | LeftParen (name (Comma name)*)* RightParen Arrow field
    ;

function_param
    : (Spread)? field
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
    | field (Question? Dot name)
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
    | field Question field (Colon field)?
    | field Literal field
    ;

array
    : LeftBracket (spread_field (Comma spread_field)*)? RightBracket
    ;

spread_field
    : Spread? field
    ;

object
    : LeftBrace (object_field (Comma object_field)*)? RightBrace
    ;

object_field
    : name Colon field
    | ESCAPED_STRING Colon field
    | Spread? field
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