package types

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
)

// JsonLambda is a struct that represents a lambda value.
type JsonLambda struct {
	JsonType
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

func (n JsonLambda) LessThan(JsonType) (bool, error) {
	return false, burrito.WrappedErrorf("Lambdas cannot be compared")
}

func (n JsonLambda) BoolValue() bool {
	return true
}

func (n JsonLambda) StringValue() string {
	return n.String
}

func (n JsonLambda) Equals(JsonType) bool {
	return false
}

func (n JsonLambda) Unbox() interface{} {
	return n.String
}

func (n JsonLambda) Negate() JsonType {
	return True
}

func (n JsonLambda) Index(i JsonType) (JsonType, error) {
	return Null, burrito.WrappedErrorf("Cannot access %s from a lambda", i.StringValue())
}

func (n JsonLambda) Add(i JsonType) JsonType {
	return NewString(n.StringValue() + i.StringValue())
}

func NewLambda(value func(this *JsonLambda, args []JsonType) (JsonType, error), stringValue string, vars, args []string) JsonLambda {
	return JsonLambda{Value: value, String: stringValue, UsedVariables: vars, Arguments: args}
}
