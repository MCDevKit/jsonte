package jsonte

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gammazero/deque"
	"jsonte/jsonte/functions"
	"jsonte/jsonte/utils"
	"jsonte/parser"
	"os"
	"reflect"
)

const DefaultName = "value"
const NaN = "NaN"

type ExpressionVisitor struct {
	parser.BaseJsonTemplateVisitor
	action       utils.JsonAction
	name         *string
	fullScope    utils.JsonObject
	extraScope   utils.JsonObject
	currentScope deque.Deque[interface{}]
	path         *string
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
	return nil
}

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

func (v *ExpressionVisitor) pushScope(scope map[string]interface{}) {
	v.currentScope.PushBack(scope)
}

func (v *ExpressionVisitor) pushScopePair(key string, value interface{}) {
	v.currentScope.PushBack(map[string]interface{}{key: value})
}

func (v *ExpressionVisitor) popScope() {
	v.currentScope.PopBack()
}

func negate(value interface{}) interface{} {
	if value == nil {
		return nil
	}
	if b, ok := value.(bool); ok {
		return !b
	}
	if b, ok := value.(int); ok {
		return -b
	}
	if b, ok := value.(float64); ok {
		return -b
	}
	if b, ok := value.(float32); ok {
		return -b
	}
	if utils.IsArray(value) {
		result := make(utils.JsonArray, len(value.(utils.JsonArray)))
		for i, v := range value.(utils.JsonArray) {
			result[i] = negate(v)
		}
		return result
	}
	return NaN
}

func isError(v interface{}) bool {
	_, ok := v.(error)
	return ok
}

func (v *ExpressionVisitor) resolveScope(name string) interface{} {
	for i := v.currentScope.Len() - 1; i >= 0; i-- {
		m := v.currentScope.At(i)
		if c, ok := m.(utils.JsonObject); ok {
			if v, ok := c[name]; ok {
				return v
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
	name := DefaultName
	if ctx.As() != nil {
		name = ctx.Name().GetText()
	}
	v.name = &name
	result := v.Visit(ctx.Field())
	if isError(result) {
		// DO YOU REALLY NEED TO RETURN AN ERROR HERE???
		// I JUST WANT TO PRINT THE ERROR AND CONTINUE
		_, err := fmt.Fprintln(os.Stderr, result)
		if err != nil {
			return nil
		}
		return nil
	}
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
			if utils.IsNumber(f1) && utils.IsNumber(f2) {
				n1 := utils.ToNumber(f1)
				n2 := utils.ToNumber(f2)
				return utils.JsonNumber{
					Value:   n1.FloatValue() + n2.FloatValue(),
					Decimal: n1.Decimal || n2.Decimal,
				}
			} else if utils.IsArray(f1) && utils.IsArray(f2) {
				array := utils.MergeArray(f1.(utils.JsonArray), f2.(utils.JsonArray))
				return array
			} else if utils.IsObject(f1) && utils.IsObject(f2) {
				var result utils.JsonArray = nil
				result = append(result, f1.(utils.JsonArray)...)
				return append(result, f2.(utils.JsonArray)...)
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
				return &utils.EvaluationError{
					Message: fmt.Sprintf("Function '%s' not found!", context.Field(0).GetText()),
					Path:    *v.path,
				}
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
				return &utils.EvaluationError{
					Message: fmt.Sprintf("Function '%s' not found!", context.Field(0).GetText()),
					Path:    *v.path,
				}
			}
			function, err := functions.CallFunction(*methodName, params)
			if err != nil {
				return err
			}
			return function
		}
	}
	if context.Name() != nil && len(context.AllField()) == 1 {
		text := context.Name().GetText()
		object := v.Visit(context.Field(0))
		if isError(object) {
			return object
		}
		var newScope utils.JsonObject = nil
		if utils.IsObject(object) {
			if v, ok := object.(utils.JsonObject)[text]; ok {
				newScope = v.(utils.JsonObject)
			}
		} else if functions.HasInstanceFunction(reflect.TypeOf(object), text) {
			return utils.JsonLambda(
				func(o []interface{}) (interface{}, error) {
					return functions.CallInstanceFunction(text, object, o)
				},
			)
		}

		if newScope == nil {
			if v.action == utils.Predicate || context.Question() != nil {
				return nil
			} else {
				return utils.EvaluationError{
					Message: "Failed to resolve \"" + text + "\"",
					Path:    *v.path,
				}
			}
		}
		return newScope
	}
	if context.Name() != nil {
		return v.Visit(context.Name())
	}
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
			if utils.IsNumber(i) {
				value := utils.ToNumber(i).IntValue()
				if value < 0 || value >= len(object.(utils.JsonArray)) {
					if context.Question() != nil {
						return nil
					} else {
						return utils.EvaluationError{
							Message: "Array index out of bounds!",
							Path:    *v.path,
						}
					}
				}
				return object.(utils.JsonArray)[value]
			} else {
				if context.Question() != nil {
					return nil
				} else {
					return utils.EvaluationError{
						Message: "Array index is not a number!",
						Path:    *v.path,
					}
				}
			}
		} else if utils.IsObject(object) {
			value := utils.ToString(i)
			if b, ok := object.(utils.JsonObject)[value]; !ok {
				if context.Question() != nil {
					return nil
				} else {
					return utils.EvaluationError{
						Message: "Object does not have key \"" + value + "\"!",
						Path:    *v.path,
					}
				}
			} else {
				return b
			}
		}
	}
	if context.NUMBER() != nil {
		return utils.ToNumber(context.NUMBER().GetText())
	}
	if context.ESCAPED_STRING() != nil {
		return utils.UnescapeString(utils.ToString(context.ESCAPED_STRING().GetText()))
	}
	if context.Array() != nil {
		return v.Visit(context.Array())
	}
	if context.Object() != nil {
		return v.Visit(context.Object())
	}
	if context.Subtract() != nil && len(context.AllField()) == 1 {
		f := v.Visit(context.Field(0))
		if isError(f) {
			return f
		}
		return negate(f)
	}
	return nil
}

