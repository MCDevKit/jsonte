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
		"script", "statement", "expression", "lambda", "function_param", "field",
		"array", "object", "object_field", "name", "index",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 44, 347, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 5, 0, 24, 8, 0, 10, 0, 12, 0, 27, 9, 0, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 5, 1, 45, 8, 1, 10, 1, 12, 1, 48, 9, 1, 3, 1, 50, 8, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 5, 1, 56, 8, 1, 10, 1, 12, 1, 59, 9, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 5, 1, 67, 8, 1, 10, 1, 12, 1, 70, 9, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 5, 1, 78, 8, 1, 10, 1, 12, 1, 81, 9, 1, 1, 1, 1, 1, 5,
		1, 85, 8, 1, 10, 1, 12, 1, 88, 9, 1, 1, 1, 1, 1, 1, 1, 5, 1, 93, 8, 1,
		10, 1, 12, 1, 96, 9, 1, 1, 1, 3, 1, 99, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5,
		1, 105, 8, 1, 10, 1, 12, 1, 108, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5,
		1, 115, 8, 1, 10, 1, 12, 1, 118, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
		1, 125, 8, 1, 1, 2, 1, 2, 3, 2, 129, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 3, 2, 137, 8, 2, 3, 2, 139, 8, 2, 1, 2, 1, 2, 3, 2, 143, 8, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 3, 2, 149, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 155,
		8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 161, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		3, 2, 167, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 173, 8, 2, 1, 2, 1, 2, 1,
		2, 3, 2, 178, 8, 2, 1, 2, 1, 2, 3, 2, 182, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 192, 8, 3, 10, 3, 12, 3, 195, 9, 3, 5, 3,
		197, 8, 3, 10, 3, 12, 3, 200, 9, 3, 1, 3, 1, 3, 1, 3, 3, 3, 205, 8, 3,
		1, 4, 1, 4, 3, 4, 209, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 228,
		8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 3, 5, 268, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 274, 8,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 285, 8,
		5, 10, 5, 12, 5, 288, 9, 5, 3, 5, 290, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 3, 5, 298, 8, 5, 5, 5, 300, 8, 5, 10, 5, 12, 5, 303, 9, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 5, 6, 309, 8, 6, 10, 6, 12, 6, 312, 9, 6, 3, 6, 314,
		8, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 322, 8, 7, 10, 7, 12, 7,
		325, 9, 7, 3, 7, 327, 8, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 3, 8, 338, 8, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 3, 10, 345,
		8, 10, 1, 10, 0, 1, 10, 11, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 0, 2,
		1, 0, 22, 23, 1, 0, 20, 21, 407, 0, 25, 1, 0, 0, 0, 2, 124, 1, 0, 0, 0,
		4, 181, 1, 0, 0, 0, 6, 204, 1, 0, 0, 0, 8, 208, 1, 0, 0, 0, 10, 227, 1,
		0, 0, 0, 12, 304, 1, 0, 0, 0, 14, 317, 1, 0, 0, 0, 16, 337, 1, 0, 0, 0,
		18, 339, 1, 0, 0, 0, 20, 344, 1, 0, 0, 0, 22, 24, 3, 2, 1, 0, 23, 22, 1,
		0, 0, 0, 24, 27, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26,
		1, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 28, 29, 3, 10, 5, 0, 29, 30, 5, 1, 0,
		0, 30, 125, 1, 0, 0, 0, 31, 32, 3, 10, 5, 0, 32, 33, 5, 30, 0, 0, 33, 34,
		3, 10, 5, 0, 34, 35, 5, 1, 0, 0, 35, 125, 1, 0, 0, 0, 36, 37, 5, 2, 0,
		0, 37, 38, 3, 10, 5, 0, 38, 39, 5, 1, 0, 0, 39, 125, 1, 0, 0, 0, 40, 49,
		5, 3, 0, 0, 41, 46, 3, 18, 9, 0, 42, 43, 5, 34, 0, 0, 43, 45, 3, 18, 9,
		0, 44, 42, 1, 0, 0, 0, 45, 48, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 46, 47,
		1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 49, 41, 1, 0, 0, 0,
		49, 50, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 52, 5, 4, 0, 0, 52, 53, 3,
		10, 5, 0, 53, 57, 5, 36, 0, 0, 54, 56, 3, 2, 1, 0, 55, 54, 1, 0, 0, 0,
		56, 59, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 60, 1,
		0, 0, 0, 59, 57, 1, 0, 0, 0, 60, 61, 5, 37, 0, 0, 61, 125, 1, 0, 0, 0,
		62, 63, 5, 5, 0, 0, 63, 64, 3, 10, 5, 0, 64, 68, 5, 36, 0, 0, 65, 67, 3,
		2, 1, 0, 66, 65, 1, 0, 0, 0, 67, 70, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68,
		69, 1, 0, 0, 0, 69, 71, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 71, 86, 5, 37,
		0, 0, 72, 73, 5, 6, 0, 0, 73, 74, 5, 5, 0, 0, 74, 75, 3, 10, 5, 0, 75,
		79, 5, 36, 0, 0, 76, 78, 3, 2, 1, 0, 77, 76, 1, 0, 0, 0, 78, 81, 1, 0,
		0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 82, 1, 0, 0, 0, 81, 79,
		1, 0, 0, 0, 82, 83, 5, 37, 0, 0, 83, 85, 1, 0, 0, 0, 84, 72, 1, 0, 0, 0,
		85, 88, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 98, 1,
		0, 0, 0, 88, 86, 1, 0, 0, 0, 89, 90, 5, 6, 0, 0, 90, 94, 5, 36, 0, 0, 91,
		93, 3, 2, 1, 0, 92, 91, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 92, 1, 0, 0,
		0, 94, 95, 1, 0, 0, 0, 95, 97, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 97, 99,
		5, 37, 0, 0, 98, 89, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 125, 1, 0, 0,
		0, 100, 101, 5, 7, 0, 0, 101, 102, 3, 10, 5, 0, 102, 106, 5, 36, 0, 0,
		103, 105, 3, 2, 1, 0, 104, 103, 1, 0, 0, 0, 105, 108, 1, 0, 0, 0, 106,
		104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 109, 1, 0, 0, 0, 108, 106,
		1, 0, 0, 0, 109, 110, 5, 37, 0, 0, 110, 125, 1, 0, 0, 0, 111, 112, 5, 8,
		0, 0, 112, 116, 5, 36, 0, 0, 113, 115, 3, 2, 1, 0, 114, 113, 1, 0, 0, 0,
		115, 118, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117,
		119, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 119, 120, 5, 37, 0, 0, 120, 121,
		5, 7, 0, 0, 121, 122, 3, 10, 5, 0, 122, 123, 5, 1, 0, 0, 123, 125, 1, 0,
		0, 0, 124, 28, 1, 0, 0, 0, 124, 31, 1, 0, 0, 0, 124, 36, 1, 0, 0, 0, 124,
		40, 1, 0, 0, 0, 124, 62, 1, 0, 0, 0, 124, 100, 1, 0, 0, 0, 124, 111, 1,
		0, 0, 0, 125, 3, 1, 0, 0, 0, 126, 127, 5, 36, 0, 0, 127, 129, 5, 36, 0,
		0, 128, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130,
		131, 5, 28, 0, 0, 131, 138, 3, 10, 5, 0, 132, 133, 5, 33, 0, 0, 133, 136,
		3, 18, 9, 0, 134, 135, 5, 34, 0, 0, 135, 137, 3, 18, 9, 0, 136, 134, 1,
		0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 139, 1, 0, 0, 0, 138, 132, 1, 0, 0,
		0, 138, 139, 1, 0, 0, 0, 139, 142, 1, 0, 0, 0, 140, 141, 5, 37, 0, 0, 141,
		143, 5, 37, 0, 0, 142, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144,
		1, 0, 0, 0, 144, 145, 5, 0, 0, 1, 145, 182, 1, 0, 0, 0, 146, 147, 5, 36,
		0, 0, 147, 149, 5, 36, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0,
		149, 150, 1, 0, 0, 0, 150, 151, 5, 29, 0, 0, 151, 154, 3, 10, 5, 0, 152,
		153, 5, 37, 0, 0, 153, 155, 5, 37, 0, 0, 154, 152, 1, 0, 0, 0, 154, 155,
		1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 157, 5, 0, 0, 1, 157, 182, 1, 0,
		0, 0, 158, 159, 5, 36, 0, 0, 159, 161, 5, 36, 0, 0, 160, 158, 1, 0, 0,
		0, 160, 161, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 163, 5, 30, 0, 0, 163,
		166, 3, 10, 5, 0, 164, 165, 5, 37, 0, 0, 165, 167, 5, 37, 0, 0, 166, 164,
		1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 169, 5, 0,
		0, 1, 169, 182, 1, 0, 0, 0, 170, 171, 5, 36, 0, 0, 171, 173, 5, 36, 0,
		0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174,
		177, 3, 10, 5, 0, 175, 176, 5, 37, 0, 0, 176, 178, 5, 37, 0, 0, 177, 175,
		1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 180, 5, 0,
		0, 1, 180, 182, 1, 0, 0, 0, 181, 128, 1, 0, 0, 0, 181, 148, 1, 0, 0, 0,
		181, 160, 1, 0, 0, 0, 181, 172, 1, 0, 0, 0, 182, 5, 1, 0, 0, 0, 183, 184,
		3, 18, 9, 0, 184, 185, 5, 35, 0, 0, 185, 186, 3, 10, 5, 0, 186, 205, 1,
		0, 0, 0, 187, 198, 5, 24, 0, 0, 188, 193, 3, 18, 9, 0, 189, 190, 5, 34,
		0, 0, 190, 192, 3, 18, 9, 0, 191, 189, 1, 0, 0, 0, 192, 195, 1, 0, 0, 0,
		193, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 197, 1, 0, 0, 0, 195,
		193, 1, 0, 0, 0, 196, 188, 1, 0, 0, 0, 197, 200, 1, 0, 0, 0, 198, 196,
		1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 201, 1, 0, 0, 0, 200, 198, 1, 0,
		0, 0, 201, 202, 5, 25, 0, 0, 202, 203, 5, 35, 0, 0, 203, 205, 3, 10, 5,
		0, 204, 183, 1, 0, 0, 0, 204, 187, 1, 0, 0, 0, 205, 7, 1, 0, 0, 0, 206,
		209, 3, 10, 5, 0, 207, 209, 3, 6, 3, 0, 208, 206, 1, 0, 0, 0, 208, 207,
		1, 0, 0, 0, 209, 9, 1, 0, 0, 0, 210, 211, 6, 5, -1, 0, 211, 212, 5, 24,
		0, 0, 212, 213, 3, 10, 5, 0, 213, 214, 5, 25, 0, 0, 214, 228, 1, 0, 0,
		0, 215, 228, 5, 40, 0, 0, 216, 228, 5, 39, 0, 0, 217, 228, 5, 38, 0, 0,
		218, 228, 5, 43, 0, 0, 219, 228, 5, 41, 0, 0, 220, 228, 3, 12, 6, 0, 221,
		228, 3, 14, 7, 0, 222, 228, 3, 18, 9, 0, 223, 224, 5, 21, 0, 0, 224, 228,
		3, 10, 5, 15, 225, 226, 5, 19, 0, 0, 226, 228, 3, 10, 5, 14, 227, 210,
		1, 0, 0, 0, 227, 215, 1, 0, 0, 0, 227, 216, 1, 0, 0, 0, 227, 217, 1, 0,
		0, 0, 227, 218, 1, 0, 0, 0, 227, 219, 1, 0, 0, 0, 227, 220, 1, 0, 0, 0,
		227, 221, 1, 0, 0, 0, 227, 222, 1, 0, 0, 0, 227, 223, 1, 0, 0, 0, 227,
		225, 1, 0, 0, 0, 228, 301, 1, 0, 0, 0, 229, 230, 10, 13, 0, 0, 230, 231,
		7, 0, 0, 0, 231, 300, 3, 10, 5, 14, 232, 233, 10, 12, 0, 0, 233, 234, 7,
		1, 0, 0, 234, 300, 3, 10, 5, 13, 235, 236, 10, 11, 0, 0, 236, 237, 5, 32,
		0, 0, 237, 300, 3, 10, 5, 12, 238, 239, 10, 10, 0, 0, 239, 240, 5, 31,
		0, 0, 240, 300, 3, 10, 5, 11, 241, 242, 10, 9, 0, 0, 242, 243, 5, 13, 0,
		0, 243, 300, 3, 10, 5, 10, 244, 245, 10, 8, 0, 0, 245, 246, 5, 11, 0, 0,
		246, 300, 3, 10, 5, 9, 247, 248, 10, 7, 0, 0, 248, 249, 5, 12, 0, 0, 249,
		300, 3, 10, 5, 8, 250, 251, 10, 6, 0, 0, 251, 252, 5, 14, 0, 0, 252, 300,
		3, 10, 5, 7, 253, 254, 10, 5, 0, 0, 254, 255, 5, 15, 0, 0, 255, 300, 3,
		10, 5, 6, 256, 257, 10, 4, 0, 0, 257, 258, 5, 16, 0, 0, 258, 300, 3, 10,
		5, 5, 259, 260, 10, 3, 0, 0, 260, 261, 5, 17, 0, 0, 261, 300, 3, 10, 5,
		4, 262, 263, 10, 2, 0, 0, 263, 264, 5, 18, 0, 0, 264, 300, 3, 10, 5, 3,
		265, 267, 10, 18, 0, 0, 266, 268, 5, 29, 0, 0, 267, 266, 1, 0, 0, 0, 267,
		268, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269, 270, 5, 9, 0, 0, 270, 300,
		3, 18, 9, 0, 271, 273, 10, 17, 0, 0, 272, 274, 5, 29, 0, 0, 273, 272, 1,
		0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 275, 1, 0, 0, 0, 275, 276, 5, 26, 0,
		0, 276, 277, 3, 20, 10, 0, 277, 278, 5, 27, 0, 0, 278, 300, 1, 0, 0, 0,
		279, 280, 10, 16, 0, 0, 280, 289, 5, 24, 0, 0, 281, 286, 3, 8, 4, 0, 282,
		283, 5, 34, 0, 0, 283, 285, 3, 8, 4, 0, 284, 282, 1, 0, 0, 0, 285, 288,
		1, 0, 0, 0, 286, 284, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 290, 1, 0,
		0, 0, 288, 286, 1, 0, 0, 0, 289, 281, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0,
		290, 291, 1, 0, 0, 0, 291, 300, 5, 25, 0, 0, 292, 293, 10, 1, 0, 0, 293,
		294, 5, 29, 0, 0, 294, 297, 3, 10, 5, 0, 295, 296, 5, 10, 0, 0, 296, 298,
		3, 10, 5, 0, 297, 295, 1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 300, 1, 0,
		0, 0, 299, 229, 1, 0, 0, 0, 299, 232, 1, 0, 0, 0, 299, 235, 1, 0, 0, 0,
		299, 238, 1, 0, 0, 0, 299, 241, 1, 0, 0, 0, 299, 244, 1, 0, 0, 0, 299,
		247, 1, 0, 0, 0, 299, 250, 1, 0, 0, 0, 299, 253, 1, 0, 0, 0, 299, 256,
		1, 0, 0, 0, 299, 259, 1, 0, 0, 0, 299, 262, 1, 0, 0, 0, 299, 265, 1, 0,
		0, 0, 299, 271, 1, 0, 0, 0, 299, 279, 1, 0, 0, 0, 299, 292, 1, 0, 0, 0,
		300, 303, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302,
		11, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 304, 313, 5, 26, 0, 0, 305, 310,
		3, 10, 5, 0, 306, 307, 5, 34, 0, 0, 307, 309, 3, 10, 5, 0, 308, 306, 1,
		0, 0, 0, 309, 312, 1, 0, 0, 0, 310, 308, 1, 0, 0, 0, 310, 311, 1, 0, 0,
		0, 311, 314, 1, 0, 0, 0, 312, 310, 1, 0, 0, 0, 313, 305, 1, 0, 0, 0, 313,
		314, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 316, 5, 27, 0, 0, 316, 13,
		1, 0, 0, 0, 317, 326, 5, 36, 0, 0, 318, 323, 3, 16, 8, 0, 319, 320, 5,
		34, 0, 0, 320, 322, 3, 16, 8, 0, 321, 319, 1, 0, 0, 0, 322, 325, 1, 0,
		0, 0, 323, 321, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 327, 1, 0, 0, 0,
		325, 323, 1, 0, 0, 0, 326, 318, 1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327,
		328, 1, 0, 0, 0, 328, 329, 5, 37, 0, 0, 329, 15, 1, 0, 0, 0, 330, 331,
		3, 18, 9, 0, 331, 332, 5, 10, 0, 0, 332, 333, 3, 10, 5, 0, 333, 338, 1,
		0, 0, 0, 334, 335, 5, 41, 0, 0, 335, 336, 5, 10, 0, 0, 336, 338, 3, 10,
		5, 0, 337, 330, 1, 0, 0, 0, 337, 334, 1, 0, 0, 0, 338, 17, 1, 0, 0, 0,
		339, 340, 5, 42, 0, 0, 340, 19, 1, 0, 0, 0, 341, 345, 3, 10, 5, 0, 342,
		345, 5, 43, 0, 0, 343, 345, 5, 41, 0, 0, 344, 341, 1, 0, 0, 0, 344, 342,
		1, 0, 0, 0, 344, 343, 1, 0, 0, 0, 345, 21, 1, 0, 0, 0, 41, 25, 46, 49,
		57, 68, 79, 86, 94, 98, 106, 116, 124, 128, 136, 138, 142, 148, 154, 160,
		166, 172, 177, 181, 193, 198, 204, 208, 227, 267, 273, 286, 289, 297, 299,
		301, 310, 313, 323, 326, 337, 344,
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
	JsonTemplateParserAs             = 33
	JsonTemplateParserComma          = 34
	JsonTemplateParserArrow          = 35
	JsonTemplateParserLeftBrace      = 36
	JsonTemplateParserRightBrace     = 37
	JsonTemplateParserNull           = 38
	JsonTemplateParserFalse          = 39
	JsonTemplateParserTrue           = 40
	JsonTemplateParserESCAPED_STRING = 41
	JsonTemplateParserSTRING         = 42
	JsonTemplateParserNUMBER         = 43
	JsonTemplateParserWS             = 44
)

