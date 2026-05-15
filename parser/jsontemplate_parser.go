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

type JsonTemplate struct {
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
		"script", "statement", "statements", "expression", "lambda", "function_param",
		"field", "array", "spread_field", "object", "object_field", "template_string",
		"template_string_part", "name", "index",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 54, 375, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 5, 0,
		32, 8, 0, 10, 0, 12, 0, 35, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
		1, 43, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
		54, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 5, 1, 68, 8, 1, 10, 1, 12, 1, 71, 9, 1, 1, 1, 1, 1, 3, 1, 75,
		8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
		87, 8, 1, 1, 2, 1, 2, 5, 2, 91, 8, 2, 10, 2, 12, 2, 94, 9, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 3, 3, 100, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3,
		3, 108, 8, 3, 3, 3, 110, 8, 3, 1, 3, 1, 3, 3, 3, 114, 8, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 3, 3, 120, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 126, 8, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 3, 3, 132, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 138,
		8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 144, 8, 3, 1, 3, 1, 3, 1, 3, 3, 3,
		149, 8, 3, 1, 3, 1, 3, 3, 3, 153, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 167, 8, 4, 10, 4, 12, 4, 170,
		9, 4, 5, 4, 172, 8, 4, 10, 4, 12, 4, 175, 9, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 5, 4, 184, 8, 4, 10, 4, 12, 4, 187, 9, 4, 5, 4, 189,
		8, 4, 10, 4, 12, 4, 192, 9, 4, 1, 4, 1, 4, 1, 4, 3, 4, 197, 8, 4, 1, 5,
		3, 5, 200, 8, 5, 1, 5, 1, 5, 3, 5, 204, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 3, 6, 225, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 3, 6, 271, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 277, 8, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 288, 8, 6, 10, 6,
		12, 6, 291, 9, 6, 3, 6, 293, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		3, 6, 301, 8, 6, 5, 6, 303, 8, 6, 10, 6, 12, 6, 306, 9, 6, 1, 7, 1, 7,
		1, 7, 1, 7, 5, 7, 312, 8, 7, 10, 7, 12, 7, 315, 9, 7, 3, 7, 317, 8, 7,
		1, 7, 1, 7, 1, 8, 3, 8, 322, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9,
		5, 9, 330, 8, 9, 10, 9, 12, 9, 333, 9, 9, 3, 9, 335, 8, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 347, 8,
		10, 1, 10, 3, 10, 350, 8, 10, 1, 11, 1, 11, 5, 11, 354, 8, 11, 10, 11,
		12, 11, 357, 9, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3,
		12, 366, 8, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 3, 14, 373, 8, 14, 1,
		14, 0, 1, 12, 15, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		0, 2, 1, 0, 12, 13, 1, 0, 10, 11, 442, 0, 33, 1, 0, 0, 0, 2, 86, 1, 0,
		0, 0, 4, 88, 1, 0, 0, 0, 6, 152, 1, 0, 0, 0, 8, 196, 1, 0, 0, 0, 10, 203,
		1, 0, 0, 0, 12, 224, 1, 0, 0, 0, 14, 307, 1, 0, 0, 0, 16, 321, 1, 0, 0,
		0, 18, 325, 1, 0, 0, 0, 20, 349, 1, 0, 0, 0, 22, 351, 1, 0, 0, 0, 24, 365,
		1, 0, 0, 0, 26, 367, 1, 0, 0, 0, 28, 372, 1, 0, 0, 0, 30, 32, 3, 2, 1,
		0, 31, 30, 1, 0, 0, 0, 32, 35, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34,
		1, 0, 0, 0, 34, 1, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 36, 87, 3, 4, 2, 0,
		37, 38, 3, 12, 6, 0, 38, 39, 5, 29, 0, 0, 39, 87, 1, 0, 0, 0, 40, 42, 5,
		36, 0, 0, 41, 43, 3, 12, 6, 0, 42, 41, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0,
		43, 44, 1, 0, 0, 0, 44, 87, 5, 29, 0, 0, 45, 46, 5, 43, 0, 0, 46, 87, 5,
		29, 0, 0, 47, 48, 5, 44, 0, 0, 48, 87, 5, 29, 0, 0, 49, 50, 5, 37, 0, 0,
		50, 53, 3, 26, 13, 0, 51, 52, 5, 26, 0, 0, 52, 54, 3, 26, 13, 0, 53, 51,
		1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 56, 5, 38, 0, 0,
		56, 57, 3, 12, 6, 0, 57, 58, 3, 4, 2, 0, 58, 87, 1, 0, 0, 0, 59, 60, 5,
		39, 0, 0, 60, 61, 3, 12, 6, 0, 61, 69, 3, 4, 2, 0, 62, 63, 5, 40, 0, 0,
		63, 64, 5, 39, 0, 0, 64, 65, 3, 12, 6, 0, 65, 66, 3, 4, 2, 0, 66, 68, 1,
		0, 0, 0, 67, 62, 1, 0, 0, 0, 68, 71, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69,
		70, 1, 0, 0, 0, 70, 74, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 73, 5, 40,
		0, 0, 73, 75, 3, 4, 2, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 87,
		1, 0, 0, 0, 76, 77, 5, 41, 0, 0, 77, 78, 3, 12, 6, 0, 78, 79, 3, 4, 2,
		0, 79, 87, 1, 0, 0, 0, 80, 81, 5, 42, 0, 0, 81, 82, 3, 4, 2, 0, 82, 83,
		5, 41, 0, 0, 83, 84, 3, 12, 6, 0, 84, 85, 5, 29, 0, 0, 85, 87, 1, 0, 0,
		0, 86, 36, 1, 0, 0, 0, 86, 37, 1, 0, 0, 0, 86, 40, 1, 0, 0, 0, 86, 45,
		1, 0, 0, 0, 86, 47, 1, 0, 0, 0, 86, 49, 1, 0, 0, 0, 86, 59, 1, 0, 0, 0,
		86, 76, 1, 0, 0, 0, 86, 80, 1, 0, 0, 0, 87, 3, 1, 0, 0, 0, 88, 92, 5, 31,
		0, 0, 89, 91, 3, 2, 1, 0, 90, 89, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 90,
		1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 95, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0,
		95, 96, 5, 32, 0, 0, 96, 5, 1, 0, 0, 0, 97, 98, 5, 31, 0, 0, 98, 100, 5,
		31, 0, 0, 99, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0,
		101, 102, 5, 18, 0, 0, 102, 109, 3, 12, 6, 0, 103, 104, 5, 25, 0, 0, 104,
		107, 3, 26, 13, 0, 105, 106, 5, 26, 0, 0, 106, 108, 3, 26, 13, 0, 107,
		105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 110, 1, 0, 0, 0, 109, 103,
		1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 112, 5, 32,
		0, 0, 112, 114, 5, 32, 0, 0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0,
		114, 115, 1, 0, 0, 0, 115, 116, 5, 0, 0, 1, 116, 153, 1, 0, 0, 0, 117,
		118, 5, 31, 0, 0, 118, 120, 5, 31, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120,
		1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 122, 5, 19, 0, 0, 122, 125, 3, 12,
		6, 0, 123, 124, 5, 32, 0, 0, 124, 126, 5, 32, 0, 0, 125, 123, 1, 0, 0,
		0, 125, 126, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128, 5, 0, 0, 1, 128,
		153, 1, 0, 0, 0, 129, 130, 5, 31, 0, 0, 130, 132, 5, 31, 0, 0, 131, 129,
		1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 5, 21,
		0, 0, 134, 137, 3, 12, 6, 0, 135, 136, 5, 32, 0, 0, 136, 138, 5, 32, 0,
		0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139,
		140, 5, 0, 0, 1, 140, 153, 1, 0, 0, 0, 141, 142, 5, 31, 0, 0, 142, 144,
		5, 31, 0, 0, 143, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 145, 1, 0,
		0, 0, 145, 148, 3, 12, 6, 0, 146, 147, 5, 32, 0, 0, 147, 149, 5, 32, 0,
		0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150,
		151, 5, 0, 0, 1, 151, 153, 1, 0, 0, 0, 152, 99, 1, 0, 0, 0, 152, 119, 1,
		0, 0, 0, 152, 131, 1, 0, 0, 0, 152, 143, 1, 0, 0, 0, 153, 7, 1, 0, 0, 0,
		154, 155, 3, 26, 13, 0, 155, 156, 5, 27, 0, 0, 156, 157, 3, 12, 6, 0, 157,
		197, 1, 0, 0, 0, 158, 159, 3, 26, 13, 0, 159, 160, 5, 27, 0, 0, 160, 161,
		3, 4, 2, 0, 161, 197, 1, 0, 0, 0, 162, 173, 5, 14, 0, 0, 163, 168, 3, 26,
		13, 0, 164, 165, 5, 26, 0, 0, 165, 167, 3, 26, 13, 0, 166, 164, 1, 0, 0,
		0, 167, 170, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169,
		172, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 171, 163, 1, 0, 0, 0, 172, 175,
		1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 176, 1, 0,
		0, 0, 175, 173, 1, 0, 0, 0, 176, 177, 5, 15, 0, 0, 177, 178, 5, 27, 0,
		0, 178, 197, 3, 12, 6, 0, 179, 190, 5, 14, 0, 0, 180, 185, 3, 26, 13, 0,
		181, 182, 5, 26, 0, 0, 182, 184, 3, 26, 13, 0, 183, 181, 1, 0, 0, 0, 184,
		187, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 189,
		1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 188, 180, 1, 0, 0, 0, 189, 192, 1, 0,
		0, 0, 190, 188, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 193, 1, 0, 0, 0,
		192, 190, 1, 0, 0, 0, 193, 194, 5, 15, 0, 0, 194, 195, 5, 27, 0, 0, 195,
		197, 3, 4, 2, 0, 196, 154, 1, 0, 0, 0, 196, 158, 1, 0, 0, 0, 196, 162,
		1, 0, 0, 0, 196, 179, 1, 0, 0, 0, 197, 9, 1, 0, 0, 0, 198, 200, 5, 24,
		0, 0, 199, 198, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0,
		201, 204, 3, 12, 6, 0, 202, 204, 3, 8, 4, 0, 203, 199, 1, 0, 0, 0, 203,
		202, 1, 0, 0, 0, 204, 11, 1, 0, 0, 0, 205, 206, 6, 6, -1, 0, 206, 207,
		5, 14, 0, 0, 207, 208, 3, 12, 6, 0, 208, 209, 5, 15, 0, 0, 209, 225, 1,
		0, 0, 0, 210, 225, 5, 35, 0, 0, 211, 225, 5, 34, 0, 0, 212, 225, 5, 33,
		0, 0, 213, 225, 5, 48, 0, 0, 214, 225, 5, 46, 0, 0, 215, 225, 3, 22, 11,
		0, 216, 225, 3, 14, 7, 0, 217, 225, 3, 18, 9, 0, 218, 225, 3, 8, 4, 0,
		219, 225, 3, 26, 13, 0, 220, 221, 5, 11, 0, 0, 221, 225, 3, 12, 6, 17,
		222, 223, 5, 9, 0, 0, 223, 225, 3, 12, 6, 16, 224, 205, 1, 0, 0, 0, 224,
		210, 1, 0, 0, 0, 224, 211, 1, 0, 0, 0, 224, 212, 1, 0, 0, 0, 224, 213,
		1, 0, 0, 0, 224, 214, 1, 0, 0, 0, 224, 215, 1, 0, 0, 0, 224, 216, 1, 0,
		0, 0, 224, 217, 1, 0, 0, 0, 224, 218, 1, 0, 0, 0, 224, 219, 1, 0, 0, 0,
		224, 220, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 225, 304, 1, 0, 0, 0, 226,
		227, 10, 15, 0, 0, 227, 228, 7, 0, 0, 0, 228, 303, 3, 12, 6, 16, 229, 230,
		10, 14, 0, 0, 230, 231, 7, 1, 0, 0, 231, 303, 3, 12, 6, 15, 232, 233, 10,
		13, 0, 0, 233, 234, 5, 23, 0, 0, 234, 303, 3, 12, 6, 14, 235, 236, 10,
		12, 0, 0, 236, 237, 5, 22, 0, 0, 237, 303, 3, 12, 6, 13, 238, 239, 10,
		11, 0, 0, 239, 240, 5, 3, 0, 0, 240, 303, 3, 12, 6, 12, 241, 242, 10, 10,
		0, 0, 242, 243, 5, 1, 0, 0, 243, 303, 3, 12, 6, 11, 244, 245, 10, 9, 0,
		0, 245, 246, 5, 2, 0, 0, 246, 303, 3, 12, 6, 10, 247, 248, 10, 8, 0, 0,
		248, 249, 5, 4, 0, 0, 249, 303, 3, 12, 6, 9, 250, 251, 10, 7, 0, 0, 251,
		252, 5, 5, 0, 0, 252, 303, 3, 12, 6, 8, 253, 254, 10, 6, 0, 0, 254, 255,
		5, 6, 0, 0, 255, 303, 3, 12, 6, 7, 256, 257, 10, 5, 0, 0, 257, 258, 5,
		7, 0, 0, 258, 303, 3, 12, 6, 6, 259, 260, 10, 4, 0, 0, 260, 261, 5, 8,
		0, 0, 261, 303, 3, 12, 6, 5, 262, 263, 10, 2, 0, 0, 263, 264, 5, 20, 0,
		0, 264, 303, 3, 12, 6, 3, 265, 266, 10, 1, 0, 0, 266, 267, 5, 21, 0, 0,
		267, 303, 3, 12, 6, 2, 268, 270, 10, 20, 0, 0, 269, 271, 5, 19, 0, 0, 270,
		269, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 273,
		5, 30, 0, 0, 273, 303, 3, 26, 13, 0, 274, 276, 10, 19, 0, 0, 275, 277,
		5, 19, 0, 0, 276, 275, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 278, 1, 0,
		0, 0, 278, 279, 5, 16, 0, 0, 279, 280, 3, 28, 14, 0, 280, 281, 5, 17, 0,
		0, 281, 303, 1, 0, 0, 0, 282, 283, 10, 18, 0, 0, 283, 292, 5, 14, 0, 0,
		284, 289, 3, 10, 5, 0, 285, 286, 5, 26, 0, 0, 286, 288, 3, 10, 5, 0, 287,
		285, 1, 0, 0, 0, 288, 291, 1, 0, 0, 0, 289, 287, 1, 0, 0, 0, 289, 290,
		1, 0, 0, 0, 290, 293, 1, 0, 0, 0, 291, 289, 1, 0, 0, 0, 292, 284, 1, 0,
		0, 0, 292, 293, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0, 294, 303, 5, 15, 0, 0,
		295, 296, 10, 3, 0, 0, 296, 297, 5, 19, 0, 0, 297, 300, 3, 12, 6, 0, 298,
		299, 5, 28, 0, 0, 299, 301, 3, 12, 6, 0, 300, 298, 1, 0, 0, 0, 300, 301,
		1, 0, 0, 0, 301, 303, 1, 0, 0, 0, 302, 226, 1, 0, 0, 0, 302, 229, 1, 0,
		0, 0, 302, 232, 1, 0, 0, 0, 302, 235, 1, 0, 0, 0, 302, 238, 1, 0, 0, 0,
		302, 241, 1, 0, 0, 0, 302, 244, 1, 0, 0, 0, 302, 247, 1, 0, 0, 0, 302,
		250, 1, 0, 0, 0, 302, 253, 1, 0, 0, 0, 302, 256, 1, 0, 0, 0, 302, 259,
		1, 0, 0, 0, 302, 262, 1, 0, 0, 0, 302, 265, 1, 0, 0, 0, 302, 268, 1, 0,
		0, 0, 302, 274, 1, 0, 0, 0, 302, 282, 1, 0, 0, 0, 302, 295, 1, 0, 0, 0,
		303, 306, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305,
		13, 1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 307, 316, 5, 16, 0, 0, 308, 313,
		3, 16, 8, 0, 309, 310, 5, 26, 0, 0, 310, 312, 3, 16, 8, 0, 311, 309, 1,
		0, 0, 0, 312, 315, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 313, 314, 1, 0, 0,
		0, 314, 317, 1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 316, 308, 1, 0, 0, 0, 316,
		317, 1, 0, 0, 0, 317, 318, 1, 0, 0, 0, 318, 319, 5, 17, 0, 0, 319, 15,
		1, 0, 0, 0, 320, 322, 5, 24, 0, 0, 321, 320, 1, 0, 0, 0, 321, 322, 1, 0,
		0, 0, 322, 323, 1, 0, 0, 0, 323, 324, 3, 12, 6, 0, 324, 17, 1, 0, 0, 0,
		325, 334, 5, 31, 0, 0, 326, 331, 3, 20, 10, 0, 327, 328, 5, 26, 0, 0, 328,
		330, 3, 20, 10, 0, 329, 327, 1, 0, 0, 0, 330, 333, 1, 0, 0, 0, 331, 329,
		1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332, 335, 1, 0, 0, 0, 333, 331, 1, 0,
		0, 0, 334, 326, 1, 0, 0, 0, 334, 335, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0,
		336, 337, 5, 32, 0, 0, 337, 19, 1, 0, 0, 0, 338, 339, 3, 26, 13, 0, 339,
		340, 5, 28, 0, 0, 340, 341, 3, 12, 6, 0, 341, 350, 1, 0, 0, 0, 342, 343,
		5, 46, 0, 0, 343, 344, 5, 28, 0, 0, 344, 350, 3, 12, 6, 0, 345, 347, 5,
		24, 0, 0, 346, 345, 1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 347, 348, 1, 0, 0,
		0, 348, 350, 3, 12, 6, 0, 349, 338, 1, 0, 0, 0, 349, 342, 1, 0, 0, 0, 349,
		346, 1, 0, 0, 0, 350, 21, 1, 0, 0, 0, 351, 355, 5, 45, 0, 0, 352, 354,
		3, 24, 12, 0, 353, 352, 1, 0, 0, 0, 354, 357, 1, 0, 0, 0, 355, 353, 1,
		0, 0, 0, 355, 356, 1, 0, 0, 0, 356, 358, 1, 0, 0, 0, 357, 355, 1, 0, 0,
		0, 358, 359, 5, 54, 0, 0, 359, 23, 1, 0, 0, 0, 360, 366, 5, 52, 0, 0, 361,
		362, 5, 53, 0, 0, 362, 363, 3, 12, 6, 0, 363, 364, 5, 32, 0, 0, 364, 366,
		1, 0, 0, 0, 365, 360, 1, 0, 0, 0, 365, 361, 1, 0, 0, 0, 366, 25, 1, 0,
		0, 0, 367, 368, 5, 47, 0, 0, 368, 27, 1, 0, 0, 0, 369, 373, 3, 12, 6, 0,
		370, 373, 5, 48, 0, 0, 371, 373, 5, 46, 0, 0, 372, 369, 1, 0, 0, 0, 372,
		370, 1, 0, 0, 0, 372, 371, 1, 0, 0, 0, 373, 29, 1, 0, 0, 0, 43, 33, 42,
		53, 69, 74, 86, 92, 99, 107, 109, 113, 119, 125, 131, 137, 143, 148, 152,
		168, 173, 185, 190, 196, 199, 203, 224, 270, 276, 289, 292, 300, 302, 304,
		313, 316, 321, 331, 334, 346, 349, 355, 365, 372,
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

// JsonTemplateInit initializes any static state used to implement JsonTemplate. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJsonTemplate(). You can call this function if you wish to initialize the static state ahead
// of time.
func JsonTemplateInit() {
	staticData := &jsontemplateParserStaticData
	staticData.once.Do(jsontemplateParserInit)
}

// NewJsonTemplate produces a new parser instance for the optional input antlr.TokenStream.
func NewJsonTemplate(input antlr.TokenStream) *JsonTemplate {
	JsonTemplateInit()
	this := new(JsonTemplate)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &jsontemplateParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// JsonTemplate tokens.
const (
	JsonTemplateEOF                   = antlr.TokenEOF
	JsonTemplateLess                  = 1
	JsonTemplateLessOrEqual           = 2
	JsonTemplateEqual                 = 3
	JsonTemplateGreater               = 4
	JsonTemplateGreaterOrEqual        = 5
	JsonTemplateNotEqual              = 6
	JsonTemplateAnd                   = 7
	JsonTemplateOr                    = 8
	JsonTemplateNot                   = 9
	JsonTemplateAdd                   = 10
	JsonTemplateSubtract              = 11
	JsonTemplateMultiply              = 12
	JsonTemplateDivide                = 13
	JsonTemplateLeftParen             = 14
	JsonTemplateRightParen            = 15
	JsonTemplateLeftBracket           = 16
	JsonTemplateRightBracket          = 17
	JsonTemplateIteration             = 18
	JsonTemplateQuestion              = 19
	JsonTemplateAddAssign             = 20
	JsonTemplateLiteral               = 21
	JsonTemplateNullCoalescing        = 22
	JsonTemplateRange                 = 23
	JsonTemplateSpread                = 24
	JsonTemplateAs                    = 25
	JsonTemplateComma                 = 26
	JsonTemplateArrow                 = 27
	JsonTemplateColon                 = 28
	JsonTemplateSemicolon             = 29
	JsonTemplateDot                   = 30
	JsonTemplateLeftBrace             = 31
	JsonTemplateRightBrace            = 32
	JsonTemplateNull                  = 33
	JsonTemplateFalse                 = 34
	JsonTemplateTrue                  = 35
	JsonTemplateReturn                = 36
	JsonTemplateFor                   = 37
	JsonTemplateIn                    = 38
	JsonTemplateIf                    = 39
	JsonTemplateElse                  = 40
	JsonTemplateWhile                 = 41
	JsonTemplateDo                    = 42
	JsonTemplateBreak                 = 43
	JsonTemplateContinue              = 44
	JsonTemplateBacktick              = 45
	JsonTemplateESCAPED_STRING        = 46
	JsonTemplateSTRING                = 47
	JsonTemplateNUMBER                = 48
	JsonTemplateLINE_COMMENT          = 49
	JsonTemplateBLOCK_COMMENT         = 50
	JsonTemplateWS                    = 51
	JsonTemplateTEMPLATE_TEXT         = 52
	JsonTemplateTEMPLATE_INTERP_START = 53
	JsonTemplateTEMPLATE_END          = 54
)

// JsonTemplate rules.
const (
	JsonTemplateRULE_script               = 0
	JsonTemplateRULE_statement            = 1
	JsonTemplateRULE_statements           = 2
	JsonTemplateRULE_expression           = 3
	JsonTemplateRULE_lambda               = 4
	JsonTemplateRULE_function_param       = 5
	JsonTemplateRULE_field                = 6
	JsonTemplateRULE_array                = 7
	JsonTemplateRULE_spread_field         = 8
	JsonTemplateRULE_object               = 9
	JsonTemplateRULE_object_field         = 10
	JsonTemplateRULE_template_string      = 11
	JsonTemplateRULE_template_string_part = 12
	JsonTemplateRULE_name                 = 13
	JsonTemplateRULE_index                = 14
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
	p.RuleIndex = JsonTemplateRULE_script
	return p
}

func (*ScriptContext) IsScriptContext() {}

func NewScriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScriptContext {
	var p = new(ScriptContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateRULE_script

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

func (p *JsonTemplate) Script() (localctx IScriptContext) {
	this := p
	_ = this

	localctx = NewScriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JsonTemplateRULE_script)
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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&561569121520128) != 0 {
		{
			p.SetState(30)
			p.Statement()
		}

		p.SetState(35)
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
	p.RuleIndex = JsonTemplateRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

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

func (s *StatementContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(JsonTemplateSemicolon, 0)
}

func (s *StatementContext) Return() antlr.TerminalNode {
	return s.GetToken(JsonTemplateReturn, 0)
}

func (s *StatementContext) Break() antlr.TerminalNode {
	return s.GetToken(JsonTemplateBreak, 0)
}

func (s *StatementContext) Continue() antlr.TerminalNode {
	return s.GetToken(JsonTemplateContinue, 0)
}

func (s *StatementContext) For() antlr.TerminalNode {
	return s.GetToken(JsonTemplateFor, 0)
}

func (s *StatementContext) In() antlr.TerminalNode {
	return s.GetToken(JsonTemplateIn, 0)
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

func (s *StatementContext) Comma() antlr.TerminalNode {
	return s.GetToken(JsonTemplateComma, 0)
}

func (s *StatementContext) AllIf() []antlr.TerminalNode {
	return s.GetTokens(JsonTemplateIf)
}

func (s *StatementContext) If(i int) antlr.TerminalNode {
	return s.GetToken(JsonTemplateIf, i)
}

func (s *StatementContext) AllElse() []antlr.TerminalNode {
	return s.GetTokens(JsonTemplateElse)
}

func (s *StatementContext) Else(i int) antlr.TerminalNode {
	return s.GetToken(JsonTemplateElse, i)
}

func (s *StatementContext) While() antlr.TerminalNode {
	return s.GetToken(JsonTemplateWhile, 0)
}

func (s *StatementContext) Do() antlr.TerminalNode {
	return s.GetToken(JsonTemplateDo, 0)
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

func (p *JsonTemplate) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JsonTemplateRULE_statement)
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

	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(36)
			p.Statements()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.field(0)
		}
		{
			p.SetState(38)
			p.Match(JsonTemplateSemicolon)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(40)
			p.Match(JsonTemplateReturn)
		}
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&527827858442752) != 0 {
			{
				p.SetState(41)
				p.field(0)
			}

		}
		{
			p.SetState(44)
			p.Match(JsonTemplateSemicolon)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(45)
			p.Match(JsonTemplateBreak)
		}
		{
			p.SetState(46)
			p.Match(JsonTemplateSemicolon)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(47)
			p.Match(JsonTemplateContinue)
		}
		{
			p.SetState(48)
			p.Match(JsonTemplateSemicolon)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(49)
			p.Match(JsonTemplateFor)
		}

		{
			p.SetState(50)
			p.Name()
		}
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateComma {
			{
				p.SetState(51)
				p.Match(JsonTemplateComma)
			}
			{
				p.SetState(52)
				p.Name()
			}

		}

		{
			p.SetState(55)
			p.Match(JsonTemplateIn)
		}
		{
			p.SetState(56)
			p.field(0)
		}
		{
			p.SetState(57)
			p.Statements()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(59)
			p.Match(JsonTemplateIf)
		}
		{
			p.SetState(60)
			p.field(0)
		}
		{
			p.SetState(61)
			p.Statements()
		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(62)
					p.Match(JsonTemplateElse)
				}
				{
					p.SetState(63)
					p.Match(JsonTemplateIf)
				}
				{
					p.SetState(64)
					p.field(0)
				}
				{
					p.SetState(65)
					p.Statements()
				}

			}
			p.SetState(71)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateElse {
			{
				p.SetState(72)
				p.Match(JsonTemplateElse)
			}
			{
				p.SetState(73)
				p.Statements()
			}

		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(76)
			p.Match(JsonTemplateWhile)
		}
		{
			p.SetState(77)
			p.field(0)
		}
		{
			p.SetState(78)
			p.Statements()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(80)
			p.Match(JsonTemplateDo)
		}
		{
			p.SetState(81)
			p.Statements()
		}
		{
			p.SetState(82)
			p.Match(JsonTemplateWhile)
		}
		{
			p.SetState(83)
			p.field(0)
		}
		{
			p.SetState(84)
			p.Match(JsonTemplateSemicolon)
		}

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
	p.RuleIndex = JsonTemplateRULE_statements
	return p
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(JsonTemplateLeftBrace, 0)
}

func (s *StatementsContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(JsonTemplateRightBrace, 0)
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

func (p *JsonTemplate) Statements() (localctx IStatementsContext) {
	this := p
	_ = this

	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JsonTemplateRULE_statements)
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
		p.SetState(88)
		p.Match(JsonTemplateLeftBrace)
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&561569121520128) != 0 {
		{
			p.SetState(89)
			p.Statement()
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(95)
		p.Match(JsonTemplateRightBrace)
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
	p.RuleIndex = JsonTemplateRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Iteration() antlr.TerminalNode {
	return s.GetToken(JsonTemplateIteration, 0)
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
	return s.GetToken(JsonTemplateEOF, 0)
}

func (s *ExpressionContext) AllLeftBrace() []antlr.TerminalNode {
	return s.GetTokens(JsonTemplateLeftBrace)
}

func (s *ExpressionContext) LeftBrace(i int) antlr.TerminalNode {
	return s.GetToken(JsonTemplateLeftBrace, i)
}

func (s *ExpressionContext) As() antlr.TerminalNode {
	return s.GetToken(JsonTemplateAs, 0)
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
	return s.GetTokens(JsonTemplateRightBrace)
}

func (s *ExpressionContext) RightBrace(i int) antlr.TerminalNode {
	return s.GetToken(JsonTemplateRightBrace, i)
}

func (s *ExpressionContext) Comma() antlr.TerminalNode {
	return s.GetToken(JsonTemplateComma, 0)
}

func (s *ExpressionContext) Question() antlr.TerminalNode {
	return s.GetToken(JsonTemplateQuestion, 0)
}

func (s *ExpressionContext) Literal() antlr.TerminalNode {
	return s.GetToken(JsonTemplateLiteral, 0)
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

func (p *JsonTemplate) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JsonTemplateRULE_expression)
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

	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateLeftBrace {
			{
				p.SetState(97)
				p.Match(JsonTemplateLeftBrace)
			}
			{
				p.SetState(98)
				p.Match(JsonTemplateLeftBrace)
			}

		}
		{
			p.SetState(101)
			p.Match(JsonTemplateIteration)
		}
		{
			p.SetState(102)
			p.field(0)
		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateAs {
			{
				p.SetState(103)
				p.Match(JsonTemplateAs)
			}
			{
				p.SetState(104)
				p.Name()
			}
			p.SetState(107)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == JsonTemplateComma {
				{
					p.SetState(105)
					p.Match(JsonTemplateComma)
				}
				{
					p.SetState(106)
					p.Name()
				}

			}

		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateRightBrace {
			{
				p.SetState(111)
				p.Match(JsonTemplateRightBrace)
			}
			{
				p.SetState(112)
				p.Match(JsonTemplateRightBrace)
			}

		}
		{
			p.SetState(115)
			p.Match(JsonTemplateEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateLeftBrace {
			{
				p.SetState(117)
				p.Match(JsonTemplateLeftBrace)
			}
			{
				p.SetState(118)
				p.Match(JsonTemplateLeftBrace)
			}

		}
		{
			p.SetState(121)
			p.Match(JsonTemplateQuestion)
		}
		{
			p.SetState(122)
			p.field(0)
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateRightBrace {
			{
				p.SetState(123)
				p.Match(JsonTemplateRightBrace)
			}
			{
				p.SetState(124)
				p.Match(JsonTemplateRightBrace)
			}

		}
		{
			p.SetState(127)
			p.Match(JsonTemplateEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateLeftBrace {
			{
				p.SetState(129)
				p.Match(JsonTemplateLeftBrace)
			}
			{
				p.SetState(130)
				p.Match(JsonTemplateLeftBrace)
			}

		}
		{
			p.SetState(133)
			p.Match(JsonTemplateLiteral)
		}
		{
			p.SetState(134)
			p.field(0)
		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateRightBrace {
			{
				p.SetState(135)
				p.Match(JsonTemplateRightBrace)
			}
			{
				p.SetState(136)
				p.Match(JsonTemplateRightBrace)
			}

		}
		{
			p.SetState(139)
			p.Match(JsonTemplateEOF)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(143)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(141)
				p.Match(JsonTemplateLeftBrace)
			}
			{
				p.SetState(142)
				p.Match(JsonTemplateLeftBrace)
			}

		}
		{
			p.SetState(145)
			p.field(0)
		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateRightBrace {
			{
				p.SetState(146)
				p.Match(JsonTemplateRightBrace)
			}
			{
				p.SetState(147)
				p.Match(JsonTemplateRightBrace)
			}

		}
		{
			p.SetState(150)
			p.Match(JsonTemplateEOF)
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
	p.RuleIndex = JsonTemplateRULE_lambda
	return p
}

func (*LambdaContext) IsLambdaContext() {}

func NewLambdaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaContext {
	var p = new(LambdaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateRULE_lambda

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
	return s.GetToken(JsonTemplateArrow, 0)
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

func (s *LambdaContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *LambdaContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(JsonTemplateLeftParen, 0)
}

func (s *LambdaContext) RightParen() antlr.TerminalNode {
	return s.GetToken(JsonTemplateRightParen, 0)
}

func (s *LambdaContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(JsonTemplateComma)
}

func (s *LambdaContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(JsonTemplateComma, i)
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

func (p *JsonTemplate) Lambda() (localctx ILambdaContext) {
	this := p
	_ = this

	localctx = NewLambdaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JsonTemplateRULE_lambda)
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

	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(154)
			p.Name()
		}
		{
			p.SetState(155)
			p.Match(JsonTemplateArrow)
		}
		{
			p.SetState(156)
			p.field(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(158)
			p.Name()
		}
		{
			p.SetState(159)
			p.Match(JsonTemplateArrow)
		}
		{
			p.SetState(160)
			p.Statements()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(162)
			p.Match(JsonTemplateLeftParen)
		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonTemplateSTRING {
			{
				p.SetState(163)
				p.Name()
			}
			p.SetState(168)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == JsonTemplateComma {
				{
					p.SetState(164)
					p.Match(JsonTemplateComma)
				}
				{
					p.SetState(165)
					p.Name()
				}

				p.SetState(170)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

			p.SetState(175)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(176)
			p.Match(JsonTemplateRightParen)
		}
		{
			p.SetState(177)
			p.Match(JsonTemplateArrow)
		}
		{
			p.SetState(178)
			p.field(0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(179)
			p.Match(JsonTemplateLeftParen)
		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonTemplateSTRING {
			{
				p.SetState(180)
				p.Name()
			}
			p.SetState(185)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == JsonTemplateComma {
				{
					p.SetState(181)
					p.Match(JsonTemplateComma)
				}
				{
					p.SetState(182)
					p.Name()
				}

				p.SetState(187)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

			p.SetState(192)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(193)
			p.Match(JsonTemplateRightParen)
		}
		{
			p.SetState(194)
			p.Match(JsonTemplateArrow)
		}
		{
			p.SetState(195)
			p.Statements()
		}

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
	p.RuleIndex = JsonTemplateRULE_function_param
	return p
}

func (*Function_paramContext) IsFunction_paramContext() {}

func NewFunction_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_paramContext {
	var p = new(Function_paramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateRULE_function_param

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
	return s.GetToken(JsonTemplateSpread, 0)
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

func (p *JsonTemplate) Function_param() (localctx IFunction_paramContext) {
	this := p
	_ = this

	localctx = NewFunction_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JsonTemplateRULE_function_param)
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

	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateSpread {
			{
				p.SetState(198)
				p.Match(JsonTemplateSpread)
			}

		}
		{
			p.SetState(201)
			p.field(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(202)
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
	p.RuleIndex = JsonTemplateRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(JsonTemplateLeftParen, 0)
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
	return s.GetToken(JsonTemplateRightParen, 0)
}

func (s *FieldContext) True() antlr.TerminalNode {
	return s.GetToken(JsonTemplateTrue, 0)
}

func (s *FieldContext) False() antlr.TerminalNode {
	return s.GetToken(JsonTemplateFalse, 0)
}

func (s *FieldContext) Null() antlr.TerminalNode {
	return s.GetToken(JsonTemplateNull, 0)
}

func (s *FieldContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(JsonTemplateNUMBER, 0)
}

func (s *FieldContext) ESCAPED_STRING() antlr.TerminalNode {
	return s.GetToken(JsonTemplateESCAPED_STRING, 0)
}

func (s *FieldContext) Template_string() ITemplate_stringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITemplate_stringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITemplate_stringContext)
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

func (s *FieldContext) Lambda() ILambdaContext {
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
	return s.GetToken(JsonTemplateSubtract, 0)
}

func (s *FieldContext) Not() antlr.TerminalNode {
	return s.GetToken(JsonTemplateNot, 0)
}

func (s *FieldContext) Divide() antlr.TerminalNode {
	return s.GetToken(JsonTemplateDivide, 0)
}

func (s *FieldContext) Multiply() antlr.TerminalNode {
	return s.GetToken(JsonTemplateMultiply, 0)
}

func (s *FieldContext) Add() antlr.TerminalNode {
	return s.GetToken(JsonTemplateAdd, 0)
}

func (s *FieldContext) Range() antlr.TerminalNode {
	return s.GetToken(JsonTemplateRange, 0)
}

func (s *FieldContext) NullCoalescing() antlr.TerminalNode {
	return s.GetToken(JsonTemplateNullCoalescing, 0)
}

func (s *FieldContext) Equal() antlr.TerminalNode {
	return s.GetToken(JsonTemplateEqual, 0)
}

func (s *FieldContext) Less() antlr.TerminalNode {
	return s.GetToken(JsonTemplateLess, 0)
}

func (s *FieldContext) LessOrEqual() antlr.TerminalNode {
	return s.GetToken(JsonTemplateLessOrEqual, 0)
}

func (s *FieldContext) Greater() antlr.TerminalNode {
	return s.GetToken(JsonTemplateGreater, 0)
}

func (s *FieldContext) GreaterOrEqual() antlr.TerminalNode {
	return s.GetToken(JsonTemplateGreaterOrEqual, 0)
}

func (s *FieldContext) NotEqual() antlr.TerminalNode {
	return s.GetToken(JsonTemplateNotEqual, 0)
}

func (s *FieldContext) And() antlr.TerminalNode {
	return s.GetToken(JsonTemplateAnd, 0)
}

func (s *FieldContext) Or() antlr.TerminalNode {
	return s.GetToken(JsonTemplateOr, 0)
}

func (s *FieldContext) AddAssign() antlr.TerminalNode {
	return s.GetToken(JsonTemplateAddAssign, 0)
}

func (s *FieldContext) Literal() antlr.TerminalNode {
	return s.GetToken(JsonTemplateLiteral, 0)
}

func (s *FieldContext) Dot() antlr.TerminalNode {
	return s.GetToken(JsonTemplateDot, 0)
}

func (s *FieldContext) Question() antlr.TerminalNode {
	return s.GetToken(JsonTemplateQuestion, 0)
}

func (s *FieldContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(JsonTemplateLeftBracket, 0)
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
	return s.GetToken(JsonTemplateRightBracket, 0)
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
	return s.GetTokens(JsonTemplateComma)
}

func (s *FieldContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(JsonTemplateComma, i)
}

func (s *FieldContext) Colon() antlr.TerminalNode {
	return s.GetToken(JsonTemplateColon, 0)
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

func (p *JsonTemplate) Field() (localctx IFieldContext) {
	return p.field(0)
}

func (p *JsonTemplate) field(_p int) (localctx IFieldContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewFieldContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFieldContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, JsonTemplateRULE_field, _p)
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
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(206)
			p.Match(JsonTemplateLeftParen)
		}
		{
			p.SetState(207)
			p.field(0)
		}
		{
			p.SetState(208)
			p.Match(JsonTemplateRightParen)
		}

	case 2:
		{
			p.SetState(210)
			p.Match(JsonTemplateTrue)
		}

	case 3:
		{
			p.SetState(211)
			p.Match(JsonTemplateFalse)
		}

	case 4:
		{
			p.SetState(212)
			p.Match(JsonTemplateNull)
		}

	case 5:
		{
			p.SetState(213)
			p.Match(JsonTemplateNUMBER)
		}

	case 6:
		{
			p.SetState(214)
			p.Match(JsonTemplateESCAPED_STRING)
		}

	case 7:
		{
			p.SetState(215)
			p.Template_string()
		}

	case 8:
		{
			p.SetState(216)
			p.Array()
		}

	case 9:
		{
			p.SetState(217)
			p.Object()
		}

	case 10:
		{
			p.SetState(218)
			p.Lambda()
		}

	case 11:
		{
			p.SetState(219)
			p.Name()
		}

	case 12:
		{
			p.SetState(220)
			p.Match(JsonTemplateSubtract)
		}
		{
			p.SetState(221)
			p.field(17)
		}

	case 13:
		{
			p.SetState(222)
			p.Match(JsonTemplateNot)
		}
		{
			p.SetState(223)
			p.field(16)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(302)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
			case 1:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(226)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(227)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JsonTemplateMultiply || _la == JsonTemplateDivide) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(228)
					p.field(16)
				}

			case 2:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(229)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(230)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JsonTemplateAdd || _la == JsonTemplateSubtract) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(231)
					p.field(15)
				}

			case 3:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(232)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(233)
					p.Match(JsonTemplateRange)
				}
				{
					p.SetState(234)
					p.field(14)
				}

			case 4:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(235)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(236)
					p.Match(JsonTemplateNullCoalescing)
				}
				{
					p.SetState(237)
					p.field(13)
				}

			case 5:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(238)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(239)
					p.Match(JsonTemplateEqual)
				}
				{
					p.SetState(240)
					p.field(12)
				}

			case 6:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(241)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(242)
					p.Match(JsonTemplateLess)
				}
				{
					p.SetState(243)
					p.field(11)
				}

			case 7:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(244)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(245)
					p.Match(JsonTemplateLessOrEqual)
				}
				{
					p.SetState(246)
					p.field(10)
				}

			case 8:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(247)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(248)
					p.Match(JsonTemplateGreater)
				}
				{
					p.SetState(249)
					p.field(9)
				}

			case 9:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(250)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(251)
					p.Match(JsonTemplateGreaterOrEqual)
				}
				{
					p.SetState(252)
					p.field(8)
				}

			case 10:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(253)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(254)
					p.Match(JsonTemplateNotEqual)
				}
				{
					p.SetState(255)
					p.field(7)
				}

			case 11:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(256)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(257)
					p.Match(JsonTemplateAnd)
				}
				{
					p.SetState(258)
					p.field(6)
				}

			case 12:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(259)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(260)
					p.Match(JsonTemplateOr)
				}
				{
					p.SetState(261)
					p.field(5)
				}

			case 13:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(262)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(263)
					p.Match(JsonTemplateAddAssign)
				}
				{
					p.SetState(264)
					p.field(3)
				}

			case 14:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(265)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(266)
					p.Match(JsonTemplateLiteral)
				}
				{
					p.SetState(267)
					p.field(2)
				}

			case 15:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(268)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
				}

				p.SetState(270)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonTemplateQuestion {
					{
						p.SetState(269)
						p.Match(JsonTemplateQuestion)
					}

				}
				{
					p.SetState(272)
					p.Match(JsonTemplateDot)
				}
				{
					p.SetState(273)
					p.Name()
				}

			case 16:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(274)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}

				p.SetState(276)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonTemplateQuestion {
					{
						p.SetState(275)
						p.Match(JsonTemplateQuestion)
					}

				}
				{
					p.SetState(278)
					p.Match(JsonTemplateLeftBracket)
				}
				{
					p.SetState(279)
					p.Index()
				}
				{
					p.SetState(280)
					p.Match(JsonTemplateRightBracket)
				}

			case 17:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(282)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(283)
					p.Match(JsonTemplateLeftParen)
				}
				p.SetState(292)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&527827875219968) != 0 {
					{
						p.SetState(284)
						p.Function_param()
					}
					p.SetState(289)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == JsonTemplateComma {
						{
							p.SetState(285)
							p.Match(JsonTemplateComma)
						}
						{
							p.SetState(286)
							p.Function_param()
						}

						p.SetState(291)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(294)
					p.Match(JsonTemplateRightParen)
				}

			case 18:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(295)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(296)
					p.Match(JsonTemplateQuestion)
				}
				{
					p.SetState(297)
					p.field(0)
				}
				p.SetState(300)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(298)
						p.Match(JsonTemplateColon)
					}
					{
						p.SetState(299)
						p.field(0)
					}

				}

			}

		}
		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())
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
	p.RuleIndex = JsonTemplateRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(JsonTemplateLeftBracket, 0)
}

