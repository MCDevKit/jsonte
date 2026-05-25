package jsonte

import (
	"reflect"
	"strings"
	"sync"

	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/functions"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/MCDevKit/jsonte/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/gammazero/deque"
)

const DefaultName = "value"
const DefaultIndexName = "index"

// jsonStringInternCache caches *JsonString wrappers for frequently-used identifier names.
// This avoids repeated allocations in ResolveScope where the same names are looked up
// thousands of times during array iteration.
var jsonStringInternCache sync.Map

// lambdaStringCache caches parsed lambda AST trees for string-lambda arguments.
// Keyed by the raw lambda source string.
var lambdaStringCache sync.Map

// lambdaCtxInfo holds precomputed metadata for a LambdaContext.
type lambdaCtxInfo struct {
	text   string
	vars   []string
	args   []string // lv.arguments — all args including nested lambdas; used for JsonLambda.Arguments
	params []string // direct outer params only (ctx.AllName()); used for scope push/pop and arity check
}

// lambdaInfoCache caches metadata for LambdaContext pointers from the expression cache.
// Since the expression cache keeps parse trees alive, the *LambdaContext pointer is stable.
var lambdaInfoCache sync.Map

// nodeTextCache caches the GetText() result for stable parse-tree nodes.
// ANTLR creates an Interval on every GetText() call (never memoizes), so caching here
// saves one heap allocation per node per visit.
var nodeTextCache sync.Map

// spreadFieldCache caches the AllSpread_field() slice per ArrayContext pointer.
// ANTLR's All* methods allocate a new slice each call; caching avoids this for
// expressions that appear repeatedly inside loops.
var spreadFieldCache sync.Map

// nodeNumberCache caches *types.JsonNumber per NUMBER terminal node pointer.
// NUMBER literals in expressions are never indexed (no UpdateParent call), so
// sharing the same *JsonNumber across evaluations is safe.
var nodeNumberCache sync.Map

// cachedGetText returns GetText() for any stable parse-tree node, caching the result.
func cachedGetText(node antlr.ParseTree) string {
	if v, ok := nodeTextCache.Load(node); ok {
		return v.(string)
	}
	text := node.GetText()
	nodeTextCache.Store(node, text)
	return text
}

// cachedNodeNumber returns a *types.JsonNumber for a NUMBER terminal node, caching by node pointer.
func cachedNodeNumber(node antlr.TerminalNode) *types.JsonNumber {
	if v, ok := nodeNumberCache.Load(node); ok {
		return v.(*types.JsonNumber)
	}
	n := types.AsNumber(node.GetText())
	nodeNumberCache.Store(node, n)
	return n
}

func internedJsonString(s string) *types.JsonString {
	if v, ok := jsonStringInternCache.Load(s); ok {
		return v.(*types.JsonString)
	}
	js := types.NewString(s)
	actual, _ := jsonStringInternCache.LoadOrStore(s, js)
	return actual.(*types.JsonString)
}

var scopeObjectPool = sync.Pool{
	New: func() interface{} { return types.NewJsonObjectWithCapacity(2) },
}

func getScopeObject() *types.JsonObject {
	return scopeObjectPool.Get().(*types.JsonObject)
}

func putScopeObject(obj *types.JsonObject) {
	obj.Reset()
	scopeObjectPool.Put(obj)
}

type ExpressionVisitor struct {
	parser.BaseJsonTemplateVisitor
	action    types.JsonAction
	name      string
	indexName string
	scope     deque.Deque[*types.JsonObject]
	// extraScope holds an additional scope layer (e.g. moduleScope from TemplateVisitor)
	// searched after scope so that scope entries take priority.
	extraScope    *deque.Deque[*types.JsonObject]
	usedVariables []string
	variableScope *types.JsonObject
	// globalScope, when non-nil, enables the $scope pseudo-variable in .jsonte scripts.
	// $scope.key = value writes directly into the global scope object.
	globalScope *types.JsonObject
}

