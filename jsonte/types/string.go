package types

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"strconv"
	"strings"
)

// JsonString is a struct that represents a number, that can be either integer or decimal.
type JsonString struct {
	JsonType
	Value string
}

var EmptyString = JsonString{Value: ""}

func (n JsonString) LessThan(other JsonType) (bool, error) {
	if other == nil || other == Null {
		return false, nil
	}
	if IsString(other) {
		compare := strings.Compare(n.StringValue(), AsString(other).StringValue())
		return compare < 0, nil
	}
	return false, incompatibleTypesError(n, other)
}

func (n JsonString) BoolValue() bool {
	return strings.Trim(n.StringValue(), "\n\r") != ""
}

func (n JsonString) StringValue() string {
	return n.Value
}

func (n JsonString) Equals(value JsonType) bool {
	if value == Null {
		return false
	}
	if IsString(value) {
		return n.StringValue() == value.StringValue()
	}
	return false
}

func (n JsonString) Unbox() interface{} {
	return n.StringValue()
}

func (n JsonString) Add(i JsonType) JsonType {
	return NewString(n.StringValue() + i.StringValue())
}

func (n JsonString) Negate() JsonType {
	return AsBool(!n.BoolValue())
}

func (n JsonString) Index(index JsonType) (JsonType, error) {
	if IsNumber(index) {
		i := int(AsNumber(index).IntValue())
		if i < 0 {
			i = len(n.StringValue()) + i
		}
		if i >= 0 && i < len(n.StringValue()) {
			return NewString(string(n.StringValue()[i])), nil
		} else {
			return Null, burrito.WrappedErrorf("Index out of bounds: %d", i)
		}
	}
	return Null, burrito.WrappedErrorf("Invalid index: %s", index.StringValue())
}

// AsString converts an interface to a string.
func AsString(obj interface{}) JsonString {
	if obj == nil {
		return EmptyString
	}
	if b, ok := obj.(JsonString); ok {
		return b
	}
	if b, ok := obj.(bool); ok {
		if b {
			return NewString("true")
		}
		return NewString("false")
	}
	if b, ok := obj.(int); ok && b != 0 {
		return NewString(strconv.FormatFloat(float64(b), 'f', -1, 64))
	}
	if b, ok := obj.(float64); ok && b != 0 {
		return NewString(strconv.FormatFloat(b, 'f', -1, 64))
	}
	if b, ok := obj.(float32); ok && b != 0 {
		return NewString(strconv.FormatFloat(float64(b), 'f', -1, 64))
	}
	if b, ok := obj.(string); ok {
		return NewString(b)
	}
	if b, ok := obj.(JsonType); ok {
		return NewString(b.StringValue())
	}
	//TODO: add more types
	return EmptyString
}

// IsString returns true if the given interface is a string.
func IsString(obj interface{}) bool {
	if obj == nil {
		return false
	}
	if _, ok := obj.(string); ok {
		return true
	}
	if _, ok := obj.(JsonString); ok {
		return true
	}
	return false
}

func NewString(value string) JsonString {
	if value == "" {
		return EmptyString
	}
	return JsonString{Value: value}
}