func (s *ArrayContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(JsonTemplateRightBracket, 0)
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
	return s.GetTokens(JsonTemplateComma)
}

func (s *ArrayContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(JsonTemplateComma, i)
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

func (p *JsonTemplate) Array() (localctx IArrayContext) {
	this := p
	_ = this

	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JsonTemplateRULE_array)
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
		p.SetState(307)
		p.Match(JsonTemplateLeftBracket)
	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&527827875219968) != 0 {
		{
			p.SetState(308)
			p.Spread_field()
		}
		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonTemplateComma {
			{
				p.SetState(309)
				p.Match(JsonTemplateComma)
			}
			{
				p.SetState(310)
				p.Spread_field()
			}

			p.SetState(315)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(318)
		p.Match(JsonTemplateRightBracket)
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
	p.RuleIndex = JsonTemplateRULE_spread_field
	return p
}

func (*Spread_fieldContext) IsSpread_fieldContext() {}

func NewSpread_fieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Spread_fieldContext {
	var p = new(Spread_fieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateRULE_spread_field

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
	return s.GetToken(JsonTemplateSpread, 0)
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

func (p *JsonTemplate) Spread_field() (localctx ISpread_fieldContext) {
	this := p
	_ = this

	localctx = NewSpread_fieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JsonTemplateRULE_spread_field)
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
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JsonTemplateSpread {
		{
			p.SetState(320)
			p.Match(JsonTemplateSpread)
		}

	}
	{
		p.SetState(323)
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
	p.RuleIndex = JsonTemplateRULE_object
	return p
}

func (*ObjectContext) IsObjectContext() {}

func NewObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectContext {
	var p = new(ObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateRULE_object

	return p
}

func (s *ObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(JsonTemplateLeftBrace, 0)
}

func (s *ObjectContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(JsonTemplateRightBrace, 0)
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
	return s.GetTokens(JsonTemplateComma)
}

func (s *ObjectContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(JsonTemplateComma, i)
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

func (p *JsonTemplate) Object() (localctx IObjectContext) {
	this := p
	_ = this

	localctx = NewObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JsonTemplateRULE_object)
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
		p.SetState(325)
		p.Match(JsonTemplateLeftBrace)
	}
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&527827875219968) != 0 {
		{
			p.SetState(326)
			p.Object_field()
		}
		p.SetState(331)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonTemplateComma {
			{
				p.SetState(327)
				p.Match(JsonTemplateComma)
			}
			{
				p.SetState(328)
				p.Object_field()
			}

			p.SetState(333)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(336)
		p.Match(JsonTemplateRightBrace)
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
	p.RuleIndex = JsonTemplateRULE_object_field
	return p
}

func (*Object_fieldContext) IsObject_fieldContext() {}

func NewObject_fieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_fieldContext {
	var p = new(Object_fieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateRULE_object_field

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

func (s *Object_fieldContext) Colon() antlr.TerminalNode {
	return s.GetToken(JsonTemplateColon, 0)
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
	return s.GetToken(JsonTemplateESCAPED_STRING, 0)
}

func (s *Object_fieldContext) Spread() antlr.TerminalNode {
	return s.GetToken(JsonTemplateSpread, 0)
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

func (p *JsonTemplate) Object_field() (localctx IObject_fieldContext) {
	this := p
	_ = this

	localctx = NewObject_fieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, JsonTemplateRULE_object_field)
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

	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(338)
			p.Name()
		}
		{
			p.SetState(339)
			p.Match(JsonTemplateColon)
		}
		{
			p.SetState(340)
			p.field(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(342)
			p.Match(JsonTemplateESCAPED_STRING)
		}
		{
			p.SetState(343)
			p.Match(JsonTemplateColon)
		}
		{
			p.SetState(344)
			p.field(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateSpread {
			{
				p.SetState(345)
				p.Match(JsonTemplateSpread)
			}

		}
		{
			p.SetState(348)
			p.field(0)
		}

	}

	return localctx
}

// ITemplate_stringContext is an interface to support dynamic dispatch.
type ITemplate_stringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemplate_stringContext differentiates from other interfaces.
	IsTemplate_stringContext()
}

type Template_stringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplate_stringContext() *Template_stringContext {
	var p = new(Template_stringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonTemplateRULE_template_string
	return p
}

func (*Template_stringContext) IsTemplate_stringContext() {}

func NewTemplate_stringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Template_stringContext {
	var p = new(Template_stringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateRULE_template_string

	return p
}

func (s *Template_stringContext) GetParser() antlr.Parser { return s.parser }

func (s *Template_stringContext) Backtick() antlr.TerminalNode {
	return s.GetToken(JsonTemplateBacktick, 0)
}

func (s *Template_stringContext) TEMPLATE_END() antlr.TerminalNode {
	return s.GetToken(JsonTemplateTEMPLATE_END, 0)
}

func (s *Template_stringContext) AllTemplate_string_part() []ITemplate_string_partContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITemplate_string_partContext); ok {
			len++
		}
	}

	tst := make([]ITemplate_string_partContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITemplate_string_partContext); ok {
			tst[i] = t.(ITemplate_string_partContext)
			i++
		}
	}

	return tst
}

func (s *Template_stringContext) Template_string_part(i int) ITemplate_string_partContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITemplate_string_partContext); ok {
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

	return t.(ITemplate_string_partContext)
}

func (s *Template_stringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Template_stringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Template_stringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.EnterTemplate_string(s)
	}
}

func (s *Template_stringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.ExitTemplate_string(s)
	}
}

func (p *JsonTemplate) Template_string() (localctx ITemplate_stringContext) {
	this := p
	_ = this

	localctx = NewTemplate_stringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, JsonTemplateRULE_template_string)
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
		p.SetState(351)
		p.Match(JsonTemplateBacktick)
	}
	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JsonTemplateTEMPLATE_TEXT || _la == JsonTemplateTEMPLATE_INTERP_START {
		{
			p.SetState(352)
			p.Template_string_part()
		}

		p.SetState(357)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(358)
		p.Match(JsonTemplateTEMPLATE_END)
	}

	return localctx
}

