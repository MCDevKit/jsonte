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
		"DEFAULT_MODE", "TEMPLATE_MODE",
	}
	staticData.literalNames = []string{
		"", "'<'", "'<='", "'=='", "'>'", "'>='", "'!='", "'&&'", "'||'", "'!'",
		"'+'", "'-'", "'*'", "'/'", "'('", "')'", "'['", "']'", "'#'", "'?'",
		"'+='", "'='", "'??'", "'..'", "'...'", "'as'", "','", "'=>'", "':'",
		"';'", "'.'", "'{'", "'}'", "'null'", "'false'", "'true'", "'return'",
		"'for'", "'in'", "'if'", "'else'", "'while'", "'do'", "'break'", "'continue'",
		"", "", "", "", "", "", "", "", "'${'",
	}
	staticData.symbolicNames = []string{
		"", "Less", "LessOrEqual", "Equal", "Greater", "GreaterOrEqual", "NotEqual",
		"And", "Or", "Not", "Add", "Subtract", "Multiply", "Divide", "LeftParen",
		"RightParen", "LeftBracket", "RightBracket", "Iteration", "Question",
		"AddAssign", "Literal", "NullCoalescing", "Range", "Spread", "As", "Comma",
		"Arrow", "Colon", "Semicolon", "Dot", "LeftBrace", "RightBrace", "Null",
		"False", "True", "Return", "For", "In", "If", "Else", "While", "Do",
		"Break", "Continue", "Backtick", "ESCAPED_STRING", "STRING", "NUMBER",
		"LINE_COMMENT", "BLOCK_COMMENT", "WS", "TEMPLATE_TEXT", "TEMPLATE_INTERP_START",
		"TEMPLATE_END",
	}
	staticData.ruleNames = []string{
		"Less", "LessOrEqual", "Equal", "Greater", "GreaterOrEqual", "NotEqual",
		"And", "Or", "Not", "Add", "Subtract", "Multiply", "Divide", "LeftParen",
		"RightParen", "LeftBracket", "RightBracket", "Iteration", "Question",
		"AddAssign", "Literal", "NullCoalescing", "Range", "Spread", "As", "Comma",
		"Arrow", "Colon", "Semicolon", "Dot", "LeftBrace", "RightBrace", "Null",
		"False", "True", "Return", "For", "In", "If", "Else", "While", "Do",
		"Break", "Continue", "Backtick", "ESCAPED_STRING", "STRING", "NUMBER",
		"LINE_COMMENT", "BLOCK_COMMENT", "WS", "TEMPLATE_TEXT", "TEMPLATE_INTERP_START",
		"TEMPLATE_END",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 54, 340, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3,
		7, 3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9,
		7, 9, 2, 10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7,
		14, 2, 15, 7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19,
		2, 20, 7, 20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2,
		25, 7, 25, 2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30,
		7, 30, 2, 31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7,
		35, 2, 36, 7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40,
		2, 41, 7, 41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2,
		46, 7, 46, 2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51,
		7, 51, 2, 52, 7, 52, 2, 53, 7, 53, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41,
		1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1,
		43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45,
		1, 45, 1, 45, 5, 45, 258, 8, 45, 10, 45, 12, 45, 261, 9, 45, 1, 45, 1,
		45, 1, 45, 1, 45, 1, 45, 5, 45, 268, 8, 45, 10, 45, 12, 45, 271, 9, 45,
		1, 45, 3, 45, 274, 8, 45, 1, 46, 1, 46, 5, 46, 278, 8, 46, 10, 46, 12,
		46, 281, 9, 46, 1, 47, 4, 47, 284, 8, 47, 11, 47, 12, 47, 285, 1, 47, 1,
		47, 4, 47, 290, 8, 47, 11, 47, 12, 47, 291, 3, 47, 294, 8, 47, 1, 48, 1,
		48, 1, 48, 1, 48, 5, 48, 300, 8, 48, 10, 48, 12, 48, 303, 9, 48, 1, 48,
		1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 5, 49, 311, 8, 49, 10, 49, 12, 49, 314,
		9, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1,
		51, 1, 51, 1, 51, 1, 51, 1, 51, 4, 51, 330, 8, 51, 11, 51, 12, 51, 331,
		1, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1, 312, 0, 54, 2, 1, 4,
		2, 6, 3, 8, 4, 10, 5, 12, 6, 14, 7, 16, 8, 18, 9, 20, 10, 22, 11, 24, 12,
		26, 13, 28, 14, 30, 15, 32, 16, 34, 17, 36, 18, 38, 19, 40, 20, 42, 21,
		44, 22, 46, 23, 48, 24, 50, 25, 52, 26, 54, 27, 56, 28, 58, 29, 60, 30,
		62, 31, 64, 32, 66, 33, 68, 34, 70, 35, 72, 36, 74, 37, 76, 38, 78, 39,
		80, 40, 82, 41, 84, 42, 86, 43, 88, 44, 90, 45, 92, 46, 94, 47, 96, 48,
		98, 49, 100, 50, 102, 51, 104, 52, 106, 53, 108, 54, 2, 0, 1, 9, 2, 0,
		34, 34, 92, 92, 2, 0, 39, 39, 92, 92, 4, 0, 36, 36, 65, 90, 95, 95, 97,
		122, 5, 0, 36, 36, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 2, 0,
		10, 10, 13, 13, 3, 0, 9, 10, 13, 13, 32, 32, 3, 0, 36, 36, 92, 92, 96,
		96, 2, 0, 96, 96, 123, 123, 352, 0, 2, 1, 0, 0, 0, 0, 4, 1, 0, 0, 0, 0,
		6, 1, 0, 0, 0, 0, 8, 1, 0, 0, 0, 0, 10, 1, 0, 0, 0, 0, 12, 1, 0, 0, 0,
		0, 14, 1, 0, 0, 0, 0, 16, 1, 0, 0, 0, 0, 18, 1, 0, 0, 0, 0, 20, 1, 0, 0,
		0, 0, 22, 1, 0, 0, 0, 0, 24, 1, 0, 0, 0, 0, 26, 1, 0, 0, 0, 0, 28, 1, 0,
		0, 0, 0, 30, 1, 0, 0, 0, 0, 32, 1, 0, 0, 0, 0, 34, 1, 0, 0, 0, 0, 36, 1,
		0, 0, 0, 0, 38, 1, 0, 0, 0, 0, 40, 1, 0, 0, 0, 0, 42, 1, 0, 0, 0, 0, 44,
		1, 0, 0, 0, 0, 46, 1, 0, 0, 0, 0, 48, 1, 0, 0, 0, 0, 50, 1, 0, 0, 0, 0,
		52, 1, 0, 0, 0, 0, 54, 1, 0, 0, 0, 0, 56, 1, 0, 0, 0, 0, 58, 1, 0, 0, 0,
		0, 60, 1, 0, 0, 0, 0, 62, 1, 0, 0, 0, 0, 64, 1, 0, 0, 0, 0, 66, 1, 0, 0,
		0, 0, 68, 1, 0, 0, 0, 0, 70, 1, 0, 0, 0, 0, 72, 1, 0, 0, 0, 0, 74, 1, 0,
		0, 0, 0, 76, 1, 0, 0, 0, 0, 78, 1, 0, 0, 0, 0, 80, 1, 0, 0, 0, 0, 82, 1,
		0, 0, 0, 0, 84, 1, 0, 0, 0, 0, 86, 1, 0, 0, 0, 0, 88, 1, 0, 0, 0, 0, 90,
		1, 0, 0, 0, 0, 92, 1, 0, 0, 0, 0, 94, 1, 0, 0, 0, 0, 96, 1, 0, 0, 0, 0,
		98, 1, 0, 0, 0, 0, 100, 1, 0, 0, 0, 0, 102, 1, 0, 0, 0, 1, 104, 1, 0, 0,
		0, 1, 106, 1, 0, 0, 0, 1, 108, 1, 0, 0, 0, 2, 110, 1, 0, 0, 0, 4, 112,
		1, 0, 0, 0, 6, 115, 1, 0, 0, 0, 8, 118, 1, 0, 0, 0, 10, 120, 1, 0, 0, 0,
		12, 123, 1, 0, 0, 0, 14, 126, 1, 0, 0, 0, 16, 129, 1, 0, 0, 0, 18, 132,
		1, 0, 0, 0, 20, 134, 1, 0, 0, 0, 22, 136, 1, 0, 0, 0, 24, 138, 1, 0, 0,
		0, 26, 140, 1, 0, 0, 0, 28, 142, 1, 0, 0, 0, 30, 144, 1, 0, 0, 0, 32, 146,
		1, 0, 0, 0, 34, 148, 1, 0, 0, 0, 36, 150, 1, 0, 0, 0, 38, 152, 1, 0, 0,
		0, 40, 154, 1, 0, 0, 0, 42, 157, 1, 0, 0, 0, 44, 159, 1, 0, 0, 0, 46, 162,
		1, 0, 0, 0, 48, 165, 1, 0, 0, 0, 50, 169, 1, 0, 0, 0, 52, 172, 1, 0, 0,
		0, 54, 174, 1, 0, 0, 0, 56, 177, 1, 0, 0, 0, 58, 179, 1, 0, 0, 0, 60, 181,
		1, 0, 0, 0, 62, 183, 1, 0, 0, 0, 64, 185, 1, 0, 0, 0, 66, 187, 1, 0, 0,
		0, 68, 192, 1, 0, 0, 0, 70, 198, 1, 0, 0, 0, 72, 203, 1, 0, 0, 0, 74, 210,
		1, 0, 0, 0, 76, 214, 1, 0, 0, 0, 78, 217, 1, 0, 0, 0, 80, 220, 1, 0, 0,
		0, 82, 225, 1, 0, 0, 0, 84, 231, 1, 0, 0, 0, 86, 234, 1, 0, 0, 0, 88, 240,
		1, 0, 0, 0, 90, 249, 1, 0, 0, 0, 92, 273, 1, 0, 0, 0, 94, 275, 1, 0, 0,
		0, 96, 283, 1, 0, 0, 0, 98, 295, 1, 0, 0, 0, 100, 306, 1, 0, 0, 0, 102,
		320, 1, 0, 0, 0, 104, 329, 1, 0, 0, 0, 106, 333, 1, 0, 0, 0, 108, 336,
		1, 0, 0, 0, 110, 111, 5, 60, 0, 0, 111, 3, 1, 0, 0, 0, 112, 113, 5, 60,
		0, 0, 113, 114, 5, 61, 0, 0, 114, 5, 1, 0, 0, 0, 115, 116, 5, 61, 0, 0,
		116, 117, 5, 61, 0, 0, 117, 7, 1, 0, 0, 0, 118, 119, 5, 62, 0, 0, 119,
		9, 1, 0, 0, 0, 120, 121, 5, 62, 0, 0, 121, 122, 5, 61, 0, 0, 122, 11, 1,
		0, 0, 0, 123, 124, 5, 33, 0, 0, 124, 125, 5, 61, 0, 0, 125, 13, 1, 0, 0,
		0, 126, 127, 5, 38, 0, 0, 127, 128, 5, 38, 0, 0, 128, 15, 1, 0, 0, 0, 129,
		130, 5, 124, 0, 0, 130, 131, 5, 124, 0, 0, 131, 17, 1, 0, 0, 0, 132, 133,
		5, 33, 0, 0, 133, 19, 1, 0, 0, 0, 134, 135, 5, 43, 0, 0, 135, 21, 1, 0,
		0, 0, 136, 137, 5, 45, 0, 0, 137, 23, 1, 0, 0, 0, 138, 139, 5, 42, 0, 0,
		139, 25, 1, 0, 0, 0, 140, 141, 5, 47, 0, 0, 141, 27, 1, 0, 0, 0, 142, 143,
		5, 40, 0, 0, 143, 29, 1, 0, 0, 0, 144, 145, 5, 41, 0, 0, 145, 31, 1, 0,
		0, 0, 146, 147, 5, 91, 0, 0, 147, 33, 1, 0, 0, 0, 148, 149, 5, 93, 0, 0,
		149, 35, 1, 0, 0, 0, 150, 151, 5, 35, 0, 0, 151, 37, 1, 0, 0, 0, 152, 153,
		5, 63, 0, 0, 153, 39, 1, 0, 0, 0, 154, 155, 5, 43, 0, 0, 155, 156, 5, 61,
		0, 0, 156, 41, 1, 0, 0, 0, 157, 158, 5, 61, 0, 0, 158, 43, 1, 0, 0, 0,
		159, 160, 5, 63, 0, 0, 160, 161, 5, 63, 0, 0, 161, 45, 1, 0, 0, 0, 162,
		163, 5, 46, 0, 0, 163, 164, 5, 46, 0, 0, 164, 47, 1, 0, 0, 0, 165, 166,
		5, 46, 0, 0, 166, 167, 5, 46, 0, 0, 167, 168, 5, 46, 0, 0, 168, 49, 1,
		0, 0, 0, 169, 170, 5, 97, 0, 0, 170, 171, 5, 115, 0, 0, 171, 51, 1, 0,
		0, 0, 172, 173, 5, 44, 0, 0, 173, 53, 1, 0, 0, 0, 174, 175, 5, 61, 0, 0,
		175, 176, 5, 62, 0, 0, 176, 55, 1, 0, 0, 0, 177, 178, 5, 58, 0, 0, 178,
		57, 1, 0, 0, 0, 179, 180, 5, 59, 0, 0, 180, 59, 1, 0, 0, 0, 181, 182, 5,
		46, 0, 0, 182, 61, 1, 0, 0, 0, 183, 184, 5, 123, 0, 0, 184, 63, 1, 0, 0,
		0, 185, 186, 5, 125, 0, 0, 186, 65, 1, 0, 0, 0, 187, 188, 5, 110, 0, 0,
		188, 189, 5, 117, 0, 0, 189, 190, 5, 108, 0, 0, 190, 191, 5, 108, 0, 0,
		191, 67, 1, 0, 0, 0, 192, 193, 5, 102, 0, 0, 193, 194, 5, 97, 0, 0, 194,
		195, 5, 108, 0, 0, 195, 196, 5, 115, 0, 0, 196, 197, 5, 101, 0, 0, 197,
		69, 1, 0, 0, 0, 198, 199, 5, 116, 0, 0, 199, 200, 5, 114, 0, 0, 200, 201,
		5, 117, 0, 0, 201, 202, 5, 101, 0, 0, 202, 71, 1, 0, 0, 0, 203, 204, 5,
		114, 0, 0, 204, 205, 5, 101, 0, 0, 205, 206, 5, 116, 0, 0, 206, 207, 5,
		117, 0, 0, 207, 208, 5, 114, 0, 0, 208, 209, 5, 110, 0, 0, 209, 73, 1,
		0, 0, 0, 210, 211, 5, 102, 0, 0, 211, 212, 5, 111, 0, 0, 212, 213, 5, 114,
		0, 0, 213, 75, 1, 0, 0, 0, 214, 215, 5, 105, 0, 0, 215, 216, 5, 110, 0,
		0, 216, 77, 1, 0, 0, 0, 217, 218, 5, 105, 0, 0, 218, 219, 5, 102, 0, 0,
		219, 79, 1, 0, 0, 0, 220, 221, 5, 101, 0, 0, 221, 222, 5, 108, 0, 0, 222,
		223, 5, 115, 0, 0, 223, 224, 5, 101, 0, 0, 224, 81, 1, 0, 0, 0, 225, 226,
		5, 119, 0, 0, 226, 227, 5, 104, 0, 0, 227, 228, 5, 105, 0, 0, 228, 229,
		5, 108, 0, 0, 229, 230, 5, 101, 0, 0, 230, 83, 1, 0, 0, 0, 231, 232, 5,
		100, 0, 0, 232, 233, 5, 111, 0, 0, 233, 85, 1, 0, 0, 0, 234, 235, 5, 98,
		0, 0, 235, 236, 5, 114, 0, 0, 236, 237, 5, 101, 0, 0, 237, 238, 5, 97,
		0, 0, 238, 239, 5, 107, 0, 0, 239, 87, 1, 0, 0, 0, 240, 241, 5, 99, 0,
		0, 241, 242, 5, 111, 0, 0, 242, 243, 5, 110, 0, 0, 243, 244, 5, 116, 0,
		0, 244, 245, 5, 105, 0, 0, 245, 246, 5, 110, 0, 0, 246, 247, 5, 117, 0,
		0, 247, 248, 5, 101, 0, 0, 248, 89, 1, 0, 0, 0, 249, 250, 5, 96, 0, 0,
		250, 251, 1, 0, 0, 0, 251, 252, 6, 44, 0, 0, 252, 91, 1, 0, 0, 0, 253,
		259, 5, 34, 0, 0, 254, 255, 5, 92, 0, 0, 255, 258, 9, 0, 0, 0, 256, 258,
		8, 0, 0, 0, 257, 254, 1, 0, 0, 0, 257, 256, 1, 0, 0, 0, 258, 261, 1, 0,
		0, 0, 259, 257, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 262, 1, 0, 0, 0,
		261, 259, 1, 0, 0, 0, 262, 274, 5, 34, 0, 0, 263, 269, 5, 39, 0, 0, 264,
		265, 5, 92, 0, 0, 265, 268, 9, 0, 0, 0, 266, 268, 8, 1, 0, 0, 267, 264,
		1, 0, 0, 0, 267, 266, 1, 0, 0, 0, 268, 271, 1, 0, 0, 0, 269, 267, 1, 0,
		0, 0, 269, 270, 1, 0, 0, 0, 270, 272, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0,
		272, 274, 5, 39, 0, 0, 273, 253, 1, 0, 0, 0, 273, 263, 1, 0, 0, 0, 274,
		93, 1, 0, 0, 0, 275, 279, 7, 2, 0, 0, 276, 278, 7, 3, 0, 0, 277, 276, 1,
		0, 0, 0, 278, 281, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0, 279, 280, 1, 0, 0,
		0, 280, 95, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 282, 284, 7, 4, 0, 0, 283,
		282, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 285, 286,
		1, 0, 0, 0, 286, 293, 1, 0, 0, 0, 287, 289, 5, 46, 0, 0, 288, 290, 7, 4,
		0, 0, 289, 288, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 289, 1, 0, 0, 0,
		291, 292, 1, 0, 0, 0, 292, 294, 1, 0, 0, 0, 293, 287, 1, 0, 0, 0, 293,
		294, 1, 0, 0, 0, 294, 97, 1, 0, 0, 0, 295, 296, 5, 47, 0, 0, 296, 297,
		5, 47, 0, 0, 297, 301, 1, 0, 0, 0, 298, 300, 8, 5, 0, 0, 299, 298, 1, 0,
		0, 0, 300, 303, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0,
		302, 304, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 304, 305, 6, 48, 1, 0, 305,
		99, 1, 0, 0, 0, 306, 307, 5, 47, 0, 0, 307, 308, 5, 42, 0, 0, 308, 312,
		1, 0, 0, 0, 309, 311, 9, 0, 0, 0, 310, 309, 1, 0, 0, 0, 311, 314, 1, 0,
		0, 0, 312, 313, 1, 0, 0, 0, 312, 310, 1, 0, 0, 0, 313, 315, 1, 0, 0, 0,
		314, 312, 1, 0, 0, 0, 315, 316, 5, 42, 0, 0, 316, 317, 5, 47, 0, 0, 317,
		318, 1, 0, 0, 0, 318, 319, 6, 49, 1, 0, 319, 101, 1, 0, 0, 0, 320, 321,
		7, 6, 0, 0, 321, 322, 1, 0, 0, 0, 322, 323, 6, 50, 1, 0, 323, 103, 1, 0,
		0, 0, 324, 330, 8, 7, 0, 0, 325, 326, 5, 92, 0, 0, 326, 330, 9, 0, 0, 0,
		327, 328, 5, 36, 0, 0, 328, 330, 8, 8, 0, 0, 329, 324, 1, 0, 0, 0, 329,
		325, 1, 0, 0, 0, 329, 327, 1, 0, 0, 0, 330, 331, 1, 0, 0, 0, 331, 329,
		1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332, 105, 1, 0, 0, 0, 333, 334, 5, 36,
		0, 0, 334, 335, 5, 123, 0, 0, 335, 107, 1, 0, 0, 0, 336, 337, 5, 96, 0,
		0, 337, 338, 1, 0, 0, 0, 338, 339, 6, 53, 2, 0, 339, 109, 1, 0, 0, 0, 15,
		0, 1, 257, 259, 267, 269, 273, 279, 285, 291, 293, 301, 312, 329, 331,
		3, 5, 1, 0, 0, 1, 0, 4, 0, 0,
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
	l.GrammarFileName = "JsonTemplateLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JsonTemplateLexer tokens.
const (
	JsonTemplateLexerLess                  = 1
	JsonTemplateLexerLessOrEqual           = 2
	JsonTemplateLexerEqual                 = 3
	JsonTemplateLexerGreater               = 4
	JsonTemplateLexerGreaterOrEqual        = 5
	JsonTemplateLexerNotEqual              = 6
	JsonTemplateLexerAnd                   = 7
	JsonTemplateLexerOr                    = 8
	JsonTemplateLexerNot                   = 9
	JsonTemplateLexerAdd                   = 10
	JsonTemplateLexerSubtract              = 11
	JsonTemplateLexerMultiply              = 12
	JsonTemplateLexerDivide                = 13
	JsonTemplateLexerLeftParen             = 14
	JsonTemplateLexerRightParen            = 15
	JsonTemplateLexerLeftBracket           = 16
	JsonTemplateLexerRightBracket          = 17
	JsonTemplateLexerIteration             = 18
	JsonTemplateLexerQuestion              = 19
	JsonTemplateLexerAddAssign             = 20
	JsonTemplateLexerLiteral               = 21
	JsonTemplateLexerNullCoalescing        = 22
	JsonTemplateLexerRange                 = 23
	JsonTemplateLexerSpread                = 24
	JsonTemplateLexerAs                    = 25
	JsonTemplateLexerComma                 = 26
	JsonTemplateLexerArrow                 = 27
	JsonTemplateLexerColon                 = 28
	JsonTemplateLexerSemicolon             = 29
	JsonTemplateLexerDot                   = 30
	JsonTemplateLexerLeftBrace             = 31
	JsonTemplateLexerRightBrace            = 32
	JsonTemplateLexerNull                  = 33
	JsonTemplateLexerFalse                 = 34
	JsonTemplateLexerTrue                  = 35
	JsonTemplateLexerReturn                = 36
	JsonTemplateLexerFor                   = 37
	JsonTemplateLexerIn                    = 38
	JsonTemplateLexerIf                    = 39
	JsonTemplateLexerElse                  = 40
	JsonTemplateLexerWhile                 = 41
	JsonTemplateLexerDo                    = 42
	JsonTemplateLexerBreak                 = 43
	JsonTemplateLexerContinue              = 44
	JsonTemplateLexerBacktick              = 45
	JsonTemplateLexerESCAPED_STRING        = 46
	JsonTemplateLexerSTRING                = 47
	JsonTemplateLexerNUMBER                = 48
	JsonTemplateLexerLINE_COMMENT          = 49
	JsonTemplateLexerBLOCK_COMMENT         = 50
	JsonTemplateLexerWS                    = 51
	JsonTemplateLexerTEMPLATE_TEXT         = 52
	JsonTemplateLexerTEMPLATE_INTERP_START = 53
	JsonTemplateLexerTEMPLATE_END          = 54
)

// JsonTemplateLexerTEMPLATE_MODE is the JsonTemplateLexer mode.
const JsonTemplateLexerTEMPLATE_MODE = 1