func (v *ExpressionVisitor) Visit(tree antlr.ParseTree) (types.JsonType, error) {
	var result types.JsonType
	var err error
	switch val := tree.(type) {
	case *parser.FieldContext:
		result, err = v.VisitField(val)
		break
	case *parser.ArrayContext:
		result, err = v.VisitArray(val)
		break
	case *parser.ObjectContext:
		result, err = v.VisitObject(val)
		break
	case *parser.Object_fieldContext:
		result, err = v.VisitObject_field(val)
		break
	case *parser.ExpressionContext:
		result, err = v.VisitExpression(val)
		break
	case *parser.Function_paramContext:
		result, err = v.VisitFunction_param(val)
		break
	case *parser.LambdaContext:
		result, err = v.VisitLambda(val)
		break
	case *parser.NameContext:
		result, err = v.VisitName(val)
		break
	case *parser.IndexContext:
		result, err = v.VisitIndex(val)
		break
	case *parser.StatementContext:
		result, err = v.VisitStatement(val)
		break
	case *parser.StatementsContext:
		result, err = v.VisitStatements(val)
		break
	case *parser.ScriptContext:
		result, err = v.VisitScript(val)
		break
	case *parser.Template_stringContext:
		result, err = v.VisitTemplate_string(val)
		break
	case *parser.Template_string_partContext:
		result, err = v.VisitTemplate_string_part(val)
		break
	default:
		utils.BadDeveloperError("Unknown tree type " + reflect.TypeOf(tree).String())
	}
	//utils.Logger.Debugf("Expression %s evaluated to %s", tree.GetText(), result.StringValue())
	return result, err
}

func treeMatches(context antlr.Tree, matchFunction func(ctx, parent antlr.Tree) bool) bool {
	if context == nil {
		return false
	}
	counter := 0
	for {
		if context.GetParent() == nil {
			return false
		}
		if matchFunction(context, context.GetParent()) {
			return true
		}
		context = context.GetParent()
		counter++
		if counter > 100 {
			utils.BadDeveloperError("Too many loops in treeMatches")
		}
	}
}

func isInLeftSideOfAssignment(context antlr.Tree) bool {
	return treeMatches(context, func(ctx, parent antlr.Tree) bool {
		if f, ok := parent.(*parser.FieldContext); ok && len(f.AllField()) == 2 && f.Field(0) == ctx {
			if f.Literal() != nil || f.AddAssign() != nil {
				return true
			}
		}
		return false
	})
}

// resolveLambdaTree resolves a string to an AST tree
func (v *ExpressionVisitor) resolveLambdaTree(src string) parser.ILambdaContext {
	if cached, ok := lambdaStringCache.Load(src); ok {
		return cached.(parser.ILambdaContext)
	}
	is := antlr.NewInputStream(src)
	lexer := NewTemplateAwareLexer(is)
	lexer.JsonTemplateLexer.RemoveErrorListeners()
	lexer.JsonTemplateLexer.AddErrorListener(antlr.NewConsoleErrorListener())
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewJsonTemplate(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewConsoleErrorListener())
	p.BuildParseTrees = true
	ctx := p.Lambda()
	lambdaStringCache.Store(src, ctx)
	return ctx
}

// pushScope pushes a new scope to the stack
func (v *ExpressionVisitor) pushScope(scope map[string]interface{}) {
	v.scope.PushBack(types.AsObject(scope))
}

// pushScopePair pushes a single key-value pair onto the scope stack without reflection overhead.
func (v *ExpressionVisitor) pushScopePair(key string, value interface{}) {
	obj := getScopeObject()
	obj.Put(key, types.Box(value))
	v.scope.PushBack(obj)
}

// popScope pops the last scope from the stack
func (v *ExpressionVisitor) popScope() {
	v.scope.PopBack()
}

func (v *ExpressionVisitor) popScopeObject() {
	obj := v.scope.Back()
	v.scope.PopBack()
	putScopeObject(obj)
}

func isError(v interface{}) bool {
	_, err := v.(error)
	return err
}

// initVarScope lazily initializes the variable scope on first use.
func (v *ExpressionVisitor) initVarScope() *types.JsonObject {
	if v.variableScope == nil {
		v.variableScope = types.NewJsonObject()
	}
	return v.variableScope
}

// ResolveScope resolves a value from the scope by name
func (v *ExpressionVisitor) ResolveScope(name string) types.JsonType {
	// $scope is a write-through proxy to the global scope, available only in .jsonte scripts.
	if name == "$scope" && v.globalScope != nil {
		return &types.JsonObject{
			StackValue:  &v.scope,
			StackTarget: v.globalScope,
		}
	}
	if name == "this" {
		return &types.JsonObject{
			StackValue:  &v.scope,
			StackTarget: v.initVarScope(),
		}
	}
	nameStr := internedJsonString(name)
	if v.variableScope != nil && v.variableScope.ContainsKey(name) {
		get := v.variableScope.Get(name)
		get.UpdateParent(v.variableScope, nameStr)
		return get
	}
	for i := v.scope.Len() - 1; i >= 0; i-- {
		m := v.scope.At(i)
		if m.ContainsKey(name) {
			get := m.Get(name)
			get.UpdateParent(m, nameStr)
			return get
		}
	}
	if v.extraScope != nil {
		for i := v.extraScope.Len() - 1; i >= 0; i-- {
			m := v.extraScope.At(i)
			if m.ContainsKey(name) {
				get := m.Get(name)
				get.UpdateParent(m, nameStr)
				return get
			}
		}
	}
	return types.NullWithParent(v.initVarScope(), nameStr)
}

