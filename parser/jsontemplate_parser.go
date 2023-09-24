// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // JsonTemplate

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type JsonTemplateParser struct {
	*antlr.BaseParser
}

var jsontemplateParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jsontemplateParserInit() {
	staticData := &jsontemplateParserStaticData
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
		"script", "statement", "statements", "expression", "lambda", "function_param",
		"field", "array", "spread_field", "object", "object_field", "name",
		"index",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 45, 328, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 5, 0, 28, 8, 0, 10, 0, 12, 0, 31,
		9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		5, 1, 44, 8, 1, 10, 1, 12, 1, 47, 9, 1, 3, 1, 49, 8, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 63, 8, 1,
		10, 1, 12, 1, 66, 9, 1, 1, 1, 1, 1, 3, 1, 70, 8, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 82, 8, 1, 1, 2, 1, 2, 5, 2,
		86, 8, 2, 10, 2, 12, 2, 89, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 95, 8,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 103, 8, 3, 3, 3, 105, 8, 3,
		1, 3, 1, 3, 3, 3, 109, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 115, 8, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 3, 3, 121, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 127,
		8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 133, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		3, 3, 139, 8, 3, 1, 3, 1, 3, 1, 3, 3, 3, 144, 8, 3, 1, 3, 1, 3, 3, 3, 148,
		8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 158, 8, 4,
		10, 4, 12, 4, 161, 9, 4, 5, 4, 163, 8, 4, 10, 4, 12, 4, 166, 9, 4, 1, 4,
		1, 4, 1, 4, 3, 4, 171, 8, 4, 1, 5, 1, 5, 3, 5, 175, 8, 5, 1, 5, 3, 5, 178,
		8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 197, 8, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 3, 6, 240, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 246, 8, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 257, 8, 6, 10,
		6, 12, 6, 260, 9, 6, 3, 6, 262, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 3, 6, 270, 8, 6, 5, 6, 272, 8, 6, 10, 6, 12, 6, 275, 9, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 5, 7, 281, 8, 7, 10, 7, 12, 7, 284, 9, 7, 3, 7, 286, 8,
		7, 1, 7, 1, 7, 1, 8, 3, 8, 291, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1,
		9, 5, 9, 299, 8, 9, 10, 9, 12, 9, 302, 9, 9, 3, 9, 304, 8, 9, 1, 9, 1,
		9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 316,
		8, 10, 1, 10, 3, 10, 319, 8, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 3,
		12, 326, 8, 12, 1, 12, 0, 1, 12, 13, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 0, 2, 1, 0, 22, 23, 1, 0, 20, 21, 385, 0, 29, 1, 0, 0, 0, 2,
		81, 1, 0, 0, 0, 4, 83, 1, 0, 0, 0, 6, 147, 1, 0, 0, 0, 8, 170, 1, 0, 0,
		0, 10, 177, 1, 0, 0, 0, 12, 196, 1, 0, 0, 0, 14, 276, 1, 0, 0, 0, 16, 290,
		1, 0, 0, 0, 18, 294, 1, 0, 0, 0, 20, 318, 1, 0, 0, 0, 22, 320, 1, 0, 0,
		0, 24, 325, 1, 0, 0, 0, 26, 28, 3, 2, 1, 0, 27, 26, 1, 0, 0, 0, 28, 31,
		1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 1, 1, 0, 0, 0,
		31, 29, 1, 0, 0, 0, 32, 33, 3, 12, 6, 0, 33, 34, 5, 1, 0, 0, 34, 82, 1,
		0, 0, 0, 35, 36, 5, 2, 0, 0, 36, 37, 3, 12, 6, 0, 37, 38, 5, 1, 0, 0, 38,
		82, 1, 0, 0, 0, 39, 48, 5, 3, 0, 0, 40, 45, 3, 22, 11, 0, 41, 42, 5, 35,
		0, 0, 42, 44, 3, 22, 11, 0, 43, 41, 1, 0, 0, 0, 44, 47, 1, 0, 0, 0, 45,
		43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 49, 1, 0, 0, 0, 47, 45, 1, 0, 0,
		0, 48, 40, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51,
		5, 4, 0, 0, 51, 52, 3, 12, 6, 0, 52, 53, 3, 4, 2, 0, 53, 82, 1, 0, 0, 0,
		54, 55, 5, 5, 0, 0, 55, 56, 3, 12, 6, 0, 56, 64, 3, 4, 2, 0, 57, 58, 5,
		6, 0, 0, 58, 59, 5, 5, 0, 0, 59, 60, 3, 12, 6, 0, 60, 61, 3, 4, 2, 0, 61,
		63, 1, 0, 0, 0, 62, 57, 1, 0, 0, 0, 63, 66, 1, 0, 0, 0, 64, 62, 1, 0, 0,
		0, 64, 65, 1, 0, 0, 0, 65, 69, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 67, 68,
		5, 6, 0, 0, 68, 70, 3, 4, 2, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0,
		70, 82, 1, 0, 0, 0, 71, 72, 5, 7, 0, 0, 72, 73, 3, 12, 6, 0, 73, 74, 3,
		4, 2, 0, 74, 82, 1, 0, 0, 0, 75, 76, 5, 8, 0, 0, 76, 77, 3, 4, 2, 0, 77,
		78, 5, 7, 0, 0, 78, 79, 3, 12, 6, 0, 79, 80, 5, 1, 0, 0, 80, 82, 1, 0,
		0, 0, 81, 32, 1, 0, 0, 0, 81, 35, 1, 0, 0, 0, 81, 39, 1, 0, 0, 0, 81, 54,
		1, 0, 0, 0, 81, 71, 1, 0, 0, 0, 81, 75, 1, 0, 0, 0, 82, 3, 1, 0, 0, 0,
		83, 87, 5, 37, 0, 0, 84, 86, 3, 2, 1, 0, 85, 84, 1, 0, 0, 0, 86, 89, 1,
		0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 90, 1, 0, 0, 0, 89,
		87, 1, 0, 0, 0, 90, 91, 5, 38, 0, 0, 91, 5, 1, 0, 0, 0, 92, 93, 5, 37,
		0, 0, 93, 95, 5, 37, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95,
		96, 1, 0, 0, 0, 96, 97, 5, 28, 0, 0, 97, 104, 3, 12, 6, 0, 98, 99, 5, 34,
		0, 0, 99, 102, 3, 22, 11, 0, 100, 101, 5, 35, 0, 0, 101, 103, 3, 22, 11,
		0, 102, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 105, 1, 0, 0, 0, 104,
		98, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 108, 1, 0, 0, 0, 106, 107, 5,
		38, 0, 0, 107, 109, 5, 38, 0, 0, 108, 106, 1, 0, 0, 0, 108, 109, 1, 0,
		0, 0, 109, 110, 1, 0, 0, 0, 110, 111, 5, 0, 0, 1, 111, 148, 1, 0, 0, 0,
		112, 113, 5, 37, 0, 0, 113, 115, 5, 37, 0, 0, 114, 112, 1, 0, 0, 0, 114,
		115, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 117, 5, 29, 0, 0, 117, 120,
		3, 12, 6, 0, 118, 119, 5, 38, 0, 0, 119, 121, 5, 38, 0, 0, 120, 118, 1,
		0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 123, 5, 0, 0,
		1, 123, 148, 1, 0, 0, 0, 124, 125, 5, 37, 0, 0, 125, 127, 5, 37, 0, 0,
		126, 124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128,
		129, 5, 30, 0, 0, 129, 132, 3, 12, 6, 0, 130, 131, 5, 38, 0, 0, 131, 133,
		5, 38, 0, 0, 132, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 1, 0,
		0, 0, 134, 135, 5, 0, 0, 1, 135, 148, 1, 0, 0, 0, 136, 137, 5, 37, 0, 0,
		137, 139, 5, 37, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139,
		140, 1, 0, 0, 0, 140, 143, 3, 12, 6, 0, 141, 142, 5, 38, 0, 0, 142, 144,
		5, 38, 0, 0, 143, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 145, 1, 0,
		0, 0, 145, 146, 5, 0, 0, 1, 146, 148, 1, 0, 0, 0, 147, 94, 1, 0, 0, 0,
		147, 114, 1, 0, 0, 0, 147, 126, 1, 0, 0, 0, 147, 138, 1, 0, 0, 0, 148,
		7, 1, 0, 0, 0, 149, 150, 3, 22, 11, 0, 150, 151, 5, 36, 0, 0, 151, 152,
		3, 12, 6, 0, 152, 171, 1, 0, 0, 0, 153, 164, 5, 24, 0, 0, 154, 159, 3,
		22, 11, 0, 155, 156, 5, 35, 0, 0, 156, 158, 3, 22, 11, 0, 157, 155, 1,
		0, 0, 0, 158, 161, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159, 160, 1, 0, 0,
		0, 160, 163, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 162, 154, 1, 0, 0, 0, 163,
		166, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 167,
		1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 167, 168, 5, 25, 0, 0, 168, 169, 5, 36,
		0, 0, 169, 171, 3, 12, 6, 0, 170, 149, 1, 0, 0, 0, 170, 153, 1, 0, 0, 0,
		171, 9, 1, 0, 0, 0, 172, 174, 3, 12, 6, 0, 173, 175, 5, 33, 0, 0, 174,
		173, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 178, 1, 0, 0, 0, 176, 178,
		3, 8, 4, 0, 177, 172, 1, 0, 0, 0, 177, 176, 1, 0, 0, 0, 178, 11, 1, 0,
		0, 0, 179, 180, 6, 6, -1, 0, 180, 181, 5, 24, 0, 0, 181, 182, 3, 12, 6,
		0, 182, 183, 5, 25, 0, 0, 183, 197, 1, 0, 0, 0, 184, 197, 5, 41, 0, 0,
		185, 197, 5, 40, 0, 0, 186, 197, 5, 39, 0, 0, 187, 197, 5, 44, 0, 0, 188,
		197, 5, 42, 0, 0, 189, 197, 3, 14, 7, 0, 190, 197, 3, 18, 9, 0, 191, 197,
		3, 22, 11, 0, 192, 193, 5, 21, 0, 0, 193, 197, 3, 12, 6, 16, 194, 195,
		5, 19, 0, 0, 195, 197, 3, 12, 6, 15, 196, 179, 1, 0, 0, 0, 196, 184, 1,
		0, 0, 0, 196, 185, 1, 0, 0, 0, 196, 186, 1, 0, 0, 0, 196, 187, 1, 0, 0,
		0, 196, 188, 1, 0, 0, 0, 196, 189, 1, 0, 0, 0, 196, 190, 1, 0, 0, 0, 196,
		191, 1, 0, 0, 0, 196, 192, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 197, 273,
		1, 0, 0, 0, 198, 199, 10, 14, 0, 0, 199, 200, 7, 0, 0, 0, 200, 272, 3,
		12, 6, 15, 201, 202, 10, 13, 0, 0, 202, 203, 7, 1, 0, 0, 203, 272, 3, 12,
		6, 14, 204, 205, 10, 12, 0, 0, 205, 206, 5, 32, 0, 0, 206, 272, 3, 12,
		6, 13, 207, 208, 10, 11, 0, 0, 208, 209, 5, 31, 0, 0, 209, 272, 3, 12,
		6, 12, 210, 211, 10, 10, 0, 0, 211, 212, 5, 13, 0, 0, 212, 272, 3, 12,
		6, 11, 213, 214, 10, 9, 0, 0, 214, 215, 5, 11, 0, 0, 215, 272, 3, 12, 6,
		10, 216, 217, 10, 8, 0, 0, 217, 218, 5, 12, 0, 0, 218, 272, 3, 12, 6, 9,
		219, 220, 10, 7, 0, 0, 220, 221, 5, 14, 0, 0, 221, 272, 3, 12, 6, 8, 222,
		223, 10, 6, 0, 0, 223, 224, 5, 15, 0, 0, 224, 272, 3, 12, 6, 7, 225, 226,
		10, 5, 0, 0, 226, 227, 5, 16, 0, 0, 227, 272, 3, 12, 6, 6, 228, 229, 10,
		4, 0, 0, 229, 230, 5, 17, 0, 0, 230, 272, 3, 12, 6, 5, 231, 232, 10, 3,
		0, 0, 232, 233, 5, 18, 0, 0, 233, 272, 3, 12, 6, 4, 234, 235, 10, 1, 0,
		0, 235, 236, 5, 30, 0, 0, 236, 272, 3, 12, 6, 2, 237, 239, 10, 19, 0, 0,
		238, 240, 5, 29, 0, 0, 239, 238, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240,
		241, 1, 0, 0, 0, 241, 242, 5, 9, 0, 0, 242, 272, 3, 22, 11, 0, 243, 245,
		10, 18, 0, 0, 244, 246, 5, 29, 0, 0, 245, 244, 1, 0, 0, 0, 245, 246, 1,
		0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 248, 5, 26, 0, 0, 248, 249, 3, 24,
		12, 0, 249, 250, 5, 27, 0, 0, 250, 272, 1, 0, 0, 0, 251, 252, 10, 17, 0,
		0, 252, 261, 5, 24, 0, 0, 253, 258, 3, 10, 5, 0, 254, 255, 5, 35, 0, 0,
		255, 257, 3, 10, 5, 0, 256, 254, 1, 0, 0, 0, 257, 260, 1, 0, 0, 0, 258,
		256, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 262, 1, 0, 0, 0, 260, 258,
		1, 0, 0, 0, 261, 253, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 263, 1, 0,
		0, 0, 263, 272, 5, 25, 0, 0, 264, 265, 10, 2, 0, 0, 265, 266, 5, 29, 0,
		0, 266, 269, 3, 12, 6, 0, 267, 268, 5, 10, 0, 0, 268, 270, 3, 12, 6, 0,
		269, 267, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 272, 1, 0, 0, 0, 271,
		198, 1, 0, 0, 0, 271, 201, 1, 0, 0, 0, 271, 204, 1, 0, 0, 0, 271, 207,
		1, 0, 0, 0, 271, 210, 1, 0, 0, 0, 271, 213, 1, 0, 0, 0, 271, 216, 1, 0,
		0, 0, 271, 219, 1, 0, 0, 0, 271, 222, 1, 0, 0, 0, 271, 225, 1, 0, 0, 0,
		271, 228, 1, 0, 0, 0, 271, 231, 1, 0, 0, 0, 271, 234, 1, 0, 0, 0, 271,
		237, 1, 0, 0, 0, 271, 243, 1, 0, 0, 0, 271, 251, 1, 0, 0, 0, 271, 264,
		1, 0, 0, 0, 272, 275, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 273, 274, 1, 0,
		0, 0, 274, 13, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 276, 285, 5, 26, 0, 0,
		277, 282, 3, 16, 8, 0, 278, 279, 5, 35, 0, 0, 279, 281, 3, 16, 8, 0, 280,
		278, 1, 0, 0, 0, 281, 284, 1, 0, 0, 0, 282, 280, 1, 0, 0, 0, 282, 283,
		1, 0, 0, 0, 283, 286, 1, 0, 0, 0, 284, 282, 1, 0, 0, 0, 285, 277, 1, 0,
		0, 0, 285, 286, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 288, 5, 27, 0, 0,
		288, 15, 1, 0, 0, 0, 289, 291, 5, 33, 0, 0, 290, 289, 1, 0, 0, 0, 290,
		291, 1, 0, 0, 0, 291, 292, 1, 0, 0, 0, 292, 293, 3, 12, 6, 0, 293, 17,
		1, 0, 0, 0, 294, 303, 5, 37, 0, 0, 295, 300, 3, 20, 10, 0, 296, 297, 5,
		35, 0, 0, 297, 299, 3, 20, 10, 0, 298, 296, 1, 0, 0, 0, 299, 302, 1, 0,
		0, 0, 300, 298, 1, 0, 0, 0, 300, 301, 1, 0, 0, 0, 301, 304, 1, 0, 0, 0,
		302, 300, 1, 0, 0, 0, 303, 295, 1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304,
		305, 1, 0, 0, 0, 305, 306, 5, 38, 0, 0, 306, 19, 1, 0, 0, 0, 307, 308,
		3, 22, 11, 0, 308, 309, 5, 10, 0, 0, 309, 310, 3, 12, 6, 0, 310, 319, 1,
		0, 0, 0, 311, 312, 5, 42, 0, 0, 312, 313, 5, 10, 0, 0, 313, 319, 3, 12,
		6, 0, 314, 316, 5, 33, 0, 0, 315, 314, 1, 0, 0, 0, 315, 316, 1, 0, 0, 0,
		316, 317, 1, 0, 0, 0, 317, 319, 3, 12, 6, 0, 318, 307, 1, 0, 0, 0, 318,
		311, 1, 0, 0, 0, 318, 315, 1, 0, 0, 0, 319, 21, 1, 0, 0, 0, 320, 321, 5,
		43, 0, 0, 321, 23, 1, 0, 0, 0, 322, 326, 3, 12, 6, 0, 323, 326, 5, 44,
		0, 0, 324, 326, 5, 42, 0, 0, 325, 322, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0,
		325, 324, 1, 0, 0, 0, 326, 25, 1, 0, 0, 0, 39, 29, 45, 48, 64, 69, 81,
		87, 94, 102, 104, 108, 114, 120, 126, 132, 138, 143, 147, 159, 164, 170,
		174, 177, 196, 239, 245, 258, 261, 269, 271, 273, 282, 285, 290, 300, 303,
		315, 318, 325,
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

// JsonTemplateParserInit initializes any static state used to implement JsonTemplateParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJsonTemplateParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JsonTemplateParserInit() {
	staticData := &jsontemplateParserStaticData
	staticData.once.Do(jsontemplateParserInit)
}

// NewJsonTemplateParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJsonTemplateParser(input antlr.TokenStream) *JsonTemplateParser {
	JsonTemplateParserInit()
	this := new(JsonTemplateParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &jsontemplateParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// JsonTemplateParser tokens.
const (
	JsonTemplateParserEOF            = antlr.TokenEOF
	JsonTemplateParserT__0           = 1
	JsonTemplateParserT__1           = 2
	JsonTemplateParserT__2           = 3
	JsonTemplateParserT__3           = 4
	JsonTemplateParserT__4           = 5
	JsonTemplateParserT__5           = 6
	JsonTemplateParserT__6           = 7
	JsonTemplateParserT__7           = 8
	JsonTemplateParserT__8           = 9
	JsonTemplateParserT__9           = 10
	JsonTemplateParserLess           = 11
	JsonTemplateParserLessOrEqual    = 12
	JsonTemplateParserEqual          = 13
	JsonTemplateParserGreater        = 14
	JsonTemplateParserGreaterOrEqual = 15
	JsonTemplateParserNotEqual       = 16
	JsonTemplateParserAnd            = 17
	JsonTemplateParserOr             = 18
	JsonTemplateParserNot            = 19
	JsonTemplateParserAdd            = 20
	JsonTemplateParserSubtract       = 21
	JsonTemplateParserMultiply       = 22
	JsonTemplateParserDivide         = 23
	JsonTemplateParserLeftParen      = 24
	JsonTemplateParserRightParen     = 25
	JsonTemplateParserLeftBracket    = 26
	JsonTemplateParserRightBracket   = 27
	JsonTemplateParserIteration      = 28
	JsonTemplateParserQuestion       = 29
	JsonTemplateParserLiteral        = 30
	JsonTemplateParserNullCoalescing = 31
	JsonTemplateParserRange          = 32
	JsonTemplateParserSpread         = 33
	JsonTemplateParserAs             = 34
	JsonTemplateParserComma          = 35
	JsonTemplateParserArrow          = 36
	JsonTemplateParserLeftBrace      = 37
	JsonTemplateParserRightBrace     = 38
	JsonTemplateParserNull           = 39
	JsonTemplateParserFalse          = 40
	JsonTemplateParserTrue           = 41
	JsonTemplateParserESCAPED_STRING = 42
	JsonTemplateParserSTRING         = 43
	JsonTemplateParserNUMBER         = 44
	JsonTemplateParserWS             = 45
)

// JsonTemplateParser rules.
const (
	JsonTemplateParserRULE_script         = 0
	JsonTemplateParserRULE_statement      = 1
	JsonTemplateParserRULE_statements     = 2
	JsonTemplateParserRULE_expression     = 3
	JsonTemplateParserRULE_lambda         = 4
	JsonTemplateParserRULE_function_param = 5
	JsonTemplateParserRULE_field          = 6
	JsonTemplateParserRULE_array          = 7
	JsonTemplateParserRULE_spread_field   = 8
	JsonTemplateParserRULE_object         = 9
	JsonTemplateParserRULE_object_field   = 10
	JsonTemplateParserRULE_name           = 11
	JsonTemplateParserRULE_index          = 12
)

// IScriptContext is an interface to support dynamic dispatch.
type IScriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScriptContext differentiates from other interfaces.
	IsScriptContext()
}

type ScriptContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScriptContext() *ScriptContext {
	var p = new(ScriptContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonTemplateParserRULE_script
	return p
}

func (*ScriptContext) IsScriptContext() {}

func NewScriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScriptContext {
	var p = new(ScriptContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateParserRULE_script

	return p
}

func (s *ScriptContext) GetParser() antlr.Parser { return s.parser }

func (s *ScriptContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ScriptContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ScriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScriptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.EnterScript(s)
	}
}

func (s *ScriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.ExitScript(s)
	}
}

func (s *ScriptContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitScript(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonTemplateParser) Script() (localctx IScriptContext) {
	this := p
	_ = this

	localctx = NewScriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JsonTemplateParserRULE_script)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34772141736364) != 0 {
		{
			p.SetState(26)
			p.Statement()
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonTemplateParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *StatementContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *StatementContext) AllStatements() []IStatementsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementsContext); ok {
			len++
		}
	}

	tst := make([]IStatementsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementsContext); ok {
			tst[i] = t.(IStatementsContext)
			i++
		}
	}

	return tst
}

func (s *StatementContext) Statements(i int) IStatementsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *StatementContext) AllName() []INameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameContext); ok {
			len++
		}
	}

	tst := make([]INameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameContext); ok {
			tst[i] = t.(INameContext)
			i++
		}
	}

	return tst
}

