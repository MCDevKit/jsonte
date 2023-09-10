// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
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
		"", "';'", "'return'", "'for'", "'in'", "'if'", "'else'", "'while'",
		"'do'", "'.'", "':'", "'<'", "'<='", "'=='", "'>'", "'>='", "'!='",
		"'&&'", "'||'", "'!'", "'+'", "'-'", "'*'", "'/'", "'('", "')'", "'['",
		"']'", "'#'", "'?'", "'='", "'??'", "'..'", "'as'", "','", "'=>'", "'{'",
		"'}'", "'null'", "'false'", "'true'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "Less", "LessOrEqual", "Equal",
		"Greater", "GreaterOrEqual", "NotEqual", "And", "Or", "Not", "Add",
		"Subtract", "Multiply", "Divide", "LeftParen", "RightParen", "LeftBracket",
		"RightBracket", "Iteration", "Question", "Literal", "NullCoalescing",
		"Range", "As", "Comma", "Arrow", "LeftBrace", "RightBrace", "Null",
		"False", "True", "ESCAPED_STRING", "STRING", "NUMBER", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "Less", "LessOrEqual", "Equal", "Greater", "GreaterOrEqual",
		"NotEqual", "And", "Or", "Not", "Add", "Subtract", "Multiply", "Divide",
		"LeftParen", "RightParen", "LeftBracket", "RightBracket", "Iteration",
		"Question", "Literal", "NullCoalescing", "Range", "As", "Comma", "Arrow",
		"LeftBrace", "RightBrace", "Null", "False", "True", "ESCAPED_STRING",
		"STRING", "NUMBER", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 44, 252, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30,
		1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1,
		34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 5, 40, 211, 8, 40, 10, 40, 12, 40,
		214, 9, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 5, 40, 221, 8, 40, 10, 40,
		12, 40, 224, 9, 40, 1, 40, 3, 40, 227, 8, 40, 1, 41, 1, 41, 5, 41, 231,
		8, 41, 10, 41, 12, 41, 234, 9, 41, 1, 42, 4, 42, 237, 8, 42, 11, 42, 12,
		42, 238, 1, 42, 1, 42, 4, 42, 243, 8, 42, 11, 42, 12, 42, 244, 3, 42, 247,
		8, 42, 1, 43, 1, 43, 1, 43, 1, 43, 0, 0, 44, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32,
		65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41,
		83, 42, 85, 43, 87, 44, 1, 0, 6, 2, 0, 34, 34, 92, 92, 2, 0, 39, 39, 92,
		92, 4, 0, 36, 36, 65, 90, 95, 95, 97, 122, 5, 0, 36, 36, 48, 57, 65, 90,
		95, 95, 97, 122, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 260, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0,
		0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1,
		0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55,
		1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0,
		63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0,
		0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0,
		0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0,
		0, 0, 0, 87, 1, 0, 0, 0, 1, 89, 1, 0, 0, 0, 3, 91, 1, 0, 0, 0, 5, 98, 1,
		0, 0, 0, 7, 102, 1, 0, 0, 0, 9, 105, 1, 0, 0, 0, 11, 108, 1, 0, 0, 0, 13,
		113, 1, 0, 0, 0, 15, 119, 1, 0, 0, 0, 17, 122, 1, 0, 0, 0, 19, 124, 1,
		0, 0, 0, 21, 126, 1, 0, 0, 0, 23, 128, 1, 0, 0, 0, 25, 131, 1, 0, 0, 0,
		27, 134, 1, 0, 0, 0, 29, 136, 1, 0, 0, 0, 31, 139, 1, 0, 0, 0, 33, 142,
		1, 0, 0, 0, 35, 145, 1, 0, 0, 0, 37, 148, 1, 0, 0, 0, 39, 150, 1, 0, 0,
		0, 41, 152, 1, 0, 0, 0, 43, 154, 1, 0, 0, 0, 45, 156, 1, 0, 0, 0, 47, 158,
		1, 0, 0, 0, 49, 160, 1, 0, 0, 0, 51, 162, 1, 0, 0, 0, 53, 164, 1, 0, 0,
		0, 55, 166, 1, 0, 0, 0, 57, 168, 1, 0, 0, 0, 59, 170, 1, 0, 0, 0, 61, 172,
		1, 0, 0, 0, 63, 175, 1, 0, 0, 0, 65, 178, 1, 0, 0, 0, 67, 181, 1, 0, 0,
		0, 69, 183, 1, 0, 0, 0, 71, 186, 1, 0, 0, 0, 73, 188, 1, 0, 0, 0, 75, 190,
		1, 0, 0, 0, 77, 195, 1, 0, 0, 0, 79, 201, 1, 0, 0, 0, 81, 226, 1, 0, 0,
		0, 83, 228, 1, 0, 0, 0, 85, 236, 1, 0, 0, 0, 87, 248, 1, 0, 0, 0, 89, 90,
		5, 59, 0, 0, 90, 2, 1, 0, 0, 0, 91, 92, 5, 114, 0, 0, 92, 93, 5, 101, 0,
		0, 93, 94, 5, 116, 0, 0, 94, 95, 5, 117, 0, 0, 95, 96, 5, 114, 0, 0, 96,
		97, 5, 110, 0, 0, 97, 4, 1, 0, 0, 0, 98, 99, 5, 102, 0, 0, 99, 100, 5,
		111, 0, 0, 100, 101, 5, 114, 0, 0, 101, 6, 1, 0, 0, 0, 102, 103, 5, 105,
		0, 0, 103, 104, 5, 110, 0, 0, 104, 8, 1, 0, 0, 0, 105, 106, 5, 105, 0,
		0, 106, 107, 5, 102, 0, 0, 107, 10, 1, 0, 0, 0, 108, 109, 5, 101, 0, 0,
		109, 110, 5, 108, 0, 0, 110, 111, 5, 115, 0, 0, 111, 112, 5, 101, 0, 0,
		112, 12, 1, 0, 0, 0, 113, 114, 5, 119, 0, 0, 114, 115, 5, 104, 0, 0, 115,
		116, 5, 105, 0, 0, 116, 117, 5, 108, 0, 0, 117, 118, 5, 101, 0, 0, 118,
		14, 1, 0, 0, 0, 119, 120, 5, 100, 0, 0, 120, 121, 5, 111, 0, 0, 121, 16,
		1, 0, 0, 0, 122, 123, 5, 46, 0, 0, 123, 18, 1, 0, 0, 0, 124, 125, 5, 58,
		0, 0, 125, 20, 1, 0, 0, 0, 126, 127, 5, 60, 0, 0, 127, 22, 1, 0, 0, 0,
		128, 129, 5, 60, 0, 0, 129, 130, 5, 61, 0, 0, 130, 24, 1, 0, 0, 0, 131,
		132, 5, 61, 0, 0, 132, 133, 5, 61, 0, 0, 133, 26, 1, 0, 0, 0, 134, 135,
		5, 62, 0, 0, 135, 28, 1, 0, 0, 0, 136, 137, 5, 62, 0, 0, 137, 138, 5, 61,
		0, 0, 138, 30, 1, 0, 0, 0, 139, 140, 5, 33, 0, 0, 140, 141, 5, 61, 0, 0,
		141, 32, 1, 0, 0, 0, 142, 143, 5, 38, 0, 0, 143, 144, 5, 38, 0, 0, 144,
		34, 1, 0, 0, 0, 145, 146, 5, 124, 0, 0, 146, 147, 5, 124, 0, 0, 147, 36,
		1, 0, 0, 0, 148, 149, 5, 33, 0, 0, 149, 38, 1, 0, 0, 0, 150, 151, 5, 43,
		0, 0, 151, 40, 1, 0, 0, 0, 152, 153, 5, 45, 0, 0, 153, 42, 1, 0, 0, 0,
		154, 155, 5, 42, 0, 0, 155, 44, 1, 0, 0, 0, 156, 157, 5, 47, 0, 0, 157,
		46, 1, 0, 0, 0, 158, 159, 5, 40, 0, 0, 159, 48, 1, 0, 0, 0, 160, 161, 5,
		41, 0, 0, 161, 50, 1, 0, 0, 0, 162, 163, 5, 91, 0, 0, 163, 52, 1, 0, 0,
		0, 164, 165, 5, 93, 0, 0, 165, 54, 1, 0, 0, 0, 166, 167, 5, 35, 0, 0, 167,
		56, 1, 0, 0, 0, 168, 169, 5, 63, 0, 0, 169, 58, 1, 0, 0, 0, 170, 171, 5,
		61, 0, 0, 171, 60, 1, 0, 0, 0, 172, 173, 5, 63, 0, 0, 173, 174, 5, 63,
		0, 0, 174, 62, 1, 0, 0, 0, 175, 176, 5, 46, 0, 0, 176, 177, 5, 46, 0, 0,
		177, 64, 1, 0, 0, 0, 178, 179, 5, 97, 0, 0, 179, 180, 5, 115, 0, 0, 180,
		66, 1, 0, 0, 0, 181, 182, 5, 44, 0, 0, 182, 68, 1, 0, 0, 0, 183, 184, 5,
		61, 0, 0, 184, 185, 5, 62, 0, 0, 185, 70, 1, 0, 0, 0, 186, 187, 5, 123,
		0, 0, 187, 72, 1, 0, 0, 0, 188, 189, 5, 125, 0, 0, 189, 74, 1, 0, 0, 0,
		190, 191, 5, 110, 0, 0, 191, 192, 5, 117, 0, 0, 192, 193, 5, 108, 0, 0,
		193, 194, 5, 108, 0, 0, 194, 76, 1, 0, 0, 0, 195, 196, 5, 102, 0, 0, 196,
		197, 5, 97, 0, 0, 197, 198, 5, 108, 0, 0, 198, 199, 5, 115, 0, 0, 199,
		200, 5, 101, 0, 0, 200, 78, 1, 0, 0, 0, 201, 202, 5, 116, 0, 0, 202, 203,
		5, 114, 0, 0, 203, 204, 5, 117, 0, 0, 204, 205, 5, 101, 0, 0, 205, 80,
		1, 0, 0, 0, 206, 212, 5, 34, 0, 0, 207, 208, 5, 92, 0, 0, 208, 211, 9,
		0, 0, 0, 209, 211, 8, 0, 0, 0, 210, 207, 1, 0, 0, 0, 210, 209, 1, 0, 0,
		0, 211, 214, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213,
		215, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 215, 227, 5, 34, 0, 0, 216, 222,
		5, 39, 0, 0, 217, 218, 5, 92, 0, 0, 218, 221, 9, 0, 0, 0, 219, 221, 8,
		1, 0, 0, 220, 217, 1, 0, 0, 0, 220, 219, 1, 0, 0, 0, 221, 224, 1, 0, 0,
		0, 222, 220, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 225, 1, 0, 0, 0, 224,
		222, 1, 0, 0, 0, 225, 227, 5, 39, 0, 0, 226, 206, 1, 0, 0, 0, 226, 216,
		1, 0, 0, 0, 227, 82, 1, 0, 0, 0, 228, 232, 7, 2, 0, 0, 229, 231, 7, 3,
		0, 0, 230, 229, 1, 0, 0, 0, 231, 234, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0,
		232, 233, 1, 0, 0, 0, 233, 84, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 235, 237,
		7, 4, 0, 0, 236, 235, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 236, 1, 0,
		0, 0, 238, 239, 1, 0, 0, 0, 239, 246, 1, 0, 0, 0, 240, 242, 5, 46, 0, 0,
		241, 243, 7, 4, 0, 0, 242, 241, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244,
		242, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 247, 1, 0, 0, 0, 246, 240,
		1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 86, 1, 0, 0, 0, 248, 249, 7, 5,
		0, 0, 249, 250, 1, 0, 0, 0, 250, 251, 6, 43, 0, 0, 251, 88, 1, 0, 0, 0,
		10, 0, 210, 212, 220, 222, 226, 232, 238, 244, 246, 1, 0, 1, 0,
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
	JsonTemplateLexerT__2           = 3
	JsonTemplateLexerT__3           = 4
	JsonTemplateLexerT__4           = 5
	JsonTemplateLexerT__5           = 6
	JsonTemplateLexerT__6           = 7
	JsonTemplateLexerT__7           = 8
	JsonTemplateLexerT__8           = 9
	JsonTemplateLexerT__9           = 10
	JsonTemplateLexerLess           = 11
	JsonTemplateLexerLessOrEqual    = 12
	JsonTemplateLexerEqual          = 13
	JsonTemplateLexerGreater        = 14
	JsonTemplateLexerGreaterOrEqual = 15
	JsonTemplateLexerNotEqual       = 16
	JsonTemplateLexerAnd            = 17
	JsonTemplateLexerOr             = 18
	JsonTemplateLexerNot            = 19
	JsonTemplateLexerAdd            = 20
	JsonTemplateLexerSubtract       = 21
	JsonTemplateLexerMultiply       = 22
	JsonTemplateLexerDivide         = 23
	JsonTemplateLexerLeftParen      = 24
	JsonTemplateLexerRightParen     = 25
	JsonTemplateLexerLeftBracket    = 26
	JsonTemplateLexerRightBracket   = 27
	JsonTemplateLexerIteration      = 28
	JsonTemplateLexerQuestion       = 29
	JsonTemplateLexerLiteral        = 30
	JsonTemplateLexerNullCoalescing = 31
	JsonTemplateLexerRange          = 32
	JsonTemplateLexerAs             = 33
	JsonTemplateLexerComma          = 34
	JsonTemplateLexerArrow          = 35
	JsonTemplateLexerLeftBrace      = 36
	JsonTemplateLexerRightBrace     = 37
	JsonTemplateLexerNull           = 38
	JsonTemplateLexerFalse          = 39
	JsonTemplateLexerTrue           = 40
	JsonTemplateLexerESCAPED_STRING = 41
	JsonTemplateLexerSTRING         = 42
	JsonTemplateLexerNUMBER         = 43
	JsonTemplateLexerWS             = 44
)