func (v *ExpressionVisitor) VisitScript(ctx *parser.ScriptContext) (types.JsonType, error) {
	for _, statement := range ctx.AllStatement() {
		visit, err := v.Visit(statement)
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		if types.IsReturn(visit) {
			return visit, nil
		}
		if visit == types.Break || visit == types.Continue {
			return types.Null, burrito.WrappedErrorf("%s outside of loop", visit.StringValue())
		}
	}
	return types.Null, nil
}

func (v *ExpressionVisitor) VisitStatement(ctx *parser.StatementContext) (types.JsonType, error) {
	if ctx.Return() != nil {
		if len(ctx.AllField()) == 0 {
			return types.NewReturn(types.Null), nil
		}
		visit, err := v.Visit(ctx.Field(0))
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		return types.NewReturn(visit), nil
	} else if ctx.Break() != nil {
		return types.Break, nil
	} else if ctx.Continue() != nil {
		return types.Continue, nil
	} else if ctx.For() != nil && ctx.In() != nil {
		arr, err := v.Visit(ctx.Field(0))
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		if _, ok := arr.(*types.JsonArray); !ok {
			return types.Null, burrito.WrappedErrorf("Cannot iterate over %s", arr.StringValue())
		}
		valueName := ctx.Name(0).GetText()
		indexName := ""
		hasIndex := false
		if len(ctx.AllName()) > 1 {
			indexName = ctx.Name(1).GetText()
			hasIndex = true
		}
		for i, value := range arr.(*types.JsonArray).Value {
			v.pushScopePair(valueName, value)
			if hasIndex {
				v.pushScopePair(indexName, types.NewIntNumber(int64(i)))
			}
			val, err := v.Visit(ctx.Statements(0))
			if hasIndex {
				v.popScopeObject()
			}
			v.popScopeObject()
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			if val != nil && types.IsReturn(val) {
				return val, nil
			}
			if val == types.Break {
				break
			}
			if val == types.Continue {
				continue
			}
		}
	} else if len(ctx.AllIf()) > 0 {
		for i := range ctx.AllIf() {
			condition, err := v.Visit(ctx.Field(i))
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			if condition.BoolValue() {
				val, err := v.Visit(ctx.Statements(i))
				if err != nil {
					return types.Null, burrito.PassError(err)
				}
				return val, nil
			}
		}
		if len(ctx.AllElse()) == len(ctx.AllIf()) {
			val, err := v.Visit(ctx.Statements(len(ctx.AllIf())))
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			return val, nil
		}
	} else if ctx.Do() != nil {
		for {
			val, err := v.Visit(ctx.Statements(0))
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			if val != nil && types.IsReturn(val) {
				return val, nil
			}
			if val == types.Break {
				break
			}
			if val == types.Continue {
				continue
			}
			condition, err := v.Visit(ctx.Field(0))
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			if !condition.BoolValue() {
				return types.Null, nil
			}
		}
	} else if ctx.While() != nil {
		for {
			condition, err := v.Visit(ctx.Field(0))
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			if !condition.BoolValue() {
				return types.Null, nil
			}
			val, err := v.Visit(ctx.Statements(0))
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			if val != nil && types.IsReturn(val) {
				return val, nil
			}
			if val == types.Break {
				break
			}
			if val == types.Continue {
				continue
			}
		}
	} else if len(ctx.AllField()) == 1 {
		return v.Visit(ctx.Field(0))
	} else if len(ctx.AllStatements()) == 1 {
		return v.Visit(ctx.Statements(0))
	}
	return types.Null, nil
}

func (v *ExpressionVisitor) VisitStatements(ctx *parser.StatementsContext) (types.JsonType, error) {
	for _, statement := range ctx.AllStatement() {
		v, err := v.Visit(statement)
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		if types.IsReturn(v) {
			return v, nil
		}
		if v == types.Break || v == types.Continue {
			return v, nil
		}
	}
	return types.Null, nil
}

