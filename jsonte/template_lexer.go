package jsonte

import (
	"github.com/MCDevKit/jsonte/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// interpState tracks brace depth for one active ${...} interpolation.
type interpState struct {
	depth int
}

// TemplateAwareLexer wraps the generated lexer and manages the mode stack for
// template string interpolations. ANTLR combined/split grammars do not support
// tracking brace depth in lexer modes, so we handle it here:
//
//   - When TEMPLATE_INTERP_START ("${") is emitted, push DEFAULT_MODE so that
//     expression tokens are lexed correctly.
//   - Track LeftBrace/RightBrace depth so that object literals inside ${...}
//     do not prematurely close the interpolation.
//   - When the matching RightBrace is found (depth == 0), pop DEFAULT_MODE back
//     to TEMPLATE_MODE so subsequent characters are treated as template text.
type TemplateAwareLexer struct {
	*parser.JsonTemplateLexer
	interpStack []interpState
}

func NewTemplateAwareLexer(input antlr.CharStream) *TemplateAwareLexer {
	return &TemplateAwareLexer{
		JsonTemplateLexer: parser.NewJsonTemplateLexer(input),
	}
}

func (l *TemplateAwareLexer) NextToken() antlr.Token {
	t := l.JsonTemplateLexer.NextToken()
	switch t.GetTokenType() {
	case parser.JsonTemplateLexerTEMPLATE_INTERP_START:
		l.interpStack = append(l.interpStack, interpState{depth: 0})
		l.JsonTemplateLexer.PushMode(antlr.LexerDefaultMode)

	case parser.JsonTemplateLexerLeftBrace:
		if len(l.interpStack) > 0 {
			l.interpStack[len(l.interpStack)-1].depth++
		}

	case parser.JsonTemplateLexerRightBrace:
		if len(l.interpStack) > 0 {
			top := &l.interpStack[len(l.interpStack)-1]
			if top.depth == 0 {
				l.interpStack = l.interpStack[:len(l.interpStack)-1]
				l.JsonTemplateLexer.PopMode()
			} else {
				top.depth--
			}
		}
	}
	return t
}
