package types

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"strconv"
)

// JsonBool is a struct that represents a number, that can be either integer or decimal.
type JsonBool struct {
	JsonType
	Value bool
}

var True = JsonBool{Value: true}
var False = JsonBool{Value: false}

func (n JsonBool) IsNull() bool {
	return false
}

func (n JsonBool) LessThan(other JsonType) (bool, error) {
	if other == nil {
		return false, nil
	}
	if IsBool(other) {
		if n.BoolValue() == AsBool(other).BoolValue() || n.BoolValue() {
			return false, nil
		}
		return true, nil
	}
	if IsNumber(other) {
		if n.Value {
			return 1 < AsNumber(other).FloatValue(), nil
		} else {
			return 0 < AsNumber(other).FloatValue(), nil
		}
	}
	return false, incompatibleTypesError(n, other)
}

func (n JsonBool) BoolValue() bool {
	return n.Value
}

func (n JsonBool) StringValue() string {
	return strconv.FormatBool(n.BoolValue())
}

func (n JsonBool) Equals(value JsonType) bool {
	if value == Null {
		return false
	}
	if IsBool(value) {
		return n.BoolValue() == value.BoolValue()
	}
	if IsNumber(value) {
		return n.BoolValue() == value.BoolValue()
	}
	return false
}

func (n JsonBool) Unwrap() interface{} {
	return n.BoolValue()
}

func (n JsonBool) Negate() JsonType {
	return NewBool(!n.BoolValue())
}

func (n JsonBool) Index(i JsonType) (JsonType, error) {
	return Null, burrito.WrappedErrorf("Cannot access %s from a boolean", i.StringValue())
}

func (n JsonBool) Add(i JsonType) JsonType {
	if IsNumber(i) {
		return AsNumber(i).Add(n)
	}
	if IsBool(i) {
		result := 0
		if n.BoolValue() {
			result++
		}
		if AsBool(i).BoolValue() {
			result++
		}
		return JsonNumber{
			Value:   float64(result),
			Decimal: false,
		}
	}
	return NewString(n.StringValue() + i.StringValue())
}

// AsBool converts an interface to a boolean.
func AsBool(obj interface{}) JsonBool {
	if obj == nil {
		return False
	}
	if b, ok := obj.(bool); ok {
		return NewBool(b)
	}
	if b, ok := obj.(int); ok && b != 0 {
		return True
	}
	if b, ok := obj.(float64); ok && b != 0 {
		return True
	}
	if b, ok := obj.(float32); ok && b != 0 {
		return True
	}
	if b, ok := obj.(JsonType); ok {
		return NewBool(b.BoolValue())
	}
	return NewBool(obj != nil)
}

// IsBool returns true if the given interface is a boolean.
func IsBool(obj interface{}) bool {
	if obj == nil {
		return false
	}
	if _, ok := obj.(bool); ok {
		return true
	}
	if _, ok := obj.(JsonBool); ok {
		return true
	}
	return false
}

func NewBool(value bool) JsonBool {
	if value {
		return True
	}
	return False
}