func (v *ExpressionVisitor) VisitName(context *parser.NameContext) interface{} {
	text := context.GetText()
	if text == "this" {
		scope := utils.JsonObject{}
		for key, value := range v.fullScope {
			scope[key] = value
		}
		for key, value := range v.extraScope {
			scope[key] = value
		}
		for i := 0; i < v.currentScope.Len(); i++ {
			s := v.currentScope.At(i)
			if c, ok := s.(utils.JsonObject); ok {
				for key, value := range c {
					scope[key] = value
				}
			}
		}
		return scope
	}
	if text == "value" {
		if v.currentScope.Len() > 0 {
			return v.currentScope.Back()
		} else {
			return nil
		}
	}
	newScope := v.resolveScope(text)

	back := v.currentScope.Len() - 1
	for newScope == nil && back >= 0 {
		m := v.currentScope.At(back)
		if c, ok := m.(utils.JsonObject); ok {
			if v, ok := c[text]; ok {
				newScope = v
			}
		}
		back--
	}
	if newScope == nil {
		if b, ok := v.extraScope[text]; ok {
			newScope = b
		}
	}
	if newScope == nil {
		if b, ok := v.fullScope[text]; ok {
			newScope = b
		}
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
	return utils.JsonNumber{
		Value:   -1,
		Decimal: false,
	}
}

func (v *ExpressionVisitor) VisitArray(context *parser.ArrayContext) interface{} {
	result := make(utils.JsonArray, len(context.AllField()))
	for i, f := range context.AllField() {
		result[i] = v.Visit(f)
		if isError(result[i]) {
			return result[i]
		}
	}
	return result
}

func (v *ExpressionVisitor) VisitObject(context *parser.ObjectContext) interface{} {
	result := utils.JsonObject{}
	for _, f := range context.AllObject_field() {
		obj := v.Visit(f)
		if isError(obj) {
			return obj
		}
		for key, value := range obj.(utils.JsonObject) {
			result[key] = value
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
	return utils.JsonObject{
		name: field,
	}
}

func (v *ExpressionVisitor) VisitLambda(ctx *parser.LambdaContext) interface{} {
	return utils.JsonLambda(
		func(o []interface{}) (interface{}, error) {
			if len(ctx.AllName()) > len(o) {
				return nil, &utils.EvaluationError{
					Message: fmt.Sprintf("Lambda requires %d parameters, but only %d were supplied!", len(ctx.AllName()), len(o)),
					Path:    *v.path,
				}
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
	return nil
}
