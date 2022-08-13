package jsonte

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gammazero/deque"
	"jsonte/jsonte/utils"
	"jsonte/parser"
)

type Result struct {
	Value  interface{}
	Action utils.JsonAction
}

func (r *Result) GetError() error {
	if isError(r.Value) {
		return r.Value.(error)
	} else {
		return nil
	}
}

func Eval(text string, extraScope, fullScope utils.JsonObject, thisInstance deque.Deque[interface{}], path string) (Result, error) {
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
	visitor := ExpressionVisitor{
		fullScope:    fullScope,
		extraScope:   extraScope,
		currentScope: thisInstance,
		path:         &path,
	}
	r := visitor.Visit(tree)
	var err error
	if isError(r) {
		err = r.(error)
	}
	return Result{
		Value:  r,
		Action: visitor.action,
	}, err
}
