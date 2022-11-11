package jsonte

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/functions"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gammazero/deque"
	"reflect"
	"strings"
)

const DefaultName = "value"
const DefaultIndexName = "index"

type ExpressionVisitor struct {
	parser.BaseJsonTemplateVisitor
	action    types.JsonAction
	name      *string
	indexName *string
	scope     deque.Deque[types.JsonObject]
	path      *string
}

func (v *ExpressionVisitor) Visit(tree antlr.ParseTree) (types.JsonType, error) {
	switch val := tree.(type) {
	case *parser.FieldContext:
		return v.VisitField(val)
	case *parser.ArrayContext:
		return v.VisitArray(val)
	case *parser.ObjectContext:
		return v.VisitObject(val)
	case *parser.Object_fieldContext:
		return v.VisitObject_field(val)
	case *parser.ExpressionContext:
		return v.VisitExpression(val)
	case *parser.Function_paramContext:
		return v.VisitFunction_param(val)
	case *parser.LambdaContext:
		return v.VisitLambda(val)
	case *parser.NameContext:
		return v.VisitName(val)
	case *parser.IndexContext:
		return v.VisitIndex(val)
	}
	panic("Unknown tree type " + reflect.TypeOf(tree).String())
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

// resolveScope resolves a value from the scope by name
func (v *ExpressionVisitor) resolveScope(name string) types.JsonType {
	for i := v.scope.Len() - 1; i >= 0; i-- {
		m := v.scope.At(i)
		if m.ContainsKey(name) {
			return m.Get(name)
		}
	}
	return types.Null
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
		return types.True, nil
	}
	if context.False() != nil {
		return types.False, nil
	}
	if context.Not() != nil {
		visit, err := v.Visit(context.Field(0))
		if err != nil {
			return types.Null, err
		}
		return visit.Negate(), nil
	}
	// process field composed of two other fields
	if len(context.AllField()) == 2 {
		f1, err := v.Visit(context.Field(0))
		if err != nil {
			return types.Null, err
		}
		// Move AND and OR here to make those operators short-circuiting
		if context.And() != nil {
			if f1.BoolValue() {
				f2, err := v.Visit(context.Field(1))
				if err != nil {
					return types.Null, err
				}
				return types.AsBool(f2.BoolValue()), nil
			} else {
				return types.False, nil
			}
		} else if context.Or() != nil {
			if f1.Negate().BoolValue() {
				f2, err := v.Visit(context.Field(1))
				if err != nil {
					return types.Null, err
				}
				return types.AsBool(f2.BoolValue()), nil
			} else {
				return types.True, nil
			}
		} else if context.Question() != nil {
			if f1.BoolValue() {
				f2, err := v.Visit(context.Field(1))
				if err != nil {
					return types.Null, err
				}
				return f2, nil
			} else {
				return types.Null, nil
			}
		}
		f2, err := v.Visit(context.Field(1))
		if err != nil {
			return types.Null, err
		}
		if context.NullCoalescing() != nil {
			if f1.IsNull() {
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
				return types.Null, err
			}
			return types.AsBool(than), nil
		} else if context.Greater() != nil {
			than, err := f1.LessThan(f2)
			if err != nil {
				return types.Null, err
			}
			return types.AsBool(!than && !f1.Equals(f2)), nil
		} else if context.LessOrEqual() != nil {
			than, err := f1.LessThan(f2)
			if err != nil {
				return types.Null, err
			}
			return types.AsBool(than || f1.Equals(f2)), nil
		} else if context.GreaterOrEqual() != nil {
			than, err := f1.LessThan(f2)
			if err != nil {
				return types.Null, err
			}
			return types.AsBool(!than), nil
		} else if types.IsNumber(f1) && types.IsNumber(f2) {
			n1 := types.AsNumber(f1)
			n2 := types.AsNumber(f2)
			if context.Range() != nil {
				return types.CreateRange(n1.IntValue(), n2.IntValue()), nil
			} else if n1.Decimal || n2.Decimal {
				if context.Subtract() != nil {
					return types.JsonNumber{
						Value:   n1.FloatValue() - n2.FloatValue(),
						Decimal: true,
					}, nil
				} else if context.Divide() != nil {
					return types.JsonNumber{
						Value:   n1.FloatValue() / n2.FloatValue(),
						Decimal: true,
					}, nil
				} else if context.Multiply() != nil {
					return types.JsonNumber{
						Value:   n1.FloatValue() * n2.FloatValue(),
						Decimal: true,
					}, nil
				}
			} else {
				if context.Subtract() != nil {
					return types.JsonNumber{
						Value:   float64(n1.IntValue() - n2.IntValue()),
						Decimal: false,
					}, nil
				} else if context.Divide() != nil {
					return types.JsonNumber{
						Value:   float64(n1.IntValue() / n2.IntValue()),
						Decimal: false,
					}, nil
				} else if context.Multiply() != nil {
					return types.JsonNumber{
						Value:   float64(n1.IntValue() * n2.IntValue()),
						Decimal: false,
					}, nil
				}
			}
		} else {
			return types.NaN, nil
		}
	} else if len(context.AllField()) == 3 {
		// handle ternary operator
		if context.Question() != nil {
			f1, err := v.Visit(context.Field(0))
			if err != nil {
				return types.Null, err
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
			return types.Null, err
		}
		params := make([]types.JsonType, len(context.AllFunction_param()))
		for i, param := range context.AllFunction_param() {
			p, err := v.Visit(param)
			if err != nil {
				return types.Null, err
			}
			params[i] = p
		}
		if _, ok := lambda.(types.JsonString); ok {
			lambdaContext := v.resolveLambdaTree(lambda.StringValue())
			fun, err := v.Visit(lambdaContext)
			if err != nil {
				return types.Null, err
			}
			if _, ok := fun.(types.JsonLambda); ok {
				i, err := fun.(types.JsonLambda).Value(params)
				if err != nil {
					return types.Null, err
				}
				return i, nil
			} else {
				return types.Null, burrito.WrappedErrorf("%s is not a function", lambda)
			}
		} else if _, ok := lambda.(types.JsonLambda); ok {
			i, err := lambda.(types.JsonLambda).Value(params)
			if err != nil {
				return types.Null, err
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
				return types.Null, burrito.WrappedErrorf("Function '%s' not found!", context.Field(0).GetText())
			}
			function, err := functions.CallFunction(*methodName, params)
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
			return types.Null, err
		}
		if object.IsNull() {
			// handle null-forgiving operator
			if context.Question() != nil {
				return types.Null, nil
			} else {
				return types.Null, burrito.WrappedErrorf("Cannot access %s because %s is %s", context.GetText(), context.Field(0).GetText(), types.ToString(object))
			}
		}
		if functions.HasInstanceFunction(reflect.TypeOf(object), text) {
			return types.JsonLambda{
				Value: func(o []types.JsonType) (types.JsonType, error) {
					function, err := functions.CallInstanceFunction(text, object.(types.JsonType), o)
					if err != nil {
						return types.Null, burrito.WrapErrorf(err, "Error calling function '%s' on %s", text, types.ToString(object))
					}
					return function, nil
				},
			}, nil
		} else {
			index, err := object.Index(types.NewString(text))
			if err != nil {
				if context.Question() != nil || v.action == types.Predicate {
					return types.Null, nil
				} else {
					return types.Null, burrito.WrapErrorf(err, "Cannot access %s because %s is %s", context.GetText(), context.Field(0).GetText(), types.ToString(object))
				}
			}
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
			return types.Null, err
		}
		object, err := v.Visit(context.Field(0))
		if err != nil {
			return types.Null, err
		}
		index, err := object.Index(i)
		if err != nil {
			if context.Question() != nil {
				return types.Null, nil
			} else {
				return types.Null, burrito.WrapErrorf(err, "Cannot access %s because %s is %s", context.GetText(), context.Field(0).GetText(), types.ToString(object))
			}
		}
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
			return types.Null, err
		}
		return f.Negate(), nil
	}
	return types.Null, burrito.WrappedErrorf("Failed to resolve '%s'", context.GetText())
}

