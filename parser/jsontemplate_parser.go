// Code generated from ../grammar/JsonTemplate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // JsonTemplate

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
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
		"expression", "lambda", "function_param", "field", "array", "object",
		"object_field", "name", "index",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 36, 227, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 3, 0, 21,
		8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 27, 8, 0, 1, 0, 1, 0, 3, 0, 31, 8,
		0, 1, 0, 1, 0, 3, 0, 35, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 41, 8, 0,
		1, 0, 1, 0, 3, 0, 45, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 51, 8, 0, 1,
		0, 1, 0, 3, 0, 55, 8, 0, 1, 0, 1, 0, 1, 0, 3, 0, 60, 8, 0, 3, 0, 62, 8,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 72, 8, 1, 10,
		1, 12, 1, 75, 9, 1, 5, 1, 77, 8, 1, 10, 1, 12, 1, 80, 9, 1, 1, 1, 1, 1,
		1, 1, 3, 1, 85, 8, 1, 1, 2, 1, 2, 3, 2, 89, 8, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 3, 3, 108, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 148, 8, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 3, 3, 154, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 5, 3, 165, 8, 3, 10, 3, 12, 3, 168, 9, 3, 3, 3, 170, 8, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 178, 8, 3, 5, 3, 180, 8, 3, 10,
		3, 12, 3, 183, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 189, 8, 4, 10, 4, 12,
		4, 192, 9, 4, 3, 4, 194, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5,
		202, 8, 5, 10, 5, 12, 5, 205, 9, 5, 3, 5, 207, 8, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 218, 8, 6, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 3, 8, 225, 8, 8, 1, 8, 0, 1, 6, 9, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 0, 2, 1, 0, 14, 15, 1, 0, 12, 13, 271, 0, 61, 1, 0, 0, 0, 2, 84, 1,
		0, 0, 0, 4, 88, 1, 0, 0, 0, 6, 107, 1, 0, 0, 0, 8, 184, 1, 0, 0, 0, 10,
		197, 1, 0, 0, 0, 12, 217, 1, 0, 0, 0, 14, 219, 1, 0, 0, 0, 16, 224, 1,
		0, 0, 0, 18, 19, 5, 28, 0, 0, 19, 21, 5, 28, 0, 0, 20, 18, 1, 0, 0, 0,
		20, 21, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 23, 5, 20, 0, 0, 23, 26, 3,
		6, 3, 0, 24, 25, 5, 25, 0, 0, 25, 27, 3, 14, 7, 0, 26, 24, 1, 0, 0, 0,
		26, 27, 1, 0, 0, 0, 27, 30, 1, 0, 0, 0, 28, 29, 5, 29, 0, 0, 29, 31, 5,
		29, 0, 0, 30, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 62, 1, 0, 0, 0, 32,
		33, 5, 28, 0, 0, 33, 35, 5, 28, 0, 0, 34, 32, 1, 0, 0, 0, 34, 35, 1, 0,
		0, 0, 35, 36, 1, 0, 0, 0, 36, 37, 5, 21, 0, 0, 37, 40, 3, 6, 3, 0, 38,
		39, 5, 29, 0, 0, 39, 41, 5, 29, 0, 0, 40, 38, 1, 0, 0, 0, 40, 41, 1, 0,
		0, 0, 41, 62, 1, 0, 0, 0, 42, 43, 5, 28, 0, 0, 43, 45, 5, 28, 0, 0, 44,
		42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 47, 5, 22,
		0, 0, 47, 50, 3, 6, 3, 0, 48, 49, 5, 29, 0, 0, 49, 51, 5, 29, 0, 0, 50,
		48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 62, 1, 0, 0, 0, 52, 53, 5, 28,
		0, 0, 53, 55, 5, 28, 0, 0, 54, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55,
		56, 1, 0, 0, 0, 56, 59, 3, 6, 3, 0, 57, 58, 5, 29, 0, 0, 58, 60, 5, 29,
		0, 0, 59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 62, 1, 0, 0, 0, 61, 20,
		1, 0, 0, 0, 61, 34, 1, 0, 0, 0, 61, 44, 1, 0, 0, 0, 61, 54, 1, 0, 0, 0,
		62, 1, 1, 0, 0, 0, 63, 64, 3, 14, 7, 0, 64, 65, 5, 27, 0, 0, 65, 66, 3,
		6, 3, 0, 66, 85, 1, 0, 0, 0, 67, 78, 5, 16, 0, 0, 68, 73, 3, 14, 7, 0,
		69, 70, 5, 26, 0, 0, 70, 72, 3, 14, 7, 0, 71, 69, 1, 0, 0, 0, 72, 75, 1,
		0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 77, 1, 0, 0, 0, 75,
		73, 1, 0, 0, 0, 76, 68, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0, 78, 76, 1, 0, 0,
		0, 78, 79, 1, 0, 0, 0, 79, 81, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 81, 82,
		5, 17, 0, 0, 82, 83, 5, 27, 0, 0, 83, 85, 3, 6, 3, 0, 84, 63, 1, 0, 0,
		0, 84, 67, 1, 0, 0, 0, 85, 3, 1, 0, 0, 0, 86, 89, 3, 6, 3, 0, 87, 89, 3,
		2, 1, 0, 88, 86, 1, 0, 0, 0, 88, 87, 1, 0, 0, 0, 89, 5, 1, 0, 0, 0, 90,
		91, 6, 3, -1, 0, 91, 92, 5, 16, 0, 0, 92, 93, 3, 6, 3, 0, 93, 94, 5, 17,
		0, 0, 94, 108, 1, 0, 0, 0, 95, 108, 5, 32, 0, 0, 96, 108, 5, 31, 0, 0,
		97, 108, 5, 30, 0, 0, 98, 108, 5, 35, 0, 0, 99, 108, 5, 33, 0, 0, 100,
		108, 3, 8, 4, 0, 101, 108, 3, 10, 5, 0, 102, 108, 3, 14, 7, 0, 103, 104,
		5, 13, 0, 0, 104, 108, 3, 6, 3, 15, 105, 106, 5, 11, 0, 0, 106, 108, 3,
		6, 3, 2, 107, 90, 1, 0, 0, 0, 107, 95, 1, 0, 0, 0, 107, 96, 1, 0, 0, 0,
		107, 97, 1, 0, 0, 0, 107, 98, 1, 0, 0, 0, 107, 99, 1, 0, 0, 0, 107, 100,
		1, 0, 0, 0, 107, 101, 1, 0, 0, 0, 107, 102, 1, 0, 0, 0, 107, 103, 1, 0,
		0, 0, 107, 105, 1, 0, 0, 0, 108, 181, 1, 0, 0, 0, 109, 110, 10, 14, 0,
		0, 110, 111, 7, 0, 0, 0, 111, 180, 3, 6, 3, 15, 112, 113, 10, 13, 0, 0,
		113, 114, 7, 1, 0, 0, 114, 180, 3, 6, 3, 14, 115, 116, 10, 12, 0, 0, 116,
		117, 5, 24, 0, 0, 117, 180, 3, 6, 3, 13, 118, 119, 10, 11, 0, 0, 119, 120,
		5, 23, 0, 0, 120, 180, 3, 6, 3, 12, 121, 122, 10, 10, 0, 0, 122, 123, 5,
		5, 0, 0, 123, 180, 3, 6, 3, 11, 124, 125, 10, 9, 0, 0, 125, 126, 5, 3,
		0, 0, 126, 180, 3, 6, 3, 10, 127, 128, 10, 8, 0, 0, 128, 129, 5, 4, 0,
		0, 129, 180, 3, 6, 3, 9, 130, 131, 10, 7, 0, 0, 131, 132, 5, 6, 0, 0, 132,
		180, 3, 6, 3, 8, 133, 134, 10, 6, 0, 0, 134, 135, 5, 7, 0, 0, 135, 180,
		3, 6, 3, 7, 136, 137, 10, 5, 0, 0, 137, 138, 5, 8, 0, 0, 138, 180, 3, 6,
		3, 6, 139, 140, 10, 4, 0, 0, 140, 141, 5, 9, 0, 0, 141, 180, 3, 6, 3, 5,
		142, 143, 10, 3, 0, 0, 143, 144, 5, 10, 0, 0, 144, 180, 3, 6, 3, 4, 145,
		147, 10, 18, 0, 0, 146, 148, 5, 21, 0, 0, 147, 146, 1, 0, 0, 0, 147, 148,
		1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 150, 5, 1, 0, 0, 150, 180, 3, 14,
		7, 0, 151, 153, 10, 17, 0, 0, 152, 154, 5, 21, 0, 0, 153, 152, 1, 0, 0,
		0, 153, 154, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 156, 5, 18, 0, 0, 156,
		157, 3, 16, 8, 0, 157, 158, 5, 19, 0, 0, 158, 180, 1, 0, 0, 0, 159, 160,
		10, 16, 0, 0, 160, 169, 5, 16, 0, 0, 161, 166, 3, 4, 2, 0, 162, 163, 5,
		26, 0, 0, 163, 165, 3, 4, 2, 0, 164, 162, 1, 0, 0, 0, 165, 168, 1, 0, 0,
		0, 166, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 170, 1, 0, 0, 0, 168,
		166, 1, 0, 0, 0, 169, 161, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 171,
		1, 0, 0, 0, 171, 180, 5, 17, 0, 0, 172, 173, 10, 1, 0, 0, 173, 174, 5,
		21, 0, 0, 174, 177, 3, 6, 3, 0, 175, 176, 5, 2, 0, 0, 176, 178, 3, 6, 3,
		0, 177, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 180, 1, 0, 0, 0, 179,
		109, 1, 0, 0, 0, 179, 112, 1, 0, 0, 0, 179, 115, 1, 0, 0, 0, 179, 118,
		1, 0, 0, 0, 179, 121, 1, 0, 0, 0, 179, 124, 1, 0, 0, 0, 179, 127, 1, 0,
		0, 0, 179, 130, 1, 0, 0, 0, 179, 133, 1, 0, 0, 0, 179, 136, 1, 0, 0, 0,
		179, 139, 1, 0, 0, 0, 179, 142, 1, 0, 0, 0, 179, 145, 1, 0, 0, 0, 179,
		151, 1, 0, 0, 0, 179, 159, 1, 0, 0, 0, 179, 172, 1, 0, 0, 0, 180, 183,
		1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 7, 1, 0, 0,
		0, 183, 181, 1, 0, 0, 0, 184, 193, 5, 18, 0, 0, 185, 190, 3, 6, 3, 0, 186,
		187, 5, 26, 0, 0, 187, 189, 3, 6, 3, 0, 188, 186, 1, 0, 0, 0, 189, 192,
		1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 194, 1, 0,
		0, 0, 192, 190, 1, 0, 0, 0, 193, 185, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0,
		194, 195, 1, 0, 0, 0, 195, 196, 5, 19, 0, 0, 196, 9, 1, 0, 0, 0, 197, 206,
		5, 28, 0, 0, 198, 203, 3, 12, 6, 0, 199, 200, 5, 26, 0, 0, 200, 202, 3,
		12, 6, 0, 201, 199, 1, 0, 0, 0, 202, 205, 1, 0, 0, 0, 203, 201, 1, 0, 0,
		0, 203, 204, 1, 0, 0, 0, 204, 207, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 206,
		198, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 209,
		5, 29, 0, 0, 209, 11, 1, 0, 0, 0, 210, 211, 3, 14, 7, 0, 211, 212, 5, 2,
		0, 0, 212, 213, 3, 6, 3, 0, 213, 218, 1, 0, 0, 0, 214, 215, 5, 33, 0, 0,
		215, 216, 5, 2, 0, 0, 216, 218, 3, 6, 3, 0, 217, 210, 1, 0, 0, 0, 217,
		214, 1, 0, 0, 0, 218, 13, 1, 0, 0, 0, 219, 220, 5, 34, 0, 0, 220, 15, 1,
		0, 0, 0, 221, 225, 3, 6, 3, 0, 222, 225, 5, 35, 0, 0, 223, 225, 5, 33,
		0, 0, 224, 221, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224, 223, 1, 0, 0, 0,
		225, 17, 1, 0, 0, 0, 28, 20, 26, 30, 34, 40, 44, 50, 54, 59, 61, 73, 78,
		84, 88, 107, 147, 153, 166, 169, 177, 179, 181, 190, 193, 203, 206, 217,
		224,
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
	this.GrammarFileName = "JsonTemplate.g4"

	return this
}