func (s *StatementContext) Name(i int) INameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *StatementContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(JsonTemplateParserComma)
}

func (s *StatementContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserComma, i)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonTemplateParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JsonTemplateParserRULE_statement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(81)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonTemplateParserNot, JsonTemplateParserSubtract, JsonTemplateParserLeftParen, JsonTemplateParserLeftBracket, JsonTemplateParserLeftBrace, JsonTemplateParserNull, JsonTemplateParserFalse, JsonTemplateParserTrue, JsonTemplateParserESCAPED_STRING, JsonTemplateParserSTRING, JsonTemplateParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(32)
			p.field(0)
		}
		{
			p.SetState(33)
			p.Match(JsonTemplateParserT__0)
		}

	case JsonTemplateParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.Match(JsonTemplateParserT__1)
		}
		{
			p.SetState(36)
			p.field(0)
		}
		{
			p.SetState(37)
			p.Match(JsonTemplateParserT__0)
		}

	case JsonTemplateParserT__2:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(39)
			p.Match(JsonTemplateParserT__2)
		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserSTRING {
			{
				p.SetState(40)
				p.Name()
			}
			p.SetState(45)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == JsonTemplateParserComma {
				{
					p.SetState(41)
					p.Match(JsonTemplateParserComma)
				}
				{
					p.SetState(42)
					p.Name()
				}

				p.SetState(47)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(50)
			p.Match(JsonTemplateParserT__3)
		}
		{
			p.SetState(51)
			p.field(0)
		}
		{
			p.SetState(52)
			p.Statements()
		}

	case JsonTemplateParserT__4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(54)
			p.Match(JsonTemplateParserT__4)
		}
		{
			p.SetState(55)
			p.field(0)
		}
		{
			p.SetState(56)
			p.Statements()
		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(57)
					p.Match(JsonTemplateParserT__5)
				}
				{
					p.SetState(58)
					p.Match(JsonTemplateParserT__4)
				}
				{
					p.SetState(59)
					p.field(0)
				}
				{
					p.SetState(60)
					p.Statements()
				}

			}
			p.SetState(66)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserT__5 {
			{
				p.SetState(67)
				p.Match(JsonTemplateParserT__5)
			}
			{
				p.SetState(68)
				p.Statements()
			}

		}

	case JsonTemplateParserT__6:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(71)
			p.Match(JsonTemplateParserT__6)
		}
		{
			p.SetState(72)
			p.field(0)
		}
		{
			p.SetState(73)
			p.Statements()
		}

	case JsonTemplateParserT__7:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(75)
			p.Match(JsonTemplateParserT__7)
		}
		{
			p.SetState(76)
			p.Statements()
		}
		{
			p.SetState(77)
			p.Match(JsonTemplateParserT__6)
		}
		{
			p.SetState(78)
			p.field(0)
		}
		{
			p.SetState(79)
			p.Match(JsonTemplateParserT__0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonTemplateParserRULE_statements
	return p
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserLeftBrace, 0)
}

