parser grammar JsonTemplate;

options { tokenVocab=JsonTemplateLexer; }

script
    : statement*
    ;

statement
    : statements
    | field Semicolon
    | Return field? Semicolon
    | Break Semicolon
    | Continue Semicolon
    | For (name (Comma name)?) In field statement
    | If field statement (Else If field statement)* (Else statement)?
    | While field statement
    | Do statement While field Semicolon
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
    | name Arrow statements
    | LeftParen (name (Comma name)*)* RightParen Arrow field
    | LeftParen (name (Comma name)*)* RightParen Arrow statements
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
    | template_string
    | array
    | object
    | lambda
    | name
    | field (Question? Dot name)
    | field (Question? LeftBracket index RightBracket)
    | field LeftParen (function_param (Comma function_param)*)? RightParen
    | field Increment
    | field Decrement
    | Subtract field
    | Add field
    | Not field
    | field (Divide | Multiply) field
    | field (Add | Subtract) field
    | field Range field
    | field (Less | LessOrEqual | Greater | GreaterOrEqual | In) field
    | field (Equal | NotEqual) field
    | field And field
    | field Or field
    | field NullCoalescing field
    | field Question field (Colon field)?
    | field (AddAssign | SubtractAssign | MultiplyAssign | DivideAssign | ModuloAssign) field
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
    | LeftBracket field RightBracket Colon field
    | Spread? field
    ;

template_string
    : Backtick template_string_part* TEMPLATE_END
    ;

template_string_part
    : TEMPLATE_TEXT
    | TEMPLATE_INTERP_START field RightBrace
    ;

name
    : STRING
    ;

index
    : field
    | NUMBER
    | ESCAPED_STRING
    ;
