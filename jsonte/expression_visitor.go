package jsonte

import (
	"github.com/MCDevKit/jsonte/jsonte/functions"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/MCDevKit/jsonte/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gammazero/deque"
	"reflect"
)

const DefaultName = "value"
const DefaultIndexName = "index"
const NaN = "NaN"

type ExpressionVisitor struct {
	parser.BaseJsonTemplateVisitor
	action    utils.JsonAction
	name      *string
	indexName *string
	scope     deque.Deque[interface{}]
	path      *string
}

func (v *ExpressionVisitor) Visit(tree antlr.ParseTree) interface{} {
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
	v.scope.PushBack(scope)
}

// pushScopePair pushes a new entry to the stack
func (v *ExpressionVisitor) pushScopePair(key string, value interface{}) {
	v.scope.PushBack(map[string]interface{}{key: value})
}

// popScope pops the last scope from the stack
func (v *ExpressionVisitor) popScope() {
	v.scope.PopBack()
}

// negate negates a value
func negate(value interface{}) interface{} {
	if value == nil {
		return nil
	}
	if b, ok := value.(utils.JsonNumber); ok {
		return utils.JsonNumber{
			Value:   -b.Value,
			Decimal: b.Decimal,
		}
	}
	if b, ok := value.(bool); ok {
		return !b
	}
	if b, ok := value.(int); ok {
		return utils.ToNumber(-b)
	}
	if b, ok := value.(float64); ok {
		return utils.ToNumber(-b)
	}
	if b, ok := value.(float32); ok {
		return utils.ToNumber(-b)
	}
	if utils.IsArray(value) {
		result := make([]interface{}, len(value.([]interface{})))
		for i, v := range value.([]interface{}) {
			result[i] = negate(v)
		}
		return result
	}
	return NaN
}

func isError(v interface{}) bool {
	_, err := v.(error)
	return err
}

func getError(v interface{}) error {
	if err, ok := v.(error); ok {
		return err
	}
	return nil
}

// resolveScope resolves a value from the scope by name
func (v *ExpressionVisitor) resolveScope(name string) interface{} {
	for i := v.scope.Len() - 1; i >= 0; i-- {
		m := v.scope.At(i)
		if c, ok := m.(utils.NavigableMap[string, interface{}]); ok {
			if c.ContainsKey(name) {
				return c.Get(name)
			}
		}
		// Seems like sometimes above cast fails, so we need to check for map[string]interface{} as well
		if c, ok := m.(map[string]interface{}); ok {
			if v, ok := c[name]; ok {
				return v
			}
		}
	}
	return nil
}