func (v *ExpressionVisitor) VisitExpression(ctx *parser.ExpressionContext) (types.JsonType, error) {
	// check the type of the expression
	v.action = types.Value
	if ctx.Iteration() != nil {
		v.action = types.Iteration
	}
	if ctx.Literal() != nil {
		v.action = types.Literal
	}
	if ctx.Question() != nil {
		v.action = types.Predicate
	}
	// check if the iteration value is named
	name := DefaultName
	indexName := DefaultIndexName
	if ctx.As() != nil {
		if len(ctx.AllName()) > 0 {
			name = ctx.Name(0).GetText()
		}
		if len(ctx.AllName()) > 1 {
			indexName = ctx.Name(1).GetText()
		}
	}
	v.name = name
	v.indexName = indexName
	result, err := v.Visit(ctx.Field())
	if err != nil {
		return types.Null, burrito.PassError(err)
	}
	if _, ok := result.(*types.JsonLambda); ok {
		switch v.action {
		case types.Iteration:
			return types.Null, burrito.WrappedErrorf("Cannot iterate over a lambda")
		case types.Predicate:
			return types.Null, burrito.WrappedErrorf("Cannot use a lambda as a predicate")
		case types.Literal:
			return types.Null, burrito.WrappedErrorf("Cannot use a lambda as a literal value")
		case types.Value:
			return types.Null, burrito.WrappedErrorf("Cannot use a lambda as a template value")
		}
	}
	return result, nil
}

