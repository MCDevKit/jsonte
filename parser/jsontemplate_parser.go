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
		"Break", "Continue", "ESCAPED_STRING", "STRING", "NUMBER", "LINE_COMMENT",
		"BLOCK_COMMENT", "WS",
	}
	staticData.ruleNames = []string{
		"script", "statement", "statements", "expression", "lambda", "function_param",
		"field", "array", "spread_field", "object", "object_field", "name",
		"index",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 50, 354, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 5, 0, 28, 8, 0, 10, 0, 12, 0, 31,
		9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 39, 8, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 50, 8, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 64, 8, 1, 10,
		1, 12, 1, 67, 9, 1, 1, 1, 1, 1, 3, 1, 71, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 83, 8, 1, 1, 2, 1, 2, 5, 2, 87,
		8, 2, 10, 2, 12, 2, 90, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 96, 8, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 104, 8, 3, 3, 3, 106, 8, 3, 1, 3,
		1, 3, 3, 3, 110, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 116, 8, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 3, 3, 122, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 128, 8, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 134, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3,
		140, 8, 3, 1, 3, 1, 3, 1, 3, 3, 3, 145, 8, 3, 1, 3, 1, 3, 3, 3, 149, 8,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 5, 4, 163, 8, 4, 10, 4, 12, 4, 166, 9, 4, 5, 4, 168, 8, 4, 10, 4, 12,
		4, 171, 9, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 180, 8, 4,
		10, 4, 12, 4, 183, 9, 4, 5, 4, 185, 8, 4, 10, 4, 12, 4, 188, 9, 4, 1, 4,
		1, 4, 1, 4, 3, 4, 193, 8, 4, 1, 5, 3, 5, 196, 8, 5, 1, 5, 1, 5, 3, 5, 200,
		8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 220, 8, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 266, 8, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 3, 6, 272, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 5, 6, 283, 8, 6, 10, 6, 12, 6, 286, 9, 6, 3, 6, 288, 8, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 296, 8, 6, 5, 6, 298, 8, 6, 10, 6,
		12, 6, 301, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 307, 8, 7, 10, 7, 12, 7,
		310, 9, 7, 3, 7, 312, 8, 7, 1, 7, 1, 7, 1, 8, 3, 8, 317, 8, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 325, 8, 9, 10, 9, 12, 9, 328, 9, 9, 3,
		9, 330, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 3, 10, 342, 8, 10, 1, 10, 3, 10, 345, 8, 10, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 12, 3, 12, 352, 8, 12, 1, 12, 0, 1, 12, 13, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 0, 2, 1, 0, 12, 13, 1, 0, 10, 11, 420,
		0, 29, 1, 0, 0, 0, 2, 82, 1, 0, 0, 0, 4, 84, 1, 0, 0, 0, 6, 148, 1, 0,
		0, 0, 8, 192, 1, 0, 0, 0, 10, 199, 1, 0, 0, 0, 12, 219, 1, 0, 0, 0, 14,
		302, 1, 0, 0, 0, 16, 316, 1, 0, 0, 0, 18, 320, 1, 0, 0, 0, 20, 344, 1,
		0, 0, 0, 22, 346, 1, 0, 0, 0, 24, 351, 1, 0, 0, 0, 26, 28, 3, 2, 1, 0,
		27, 26, 1, 0, 0, 0, 28, 31, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 29, 30, 1,
		0, 0, 0, 30, 1, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 32, 83, 3, 4, 2, 0, 33,
		34, 3, 12, 6, 0, 34, 35, 5, 29, 0, 0, 35, 83, 1, 0, 0, 0, 36, 38, 5, 36,
		0, 0, 37, 39, 3, 12, 6, 0, 38, 37, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39,
		40, 1, 0, 0, 0, 40, 83, 5, 29, 0, 0, 41, 42, 5, 43, 0, 0, 42, 83, 5, 29,
		0, 0, 43, 44, 5, 44, 0, 0, 44, 83, 5, 29, 0, 0, 45, 46, 5, 37, 0, 0, 46,
		49, 3, 22, 11, 0, 47, 48, 5, 26, 0, 0, 48, 50, 3, 22, 11, 0, 49, 47, 1,
		0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 52, 5, 38, 0, 0, 52,
		53, 3, 12, 6, 0, 53, 54, 3, 4, 2, 0, 54, 83, 1, 0, 0, 0, 55, 56, 5, 39,
		0, 0, 56, 57, 3, 12, 6, 0, 57, 65, 3, 4, 2, 0, 58, 59, 5, 40, 0, 0, 59,
		60, 5, 39, 0, 0, 60, 61, 3, 12, 6, 0, 61, 62, 3, 4, 2, 0, 62, 64, 1, 0,
		0, 0, 63, 58, 1, 0, 0, 0, 64, 67, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66,
		1, 0, 0, 0, 66, 70, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 68, 69, 5, 40, 0, 0,
		69, 71, 3, 4, 2, 0, 70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 83, 1,
		0, 0, 0, 72, 73, 5, 41, 0, 0, 73, 74, 3, 12, 6, 0, 74, 75, 3, 4, 2, 0,
		75, 83, 1, 0, 0, 0, 76, 77, 5, 42, 0, 0, 77, 78, 3, 4, 2, 0, 78, 79, 5,
		41, 0, 0, 79, 80, 3, 12, 6, 0, 80, 81, 5, 29, 0, 0, 81, 83, 1, 0, 0, 0,
		82, 32, 1, 0, 0, 0, 82, 33, 1, 0, 0, 0, 82, 36, 1, 0, 0, 0, 82, 41, 1,
		0, 0, 0, 82, 43, 1, 0, 0, 0, 82, 45, 1, 0, 0, 0, 82, 55, 1, 0, 0, 0, 82,
		72, 1, 0, 0, 0, 82, 76, 1, 0, 0, 0, 83, 3, 1, 0, 0, 0, 84, 88, 5, 31, 0,
		0, 85, 87, 3, 2, 1, 0, 86, 85, 1, 0, 0, 0, 87, 90, 1, 0, 0, 0, 88, 86,
		1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 91, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0,
		91, 92, 5, 32, 0, 0, 92, 5, 1, 0, 0, 0, 93, 94, 5, 31, 0, 0, 94, 96, 5,
		31, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97,
		98, 5, 18, 0, 0, 98, 105, 3, 12, 6, 0, 99, 100, 5, 25, 0, 0, 100, 103,
		3, 22, 11, 0, 101, 102, 5, 26, 0, 0, 102, 104, 3, 22, 11, 0, 103, 101,
		1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 106, 1, 0, 0, 0, 105, 99, 1, 0,
		0, 0, 105, 106, 1, 0, 0, 0, 106, 109, 1, 0, 0, 0, 107, 108, 5, 32, 0, 0,
		108, 110, 5, 32, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110,
		111, 1, 0, 0, 0, 111, 112, 5, 0, 0, 1, 112, 149, 1, 0, 0, 0, 113, 114,
		5, 31, 0, 0, 114, 116, 5, 31, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1,
		0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 118, 5, 19, 0, 0, 118, 121, 3, 12,
		6, 0, 119, 120, 5, 32, 0, 0, 120, 122, 5, 32, 0, 0, 121, 119, 1, 0, 0,
		0, 121, 122, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 124, 5, 0, 0, 1, 124,
		149, 1, 0, 0, 0, 125, 126, 5, 31, 0, 0, 126, 128, 5, 31, 0, 0, 127, 125,
		1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 130, 5, 21,
		0, 0, 130, 133, 3, 12, 6, 0, 131, 132, 5, 32, 0, 0, 132, 134, 5, 32, 0,
		0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135,
		136, 5, 0, 0, 1, 136, 149, 1, 0, 0, 0, 137, 138, 5, 31, 0, 0, 138, 140,
		5, 31, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 141, 1, 0,
		0, 0, 141, 144, 3, 12, 6, 0, 142, 143, 5, 32, 0, 0, 143, 145, 5, 32, 0,
		0, 144, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146,
		147, 5, 0, 0, 1, 147, 149, 1, 0, 0, 0, 148, 95, 1, 0, 0, 0, 148, 115, 1,
		0, 0, 0, 148, 127, 1, 0, 0, 0, 148, 139, 1, 0, 0, 0, 149, 7, 1, 0, 0, 0,
		150, 151, 3, 22, 11, 0, 151, 152, 5, 27, 0, 0, 152, 153, 3, 12, 6, 0, 153,
		193, 1, 0, 0, 0, 154, 155, 3, 22, 11, 0, 155, 156, 5, 27, 0, 0, 156, 157,
		3, 4, 2, 0, 157, 193, 1, 0, 0, 0, 158, 169, 5, 14, 0, 0, 159, 164, 3, 22,
		11, 0, 160, 161, 5, 26, 0, 0, 161, 163, 3, 22, 11, 0, 162, 160, 1, 0, 0,
		0, 163, 166, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165,
		168, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 167, 159, 1, 0, 0, 0, 168, 171,
		1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 172, 1, 0,
		0, 0, 171, 169, 1, 0, 0, 0, 172, 173, 5, 15, 0, 0, 173, 174, 5, 27, 0,
		0, 174, 193, 3, 12, 6, 0, 175, 186, 5, 14, 0, 0, 176, 181, 3, 22, 11, 0,
		177, 178, 5, 26, 0, 0, 178, 180, 3, 22, 11, 0, 179, 177, 1, 0, 0, 0, 180,
		183, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 185,
		1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 184, 176, 1, 0, 0, 0, 185, 188, 1, 0,
		0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 189, 1, 0, 0, 0,
		188, 186, 1, 0, 0, 0, 189, 190, 5, 15, 0, 0, 190, 191, 5, 27, 0, 0, 191,
		193, 3, 4, 2, 0, 192, 150, 1, 0, 0, 0, 192, 154, 1, 0, 0, 0, 192, 158,
		1, 0, 0, 0, 192, 175, 1, 0, 0, 0, 193, 9, 1, 0, 0, 0, 194, 196, 5, 24,
		0, 0, 195, 194, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0,
		197, 200, 3, 12, 6, 0, 198, 200, 3, 8, 4, 0, 199, 195, 1, 0, 0, 0, 199,
		198, 1, 0, 0, 0, 200, 11, 1, 0, 0, 0, 201, 202, 6, 6, -1, 0, 202, 203,
		5, 14, 0, 0, 203, 204, 3, 12, 6, 0, 204, 205, 5, 15, 0, 0, 205, 220, 1,
		0, 0, 0, 206, 220, 5, 35, 0, 0, 207, 220, 5, 34, 0, 0, 208, 220, 5, 33,
		0, 0, 209, 220, 5, 47, 0, 0, 210, 220, 5, 45, 0, 0, 211, 220, 3, 14, 7,
		0, 212, 220, 3, 18, 9, 0, 213, 220, 3, 8, 4, 0, 214, 220, 3, 22, 11, 0,
		215, 216, 5, 11, 0, 0, 216, 220, 3, 12, 6, 17, 217, 218, 5, 9, 0, 0, 218,
		220, 3, 12, 6, 16, 219, 201, 1, 0, 0, 0, 219, 206, 1, 0, 0, 0, 219, 207,
		1, 0, 0, 0, 219, 208, 1, 0, 0, 0, 219, 209, 1, 0, 0, 0, 219, 210, 1, 0,
		0, 0, 219, 211, 1, 0, 0, 0, 219, 212, 1, 0, 0, 0, 219, 213, 1, 0, 0, 0,
		219, 214, 1, 0, 0, 0, 219, 215, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 220,
		299, 1, 0, 0, 0, 221, 222, 10, 15, 0, 0, 222, 223, 7, 0, 0, 0, 223, 298,
		3, 12, 6, 16, 224, 225, 10, 14, 0, 0, 225, 226, 7, 1, 0, 0, 226, 298, 3,
		12, 6, 15, 227, 228, 10, 13, 0, 0, 228, 229, 5, 23, 0, 0, 229, 298, 3,
		12, 6, 14, 230, 231, 10, 12, 0, 0, 231, 232, 5, 22, 0, 0, 232, 298, 3,
		12, 6, 13, 233, 234, 10, 11, 0, 0, 234, 235, 5, 3, 0, 0, 235, 298, 3, 12,
		6, 12, 236, 237, 10, 10, 0, 0, 237, 238, 5, 1, 0, 0, 238, 298, 3, 12, 6,
		11, 239, 240, 10, 9, 0, 0, 240, 241, 5, 2, 0, 0, 241, 298, 3, 12, 6, 10,
		242, 243, 10, 8, 0, 0, 243, 244, 5, 4, 0, 0, 244, 298, 3, 12, 6, 9, 245,
		246, 10, 7, 0, 0, 246, 247, 5, 5, 0, 0, 247, 298, 3, 12, 6, 8, 248, 249,
		10, 6, 0, 0, 249, 250, 5, 6, 0, 0, 250, 298, 3, 12, 6, 7, 251, 252, 10,
		5, 0, 0, 252, 253, 5, 7, 0, 0, 253, 298, 3, 12, 6, 6, 254, 255, 10, 4,
		0, 0, 255, 256, 5, 8, 0, 0, 256, 298, 3, 12, 6, 5, 257, 258, 10, 2, 0,
		0, 258, 259, 5, 20, 0, 0, 259, 298, 3, 12, 6, 3, 260, 261, 10, 1, 0, 0,
		261, 262, 5, 21, 0, 0, 262, 298, 3, 12, 6, 2, 263, 265, 10, 20, 0, 0, 264,
		266, 5, 19, 0, 0, 265, 264, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 267,
		1, 0, 0, 0, 267, 268, 5, 30, 0, 0, 268, 298, 3, 22, 11, 0, 269, 271, 10,
		19, 0, 0, 270, 272, 5, 19, 0, 0, 271, 270, 1, 0, 0, 0, 271, 272, 1, 0,
		0, 0, 272, 273, 1, 0, 0, 0, 273, 274, 5, 16, 0, 0, 274, 275, 3, 24, 12,
		0, 275, 276, 5, 17, 0, 0, 276, 298, 1, 0, 0, 0, 277, 278, 10, 18, 0, 0,
		278, 287, 5, 14, 0, 0, 279, 284, 3, 10, 5, 0, 280, 281, 5, 26, 0, 0, 281,
		283, 3, 10, 5, 0, 282, 280, 1, 0, 0, 0, 283, 286, 1, 0, 0, 0, 284, 282,
		1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 288, 1, 0, 0, 0, 286, 284, 1, 0,
		0, 0, 287, 279, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0,
		289, 298, 5, 15, 0, 0, 290, 291, 10, 3, 0, 0, 291, 292, 5, 19, 0, 0, 292,
		295, 3, 12, 6, 0, 293, 294, 5, 28, 0, 0, 294, 296, 3, 12, 6, 0, 295, 293,
		1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 298, 1, 0, 0, 0, 297, 221, 1, 0,
		0, 0, 297, 224, 1, 0, 0, 0, 297, 227, 1, 0, 0, 0, 297, 230, 1, 0, 0, 0,
		297, 233, 1, 0, 0, 0, 297, 236, 1, 0, 0, 0, 297, 239, 1, 0, 0, 0, 297,
		242, 1, 0, 0, 0, 297, 245, 1, 0, 0, 0, 297, 248, 1, 0, 0, 0, 297, 251,
		1, 0, 0, 0, 297, 254, 1, 0, 0, 0, 297, 257, 1, 0, 0, 0, 297, 260, 1, 0,
		0, 0, 297, 263, 1, 0, 0, 0, 297, 269, 1, 0, 0, 0, 297, 277, 1, 0, 0, 0,
		297, 290, 1, 0, 0, 0, 298, 301, 1, 0, 0, 0, 299, 297, 1, 0, 0, 0, 299,
		300, 1, 0, 0, 0, 300, 13, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 302, 311, 5,
		16, 0, 0, 303, 308, 3, 16, 8, 0, 304, 305, 5, 26, 0, 0, 305, 307, 3, 16,
		8, 0, 306, 304, 1, 0, 0, 0, 307, 310, 1, 0, 0, 0, 308, 306, 1, 0, 0, 0,
		308, 309, 1, 0, 0, 0, 309, 312, 1, 0, 0, 0, 310, 308, 1, 0, 0, 0, 311,
		303, 1, 0, 0, 0, 311, 312, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 314,
		5, 17, 0, 0, 314, 15, 1, 0, 0, 0, 315, 317, 5, 24, 0, 0, 316, 315, 1, 0,
		0, 0, 316, 317, 1, 0, 0, 0, 317, 318, 1, 0, 0, 0, 318, 319, 3, 12, 6, 0,
		319, 17, 1, 0, 0, 0, 320, 329, 5, 31, 0, 0, 321, 326, 3, 20, 10, 0, 322,
		323, 5, 26, 0, 0, 323, 325, 3, 20, 10, 0, 324, 322, 1, 0, 0, 0, 325, 328,
		1, 0, 0, 0, 326, 324, 1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327, 330, 1, 0,
		0, 0, 328, 326, 1, 0, 0, 0, 329, 321, 1, 0, 0, 0, 329, 330, 1, 0, 0, 0,
		330, 331, 1, 0, 0, 0, 331, 332, 5, 32, 0, 0, 332, 19, 1, 0, 0, 0, 333,
		334, 3, 22, 11, 0, 334, 335, 5, 28, 0, 0, 335, 336, 3, 12, 6, 0, 336, 345,
		1, 0, 0, 0, 337, 338, 5, 45, 0, 0, 338, 339, 5, 28, 0, 0, 339, 345, 3,
		12, 6, 0, 340, 342, 5, 24, 0, 0, 341, 340, 1, 0, 0, 0, 341, 342, 1, 0,
		0, 0, 342, 343, 1, 0, 0, 0, 343, 345, 3, 12, 6, 0, 344, 333, 1, 0, 0, 0,
		344, 337, 1, 0, 0, 0, 344, 341, 1, 0, 0, 0, 345, 21, 1, 0, 0, 0, 346, 347,
		5, 46, 0, 0, 347, 23, 1, 0, 0, 0, 348, 352, 3, 12, 6, 0, 349, 352, 5, 47,
		0, 0, 350, 352, 5, 45, 0, 0, 351, 348, 1, 0, 0, 0, 351, 349, 1, 0, 0, 0,
		351, 350, 1, 0, 0, 0, 352, 25, 1, 0, 0, 0, 41, 29, 38, 49, 65, 70, 82,
		88, 95, 103, 105, 109, 115, 121, 127, 133, 139, 144, 148, 164, 169, 181,
		186, 192, 195, 199, 219, 265, 271, 284, 287, 295, 297, 299, 308, 311, 316,
		326, 329, 341, 344, 351,
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
	JsonTemplateParserLess           = 1
	JsonTemplateParserLessOrEqual    = 2
	JsonTemplateParserEqual          = 3
	JsonTemplateParserGreater        = 4
	JsonTemplateParserGreaterOrEqual = 5
	JsonTemplateParserNotEqual       = 6
	JsonTemplateParserAnd            = 7
	JsonTemplateParserOr             = 8
	JsonTemplateParserNot            = 9
	JsonTemplateParserAdd            = 10
	JsonTemplateParserSubtract       = 11
	JsonTemplateParserMultiply       = 12
	JsonTemplateParserDivide         = 13
	JsonTemplateParserLeftParen      = 14
	JsonTemplateParserRightParen     = 15
	JsonTemplateParserLeftBracket    = 16
	JsonTemplateParserRightBracket   = 17
	JsonTemplateParserIteration      = 18
	JsonTemplateParserQuestion       = 19
	JsonTemplateParserAddAssign      = 20
	JsonTemplateParserLiteral        = 21
	JsonTemplateParserNullCoalescing = 22
	JsonTemplateParserRange          = 23
	JsonTemplateParserSpread         = 24
	JsonTemplateParserAs             = 25
	JsonTemplateParserComma          = 26
	JsonTemplateParserArrow          = 27
	JsonTemplateParserColon          = 28
	JsonTemplateParserSemicolon      = 29
	JsonTemplateParserDot            = 30
	JsonTemplateParserLeftBrace      = 31
	JsonTemplateParserRightBrace     = 32
	JsonTemplateParserNull           = 33
	JsonTemplateParserFalse          = 34
	JsonTemplateParserTrue           = 35
	JsonTemplateParserReturn         = 36
	JsonTemplateParserFor            = 37
	JsonTemplateParserIn             = 38
	JsonTemplateParserIf             = 39
	JsonTemplateParserElse           = 40
	JsonTemplateParserWhile          = 41
	JsonTemplateParserDo             = 42
	JsonTemplateParserBreak          = 43
	JsonTemplateParserContinue       = 44
	JsonTemplateParserESCAPED_STRING = 45
	JsonTemplateParserSTRING         = 46
	JsonTemplateParserNUMBER         = 47
	JsonTemplateParserLINE_COMMENT   = 48
	JsonTemplateParserBLOCK_COMMENT  = 49
	JsonTemplateParserWS             = 50
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

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&280094144809472) != 0 {
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
	return s.GetToken(JsonTemplateParserSemicolon, 0)
}

