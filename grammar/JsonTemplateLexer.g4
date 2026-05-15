lexer grammar JsonTemplateLexer;

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
AddAssign: '+=';
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

Backtick: '`' -> pushMode(TEMPLATE_MODE);

ESCAPED_STRING : ('"' ('\\' . | ~["\\])* '"') | ('\'' ('\\' . | ~['\\])* '\'');

STRING
    : [a-zA-Z_$][a-zA-Z0-9_$]*
    ;

NUMBER
    : [0-9]+('.'[0-9]+)?
    ;

LINE_COMMENT
    : '//' ~[\r\n]* -> channel(HIDDEN)
    ;

BLOCK_COMMENT
    : '/*' .*? '*/' -> channel(HIDDEN)
    ;

WS
    : [ \r\n\t] -> channel(HIDDEN)
    ;

// Template string mode — entered when a backtick is encountered
mode TEMPLATE_MODE;

// Literal text: any char except backtick, backslash, dollar; escape sequences; dollar not followed by {
TEMPLATE_TEXT
    : (~[`$\\] | '\\' . | '$' ~[`{])+
    ;

// Start of an interpolation — mode switch is handled by the TemplateAwareLexer wrapper
TEMPLATE_INTERP_START
    : '${'
    ;

// Closing backtick — returns to previous mode
TEMPLATE_END
    : '`' -> popMode
    ;