func (v *ExpressionVisitor) VisitField(context *parser.FieldContext) (types.JsonType, error) {
	if context.Null() != nil {
		return types.Null, nil
	}
	if context.True() != nil {
		return types.True(), nil
	}
	if context.False() != nil {
		return types.False(), nil
	}
	if context.Not() != nil {
		visit, err := v.Visit(context.Field(0))
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		return types.AsBool(!visit.BoolValue()), nil
	}
	if context.AddAssign() != nil {
		left, err := v.Visit(context.Field(0))
		if err != nil {
			return types.Null, burrito.WrapErrorf(err, "Error resolving left side of the statement")
		}
		right, err := v.Visit(context.Field(1))
		if err != nil {
			return types.Null, burrito.WrapErrorf(err, "Error resolving right side of the statement")
		}
		right = types.Box(left.Add(right).Unbox())
		if left.Parent() == nil {
			return types.Null, burrito.WrappedErrorf("Cannot assign a value to this expression")
		}
		if left.ParentIndex() == nil {
			utils.BadDeveloperError("Invalid parent index type")
		}
		if b, ok := (left.Parent()).(*types.JsonObject); ok {
			if i, ok1 := left.ParentIndex().(*types.JsonString); ok1 {
				b.Put(i.StringValue(), right)
				return right, nil
			} else if i, ok1 := left.ParentIndex().(*types.JsonPath); ok1 {
				return i.Set(left.Parent(), right)
			} else {
				return types.Null, burrito.WrappedErrorf("Invalid index type. Expected string or Json Path, but got %s", reflect.TypeOf(left.ParentIndex()).String())
			}
		} else if b, ok := (left.Parent()).(*types.JsonArray); ok {
			if i, ok1 := left.ParentIndex().(*types.JsonNumber); ok1 {
				index := int(i.IntValue())
				if index < 0 {
					index = len(b.Value) + index
				}
				if index >= 0 && index < len(b.Value) {
					b.Value[index] = right
					return right, nil
				} else {
					return types.Null, burrito.WrappedErrorf("Index out of bounds: %d", index)
				}
			} else if i, ok1 := left.ParentIndex().(*types.JsonPath); ok1 {
				return i.Set(left.Parent(), right)
			} else {
				return types.Null, burrito.WrappedErrorf("Invalid index type. Expected number or Json Path, but got %s", reflect.TypeOf(left.ParentIndex()).String())
			}
		} else {
			return types.Null, burrito.WrappedErrorf("Cannot assign a value to this expression")
		}
	}
	if context.Literal() != nil {
		left, err := v.Visit(context.Field(0))
		if err != nil {
			return types.Null, burrito.WrapErrorf(err, "Error resolving left side of the statement")
		}
		right, err := v.Visit(context.Field(1))
		if err != nil {
			return types.Null, burrito.WrapErrorf(err, "Error resolving right side of the statement")
		}
		// Preserve lambdas as-is; for other types, re-box to strip parent references.
		if _, ok := right.(*types.JsonLambda); !ok {
			right = types.Box(right.Unbox())
		}
		if left.Parent() == nil {
			return types.Null, burrito.WrappedErrorf("Cannot assign a value to this expression")
		}
		if left.ParentIndex() == nil {
			utils.BadDeveloperError("Invalid parent index type")
		}
		if b, ok := (left.Parent()).(*types.JsonObject); ok {
			if i, ok1 := left.ParentIndex().(*types.JsonString); ok1 {
				b.Put(i.StringValue(), right)
				return right, nil
			} else if i, ok1 := left.ParentIndex().(*types.JsonPath); ok1 {
				return i.Set(left.Parent(), right)
			} else {
				return types.Null, burrito.WrappedErrorf("Invalid index type. Expected string or Json Path, but got %s", reflect.TypeOf(left.ParentIndex()).String())
			}
		} else if b, ok := (left.Parent()).(*types.JsonArray); ok {
			if i, ok1 := left.ParentIndex().(*types.JsonNumber); ok1 {
				index := int(i.IntValue())
				if index < 0 {
					index = len(b.Value) + index
				}
				if index >= 0 && index < len(b.Value) {
					b.Value[index] = right
					return right, nil
				} else {
					return types.Null, burrito.WrappedErrorf("Index out of bounds: %d", index)
				}
			} else if i, ok1 := left.ParentIndex().(*types.JsonPath); ok1 {
				return i.Set(left.Parent(), right)
			} else {
				return types.Null, burrito.WrappedErrorf("Invalid index type. Expected number or Json Path, but got %s", reflect.TypeOf(left.ParentIndex()).String())
			}
		} else if b, ok := (left.Parent()).(*types.JsonPath); ok {
			if i, ok1 := left.ParentIndex().(*types.JsonNumber); ok1 {
				index := int(i.IntValue())
				if index < 0 {
					index = len(b.Path) + index
				}
				if index >= 0 && index < len(b.Path) {
					b.Path[index] = right
					return right, nil
				} else {
					return types.Null, burrito.WrappedErrorf("Index out of bounds: %d", index)
				}
			} else {
				utils.BadDeveloperError("Invalid parent index type")
			}
		} else {
			return types.Null, burrito.WrappedErrorf("Cannot assign a value to this expression")
		}
	}
	// process field composed of two other fields
	if len(context.AllField()) == 2 {
		f1, err := v.Visit(context.Field(0))
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		// Move AND and OR here to make those operators short-circuiting
		if context.And() != nil {
			if f1.BoolValue() {
				f2, err := v.Visit(context.Field(1))
				if err != nil {
					return types.Null, burrito.PassError(err)
				}
				return types.AsBool(f2.BoolValue()), nil
			} else {
				return types.False(), nil
			}
		} else if context.Or() != nil {
			if !f1.BoolValue() {
				f2, err := v.Visit(context.Field(1))
				if err != nil {
					return types.Null, burrito.PassError(err)
				}
				return types.AsBool(f2.BoolValue()), nil
			} else {
				return types.True(), nil
			}
		}
		f2, err := v.Visit(context.Field(1))
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		_, f1IsLambda := f1.(*types.JsonLambda)
		_, f2IsLambda := f2.(*types.JsonLambda)
		if context.NullCoalescing() != nil {
			if types.IsNull(f1) {
				return f2, nil
			} else {
				return f1, nil
			}
		} else if context.Add() != nil {
			if f1IsLambda || f2IsLambda {
				return types.Null, burrito.WrappedErrorf("Cannot use a lambda in an addition or concatenation")
			}
			return f1.Add(f2), nil
		} else if context.Equal() != nil {
			return types.AsBool(f1.Equals(f2)), nil
		} else if context.NotEqual() != nil {
			return types.AsBool(!f1.Equals(f2)), nil
		} else if context.Less() != nil {
			if f1IsLambda || f2IsLambda {
				return types.Null, burrito.WrappedErrorf("Cannot compare a lambda")
			}
			than, err := f1.LessThan(f2)
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			return types.AsBool(than), nil
		} else if context.Greater() != nil {
			if f1IsLambda || f2IsLambda {
				return types.Null, burrito.WrappedErrorf("Cannot compare a lambda")
			}
			than, err := f1.LessThan(f2)
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			return types.AsBool(!than && !f1.Equals(f2)), nil
		} else if context.LessOrEqual() != nil {
			if f1IsLambda || f2IsLambda {
				return types.Null, burrito.WrappedErrorf("Cannot compare a lambda")
			}
			than, err := f1.LessThan(f2)
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			return types.AsBool(than || f1.Equals(f2)), nil
		} else if context.GreaterOrEqual() != nil {
			if f1IsLambda || f2IsLambda {
				return types.Null, burrito.WrappedErrorf("Cannot compare a lambda")
			}
			than, err := f1.LessThan(f2)
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			return types.AsBool(!than), nil
		} else if types.IsNumber(f1) && types.IsNumber(f2) {
			n1 := types.AsNumber(f1)
			n2 := types.AsNumber(f2)
			if context.Range() != nil {
				return types.CreateRange(n1.IntValue(), n2.IntValue()), nil
			} else if n1.Decimal || n2.Decimal {
				if context.Subtract() != nil {
					return types.NewFloatNumber(n1.FloatValue() - n2.FloatValue()), nil
				} else if context.Divide() != nil {
					return types.NewFloatNumber(n1.FloatValue() / n2.FloatValue()), nil
				} else if context.Multiply() != nil {
					return types.NewFloatNumber(n1.FloatValue() * n2.FloatValue()), nil
				}
			} else {
				if context.Subtract() != nil {
					return types.NewIntNumber(int64(n1.IntValue()) - int64(n2.IntValue())), nil
				} else if context.Divide() != nil {
					return types.NewIntNumber(int64(n1.IntValue()) / int64(n2.IntValue())), nil
				} else if context.Multiply() != nil {
					return types.NewIntNumber(int64(n1.IntValue()) * int64(n2.IntValue())), nil
				}
			}
		} else {
			if f1IsLambda || f2IsLambda {
				if context.Range() != nil {
					return types.Null, burrito.WrappedErrorf("Cannot use a lambda in a range expression")
				}
				return types.Null, burrito.WrappedErrorf("Cannot use a lambda in an arithmetic operation")
			}
			return types.NaN(), nil
		}
	} else if len(context.AllField()) == 3 {
		// handle ternary operator
		if context.Question() != nil {
			f1, err := v.Visit(context.Field(0))
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			if types.AsBool(f1).BoolValue() {
				return v.Visit(context.Field(1))
			} else {
				return v.Visit(context.Field(2))
			}
		}
	}
	// if the field is another field in parentheses, return the value of that field
	// we need to also check if the first element is not a field, because if it is, it will be a function call
	if context.LeftParen() != nil && len(context.AllField()) == 1 && context.GetChild(0) != context.Field(0) {
		return v.Visit(context.Field(0))
	} else if context.LeftParen() != nil && len(context.AllField()) == 1 && context.GetChild(0) == context.Field(0) {
		lambda, err := v.Visit(context.Field(0))
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		params := make([]types.JsonType, 0)
		for _, param := range context.AllFunction_param() {
			p, err := v.Visit(param)
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			if param.(*parser.Function_paramContext).Spread() != nil {
				if a, ok := p.(*types.JsonArray); ok {
					params = append(params, a.Value...)
				} else {
					return types.Null, burrito.WrappedErrorf("Cannot spread %s", p.StringValue())
				}
			} else {
				params = append(params, p)
			}
		}
		if _, ok := lambda.(*types.JsonString); ok {
			lambdaContext := v.resolveLambdaTree(lambda.StringValue())
			fun, err := v.Visit(lambdaContext)
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			if f, ok := fun.(*types.JsonLambda); ok {
				i, err := f.Value(f, params)
				if err != nil {
					return types.Null, burrito.PassError(err)
				}
				return i, nil
			} else {
				return types.Null, burrito.WrappedErrorf("%s is not a function", lambda.StringValue())
			}
		} else if l, ok := lambda.(*types.JsonLambda); ok {
			i, err := l.Value(l, params)
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			return i, nil
		} else {
			var methodName *string = nil
			for _, child := range context.Field(0).GetChildren() {
				if b, ok := child.(parser.INameContext); ok {
					text := cachedGetText(b.(antlr.ParseTree))
					methodName = &text
					break
				}
			}
			if methodName == nil || !functions.HasFunction(*methodName) {
				if methodName != nil {
					find := functions.FindMisspelling(*methodName)
					if find != nil {
						return types.Null, burrito.WrappedErrorf("Function '%s' not found, did you mean '%s'?", *methodName, *find)
					}
				}
				return types.Null, burrito.WrappedErrorf("Function '%s' not found!", context.Field(0).GetText())
			}
			function, err := functions.CallFunction(*methodName, params, v.ResolveScope)
			if err != nil {
				return types.Null, burrito.WrapErrorf(err, "Error calling function '%s'", *methodName)
			}
			return function, nil
		}
	}
	// handle accessing a property of an object or calling an instance function
	if context.Name() != nil && len(context.AllField()) == 1 {
		text := cachedGetText(context.Name())
		object, err := v.Visit(context.Field(0))
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		if functions.HasInstanceFunction(reflect.TypeOf(object), text) {
			return types.NewLambda(
				func(this *types.JsonLambda, o []types.JsonType) (types.JsonType, error) {
					function, err := functions.CallInstanceFunction(text, object.(types.JsonType), o, v.ResolveScope)
					if err != nil {
						return types.Null, burrito.WrapErrorf(err, "Error calling function '%s' on %s", text, object.StringValue())
					}
					return function, nil
				},
				text,
				[]string{},
				[]string{},
			), nil
		} else {
			textStr := internedJsonString(text)
			index, err := object.Index(textStr)
			if err != nil {
				if context.Question() != nil || v.action == types.Predicate || isInLeftSideOfAssignment(context) {
					return types.NullWithParent(object, textStr), nil
				} else {
					return types.Null, burrito.WrapErrorf(err, "Cannot access %s", context.GetText())
				}
			}
			index.UpdateParent(object, textStr)
			return index, nil
		}
	}
	if context.Name() != nil {
		return v.Visit(context.Name())
	}
	// handle indexed access
	if context.Index() != nil && len(context.AllField()) == 1 {
		i, err := v.Visit(context.Index())
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		if _, ok := i.(*types.JsonLambda); ok {
			return types.Null, burrito.WrappedErrorf("Cannot use a lambda as an index")
		}
		object, err := v.Visit(context.Field(0))
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		if _, ok := object.(*types.JsonLambda); ok {
			return types.Null, burrito.WrappedErrorf("Cannot index a lambda")
		}
		index, err := object.Index(i)
		if err != nil {
			if context.Question() != nil || v.action == types.Predicate || isInLeftSideOfAssignment(context) {
				return types.NullWithParent(object, i), nil
			} else {
				return types.Null, burrito.WrapErrorf(err, "Cannot access %s", context.GetText())
			}
		}
		index.UpdateParent(object, i)
		return index, nil
	}
	if context.NUMBER() != nil {
		return cachedNodeNumber(context.NUMBER()), nil
	}
	if context.ESCAPED_STRING() != nil {
		return types.NewString(unescapeString(types.ToString(cachedGetText(context.ESCAPED_STRING())))), nil
	}
	// template string literal
	if context.Template_string() != nil {
		return v.Visit(context.Template_string())
	}
	// literal array notation
	if context.Array() != nil {
		return v.Visit(context.Array())
	}
	// literal object notation
	if context.Object() != nil {
		return v.Visit(context.Object())
	}
	// lambda literal in field position
	if context.Lambda() != nil {
		return v.Visit(context.Lambda())
	}
	// negation
	if context.Subtract() != nil && len(context.AllField()) == 1 {
		f, err := v.Visit(context.Field(0))
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		if _, ok := f.(*types.JsonLambda); ok {
			return types.Null, burrito.WrappedErrorf("Cannot negate a lambda")
		}
		return f.Negate(), nil
	}
	return types.Null, burrito.WrappedErrorf("Failed to resolve '%s'", context.GetText())
}

