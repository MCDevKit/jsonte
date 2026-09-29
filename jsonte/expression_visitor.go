package jsonte

import (
	"math"
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
	// body is ctx.Statements() when hasBlock, otherwise ctx.Field().
	body     antlr.ParseTree
	hasBlock bool
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

// fieldInfoCache shares immutable *fieldCtxInfo values across visitors (most
// evaluations use a fresh visitor) for cached parse trees. It is keyed by
// *parser.FieldContext and cleared together with the expression cache.
var fieldInfoCache sync.Map

// fieldCtxInfo holds precomputed child/token lookups for a FieldContext.
// ANTLR's generated accessors (GetToken, AllField, Name, ...) scan and
// type-assert the children slice on every call; since parse trees are cached
// and reused across evaluations, this metadata is computed once per node.
type fieldCtxInfo struct {
	fields         []parser.IFieldContext
	params         []parser.IFunction_paramContext
	name           parser.INameContext
	index          parser.IIndexContext
	number         antlr.TerminalNode
	escapedString  antlr.TerminalNode
	templateString parser.ITemplate_stringContext
	array          parser.IArrayContext
	object         parser.IObjectContext
	lambda         parser.ILambdaContext

	hasNull, hasTrue, hasFalse        bool
	hasNot, hasAddAssign, hasLiteral  bool
	hasAnd, hasOr, hasNullCoalescing  bool
	hasAdd, hasEqual, hasNotEqual     bool
	hasLess, hasGreater               bool
	hasLessOrEqual, hasGreaterOrEqual bool
	hasRange, hasSubtract             bool
	hasDivide, hasMultiply            bool
	hasQuestion, hasLeftParen         bool
	hasIn                             bool
	hasIncrement, hasDecrement        bool
	hasSubtractAssign                 bool
	hasMultiplyAssign                 bool
	hasDivideAssign, hasModuloAssign  bool

	// nameText/nameStr are the property name and its interned JsonString when
	// name is non-nil; numberVal is the parsed literal when number is non-nil.
	nameText  string
	nameStr   *types.JsonString
	numberVal *types.JsonNumber
	// escapedVal is the unescaped ESCAPED_STRING literal; indexInfo is the
	// precomputed child of index.
	escapedVal string
	indexInfo  indexCtxInfo
}

// exprCtxInfo holds precomputed lookups for an ExpressionContext.
type exprCtxInfo struct {
	action    types.JsonAction
	name      string
	indexName string
	field     parser.IFieldContext
}

func computeExprCtxInfo(ctx *parser.ExpressionContext) *exprCtxInfo {
	info := &exprCtxInfo{
		action:    types.Value,
		name:      DefaultName,
		indexName: DefaultIndexName,
		field:     ctx.Field(),
	}
	if ctx.Iteration() != nil {
		info.action = types.Iteration
	}
	if ctx.Literal() != nil {
		info.action = types.Literal
	}
	if ctx.Question() != nil {
		info.action = types.Predicate
	}
	if ctx.As() != nil {
		if len(ctx.AllName()) > 0 {
			info.name = ctx.Name(0).GetText()
		}
		if len(ctx.AllName()) > 1 {
			info.indexName = ctx.Name(1).GetText()
		}
	}
	return info
}

func computeFieldCtxInfo(ctx *parser.FieldContext) *fieldCtxInfo {
	// Single pass over the children, mirroring the generated accessors: the
	// singular accessors return the first match, All* return every match in order.
	children := ctx.GetChildren()
	info := &fieldCtxInfo{}
	for _, child := range children {
		switch c := child.(type) {
		case *parser.FieldContext:
			if info.fields == nil {
				info.fields = make([]parser.IFieldContext, 0, len(children))
			}
			info.fields = append(info.fields, c)
		case *parser.Function_paramContext:
			if info.params == nil {
				info.params = make([]parser.IFunction_paramContext, 0, len(children))
			}
			info.params = append(info.params, c)
		case *parser.NameContext:
			if info.name == nil {
				info.name = c
			}
		case *parser.IndexContext:
			if info.index == nil {
				info.index = c
			}
		case *parser.Template_stringContext:
			if info.templateString == nil {
				info.templateString = c
			}
		case *parser.ArrayContext:
			if info.array == nil {
				info.array = c
			}
		case *parser.ObjectContext:
			if info.object == nil {
				info.object = c
			}
		case *parser.LambdaContext:
			if info.lambda == nil {
				info.lambda = c
			}
		case antlr.TerminalNode:
			switch c.GetSymbol().GetTokenType() {
			case parser.JsonTemplateNUMBER:
				if info.number == nil {
					info.number = c
				}
			case parser.JsonTemplateESCAPED_STRING:
				if info.escapedString == nil {
					info.escapedString = c
				}
			case parser.JsonTemplateNull:
				info.hasNull = true
			case parser.JsonTemplateTrue:
				info.hasTrue = true
			case parser.JsonTemplateFalse:
				info.hasFalse = true
			case parser.JsonTemplateNot:
				info.hasNot = true
			case parser.JsonTemplateAddAssign:
				info.hasAddAssign = true
			case parser.JsonTemplateLiteral:
				info.hasLiteral = true
			case parser.JsonTemplateAnd:
				info.hasAnd = true
			case parser.JsonTemplateOr:
				info.hasOr = true
			case parser.JsonTemplateNullCoalescing:
				info.hasNullCoalescing = true
			case parser.JsonTemplateAdd:
				info.hasAdd = true
			case parser.JsonTemplateEqual:
				info.hasEqual = true
			case parser.JsonTemplateNotEqual:
				info.hasNotEqual = true
			case parser.JsonTemplateLess:
				info.hasLess = true
			case parser.JsonTemplateGreater:
				info.hasGreater = true
			case parser.JsonTemplateLessOrEqual:
				info.hasLessOrEqual = true
			case parser.JsonTemplateGreaterOrEqual:
				info.hasGreaterOrEqual = true
			case parser.JsonTemplateRange:
				info.hasRange = true
			case parser.JsonTemplateSubtract:
				info.hasSubtract = true
			case parser.JsonTemplateDivide:
				info.hasDivide = true
			case parser.JsonTemplateMultiply:
				info.hasMultiply = true
			case parser.JsonTemplateQuestion:
				info.hasQuestion = true
			case parser.JsonTemplateLeftParen:
				info.hasLeftParen = true
			case parser.JsonTemplateIn:
				info.hasIn = true
			case parser.JsonTemplateIncrement:
				info.hasIncrement = true
			case parser.JsonTemplateDecrement:
				info.hasDecrement = true
			case parser.JsonTemplateSubtractAssign:
				info.hasSubtractAssign = true
			case parser.JsonTemplateMultiplyAssign:
				info.hasMultiplyAssign = true
			case parser.JsonTemplateDivideAssign:
				info.hasDivideAssign = true
			case parser.JsonTemplateModuloAssign:
				info.hasModuloAssign = true
			}
		}
	}
	if info.fields == nil {
		info.fields = []parser.IFieldContext{}
	}
	if info.params == nil {
		info.params = []parser.IFunction_paramContext{}
	}
	if info.name != nil {
		info.nameText = info.name.GetText()
		info.nameStr = internedJsonString(info.nameText)
	}
	if info.number != nil {
		info.numberVal = types.AsNumber(info.number.GetText())
	}
	if info.escapedString != nil {
		info.escapedVal = unescapeString(info.escapedString.GetText())
	}
	if info.index != nil {
		info.indexInfo = computeIndexCtxInfo(info.index.(*parser.IndexContext))
	}
	return info
}

// indexCtxInfo holds the precomputed child of an IndexContext. Exactly one of
// number, hasString or field is used, in the same priority as VisitIndex.
type indexCtxInfo struct {
	number    *types.JsonNumber
	str       string
	hasString bool
	field     parser.IFieldContext
}

func computeIndexCtxInfo(ctx *parser.IndexContext) indexCtxInfo {
	var info indexCtxInfo
	if n := ctx.NUMBER(); n != nil {
		info.number = types.AsNumber(n.GetText())
	} else if s := ctx.ESCAPED_STRING(); s != nil {
		info.str = unescapeString(s.GetText())
		info.hasString = true
	} else {
		info.field = ctx.Field()
	}
	return info
}

// visitIndexInfo evaluates a precomputed index; equivalent to VisitIndex.
func (v *ExpressionVisitor) visitIndexInfo(ctx parser.IIndexContext, info *indexCtxInfo) (types.JsonType, error) {
	if info.number != nil {
		return info.number, nil
	}
	if info.hasString {
		return types.NewString(info.str), nil
	}
	if info.field != nil {
		return v.Visit(info.field)
	}
	return types.Null, burrito.WrappedErrorf("Invalid index: %s", ctx.GetText())
}

// cachedGetText returns GetText() for any stable parse-tree node, caching the result.
func cachedGetText(node antlr.ParseTree) string {
	if v, ok := nodeTextCache.Load(node); ok {
		return v.(string)
	}
	text := node.GetText()
	nodeTextCache.Store(node, text)
	return text
}

func internedJsonString(s string) *types.JsonString {
	if v, ok := jsonStringInternCache.Load(s); ok {
		return v.(*types.JsonString)
	}
	js := types.NewString(s)
	actual, loaded := jsonStringInternCache.LoadOrStore(s, js)
	if loaded {
		return actual.(*types.JsonString)
	}
	return js
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

var localVarPool = sync.Pool{
	New: func() interface{} { return types.NewJsonObjectWithCapacity(4) },
}

func getLocalVarObject() *types.JsonObject {
	return localVarPool.Get().(*types.JsonObject)
}

func putLocalVarObject(obj *types.JsonObject) {
	obj.Reset()
	localVarPool.Put(obj)
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
	// localVarTarget, when non-nil, is the object where new variable assignments land.
	// It is set to a fresh object on each lambda call and restored on return, so that
	// variables assigned inside a lambda body are local to that call and don't leak into
	// variableScope (which is shared across all lambda calls on the same visitor).
	localVarTarget *types.JsonObject
	// globalScope, when non-nil, enables the $scope pseudo-variable in .jsonte scripts.
	// $scope.key = value writes directly into the global scope object.
	globalScope *types.JsonObject
	// fieldInfos/exprInfos memoize per-node parse-tree metadata. They are plain
	// maps (the visitor is single-threaded) keyed by node pointer, which is much
	// cheaper than an interface-keyed sync.Map. Pooled visitors keep them across
	// evaluations on purpose; parse trees are cached and their nodes are stable.
	fieldInfos map[*parser.FieldContext]*fieldCtxInfo
	exprInfos  map[*parser.ExpressionContext]*exprCtxInfo
	nameTexts  map[*parser.NameContext]string
	// resolveFn caches the v.ResolveScope method value (see resolveScopeFn).
	resolveFn func(string) types.JsonType
}

func (v *ExpressionVisitor) nameText(ctx *parser.NameContext) string {
	if text, ok := v.nameTexts[ctx]; ok {
		return text
	}
	text := cachedGetText(ctx)
	if v.nameTexts == nil {
		v.nameTexts = make(map[*parser.NameContext]string, 64)
	}
	v.nameTexts[ctx] = text
	return text
}

func (v *ExpressionVisitor) fieldInfo(ctx *parser.FieldContext) *fieldCtxInfo {
	if info, ok := v.fieldInfos[ctx]; ok {
		return info
	}
	var info *fieldCtxInfo
	if cached, ok := fieldInfoCache.Load(ctx); ok {
		info = cached.(*fieldCtxInfo)
	} else {
		info = computeFieldCtxInfo(ctx)
		if actual, loaded := fieldInfoCache.LoadOrStore(ctx, info); loaded {
			info = actual.(*fieldCtxInfo)
		}
	}
	if v.fieldInfos == nil {
		v.fieldInfos = make(map[*parser.FieldContext]*fieldCtxInfo, 64)
	}
	v.fieldInfos[ctx] = info
	return info
}

func (v *ExpressionVisitor) exprInfo(ctx *parser.ExpressionContext) *exprCtxInfo {
	if info, ok := v.exprInfos[ctx]; ok {
		return info
	}
	info := computeExprCtxInfo(ctx)
	if v.exprInfos == nil {
		v.exprInfos = make(map[*parser.ExpressionContext]*exprCtxInfo, 16)
	}
	v.exprInfos[ctx] = info
	return info
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

func (v *ExpressionVisitor) isInLeftSideOfAssignment(context antlr.Tree) bool {
	return treeMatches(context, func(ctx, parent antlr.Tree) bool {
		if f, ok := parent.(*parser.FieldContext); ok {
			fi := v.fieldInfo(f)
			if len(fi.fields) == 2 && antlr.Tree(fi.fields[0]) == ctx && (fi.hasLiteral || fi.hasAddAssign) {
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
	ctx := parseTree(src, lambdaRule, antlr.NewConsoleErrorListener()).(parser.ILambdaContext)
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
	// Fast path for JsonNumber and JsonString to avoid boxing overhead when we already have a JsonType.
	switch val := value.(type) {
	case types.JsonType:
		obj.Put(key, val)
	default:
		obj.Put(key, types.Box(value))
	}
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

// assignmentTarget returns the object where new variable assignments should land.
// Inside a lambda body it is the lambda-local scope; at script level it is variableScope.
func (v *ExpressionVisitor) assignmentTarget() *types.JsonObject {
	if v.localVarTarget != nil {
		return v.localVarTarget
	}
	return v.initVarScope()
}

// ResolveScope resolves a value from the scope by name
func (v *ExpressionVisitor) ResolveScope(name string) types.JsonType {
	return v.resolveScopeStr(name, nil)
}

// resolveScopeFn returns ResolveScope as a func value, bound once per visitor
// so passing it to function calls does not allocate a method value each time.
func (v *ExpressionVisitor) resolveScopeFn() func(string) types.JsonType {
	if v.resolveFn == nil {
		v.resolveFn = v.ResolveScope
	}
	return v.resolveFn
}

// isSpecialScopeName reports whether name is handled specially by resolveScopeStr
// ($scope and this) instead of being looked up in the scope layers.
func (v *ExpressionVisitor) isSpecialScopeName(name string) bool {
	return name == "this" || (name == "$scope" && v.globalScope != nil)
}

// resolveScopeStr resolves a value from the scope by name. nameStr, when
// non-nil, must be the interned JsonString for name; it is looked up lazily
// otherwise.
func (v *ExpressionVisitor) resolveScopeStr(name string, nameStr *types.JsonString) types.JsonType {
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
	if nameStr == nil {
		nameStr = internedJsonString(name)
	}
	if get, ok := v.lookupScope(name, nameStr); ok {
		return get
	}
	return types.NullWithParent(v.assignmentTarget(), nameStr)
}

// lookupScope searches the scope layers (lambda locals, scope stack, variable
// scope, extra scope) for name in ResolveScope order. On a hit the value's
// parent is updated exactly as ResolveScope does; on a miss nothing is
// allocated and (nil, false) is returned. It does not handle $scope or this.
func (v *ExpressionVisitor) lookupScope(name string, nameStr *types.JsonString) (types.JsonType, bool) {
	if v.localVarTarget != nil {
		if get, ok := v.localVarTarget.TryGet(name); ok {
			get.UpdateParent(v.localVarTarget, nameStr)
			return get, true
		}
	}
	for i := v.scope.Len() - 1; i >= 0; i-- {
		m := v.scope.At(i)
		if get, ok := m.TryGet(name); ok {
			get.UpdateParent(m, nameStr)
			return get, true
		}
	}
	if v.variableScope != nil {
		if get, ok := v.variableScope.TryGet(name); ok {
			get.UpdateParent(v.variableScope, nameStr)
			return get, true
		}
	}
	if v.extraScope != nil {
		for i := v.extraScope.Len() - 1; i >= 0; i-- {
			m := v.extraScope.At(i)
			if get, ok := m.TryGet(name); ok {
				get.UpdateParent(m, nameStr)
				return get, true
			}
		}
	}
	return nil, false
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
			val, err := v.Visit(ctx.Statement(0))
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
				val, err := v.Visit(ctx.Statement(i))
				if err != nil {
					return types.Null, burrito.PassError(err)
				}
				return val, nil
			}
		}
		if len(ctx.AllElse()) == len(ctx.AllIf()) {
			val, err := v.Visit(ctx.Statement(len(ctx.AllIf())))
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			return val, nil
		}
	} else if ctx.Do() != nil {
		for {
			val, err := v.Visit(ctx.Statement(0))
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
			val, err := v.Visit(ctx.Statement(0))
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
	} else if ctx.Statements() != nil {
		return v.Visit(ctx.Statements())
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
	info := v.exprInfo(ctx)
	v.action = info.action
	v.name = info.name
	v.indexName = info.indexName
	result, err := v.Visit(info.field)
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

// storeIntoParent writes value into the container that left was resolved from
// (assignment machinery shared by +=, -=, *=, /=, %=, ++ and --) and returns the stored value.
// parent and parentIndex must be read from left right after resolving it: see assignTarget.
func storeIntoParent(parent, parentIndex, value types.JsonType) (types.JsonType, error) {
	if parent == nil {
		return types.Null, burrito.WrappedErrorf("Cannot assign a value to this expression")
	}
	if parentIndex == nil {
		utils.BadDeveloperError("Invalid parent index type")
	}
	if b, ok := parent.(*types.JsonObject); ok {
		if i, ok1 := parentIndex.(*types.JsonString); ok1 {
			b.Put(i.StringValue(), value)
			return value, nil
		} else if i, ok1 := parentIndex.(*types.JsonPath); ok1 {
			return i.Set(parent, value)
		} else {
			return types.Null, burrito.WrappedErrorf("Invalid index type. Expected string or Json Path, but got %s", reflect.TypeOf(parentIndex).String())
		}
	} else if b, ok := parent.(*types.JsonArray); ok {
		if i, ok1 := parentIndex.(*types.JsonNumber); ok1 {
			index := int(i.IntValue())
			if index < 0 {
				index = len(b.Value) + index
			}
			if index >= 0 && index < len(b.Value) {
				b.Value[index] = value
				return value, nil
			} else {
				return types.Null, burrito.WrappedErrorf("Index out of bounds: %d", index)
			}
		} else if i, ok1 := parentIndex.(*types.JsonPath); ok1 {
			return i.Set(parent, value)
		} else {
			return types.Null, burrito.WrappedErrorf("Invalid index type. Expected number or Json Path, but got %s", reflect.TypeOf(parentIndex).String())
		}
	} else {
		return types.Null, burrito.WrappedErrorf("Cannot assign a value to this expression")
	}
}

// assignTarget resolves the left side of an assignment and returns where to store the new value.
// The parent must be captured before anything else is evaluated: small numbers are shared instances
// (see smallIntCache in types/number.go), so resolving another variable holding the same number
// overwrites the parent of the left side.
// TODO: This only fixes the evaluation order. Resolving a variable still mutates the shared cached
// instance, which is a data race when templates are processed by multiple workers. Revisit by not
// setting parents on cached numbers or by caching without shared mutable state.
func (v *ExpressionVisitor) assignTarget(field antlr.ParseTree) (left, parent, parentIndex types.JsonType, err error) {
	left, err = v.Visit(field)
	if err != nil {
		return nil, nil, nil, err
	}
	return left, left.Parent(), left.ParentIndex(), nil
}

// numericCompound computes the value for -=, *=, /= and %= mirroring the binary
// operator semantics: numeric operands produce a number, anything else NaN.
// Integer division or modulo by zero yields NaN instead of panicking.
func numericCompound(left, right types.JsonType, info *fieldCtxInfo) types.JsonType {
	if !types.IsNumber(left) || !types.IsNumber(right) {
		return types.NaN()
	}
	n1 := types.AsNumber(left)
	n2 := types.AsNumber(right)
	if info.hasModuloAssign {
		if n2.FloatValue() == 0 {
			return types.NaN()
		}
		return &types.JsonNumber{Value: math.Mod(n1.FloatValue(), n2.FloatValue()), Decimal: n1.Decimal || n2.Decimal}
	}
	if n1.Decimal || n2.Decimal {
		switch {
		case info.hasSubtractAssign:
			return types.NewFloatNumber(n1.FloatValue() - n2.FloatValue())
		case info.hasMultiplyAssign:
			return types.NewFloatNumber(n1.FloatValue() * n2.FloatValue())
		default:
			return types.NewFloatNumber(n1.FloatValue() / n2.FloatValue())
		}
	}
	switch {
	case info.hasSubtractAssign:
		return types.NewIntNumber(int64(n1.IntValue()) - int64(n2.IntValue()))
	case info.hasMultiplyAssign:
		return types.NewIntNumber(int64(n1.IntValue()) * int64(n2.IntValue()))
	default:
		if n2.IntValue() == 0 {
			return types.NaN()
		}
		return types.NewIntNumber(int64(n1.IntValue()) / int64(n2.IntValue()))
	}
}

// memberAccess evaluates `object.name` for a member-access field (info.name != nil
// and a single sub-field that evaluated to object). Instance functions are
// returned as a bound lambda; anything else is indexed by the property name.
func (v *ExpressionVisitor) memberAccess(context *parser.FieldContext, info *fieldCtxInfo, object types.JsonType) (types.JsonType, error) {
	text := info.nameText
	if functions.HasInstanceFunction(reflect.TypeOf(object), text) {
		return types.NewLambda(
			func(this *types.JsonLambda, o []types.JsonType) (types.JsonType, error) {
				function, err := functions.CallInstanceFunction(text, object.(types.JsonType), o, v.resolveScopeFn())
				if err != nil {
					return types.Null, burrito.WrapErrorf(err, "Error calling function '%s' on %s", text, object.StringValue())
				}
				return function, nil
			},
			text,
			[]string{},
			[]string{},
		), nil
	}
	textStr := info.nameStr
	index, err := object.Index(textStr)
	if err != nil {
		if info.hasQuestion || v.action == types.Predicate || v.isInLeftSideOfAssignment(context) {
			return types.NullWithParent(object, textStr), nil
		} else {
			return types.Null, burrito.WrapErrorf(err, "Cannot access %s", context.GetText())
		}
	}
	index.UpdateParent(object, textStr)
	return index, nil
}

// visitCall evaluates a function call field `callee(params...)`.
//
// Common callee shapes are resolved without allocating intermediates:
//   - a bare name missing from every scope goes straight to the global function
//     lookup (instead of materializing a Null placeholder first);
//   - `object.method(...)` where method is an instance function is called
//     directly instead of through a freshly allocated bound lambda.
//
// Evaluation order (callee/object first, then params) and error messages are
// unchanged.
func (v *ExpressionVisitor) visitCall(info *fieldCtxInfo) (types.JsonType, error) {
	calleeCtx := info.fields[0].(*parser.FieldContext)
	fi := v.fieldInfo(calleeCtx)

	// callee is the evaluated callee value; nil means an unresolved bare name.
	var callee types.JsonType
	// instance is non-nil when calling instance function fi.nameText on it directly.
	var instance types.JsonType
	if fi.name != nil && len(fi.fields) == 0 && !v.isSpecialScopeName(fi.nameText) {
		if val, ok := v.lookupScope(fi.nameText, fi.nameStr); ok {
			callee = val
		}
	} else if fi.name != nil && len(fi.fields) == 1 {
		object, err := v.Visit(fi.fields[0])
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		if functions.HasInstanceFunction(reflect.TypeOf(object), fi.nameText) {
			instance = object
		} else {
			callee, err = v.memberAccess(calleeCtx, fi, object)
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
		}
	} else {
		var err error
		callee, err = v.Visit(calleeCtx)
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
	}

	params := make([]types.JsonType, 0, len(info.params))
	for _, param := range info.params {
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

	if instance != nil {
		function, err := functions.CallInstanceFunction(fi.nameText, instance, params, v.resolveScopeFn())
		if err != nil {
			return types.Null, burrito.PassError(burrito.WrapErrorf(err, "Error calling function '%s' on %s", fi.nameText, instance.StringValue()))
		}
		return function, nil
	}
	if _, ok := callee.(*types.JsonString); ok {
		lambdaContext := v.resolveLambdaTree(callee.StringValue())
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
			return types.Null, burrito.WrappedErrorf("%s is not a function", callee.StringValue())
		}
	} else if l, ok := callee.(*types.JsonLambda); ok {
		i, err := l.Value(l, params)
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		return i, nil
	}
	// Not a lambda: treat the callee's name (if any) as a global function.
	if fi.name == nil || !functions.HasFunction(fi.nameText) {
		if fi.name != nil {
			find := functions.FindMisspelling(fi.nameText)
			if find != nil {
				return types.Null, burrito.WrappedErrorf("Function '%s' not found, did you mean '%s'?", fi.nameText, *find)
			}
		}
		return types.Null, burrito.WrappedErrorf("Function '%s' not found!", calleeCtx.GetText())
	}
	function, err := functions.CallFunction(fi.nameText, params, v.resolveScopeFn())
	if err != nil {
		return types.Null, burrito.WrapErrorf(err, "Error calling function '%s'", fi.nameText)
	}
	return function, nil
}

func (v *ExpressionVisitor) VisitField(context *parser.FieldContext) (types.JsonType, error) {
	info := v.fieldInfo(context)
	if info.hasNull {
		return types.Null, nil
	}
	if info.hasTrue {
		return types.True(), nil
	}
	if info.hasFalse {
		return types.False(), nil
	}
	if info.hasNot {
		visit, err := v.Visit(info.fields[0])
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		return types.AsBool(!visit.BoolValue()), nil
	}
	if info.hasIncrement || info.hasDecrement {
		left, parent, parentIndex, err := v.assignTarget(info.fields[0])
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		if !types.IsNumber(left) {
			return types.Null, burrito.WrappedErrorf("Cannot increment or decrement %s", left.StringValue())
		}
		n := types.AsNumber(left)
		delta := int64(1)
		if info.hasDecrement {
			delta = -1
		}
		var value types.JsonType
		if n.Decimal {
			value = types.NewFloatNumber(n.FloatValue() + float64(delta))
		} else {
			value = types.NewIntNumber(int64(n.IntValue()) + delta)
		}
		return storeIntoParent(parent, parentIndex, value)
	}
	if info.hasAddAssign || info.hasSubtractAssign || info.hasMultiplyAssign || info.hasDivideAssign || info.hasModuloAssign {
		left, parent, parentIndex, err := v.assignTarget(info.fields[0])
		if err != nil {
			return types.Null, burrito.WrapErrorf(err, "Error resolving left side of the statement")
		}
		right, err := v.Visit(info.fields[1])
		if err != nil {
			return types.Null, burrito.WrapErrorf(err, "Error resolving right side of the statement")
		}
		var value types.JsonType
		if info.hasAddAssign {
			value = types.Box(left.Add(right).Unbox())
		} else {
			value = numericCompound(left, right, info)
		}
		return storeIntoParent(parent, parentIndex, value)
	}
	if info.hasLiteral {
		_, parent, parentIndex, err := v.assignTarget(info.fields[0])
		if err != nil {
			return types.Null, burrito.WrapErrorf(err, "Error resolving left side of the statement")
		}
		right, err := v.Visit(info.fields[1])
		if err != nil {
			return types.Null, burrito.WrapErrorf(err, "Error resolving right side of the statement")
		}
		// Preserve lambdas as-is; for other types, re-box to strip parent references.
		if _, ok := right.(*types.JsonLambda); !ok {
			right = types.Box(right.Unbox())
		}
		if parent == nil {
			return types.Null, burrito.WrappedErrorf("Cannot assign a value to this expression")
		}
		if parentIndex == nil {
			utils.BadDeveloperError("Invalid parent index type")
		}
		if b, ok := parent.(*types.JsonObject); ok {
			if i, ok1 := parentIndex.(*types.JsonString); ok1 {
				b.Put(i.StringValue(), right)
				return right, nil
			} else if i, ok1 := parentIndex.(*types.JsonPath); ok1 {
				return i.Set(parent, right)
			} else {
				return types.Null, burrito.WrappedErrorf("Invalid index type. Expected string or Json Path, but got %s", reflect.TypeOf(parentIndex).String())
			}
		} else if b, ok := parent.(*types.JsonArray); ok {
			if i, ok1 := parentIndex.(*types.JsonNumber); ok1 {
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
			} else if i, ok1 := parentIndex.(*types.JsonPath); ok1 {
				return i.Set(parent, right)
			} else {
				return types.Null, burrito.WrappedErrorf("Invalid index type. Expected number or Json Path, but got %s", reflect.TypeOf(parentIndex).String())
			}
		} else if b, ok := parent.(*types.JsonPath); ok {
			if i, ok1 := parentIndex.(*types.JsonNumber); ok1 {
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
	if len(info.fields) == 2 {
		f1, err := v.Visit(info.fields[0])
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		// Move AND and OR here to make those operators short-circuiting
		if info.hasAnd {
			if f1.BoolValue() {
				f2, err := v.Visit(info.fields[1])
				if err != nil {
					return types.Null, burrito.PassError(err)
				}
				return types.AsBool(f2.BoolValue()), nil
			} else {
				return types.False(), nil
			}
		} else if info.hasOr {
			if !f1.BoolValue() {
				f2, err := v.Visit(info.fields[1])
				if err != nil {
					return types.Null, burrito.PassError(err)
				}
				return types.AsBool(f2.BoolValue()), nil
			} else {
				return types.True(), nil
			}
		} else if info.hasQuestion {
			// Colonless ternary: `a ? b` evaluates to b when a is truthy, null otherwise.
			if f1.BoolValue() {
				return v.Visit(info.fields[1])
			}
			return types.Null, nil
		}
		f2, err := v.Visit(info.fields[1])
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		_, f1IsLambda := f1.(*types.JsonLambda)
		_, f2IsLambda := f2.(*types.JsonLambda)
		if info.hasNullCoalescing {
			if types.IsNull(f1) {
				return f2, nil
			} else {
				return f1, nil
			}
		} else if info.hasAdd {
			if f1IsLambda || f2IsLambda {
				return types.Null, burrito.WrappedErrorf("Cannot use a lambda in an addition or concatenation")
			}
			return f1.Add(f2), nil
		} else if info.hasEqual {
			return types.AsBool(f1.Equals(f2)), nil
		} else if info.hasNotEqual {
			return types.AsBool(!f1.Equals(f2)), nil
		} else if info.hasIn {
			switch c := f2.(type) {
			case *types.JsonObject:
				return types.AsBool(c.ContainsKey(f1.StringValue())), nil
			case *types.JsonArray:
				for _, el := range c.Value {
					if el.Equals(f1) {
						return types.True(), nil
					}
				}
				return types.False(), nil
			case *types.JsonString:
				return types.AsBool(strings.Contains(c.StringValue(), f1.StringValue())), nil
			default:
				return types.Null, burrito.WrappedErrorf("The 'in' operator requires an object, array or string on the right side, but got %s", f2.StringValue())
			}
		} else if info.hasLess {
			if f1IsLambda || f2IsLambda {
				return types.Null, burrito.WrappedErrorf("Cannot compare a lambda")
			}
			than, err := f1.LessThan(f2)
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			return types.AsBool(than), nil
		} else if info.hasGreater {
			if f1IsLambda || f2IsLambda {
				return types.Null, burrito.WrappedErrorf("Cannot compare a lambda")
			}
			than, err := f1.LessThan(f2)
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			return types.AsBool(!than && !f1.Equals(f2)), nil
		} else if info.hasLessOrEqual {
			if f1IsLambda || f2IsLambda {
				return types.Null, burrito.WrappedErrorf("Cannot compare a lambda")
			}
			than, err := f1.LessThan(f2)
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			return types.AsBool(than || f1.Equals(f2)), nil
		} else if info.hasGreaterOrEqual {
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
			if info.hasRange {
				return types.CreateRange(n1.IntValue(), n2.IntValue()), nil
			} else if n1.Decimal || n2.Decimal {
				if info.hasSubtract {
					return types.NewFloatNumber(n1.FloatValue() - n2.FloatValue()), nil
				} else if info.hasDivide {
					return types.NewFloatNumber(n1.FloatValue() / n2.FloatValue()), nil
				} else if info.hasMultiply {
					return types.NewFloatNumber(n1.FloatValue() * n2.FloatValue()), nil
				}
			} else {
				if info.hasSubtract {
					return types.NewIntNumber(int64(n1.IntValue()) - int64(n2.IntValue())), nil
				} else if info.hasDivide {
					return types.NewIntNumber(int64(n1.IntValue()) / int64(n2.IntValue())), nil
				} else if info.hasMultiply {
					return types.NewIntNumber(int64(n1.IntValue()) * int64(n2.IntValue())), nil
				}
			}
		} else {
			if f1IsLambda || f2IsLambda {
				if info.hasRange {
					return types.Null, burrito.WrappedErrorf("Cannot use a lambda in a range expression")
				}
				return types.Null, burrito.WrappedErrorf("Cannot use a lambda in an arithmetic operation")
			}
			return types.NaN(), nil
		}
	} else if len(info.fields) == 3 {
		// handle ternary operator
		if info.hasQuestion {
			f1, err := v.Visit(info.fields[0])
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			if types.AsBool(f1).BoolValue() {
				return v.Visit(info.fields[1])
			} else {
				return v.Visit(info.fields[2])
			}
		}
	}
	// if the field is another field in parentheses, return the value of that field
	// we need to also check if the first element is not a field, because if it is, it will be a function call
	if info.hasLeftParen && len(info.fields) == 1 && context.GetChild(0) != info.fields[0] {
		return v.Visit(info.fields[0])
	} else if info.hasLeftParen && len(info.fields) == 1 && context.GetChild(0) == info.fields[0] {
		return v.visitCall(info)
	}
	// handle accessing a property of an object or calling an instance function
	if info.name != nil && len(info.fields) == 1 {
		object, err := v.Visit(info.fields[0])
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		return v.memberAccess(context, info, object)
	}
	if info.name != nil {
		return v.resolveScopeStr(info.nameText, info.nameStr), nil
	}
	// handle indexed access
	if info.index != nil && len(info.fields) == 1 {
		i, err := v.visitIndexInfo(info.index, &info.indexInfo)
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		if _, ok := i.(*types.JsonLambda); ok {
			return types.Null, burrito.WrappedErrorf("Cannot use a lambda as an index")
		}
		object, err := v.Visit(info.fields[0])
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		if _, ok := object.(*types.JsonLambda); ok {
			return types.Null, burrito.WrappedErrorf("Cannot index a lambda")
		}
		index, err := object.Index(i)
		if err != nil {
			if info.hasQuestion || v.action == types.Predicate || v.isInLeftSideOfAssignment(context) {
				return types.NullWithParent(object, i), nil
			} else {
				return types.Null, burrito.WrapErrorf(err, "Cannot access %s", context.GetText())
			}
		}
		index.UpdateParent(object, i)
		return index, nil
	}
	if info.number != nil {
		return info.numberVal, nil
	}
	if info.escapedString != nil {
		return types.NewString(info.escapedVal), nil
	}
	// template string literal
	if info.templateString != nil {
		return v.Visit(info.templateString)
	}
	// literal array notation
	if info.array != nil {
		return v.Visit(info.array)
	}
	// literal object notation
	if info.object != nil {
		return v.Visit(info.object)
	}
	// lambda literal in field position
	if info.lambda != nil {
		return v.Visit(info.lambda)
	}
	// negation
	if info.hasSubtract && len(info.fields) == 1 {
		f, err := v.Visit(info.fields[0])
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		if _, ok := f.(*types.JsonLambda); ok {
			return types.Null, burrito.WrappedErrorf("Cannot negate a lambda")
		}
		return f.Negate(), nil
	}
	// unary plus
	if info.hasAdd && len(info.fields) == 1 {
		f, err := v.Visit(info.fields[0])
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		if types.IsNumber(f) {
			return f, nil
		}
		return types.NaN(), nil
	}
	return types.Null, burrito.WrappedErrorf("Failed to resolve '%s'", context.GetText())
}

func (v *ExpressionVisitor) VisitName(context *parser.NameContext) (types.JsonType, error) {
	text := v.nameText(context)
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
	info := computeIndexCtxInfo(context)
	return v.visitIndexInfo(context, &info)
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
			obj, err := v.Visit(sf.Field(0))
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
			name, field, err := v.visitObjectFieldEntry(sf)
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			result.Put(name, field)
		}
	}
	return result, nil
}

// visitObjectFieldEntry resolves a non-spread object literal entry to its key and value.
// Supported shapes: `name: value`, `"string": value`, `[computed]: value` and the
// shorthand `name` (equivalent to `name: name`).
func (v *ExpressionVisitor) visitObjectFieldEntry(sf *parser.Object_fieldContext) (string, types.JsonType, error) {
	var name string
	if sf.ESCAPED_STRING() != nil {
		name = unescapeString(types.ToString(cachedGetText(sf.ESCAPED_STRING())))
	} else if sf.LeftBracket() != nil {
		key, err := v.Visit(sf.Field(0))
		if err != nil {
			return "", types.Null, burrito.PassError(err)
		}
		if types.IsNull(key) {
			return "", types.Null, burrito.WrappedErrorf("Computed object key evaluated to null")
		}
		name = key.StringValue()
		field, err := v.Visit(sf.Field(1))
		if err != nil {
			return "", types.Null, burrito.PassError(err)
		}
		return name, field, nil
	} else if sf.Name() != nil {
		name = cachedGetText(sf.Name())
	} else {
		// Bare field: shorthand `{a}` for `{a: a}` when the field is a plain name.
		fCtx, ok := sf.Field(0).(*parser.FieldContext)
		if !ok {
			return "", types.Null, burrito.WrappedErrorf("Invalid object literal entry '%s'", sf.GetText())
		}
		fi := v.fieldInfo(fCtx)
		if fi.name == nil || len(fi.fields) != 0 {
			return "", types.Null, burrito.WrappedErrorf("Object literal entry '%s' must be 'key: value', a spread, or a plain name", sf.GetText())
		}
		name = fi.nameText
	}
	field, err := v.Visit(sf.Field(0))
	if err != nil {
		return "", types.Null, burrito.PassError(err)
	}
	return name, field, nil
}

func (v *ExpressionVisitor) VisitObject_field(context *parser.Object_fieldContext) (types.JsonType, error) {
	if context.Spread() != nil {
		return v.Visit(context.Field(0))
	}
	name, field, err := v.visitObjectFieldEntry(context)
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
		if stmts := ctx.Statements(); stmts != nil {
			info.body = stmts
			info.hasBlock = true
		} else if f := ctx.Field(); f != nil {
			info.body = f
		}
		lambdaInfoCache.Store(ctx, info)
	}
	hasBlock := info.hasBlock
	body := info.body
	params := info.params
	return types.NewLambda(
		func(this *types.JsonLambda, o []types.JsonType) (types.JsonType, error) {
			if len(params) > len(o) {
				return types.Null, burrito.WrappedErrorf("Lambda expects %d arguments, but got %d", len(params), len(o))
			}
			for i, argName := range params {
				v.pushScopePair(argName, o[i])
			}
			outerLocalVarTarget := v.localVarTarget
			v.localVarTarget = getLocalVarObject()
			var result types.JsonType
			var err error
			if hasBlock {
				result, err = v.Visit(body)
				if err == nil {
					if signal, ok := result.(*types.JsonSignal); ok {
						result = signal.Value
					} else {
						result = types.Null
					}
				}
			} else {
				result, err = v.Visit(body)
			}
			putLocalVarObject(v.localVarTarget)
			v.localVarTarget = outerLocalVarTarget
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
	// ESCAPED_STRING tokens are delimited by single-byte quote characters.
	if len(str) < 3 {
		return ""
	}
	return UnescapeString(str[1 : len(str)-1])
}