func (v *ExpressionVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	// check the type of the expression
	v.action = utils.Value
	if ctx.Iteration() != nil {
		v.action = utils.Iteration
	}
	if ctx.Literal() != nil {
		v.action = utils.Literal
	}
	if ctx.Question() != nil {
		v.action = utils.Predicate
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
	result := v.Visit(ctx.Field())
	return result
}

func (v *ExpressionVisitor) VisitField(context *parser.FieldContext) interface{} {
	if context.Null() != nil {
		return nil
	}
	if context.True() != nil {
		return true
	}
	if context.False() != nil {
		return false
	}
	if context.Not() != nil {
		return !utils.ToBoolean(v.Visit(context.Field(0)))
	}
	// process field composed of two other fields
	if len(context.AllField()) == 2 {
		f1 := v.Visit(context.Field(0))
		if isError(f1) {
			return f1
		}
		// Move AND and OR here to make those operators short-circuiting
		if context.And() != nil {
			if utils.ToBoolean(f1) {
				f2 := v.Visit(context.Field(1))
				if isError(f2) {
					return f2
				}
				return utils.ToBoolean(f2)
			} else {
				return false
			}
		} else if context.Or() != nil {
			b := utils.ToBoolean(f1)
			if !b {
				f2 := v.Visit(context.Field(1))
				if isError(f2) {
					return f2
				}
				return utils.ToBoolean(f2)
			} else {
				return b
			}
		} else if context.Question() != nil {
			if utils.ToBoolean(f1) {
				f2 := v.Visit(context.Field(1))
				if isError(f2) {
					return f2
				}
				return f2
			} else {
				return nil
			}
		}
		f2 := v.Visit(context.Field(1))
		if isError(f2) {
			return f2
		}
		if context.NullCoalescing() != nil {
			if f1 == nil {
				return f2
			} else {
				return f1
			}
		} else if context.Add() != nil {
			if (utils.IsNumber(f1) && utils.IsNumber(f2)) || (utils.IsNumber(f1) && f2 == nil) || (f1 == nil && utils.IsNumber(f2)) {
				n1 := utils.ToNumber(f1)
				n2 := utils.ToNumber(f2)
				return utils.JsonNumber{
					Value:   n1.FloatValue() + n2.FloatValue(),
					Decimal: n1.Decimal || n2.Decimal,
				}
			} else if utils.IsArray(f1) && utils.IsArray(f2) {
				array := utils.MergeArray(f1.([]interface{}), f2.([]interface{}), false)
				return array
			} else if utils.IsObject(f1) && utils.IsObject(f2) {
				var result []interface{} = nil
				result = append(result, f1.([]interface{})...)
				return append(result, f2.([]interface{})...)
			} else {
				return utils.ToString(f1) + utils.ToString(f2)
			}
		} else if context.Equal() != nil {
			return utils.IsEqual(f1, f2)
		} else if context.NotEqual() != nil {
			return !utils.IsEqual(f1, f2)
		} else if utils.IsNumber(f1) && utils.IsNumber(f2) {
			n1 := utils.ToNumber(f1)
			n2 := utils.ToNumber(f2)
			if context.Range() != nil {
				return utils.CreateRange(n1.IntValue(), n2.IntValue())
			} else if context.Greater() != nil {
				return n1.FloatValue() > n2.FloatValue()
			} else if context.Less() != nil {
				return n1.FloatValue() < n2.FloatValue()
			} else if context.GreaterOrEqual() != nil {
				return n1.FloatValue() >= n2.FloatValue()
			} else if context.LessOrEqual() != nil {
				return n1.FloatValue() <= n2.FloatValue()
			} else if n1.Decimal || n2.Decimal {
				if context.Subtract() != nil {
					return utils.JsonNumber{
						Value:   n1.FloatValue() - n2.FloatValue(),
						Decimal: true,
					}
				} else if context.Divide() != nil {
					return utils.JsonNumber{
						Value:   n1.FloatValue() / n2.FloatValue(),
						Decimal: true,
					}
				} else if context.Multiply() != nil {
					return utils.JsonNumber{
						Value:   n1.FloatValue() * n2.FloatValue(),
						Decimal: true,
					}
				}
			} else {
				if context.Subtract() != nil {
					return utils.JsonNumber{
						Value:   float64(n1.IntValue() - n2.IntValue()),
						Decimal: false,
					}
				} else if context.Divide() != nil {
					return utils.JsonNumber{
						Value:   float64(n1.IntValue() / n2.IntValue()),
						Decimal: false,
					}
				} else if context.Multiply() != nil {
					return utils.JsonNumber{
						Value:   float64(n1.IntValue() * n2.IntValue()),
						Decimal: false,
					}
				}
			}
		} else if context.Greater() != nil || context.Less() != nil || context.GreaterOrEqual() != nil || context.LessOrEqual() != nil {
			return false
		} else {
			return NaN
		}
	} else if len(context.AllField()) == 3 {
		// handle ternary operator
		if context.Question() != nil {
			f1 := v.Visit(context.Field(0))
			if isError(f1) {
				return f1
			}
			if utils.ToBoolean(f1) {
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
		lambda := v.Visit(context.Field(0))
		if isError(lambda) {
			return lambda
		}
		params := make([]interface{}, len(context.AllFunction_param()))
		for i, param := range context.AllFunction_param() {
			params[i] = v.Visit(param)
			if isError(params[i]) {
				return params[i]
			}
		}
		if _, ok := lambda.(string); ok {
			lambdaContext := v.resolveLambdaTree(lambda.(string))
			fun := v.Visit(lambdaContext)
			if isError(fun) {
				return fun
			}
			if _, ok := fun.(utils.JsonLambda); ok {
				i, err := fun.(utils.JsonLambda)(params)
				if err != nil {
					return err
				}
				return i
			} else {
				return utils.WrappedErrorf("%s is not a function", lambda)
			}
		} else if _, ok := lambda.(utils.JsonLambda); ok {
			i, err := lambda.(utils.JsonLambda)(params)
			if err != nil {
				return err
			}
			return i
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
				return utils.WrappedErrorf("Function '%s' not found!", context.Field(0).GetText())
			}
			function, err := functions.CallFunction(*methodName, params)
			if err != nil {
				return utils.WrapErrorf(err, "Error calling function '%s'", *methodName)
			}
			return function
		}
	}
	// handle accessing a property of an object or calling an instance function
	if context.Name() != nil && len(context.AllField()) == 1 {
		text := context.Name().GetText()
		object := v.Visit(context.Field(0))
		if isError(object) {
			return object
		}
		var newScope interface{} = nil
		if object == nil {
			// handle null-forgiving operator
			if context.Question() != nil {
				return nil
			} else {
				return utils.WrappedErrorf("Cannot access %s because %s is %s", context.GetText(), context.Field(0).GetText(), utils.ToString(object))
			}
		}
		if utils.IsObject(object) {
			u := object.(utils.NavigableMap[string, interface{}])
			if u.ContainsKey(text) {
				newScope = u.Get(text)
			}
		} else if functions.HasInstanceFunction(reflect.TypeOf(object), text) {
			return utils.JsonLambda(
				func(o []interface{}) (interface{}, error) {
					function, err := functions.CallInstanceFunction(text, object, o)
					if err != nil {
						return nil, utils.WrapErrorf(err, "Error calling function '%s' on %s", text, utils.ToString(object))
					}
					return function, nil
				},
			)
		}

		if newScope == nil {
			// handle null-forgiving operator
			if v.action == utils.Predicate || context.Question() != nil {
				return nil
			} else {
				return utils.WrappedErrorf("Failed to resolve field '%s' (%s) in %s", text, context.GetText(), utils.ToString(object))
			}
		}
		return newScope
	}
	if context.Name() != nil {
		return v.Visit(context.Name())
	}
	// handle indexed access
	if context.Index() != nil && len(context.AllField()) == 1 {
		i := v.Visit(context.Index())
		if isError(i) {
			return i
		}
		object := v.Visit(context.Field(0))
		if isError(object) {
			return object
		}
		if utils.IsArray(object) {
			// in case of an array, we need an integer index
			if utils.IsNumber(i) {
				value := utils.ToNumber(i).IntValue()
				if value < 0 || value >= int32(len(object.([]interface{}))) {
					// handle null-forgiving operator
					if context.Question() != nil {
						return nil
					} else {
						return utils.WrappedErrorf("Index out of bounds: %d (%s)", value, context.GetText())
					}
				}
				return object.([]interface{})[value]
			} else {
				// handle null-forgiving operator
				if context.Question() != nil {
					return nil
				} else {
					return utils.WrappedErrorf("Index must be a number: %s (%s)", utils.ToString(i), context.GetText())
				}
			}
		} else if utils.IsObject(object) {
			// in case of an object, we need a string index
			value := utils.ToString(i)
			u := object.(utils.NavigableMap[string, interface{}])
			if !u.ContainsKey(value) {
				// handle null-forgiving operator
				if context.Question() != nil {
					return nil
				} else {
					return utils.WrappedErrorf("Property '%s' (%s) not found in %s", value, context.GetText(), utils.ToString(object))
				}
			} else {
				return u.Get(value)
			}
		} else if str, ok := object.(string); ok {
			// in case of a string, we need an integer index
			if utils.IsNumber(i) {
				value := utils.ToNumber(i).IntValue()
				if value < 0 || value >= int32(len(str)) {
					// handle null-forgiving operator
					if context.Question() != nil {
						return nil
					} else {
						return utils.WrappedErrorf("Index out of bounds: %d (%s)", value, context.GetText())
					}
				}
				return string(rune(str[value]))
			} else {
				// handle null-forgiving operator
				if context.Question() != nil {
					return nil
				} else {
					return utils.WrappedErrorf("Index must be a number: %s (%s)", utils.ToString(object), context.GetText())
				}
			}
		} else {
			// handle null-forgiving operator
			if context.Question() != nil {
				return nil
			} else {
				return utils.WrappedErrorf("Cannot access %s because %s is %s", context.GetText(), context.Field(0).GetText(), utils.ToString(object))
			}
		}
	}
	if context.NUMBER() != nil {
		return utils.ToNumber(context.NUMBER().GetText())
	}
	if context.ESCAPED_STRING() != nil {
		return utils.UnescapeString(utils.ToString(context.ESCAPED_STRING().GetText()))
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
		f := v.Visit(context.Field(0))
		if isError(f) {
			return f
		}
		return negate(f)
	}
	return utils.WrappedErrorf("Failed to resolve '%s'", context.GetText())
}

func (v *ExpressionVisitor) VisitName(context *parser.NameContext) interface{} {
	text := context.GetText()
	// "this" is a special case that always refers to the current full scope
	if text == "this" {
		scope := utils.NewNavigableMap[string, interface{}]()
		for i := 0; i < v.scope.Len(); i++ {
			s := v.scope.At(i)
			if c, ok := s.(utils.NavigableMap[string, interface{}]); ok {
				for _, key := range c.Keys() {
					scope.Put(key, c.Get(key))
				}
			}
		}
		return scope
	}
	newScope := v.resolveScope(text)

	back := v.scope.Len() - 1
	for newScope == nil && back >= 0 {
		m := v.scope.At(back)
		if c, ok := m.(utils.NavigableMap[string, interface{}]); ok {
			if c.ContainsKey(text) {
				newScope = c.Get(text)
			}
		}
		back--
	}
	return newScope
}

func (v *ExpressionVisitor) VisitIndex(context *parser.IndexContext) interface{} {
	if context.NUMBER() != nil {
		return utils.ToNumber(context.NUMBER().GetText())
	}
	if context.ESCAPED_STRING() != nil {
		return utils.UnescapeString(utils.ToString(context.ESCAPED_STRING().GetText()))
	}
	if context.Field() != nil {
		return v.Visit(context.Field())
	}
	return utils.WrappedErrorf("Invalid index: %s", context.GetText())
}

func (v *ExpressionVisitor) VisitArray(context *parser.ArrayContext) interface{} {
	result := make([]interface{}, len(context.AllField()))
	for i, f := range context.AllField() {
		result[i] = v.Visit(f)
		if isError(result[i]) {
			return result[i]
		}
	}
	return result
}

func (v *ExpressionVisitor) VisitObject(context *parser.ObjectContext) interface{} {
	result := utils.NewNavigableMap[string, interface{}]()
	for _, f := range context.AllObject_field() {
		obj := v.Visit(f)
		if isError(obj) {
			return obj
		}
		u := obj.(utils.NavigableMap[string, interface{}])
		for _, key := range u.Keys() {
			result.Put(key, u.Get(key))
		}
	}
	return result
}

func (v *ExpressionVisitor) VisitObject_field(context *parser.Object_fieldContext) interface{} {
	name := ""
	if context.ESCAPED_STRING() != nil {
		name = utils.UnescapeString(utils.ToString(context.ESCAPED_STRING().GetText()))
	} else {
		name = context.Name().GetText()
	}
	field := v.Visit(context.Field())
	if isError(field) {
		return field
	}
	n := utils.NewNavigableMap[string, interface{}]()
	n.Put(name, field)
	return n
}

func (v *ExpressionVisitor) VisitLambda(ctx *parser.LambdaContext) interface{} {
	return utils.JsonLambda(
		func(o []interface{}) (interface{}, error) {
			if len(ctx.AllName()) > len(o) {
				return nil, utils.WrappedErrorf("Lambda expects %d arguments, but got %d", len(ctx.AllName()), len(o))
			}
			for i, context := range ctx.AllName() {
				// Ensure we get boxed type
				if utils.IsNumber(o[i]) {
					v.pushScopePair(context.GetText(), utils.ToNumber(o[i]))
				} else {
					v.pushScopePair(context.GetText(), o[i])
				}
			}
			result := v.Visit(ctx.Field())
			if isError(result) {
				return nil, result.(error)
			}
			for i := 0; i < len(ctx.AllName()); i++ {
				v.popScope()
			}
			return result, nil
		},
	)
}

func (v *ExpressionVisitor) VisitFunction_param(ctx *parser.Function_paramContext) interface{} {
	if ctx.Field() != nil {
		return v.Visit(ctx.Field())
	}
	if ctx.Lambda() != nil {
		return v.Visit(ctx.Lambda())
	}
	return utils.WrappedErrorf("Invalid function parameter: %s", ctx.GetText())
}
