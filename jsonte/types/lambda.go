package types

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
)

// JsonLambda is a struct that represents a lambda value.
type JsonLambda struct {
	Value         func(this *JsonLambda, args []JsonType) (JsonType, error)
	Arguments     []string
	UsedVariables []string
	String        string
}

var IdentityLambda = NewLambda(func(this *JsonLambda, args []JsonType) (JsonType, error) {
	if len(args) != 1 {
		return nil, burrito.WrappedErrorf("Identity lambda must have exactly 1 argument")
	}
	return args[0], nil
}, "x=>x", []string{"x"}, []string{"x"})

func (t *JsonLambda) Call(args ...JsonType) (JsonType, error) {
	return t.Value(t, args)
}

func (t *JsonLambda) Parent() JsonType {
	return nil
}

func (t *JsonLambda) ParentIndex() JsonType {
	return nil
}

func (t *JsonLambda) UpdateParent(parent JsonType, parentIndex JsonType) {
}

func (t *JsonLambda) LessThan(JsonType) (bool, error) {
	return false, burrito.WrappedErrorf("Lambdas cannot be compared")
}

func (t *JsonLambda) BoolValue() bool {
	return true
}

func (t *JsonLambda) StringValue() string {
	return t.String
}

func (t *JsonLambda) Equals(JsonType) bool {
	return false
}

func (t *JsonLambda) Unbox() interface{} {
	return t.String
}

func (t *JsonLambda) Negate() JsonType {
	return True()
}

func (t *JsonLambda) Index(i JsonType) (JsonType, error) {
	return Null, burrito.WrappedErrorf("Cannot access %s from a lambda", i.StringValue())
}

func (t *JsonLambda) Add(i JsonType) JsonType {
	return NewString(t.StringValue() + i.StringValue())
}

func NewLambda(value func(this *JsonLambda, args []JsonType) (JsonType, error), stringValue string, vars, args []string) *JsonLambda {
	return &JsonLambda{Value: value, String: stringValue, UsedVariables: vars, Arguments: args}
}
