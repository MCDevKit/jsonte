package jsonte

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/functions"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/MCDevKit/jsonte/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/gammazero/deque"
	"reflect"
)

const DefaultName = "value"
const DefaultIndexName = "index"

type ExpressionVisitor struct {
	parser.BaseJsonTemplateVisitor
	action        types.JsonAction
	name          *string
	indexName     *string
	scope         deque.Deque[*types.JsonObject]
	path          *string
	usedVariables []string
	variableScope *types.JsonObject
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
		if f, ok := parent.(*parser.FieldContext); ok && f.Literal() != nil && len(f.AllField()) == 2 && f.Field(0) == ctx {
			return true
		}
		return false
	})
}

// resolveLambdaTree resolves a string to an AST tree
func (v *ExpressionVisitor) resolveLambdaTree(src string) parser.ILambdaContext {
	is := antlr.NewInputStream(src)
	lexer := parser.NewJsonTemplateLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(antlr.NewConsoleErrorListener())
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewJsonTemplateParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewConsoleErrorListener())
	p.BuildParseTrees = true
	return p.Lambda()
}

// pushScope pushes a new scope to the stack
func (v *ExpressionVisitor) pushScope(scope map[string]interface{}) {
	v.scope.PushBack(types.AsObject(scope))
}

// pushScopePair pushes a new entry to the stack
func (v *ExpressionVisitor) pushScopePair(key string, value interface{}) {
	v.scope.PushBack(types.AsObject(map[string]interface{}{key: value}))
}

// popScope pops the last scope from the stack
func (v *ExpressionVisitor) popScope() {
	v.scope.PopBack()
}

func isError(v interface{}) bool {
	_, err := v.(error)
	return err
}

// ResolveScope resolves a value from the scope by name
func (v *ExpressionVisitor) ResolveScope(name string) types.JsonType {
	if name == "this" {
		return &types.JsonObject{
			Value:       nil,
			StackValue:  &v.scope,
			StackTarget: v.variableScope,
		}
	}
	if v.variableScope.ContainsKey(name) {
		get := v.variableScope.Get(name)
		get.UpdateParent(v.variableScope, types.NewString(name))
		return get
	}
	for i := v.scope.Len() - 1; i >= 0; i-- {
		m := v.scope.At(i)
		if m.ContainsKey(name) {
			get := m.Get(name)
			get.UpdateParent(m, types.NewString(name))
			return get
		}
	}
	return types.NullWithParent(v.variableScope, types.NewString(name))
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
				v.pushScopePair(indexName, types.AsNumber(i))
			}
			val, err := v.Visit(ctx.Statements(0))
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			if hasIndex {
				v.popScope()
			}
			v.popScope()
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
	v.name = &name
	v.indexName = &indexName
	return v.Visit(ctx.Field())
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
	if context.Literal() != nil {
		left, err := v.Visit(context.Field(0))
		if err != nil {
			return types.Null, burrito.WrapErrorf(err, "Error resolving left side of the statement")
		}
		right, err := v.Visit(context.Field(1))
		if err != nil {
			return types.Null, burrito.WrapErrorf(err, "Error resolving right side of the statement")
		}
		right = types.Box(right.Unbox())
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
		if context.NullCoalescing() != nil {
			if types.IsNull(f1) {
				return f2, nil
			} else {
				return f1, nil
			}
		} else if context.Add() != nil {
			return f1.Add(f2), nil
		} else if context.Equal() != nil {
			return types.AsBool(f1.Equals(f2)), nil
		} else if context.NotEqual() != nil {
			return types.AsBool(!f1.Equals(f2)), nil
		} else if context.Less() != nil {
			than, err := f1.LessThan(f2)
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			return types.AsBool(than), nil
		} else if context.Greater() != nil {
			than, err := f1.LessThan(f2)
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			return types.AsBool(!than && !f1.Equals(f2)), nil
		} else if context.LessOrEqual() != nil {
			than, err := f1.LessThan(f2)
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			return types.AsBool(than || f1.Equals(f2)), nil
		} else if context.GreaterOrEqual() != nil {
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
					return &types.JsonNumber{
						Value:   n1.FloatValue() - n2.FloatValue(),
						Decimal: true,
					}, nil
				} else if context.Divide() != nil {
					return &types.JsonNumber{
						Value:   n1.FloatValue() / n2.FloatValue(),
						Decimal: true,
					}, nil
				} else if context.Multiply() != nil {
					return &types.JsonNumber{
						Value:   n1.FloatValue() * n2.FloatValue(),
						Decimal: true,
					}, nil
				}
			} else {
				if context.Subtract() != nil {
					return &types.JsonNumber{
						Value:   float64(n1.IntValue() - n2.IntValue()),
						Decimal: false,
					}, nil
				} else if context.Divide() != nil {
					return &types.JsonNumber{
						Value:   float64(n1.IntValue() / n2.IntValue()),
						Decimal: false,
					}, nil
				} else if context.Multiply() != nil {
					return &types.JsonNumber{
						Value:   float64(n1.IntValue() * n2.IntValue()),
						Decimal: false,
					}, nil
				}
			}
		} else {
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
					text := b.GetText()
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
		text := context.Name().GetText()
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
			index, err := object.Index(types.NewString(text))
			if err != nil {
				if context.Question() != nil || v.action == types.Predicate || isInLeftSideOfAssignment(context) {
					return types.NullWithParent(object, types.NewString(text)), nil
				} else {
					return types.Null, burrito.WrapErrorf(err, "Cannot access %s", context.GetText())
				}
			}
			index.UpdateParent(object, types.NewString(text))
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
		object, err := v.Visit(context.Field(0))
		if err != nil {
			return types.Null, burrito.PassError(err)
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
		return types.AsNumber(context.NUMBER().GetText()), nil
	}
	if context.ESCAPED_STRING() != nil {
		return types.NewString(unescapeString(types.ToString(context.ESCAPED_STRING().GetText()))), nil
	}
	// literal array notation
	if context.Array() != nil {
		return v.Visit(context.Array())
	}
	// literal object notation
	if context.Object() != nil {
		return v.Visit(context.Object())
	}
	// negation
	if context.Subtract() != nil && len(context.AllField()) == 1 {
		f, err := v.Visit(context.Field(0))
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		return f.Negate(), nil
	}
	return types.Null, burrito.WrappedErrorf("Failed to resolve '%s'", context.GetText())
}