// ITemplate_string_partContext is an interface to support dynamic dispatch.
type ITemplate_string_partContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemplate_string_partContext differentiates from other interfaces.
	IsTemplate_string_partContext()
}

type Template_string_partContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplate_string_partContext() *Template_string_partContext {
	var p = new(Template_string_partContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonTemplateRULE_template_string_part
	return p
}

func (*Template_string_partContext) IsTemplate_string_partContext() {}

func NewTemplate_string_partContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Template_string_partContext {
	var p = new(Template_string_partContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateRULE_template_string_part

	return p
}

func (s *Template_string_partContext) GetParser() antlr.Parser { return s.parser }

func (s *Template_string_partContext) TEMPLATE_TEXT() antlr.TerminalNode {
	return s.GetToken(JsonTemplateTEMPLATE_TEXT, 0)
}

func (s *Template_string_partContext) TEMPLATE_INTERP_START() antlr.TerminalNode {
	return s.GetToken(JsonTemplateTEMPLATE_INTERP_START, 0)
}

func (s *Template_string_partContext) Field() IFieldContext {
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

func (s *Template_string_partContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(JsonTemplateRightBrace, 0)
}

func (s *Template_string_partContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Template_string_partContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Template_string_partContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.EnterTemplate_string_part(s)
	}
}

func (s *Template_string_partContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonTemplateListener); ok {
		listenerT.ExitTemplate_string_part(s)
	}
}

func (p *JsonTemplate) Template_string_part() (localctx ITemplate_string_partContext) {
	this := p
	_ = this

	localctx = NewTemplate_string_partContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, JsonTemplateRULE_template_string_part)

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

	p.SetState(365)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonTemplateTEMPLATE_TEXT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(360)
			p.Match(JsonTemplateTEMPLATE_TEXT)
		}

	case JsonTemplateTEMPLATE_INTERP_START:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(361)
			p.Match(JsonTemplateTEMPLATE_INTERP_START)
		}
		{
			p.SetState(362)
			p.field(0)
		}
		{
			p.SetState(363)
			p.Match(JsonTemplateRightBrace)
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
	p.RuleIndex = JsonTemplateRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) STRING() antlr.TerminalNode {
	return s.GetToken(JsonTemplateSTRING, 0)
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

func (p *JsonTemplate) Name() (localctx INameContext) {
	this := p
	_ = this

	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, JsonTemplateRULE_name)

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
		p.SetState(367)
		p.Match(JsonTemplateSTRING)
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
	p.RuleIndex = JsonTemplateRULE_index
	return p
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonTemplateRULE_index

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
	return s.GetToken(JsonTemplateNUMBER, 0)
}

func (s *IndexContext) ESCAPED_STRING() antlr.TerminalNode {
	return s.GetToken(JsonTemplateESCAPED_STRING, 0)
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

func (p *JsonTemplate) Index() (localctx IIndexContext) {
	this := p
	_ = this

	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, JsonTemplateRULE_index)

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

	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(369)
			p.field(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(370)
			p.Match(JsonTemplateNUMBER)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(371)
			p.Match(JsonTemplateESCAPED_STRING)
		}

	}

	return localctx
}

func (p *JsonTemplate) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
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

func (p *JsonTemplate) Field_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 16:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 17:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