// JsonTemplateParser rules.
const (
	JsonTemplateParserRULE_script         = 0
	JsonTemplateParserRULE_statement      = 1
	JsonTemplateParserRULE_expression     = 2
	JsonTemplateParserRULE_lambda         = 3
	JsonTemplateParserRULE_function_param = 4
	JsonTemplateParserRULE_field          = 5
	JsonTemplateParserRULE_array          = 6
	JsonTemplateParserRULE_object         = 7
	JsonTemplateParserRULE_object_field   = 8
	JsonTemplateParserRULE_name           = 9
	JsonTemplateParserRULE_index          = 10
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
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17386114122156) != 0 {
		{
			p.SetState(22)
			p.Statement()
		}

		p.SetState(27)
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

func (s *StatementContext) Literal() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserLiteral, 0)
}

func (s *StatementContext) AllLeftBrace() []antlr.TerminalNode {
	return s.GetTokens(JsonTemplateParserLeftBrace)
}

func (s *StatementContext) LeftBrace(i int) antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserLeftBrace, i)
}

func (s *StatementContext) AllRightBrace() []antlr.TerminalNode {
	return s.GetTokens(JsonTemplateParserRightBrace)
}

func (s *StatementContext) RightBrace(i int) antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserRightBrace, i)
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