func (v *ExpressionVisitor) VisitName(context *parser.NameContext) (types.JsonType, error) {
	text := cachedGetText(context)
	newScope := v.ResolveScope(text)

	back := v.scope.Len() - 1
	for newScope == nil && back >= 0 {
		m := v.scope.At(back)
		if m.ContainsKey(text) {
			newScope = m.Get(text)
		}
		back--
	}
	return newScope, nil
}

func (v *ExpressionVisitor) VisitIndex(context *parser.IndexContext) (types.JsonType, error) {
	if context.NUMBER() != nil {
		return cachedNodeNumber(context.NUMBER()), nil
	}
	if context.ESCAPED_STRING() != nil {
		return types.NewString(unescapeString(types.ToString(cachedGetText(context.ESCAPED_STRING())))), nil
	}
	if context.Field() != nil {
		return v.Visit(context.Field())
	}
	return types.Null, burrito.WrappedErrorf("Invalid index: %s", context.GetText())
}

func (v *ExpressionVisitor) VisitArray(context *parser.ArrayContext) (types.JsonType, error) {
	var spreadFields []parser.ISpread_fieldContext
	if cached, ok := spreadFieldCache.Load(context); ok {
		spreadFields = cached.([]parser.ISpread_fieldContext)
	} else {
		spreadFields = context.AllSpread_field()
		spreadFieldCache.Store(context, spreadFields)
	}
	a := types.NewJsonArrayInline()
	for _, f := range spreadFields {
		sf := f.(*parser.Spread_fieldContext)
		r, err := v.Visit(sf.Field())
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		if sf.Spread() != nil {
			if _, ok := r.(*types.JsonArray); !ok {
				return types.Null, burrito.WrappedErrorf("Cannot spread %s into an array", r.StringValue())
			}
			a.Value = append(a.Value, r.(*types.JsonArray).Value...)
		} else {
			a.Value = append(a.Value, r)
		}
	}
	return a, nil
}

