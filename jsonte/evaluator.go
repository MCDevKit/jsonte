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
	Name   string
}

func (r *Result) GetError() error {
	if isError(r.Value) {
		return r.Value.(error)
	} else {
		return nil
	}
}

func QuickEval(text string, path string) (Result, error) {
	return Eval(text, deque.Deque[interface{}]{}, path)
}

func Eval(text string, scope deque.Deque[interface{}], path string) (Result, error) {
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
		scope: scope,
		path:  &path,
	}
	r := visitor.Visit(tree)
	var err error
	if isError(r) {
		err = r.(error)
	}
	return Result{
		Value:  r,
		Action: visitor.action,
		Name:   *visitor.name,
	}, err
}