func (s *StatementsContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserRightBrace, 0)
}

func (s *StatementsContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementsContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (s *StatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonTemplateParser) Statements() (localctx IStatementsContext) {
	this := p
	_ = this

	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JsonTemplateParserRULE_statements)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Match(JsonTemplateParserLeftBrace)
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34772141736364) != 0 {
		{
			p.SetState(84)
			p.Statement()
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
		p.Match(JsonTemplateParserRightBrace)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonTemplateParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Iteration() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserIteration, 0)
}

func (s *ExpressionContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserEOF, 0)
}

func (s *ExpressionContext) AllLeftBrace() []antlr.TerminalNode {
	return s.GetTokens(JsonTemplateParserLeftBrace)
}

func (s *ExpressionContext) LeftBrace(i int) antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserLeftBrace, i)
}

func (s *ExpressionContext) As() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserAs, 0)
}

func (s *ExpressionContext) AllName() []INameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameContext); ok {
			len++
		}
	}

	tst := make([]INameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameContext); ok {
			tst[i] = t.(INameContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Name(i int) INameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ExpressionContext) AllRightBrace() []antlr.TerminalNode {
	return s.GetTokens(JsonTemplateParserRightBrace)
}

func (s *ExpressionContext) RightBrace(i int) antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserRightBrace, i)
}

func (s *ExpressionContext) Comma() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserComma, 0)
}

func (s *ExpressionContext) Question() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserQuestion, 0)
}

func (s *ExpressionContext) Literal() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserLiteral, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonTemplateParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JsonTemplateParserRULE_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserLeftBrace {
			{
				p.SetState(92)
				p.Match(JsonTemplateParserLeftBrace)
			}
			{
				p.SetState(93)
				p.Match(JsonTemplateParserLeftBrace)
			}

		}
		{
			p.SetState(96)
			p.Match(JsonTemplateParserIteration)
		}
		{
			p.SetState(97)
			p.field(0)
		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserAs {
			{
				p.SetState(98)
				p.Match(JsonTemplateParserAs)
			}
			{
				p.SetState(99)
				p.Name()
			}
			p.SetState(102)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == JsonTemplateParserComma {
				{
					p.SetState(100)
					p.Match(JsonTemplateParserComma)
				}
				{
					p.SetState(101)
					p.Name()
				}

			}

		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserRightBrace {
			{
				p.SetState(106)
				p.Match(JsonTemplateParserRightBrace)
			}
			{
				p.SetState(107)
				p.Match(JsonTemplateParserRightBrace)
			}

		}
		{
			p.SetState(110)
			p.Match(JsonTemplateParserEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserLeftBrace {
			{
				p.SetState(112)
				p.Match(JsonTemplateParserLeftBrace)
			}
			{
				p.SetState(113)
				p.Match(JsonTemplateParserLeftBrace)
			}

		}
		{
			p.SetState(116)
			p.Match(JsonTemplateParserQuestion)
		}
		{
			p.SetState(117)
			p.field(0)
		}
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserRightBrace {
			{
				p.SetState(118)
				p.Match(JsonTemplateParserRightBrace)
			}
			{
				p.SetState(119)
				p.Match(JsonTemplateParserRightBrace)
			}

		}
		{
			p.SetState(122)
			p.Match(JsonTemplateParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserLeftBrace {
			{
				p.SetState(124)
				p.Match(JsonTemplateParserLeftBrace)
			}
			{
				p.SetState(125)
				p.Match(JsonTemplateParserLeftBrace)
			}

		}
		{
			p.SetState(128)
			p.Match(JsonTemplateParserLiteral)
		}
		{
			p.SetState(129)
			p.field(0)
		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserRightBrace {
			{
				p.SetState(130)
				p.Match(JsonTemplateParserRightBrace)
			}
			{
				p.SetState(131)
				p.Match(JsonTemplateParserRightBrace)
			}

		}
		{
			p.SetState(134)
			p.Match(JsonTemplateParserEOF)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(138)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(136)
				p.Match(JsonTemplateParserLeftBrace)
			}
			{
				p.SetState(137)
				p.Match(JsonTemplateParserLeftBrace)
			}

		}
		{
			p.SetState(140)
			p.field(0)
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserRightBrace {
			{
				p.SetState(141)
				p.Match(JsonTemplateParserRightBrace)
			}
			{
				p.SetState(142)
				p.Match(JsonTemplateParserRightBrace)
			}

		}
		{
			p.SetState(145)
			p.Match(JsonTemplateParserEOF)
		}

	}

	return localctx
}

// ILambdaContext is an interface to support dynamic dispatch.
type ILambdaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLambdaContext differentiates from other interfaces.
	IsLambdaContext()
}

type LambdaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLambdaContext() *LambdaContext {
	var p = new(LambdaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonTemplateParserRULE_lambda
	return p
}

func (*LambdaContext) IsLambdaContext() {}

func NewLambdaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaContext {
	var p = new(LambdaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateParserRULE_lambda

	return p
}

func (s *LambdaContext) GetParser() antlr.Parser { return s.parser }

func (s *LambdaContext) AllName() []INameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameContext); ok {
			len++
		}
	}

	tst := make([]INameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameContext); ok {
			tst[i] = t.(INameContext)
			i++
		}
	}

	return tst
}

func (s *LambdaContext) Name(i int) INameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *LambdaContext) Arrow() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserArrow, 0)
}

