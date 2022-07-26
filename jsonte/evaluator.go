package jsonte

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"jsonte/jsonte/utils"
	"jsonte/parser"
)

type Result struct {
	Value  interface{}
	Action utils.JsonAction
}

func Eval(text string) (result Result) {
	is := antlr.NewInputStream(text)
	lexer := parser.NewJsonTemplateLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(antlr.NewConsoleErrorListener())
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewJsonTemplateParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewConsoleErrorListener())
	p.BuildParseTrees = true
	tree := p.Expression()
	path := "#/"
	visitor := ExpressionVisitor{
		fullScope: map[string]interface{}{
			"test": 123,
		},
		extraScope: map[string]interface{}{
			"test2": 456,
		},
		path: &path,
	}
	r := visitor.Visit(tree)
	fmt.Println(utils.ToString(r))
	return Result{
		Value:  r,
		Action: visitor.action,
	}
}