func (s *StatementContext) AllStatement() []IStatementContext {
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

func (s *StatementContext) Statement(i int) IStatementContext {
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

	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.field(0)
		}
		{
			p.SetState(29)
			p.Match(JsonTemplateParserT__0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.field(0)
		}
		{
			p.SetState(32)
			p.Match(JsonTemplateParserLiteral)
		}
		{
			p.SetState(33)
			p.field(0)
		}
		{
			p.SetState(34)
			p.Match(JsonTemplateParserT__0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(36)
			p.Match(JsonTemplateParserT__1)
		}
		{
			p.SetState(37)
			p.field(0)
		}
		{
			p.SetState(38)
			p.Match(JsonTemplateParserT__0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(40)
			p.Match(JsonTemplateParserT__2)
		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserSTRING {
			{
				p.SetState(41)
				p.Name()
			}
			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == JsonTemplateParserComma {
				{
					p.SetState(42)
					p.Match(JsonTemplateParserComma)
				}
				{
					p.SetState(43)
					p.Name()
				}

				p.SetState(48)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(51)
			p.Match(JsonTemplateParserT__3)
		}
		{
			p.SetState(52)
			p.field(0)
		}
		{
			p.SetState(53)
			p.Match(JsonTemplateParserLeftBrace)
		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17386114122156) != 0 {
			{
				p.SetState(54)
				p.Statement()
			}

			p.SetState(59)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(60)
			p.Match(JsonTemplateParserRightBrace)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(62)
			p.Match(JsonTemplateParserT__4)
		}
		{
			p.SetState(63)
			p.field(0)
		}
		{
			p.SetState(64)
			p.Match(JsonTemplateParserLeftBrace)
		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17386114122156) != 0 {
			{
				p.SetState(65)
				p.Statement()
			}

			p.SetState(70)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(71)
			p.Match(JsonTemplateParserRightBrace)
		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(72)
					p.Match(JsonTemplateParserT__5)
				}
				{
					p.SetState(73)
					p.Match(JsonTemplateParserT__4)
				}
				{
					p.SetState(74)
					p.field(0)
				}
				{
					p.SetState(75)
					p.Match(JsonTemplateParserLeftBrace)
				}
				p.SetState(79)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17386114122156) != 0 {
					{
						p.SetState(76)
						p.Statement()
					}

					p.SetState(81)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(82)
					p.Match(JsonTemplateParserRightBrace)
				}

			}
			p.SetState(88)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserT__5 {
			{
				p.SetState(89)
				p.Match(JsonTemplateParserT__5)
			}
			{
				p.SetState(90)
				p.Match(JsonTemplateParserLeftBrace)
			}
			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17386114122156) != 0 {
				{
					p.SetState(91)
					p.Statement()
				}

				p.SetState(96)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(97)
				p.Match(JsonTemplateParserRightBrace)
			}

		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(100)
			p.Match(JsonTemplateParserT__6)
		}
		{
			p.SetState(101)
			p.field(0)
		}
		{
			p.SetState(102)
			p.Match(JsonTemplateParserLeftBrace)
		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17386114122156) != 0 {
			{
				p.SetState(103)
				p.Statement()
			}

			p.SetState(108)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(109)
			p.Match(JsonTemplateParserRightBrace)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(111)
			p.Match(JsonTemplateParserT__7)
		}
		{
			p.SetState(112)
			p.Match(JsonTemplateParserLeftBrace)
		}
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17386114122156) != 0 {
			{
				p.SetState(113)
				p.Statement()
			}

			p.SetState(118)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(119)
			p.Match(JsonTemplateParserRightBrace)
		}
		{
			p.SetState(120)
			p.Match(JsonTemplateParserT__6)
		}
		{
			p.SetState(121)
			p.field(0)
		}
		{
			p.SetState(122)
			p.Match(JsonTemplateParserT__0)
		}

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
	p.EnterRule(localctx, 4, JsonTemplateParserRULE_expression)
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

	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserLeftBrace {
			{
				p.SetState(126)
				p.Match(JsonTemplateParserLeftBrace)
			}
			{
				p.SetState(127)
				p.Match(JsonTemplateParserLeftBrace)
			}

		}
		{
			p.SetState(130)
			p.Match(JsonTemplateParserIteration)
		}
		{
			p.SetState(131)
			p.field(0)
		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserAs {
			{
				p.SetState(132)
				p.Match(JsonTemplateParserAs)
			}
			{
				p.SetState(133)
				p.Name()
			}
			p.SetState(136)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == JsonTemplateParserComma {
				{
					p.SetState(134)
					p.Match(JsonTemplateParserComma)
				}
				{
					p.SetState(135)
					p.Name()
				}

			}

		}
		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserRightBrace {
			{
				p.SetState(140)
				p.Match(JsonTemplateParserRightBrace)
			}
			{
				p.SetState(141)
				p.Match(JsonTemplateParserRightBrace)
			}

		}
		{
			p.SetState(144)
			p.Match(JsonTemplateParserEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserLeftBrace {
			{
				p.SetState(146)
				p.Match(JsonTemplateParserLeftBrace)
			}
			{
				p.SetState(147)
				p.Match(JsonTemplateParserLeftBrace)
			}

		}
		{
			p.SetState(150)
			p.Match(JsonTemplateParserQuestion)
		}
		{
			p.SetState(151)
			p.field(0)
		}
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserRightBrace {
			{
				p.SetState(152)
				p.Match(JsonTemplateParserRightBrace)
			}
			{
				p.SetState(153)
				p.Match(JsonTemplateParserRightBrace)
			}

		}
		{
			p.SetState(156)
			p.Match(JsonTemplateParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserLeftBrace {
			{
				p.SetState(158)
				p.Match(JsonTemplateParserLeftBrace)
			}
			{
				p.SetState(159)
				p.Match(JsonTemplateParserLeftBrace)
			}

		}
		{
			p.SetState(162)
			p.Match(JsonTemplateParserLiteral)
		}
		{
			p.SetState(163)
			p.field(0)
		}
		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserRightBrace {
			{
				p.SetState(164)
				p.Match(JsonTemplateParserRightBrace)
			}
			{
				p.SetState(165)
				p.Match(JsonTemplateParserRightBrace)
			}

		}
		{
			p.SetState(168)
			p.Match(JsonTemplateParserEOF)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(172)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(170)
				p.Match(JsonTemplateParserLeftBrace)
			}
			{
				p.SetState(171)
				p.Match(JsonTemplateParserLeftBrace)
			}

		}
		{
			p.SetState(174)
			p.field(0)
		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserRightBrace {
			{
				p.SetState(175)
				p.Match(JsonTemplateParserRightBrace)
			}
			{
				p.SetState(176)
				p.Match(JsonTemplateParserRightBrace)
			}

		}
		{
			p.SetState(179)
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
	p.EnterRule(localctx, 6, JsonTemplateParserRULE_lambda)
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

	p.SetState(204)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonTemplateParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(183)
			p.Name()
		}
		{
			p.SetState(184)
			p.Match(JsonTemplateParserArrow)
		}
		{
			p.SetState(185)
			p.field(0)
		}

	case JsonTemplateParserLeftParen:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(187)
			p.Match(JsonTemplateParserLeftParen)
		}
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonTemplateParserSTRING {
			{
				p.SetState(188)
				p.Name()
			}
			p.SetState(193)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == JsonTemplateParserComma {
				{
					p.SetState(189)
					p.Match(JsonTemplateParserComma)
				}
				{
					p.SetState(190)
					p.Name()
				}

				p.SetState(195)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

			p.SetState(200)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(201)
			p.Match(JsonTemplateParserRightParen)
		}
		{
			p.SetState(202)
			p.Match(JsonTemplateParserArrow)
		}
		{
			p.SetState(203)
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
	p.EnterRule(localctx, 8, JsonTemplateParserRULE_function_param)

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

	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(206)
			p.field(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(207)
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
	_startState := 10
	p.EnterRecursionRule(localctx, 10, JsonTemplateParserRULE_field, _p)
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
	p.SetState(227)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonTemplateParserLeftParen:
		{
			p.SetState(211)
			p.Match(JsonTemplateParserLeftParen)
		}
		{
			p.SetState(212)
			p.field(0)
		}
		{
			p.SetState(213)
			p.Match(JsonTemplateParserRightParen)
		}

	case JsonTemplateParserTrue:
		{
			p.SetState(215)
			p.Match(JsonTemplateParserTrue)
		}

	case JsonTemplateParserFalse:
		{
			p.SetState(216)
			p.Match(JsonTemplateParserFalse)
		}

	case JsonTemplateParserNull:
		{
			p.SetState(217)
			p.Match(JsonTemplateParserNull)
		}

	case JsonTemplateParserNUMBER:
		{
			p.SetState(218)
			p.Match(JsonTemplateParserNUMBER)
		}

	case JsonTemplateParserESCAPED_STRING:
		{
			p.SetState(219)
			p.Match(JsonTemplateParserESCAPED_STRING)
		}

	case JsonTemplateParserLeftBracket:
		{
			p.SetState(220)
			p.Array()
		}

	case JsonTemplateParserLeftBrace:
		{
			p.SetState(221)
			p.Object()
		}

	case JsonTemplateParserSTRING:
		{
			p.SetState(222)
			p.Name()
		}

	case JsonTemplateParserSubtract:
		{
			p.SetState(223)
			p.Match(JsonTemplateParserSubtract)
		}
		{
			p.SetState(224)
			p.field(15)
		}

	case JsonTemplateParserNot:
		{
			p.SetState(225)
			p.Match(JsonTemplateParserNot)
		}
		{
			p.SetState(226)
			p.field(14)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(299)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
			case 1:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(229)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(230)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JsonTemplateParserMultiply || _la == JsonTemplateParserDivide) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(231)
					p.field(14)
				}

			case 2:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(232)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(233)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JsonTemplateParserAdd || _la == JsonTemplateParserSubtract) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(234)
					p.field(13)
				}

			case 3:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(235)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(236)
					p.Match(JsonTemplateParserRange)
				}
				{
					p.SetState(237)
					p.field(12)
				}

			case 4:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(238)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(239)
					p.Match(JsonTemplateParserNullCoalescing)
				}
				{
					p.SetState(240)
					p.field(11)
				}

			case 5:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(241)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(242)
					p.Match(JsonTemplateParserEqual)
				}
				{
					p.SetState(243)
					p.field(10)
				}

			case 6:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(244)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(245)
					p.Match(JsonTemplateParserLess)
				}
				{
					p.SetState(246)
					p.field(9)
				}

			case 7:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(247)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(248)
					p.Match(JsonTemplateParserLessOrEqual)
				}
				{
					p.SetState(249)
					p.field(8)
				}

			case 8:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(250)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(251)
					p.Match(JsonTemplateParserGreater)
				}
				{
					p.SetState(252)
					p.field(7)
				}

			case 9:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(253)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(254)
					p.Match(JsonTemplateParserGreaterOrEqual)
				}
				{
					p.SetState(255)
					p.field(6)
				}

			case 10:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(256)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(257)
					p.Match(JsonTemplateParserNotEqual)
				}
				{
					p.SetState(258)
					p.field(5)
				}

			case 11:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(259)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(260)
					p.Match(JsonTemplateParserAnd)
				}
				{
					p.SetState(261)
					p.field(4)
				}

			case 12:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(262)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(263)
					p.Match(JsonTemplateParserOr)
				}
				{
					p.SetState(264)
					p.field(3)
				}

			case 13:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(265)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}

				p.SetState(267)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonTemplateParserQuestion {
					{
						p.SetState(266)
						p.Match(JsonTemplateParserQuestion)
					}

				}
				{
					p.SetState(269)
					p.Match(JsonTemplateParserT__8)
				}
				{
					p.SetState(270)
					p.Name()
				}

			case 14:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(271)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}

				p.SetState(273)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonTemplateParserQuestion {
					{
						p.SetState(272)
						p.Match(JsonTemplateParserQuestion)
					}

				}
				{
					p.SetState(275)
					p.Match(JsonTemplateParserLeftBracket)
				}
				{
					p.SetState(276)
					p.Index()
				}
				{
					p.SetState(277)
					p.Match(JsonTemplateParserRightBracket)
				}

			case 15:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(279)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(280)
					p.Match(JsonTemplateParserLeftParen)
				}
				p.SetState(289)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17386114121728) != 0 {
					{
						p.SetState(281)
						p.Function_param()
					}
					p.SetState(286)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == JsonTemplateParserComma {
						{
							p.SetState(282)
							p.Match(JsonTemplateParserComma)
						}
						{
							p.SetState(283)
							p.Function_param()
						}

						p.SetState(288)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(291)
					p.Match(JsonTemplateParserRightParen)
				}

			case 16:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(292)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(293)
					p.Match(JsonTemplateParserQuestion)
				}
				{
					p.SetState(294)
					p.field(0)
				}
				p.SetState(297)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(295)
						p.Match(JsonTemplateParserT__9)
					}
					{
						p.SetState(296)
						p.field(0)
					}

				}

			}

		}
		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
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

