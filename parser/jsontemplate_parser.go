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
		4, 1, 36, 231, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 3, 0, 21,
		8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 29, 8, 0, 3, 0, 31, 8,
		0, 1, 0, 1, 0, 3, 0, 35, 8, 0, 1, 0, 1, 0, 3, 0, 39, 8, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 3, 0, 45, 8, 0, 1, 0, 1, 0, 3, 0, 49, 8, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 3, 0, 55, 8, 0, 1, 0, 1, 0, 3, 0, 59, 8, 0, 1, 0, 1, 0, 1, 0,
		3, 0, 64, 8, 0, 3, 0, 66, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 5, 1, 76, 8, 1, 10, 1, 12, 1, 79, 9, 1, 5, 1, 81, 8, 1, 10, 1,
		12, 1, 84, 9, 1, 1, 1, 1, 1, 1, 1, 3, 1, 89, 8, 1, 1, 2, 1, 2, 3, 2, 93,
		8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 112, 8, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3,
		152, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 158, 8, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 169, 8, 3, 10, 3, 12, 3, 172, 9,
		3, 3, 3, 174, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 182, 8, 3,
		5, 3, 184, 8, 3, 10, 3, 12, 3, 187, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4,
		193, 8, 4, 10, 4, 12, 4, 196, 9, 4, 3, 4, 198, 8, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 5, 5, 206, 8, 5, 10, 5, 12, 5, 209, 9, 5, 3, 5, 211,
		8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 222,
		8, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 3, 8, 229, 8, 8, 1, 8, 0, 1, 6, 9,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 2, 1, 0, 14, 15, 1, 0, 12, 13, 276, 0,
		65, 1, 0, 0, 0, 2, 88, 1, 0, 0, 0, 4, 92, 1, 0, 0, 0, 6, 111, 1, 0, 0,
		0, 8, 188, 1, 0, 0, 0, 10, 201, 1, 0, 0, 0, 12, 221, 1, 0, 0, 0, 14, 223,
		1, 0, 0, 0, 16, 228, 1, 0, 0, 0, 18, 19, 5, 28, 0, 0, 19, 21, 5, 28, 0,
		0, 20, 18, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 23,
		5, 20, 0, 0, 23, 30, 3, 6, 3, 0, 24, 25, 5, 25, 0, 0, 25, 28, 3, 14, 7,
		0, 26, 27, 5, 26, 0, 0, 27, 29, 3, 14, 7, 0, 28, 26, 1, 0, 0, 0, 28, 29,
		1, 0, 0, 0, 29, 31, 1, 0, 0, 0, 30, 24, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0,
		31, 34, 1, 0, 0, 0, 32, 33, 5, 29, 0, 0, 33, 35, 5, 29, 0, 0, 34, 32, 1,
		0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 66, 1, 0, 0, 0, 36, 37, 5, 28, 0, 0, 37,
		39, 5, 28, 0, 0, 38, 36, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 40, 1, 0,
		0, 0, 40, 41, 5, 21, 0, 0, 41, 44, 3, 6, 3, 0, 42, 43, 5, 29, 0, 0, 43,
		45, 5, 29, 0, 0, 44, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 66, 1, 0,
		0, 0, 46, 47, 5, 28, 0, 0, 47, 49, 5, 28, 0, 0, 48, 46, 1, 0, 0, 0, 48,
		49, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51, 5, 22, 0, 0, 51, 54, 3, 6,
		3, 0, 52, 53, 5, 29, 0, 0, 53, 55, 5, 29, 0, 0, 54, 52, 1, 0, 0, 0, 54,
		55, 1, 0, 0, 0, 55, 66, 1, 0, 0, 0, 56, 57, 5, 28, 0, 0, 57, 59, 5, 28,
		0, 0, 58, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 63,
		3, 6, 3, 0, 61, 62, 5, 29, 0, 0, 62, 64, 5, 29, 0, 0, 63, 61, 1, 0, 0,
		0, 63, 64, 1, 0, 0, 0, 64, 66, 1, 0, 0, 0, 65, 20, 1, 0, 0, 0, 65, 38,
		1, 0, 0, 0, 65, 48, 1, 0, 0, 0, 65, 58, 1, 0, 0, 0, 66, 1, 1, 0, 0, 0,
		67, 68, 3, 14, 7, 0, 68, 69, 5, 27, 0, 0, 69, 70, 3, 6, 3, 0, 70, 89, 1,
		0, 0, 0, 71, 82, 5, 16, 0, 0, 72, 77, 3, 14, 7, 0, 73, 74, 5, 26, 0, 0,
		74, 76, 3, 14, 7, 0, 75, 73, 1, 0, 0, 0, 76, 79, 1, 0, 0, 0, 77, 75, 1,
		0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80,
		72, 1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 83, 1, 0, 0,
		0, 83, 85, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85, 86, 5, 17, 0, 0, 86, 87,
		5, 27, 0, 0, 87, 89, 3, 6, 3, 0, 88, 67, 1, 0, 0, 0, 88, 71, 1, 0, 0, 0,
		89, 3, 1, 0, 0, 0, 90, 93, 3, 6, 3, 0, 91, 93, 3, 2, 1, 0, 92, 90, 1, 0,
		0, 0, 92, 91, 1, 0, 0, 0, 93, 5, 1, 0, 0, 0, 94, 95, 6, 3, -1, 0, 95, 96,
		5, 16, 0, 0, 96, 97, 3, 6, 3, 0, 97, 98, 5, 17, 0, 0, 98, 112, 1, 0, 0,
		0, 99, 112, 5, 32, 0, 0, 100, 112, 5, 31, 0, 0, 101, 112, 5, 30, 0, 0,
		102, 112, 5, 35, 0, 0, 103, 112, 5, 33, 0, 0, 104, 112, 3, 8, 4, 0, 105,
		112, 3, 10, 5, 0, 106, 112, 3, 14, 7, 0, 107, 108, 5, 13, 0, 0, 108, 112,
		3, 6, 3, 15, 109, 110, 5, 11, 0, 0, 110, 112, 3, 6, 3, 2, 111, 94, 1, 0,
		0, 0, 111, 99, 1, 0, 0, 0, 111, 100, 1, 0, 0, 0, 111, 101, 1, 0, 0, 0,
		111, 102, 1, 0, 0, 0, 111, 103, 1, 0, 0, 0, 111, 104, 1, 0, 0, 0, 111,
		105, 1, 0, 0, 0, 111, 106, 1, 0, 0, 0, 111, 107, 1, 0, 0, 0, 111, 109,
		1, 0, 0, 0, 112, 185, 1, 0, 0, 0, 113, 114, 10, 14, 0, 0, 114, 115, 7,
		0, 0, 0, 115, 184, 3, 6, 3, 15, 116, 117, 10, 13, 0, 0, 117, 118, 7, 1,
		0, 0, 118, 184, 3, 6, 3, 14, 119, 120, 10, 12, 0, 0, 120, 121, 5, 24, 0,
		0, 121, 184, 3, 6, 3, 13, 122, 123, 10, 11, 0, 0, 123, 124, 5, 23, 0, 0,
		124, 184, 3, 6, 3, 12, 125, 126, 10, 10, 0, 0, 126, 127, 5, 5, 0, 0, 127,
		184, 3, 6, 3, 11, 128, 129, 10, 9, 0, 0, 129, 130, 5, 3, 0, 0, 130, 184,
		3, 6, 3, 10, 131, 132, 10, 8, 0, 0, 132, 133, 5, 4, 0, 0, 133, 184, 3,
		6, 3, 9, 134, 135, 10, 7, 0, 0, 135, 136, 5, 6, 0, 0, 136, 184, 3, 6, 3,
		8, 137, 138, 10, 6, 0, 0, 138, 139, 5, 7, 0, 0, 139, 184, 3, 6, 3, 7, 140,
		141, 10, 5, 0, 0, 141, 142, 5, 8, 0, 0, 142, 184, 3, 6, 3, 6, 143, 144,
		10, 4, 0, 0, 144, 145, 5, 9, 0, 0, 145, 184, 3, 6, 3, 5, 146, 147, 10,
		3, 0, 0, 147, 148, 5, 10, 0, 0, 148, 184, 3, 6, 3, 4, 149, 151, 10, 18,
		0, 0, 150, 152, 5, 21, 0, 0, 151, 150, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0,
		152, 153, 1, 0, 0, 0, 153, 154, 5, 1, 0, 0, 154, 184, 3, 14, 7, 0, 155,
		157, 10, 17, 0, 0, 156, 158, 5, 21, 0, 0, 157, 156, 1, 0, 0, 0, 157, 158,
		1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 160, 5, 18, 0, 0, 160, 161, 3, 16,
		8, 0, 161, 162, 5, 19, 0, 0, 162, 184, 1, 0, 0, 0, 163, 164, 10, 16, 0,
		0, 164, 173, 5, 16, 0, 0, 165, 170, 3, 4, 2, 0, 166, 167, 5, 26, 0, 0,
		167, 169, 3, 4, 2, 0, 168, 166, 1, 0, 0, 0, 169, 172, 1, 0, 0, 0, 170,
		168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 174, 1, 0, 0, 0, 172, 170,
		1, 0, 0, 0, 173, 165, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 175, 1, 0,
		0, 0, 175, 184, 5, 17, 0, 0, 176, 177, 10, 1, 0, 0, 177, 178, 5, 21, 0,
		0, 178, 181, 3, 6, 3, 0, 179, 180, 5, 2, 0, 0, 180, 182, 3, 6, 3, 0, 181,
		179, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 184, 1, 0, 0, 0, 183, 113,
		1, 0, 0, 0, 183, 116, 1, 0, 0, 0, 183, 119, 1, 0, 0, 0, 183, 122, 1, 0,
		0, 0, 183, 125, 1, 0, 0, 0, 183, 128, 1, 0, 0, 0, 183, 131, 1, 0, 0, 0,
		183, 134, 1, 0, 0, 0, 183, 137, 1, 0, 0, 0, 183, 140, 1, 0, 0, 0, 183,
		143, 1, 0, 0, 0, 183, 146, 1, 0, 0, 0, 183, 149, 1, 0, 0, 0, 183, 155,
		1, 0, 0, 0, 183, 163, 1, 0, 0, 0, 183, 176, 1, 0, 0, 0, 184, 187, 1, 0,
		0, 0, 185, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 7, 1, 0, 0, 0, 187,
		185, 1, 0, 0, 0, 188, 197, 5, 18, 0, 0, 189, 194, 3, 6, 3, 0, 190, 191,
		5, 26, 0, 0, 191, 193, 3, 6, 3, 0, 192, 190, 1, 0, 0, 0, 193, 196, 1, 0,
		0, 0, 194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 198, 1, 0, 0, 0,
		196, 194, 1, 0, 0, 0, 197, 189, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198,
		199, 1, 0, 0, 0, 199, 200, 5, 19, 0, 0, 200, 9, 1, 0, 0, 0, 201, 210, 5,
		28, 0, 0, 202, 207, 3, 12, 6, 0, 203, 204, 5, 26, 0, 0, 204, 206, 3, 12,
		6, 0, 205, 203, 1, 0, 0, 0, 206, 209, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0,
		207, 208, 1, 0, 0, 0, 208, 211, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 210,
		202, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 213,
		5, 29, 0, 0, 213, 11, 1, 0, 0, 0, 214, 215, 3, 14, 7, 0, 215, 216, 5, 2,
		0, 0, 216, 217, 3, 6, 3, 0, 217, 222, 1, 0, 0, 0, 218, 219, 5, 33, 0, 0,
		219, 220, 5, 2, 0, 0, 220, 222, 3, 6, 3, 0, 221, 214, 1, 0, 0, 0, 221,
		218, 1, 0, 0, 0, 222, 13, 1, 0, 0, 0, 223, 224, 5, 34, 0, 0, 224, 15, 1,
		0, 0, 0, 225, 229, 3, 6, 3, 0, 226, 229, 5, 35, 0, 0, 227, 229, 5, 33,
		0, 0, 228, 225, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228, 227, 1, 0, 0, 0,
		229, 17, 1, 0, 0, 0, 29, 20, 28, 30, 34, 38, 44, 48, 54, 58, 63, 65, 77,
		82, 88, 92, 111, 151, 157, 170, 173, 181, 183, 185, 194, 197, 207, 210,
		221, 228,
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

	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
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
		p.SetState(30)
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
			p.SetState(28)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == JsonTemplateParserComma {
				{
					p.SetState(26)
					p.Match(JsonTemplateParserComma)
				}
				{
					p.SetState(27)
					p.Name()
				}

			}

		}
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserRightBrace {
			{
				p.SetState(32)
				p.Match(JsonTemplateParserRightBrace)
			}
			{
				p.SetState(33)
				p.Match(JsonTemplateParserRightBrace)
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserLeftBrace {
			{
				p.SetState(36)
				p.Match(JsonTemplateParserLeftBrace)
			}
			{
				p.SetState(37)
				p.Match(JsonTemplateParserLeftBrace)
			}

		}
		{
			p.SetState(40)
			p.Match(JsonTemplateParserQuestion)
		}
		{
			p.SetState(41)
			p.field(0)
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserRightBrace {
			{
				p.SetState(42)
				p.Match(JsonTemplateParserRightBrace)
			}
			{
				p.SetState(43)
				p.Match(JsonTemplateParserRightBrace)
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserLeftBrace {
			{
				p.SetState(46)
				p.Match(JsonTemplateParserLeftBrace)
			}
			{
				p.SetState(47)
				p.Match(JsonTemplateParserLeftBrace)
			}

		}
		{
			p.SetState(50)
			p.Match(JsonTemplateParserLiteral)
		}
		{
			p.SetState(51)
			p.field(0)
		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserRightBrace {
			{
				p.SetState(52)
				p.Match(JsonTemplateParserRightBrace)
			}
			{
				p.SetState(53)
				p.Match(JsonTemplateParserRightBrace)
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(58)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(56)
				p.Match(JsonTemplateParserLeftBrace)
			}
			{
				p.SetState(57)
				p.Match(JsonTemplateParserLeftBrace)
			}

		}
		{
			p.SetState(60)
			p.field(0)
		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonTemplateParserRightBrace {
			{
				p.SetState(61)
				p.Match(JsonTemplateParserRightBrace)
			}
			{
				p.SetState(62)
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

	p.SetState(88)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonTemplateParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)
			p.Name()
		}
		{
			p.SetState(68)
			p.Match(JsonTemplateParserArrow)
		}
		{
			p.SetState(69)
			p.field(0)
		}

	case JsonTemplateParserLeftParen:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Match(JsonTemplateParserLeftParen)
		}
		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonTemplateParserSTRING {
			{
				p.SetState(72)
				p.Name()
			}
			p.SetState(77)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == JsonTemplateParserComma {
				{
					p.SetState(73)
					p.Match(JsonTemplateParserComma)
				}
				{
					p.SetState(74)
					p.Name()
				}

				p.SetState(79)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

			p.SetState(84)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(85)
			p.Match(JsonTemplateParserRightParen)
		}
		{
			p.SetState(86)
			p.Match(JsonTemplateParserArrow)
		}
		{
			p.SetState(87)
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

	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)
			p.field(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
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
	p.SetState(111)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonTemplateParserLeftParen:
		{
			p.SetState(95)
			p.Match(JsonTemplateParserLeftParen)
		}
		{
			p.SetState(96)
			p.field(0)
		}
		{
			p.SetState(97)
			p.Match(JsonTemplateParserRightParen)
		}

	case JsonTemplateParserTrue:
		{
			p.SetState(99)
			p.Match(JsonTemplateParserTrue)
		}

	case JsonTemplateParserFalse:
		{
			p.SetState(100)
			p.Match(JsonTemplateParserFalse)
		}

	case JsonTemplateParserNull:
		{
			p.SetState(101)
			p.Match(JsonTemplateParserNull)
		}

	case JsonTemplateParserNUMBER:
		{
			p.SetState(102)
			p.Match(JsonTemplateParserNUMBER)
		}

	case JsonTemplateParserESCAPED_STRING:
		{
			p.SetState(103)
			p.Match(JsonTemplateParserESCAPED_STRING)
		}

	case JsonTemplateParserLeftBracket:
		{
			p.SetState(104)
			p.Array()
		}

	case JsonTemplateParserLeftBrace:
		{
			p.SetState(105)
			p.Object()
		}

	case JsonTemplateParserSTRING:
		{
			p.SetState(106)
			p.Name()
		}

	case JsonTemplateParserSubtract:
		{
			p.SetState(107)
			p.Match(JsonTemplateParserSubtract)
		}
		{
			p.SetState(108)
			p.field(15)
		}

	case JsonTemplateParserNot:
		{
			p.SetState(109)
			p.Match(JsonTemplateParserNot)
		}
		{
			p.SetState(110)
			p.field(2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(183)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
			case 1:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(113)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(114)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JsonTemplateParserMultiply || _la == JsonTemplateParserDivide) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(115)
					p.field(15)
				}

			case 2:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(116)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(117)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JsonTemplateParserAdd || _la == JsonTemplateParserSubtract) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(118)
					p.field(14)
				}

			case 3:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(119)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(120)
					p.Match(JsonTemplateParserRange)
				}
				{
					p.SetState(121)
					p.field(13)
				}

			case 4:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(122)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(123)
					p.Match(JsonTemplateParserNullCoalescing)
				}
				{
					p.SetState(124)
					p.field(12)
				}

			case 5:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(125)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(126)
					p.Match(JsonTemplateParserEqual)
				}
				{
					p.SetState(127)
					p.field(11)
				}

			case 6:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(128)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(129)
					p.Match(JsonTemplateParserLess)
				}
				{
					p.SetState(130)
					p.field(10)
				}

			case 7:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(131)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(132)
					p.Match(JsonTemplateParserLessOrEqual)
				}
				{
					p.SetState(133)
					p.field(9)
				}

			case 8:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(134)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(135)
					p.Match(JsonTemplateParserGreater)
				}
				{
					p.SetState(136)
					p.field(8)
				}

			case 9:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(137)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(138)
					p.Match(JsonTemplateParserGreaterOrEqual)
				}
				{
					p.SetState(139)
					p.field(7)
				}

			case 10:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(140)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(141)
					p.Match(JsonTemplateParserNotEqual)
				}
				{
					p.SetState(142)
					p.field(6)
				}

			case 11:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(143)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(144)
					p.Match(JsonTemplateParserAnd)
				}
				{
					p.SetState(145)
					p.field(5)
				}

			case 12:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(146)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(147)
					p.Match(JsonTemplateParserOr)
				}
				{
					p.SetState(148)
					p.field(4)
				}

			case 13:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(149)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}

				p.SetState(151)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonTemplateParserQuestion {
					{
						p.SetState(150)
						p.Match(JsonTemplateParserQuestion)
					}

				}
				{
					p.SetState(153)
					p.Match(JsonTemplateParserT__0)
				}
				{
					p.SetState(154)
					p.Name()
				}

			case 14:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(155)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}

				p.SetState(157)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonTemplateParserQuestion {
					{
						p.SetState(156)
						p.Match(JsonTemplateParserQuestion)
					}

				}
				{
					p.SetState(159)
					p.Match(JsonTemplateParserLeftBracket)
				}
				{
					p.SetState(160)
					p.Index()
				}
				{
					p.SetState(161)
					p.Match(JsonTemplateParserRightBracket)
				}

			case 15:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(163)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(164)
					p.Match(JsonTemplateParserLeftParen)
				}
				p.SetState(173)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67914508288) != 0 {
					{
						p.SetState(165)
						p.Function_param()
					}
					p.SetState(170)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == JsonTemplateParserComma {
						{
							p.SetState(166)
							p.Match(JsonTemplateParserComma)
						}
						{
							p.SetState(167)
							p.Function_param()
						}

						p.SetState(172)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(175)
					p.Match(JsonTemplateParserRightParen)
				}

			case 16:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JsonTemplateParserRULE_field)
				p.SetState(176)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(177)
					p.Match(JsonTemplateParserQuestion)
				}
				{
					p.SetState(178)
					p.field(0)
				}
				p.SetState(181)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(179)
						p.Match(JsonTemplateParserT__1)
					}
					{
						p.SetState(180)
						p.field(0)
					}

				}

			}

		}
		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
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
		p.SetState(188)
		p.Match(JsonTemplateParserLeftBracket)
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67914508288) != 0 {
		{
			p.SetState(189)
			p.field(0)
		}
		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonTemplateParserComma {
			{
				p.SetState(190)
				p.Match(JsonTemplateParserComma)
			}
			{
				p.SetState(191)
				p.field(0)
			}

			p.SetState(196)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(199)
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
		p.SetState(201)
		p.Match(JsonTemplateParserLeftBrace)
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JsonTemplateParserESCAPED_STRING || _la == JsonTemplateParserSTRING {
		{
			p.SetState(202)
			p.Object_field()
		}
		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonTemplateParserComma {
			{
				p.SetState(203)
				p.Match(JsonTemplateParserComma)
			}
			{
				p.SetState(204)
				p.Object_field()
			}

			p.SetState(209)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(212)
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

	p.SetState(221)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonTemplateParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)
			p.Name()
		}
		{
			p.SetState(215)
			p.Match(JsonTemplateParserT__1)
		}
		{
			p.SetState(216)
			p.field(0)
		}

	case JsonTemplateParserESCAPED_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(218)
			p.Match(JsonTemplateParserESCAPED_STRING)
		}
		{
			p.SetState(219)
			p.Match(JsonTemplateParserT__1)
		}
		{
			p.SetState(220)
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
		p.SetState(223)
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

	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(225)
			p.field(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(226)
			p.Match(JsonTemplateParserNUMBER)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(227)
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
