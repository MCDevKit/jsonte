package jsonte

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/gammazero/deque"
	"sync"
)

// expressionCache stores parsed expression trees keyed by their original text.
var expressionCache sync.Map

// exprVisitorPool pools ExpressionVisitor objects to avoid per-call heap escapes.
var exprVisitorPool = sync.Pool{
	New: func() interface{} { return &ExpressionVisitor{} },
}

// Result is the result of evaluating an expression
type Result struct {
	Value         types.JsonType
	Action        types.JsonAction
	Name          string
	IndexName     string
	Scope         deque.Deque[*types.JsonObject]
	VariableScope *types.JsonObject
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
	return Eval(text, deque.Deque[*types.JsonObject]{}, path)
}

// ClearExpressionCache removes all cached parse trees. Mainly intended for tests.
func ClearExpressionCache() {
	expressionCache = sync.Map{}
}

// ExpressionCacheSize returns the number of cached parse trees.
func ExpressionCacheSize() int {
	size := 0
	expressionCache.Range(func(_, _ interface{}) bool {
		size++
		return true
	})
	return size
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
func Eval(text string, scope deque.Deque[*types.JsonObject], path string) (Result, error) {
	var tree antlr.ParseTree
	if cached, ok := expressionCache.Load(text); ok {
		tree = cached.(antlr.ParseTree)
	} else {
		listener := CollectingErrorListener{DefaultErrorListener: antlr.NewDefaultErrorListener()}
		is := antlr.NewInputStream(text)
		lexer := NewTemplateAwareLexer(is)
		lexer.JsonTemplateLexer.RemoveErrorListeners()
		lexer.JsonTemplateLexer.AddErrorListener(&listener)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewJsonTemplate(stream)
		p.RemoveErrorListeners()
		p.AddErrorListener(&listener)
		p.BuildParseTrees = true
		tree = p.Expression()
		if listener.Error != nil {
			return Result{}, burrito.WrapErrorf(listener.Error, "Failed to parse expression \"%s\"", text)
		}
		expressionCache.Store(text, tree)
	}
	visitor := ExpressionVisitor{
		scope:         scope,
		name:          DefaultName,
		indexName:     DefaultIndexName,
		variableScope: types.NewJsonObject(),

	}
	r, err := visitor.Visit(tree)
	return Result{
		Value:         r,
		Action:        visitor.action,
		Name:          visitor.name,
		IndexName:     visitor.indexName,
		VariableScope: visitor.variableScope,
		Scope:         scope,
	}, err
}

// EvalScript evaluates the given script and returns the result
func EvalScript(text string, scope deque.Deque[*types.JsonObject], path string) (Result, error) {
	listener := CollectingErrorListener{DefaultErrorListener: antlr.NewDefaultErrorListener()}
	is := antlr.NewInputStream(text)
	lexer := NewTemplateAwareLexer(is)
	lexer.JsonTemplateLexer.RemoveErrorListeners()
	lexer.JsonTemplateLexer.AddErrorListener(&listener)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewJsonTemplate(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(&listener)
	p.BuildParseTrees = true
	tree := p.Script()
	if listener.Error != nil {
		return Result{}, burrito.WrapErrorf(listener.Error, "Failed to parse script \"%s\"", text)
	}
	visitor := ExpressionVisitor{
		scope:         scope,
		name:          DefaultName,
		indexName:     DefaultIndexName,
		variableScope: types.NewJsonObject(),

	}
	r, err := visitor.Visit(tree)
	return Result{
		Value:         r,
		Action:        visitor.action,
		VariableScope: visitor.variableScope,
		Scope:         scope,
	}, err
}

// evalWithExtraScope evaluates text using scope as the primary scope and extraScope as secondary,
// avoiding the allocation of a merged deque copy.
func evalWithExtraScope(text string, scope *deque.Deque[*types.JsonObject], extraScope *deque.Deque[*types.JsonObject]) (Result, error) {
	var tree antlr.ParseTree
	if cached, ok := expressionCache.Load(text); ok {
		tree = cached.(antlr.ParseTree)
	} else {
		listener := CollectingErrorListener{DefaultErrorListener: antlr.NewDefaultErrorListener()}
		is := antlr.NewInputStream(text)
		lexer := NewTemplateAwareLexer(is)
		lexer.JsonTemplateLexer.RemoveErrorListeners()
		lexer.JsonTemplateLexer.AddErrorListener(&listener)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewJsonTemplate(stream)
		p.RemoveErrorListeners()
		p.AddErrorListener(&listener)
		p.BuildParseTrees = true
		tree = p.Expression()
		if listener.Error != nil {
			return Result{}, burrito.WrapErrorf(listener.Error, "Failed to parse expression \"%s\"", text)
		}
		expressionCache.Store(text, tree)
	}
	v := exprVisitorPool.Get().(*ExpressionVisitor)
	v.scope = *scope
	v.extraScope = extraScope
	v.name = DefaultName
	v.indexName = DefaultIndexName
	v.action = 0
	v.variableScope = nil
	v.globalScope = nil
	v.usedVariables = v.usedVariables[:0]

	r, err := v.Visit(tree)
	result := Result{
		Value:         r,
		Action:        v.action,
		Name:          v.name,
		IndexName:     v.indexName,
		VariableScope: v.variableScope,
		Scope:         *scope,
	}
	// Wyczyść referencje przed zwrotem do puli.
	v.scope = deque.Deque[*types.JsonObject]{}
	v.extraScope = nil
	v.variableScope = nil
	exprVisitorPool.Put(v)
	return result, err
}

// EvalWithTempScope evaluates the given expression and returns the result
func EvalWithTempScope(text string, scope deque.Deque[*types.JsonObject], path string, temp ...*types.JsonObject) (Result, error) {
	d := deque.Deque[*types.JsonObject]{}
	for i := 0; i < scope.Len(); i++ {
		d.PushBack(scope.At(i))
	}
	for _, t := range temp {
		d.PushBack(t)
	}
	return Eval(text, d, path)
}

// EvalJsonteFile runs a .jsonte script against the provided global scope object.
// The scope is NOT deep-copied, so mutations to existing scope keys persist automatically.
// New top-level scope keys can be added via `$scope.key = value`.
// Script-local variables (loop counters, temporaries) go to variableScope and are discarded.
func EvalJsonteFile(text string, scope *types.JsonObject, path string) error {
	listener := CollectingErrorListener{DefaultErrorListener: antlr.NewDefaultErrorListener()}
	is := antlr.NewInputStream(text)
	lexer := NewTemplateAwareLexer(is)
	lexer.JsonTemplateLexer.RemoveErrorListeners()
	lexer.JsonTemplateLexer.AddErrorListener(&listener)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewJsonTemplate(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(&listener)
	p.BuildParseTrees = true
	tree := p.Script()
	if listener.Error != nil {
		return burrito.WrapErrorf(listener.Error, "Failed to parse .jsonte file \"%s\"", path)
	}
	d := deque.Deque[*types.JsonObject]{}
	d.PushBack(scope)
	visitor := ExpressionVisitor{
		scope:         d,
		name:          DefaultName,
		indexName:     DefaultIndexName,
		variableScope: types.NewJsonObject(),

		globalScope:   scope,
	}
	_, err := visitor.Visit(tree)
	if err != nil {
		return burrito.PassError(err)
	}
	return nil
}

func ParseLambda(text string) ([]string, []string, error) {
	listener := CollectingErrorListener{DefaultErrorListener: antlr.NewDefaultErrorListener()}
	is := antlr.NewInputStream(text)
	lexer := NewTemplateAwareLexer(is)
	lexer.JsonTemplateLexer.RemoveErrorListeners()
	lexer.JsonTemplateLexer.AddErrorListener(&listener)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewJsonTemplate(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(&listener)
	p.BuildParseTrees = true
	tree := p.Lambda()
	if listener.Error != nil {
		return nil, nil, burrito.WrapErrorf(listener.Error, "Failed to parse lambda \"%s\"", text)
	}
	visitor := LambdaVisitor{}
	visitor.Visit(tree)
	return visitor.usedVariables, visitor.arguments, nil
}


