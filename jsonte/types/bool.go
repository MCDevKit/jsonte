package types

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"strconv"
)

// JsonBool is a struct that represents a boolean JSON value.
type JsonBool struct {
	Value       bool
	parent      JsonType
	parentIndex JsonType
}

func True() *JsonBool {
	return &JsonBool{Value: true}
}

func False() *JsonBool {
	return &JsonBool{Value: false}
}

func (t *JsonBool) Parent() JsonType {
	return t.parent
}

func (t *JsonBool) ParentIndex() JsonType {
	return t.parentIndex
}

func (t *JsonBool) UpdateParent(parent JsonType, parentIndex JsonType) {
	t.parent = parent
	t.parentIndex = parentIndex
}

// LessThan compares the JsonBool with another JsonType, returning true if the boolean is false and the other value is true.
func (t *JsonBool) LessThan(other JsonType) (bool, error) {
	if other == nil {
		return false, nil
	}
	if IsBool(other) {
		if t.BoolValue() == AsBool(other).BoolValue() || t.BoolValue() {
			return false, nil
		}
		return true, nil
	}
	if IsNumber(other) {
		if t.Value {
			return 1 < AsNumber(other).FloatValue(), nil
		} else {
			return 0 < AsNumber(other).FloatValue(), nil
		}
	}
	return false, incompatibleTypesError(t, other)
}

// BoolValue returns the boolean value of the JsonBool.
func (t *JsonBool) BoolValue() bool {
	return t.Value
}

// StringValue returns the string representation of the JsonBool.
func (t *JsonBool) StringValue() string {
	return strconv.FormatBool(t.BoolValue())
}

// Equals checks if the JsonBool is equal to another JsonType.
func (t *JsonBool) Equals(value JsonType) bool {
	if IsNull(value) {
		return false
	}
	if IsBool(value) {
		return t.BoolValue() == value.BoolValue()
	}
	if IsNumber(value) {
		return t.BoolValue() == value.BoolValue()
	}
	return false
}

// Unbox returns the JsonBool as a native Go bool.
func (t *JsonBool) Unbox() interface{} {
	return t.BoolValue()
}

// Negate returns a new JsonBool with the opposite value.
func (t *JsonBool) Negate() JsonType {
	return NewBool(!t.BoolValue())
}

// Index returns an error since indexing is not supported for booleans.
func (t *JsonBool) Index(i JsonType) (JsonType, error) {
	return Null, burrito.WrappedErrorf("Cannot access %s from a boolean", i.StringValue())
}

// Add performs addition of the JsonBool with another JsonType.
func (t *JsonBool) Add(i JsonType) JsonType {
	if IsNumber(i) {
		return AsNumber(i).Add(t)
	}
	if IsBool(i) {
		result := 0
		if t.BoolValue() {
			result++
		}
		if AsBool(i).BoolValue() {
			result++
		}
		return &JsonNumber{
			Value:   float64(result),
			Decimal: false,
		}
	}
	return NewString(t.StringValue() + i.StringValue())
}

// AsBool converts an interface to a JsonBool.
func AsBool(obj interface{}) *JsonBool {
	if obj == nil {
		return False()
	}
	if b, ok := obj.(bool); ok {
		return NewBool(b)
	}
	if b, ok := obj.(int); ok && b != 0 {
		return True()
	}
	if b, ok := obj.(float64); ok && b != 0 {
		return True()
	}
	if b, ok := obj.(float32); ok && b != 0 {
		return True()
	}
	if b, ok := obj.(JsonType); ok {
		return NewBool(b.BoolValue())
	}
	return NewBool(obj != nil)
}

// IsBool returns true if the given interface is a boolean or a JsonBool.
func IsBool(obj interface{}) bool {
	if obj == nil {
		return false
	}
	if _, ok := obj.(bool); ok {
		return true
	}
	if _, ok := obj.(*JsonBool); ok {
		return true
	}
	return false
}

// NewBool creates a new JsonBool with the specified value.
func NewBool(value bool) *JsonBool {
	return &JsonBool{Value: value}
}
