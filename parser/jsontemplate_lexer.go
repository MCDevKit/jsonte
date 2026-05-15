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
		"", "'<'", "'<='", "'=='", "'>'", "'>='", "'!='", "'&&'", "'||'", "'!'",
		"'+'", "'-'", "'*'", "'/'", "'('", "')'", "'['", "']'", "'#'", "'?'",
		"'+='", "'='", "'??'", "'..'", "'...'", "'as'", "','", "'=>'", "':'",
		"';'", "'.'", "'{'", "'}'", "'null'", "'false'", "'true'", "'return'",
		"'for'", "'in'", "'if'", "'else'", "'while'", "'do'", "'break'", "'continue'",
	}
	staticData.symbolicNames = []string{
		"", "Less", "LessOrEqual", "Equal", "Greater", "GreaterOrEqual", "NotEqual",
		"And", "Or", "Not", "Add", "Subtract", "Multiply", "Divide", "LeftParen",
		"RightParen", "LeftBracket", "RightBracket", "Iteration", "Question",
		"AddAssign", "Literal", "NullCoalescing", "Range", "Spread", "As", "Comma",
		"Arrow", "Colon", "Semicolon", "Dot", "LeftBrace", "RightBrace", "Null",
		"False", "True", "Return", "For", "In", "If", "Else", "While", "Do",
		"Break", "Continue", "ESCAPED_STRING", "STRING", "NUMBER", "WS",
	}
	staticData.ruleNames = []string{
		"Less", "LessOrEqual", "Equal", "Greater", "GreaterOrEqual", "NotEqual",
		"And", "Or", "Not", "Add", "Subtract", "Multiply", "Divide", "LeftParen",
		"RightParen", "LeftBracket", "RightBracket", "Iteration", "Question",
		"AddAssign", "Literal", "NullCoalescing", "Range", "Spread", "As", "Comma",
		"Arrow", "Colon", "Semicolon", "Dot", "LeftBrace", "RightBrace", "Null",
		"False", "True", "Return", "For", "In", "If", "Else", "While", "Do",
		"Break", "Continue", "ESCAPED_STRING", "STRING", "NUMBER", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 48, 282, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1,
		30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37,
		1, 37, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42,
		1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1,
		43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 5, 44, 241, 8, 44, 10, 44, 12, 44,
		244, 9, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 5, 44, 251, 8, 44, 10, 44,
		12, 44, 254, 9, 44, 1, 44, 3, 44, 257, 8, 44, 1, 45, 1, 45, 5, 45, 261,
		8, 45, 10, 45, 12, 45, 264, 9, 45, 1, 46, 4, 46, 267, 8, 46, 11, 46, 12,
		46, 268, 1, 46, 1, 46, 4, 46, 273, 8, 46, 11, 46, 12, 46, 274, 3, 46, 277,
		8, 46, 1, 47, 1, 47, 1, 47, 1, 47, 0, 0, 48, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32,
		65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41,
		83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 1, 0, 6, 2, 0,
		34, 34, 92, 92, 2, 0, 39, 39, 92, 92, 4, 0, 36, 36, 65, 90, 95, 95, 97,
		122, 5, 0, 36, 36, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 3, 0,
		9, 10, 13, 13, 32, 32, 290, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1,
		0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59,
		1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0,
		67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0,
		0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0,
		0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0,
		0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 1, 97, 1,
		0, 0, 0, 3, 99, 1, 0, 0, 0, 5, 102, 1, 0, 0, 0, 7, 105, 1, 0, 0, 0, 9,
		107, 1, 0, 0, 0, 11, 110, 1, 0, 0, 0, 13, 113, 1, 0, 0, 0, 15, 116, 1,
		0, 0, 0, 17, 119, 1, 0, 0, 0, 19, 121, 1, 0, 0, 0, 21, 123, 1, 0, 0, 0,
		23, 125, 1, 0, 0, 0, 25, 127, 1, 0, 0, 0, 27, 129, 1, 0, 0, 0, 29, 131,
		1, 0, 0, 0, 31, 133, 1, 0, 0, 0, 33, 135, 1, 0, 0, 0, 35, 137, 1, 0, 0,
		0, 37, 139, 1, 0, 0, 0, 39, 141, 1, 0, 0, 0, 41, 144, 1, 0, 0, 0, 43, 146,
		1, 0, 0, 0, 45, 149, 1, 0, 0, 0, 47, 152, 1, 0, 0, 0, 49, 156, 1, 0, 0,
		0, 51, 159, 1, 0, 0, 0, 53, 161, 1, 0, 0, 0, 55, 164, 1, 0, 0, 0, 57, 166,
		1, 0, 0, 0, 59, 168, 1, 0, 0, 0, 61, 170, 1, 0, 0, 0, 63, 172, 1, 0, 0,
		0, 65, 174, 1, 0, 0, 0, 67, 179, 1, 0, 0, 0, 69, 185, 1, 0, 0, 0, 71, 190,
		1, 0, 0, 0, 73, 197, 1, 0, 0, 0, 75, 201, 1, 0, 0, 0, 77, 204, 1, 0, 0,
		0, 79, 207, 1, 0, 0, 0, 81, 212, 1, 0, 0, 0, 83, 218, 1, 0, 0, 0, 85, 221,
		1, 0, 0, 0, 87, 227, 1, 0, 0, 0, 89, 256, 1, 0, 0, 0, 91, 258, 1, 0, 0,
		0, 93, 266, 1, 0, 0, 0, 95, 278, 1, 0, 0, 0, 97, 98, 5, 60, 0, 0, 98, 2,
		1, 0, 0, 0, 99, 100, 5, 60, 0, 0, 100, 101, 5, 61, 0, 0, 101, 4, 1, 0,
		0, 0, 102, 103, 5, 61, 0, 0, 103, 104, 5, 61, 0, 0, 104, 6, 1, 0, 0, 0,
		105, 106, 5, 62, 0, 0, 106, 8, 1, 0, 0, 0, 107, 108, 5, 62, 0, 0, 108,
		109, 5, 61, 0, 0, 109, 10, 1, 0, 0, 0, 110, 111, 5, 33, 0, 0, 111, 112,
		5, 61, 0, 0, 112, 12, 1, 0, 0, 0, 113, 114, 5, 38, 0, 0, 114, 115, 5, 38,
		0, 0, 115, 14, 1, 0, 0, 0, 116, 117, 5, 124, 0, 0, 117, 118, 5, 124, 0,
		0, 118, 16, 1, 0, 0, 0, 119, 120, 5, 33, 0, 0, 120, 18, 1, 0, 0, 0, 121,
		122, 5, 43, 0, 0, 122, 20, 1, 0, 0, 0, 123, 124, 5, 45, 0, 0, 124, 22,
		1, 0, 0, 0, 125, 126, 5, 42, 0, 0, 126, 24, 1, 0, 0, 0, 127, 128, 5, 47,
		0, 0, 128, 26, 1, 0, 0, 0, 129, 130, 5, 40, 0, 0, 130, 28, 1, 0, 0, 0,
		131, 132, 5, 41, 0, 0, 132, 30, 1, 0, 0, 0, 133, 134, 5, 91, 0, 0, 134,
		32, 1, 0, 0, 0, 135, 136, 5, 93, 0, 0, 136, 34, 1, 0, 0, 0, 137, 138, 5,
		35, 0, 0, 138, 36, 1, 0, 0, 0, 139, 140, 5, 63, 0, 0, 140, 38, 1, 0, 0,
		0, 141, 142, 5, 43, 0, 0, 142, 143, 5, 61, 0, 0, 143, 40, 1, 0, 0, 0, 144,
		145, 5, 61, 0, 0, 145, 42, 1, 0, 0, 0, 146, 147, 5, 63, 0, 0, 147, 148,
		5, 63, 0, 0, 148, 44, 1, 0, 0, 0, 149, 150, 5, 46, 0, 0, 150, 151, 5, 46,
		0, 0, 151, 46, 1, 0, 0, 0, 152, 153, 5, 46, 0, 0, 153, 154, 5, 46, 0, 0,
		154, 155, 5, 46, 0, 0, 155, 48, 1, 0, 0, 0, 156, 157, 5, 97, 0, 0, 157,
		158, 5, 115, 0, 0, 158, 50, 1, 0, 0, 0, 159, 160, 5, 44, 0, 0, 160, 52,
		1, 0, 0, 0, 161, 162, 5, 61, 0, 0, 162, 163, 5, 62, 0, 0, 163, 54, 1, 0,
		0, 0, 164, 165, 5, 58, 0, 0, 165, 56, 1, 0, 0, 0, 166, 167, 5, 59, 0, 0,
		167, 58, 1, 0, 0, 0, 168, 169, 5, 46, 0, 0, 169, 60, 1, 0, 0, 0, 170, 171,
		5, 123, 0, 0, 171, 62, 1, 0, 0, 0, 172, 173, 5, 125, 0, 0, 173, 64, 1,
		0, 0, 0, 174, 175, 5, 110, 0, 0, 175, 176, 5, 117, 0, 0, 176, 177, 5, 108,
		0, 0, 177, 178, 5, 108, 0, 0, 178, 66, 1, 0, 0, 0, 179, 180, 5, 102, 0,
		0, 180, 181, 5, 97, 0, 0, 181, 182, 5, 108, 0, 0, 182, 183, 5, 115, 0,
		0, 183, 184, 5, 101, 0, 0, 184, 68, 1, 0, 0, 0, 185, 186, 5, 116, 0, 0,
		186, 187, 5, 114, 0, 0, 187, 188, 5, 117, 0, 0, 188, 189, 5, 101, 0, 0,
		189, 70, 1, 0, 0, 0, 190, 191, 5, 114, 0, 0, 191, 192, 5, 101, 0, 0, 192,
		193, 5, 116, 0, 0, 193, 194, 5, 117, 0, 0, 194, 195, 5, 114, 0, 0, 195,
		196, 5, 110, 0, 0, 196, 72, 1, 0, 0, 0, 197, 198, 5, 102, 0, 0, 198, 199,
		5, 111, 0, 0, 199, 200, 5, 114, 0, 0, 200, 74, 1, 0, 0, 0, 201, 202, 5,
		105, 0, 0, 202, 203, 5, 110, 0, 0, 203, 76, 1, 0, 0, 0, 204, 205, 5, 105,
		0, 0, 205, 206, 5, 102, 0, 0, 206, 78, 1, 0, 0, 0, 207, 208, 5, 101, 0,
		0, 208, 209, 5, 108, 0, 0, 209, 210, 5, 115, 0, 0, 210, 211, 5, 101, 0,
		0, 211, 80, 1, 0, 0, 0, 212, 213, 5, 119, 0, 0, 213, 214, 5, 104, 0, 0,
		214, 215, 5, 105, 0, 0, 215, 216, 5, 108, 0, 0, 216, 217, 5, 101, 0, 0,
		217, 82, 1, 0, 0, 0, 218, 219, 5, 100, 0, 0, 219, 220, 5, 111, 0, 0, 220,
		84, 1, 0, 0, 0, 221, 222, 5, 98, 0, 0, 222, 223, 5, 114, 0, 0, 223, 224,
		5, 101, 0, 0, 224, 225, 5, 97, 0, 0, 225, 226, 5, 107, 0, 0, 226, 86, 1,
		0, 0, 0, 227, 228, 5, 99, 0, 0, 228, 229, 5, 111, 0, 0, 229, 230, 5, 110,
		0, 0, 230, 231, 5, 116, 0, 0, 231, 232, 5, 105, 0, 0, 232, 233, 5, 110,
		0, 0, 233, 234, 5, 117, 0, 0, 234, 235, 5, 101, 0, 0, 235, 88, 1, 0, 0,
		0, 236, 242, 5, 34, 0, 0, 237, 238, 5, 92, 0, 0, 238, 241, 9, 0, 0, 0,
		239, 241, 8, 0, 0, 0, 240, 237, 1, 0, 0, 0, 240, 239, 1, 0, 0, 0, 241,
		244, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0, 243, 245,
		1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 245, 257, 5, 34, 0, 0, 246, 252, 5, 39,
		0, 0, 247, 248, 5, 92, 0, 0, 248, 251, 9, 0, 0, 0, 249, 251, 8, 1, 0, 0,
		250, 247, 1, 0, 0, 0, 250, 249, 1, 0, 0, 0, 251, 254, 1, 0, 0, 0, 252,
		250, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 255, 1, 0, 0, 0, 254, 252,
		1, 0, 0, 0, 255, 257, 5, 39, 0, 0, 256, 236, 1, 0, 0, 0, 256, 246, 1, 0,
		0, 0, 257, 90, 1, 0, 0, 0, 258, 262, 7, 2, 0, 0, 259, 261, 7, 3, 0, 0,
		260, 259, 1, 0, 0, 0, 261, 264, 1, 0, 0, 0, 262, 260, 1, 0, 0, 0, 262,
		263, 1, 0, 0, 0, 263, 92, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0, 265, 267, 7,
		4, 0, 0, 266, 265, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 266, 1, 0, 0,
		0, 268, 269, 1, 0, 0, 0, 269, 276, 1, 0, 0, 0, 270, 272, 5, 46, 0, 0, 271,
		273, 7, 4, 0, 0, 272, 271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 272,
		1, 0, 0, 0, 274, 275, 1, 0, 0, 0, 275, 277, 1, 0, 0, 0, 276, 270, 1, 0,
		0, 0, 276, 277, 1, 0, 0, 0, 277, 94, 1, 0, 0, 0, 278, 279, 7, 5, 0, 0,
		279, 280, 1, 0, 0, 0, 280, 281, 6, 47, 0, 0, 281, 96, 1, 0, 0, 0, 10, 0,
		240, 242, 250, 252, 256, 262, 268, 274, 276, 1, 0, 1, 0,
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
	JsonTemplateLexerLess           = 1
	JsonTemplateLexerLessOrEqual    = 2
	JsonTemplateLexerEqual          = 3
	JsonTemplateLexerGreater        = 4
	JsonTemplateLexerGreaterOrEqual = 5
	JsonTemplateLexerNotEqual       = 6
	JsonTemplateLexerAnd            = 7
	JsonTemplateLexerOr             = 8
	JsonTemplateLexerNot            = 9
	JsonTemplateLexerAdd            = 10
	JsonTemplateLexerSubtract       = 11
	JsonTemplateLexerMultiply       = 12
	JsonTemplateLexerDivide         = 13
	JsonTemplateLexerLeftParen      = 14
	JsonTemplateLexerRightParen     = 15
	JsonTemplateLexerLeftBracket    = 16
	JsonTemplateLexerRightBracket   = 17
	JsonTemplateLexerIteration      = 18
	JsonTemplateLexerQuestion       = 19
	JsonTemplateLexerAddAssign      = 20
	JsonTemplateLexerLiteral        = 21
	JsonTemplateLexerNullCoalescing = 22
	JsonTemplateLexerRange          = 23
	JsonTemplateLexerSpread         = 24
	JsonTemplateLexerAs             = 25
	JsonTemplateLexerComma          = 26
	JsonTemplateLexerArrow          = 27
	JsonTemplateLexerColon          = 28
	JsonTemplateLexerSemicolon      = 29
	JsonTemplateLexerDot            = 30
	JsonTemplateLexerLeftBrace      = 31
	JsonTemplateLexerRightBrace     = 32
	JsonTemplateLexerNull           = 33
	JsonTemplateLexerFalse          = 34
	JsonTemplateLexerTrue           = 35
	JsonTemplateLexerReturn         = 36
	JsonTemplateLexerFor            = 37
	JsonTemplateLexerIn             = 38
	JsonTemplateLexerIf             = 39
	JsonTemplateLexerElse           = 40
	JsonTemplateLexerWhile          = 41
	JsonTemplateLexerDo             = 42
	JsonTemplateLexerBreak          = 43
	JsonTemplateLexerContinue       = 44
	JsonTemplateLexerESCAPED_STRING = 45
	JsonTemplateLexerSTRING         = 46
	JsonTemplateLexerNUMBER         = 47
	JsonTemplateLexerWS             = 48
)