// JsonTemplateParser tokens.
const (
	JsonTemplateParserEOF            = antlr.TokenEOF
	JsonTemplateParserT__0           = 1
	JsonTemplateParserT__1           = 2
	JsonTemplateParserLess           = 3
	JsonTemplateParserLessOrEqual    = 4
	JsonTemplateParserEqual          = 5
	JsonTemplateParserGreater        = 6
	JsonTemplateParserGreaterOrEqual = 7
	JsonTemplateParserNotEqual       = 8
	JsonTemplateParserAnd            = 9
	JsonTemplateParserOr             = 10
	JsonTemplateParserNot            = 11
	JsonTemplateParserAdd            = 12
	JsonTemplateParserSubtract       = 13
	JsonTemplateParserMultiply       = 14
	JsonTemplateParserDivide         = 15
	JsonTemplateParserLeftParen      = 16
	JsonTemplateParserRightParen     = 17
	JsonTemplateParserLeftBracket    = 18
	JsonTemplateParserRightBracket   = 19
	JsonTemplateParserIteration      = 20
	JsonTemplateParserQuestion       = 21
	JsonTemplateParserLiteral        = 22
	JsonTemplateParserNullCoalescing = 23
	JsonTemplateParserRange          = 24
	JsonTemplateParserAs             = 25
	JsonTemplateParserComma          = 26
	JsonTemplateParserArrow          = 27
	JsonTemplateParserLeftBrace      = 28
	JsonTemplateParserRightBrace     = 29
	JsonTemplateParserNull           = 30
	JsonTemplateParserFalse          = 31
	JsonTemplateParserTrue           = 32
	JsonTemplateParserESCAPED_STRING = 33
	JsonTemplateParserSTRING         = 34
	JsonTemplateParserNUMBER         = 35
	JsonTemplateParserWS             = 36
)