func (v *ExpressionVisitor) VisitObject(context *parser.ObjectContext) (types.JsonType, error) {
	result := types.NewJsonObject()
	for _, f := range context.AllObject_field() {
		sf := f.(*parser.Object_fieldContext)
		if sf.Spread() != nil {
			obj, err := v.Visit(sf.Field())
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			if u, ok := obj.(*types.JsonObject); ok {
				u.RangeEntries(func(key string, value types.JsonType) bool {
					result.Put(key, value)
					return false
				})
			}
		} else {
			var name string
			if sf.ESCAPED_STRING() != nil {
				name = unescapeString(types.ToString(cachedGetText(sf.ESCAPED_STRING())))
			} else {
				name = cachedGetText(sf.Name())
			}
			field, err := v.Visit(sf.Field())
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			result.Put(name, field)
		}
	}
	return result, nil
}

func (v *ExpressionVisitor) VisitObject_field(context *parser.Object_fieldContext) (types.JsonType, error) {
	name := ""
	if context.ESCAPED_STRING() != nil {
		name = unescapeString(types.ToString(cachedGetText(context.ESCAPED_STRING())))
	} else if context.Spread() != nil {
		return v.Visit(context.Field())
	} else {
		name = cachedGetText(context.Name())
	}
	field, err := v.Visit(context.Field())
	if err != nil {
		return types.Null, burrito.PassError(err)
	}
	n := types.NewJsonObject()
	n.Put(name, field)
	return n, nil
}

