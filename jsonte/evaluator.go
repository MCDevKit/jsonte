package jsonte

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/gammazero/deque"
)

// Result is the result of evaluating an expression
type Result struct {
	Value     types.JsonType
	Action    types.JsonAction
	Name      string
	IndexName string
}

// GetError returns the error from the result or nil if the expression evaluated correctly
func (r *Result) GetError() error {
	if isError(r.Value) {
		return r.Value.(error)
	} else {
		return nil
	}
}

// QuickEval is a convenience function for evaluating a single expression
func QuickEval(text string, path string) (Result, error) {
	return Eval(text, deque.Deque[types.JsonObject]{}, path)
}

// CollectingErrorListener is an error listener that collects all errors by appending them to the Error field
type CollectingErrorListener struct {
	*antlr.DefaultErrorListener
	Error error
}

func (l *CollectingErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.Error = burrito.WrapErrorf(l.Error, "column: %d %s", column, msg)
}

// Eval evaluates the given expression and returns the result
func Eval(text string, scope deque.Deque[types.JsonObject], path string) (Result, error) {
	listener := CollectingErrorListener{DefaultErrorListener: antlr.NewDefaultErrorListener()}
	is := antlr.NewInputStream(text)
	lexer := parser.NewJsonTemplateLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&listener)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewJsonTemplateParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(&listener)
	p.BuildParseTrees = true
	tree := p.Expression()
	if listener.Error != nil {
		return Result{}, burrito.WrapErrorf(listener.Error, "Failed to parse expression \"%s\"", text)
	}
	visitor := ExpressionVisitor{
		scope: scope,
		path:  &path,
	}
	r, err := visitor.Visit(tree)
	return Result{
		Value:     r,
		Action:    visitor.action,
		Name:      *visitor.name,
		IndexName: *visitor.indexName,
	}, err
}

// EvalWithTempScope evaluates the given expression and returns the result
func EvalWithTempScope(text string, scope deque.Deque[types.JsonObject], path string, temp ...types.JsonObject) (Result, error) {
	d := deque.Deque[types.JsonObject]{}
	for i := 0; i < scope.Len(); i++ {
		d.PushBack(scope.At(i))
	}
	for _, t := range temp {
		d.PushBack(t)
	}
	return Eval(text, d, path)
}

func ParseLambda(text string) ([]string, []string, error) {
	listener := CollectingErrorListener{DefaultErrorListener: antlr.NewDefaultErrorListener()}
	is := antlr.NewInputStream(text)
	lexer := parser.NewJsonTemplateLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&listener)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewJsonTemplateParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(&listener)
	p.BuildParseTrees = true
	tree := p.Lambda()
	if listener.Error != nil {
		return nil, nil, burrito.WrapErrorf(listener.Error, "Failed to parse expression \"%s\"", text)
	}
	visitor := LambdaVisitor{}
	visitor.Visit(tree)
	return visitor.usedVariables, visitor.arguments, nil
}