// JsonTemplateParser rules.
const (
	JsonTemplateParserRULE_expression     = 0
	JsonTemplateParserRULE_lambda         = 1
	JsonTemplateParserRULE_function_param = 2
	JsonTemplateParserRULE_field          = 3
	JsonTemplateParserRULE_array          = 4
	JsonTemplateParserRULE_object         = 5
	JsonTemplateParserRULE_object_field   = 6
	JsonTemplateParserRULE_name           = 7
	JsonTemplateParserRULE_index          = 8
)

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

func (s *ExpressionContext) AllLeftBrace() []antlr.TerminalNode {
	return s.GetTokens(JsonTemplateParserLeftBrace)
}

func (s *ExpressionContext) LeftBrace(i int) antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserLeftBrace, i)
}

func (s *ExpressionContext) As() antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserAs, 0)
}

func (s *ExpressionContext) Name() INameContext {
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

func (s *ExpressionContext) AllRightBrace() []antlr.TerminalNode {
	return s.GetTokens(JsonTemplateParserRightBrace)
}

func (s *ExpressionContext) RightBrace(i int) antlr.TerminalNode {
	return s.GetToken(JsonTemplateParserRightBrace, i)
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
	p.EnterRule(localctx, 0, JsonTemplateParserRULE_expression)
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

	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserLeftBrace {
			{
				p.SetState(18)
				p.Match(JsonTemplateParserLeftBrace)
			}
			{
				p.SetState(19)
				p.Match(JsonTemplateParserLeftBrace)
			}

		}
		{
			p.SetState(22)
			p.Match(JsonTemplateParserIteration)
		}
		{
			p.SetState(23)
			p.field(0)
		}
		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserAs {
			{
				p.SetState(24)
				p.Match(JsonTemplateParserAs)
			}
			{
				p.SetState(25)
				p.Name()
			}

		}
		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserRightBrace {
			{
				p.SetState(28)
				p.Match(JsonTemplateParserRightBrace)
			}
			{
				p.SetState(29)
				p.Match(JsonTemplateParserRightBrace)
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserLeftBrace {
			{
				p.SetState(32)
				p.Match(JsonTemplateParserLeftBrace)
			}
			{
				p.SetState(33)
				p.Match(JsonTemplateParserLeftBrace)
			}

		}
		{
			p.SetState(36)
			p.Match(JsonTemplateParserQuestion)
		}
		{
			p.SetState(37)
			p.field(0)
		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserRightBrace {
			{
				p.SetState(38)
				p.Match(JsonTemplateParserRightBrace)
			}
			{
				p.SetState(39)
				p.Match(JsonTemplateParserRightBrace)
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserLeftBrace {
			{
				p.SetState(42)
				p.Match(JsonTemplateParserLeftBrace)
			}
			{
				p.SetState(43)
				p.Match(JsonTemplateParserLeftBrace)
			}

		}
		{
			p.SetState(46)
			p.Match(JsonTemplateParserLiteral)
		}
		{
			p.SetState(47)
			p.field(0)
		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserRightBrace {
			{
				p.SetState(48)
				p.Match(JsonTemplateParserRightBrace)
			}
			{
				p.SetState(49)
				p.Match(JsonTemplateParserRightBrace)
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(54)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(52)
				p.Match(JsonTemplateParserLeftBrace)
			}
			{
				p.SetState(53)
				p.Match(JsonTemplateParserLeftBrace)
			}

		}
		{
			p.SetState(56)
			p.field(0)
		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserRightBrace {
			{
				p.SetState(57)
				p.Match(JsonTemplateParserRightBrace)
			}
			{
				p.SetState(58)
				p.Match(JsonTemplateParserRightBrace)
			}

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
	p.EnterRule(localctx, 2, JsonTemplateParserRULE_lambda)
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

	p.SetState(84)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonTemplateParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.Name()
		}
		{
			p.SetState(64)
			p.Match(JsonTemplateParserArrow)
		}
		{
			p.SetState(65)
			p.field(0)
		}

	case JsonTemplateParserLeftParen:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.Match(JsonTemplateParserLeftParen)
		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonTemplateParserSTRING {
			{
				p.SetState(68)
				p.Name()
			}
			p.SetState(73)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == JsonTemplateParserComma {
				{
					p.SetState(69)
					p.Match(JsonTemplateParserComma)
				}
				{
					p.SetState(70)
					p.Name()
				}

				p.SetState(75)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

			p.SetState(80)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(81)
			p.Match(JsonTemplateParserRightParen)
		}
		{
			p.SetState(82)
			p.Match(JsonTemplateParserArrow)
		}
		{
			p.SetState(83)
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
	p.EnterRule(localctx, 4, JsonTemplateParserRULE_function_param)

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

	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.field(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(87)
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
	_startState := 6
	p.EnterRecursionRule(localctx, 6, JsonTemplateParserRULE_field, _p)
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
	p.SetState(107)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonTemplateParserLeftParen:
		{
			p.SetState(91)
			p.Match(JsonTemplateParserLeftParen)
		}
		{
			p.SetState(92)
			p.field(0)
		}
		{
			p.SetState(93)
			p.Match(JsonTemplateParserRightParen)
		}

	case JsonTemplateParserTrue:
		{
			p.SetState(95)
			p.Match(JsonTemplateParserTrue)
		}

	case JsonTemplateParserFalse:
		{
			p.SetState(96)
			p.Match(JsonTemplateParserFalse)
		}

	case JsonTemplateParserNull:
		{
			p.SetState(97)
			p.Match(JsonTemplateParserNull)
		}

	case JsonTemplateParserNUMBER:
		{
			p.SetState(98)
			p.Match(JsonTemplateParserNUMBER)
		}

	case JsonTemplateParserESCAPED_STRING:
		{
			p.SetState(99)
			p.Match(JsonTemplateParserESCAPED_STRING)
		}

	case JsonTemplateParserLeftBracket:
		{
			p.SetState(100)
			p.Array()
		}

	case JsonTemplateParserLeftBrace:
		{
			p.SetState(101)
			p.Object()
		}

	case JsonTemplateParserSTRING:
		{
			p.SetState(102)
			p.Name()
		}

	case JsonTemplateParserSubtract:
		{
			p.SetState(103)
			p.Match(JsonTemplateParserSubtract)
		}
		{
			p.SetState(104)
			p.field(15)
		}

	case JsonTemplateParserNot:
		{
			p.SetState(105)
			p.Match(JsonTemplateParserNot)
		}
		{
			p.SetState(106)
			p.field(2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(179)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
			case 1:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(109)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(110)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JsonTemplateParserMultiply || _la == JsonTemplateParserDivide) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(111)
					p.field(15)
				}

			case 2:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(112)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(113)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JsonTemplateParserAdd || _la == JsonTemplateParserSubtract) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(114)
					p.field(14)
				}

			case 3:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(115)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(116)
					p.Match(JsonTemplateParserRange)
				}
				{
					p.SetState(117)
					p.field(13)
				}

			case 4:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(118)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(119)
					p.Match(JsonTemplateParserNullCoalescing)
				}
				{
					p.SetState(120)
					p.field(12)
				}

			case 5:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(121)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(122)
					p.Match(JsonTemplateParserEqual)
				}
				{
					p.SetState(123)
					p.field(11)
				}

			case 6:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(124)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(125)
					p.Match(JsonTemplateParserLess)
				}
				{
					p.SetState(126)
					p.field(10)
				}

			case 7:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(127)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(128)
					p.Match(JsonTemplateParserLessOrEqual)
				}
				{
					p.SetState(129)
					p.field(9)
				}

			case 8:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(130)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(131)
					p.Match(JsonTemplateParserGreater)
				}
				{
					p.SetState(132)
					p.field(8)
				}

			case 9:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(133)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(134)
					p.Match(JsonTemplateParserGreaterOrEqual)
				}
				{
					p.SetState(135)
					p.field(7)
				}

			case 10:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(136)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(137)
					p.Match(JsonTemplateParserNotEqual)
				}
				{
					p.SetState(138)
					p.field(6)
				}

			case 11:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(139)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(140)
					p.Match(JsonTemplateParserAnd)
				}
				{
					p.SetState(141)
					p.field(5)
				}

			case 12:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(142)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(143)
					p.Match(JsonTemplateParserOr)
				}
				{
					p.SetState(144)
					p.field(4)
				}

			case 13:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(145)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}

				p.SetState(147)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonTemplateParserQuestion {
					{
						p.SetState(146)
						p.Match(JsonTemplateParserQuestion)
					}

				}
				{
					p.SetState(149)
					p.Match(JsonTemplateParserT__0)
				}
				{
					p.SetState(150)
					p.Name()
				}

			case 14:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(151)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}

				p.SetState(153)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonTemplateParserQuestion {
					{
						p.SetState(152)
						p.Match(JsonTemplateParserQuestion)
					}

				}
				{
					p.SetState(155)
					p.Match(JsonTemplateParserLeftBracket)
				}
				{
					p.SetState(156)
					p.Index()
				}
				{
					p.SetState(157)
					p.Match(JsonTemplateParserRightBracket)
				}

			case 15:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(159)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(160)
					p.Match(JsonTemplateParserLeftParen)
				}
				p.SetState(169)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if ((_la-11)&-(0x1f+1)) == 0 && ((1<<uint((_la-11)))&((1<<(JsonTemplateParserNot-11))|(1<<(JsonTemplateParserSubtract-11))|(1<<(JsonTemplateParserLeftParen-11))|(1<<(JsonTemplateParserLeftBracket-11))|(1<<(JsonTemplateParserLeftBrace-11))|(1<<(JsonTemplateParserNull-11))|(1<<(JsonTemplateParserFalse-11))|(1<<(JsonTemplateParserTrue-11))|(1<<(JsonTemplateParserESCAPED_STRING-11))|(1<<(JsonTemplateParserSTRING-11))|(1<<(JsonTemplateParserNUMBER-11)))) != 0 {
					{
						p.SetState(161)
						p.Function_param()
					}
					p.SetState(166)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == JsonTemplateParserComma {
						{
							p.SetState(162)
							p.Match(JsonTemplateParserComma)
						}
						{
							p.SetState(163)
							p.Function_param()
						}

						p.SetState(168)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(171)
					p.Match(JsonTemplateParserRightParen)
				}

			case 16:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(172)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(173)
					p.Match(JsonTemplateParserQuestion)
				}
				{
					p.SetState(174)
					p.field(0)
				}
				p.SetState(177)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(175)
						p.Match(JsonTemplateParserT__1)
					}
					{
						p.SetState(176)
						p.field(0)
					}

				}

			}

		}
		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 8, JsonTemplateParserRULE_array)
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
		p.SetState(184)
		p.Match(JsonTemplateParserLeftBracket)
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-11)&-(0x1f+1)) == 0 && ((1<<uint((_la-11)))&((1<<(JsonTemplateParserNot-11))|(1<<(JsonTemplateParserSubtract-11))|(1<<(JsonTemplateParserLeftParen-11))|(1<<(JsonTemplateParserLeftBracket-11))|(1<<(JsonTemplateParserLeftBrace-11))|(1<<(JsonTemplateParserNull-11))|(1<<(JsonTemplateParserFalse-11))|(1<<(JsonTemplateParserTrue-11))|(1<<(JsonTemplateParserESCAPED_STRING-11))|(1<<(JsonTemplateParserSTRING-11))|(1<<(JsonTemplateParserNUMBER-11)))) != 0 {
		{
			p.SetState(185)
			p.field(0)
		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonTemplateParserComma {
			{
				p.SetState(186)
				p.Match(JsonTemplateParserComma)
			}
			{
				p.SetState(187)
				p.field(0)
			}

			p.SetState(192)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(195)
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
	p.EnterRule(localctx, 10, JsonTemplateParserRULE_object)
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
		p.SetState(197)
		p.Match(JsonTemplateParserLeftBrace)
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JsonTemplateParserESCAPED_STRING || _la == JsonTemplateParserSTRING {
		{
			p.SetState(198)
			p.Object_field()
		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonTemplateParserComma {
			{
				p.SetState(199)
				p.Match(JsonTemplateParserComma)
			}
			{
				p.SetState(200)
				p.Object_field()
			}

			p.SetState(205)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(208)
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
	p.EnterRule(localctx, 12, JsonTemplateParserRULE_object_field)

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

	p.SetState(217)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonTemplateParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)
			p.Name()
		}
		{
			p.SetState(211)
			p.Match(JsonTemplateParserT__1)
		}
		{
			p.SetState(212)
			p.field(0)
		}

	case JsonTemplateParserESCAPED_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(214)
			p.Match(JsonTemplateParserESCAPED_STRING)
		}
		{
			p.SetState(215)
			p.Match(JsonTemplateParserT__1)
		}
		{
			p.SetState(216)
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
	p.EnterRule(localctx, 14, JsonTemplateParserRULE_name)

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
		p.SetState(219)
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
	p.EnterRule(localctx, 16, JsonTemplateParserRULE_index)

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

	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(221)
			p.field(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(222)
			p.Match(JsonTemplateParserNUMBER)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(223)
			p.Match(JsonTemplateParserESCAPED_STRING)
		}

	}

	return localctx
}

func (p *JsonTemplateParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
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