func (s *StatementContext) Return() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserReturn, 0)
}

func (s *StatementContext) Break() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserBreak, 0)
}

func (s *StatementContext) Continue() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserContinue, 0)
}

func (s *StatementContext) For() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserFor, 0)
}

func (s *StatementContext) In() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserIn, 0)
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
	return s.GetToken(JsonTemplateParserComma, 0)
}

func (s *StatementContext) AllIf() []antlr.TerminalNode {
	return s.GetTokens(JsonTemplateParserIf)
}

func (s *StatementContext) If(i int) antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserIf, i)
}

func (s *StatementContext) AllElse() []antlr.TerminalNode {
	return s.GetTokens(JsonTemplateParserElse)
}

func (s *StatementContext) Else(i int) antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserElse, i)
}

func (s *StatementContext) While() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserWhile, 0)
}

func (s *StatementContext) Do() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserDo, 0)
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

	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(32)
			p.Statements()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(33)
			p.field(0)
		}
		{
			p.SetState(34)
			p.Match(JsonTemplateParserSemicolon)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(36)
			p.Match(JsonTemplateParserReturn)
		}
		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&246352881732096) != 0 {
			{
				p.SetState(37)
				p.field(0)
			}

		}
		{
			p.SetState(40)
			p.Match(JsonTemplateParserSemicolon)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(41)
			p.Match(JsonTemplateParserBreak)
		}
		{
			p.SetState(42)
			p.Match(JsonTemplateParserSemicolon)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(43)
			p.Match(JsonTemplateParserContinue)
		}
		{
			p.SetState(44)
			p.Match(JsonTemplateParserSemicolon)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(45)
			p.Match(JsonTemplateParserFor)
		}

		{
			p.SetState(46)
			p.Name()
		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserComma {
			{
				p.SetState(47)
				p.Match(JsonTemplateParserComma)
			}
			{
				p.SetState(48)
				p.Name()
			}

		}

		{
			p.SetState(51)
			p.Match(JsonTemplateParserIn)
		}
		{
			p.SetState(52)
			p.field(0)
		}
		{
			p.SetState(53)
			p.Statements()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(55)
			p.Match(JsonTemplateParserIf)
		}
		{
			p.SetState(56)
			p.field(0)
		}
		{
			p.SetState(57)
			p.Statements()
		}
		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(58)
					p.Match(JsonTemplateParserElse)
				}
				{
					p.SetState(59)
					p.Match(JsonTemplateParserIf)
				}
				{
					p.SetState(60)
					p.field(0)
				}
				{
					p.SetState(61)
					p.Statements()
				}

			}
			p.SetState(67)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserElse {
			{
				p.SetState(68)
				p.Match(JsonTemplateParserElse)
			}
			{
				p.SetState(69)
				p.Statements()
			}

		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(72)
			p.Match(JsonTemplateParserWhile)
		}
		{
			p.SetState(73)
			p.field(0)
		}
		{
			p.SetState(74)
			p.Statements()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(76)
			p.Match(JsonTemplateParserDo)
		}
		{
			p.SetState(77)
			p.Statements()
		}
		{
			p.SetState(78)
			p.Match(JsonTemplateParserWhile)
		}
		{
			p.SetState(79)
			p.field(0)
		}
		{
			p.SetState(80)
			p.Match(JsonTemplateParserSemicolon)
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
		p.SetState(84)
		p.Match(JsonTemplateParserLeftBrace)
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&280094144809472) != 0 {
		{
			p.SetState(85)
			p.Statement()
		}

		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(91)
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

	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserLeftBrace {
			{
				p.SetState(93)
				p.Match(JsonTemplateParserLeftBrace)
			}
			{
				p.SetState(94)
				p.Match(JsonTemplateParserLeftBrace)
			}

		}
		{
			p.SetState(97)
			p.Match(JsonTemplateParserIteration)
		}
		{
			p.SetState(98)
			p.field(0)
		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserAs {
			{
				p.SetState(99)
				p.Match(JsonTemplateParserAs)
			}
			{
				p.SetState(100)
				p.Name()
			}
			p.SetState(103)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == JsonTemplateParserComma {
				{
					p.SetState(101)
					p.Match(JsonTemplateParserComma)
				}
				{
					p.SetState(102)
					p.Name()
				}

			}

		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserRightBrace {
			{
				p.SetState(107)
				p.Match(JsonTemplateParserRightBrace)
			}
			{
				p.SetState(108)
				p.Match(JsonTemplateParserRightBrace)
			}

		}
		{
			p.SetState(111)
			p.Match(JsonTemplateParserEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserLeftBrace {
			{
				p.SetState(113)
				p.Match(JsonTemplateParserLeftBrace)
			}
			{
				p.SetState(114)
				p.Match(JsonTemplateParserLeftBrace)
			}

		}
		{
			p.SetState(117)
			p.Match(JsonTemplateParserQuestion)
		}
		{
			p.SetState(118)
			p.field(0)
		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserRightBrace {
			{
				p.SetState(119)
				p.Match(JsonTemplateParserRightBrace)
			}
			{
				p.SetState(120)
				p.Match(JsonTemplateParserRightBrace)
			}

		}
		{
			p.SetState(123)
			p.Match(JsonTemplateParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserLeftBrace {
			{
				p.SetState(125)
				p.Match(JsonTemplateParserLeftBrace)
			}
			{
				p.SetState(126)
				p.Match(JsonTemplateParserLeftBrace)
			}

		}
		{
			p.SetState(129)
			p.Match(JsonTemplateParserLiteral)
		}
		{
			p.SetState(130)
			p.field(0)
		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserRightBrace {
			{
				p.SetState(131)
				p.Match(JsonTemplateParserRightBrace)
			}
			{
				p.SetState(132)
				p.Match(JsonTemplateParserRightBrace)
			}

		}
		{
			p.SetState(135)
			p.Match(JsonTemplateParserEOF)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(139)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(137)
				p.Match(JsonTemplateParserLeftBrace)
			}
			{
				p.SetState(138)
				p.Match(JsonTemplateParserLeftBrace)
			}

		}
		{
			p.SetState(141)
			p.field(0)
		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserRightBrace {
			{
				p.SetState(142)
				p.Match(JsonTemplateParserRightBrace)
			}
			{
				p.SetState(143)
				p.Match(JsonTemplateParserRightBrace)
			}

		}
		{
			p.SetState(146)
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

	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(150)
			p.Name()
		}
		{
			p.SetState(151)
			p.Match(JsonTemplateParserArrow)
		}
		{
			p.SetState(152)
			p.field(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.Name()
		}
		{
			p.SetState(155)
			p.Match(JsonTemplateParserArrow)
		}
		{
			p.SetState(156)
			p.Statements()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(158)
			p.Match(JsonTemplateParserLeftParen)
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonTemplateParserSTRING {
			{
				p.SetState(159)
				p.Name()
			}
			p.SetState(164)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == JsonTemplateParserComma {
				{
					p.SetState(160)
					p.Match(JsonTemplateParserComma)
				}
				{
					p.SetState(161)
					p.Name()
				}

				p.SetState(166)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

			p.SetState(171)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(172)
			p.Match(JsonTemplateParserRightParen)
		}
		{
			p.SetState(173)
			p.Match(JsonTemplateParserArrow)
		}
		{
			p.SetState(174)
			p.field(0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(175)
			p.Match(JsonTemplateParserLeftParen)
		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonTemplateParserSTRING {
			{
				p.SetState(176)
				p.Name()
			}
			p.SetState(181)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == JsonTemplateParserComma {
				{
					p.SetState(177)
					p.Match(JsonTemplateParserComma)
				}
				{
					p.SetState(178)
					p.Name()
				}

				p.SetState(183)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

			p.SetState(188)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(189)
			p.Match(JsonTemplateParserRightParen)
		}
		{
			p.SetState(190)
			p.Match(JsonTemplateParserArrow)
		}
		{
			p.SetState(191)
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

	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserSpread {
			{
				p.SetState(194)
				p.Match(JsonTemplateParserSpread)
			}

		}
		{
			p.SetState(197)
			p.field(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(198)
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

func (s *FieldContext) AddAssign() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserAddAssign, 0)
}

func (s *FieldContext) Literal() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserLiteral, 0)
}

func (s *FieldContext) Dot() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserDot, 0)
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

func (s *FieldContext) Colon() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserColon, 0)
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
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(202)
			p.Match(JsonTemplateParserLeftParen)
		}
		{
			p.SetState(203)
			p.field(0)
		}
		{
			p.SetState(204)
			p.Match(JsonTemplateParserRightParen)
		}

	case 2:
		{
			p.SetState(206)
			p.Match(JsonTemplateParserTrue)
		}

	case 3:
		{
			p.SetState(207)
			p.Match(JsonTemplateParserFalse)
		}

	case 4:
		{
			p.SetState(208)
			p.Match(JsonTemplateParserNull)
		}

	case 5:
		{
			p.SetState(209)
			p.Match(JsonTemplateParserNUMBER)
		}

	case 6:
		{
			p.SetState(210)
			p.Match(JsonTemplateParserESCAPED_STRING)
		}

	case 7:
		{
			p.SetState(211)
			p.Array()
		}

	case 8:
		{
			p.SetState(212)
			p.Object()
		}

	case 9:
		{
			p.SetState(213)
			p.Lambda()
		}

	case 10:
		{
			p.SetState(214)
			p.Name()
		}

	case 11:
		{
			p.SetState(215)
			p.Match(JsonTemplateParserSubtract)
		}
		{
			p.SetState(216)
			p.field(17)
		}

	case 12:
		{
			p.SetState(217)
			p.Match(JsonTemplateParserNot)
		}
		{
			p.SetState(218)
			p.field(16)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(297)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
			case 1:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(221)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(222)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JsonTemplateParserMultiply || _la == JsonTemplateParserDivide) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(223)
					p.field(16)
				}

			case 2:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(224)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(225)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JsonTemplateParserAdd || _la == JsonTemplateParserSubtract) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(226)
					p.field(15)
				}

			case 3:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(227)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(228)
					p.Match(JsonTemplateParserRange)
				}
				{
					p.SetState(229)
					p.field(14)
				}

			case 4:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(230)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(231)
					p.Match(JsonTemplateParserNullCoalescing)
				}
				{
					p.SetState(232)
					p.field(13)
				}

			case 5:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(233)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(234)
					p.Match(JsonTemplateParserEqual)
				}
				{
					p.SetState(235)
					p.field(12)
				}

			case 6:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(236)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(237)
					p.Match(JsonTemplateParserLess)
				}
				{
					p.SetState(238)
					p.field(11)
				}

			case 7:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(239)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(240)
					p.Match(JsonTemplateParserLessOrEqual)
				}
				{
					p.SetState(241)
					p.field(10)
				}

			case 8:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(242)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(243)
					p.Match(JsonTemplateParserGreater)
				}
				{
					p.SetState(244)
					p.field(9)
				}

			case 9:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(245)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(246)
					p.Match(JsonTemplateParserGreaterOrEqual)
				}
				{
					p.SetState(247)
					p.field(8)
				}

			case 10:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(248)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(249)
					p.Match(JsonTemplateParserNotEqual)
				}
				{
					p.SetState(250)
					p.field(7)
				}

			case 11:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(251)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(252)
					p.Match(JsonTemplateParserAnd)
				}
				{
					p.SetState(253)
					p.field(6)
				}

			case 12:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(254)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(255)
					p.Match(JsonTemplateParserOr)
				}
				{
					p.SetState(256)
					p.field(5)
				}

			case 13:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(257)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(258)
					p.Match(JsonTemplateParserAddAssign)
				}
				{
					p.SetState(259)
					p.field(3)
				}

			case 14:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(260)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(261)
					p.Match(JsonTemplateParserLiteral)
				}
				{
					p.SetState(262)
					p.field(2)
				}

			case 15:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(263)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
				}

				p.SetState(265)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonTemplateParserQuestion {
					{
						p.SetState(264)
						p.Match(JsonTemplateParserQuestion)
					}

				}
				{
					p.SetState(267)
					p.Match(JsonTemplateParserDot)
				}
				{
					p.SetState(268)
					p.Name()
				}

			case 16:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(269)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}

				p.SetState(271)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonTemplateParserQuestion {
					{
						p.SetState(270)
						p.Match(JsonTemplateParserQuestion)
					}

				}
				{
					p.SetState(273)
					p.Match(JsonTemplateParserLeftBracket)
				}
				{
					p.SetState(274)
					p.Index()
				}
				{
					p.SetState(275)
					p.Match(JsonTemplateParserRightBracket)
				}

			case 17:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(277)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(278)
					p.Match(JsonTemplateParserLeftParen)
				}
				p.SetState(287)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&246352898509312) != 0 {
					{
						p.SetState(279)
						p.Function_param()
					}
					p.SetState(284)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == JsonTemplateParserComma {
						{
							p.SetState(280)
							p.Match(JsonTemplateParserComma)
						}
						{
							p.SetState(281)
							p.Function_param()
						}

						p.SetState(286)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(289)
					p.Match(JsonTemplateParserRightParen)
				}

			case 18:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(290)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(291)
					p.Match(JsonTemplateParserQuestion)
				}
				{
					p.SetState(292)
					p.field(0)
				}
				p.SetState(295)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(293)
						p.Match(JsonTemplateParserColon)
					}
					{
						p.SetState(294)
						p.field(0)
					}

				}

			}

		}
		p.SetState(301)
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
		p.SetState(302)
		p.Match(JsonTemplateParserLeftBracket)
	}
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&246352898509312) != 0 {
		{
			p.SetState(303)
			p.Spread_field()
		}
		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonTemplateParserComma {
			{
				p.SetState(304)
				p.Match(JsonTemplateParserComma)
			}
			{
				p.SetState(305)
				p.Spread_field()
			}

			p.SetState(310)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(313)
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
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JsonTemplateParserSpread {
		{
			p.SetState(315)
			p.Match(JsonTemplateParserSpread)
		}

	}
	{
		p.SetState(318)
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
		p.SetState(320)
		p.Match(JsonTemplateParserLeftBrace)
	}
	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&246352898509312) != 0 {
		{
			p.SetState(321)
			p.Object_field()
		}
		p.SetState(326)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonTemplateParserComma {
			{
				p.SetState(322)
				p.Match(JsonTemplateParserComma)
			}
			{
				p.SetState(323)
				p.Object_field()
			}

			p.SetState(328)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(331)
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

func (s *Object_fieldContext) Colon() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserColon, 0)
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

	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(333)
			p.Name()
		}
		{
			p.SetState(334)
			p.Match(JsonTemplateParserColon)
		}
		{
			p.SetState(335)
			p.field(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(337)
			p.Match(JsonTemplateParserESCAPED_STRING)
		}
		{
			p.SetState(338)
			p.Match(JsonTemplateParserColon)
		}
		{
			p.SetState(339)
			p.field(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(341)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserSpread {
			{
				p.SetState(340)
				p.Match(JsonTemplateParserSpread)
			}

		}
		{
			p.SetState(343)
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
		p.SetState(346)
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

	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(348)
			p.field(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(349)
			p.Match(JsonTemplateParserNUMBER)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(350)
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