func (v *ExpressionVisitor) VisitName(context *parser.NameContext) (types.JsonType, error) {
	text := context.GetText()
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
		return types.AsNumber(context.NUMBER().GetText()), nil
	}
	if context.ESCAPED_STRING() != nil {
		return types.NewString(unescapeString(types.ToString(context.ESCAPED_STRING().GetText()))), nil
	}
	if context.Field() != nil {
		return v.Visit(context.Field())
	}
	return types.Null, burrito.WrappedErrorf("Invalid index: %s", context.GetText())
}

func (v *ExpressionVisitor) VisitArray(context *parser.ArrayContext) (types.JsonType, error) {
	result := make([]types.JsonType, 0)
	for _, f := range context.AllSpread_field() {
		sf := f.(*parser.Spread_fieldContext)
		r, err := v.Visit(sf.Field())
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		if sf.Spread() != nil {
			if _, ok := r.(*types.JsonArray); !ok {
				return types.Null, burrito.WrappedErrorf("Cannot spread %s into an array", r.StringValue())
			}
			result = append(result, r.(*types.JsonArray).Value...)
		} else {
			result = append(result, r)
		}
	}
	return &types.JsonArray{Value: result}, nil
}

func (v *ExpressionVisitor) VisitObject(context *parser.ObjectContext) (types.JsonType, error) {
	result := types.NewJsonObject()
	for _, f := range context.AllObject_field() {
		obj, err := v.Visit(f)
		if err != nil {
			return types.Null, burrito.PassError(err)
		}
		u := obj.(*types.JsonObject)
		for _, key := range u.Keys() {
			result.Put(key, u.Get(key))
		}
	}
	return result, nil
}

func (v *ExpressionVisitor) VisitObject_field(context *parser.Object_fieldContext) (types.JsonType, error) {
	name := ""
	if context.ESCAPED_STRING() != nil {
		name = unescapeString(types.ToString(context.ESCAPED_STRING().GetText()))
	} else if context.Spread() != nil {
		return v.Visit(context.Field())
	} else {
		name = context.Name().GetText()
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
	vars, args, err := ParseLambda(ctx.GetText())
	if err != nil {
		return types.Null, burrito.PassError(err)
	}
	return types.NewLambda(
		func(this *types.JsonLambda, o []types.JsonType) (types.JsonType, error) {
			if len(ctx.AllName()) > len(o) {
				return types.Null, burrito.WrappedErrorf("Lambda expects %d arguments, but got %d", len(ctx.AllName()), len(o))
			}
			for i, context := range ctx.AllName() {
				// Ensure we get boxed type
				v.pushScopePair(context.GetText(), o[i])
			}
			result, err := v.Visit(ctx.Field())
			if err != nil {
				return types.Null, burrito.PassError(err)
			}
			for i := 0; i < len(ctx.AllName()); i++ {
				v.popScope()
			}
			return result, nil
		},
		ctx.GetText(),
		vars,
		args,
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

// unescapeString removes quotes and unescapes a string.
func unescapeString(str string) string {
	runes := []rune(str)
	if len(runes) < 3 {
		return ""
	}
	str = string(runes[1 : len(runes)-1])
	return UnescapeString(str)
}
