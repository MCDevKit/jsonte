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
		"'+='", "'-='", "'*='", "'/='", "'%='", "'++'", "'--'", "'='", "'??'",
		"'..'", "'...'", "'as'", "','", "'=>'", "':'", "';'", "'.'", "'{'",
		"'}'", "'null'", "'false'", "'true'", "'return'", "'for'", "'in'", "'if'",
		"'else'", "'while'", "'do'", "'break'", "'continue'", "", "", "", "",
		"", "", "", "", "'${'",
	}
	staticData.symbolicNames = []string{
		"", "Less", "LessOrEqual", "Equal", "Greater", "GreaterOrEqual", "NotEqual",
		"And", "Or", "Not", "Add", "Subtract", "Multiply", "Divide", "LeftParen",
		"RightParen", "LeftBracket", "RightBracket", "Iteration", "Question",
		"AddAssign", "SubtractAssign", "MultiplyAssign", "DivideAssign", "ModuloAssign",
		"Increment", "Decrement", "Literal", "NullCoalescing", "Range", "Spread",
		"As", "Comma", "Arrow", "Colon", "Semicolon", "Dot", "LeftBrace", "RightBrace",
		"Null", "False", "True", "Return", "For", "In", "If", "Else", "While",
		"Do", "Break", "Continue", "Backtick", "ESCAPED_STRING", "STRING", "NUMBER",
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
		4, 1, 60, 375, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
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
		6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 227, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 3, 6, 261, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 267,
		8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 278,
		8, 6, 10, 6, 12, 6, 281, 9, 6, 3, 6, 283, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 295, 8, 6, 5, 6, 297, 8, 6, 10,
		6, 12, 6, 300, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 306, 8, 7, 10, 7, 12,
		7, 309, 9, 7, 3, 7, 311, 8, 7, 1, 7, 1, 7, 1, 8, 3, 8, 316, 8, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 324, 8, 9, 10, 9, 12, 9, 327, 9, 9,
		3, 9, 329, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 347, 8,
		10, 1, 10, 3, 10, 350, 8, 10, 1, 11, 1, 11, 5, 11, 354, 8, 11, 10, 11,
		12, 11, 357, 9, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3,
		12, 366, 8, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 3, 14, 373, 8, 14, 1,
		14, 0, 1, 12, 15, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		0, 5, 1, 0, 12, 13, 1, 0, 10, 11, 3, 0, 1, 2, 4, 5, 44, 44, 2, 0, 3, 3,
		6, 6, 1, 0, 20, 24, 442, 0, 33, 1, 0, 0, 0, 2, 86, 1, 0, 0, 0, 4, 88, 1,
		0, 0, 0, 6, 152, 1, 0, 0, 0, 8, 196, 1, 0, 0, 0, 10, 203, 1, 0, 0, 0, 12,
		226, 1, 0, 0, 0, 14, 301, 1, 0, 0, 0, 16, 315, 1, 0, 0, 0, 18, 319, 1,
		0, 0, 0, 20, 349, 1, 0, 0, 0, 22, 351, 1, 0, 0, 0, 24, 365, 1, 0, 0, 0,
		26, 367, 1, 0, 0, 0, 28, 372, 1, 0, 0, 0, 30, 32, 3, 2, 1, 0, 31, 30, 1,
		0, 0, 0, 32, 35, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34,
		1, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 36, 87, 3, 4, 2, 0, 37, 38, 3, 12, 6,
		0, 38, 39, 5, 35, 0, 0, 39, 87, 1, 0, 0, 0, 40, 42, 5, 42, 0, 0, 41, 43,
		3, 12, 6, 0, 42, 41, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0,
		44, 87, 5, 35, 0, 0, 45, 46, 5, 49, 0, 0, 46, 87, 5, 35, 0, 0, 47, 48,
		5, 50, 0, 0, 48, 87, 5, 35, 0, 0, 49, 50, 5, 43, 0, 0, 50, 53, 3, 26, 13,
		0, 51, 52, 5, 32, 0, 0, 52, 54, 3, 26, 13, 0, 53, 51, 1, 0, 0, 0, 53, 54,
		1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 56, 5, 44, 0, 0, 56, 57, 3, 12, 6,
		0, 57, 58, 3, 2, 1, 0, 58, 87, 1, 0, 0, 0, 59, 60, 5, 45, 0, 0, 60, 61,
		3, 12, 6, 0, 61, 69, 3, 2, 1, 0, 62, 63, 5, 46, 0, 0, 63, 64, 5, 45, 0,
		0, 64, 65, 3, 12, 6, 0, 65, 66, 3, 2, 1, 0, 66, 68, 1, 0, 0, 0, 67, 62,
		1, 0, 0, 0, 68, 71, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0,
		70, 74, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 73, 5, 46, 0, 0, 73, 75, 3,
		2, 1, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 87, 1, 0, 0, 0, 76,
		77, 5, 47, 0, 0, 77, 78, 3, 12, 6, 0, 78, 79, 3, 2, 1, 0, 79, 87, 1, 0,
		0, 0, 80, 81, 5, 48, 0, 0, 81, 82, 3, 2, 1, 0, 82, 83, 5, 47, 0, 0, 83,
		84, 3, 12, 6, 0, 84, 85, 5, 35, 0, 0, 85, 87, 1, 0, 0, 0, 86, 36, 1, 0,
		0, 0, 86, 37, 1, 0, 0, 0, 86, 40, 1, 0, 0, 0, 86, 45, 1, 0, 0, 0, 86, 47,
		1, 0, 0, 0, 86, 49, 1, 0, 0, 0, 86, 59, 1, 0, 0, 0, 86, 76, 1, 0, 0, 0,
		86, 80, 1, 0, 0, 0, 87, 3, 1, 0, 0, 0, 88, 92, 5, 37, 0, 0, 89, 91, 3,
		2, 1, 0, 90, 89, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92,
		93, 1, 0, 0, 0, 93, 95, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 95, 96, 5, 38,
		0, 0, 96, 5, 1, 0, 0, 0, 97, 98, 5, 37, 0, 0, 98, 100, 5, 37, 0, 0, 99,
		97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 102, 5,
		18, 0, 0, 102, 109, 3, 12, 6, 0, 103, 104, 5, 31, 0, 0, 104, 107, 3, 26,
		13, 0, 105, 106, 5, 32, 0, 0, 106, 108, 3, 26, 13, 0, 107, 105, 1, 0, 0,
		0, 107, 108, 1, 0, 0, 0, 108, 110, 1, 0, 0, 0, 109, 103, 1, 0, 0, 0, 109,
		110, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 112, 5, 38, 0, 0, 112, 114,
		5, 38, 0, 0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 115, 1, 0,
		0, 0, 115, 116, 5, 0, 0, 1, 116, 153, 1, 0, 0, 0, 117, 118, 5, 37, 0, 0,
		118, 120, 5, 37, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120,
		121, 1, 0, 0, 0, 121, 122, 5, 19, 0, 0, 122, 125, 3, 12, 6, 0, 123, 124,
		5, 38, 0, 0, 124, 126, 5, 38, 0, 0, 125, 123, 1, 0, 0, 0, 125, 126, 1,
		0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128, 5, 0, 0, 1, 128, 153, 1, 0, 0,
		0, 129, 130, 5, 37, 0, 0, 130, 132, 5, 37, 0, 0, 131, 129, 1, 0, 0, 0,
		131, 132, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 5, 27, 0, 0, 134,
		137, 3, 12, 6, 0, 135, 136, 5, 38, 0, 0, 136, 138, 5, 38, 0, 0, 137, 135,
		1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 140, 5, 0,
		0, 1, 140, 153, 1, 0, 0, 0, 141, 142, 5, 37, 0, 0, 142, 144, 5, 37, 0,
		0, 143, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145,
		148, 3, 12, 6, 0, 146, 147, 5, 38, 0, 0, 147, 149, 5, 38, 0, 0, 148, 146,
		1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 151, 5, 0,
		0, 1, 151, 153, 1, 0, 0, 0, 152, 99, 1, 0, 0, 0, 152, 119, 1, 0, 0, 0,
		152, 131, 1, 0, 0, 0, 152, 143, 1, 0, 0, 0, 153, 7, 1, 0, 0, 0, 154, 155,
		3, 26, 13, 0, 155, 156, 5, 33, 0, 0, 156, 157, 3, 12, 6, 0, 157, 197, 1,
		0, 0, 0, 158, 159, 3, 26, 13, 0, 159, 160, 5, 33, 0, 0, 160, 161, 3, 4,
		2, 0, 161, 197, 1, 0, 0, 0, 162, 173, 5, 14, 0, 0, 163, 168, 3, 26, 13,
		0, 164, 165, 5, 32, 0, 0, 165, 167, 3, 26, 13, 0, 166, 164, 1, 0, 0, 0,
		167, 170, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169,
		172, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 171, 163, 1, 0, 0, 0, 172, 175,
		1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 176, 1, 0,
		0, 0, 175, 173, 1, 0, 0, 0, 176, 177, 5, 15, 0, 0, 177, 178, 5, 33, 0,
		0, 178, 197, 3, 12, 6, 0, 179, 190, 5, 14, 0, 0, 180, 185, 3, 26, 13, 0,
		181, 182, 5, 32, 0, 0, 182, 184, 3, 26, 13, 0, 183, 181, 1, 0, 0, 0, 184,
		187, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 189,
		1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 188, 180, 1, 0, 0, 0, 189, 192, 1, 0,
		0, 0, 190, 188, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 193, 1, 0, 0, 0,
		192, 190, 1, 0, 0, 0, 193, 194, 5, 15, 0, 0, 194, 195, 5, 33, 0, 0, 195,
		197, 3, 4, 2, 0, 196, 154, 1, 0, 0, 0, 196, 158, 1, 0, 0, 0, 196, 162,
		1, 0, 0, 0, 196, 179, 1, 0, 0, 0, 197, 9, 1, 0, 0, 0, 198, 200, 5, 30,
		0, 0, 199, 198, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0,
		201, 204, 3, 12, 6, 0, 202, 204, 3, 8, 4, 0, 203, 199, 1, 0, 0, 0, 203,
		202, 1, 0, 0, 0, 204, 11, 1, 0, 0, 0, 205, 206, 6, 6, -1, 0, 206, 207,
		5, 14, 0, 0, 207, 208, 3, 12, 6, 0, 208, 209, 5, 15, 0, 0, 209, 227, 1,
		0, 0, 0, 210, 227, 5, 41, 0, 0, 211, 227, 5, 40, 0, 0, 212, 227, 5, 39,
		0, 0, 213, 227, 5, 54, 0, 0, 214, 227, 5, 52, 0, 0, 215, 227, 3, 22, 11,
		0, 216, 227, 3, 14, 7, 0, 217, 227, 3, 18, 9, 0, 218, 227, 3, 8, 4, 0,
		219, 227, 3, 26, 13, 0, 220, 221, 5, 11, 0, 0, 221, 227, 3, 12, 6, 14,
		222, 223, 5, 10, 0, 0, 223, 227, 3, 12, 6, 13, 224, 225, 5, 9, 0, 0, 225,
		227, 3, 12, 6, 12, 226, 205, 1, 0, 0, 0, 226, 210, 1, 0, 0, 0, 226, 211,
		1, 0, 0, 0, 226, 212, 1, 0, 0, 0, 226, 213, 1, 0, 0, 0, 226, 214, 1, 0,
		0, 0, 226, 215, 1, 0, 0, 0, 226, 216, 1, 0, 0, 0, 226, 217, 1, 0, 0, 0,
		226, 218, 1, 0, 0, 0, 226, 219, 1, 0, 0, 0, 226, 220, 1, 0, 0, 0, 226,
		222, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 227, 298, 1, 0, 0, 0, 228, 229,
		10, 11, 0, 0, 229, 230, 7, 0, 0, 0, 230, 297, 3, 12, 6, 12, 231, 232, 10,
		10, 0, 0, 232, 233, 7, 1, 0, 0, 233, 297, 3, 12, 6, 11, 234, 235, 10, 9,
		0, 0, 235, 236, 5, 29, 0, 0, 236, 297, 3, 12, 6, 10, 237, 238, 10, 8, 0,
		0, 238, 239, 7, 2, 0, 0, 239, 297, 3, 12, 6, 9, 240, 241, 10, 7, 0, 0,
		241, 242, 7, 3, 0, 0, 242, 297, 3, 12, 6, 8, 243, 244, 10, 6, 0, 0, 244,
		245, 5, 7, 0, 0, 245, 297, 3, 12, 6, 7, 246, 247, 10, 5, 0, 0, 247, 248,
		5, 8, 0, 0, 248, 297, 3, 12, 6, 6, 249, 250, 10, 4, 0, 0, 250, 251, 5,
		28, 0, 0, 251, 297, 3, 12, 6, 5, 252, 253, 10, 2, 0, 0, 253, 254, 7, 4,
		0, 0, 254, 297, 3, 12, 6, 3, 255, 256, 10, 1, 0, 0, 256, 257, 5, 27, 0,
		0, 257, 297, 3, 12, 6, 2, 258, 260, 10, 19, 0, 0, 259, 261, 5, 19, 0, 0,
		260, 259, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262,
		263, 5, 36, 0, 0, 263, 297, 3, 26, 13, 0, 264, 266, 10, 18, 0, 0, 265,
		267, 5, 19, 0, 0, 266, 265, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 268,
		1, 0, 0, 0, 268, 269, 5, 16, 0, 0, 269, 270, 3, 28, 14, 0, 270, 271, 5,
		17, 0, 0, 271, 297, 1, 0, 0, 0, 272, 273, 10, 17, 0, 0, 273, 282, 5, 14,
		0, 0, 274, 279, 3, 10, 5, 0, 275, 276, 5, 32, 0, 0, 276, 278, 3, 10, 5,
		0, 277, 275, 1, 0, 0, 0, 278, 281, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0, 279,
		280, 1, 0, 0, 0, 280, 283, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 282, 274,
		1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284, 297, 5, 15,
		0, 0, 285, 286, 10, 16, 0, 0, 286, 297, 5, 25, 0, 0, 287, 288, 10, 15,
		0, 0, 288, 297, 5, 26, 0, 0, 289, 290, 10, 3, 0, 0, 290, 291, 5, 19, 0,
		0, 291, 294, 3, 12, 6, 0, 292, 293, 5, 34, 0, 0, 293, 295, 3, 12, 6, 0,
		294, 292, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295, 297, 1, 0, 0, 0, 296,
		228, 1, 0, 0, 0, 296, 231, 1, 0, 0, 0, 296, 234, 1, 0, 0, 0, 296, 237,
		1, 0, 0, 0, 296, 240, 1, 0, 0, 0, 296, 243, 1, 0, 0, 0, 296, 246, 1, 0,
		0, 0, 296, 249, 1, 0, 0, 0, 296, 252, 1, 0, 0, 0, 296, 255, 1, 0, 0, 0,
		296, 258, 1, 0, 0, 0, 296, 264, 1, 0, 0, 0, 296, 272, 1, 0, 0, 0, 296,
		285, 1, 0, 0, 0, 296, 287, 1, 0, 0, 0, 296, 289, 1, 0, 0, 0, 297, 300,
		1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299, 13, 1, 0,
		0, 0, 300, 298, 1, 0, 0, 0, 301, 310, 5, 16, 0, 0, 302, 307, 3, 16, 8,
		0, 303, 304, 5, 32, 0, 0, 304, 306, 3, 16, 8, 0, 305, 303, 1, 0, 0, 0,
		306, 309, 1, 0, 0, 0, 307, 305, 1, 0, 0, 0, 307, 308, 1, 0, 0, 0, 308,
		311, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 310, 302, 1, 0, 0, 0, 310, 311,
		1, 0, 0, 0, 311, 312, 1, 0, 0, 0, 312, 313, 5, 17, 0, 0, 313, 15, 1, 0,
		0, 0, 314, 316, 5, 30, 0, 0, 315, 314, 1, 0, 0, 0, 315, 316, 1, 0, 0, 0,
		316, 317, 1, 0, 0, 0, 317, 318, 3, 12, 6, 0, 318, 17, 1, 0, 0, 0, 319,
		328, 5, 37, 0, 0, 320, 325, 3, 20, 10, 0, 321, 322, 5, 32, 0, 0, 322, 324,
		3, 20, 10, 0, 323, 321, 1, 0, 0, 0, 324, 327, 1, 0, 0, 0, 325, 323, 1,
		0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 329, 1, 0, 0, 0, 327, 325, 1, 0, 0,
		0, 328, 320, 1, 0, 0, 0, 328, 329, 1, 0, 0, 0, 329, 330, 1, 0, 0, 0, 330,
		331, 5, 38, 0, 0, 331, 19, 1, 0, 0, 0, 332, 333, 3, 26, 13, 0, 333, 334,
		5, 34, 0, 0, 334, 335, 3, 12, 6, 0, 335, 350, 1, 0, 0, 0, 336, 337, 5,
		52, 0, 0, 337, 338, 5, 34, 0, 0, 338, 350, 3, 12, 6, 0, 339, 340, 5, 16,
		0, 0, 340, 341, 3, 12, 6, 0, 341, 342, 5, 17, 0, 0, 342, 343, 5, 34, 0,
		0, 343, 344, 3, 12, 6, 0, 344, 350, 1, 0, 0, 0, 345, 347, 5, 30, 0, 0,
		346, 345, 1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 347, 348, 1, 0, 0, 0, 348,
		350, 3, 12, 6, 0, 349, 332, 1, 0, 0, 0, 349, 336, 1, 0, 0, 0, 349, 339,
		1, 0, 0, 0, 349, 346, 1, 0, 0, 0, 350, 21, 1, 0, 0, 0, 351, 355, 5, 51,
		0, 0, 352, 354, 3, 24, 12, 0, 353, 352, 1, 0, 0, 0, 354, 357, 1, 0, 0,
		0, 355, 353, 1, 0, 0, 0, 355, 356, 1, 0, 0, 0, 356, 358, 1, 0, 0, 0, 357,
		355, 1, 0, 0, 0, 358, 359, 5, 60, 0, 0, 359, 23, 1, 0, 0, 0, 360, 366,
		5, 58, 0, 0, 361, 362, 5, 59, 0, 0, 362, 363, 3, 12, 6, 0, 363, 364, 5,
		38, 0, 0, 364, 366, 1, 0, 0, 0, 365, 360, 1, 0, 0, 0, 365, 361, 1, 0, 0,
		0, 366, 25, 1, 0, 0, 0, 367, 368, 5, 53, 0, 0, 368, 27, 1, 0, 0, 0, 369,
		373, 3, 12, 6, 0, 370, 373, 5, 54, 0, 0, 371, 373, 5, 52, 0, 0, 372, 369,
		1, 0, 0, 0, 372, 370, 1, 0, 0, 0, 372, 371, 1, 0, 0, 0, 373, 29, 1, 0,
		0, 0, 43, 33, 42, 53, 69, 74, 86, 92, 99, 107, 109, 113, 119, 125, 131,
		137, 143, 148, 152, 168, 173, 185, 190, 196, 199, 203, 226, 260, 266, 279,
		282, 294, 296, 298, 307, 310, 315, 325, 328, 346, 349, 355, 365, 372,
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
	JsonTemplateSubtractAssign        = 21
	JsonTemplateMultiplyAssign        = 22
	JsonTemplateDivideAssign          = 23
	JsonTemplateModuloAssign          = 24
	JsonTemplateIncrement             = 25
	JsonTemplateDecrement             = 26
	JsonTemplateLiteral               = 27
	JsonTemplateNullCoalescing        = 28
	JsonTemplateRange                 = 29
	JsonTemplateSpread                = 30
	JsonTemplateAs                    = 31
	JsonTemplateComma                 = 32
	JsonTemplateArrow                 = 33
	JsonTemplateColon                 = 34
	JsonTemplateSemicolon             = 35
	JsonTemplateDot                   = 36
	JsonTemplateLeftBrace             = 37
	JsonTemplateRightBrace            = 38
	JsonTemplateNull                  = 39
	JsonTemplateFalse                 = 40
	JsonTemplateTrue                  = 41
	JsonTemplateReturn                = 42
	JsonTemplateFor                   = 43
	JsonTemplateIn                    = 44
	JsonTemplateIf                    = 45
	JsonTemplateElse                  = 46
	JsonTemplateWhile                 = 47
	JsonTemplateDo                    = 48
	JsonTemplateBreak                 = 49
	JsonTemplateContinue              = 50
	JsonTemplateBacktick              = 51
	JsonTemplateESCAPED_STRING        = 52
	JsonTemplateSTRING                = 53
	JsonTemplateNUMBER                = 54
	JsonTemplateLINE_COMMENT          = 55
	JsonTemplateBLOCK_COMMENT         = 56
	JsonTemplateWS                    = 57
	JsonTemplateTEMPLATE_TEXT         = 58
	JsonTemplateTEMPLATE_INTERP_START = 59
	JsonTemplateTEMPLATE_END          = 60
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

func (s *ScriptContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitScript(s)

	default:
		return t.VisitChildren(s)
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

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&35940423771966976) != 0 {
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

func (s *StatementContext) Statements() IStatementsContext {
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

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
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

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33780982935014912) != 0 {
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
			p.Statement()
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
			p.Statement()
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
					p.Statement()
				}

			}
			p.SetState(71)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(72)
				p.Match(JsonTemplateElse)
			}
			{
				p.SetState(73)
				p.Statement()
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
			p.Statement()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(80)
			p.Match(JsonTemplateDo)
		}
		{
			p.SetState(81)
			p.Statement()
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

func (s *StatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitStatements(s)

	default:
		return t.VisitChildren(s)
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

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&35940423771966976) != 0 {
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

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
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

func (s *LambdaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitLambda(s)

	default:
		return t.VisitChildren(s)
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

func (s *Function_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitFunction_param(s)

	default:
		return t.VisitChildren(s)
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

func (s *FieldContext) Add() antlr.TerminalNode {
	return s.GetToken(JsonTemplateAdd, 0)
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

func (s *FieldContext) Range() antlr.TerminalNode {
	return s.GetToken(JsonTemplateRange, 0)
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

func (s *FieldContext) In() antlr.TerminalNode {
	return s.GetToken(JsonTemplateIn, 0)
}

func (s *FieldContext) Equal() antlr.TerminalNode {
	return s.GetToken(JsonTemplateEqual, 0)
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

func (s *FieldContext) NullCoalescing() antlr.TerminalNode {
	return s.GetToken(JsonTemplateNullCoalescing, 0)
}

func (s *FieldContext) AddAssign() antlr.TerminalNode {
	return s.GetToken(JsonTemplateAddAssign, 0)
}

func (s *FieldContext) SubtractAssign() antlr.TerminalNode {
	return s.GetToken(JsonTemplateSubtractAssign, 0)
}

func (s *FieldContext) MultiplyAssign() antlr.TerminalNode {
	return s.GetToken(JsonTemplateMultiplyAssign, 0)
}

func (s *FieldContext) DivideAssign() antlr.TerminalNode {
	return s.GetToken(JsonTemplateDivideAssign, 0)
}

func (s *FieldContext) ModuloAssign() antlr.TerminalNode {
	return s.GetToken(JsonTemplateModuloAssign, 0)
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

func (s *FieldContext) Increment() antlr.TerminalNode {
	return s.GetToken(JsonTemplateIncrement, 0)
}

func (s *FieldContext) Decrement() antlr.TerminalNode {
	return s.GetToken(JsonTemplateDecrement, 0)
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

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
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
	p.SetState(226)
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
			p.field(14)
		}

	case 13:
		{
			p.SetState(222)
			p.Match(JsonTemplateAdd)
		}
		{
			p.SetState(223)
			p.field(13)
		}

	case 14:
		{
			p.SetState(224)
			p.Match(JsonTemplateNot)
		}
		{
			p.SetState(225)
			p.field(12)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(296)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
			case 1:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(228)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(229)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JsonTemplateMultiply || _la == JsonTemplateDivide) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(230)
					p.field(12)
				}

			case 2:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(231)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(232)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JsonTemplateAdd || _la == JsonTemplateSubtract) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(233)
					p.field(11)
				}

			case 3:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(234)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(235)
					p.Match(JsonTemplateRange)
				}
				{
					p.SetState(236)
					p.field(10)
				}

			case 4:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(237)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(238)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17592186044470) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(239)
					p.field(9)
				}

			case 5:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(240)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(241)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JsonTemplateEqual || _la == JsonTemplateNotEqual) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(242)
					p.field(8)
				}

			case 6:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(243)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(244)
					p.Match(JsonTemplateAnd)
				}
				{
					p.SetState(245)
					p.field(7)
				}

			case 7:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(246)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(247)
					p.Match(JsonTemplateOr)
				}
				{
					p.SetState(248)
					p.field(6)
				}

			case 8:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(249)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(250)
					p.Match(JsonTemplateNullCoalescing)
				}
				{
					p.SetState(251)
					p.field(5)
				}

			case 9:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(252)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(253)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32505856) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(254)
					p.field(3)
				}

			case 10:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(255)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(256)
					p.Match(JsonTemplateLiteral)
				}
				{
					p.SetState(257)
					p.field(2)
				}

			case 11:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(258)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}

				p.SetState(260)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonTemplateQuestion {
					{
						p.SetState(259)
						p.Match(JsonTemplateQuestion)
					}

				}
				{
					p.SetState(262)
					p.Match(JsonTemplateDot)
				}
				{
					p.SetState(263)
					p.Name()
				}

			case 12:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(264)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}

				p.SetState(266)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonTemplateQuestion {
					{
						p.SetState(265)
						p.Match(JsonTemplateQuestion)
					}

				}
				{
					p.SetState(268)
					p.Match(JsonTemplateLeftBracket)
				}
				{
					p.SetState(269)
					p.Index()
				}
				{
					p.SetState(270)
					p.Match(JsonTemplateRightBracket)
				}

			case 13:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(272)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(273)
					p.Match(JsonTemplateLeftParen)
				}
				p.SetState(282)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33780984008756736) != 0 {
					{
						p.SetState(274)
						p.Function_param()
					}
					p.SetState(279)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == JsonTemplateComma {
						{
							p.SetState(275)
							p.Match(JsonTemplateComma)
						}
						{
							p.SetState(276)
							p.Function_param()
						}

						p.SetState(281)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(284)
					p.Match(JsonTemplateRightParen)
				}

			case 14:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(285)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(286)
					p.Match(JsonTemplateIncrement)
				}

			case 15:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(287)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(288)
					p.Match(JsonTemplateDecrement)
				}

			case 16:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateRULE_field)
				p.SetState(289)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(290)
					p.Match(JsonTemplateQuestion)
				}
				{
					p.SetState(291)
					p.field(0)
				}
				p.SetState(294)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(292)
						p.Match(JsonTemplateColon)
					}
					{
						p.SetState(293)
						p.field(0)
					}

				}

			}

		}
		p.SetState(300)
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

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(301)
		p.Match(JsonTemplateLeftBracket)
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33780984008756736) != 0 {
		{
			p.SetState(302)
			p.Spread_field()
		}
		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonTemplateComma {
			{
				p.SetState(303)
				p.Match(JsonTemplateComma)
			}
			{
				p.SetState(304)
				p.Spread_field()
			}

			p.SetState(309)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(312)
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

func (s *Spread_fieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitSpread_field(s)

	default:
		return t.VisitChildren(s)
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
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JsonTemplateSpread {
		{
			p.SetState(314)
			p.Match(JsonTemplateSpread)
		}

	}
	{
		p.SetState(317)
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

func (s *ObjectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitObject(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(319)
		p.Match(JsonTemplateLeftBrace)
	}
	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33780984008756736) != 0 {
		{
			p.SetState(320)
			p.Object_field()
		}
		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonTemplateComma {
			{
				p.SetState(321)
				p.Match(JsonTemplateComma)
			}
			{
				p.SetState(322)
				p.Object_field()
			}

			p.SetState(327)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(330)
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

func (s *Object_fieldContext) AllField() []IFieldContext {
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

func (s *Object_fieldContext) Field(i int) IFieldContext {
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

func (s *Object_fieldContext) ESCAPED_STRING() antlr.TerminalNode {
	return s.GetToken(JsonTemplateESCAPED_STRING, 0)
}

func (s *Object_fieldContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(JsonTemplateLeftBracket, 0)
}

func (s *Object_fieldContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(JsonTemplateRightBracket, 0)
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

func (s *Object_fieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitObject_field(s)

	default:
		return t.VisitChildren(s)
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
			p.SetState(332)
			p.Name()
		}
		{
			p.SetState(333)
			p.Match(JsonTemplateColon)
		}
		{
			p.SetState(334)
			p.field(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(336)
			p.Match(JsonTemplateESCAPED_STRING)
		}
		{
			p.SetState(337)
			p.Match(JsonTemplateColon)
		}
		{
			p.SetState(338)
			p.field(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(339)
			p.Match(JsonTemplateLeftBracket)
		}
		{
			p.SetState(340)
			p.field(0)
		}
		{
			p.SetState(341)
			p.Match(JsonTemplateRightBracket)
		}
		{
			p.SetState(342)
			p.Match(JsonTemplateColon)
		}
		{
			p.SetState(343)
			p.field(0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
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

func (s *Template_stringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitTemplate_string(s)

	default:
		return t.VisitChildren(s)
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

func (s *Template_string_partContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitTemplate_string_part(s)

	default:
		return t.VisitChildren(s)
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

func (s *NameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitName(s)

	default:
		return t.VisitChildren(s)
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

func (s *IndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonTemplateVisitor:
		return t.VisitIndex(s)

	default:
		return t.VisitChildren(s)
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
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