func (s *LambdaContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *LambdaContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserLeftParen, 0)
}

func (s *LambdaContext) RightParen() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserRightParen, 0)
}

func (s *LambdaContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(JsonTemplateParserComma)
}

func (s *LambdaContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserComma, i)
}

func (s *LambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LambdaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.EnterLambda(s)
	}
}

func (s *LambdaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.ExitLambda(s)
	}
}

func (s *LambdaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitLambda(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonTemplateParser) Lambda() (localctx ILambdaContext) {
	this := p
	_ = this

	localctx = NewLambdaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JsonTemplateParserRULE_lambda)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(170)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonTemplateParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.Name()
		}
		{
			p.SetState(150)
			p.Match(JsonTemplateParserArrow)
		}
		{
			p.SetState(151)
			p.field(0)
		}

	case JsonTemplateParserLeftParen:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(153)
			p.Match(JsonTemplateParserLeftParen)
		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonTemplateParserSTRING {
			{
				p.SetState(154)
				p.Name()
			}
			p.SetState(159)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == JsonTemplateParserComma {
				{
					p.SetState(155)
					p.Match(JsonTemplateParserComma)
				}
				{
					p.SetState(156)
					p.Name()
				}

				p.SetState(161)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

			p.SetState(166)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(167)
			p.Match(JsonTemplateParserRightParen)
		}
		{
			p.SetState(168)
			p.Match(JsonTemplateParserArrow)
		}
		{
			p.SetState(169)
			p.field(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunction_paramContext is an interface to support dynamic dispatch.
type IFunction_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_paramContext differentiates from other interfaces.
	IsFunction_paramContext()
}

type Function_paramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_paramContext() *Function_paramContext {
	var p = new(Function_paramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonTemplateParserRULE_function_param
	return p
}

func (*Function_paramContext) IsFunction_paramContext() {}

func NewFunction_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_paramContext {
	var p = new(Function_paramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateParserRULE_function_param

	return p
}

func (s *Function_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_paramContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *Function_paramContext) Spread() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserSpread, 0)
}

func (s *Function_paramContext) Lambda() ILambdaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILambdaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILambdaContext)
}

func (s *Function_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.EnterFunction_param(s)
	}
}

func (s *Function_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.ExitFunction_param(s)
	}
}

func (s *Function_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitFunction_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonTemplateParser) Function_param() (localctx IFunction_paramContext) {
	this := p
	_ = this

	localctx = NewFunction_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JsonTemplateParserRULE_function_param)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(172)
			p.field(0)
		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserSpread {
			{
				p.SetState(173)
				p.Match(JsonTemplateParserSpread)
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(176)
			p.Lambda()
		}

	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonTemplateParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserLeftParen, 0)
}

func (s *FieldContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *FieldContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *FieldContext) RightParen() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserRightParen, 0)
}

func (s *FieldContext) True() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserTrue, 0)
}

func (s *FieldContext) False() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserFalse, 0)
}

func (s *FieldContext) Null() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserNull, 0)
}

func (s *FieldContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserNUMBER, 0)
}

func (s *FieldContext) ESCAPED_STRING() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserESCAPED_STRING, 0)
}

func (s *FieldContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *FieldContext) Object() IObjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *FieldContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FieldContext) Subtract() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserSubtract, 0)
}

func (s *FieldContext) Not() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserNot, 0)
}

func (s *FieldContext) Divide() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserDivide, 0)
}

func (s *FieldContext) Multiply() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserMultiply, 0)
}

func (s *FieldContext) Add() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserAdd, 0)
}

func (s *FieldContext) Range() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserRange, 0)
}

func (s *FieldContext) NullCoalescing() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserNullCoalescing, 0)
}

func (s *FieldContext) Equal() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserEqual, 0)
}

func (s *FieldContext) Less() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserLess, 0)
}

func (s *FieldContext) LessOrEqual() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserLessOrEqual, 0)
}

func (s *FieldContext) Greater() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserGreater, 0)
}

func (s *FieldContext) GreaterOrEqual() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserGreaterOrEqual, 0)
}

func (s *FieldContext) NotEqual() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserNotEqual, 0)
}

func (s *FieldContext) And() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserAnd, 0)
}

func (s *FieldContext) Or() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserOr, 0)
}

func (s *FieldContext) Literal() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserLiteral, 0)
}

func (s *FieldContext) Question() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserQuestion, 0)
}

func (s *FieldContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserLeftBracket, 0)
}

func (s *FieldContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *FieldContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserRightBracket, 0)
}

func (s *FieldContext) AllFunction_param() []IFunction_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunction_paramContext); ok {
			len++
		}
	}

	tst := make([]IFunction_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunction_paramContext); ok {
			tst[i] = t.(IFunction_paramContext)
			i++
		}
	}

	return tst
}

func (s *FieldContext) Function_param(i int) IFunction_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_paramContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_paramContext)
}

func (s *FieldContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(JsonTemplateParserComma)
}

func (s *FieldContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserComma, i)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonTemplateParser) Field() (localctx IFieldContext) {
	return p.field(0)
}

