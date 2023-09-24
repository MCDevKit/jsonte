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
		"']'", "'#'", "'?'", "'='", "'??'", "'..'", "'...'", "'as'", "','",
		"'=>'", "'{'", "'}'", "'null'", "'false'", "'true'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "Less", "LessOrEqual", "Equal",
		"Greater", "GreaterOrEqual", "NotEqual", "And", "Or", "Not", "Add",
		"Subtract", "Multiply", "Divide", "LeftParen", "RightParen", "LeftBracket",
		"RightBracket", "Iteration", "Question", "Literal", "NullCoalescing",
		"Range", "Spread", "As", "Comma", "Arrow", "LeftBrace", "RightBrace",
		"Null", "False", "True", "ESCAPED_STRING", "STRING", "NUMBER", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "Less", "LessOrEqual", "Equal", "Greater", "GreaterOrEqual",
		"NotEqual", "And", "Or", "Not", "Add", "Subtract", "Multiply", "Divide",
		"LeftParen", "RightParen", "LeftBracket", "RightBracket", "Iteration",
		"Question", "Literal", "NullCoalescing", "Range", "Spread", "As", "Comma",
		"Arrow", "LeftBrace", "RightBrace", "Null", "False", "True", "ESCAPED_STRING",
		"STRING", "NUMBER", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 45, 258, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29,
		1, 29, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36,
		1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41,
		1, 41, 5, 41, 217, 8, 41, 10, 41, 12, 41, 220, 9, 41, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 41, 5, 41, 227, 8, 41, 10, 41, 12, 41, 230, 9, 41, 1, 41,
		3, 41, 233, 8, 41, 1, 42, 1, 42, 5, 42, 237, 8, 42, 10, 42, 12, 42, 240,
		9, 42, 1, 43, 4, 43, 243, 8, 43, 11, 43, 12, 43, 244, 1, 43, 1, 43, 4,
		43, 249, 8, 43, 11, 43, 12, 43, 250, 3, 43, 253, 8, 43, 1, 44, 1, 44, 1,
		44, 1, 44, 0, 0, 45, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8,
		17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35,
		71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44,
		89, 45, 1, 0, 6, 2, 0, 34, 34, 92, 92, 2, 0, 39, 39, 92, 92, 4, 0, 36,
		36, 65, 90, 95, 95, 97, 122, 5, 0, 36, 36, 48, 57, 65, 90, 95, 95, 97,
		122, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 266, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1,
		0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41,
		1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0,
		49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0,
		0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0,
		0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0,
		0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1,
		0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87,
		1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 1, 91, 1, 0, 0, 0, 3, 93, 1, 0, 0, 0, 5,
		100, 1, 0, 0, 0, 7, 104, 1, 0, 0, 0, 9, 107, 1, 0, 0, 0, 11, 110, 1, 0,
		0, 0, 13, 115, 1, 0, 0, 0, 15, 121, 1, 0, 0, 0, 17, 124, 1, 0, 0, 0, 19,
		126, 1, 0, 0, 0, 21, 128, 1, 0, 0, 0, 23, 130, 1, 0, 0, 0, 25, 133, 1,
		0, 0, 0, 27, 136, 1, 0, 0, 0, 29, 138, 1, 0, 0, 0, 31, 141, 1, 0, 0, 0,
		33, 144, 1, 0, 0, 0, 35, 147, 1, 0, 0, 0, 37, 150, 1, 0, 0, 0, 39, 152,
		1, 0, 0, 0, 41, 154, 1, 0, 0, 0, 43, 156, 1, 0, 0, 0, 45, 158, 1, 0, 0,
		0, 47, 160, 1, 0, 0, 0, 49, 162, 1, 0, 0, 0, 51, 164, 1, 0, 0, 0, 53, 166,
		1, 0, 0, 0, 55, 168, 1, 0, 0, 0, 57, 170, 1, 0, 0, 0, 59, 172, 1, 0, 0,
		0, 61, 174, 1, 0, 0, 0, 63, 177, 1, 0, 0, 0, 65, 180, 1, 0, 0, 0, 67, 184,
		1, 0, 0, 0, 69, 187, 1, 0, 0, 0, 71, 189, 1, 0, 0, 0, 73, 192, 1, 0, 0,
		0, 75, 194, 1, 0, 0, 0, 77, 196, 1, 0, 0, 0, 79, 201, 1, 0, 0, 0, 81, 207,
		1, 0, 0, 0, 83, 232, 1, 0, 0, 0, 85, 234, 1, 0, 0, 0, 87, 242, 1, 0, 0,
		0, 89, 254, 1, 0, 0, 0, 91, 92, 5, 59, 0, 0, 92, 2, 1, 0, 0, 0, 93, 94,
		5, 114, 0, 0, 94, 95, 5, 101, 0, 0, 95, 96, 5, 116, 0, 0, 96, 97, 5, 117,
		0, 0, 97, 98, 5, 114, 0, 0, 98, 99, 5, 110, 0, 0, 99, 4, 1, 0, 0, 0, 100,
		101, 5, 102, 0, 0, 101, 102, 5, 111, 0, 0, 102, 103, 5, 114, 0, 0, 103,
		6, 1, 0, 0, 0, 104, 105, 5, 105, 0, 0, 105, 106, 5, 110, 0, 0, 106, 8,
		1, 0, 0, 0, 107, 108, 5, 105, 0, 0, 108, 109, 5, 102, 0, 0, 109, 10, 1,
		0, 0, 0, 110, 111, 5, 101, 0, 0, 111, 112, 5, 108, 0, 0, 112, 113, 5, 115,
		0, 0, 113, 114, 5, 101, 0, 0, 114, 12, 1, 0, 0, 0, 115, 116, 5, 119, 0,
		0, 116, 117, 5, 104, 0, 0, 117, 118, 5, 105, 0, 0, 118, 119, 5, 108, 0,
		0, 119, 120, 5, 101, 0, 0, 120, 14, 1, 0, 0, 0, 121, 122, 5, 100, 0, 0,
		122, 123, 5, 111, 0, 0, 123, 16, 1, 0, 0, 0, 124, 125, 5, 46, 0, 0, 125,
		18, 1, 0, 0, 0, 126, 127, 5, 58, 0, 0, 127, 20, 1, 0, 0, 0, 128, 129, 5,
		60, 0, 0, 129, 22, 1, 0, 0, 0, 130, 131, 5, 60, 0, 0, 131, 132, 5, 61,
		0, 0, 132, 24, 1, 0, 0, 0, 133, 134, 5, 61, 0, 0, 134, 135, 5, 61, 0, 0,
		135, 26, 1, 0, 0, 0, 136, 137, 5, 62, 0, 0, 137, 28, 1, 0, 0, 0, 138, 139,
		5, 62, 0, 0, 139, 140, 5, 61, 0, 0, 140, 30, 1, 0, 0, 0, 141, 142, 5, 33,
		0, 0, 142, 143, 5, 61, 0, 0, 143, 32, 1, 0, 0, 0, 144, 145, 5, 38, 0, 0,
		145, 146, 5, 38, 0, 0, 146, 34, 1, 0, 0, 0, 147, 148, 5, 124, 0, 0, 148,
		149, 5, 124, 0, 0, 149, 36, 1, 0, 0, 0, 150, 151, 5, 33, 0, 0, 151, 38,
		1, 0, 0, 0, 152, 153, 5, 43, 0, 0, 153, 40, 1, 0, 0, 0, 154, 155, 5, 45,
		0, 0, 155, 42, 1, 0, 0, 0, 156, 157, 5, 42, 0, 0, 157, 44, 1, 0, 0, 0,
		158, 159, 5, 47, 0, 0, 159, 46, 1, 0, 0, 0, 160, 161, 5, 40, 0, 0, 161,
		48, 1, 0, 0, 0, 162, 163, 5, 41, 0, 0, 163, 50, 1, 0, 0, 0, 164, 165, 5,
		91, 0, 0, 165, 52, 1, 0, 0, 0, 166, 167, 5, 93, 0, 0, 167, 54, 1, 0, 0,
		0, 168, 169, 5, 35, 0, 0, 169, 56, 1, 0, 0, 0, 170, 171, 5, 63, 0, 0, 171,
		58, 1, 0, 0, 0, 172, 173, 5, 61, 0, 0, 173, 60, 1, 0, 0, 0, 174, 175, 5,
		63, 0, 0, 175, 176, 5, 63, 0, 0, 176, 62, 1, 0, 0, 0, 177, 178, 5, 46,
		0, 0, 178, 179, 5, 46, 0, 0, 179, 64, 1, 0, 0, 0, 180, 181, 5, 46, 0, 0,
		181, 182, 5, 46, 0, 0, 182, 183, 5, 46, 0, 0, 183, 66, 1, 0, 0, 0, 184,
		185, 5, 97, 0, 0, 185, 186, 5, 115, 0, 0, 186, 68, 1, 0, 0, 0, 187, 188,
		5, 44, 0, 0, 188, 70, 1, 0, 0, 0, 189, 190, 5, 61, 0, 0, 190, 191, 5, 62,
		0, 0, 191, 72, 1, 0, 0, 0, 192, 193, 5, 123, 0, 0, 193, 74, 1, 0, 0, 0,
		194, 195, 5, 125, 0, 0, 195, 76, 1, 0, 0, 0, 196, 197, 5, 110, 0, 0, 197,
		198, 5, 117, 0, 0, 198, 199, 5, 108, 0, 0, 199, 200, 5, 108, 0, 0, 200,
		78, 1, 0, 0, 0, 201, 202, 5, 102, 0, 0, 202, 203, 5, 97, 0, 0, 203, 204,
		5, 108, 0, 0, 204, 205, 5, 115, 0, 0, 205, 206, 5, 101, 0, 0, 206, 80,
		1, 0, 0, 0, 207, 208, 5, 116, 0, 0, 208, 209, 5, 114, 0, 0, 209, 210, 5,
		117, 0, 0, 210, 211, 5, 101, 0, 0, 211, 82, 1, 0, 0, 0, 212, 218, 5, 34,
		0, 0, 213, 214, 5, 92, 0, 0, 214, 217, 9, 0, 0, 0, 215, 217, 8, 0, 0, 0,
		216, 213, 1, 0, 0, 0, 216, 215, 1, 0, 0, 0, 217, 220, 1, 0, 0, 0, 218,
		216, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 221, 1, 0, 0, 0, 220, 218,
		1, 0, 0, 0, 221, 233, 5, 34, 0, 0, 222, 228, 5, 39, 0, 0, 223, 224, 5,
		92, 0, 0, 224, 227, 9, 0, 0, 0, 225, 227, 8, 1, 0, 0, 226, 223, 1, 0, 0,
		0, 226, 225, 1, 0, 0, 0, 227, 230, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228,
		229, 1, 0, 0, 0, 229, 231, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 231, 233,
		5, 39, 0, 0, 232, 212, 1, 0, 0, 0, 232, 222, 1, 0, 0, 0, 233, 84, 1, 0,
		0, 0, 234, 238, 7, 2, 0, 0, 235, 237, 7, 3, 0, 0, 236, 235, 1, 0, 0, 0,
		237, 240, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239,
		86, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 241, 243, 7, 4, 0, 0, 242, 241, 1,
		0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 244, 245, 1, 0, 0,
		0, 245, 252, 1, 0, 0, 0, 246, 248, 5, 46, 0, 0, 247, 249, 7, 4, 0, 0, 248,
		247, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 250, 251,
		1, 0, 0, 0, 251, 253, 1, 0, 0, 0, 252, 246, 1, 0, 0, 0, 252, 253, 1, 0,
		0, 0, 253, 88, 1, 0, 0, 0, 254, 255, 7, 5, 0, 0, 255, 256, 1, 0, 0, 0,
		256, 257, 6, 44, 0, 0, 257, 90, 1, 0, 0, 0, 10, 0, 216, 218, 226, 228,
		232, 238, 244, 250, 252, 1, 0, 1, 0,
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
	JsonTemplateLexerSpread         = 33
	JsonTemplateLexerAs             = 34
	JsonTemplateLexerComma          = 35
	JsonTemplateLexerArrow          = 36
	JsonTemplateLexerLeftBrace      = 37
	JsonTemplateLexerRightBrace     = 38
	JsonTemplateLexerNull           = 39
	JsonTemplateLexerFalse          = 40
	JsonTemplateLexerTrue           = 41
	JsonTemplateLexerESCAPED_STRING = 42
	JsonTemplateLexerSTRING         = 43
	JsonTemplateLexerNUMBER         = 44
	JsonTemplateLexerWS             = 45
)
