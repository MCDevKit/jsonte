package types

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"strconv"
	"strings"
)

// JsonString is a struct that represents a number, that can be either integer or decimal.
type JsonString struct {
	Value       string
	parent      JsonType
	parentIndex JsonType
}

func (t *JsonString) Parent() JsonType {
	return t.parent
}

func (t *JsonString) ParentIndex() JsonType {
	return t.parentIndex
}

func (t *JsonString) UpdateParent(parent JsonType, parentIndex JsonType) {
	t.parent = parent
	t.parentIndex = parentIndex
}

func (t *JsonString) LessThan(other JsonType) (bool, error) {
	if other == nil || other == Null {
		return false, nil
	}
	if IsString(other) {
		compare := strings.Compare(t.StringValue(), AsString(other).StringValue())
		return compare < 0, nil
	}
	return false, incompatibleTypesError(t, other)
}

func (t *JsonString) BoolValue() bool {
	return strings.Trim(t.StringValue(), "\n\r") != ""
}

func (t *JsonString) StringValue() string {
	return t.Value
}

func (t *JsonString) Equals(value JsonType) bool {
	if value == Null {
		return false
	}
	if IsString(value) {
		return t.StringValue() == value.StringValue()
	}
	return false
}

func (t *JsonString) Unbox() interface{} {
	return t.StringValue()
}

func (t *JsonString) Add(i JsonType) JsonType {
	return NewString(t.StringValue() + i.StringValue())
}

func (t *JsonString) Negate() JsonType {
	return AsBool(!t.BoolValue())
}

func (t *JsonString) Index(index JsonType) (JsonType, error) {
	if IsNumber(index) {
		i := int(AsNumber(index).IntValue())
		runes := []rune(t.StringValue())
		if i < 0 {
			i = len(runes) + i
		}
		if i >= 0 && i < len(runes) {
			return NewString(string(runes[i])), nil
		} else {
			return Null, burrito.WrappedErrorf("Index out of bounds: %d", i)
		}
	}
	return Null, burrito.WrappedErrorf("Invalid index: %s", index.StringValue())
}

// AsString converts an interface to a string.
func AsString(obj interface{}) *JsonString {
	if obj == nil {
		return NewString("")
	}
	if b, ok := obj.(*JsonString); ok {
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
	return NewString("")
}

// IsString returns true if the given interface is a string.
func IsString(obj interface{}) bool {
	if obj == nil {
		return false
	}
	if _, ok := obj.(string); ok {
		return true
	}
	if _, ok := obj.(*JsonString); ok {
		return true
	}
	return false
}

func NewString(value string) *JsonString {
	return &JsonString{Value: value}
}