func (v *ExpressionVisitor) VisitName(context *parser.NameContext) (types.JsonType, error) {
	text := context.GetText()
	// "this" is a special case that always refers to the current full scope
	if text == "this" {
		scope := types.NewJsonObject()
		for i := 0; i < v.scope.Len(); i++ {
			s := v.scope.At(i)
			for _, key := range s.Keys() {
				scope.Put(key, s.Get(key))
			}
		}
		return scope, nil
	}
	newScope := v.resolveScope(text)

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
	result := make([]types.JsonType, len(context.AllField()))
	for i, f := range context.AllField() {
		r, err := v.Visit(f)
		if err != nil {
			return types.Null, err
		}
		result[i] = r
	}
	return types.JsonArray{Value: result}, nil
}

func (v *ExpressionVisitor) VisitObject(context *parser.ObjectContext) (types.JsonType, error) {
	result := types.NewJsonObject()
	for _, f := range context.AllObject_field() {
		obj, err := v.Visit(f)
		if err != nil {
			return types.Null, err
		}
		u := obj.(types.JsonObject)
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
	} else {
		name = context.Name().GetText()
	}
	field, err := v.Visit(context.Field())
	if err != nil {
		return types.Null, err
	}
	n := types.NewJsonObject()
	n.Put(name, field)
	return n, nil
}

func (v *ExpressionVisitor) VisitLambda(ctx *parser.LambdaContext) (types.JsonType, error) {
	return types.JsonLambda{
		Value: func(o []types.JsonType) (types.JsonType, error) {
			if len(ctx.AllName()) > len(o) {
				return types.Null, burrito.WrappedErrorf("Lambda expects %d arguments, but got %d", len(ctx.AllName()), len(o))
			}
			for i, context := range ctx.AllName() {
				// Ensure we get boxed type
				v.pushScopePair(context.GetText(), o[i])
			}
			result, err := v.Visit(ctx.Field())
			if err != nil {
				return types.Null, err
			}
			for i := 0; i < len(ctx.AllName()); i++ {
				v.popScope()
			}
			return result, nil
		},
	}, nil
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
	if len(str) < 3 {
		return ""
	}
	str = str[1 : len(str)-1]
	str = strings.ReplaceAll(str, "\\\\\"", "\"")
	str = strings.ReplaceAll(str, "\\\\'", "'")
	str = strings.ReplaceAll(str, "\\\\n", "\n")
	str = strings.ReplaceAll(str, "\\\\\\\\", "\\\\")
	return str
}
