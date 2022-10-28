// Code generated from ../grammar/JsonTemplate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type JsonTemplateLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var jsontemplatelexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jsontemplatelexerLexerInit() {
	staticData := &jsontemplatelexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'.'", "':'", "'<'", "'<='", "'=='", "'>'", "'>='", "'!='", "'&&'",
		"'||'", "'!'", "'+'", "'-'", "'*'", "'/'", "'('", "')'", "'['", "']'",
		"'#'", "'?'", "'='", "'??'", "'..'", "'as'", "','", "'=>'", "'{'", "'}'",
		"'null'", "'false'", "'true'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "Less", "LessOrEqual", "Equal", "Greater", "GreaterOrEqual",
		"NotEqual", "And", "Or", "Not", "Add", "Subtract", "Multiply", "Divide",
		"LeftParen", "RightParen", "LeftBracket", "RightBracket", "Iteration",
		"Question", "Literal", "NullCoalescing", "Range", "As", "Comma", "Arrow",
		"LeftBrace", "RightBrace", "Null", "False", "True", "ESCAPED_STRING",
		"STRING", "NUMBER", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "Less", "LessOrEqual", "Equal", "Greater", "GreaterOrEqual",
		"NotEqual", "And", "Or", "Not", "Add", "Subtract", "Multiply", "Divide",
		"LeftParen", "RightParen", "LeftBracket", "RightBracket", "Iteration",
		"Question", "Literal", "NullCoalescing", "Range", "As", "Comma", "Arrow",
		"LeftBrace", "RightBrace", "Null", "False", "True", "ESCAPED_STRING",
		"STRING", "NUMBER", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 36, 203, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1,
		24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28,
		1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32,
		5, 32, 162, 8, 32, 10, 32, 12, 32, 165, 9, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 5, 32, 172, 8, 32, 10, 32, 12, 32, 175, 9, 32, 1, 32, 3, 32,
		178, 8, 32, 1, 33, 1, 33, 5, 33, 182, 8, 33, 10, 33, 12, 33, 185, 9, 33,
		1, 34, 4, 34, 188, 8, 34, 11, 34, 12, 34, 189, 1, 34, 1, 34, 4, 34, 194,
		8, 34, 11, 34, 12, 34, 195, 3, 34, 198, 8, 34, 1, 35, 1, 35, 1, 35, 1,
		35, 0, 0, 36, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9,
		19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18,
		37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27,
		55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36,
		1, 0, 6, 2, 0, 34, 34, 92, 92, 2, 0, 39, 39, 92, 92, 4, 0, 36, 36, 65,
		90, 95, 95, 97, 122, 5, 0, 36, 36, 48, 57, 65, 90, 95, 95, 97, 122, 1,
		0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 211, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19,
		1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0,
		27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0,
		0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0,
		0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0,
		0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1,
		0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65,
		1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 1,
		73, 1, 0, 0, 0, 3, 75, 1, 0, 0, 0, 5, 77, 1, 0, 0, 0, 7, 79, 1, 0, 0, 0,
		9, 82, 1, 0, 0, 0, 11, 85, 1, 0, 0, 0, 13, 87, 1, 0, 0, 0, 15, 90, 1, 0,
		0, 0, 17, 93, 1, 0, 0, 0, 19, 96, 1, 0, 0, 0, 21, 99, 1, 0, 0, 0, 23, 101,
		1, 0, 0, 0, 25, 103, 1, 0, 0, 0, 27, 105, 1, 0, 0, 0, 29, 107, 1, 0, 0,
		0, 31, 109, 1, 0, 0, 0, 33, 111, 1, 0, 0, 0, 35, 113, 1, 0, 0, 0, 37, 115,
		1, 0, 0, 0, 39, 117, 1, 0, 0, 0, 41, 119, 1, 0, 0, 0, 43, 121, 1, 0, 0,
		0, 45, 123, 1, 0, 0, 0, 47, 126, 1, 0, 0, 0, 49, 129, 1, 0, 0, 0, 51, 132,
		1, 0, 0, 0, 53, 134, 1, 0, 0, 0, 55, 137, 1, 0, 0, 0, 57, 139, 1, 0, 0,
		0, 59, 141, 1, 0, 0, 0, 61, 146, 1, 0, 0, 0, 63, 152, 1, 0, 0, 0, 65, 177,
		1, 0, 0, 0, 67, 179, 1, 0, 0, 0, 69, 187, 1, 0, 0, 0, 71, 199, 1, 0, 0,
		0, 73, 74, 5, 46, 0, 0, 74, 2, 1, 0, 0, 0, 75, 76, 5, 58, 0, 0, 76, 4,
		1, 0, 0, 0, 77, 78, 5, 60, 0, 0, 78, 6, 1, 0, 0, 0, 79, 80, 5, 60, 0, 0,
		80, 81, 5, 61, 0, 0, 81, 8, 1, 0, 0, 0, 82, 83, 5, 61, 0, 0, 83, 84, 5,
		61, 0, 0, 84, 10, 1, 0, 0, 0, 85, 86, 5, 62, 0, 0, 86, 12, 1, 0, 0, 0,
		87, 88, 5, 62, 0, 0, 88, 89, 5, 61, 0, 0, 89, 14, 1, 0, 0, 0, 90, 91, 5,
		33, 0, 0, 91, 92, 5, 61, 0, 0, 92, 16, 1, 0, 0, 0, 93, 94, 5, 38, 0, 0,
		94, 95, 5, 38, 0, 0, 95, 18, 1, 0, 0, 0, 96, 97, 5, 124, 0, 0, 97, 98,
		5, 124, 0, 0, 98, 20, 1, 0, 0, 0, 99, 100, 5, 33, 0, 0, 100, 22, 1, 0,
		0, 0, 101, 102, 5, 43, 0, 0, 102, 24, 1, 0, 0, 0, 103, 104, 5, 45, 0, 0,
		104, 26, 1, 0, 0, 0, 105, 106, 5, 42, 0, 0, 106, 28, 1, 0, 0, 0, 107, 108,
		5, 47, 0, 0, 108, 30, 1, 0, 0, 0, 109, 110, 5, 40, 0, 0, 110, 32, 1, 0,
		0, 0, 111, 112, 5, 41, 0, 0, 112, 34, 1, 0, 0, 0, 113, 114, 5, 91, 0, 0,
		114, 36, 1, 0, 0, 0, 115, 116, 5, 93, 0, 0, 116, 38, 1, 0, 0, 0, 117, 118,
		5, 35, 0, 0, 118, 40, 1, 0, 0, 0, 119, 120, 5, 63, 0, 0, 120, 42, 1, 0,
		0, 0, 121, 122, 5, 61, 0, 0, 122, 44, 1, 0, 0, 0, 123, 124, 5, 63, 0, 0,
		124, 125, 5, 63, 0, 0, 125, 46, 1, 0, 0, 0, 126, 127, 5, 46, 0, 0, 127,
		128, 5, 46, 0, 0, 128, 48, 1, 0, 0, 0, 129, 130, 5, 97, 0, 0, 130, 131,
		5, 115, 0, 0, 131, 50, 1, 0, 0, 0, 132, 133, 5, 44, 0, 0, 133, 52, 1, 0,
		0, 0, 134, 135, 5, 61, 0, 0, 135, 136, 5, 62, 0, 0, 136, 54, 1, 0, 0, 0,
		137, 138, 5, 123, 0, 0, 138, 56, 1, 0, 0, 0, 139, 140, 5, 125, 0, 0, 140,
		58, 1, 0, 0, 0, 141, 142, 5, 110, 0, 0, 142, 143, 5, 117, 0, 0, 143, 144,
		5, 108, 0, 0, 144, 145, 5, 108, 0, 0, 145, 60, 1, 0, 0, 0, 146, 147, 5,
		102, 0, 0, 147, 148, 5, 97, 0, 0, 148, 149, 5, 108, 0, 0, 149, 150, 5,
		115, 0, 0, 150, 151, 5, 101, 0, 0, 151, 62, 1, 0, 0, 0, 152, 153, 5, 116,
		0, 0, 153, 154, 5, 114, 0, 0, 154, 155, 5, 117, 0, 0, 155, 156, 5, 101,
		0, 0, 156, 64, 1, 0, 0, 0, 157, 163, 5, 34, 0, 0, 158, 159, 5, 92, 0, 0,
		159, 162, 9, 0, 0, 0, 160, 162, 8, 0, 0, 0, 161, 158, 1, 0, 0, 0, 161,
		160, 1, 0, 0, 0, 162, 165, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164,
		1, 0, 0, 0, 164, 166, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 166, 178, 5, 34,
		0, 0, 167, 173, 5, 39, 0, 0, 168, 169, 5, 92, 0, 0, 169, 172, 9, 0, 0,
		0, 170, 172, 8, 1, 0, 0, 171, 168, 1, 0, 0, 0, 171, 170, 1, 0, 0, 0, 172,
		175, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 176,
		1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 176, 178, 5, 39, 0, 0, 177, 157, 1, 0,
		0, 0, 177, 167, 1, 0, 0, 0, 178, 66, 1, 0, 0, 0, 179, 183, 7, 2, 0, 0,
		180, 182, 7, 3, 0, 0, 181, 180, 1, 0, 0, 0, 182, 185, 1, 0, 0, 0, 183,
		181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 68, 1, 0, 0, 0, 185, 183, 1,
		0, 0, 0, 186, 188, 7, 4, 0, 0, 187, 186, 1, 0, 0, 0, 188, 189, 1, 0, 0,
		0, 189, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 197, 1, 0, 0, 0, 191,
		193, 5, 46, 0, 0, 192, 194, 7, 4, 0, 0, 193, 192, 1, 0, 0, 0, 194, 195,
		1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 198, 1, 0,
		0, 0, 197, 191, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 70, 1, 0, 0, 0,
		199, 200, 7, 5, 0, 0, 200, 201, 1, 0, 0, 0, 201, 202, 6, 35, 0, 0, 202,
		72, 1, 0, 0, 0, 10, 0, 161, 163, 171, 173, 177, 183, 189, 195, 197, 1,
		0, 1, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// JsonTemplateLexerInit initializes any static state used to implement JsonTemplateLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewJsonTemplateLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func JsonTemplateLexerInit() {
	staticData := &jsontemplatelexerLexerStaticData
	staticData.once.Do(jsontemplatelexerLexerInit)
}

// NewJsonTemplateLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewJsonTemplateLexer(input antlr.CharStream) *JsonTemplateLexer {
	JsonTemplateLexerInit()
	l := new(JsonTemplateLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &jsontemplatelexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "JsonTemplate.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JsonTemplateLexer tokens.
const (
	JsonTemplateLexerT__0           = 1
	JsonTemplateLexerT__1           = 2
	JsonTemplateLexerLess           = 3
	JsonTemplateLexerLessOrEqual    = 4
	JsonTemplateLexerEqual          = 5
	JsonTemplateLexerGreater        = 6
	JsonTemplateLexerGreaterOrEqual = 7
	JsonTemplateLexerNotEqual       = 8
	JsonTemplateLexerAnd            = 9
	JsonTemplateLexerOr             = 10
	JsonTemplateLexerNot            = 11
	JsonTemplateLexerAdd            = 12
	JsonTemplateLexerSubtract       = 13
	JsonTemplateLexerMultiply       = 14
	JsonTemplateLexerDivide         = 15
	JsonTemplateLexerLeftParen      = 16
	JsonTemplateLexerRightParen     = 17
	JsonTemplateLexerLeftBracket    = 18
	JsonTemplateLexerRightBracket   = 19
	JsonTemplateLexerIteration      = 20
	JsonTemplateLexerQuestion       = 21
	JsonTemplateLexerLiteral        = 22
	JsonTemplateLexerNullCoalescing = 23
	JsonTemplateLexerRange          = 24
	JsonTemplateLexerAs             = 25
	JsonTemplateLexerComma          = 26
	JsonTemplateLexerArrow          = 27
	JsonTemplateLexerLeftBrace      = 28
	JsonTemplateLexerRightBrace     = 29
	JsonTemplateLexerNull           = 30
	JsonTemplateLexerFalse          = 31
	JsonTemplateLexerTrue           = 32
	JsonTemplateLexerESCAPED_STRING = 33
	JsonTemplateLexerSTRING         = 34
	JsonTemplateLexerNUMBER         = 35
	JsonTemplateLexerWS             = 36
)