func (p *JsonTemplateParser) field(_p int) (localctx IFieldContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewFieldContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFieldContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, JsonTemplateParserRULE_field, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(196)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonTemplateParserLeftParen:
		{
			p.SetState(180)
			p.Match(JsonTemplateParserLeftParen)
		}
		{
			p.SetState(181)
			p.field(0)
		}
		{
			p.SetState(182)
			p.Match(JsonTemplateParserRightParen)
		}

	case JsonTemplateParserTrue:
		{
			p.SetState(184)
			p.Match(JsonTemplateParserTrue)
		}

	case JsonTemplateParserFalse:
		{
			p.SetState(185)
			p.Match(JsonTemplateParserFalse)
		}

	case JsonTemplateParserNull:
		{
			p.SetState(186)
			p.Match(JsonTemplateParserNull)
		}

	case JsonTemplateParserNUMBER:
		{
			p.SetState(187)
			p.Match(JsonTemplateParserNUMBER)
		}

	case JsonTemplateParserESCAPED_STRING:
		{
			p.SetState(188)
			p.Match(JsonTemplateParserESCAPED_STRING)
		}

	case JsonTemplateParserLeftBracket:
		{
			p.SetState(189)
			p.Array()
		}

	case JsonTemplateParserLeftBrace:
		{
			p.SetState(190)
			p.Object()
		}

	case JsonTemplateParserSTRING:
		{
			p.SetState(191)
			p.Name()
		}

	case JsonTemplateParserSubtract:
		{
			p.SetState(192)
			p.Match(JsonTemplateParserSubtract)
		}
		{
			p.SetState(193)
			p.field(16)
		}

	case JsonTemplateParserNot:
		{
			p.SetState(194)
			p.Match(JsonTemplateParserNot)
		}
		{
			p.SetState(195)
			p.field(15)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(271)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
			case 1:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(198)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(199)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JsonTemplateParserMultiply || _la == JsonTemplateParserDivide) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(200)
					p.field(15)
				}

			case 2:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(201)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(202)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JsonTemplateParserAdd || _la == JsonTemplateParserSubtract) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(203)
					p.field(14)
				}

			case 3:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(204)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(205)
					p.Match(JsonTemplateParserRange)
				}
				{
					p.SetState(206)
					p.field(13)
				}

			case 4:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(207)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(208)
					p.Match(JsonTemplateParserNullCoalescing)
				}
				{
					p.SetState(209)
					p.field(12)
				}

			case 5:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(210)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(211)
					p.Match(JsonTemplateParserEqual)
				}
				{
					p.SetState(212)
					p.field(11)
				}

			case 6:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(213)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(214)
					p.Match(JsonTemplateParserLess)
				}
				{
					p.SetState(215)
					p.field(10)
				}

			case 7:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(216)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(217)
					p.Match(JsonTemplateParserLessOrEqual)
				}
				{
					p.SetState(218)
					p.field(9)
				}

			case 8:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(219)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(220)
					p.Match(JsonTemplateParserGreater)
				}
				{
					p.SetState(221)
					p.field(8)
				}

			case 9:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(222)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(223)
					p.Match(JsonTemplateParserGreaterOrEqual)
				}
				{
					p.SetState(224)
					p.field(7)
				}

			case 10:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(225)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(226)
					p.Match(JsonTemplateParserNotEqual)
				}
				{
					p.SetState(227)
					p.field(6)
				}

			case 11:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(228)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(229)
					p.Match(JsonTemplateParserAnd)
				}
				{
					p.SetState(230)
					p.field(5)
				}

			case 12:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(231)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(232)
					p.Match(JsonTemplateParserOr)
				}
				{
					p.SetState(233)
					p.field(4)
				}

			case 13:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(234)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(235)
					p.Match(JsonTemplateParserLiteral)
				}
				{
					p.SetState(236)
					p.field(2)
				}

			case 14:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(237)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}

				p.SetState(239)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonTemplateParserQuestion {
					{
						p.SetState(238)
						p.Match(JsonTemplateParserQuestion)
					}

				}
				{
					p.SetState(241)
					p.Match(JsonTemplateParserT__8)
				}
				{
					p.SetState(242)
					p.Name()
				}

			case 15:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(243)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}

				p.SetState(245)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonTemplateParserQuestion {
					{
						p.SetState(244)
						p.Match(JsonTemplateParserQuestion)
					}

				}
				{
					p.SetState(247)
					p.Match(JsonTemplateParserLeftBracket)
				}
				{
					p.SetState(248)
					p.Index()
				}
				{
					p.SetState(249)
					p.Match(JsonTemplateParserRightBracket)
				}

			case 16:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(251)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(252)
					p.Match(JsonTemplateParserLeftParen)
				}
				p.SetState(261)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34772141735936) != 0 {
					{
						p.SetState(253)
						p.Function_param()
					}
					p.SetState(258)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == JsonTemplateParserComma {
						{
							p.SetState(254)
							p.Match(JsonTemplateParserComma)
						}
						{
							p.SetState(255)
							p.Function_param()
						}

						p.SetState(260)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(263)
					p.Match(JsonTemplateParserRightParen)
				}

			case 17:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(264)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(265)
					p.Match(JsonTemplateParserQuestion)
				}
				{
					p.SetState(266)
					p.field(0)
				}
				p.SetState(269)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(267)
						p.Match(JsonTemplateParserT__9)
					}
					{
						p.SetState(268)
						p.field(0)
					}

				}

			}

		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonTemplateParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserLeftBracket, 0)
}

func (s *ArrayContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserRightBracket, 0)
}

func (s *ArrayContext) AllSpread_field() []ISpread_fieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISpread_fieldContext); ok {
			len++
		}
	}

	tst := make([]ISpread_fieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISpread_fieldContext); ok {
			tst[i] = t.(ISpread_fieldContext)
			i++
		}
	}

	return tst
}

func (s *ArrayContext) Spread_field(i int) ISpread_fieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpread_fieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpread_fieldContext)
}

func (s *ArrayContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(JsonTemplateParserComma)
}

func (s *ArrayContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserComma, i)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.ExitArray(s)
	}
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonTemplateParser) Array() (localctx IArrayContext) {
	this := p
	_ = this

	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JsonTemplateParserRULE_array)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.Match(JsonTemplateParserLeftBracket)
	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34780731670528) != 0 {
		{
			p.SetState(277)
			p.Spread_field()
		}
		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonTemplateParserComma {
			{
				p.SetState(278)
				p.Match(JsonTemplateParserComma)
			}
			{
				p.SetState(279)
				p.Spread_field()
			}

			p.SetState(284)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(287)
		p.Match(JsonTemplateParserRightBracket)
	}

	return localctx
}