func (v *ExpressionVisitor) VisitLambda(ctx *parser.LambdaContext) (types.JsonType, error) {
	var info lambdaCtxInfo
	if cached, ok := lambdaInfoCache.Load(ctx); ok {
		info = cached.(lambdaCtxInfo)
	} else {
		lv := LambdaVisitor{}
		lv.Visit(ctx)
		directNames := ctx.AllName()
		params := make([]string, len(directNames))
		for i, n := range directNames {
			params[i] = n.GetText()
		}
		info = lambdaCtxInfo{
			text:   ctx.GetStart().GetInputStream().GetText(ctx.GetStart().GetStart(), ctx.GetStop().GetStop()),
			vars:   lv.usedVariables,
			args:   lv.arguments,
			params: params,
		}
		lambdaInfoCache.Store(ctx, info)
	}
	hasBlock := ctx.Statements() != nil
	params := info.params
	return types.NewLambda(
		func(this *types.JsonLambda, o []types.JsonType) (types.JsonType, error) {
			if len(params) > len(o) {
				return types.Null, burrito.WrappedErrorf("Lambda expects %d arguments, but got %d", len(params), len(o))
			}
			for i, argName := range params {
				v.pushScopePair(argName, o[i])
			}
			var result types.JsonType
			var err error
			if hasBlock {
				result, err = v.Visit(ctx.Statements())
				if err == nil {
					if signal, ok := result.(*types.JsonSignal); ok {
						result = signal.Value
					} else {
						result = types.Null
					}
				}
			} else {
				result, err = v.Visit(ctx.Field())
			}
			for range params {
				v.popScopeObject()
			}
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			return result, nil
		},
		info.text,
		info.vars,
		info.args,
	), nil
}

func (v *ExpressionVisitor) VisitFunction_param(ctx *parser.Function_paramContext) (types.JsonType, error) {
	if ctx.Field() != nil {
		return v.Visit(ctx.Field())
	}
	if ctx.Lambda() != nil {
		return v.Visit(ctx.Lambda())
	}
	return types.Null, burrito.WrappedErrorf("Invalid function parameter: %s", ctx.GetText())
}

func (v *ExpressionVisitor) VisitTemplate_string(ctx *parser.Template_stringContext) (types.JsonType, error) {
	var sb strings.Builder
	for _, part := range ctx.AllTemplate_string_part() {
		partCtx := part.(*parser.Template_string_partContext)
		if partCtx.TEMPLATE_TEXT() != nil {
			sb.WriteString(unescapeTemplateText(cachedGetText(partCtx.TEMPLATE_TEXT())))
		} else {
			val, err := v.Visit(partCtx.Field())
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			sb.WriteString(val.StringValue())
		}
	}
	return types.NewString(sb.String()), nil
}

func (v *ExpressionVisitor) VisitTemplate_string_part(ctx *parser.Template_string_partContext) (types.JsonType, error) {
	if ctx.Field() != nil {
		return v.Visit(ctx.Field())
	}
	if ctx.TEMPLATE_TEXT() != nil {
		return types.NewString(unescapeTemplateText(cachedGetText(ctx.TEMPLATE_TEXT()))), nil
	}
	return types.NewString(""), nil
}

// unescapeTemplateText handles backslash escapes in template string literal segments.
func unescapeTemplateText(text string) string {
	return UnescapeString(text)
}

// unescapeString removes quotes and unescapes a string.
func unescapeString(str string) string {
	runes := []rune(str)
	if len(runes) < 3 {
		return ""
	}
	str = string(runes[1 : len(runes)-1])
	return UnescapeString(str)
}