func (s *ArrayContext) AllField() []IFieldContext {
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

func (s *ArrayContext) Field(i int) IFieldContext {
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
	p.EnterRule(localctx, 12, JsonTemplateParserRULE_array)
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
		p.SetState(304)
		p.Match(JsonTemplateParserLeftBracket)
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17386114121728) != 0 {
		{
			p.SetState(305)
			p.field(0)
		}
		p.SetState(310)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonTemplateParserComma {
			{
				p.SetState(306)
				p.Match(JsonTemplateParserComma)
			}
			{
				p.SetState(307)
				p.field(0)
			}

			p.SetState(312)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(315)
		p.Match(JsonTemplateParserRightBracket)
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
	p.EnterRule(localctx, 14, JsonTemplateParserRULE_object)
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
		p.SetState(317)
		p.Match(JsonTemplateParserLeftBrace)
	}
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JsonTemplateParserESCAPED_STRING || _la == JsonTemplateParserSTRING {
		{
			p.SetState(318)
			p.Object_field()
		}
		p.SetState(323)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonTemplateParserComma {
			{
				p.SetState(319)
				p.Match(JsonTemplateParserComma)
			}
			{
				p.SetState(320)
				p.Object_field()
			}

			p.SetState(325)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(328)
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
	p.EnterRule(localctx, 16, JsonTemplateParserRULE_object_field)

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

	p.SetState(337)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonTemplateParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(330)
			p.Name()
		}
		{
			p.SetState(331)
			p.Match(JsonTemplateParserT__9)
		}
		{
			p.SetState(332)
			p.field(0)
		}

	case JsonTemplateParserESCAPED_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(334)
			p.Match(JsonTemplateParserESCAPED_STRING)
		}
		{
			p.SetState(335)
			p.Match(JsonTemplateParserT__9)
		}
		{
			p.SetState(336)
			p.field(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 18, JsonTemplateParserRULE_name)

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
		p.SetState(339)
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
	p.EnterRule(localctx, 20, JsonTemplateParserRULE_index)

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

	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(341)
			p.field(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(342)
			p.Match(JsonTemplateParserNUMBER)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(343)
			p.Match(JsonTemplateParserESCAPED_STRING)
		}

	}

	return localctx
}

func (p *JsonTemplateParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
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
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