// ISpread_fieldContext is an interface to support dynamic dispatch.
type ISpread_fieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpread_fieldContext differentiates from other interfaces.
	IsSpread_fieldContext()
}

type Spread_fieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpread_fieldContext() *Spread_fieldContext {
	var p = new(Spread_fieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonTemplateParserRULE_spread_field
	return p
}

func (*Spread_fieldContext) IsSpread_fieldContext() {}

func NewSpread_fieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Spread_fieldContext {
	var p = new(Spread_fieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateParserRULE_spread_field

	return p
}

func (s *Spread_fieldContext) GetParser() antlr.Parser { return s.parser }

func (s *Spread_fieldContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *Spread_fieldContext) Spread() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserSpread, 0)
}

func (s *Spread_fieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Spread_fieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Spread_fieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.EnterSpread_field(s)
	}
}

func (s *Spread_fieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.ExitSpread_field(s)
	}
}

func (s *Spread_fieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitSpread_field(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonTemplateParser) Spread_field() (localctx ISpread_fieldContext) {
	this := p
	_ = this

	localctx = NewSpread_fieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JsonTemplateParserRULE_spread_field)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JsonTemplateParserSpread {
		{
			p.SetState(289)
			p.Match(JsonTemplateParserSpread)
		}

	}
	{
		p.SetState(292)
		p.field(0)
	}

	return localctx
}

// IObjectContext is an interface to support dynamic dispatch.
type IObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectContext differentiates from other interfaces.
	IsObjectContext()
}

type ObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectContext() *ObjectContext {
	var p = new(ObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonTemplateParserRULE_object
	return p
}

func (*ObjectContext) IsObjectContext() {}

func NewObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectContext {
	var p = new(ObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateParserRULE_object

	return p
}

func (s *ObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserLeftBrace, 0)
}

func (s *ObjectContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserRightBrace, 0)
}

func (s *ObjectContext) AllObject_field() []IObject_fieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObject_fieldContext); ok {
			len++
		}
	}

	tst := make([]IObject_fieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObject_fieldContext); ok {
			tst[i] = t.(IObject_fieldContext)
			i++
		}
	}

	return tst
}

func (s *ObjectContext) Object_field(i int) IObject_fieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObject_fieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObject_fieldContext)
}

func (s *ObjectContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(JsonTemplateParserComma)
}

func (s *ObjectContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserComma, i)
}

func (s *ObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.EnterObject(s)
	}
}

func (s *ObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.ExitObject(s)
	}
}

func (s *ObjectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitObject(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonTemplateParser) Object() (localctx IObjectContext) {
	this := p
	_ = this

	localctx = NewObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JsonTemplateParserRULE_object)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(294)
		p.Match(JsonTemplateParserLeftBrace)
	}
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34780731670528) != 0 {
		{
			p.SetState(295)
			p.Object_field()
		}
		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonTemplateParserComma {
			{
				p.SetState(296)
				p.Match(JsonTemplateParserComma)
			}
			{
				p.SetState(297)
				p.Object_field()
			}

			p.SetState(302)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(305)
		p.Match(JsonTemplateParserRightBrace)
	}

	return localctx
}

// IObject_fieldContext is an interface to support dynamic dispatch.
type IObject_fieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObject_fieldContext differentiates from other interfaces.
	IsObject_fieldContext()
}

type Object_fieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_fieldContext() *Object_fieldContext {
	var p = new(Object_fieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonTemplateParserRULE_object_field
	return p
}

func (*Object_fieldContext) IsObject_fieldContext() {}

func NewObject_fieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_fieldContext {
	var p = new(Object_fieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateParserRULE_object_field

	return p
}

func (s *Object_fieldContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_fieldContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Object_fieldContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *Object_fieldContext) ESCAPED_STRING() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserESCAPED_STRING, 0)
}

func (s *Object_fieldContext) Spread() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserSpread, 0)
}

func (s *Object_fieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_fieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_fieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.EnterObject_field(s)
	}
}

func (s *Object_fieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.ExitObject_field(s)
	}
}

func (s *Object_fieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitObject_field(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonTemplateParser) Object_field() (localctx IObject_fieldContext) {
	this := p
	_ = this

	localctx = NewObject_fieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, JsonTemplateParserRULE_object_field)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(307)
			p.Name()
		}
		{
			p.SetState(308)
			p.Match(JsonTemplateParserT__9)
		}
		{
			p.SetState(309)
			p.field(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(311)
			p.Match(JsonTemplateParserESCAPED_STRING)
		}
		{
			p.SetState(312)
			p.Match(JsonTemplateParserT__9)
		}
		{
			p.SetState(313)
			p.field(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserSpread {
			{
				p.SetState(314)
				p.Match(JsonTemplateParserSpread)
			}

		}
		{
			p.SetState(317)
			p.field(0)
		}

	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonTemplateParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) STRING() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserSTRING, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.ExitName(s)
	}
}

func (s *NameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonTemplateParser) Name() (localctx INameContext) {
	this := p
	_ = this

	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, JsonTemplateParserRULE_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.Match(JsonTemplateParserSTRING)
	}

	return localctx
}

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonTemplateParserRULE_index
	return p
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *IndexContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserNUMBER, 0)
}

func (s *IndexContext) ESCAPED_STRING() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserESCAPED_STRING, 0)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.ExitIndex(s)
	}
}

func (s *IndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonTemplateParser) Index() (localctx IIndexContext) {
	this := p
	_ = this

	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, JsonTemplateParserRULE_index)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(322)
			p.field(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(323)
			p.Match(JsonTemplateParserNUMBER)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(324)
			p.Match(JsonTemplateParserESCAPED_STRING)
		}

	}

	return localctx
}

func (p *JsonTemplateParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 6:
		var t *FieldContext = nil
		if localctx != nil {
			t = localctx.(*FieldContext)
		}
		return p.Field_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *JsonTemplateParser) Field_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 16:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
