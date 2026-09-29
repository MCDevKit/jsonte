package jsonte

import (
	"sync"

	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/gammazero/deque"
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
	fieldInfoCache = sync.Map{}
}

// startRule invokes a parser start rule and returns the resulting tree.
type startRule func(p *parser.JsonTemplate) antlr.ParseTree

func expressionRule(p *parser.JsonTemplate) antlr.ParseTree { return p.Expression() }
func scriptRule(p *parser.JsonTemplate) antlr.ParseTree     { return p.Script() }
func lambdaRule(p *parser.JsonTemplate) antlr.ParseTree     { return p.Lambda() }

// newParser builds a lexer, token stream and parser for text with listener attached to both.
func newParser(text string, listener antlr.ErrorListener) *parser.JsonTemplate {
	is := antlr.NewInputStream(text)
	lexer := NewTemplateAwareLexer(is)
	lexer.JsonTemplateLexer.RemoveErrorListeners()
	lexer.JsonTemplateLexer.AddErrorListener(listener)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewJsonTemplate(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(listener)
	p.BuildParseTrees = true
	return p
}

// parseSLL tries to parse text with the fast SLL prediction mode and a bail-out error strategy.
// It returns nil if the parse failed for any reason (lexer error, syntax error or an ambiguity
// that SLL cannot resolve), in which case the caller must re-parse in full LL mode.
func parseSLL(text string, rule startRule) (tree antlr.ParseTree) {
	listener := CollectingErrorListener{DefaultErrorListener: antlr.NewDefaultErrorListener()}
	p := newParser(text, &listener)
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	p.SetErrorHandler(antlr.NewBailErrorStrategy())
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(*antlr.ParseCancellationException); !ok {
				panic(r)
			}
			tree = nil
		}
	}()
	tree = rule(p)
	if listener.Error != nil {
		return nil
	}
	return tree
}

// parseTree parses text using rule. It first tries the fast SLL mode; if that fails, it
// re-parses from scratch in the default LL mode with the default error strategy and reports
// errors to listener, so error messages are identical to a plain LL parse.
func parseTree(text string, rule startRule, listener antlr.ErrorListener) antlr.ParseTree {
	if tree := parseSLL(text, rule); tree != nil {
		return tree
	}
	return rule(newParser(text, listener))
}

// parseExpressionTree parses text as an expression.
func parseExpressionTree(text string) (antlr.ParseTree, error) {
	listener := CollectingErrorListener{DefaultErrorListener: antlr.NewDefaultErrorListener()}
	tree := parseTree(text, expressionRule, &listener)
	return tree, listener.Error
}

// cachedExpressionTree returns the parse tree of the expression text, parsing and caching it if needed.
func cachedExpressionTree(text string) (antlr.ParseTree, error) {
	if cached, ok := expressionCache.Load(text); ok {
		return cached.(antlr.ParseTree), nil
	}
	tree, err := parseExpressionTree(text)
	if err != nil {
		return nil, burrito.WrapErrorf(err, "Failed to parse expression \"%s\"", text)
	}
	expressionCache.Store(text, tree)
	return tree, nil
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
	tree, err := cachedExpressionTree(text)
	if err != nil {
		return Result{}, err
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
	tree := parseTree(text, scriptRule, &listener)
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
	tree, err := cachedExpressionTree(text)
	if err != nil {
		return Result{}, err
	}
	v := exprVisitorPool.Get().(*ExpressionVisitor)
	v.scope = *scope
	v.extraScope = extraScope
	v.name = DefaultName
	v.indexName = DefaultIndexName
	v.action = 0
	v.variableScope = nil
	v.localVarTarget = nil
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
	// Clear reference before returning to the pool
	v.scope = deque.Deque[*types.JsonObject]{}
	v.extraScope = nil
	v.variableScope = nil
	v.localVarTarget = nil
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
	tree := parseTree(text, scriptRule, &listener)
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

		globalScope: scope,
	}
	_, err := visitor.Visit(tree)
	if err != nil {
		return burrito.PassError(err)
	}
	return nil
}

func ParseLambda(text string) ([]string, []string, error) {
	listener := CollectingErrorListener{DefaultErrorListener: antlr.NewDefaultErrorListener()}
	tree := parseTree(text, lambdaRule, &listener)
	if listener.Error != nil {
		return nil, nil, burrito.WrapErrorf(listener.Error, "Failed to parse lambda \"%s\"", text)
	}
	visitor := LambdaVisitor{}
	visitor.Visit(tree)
	return visitor.usedVariables, visitor.arguments, nil
}
