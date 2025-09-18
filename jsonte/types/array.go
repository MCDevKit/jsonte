package types

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"reflect"
)

type JsonArray struct {
	Value       []JsonType
	parent      JsonType
	parentIndex JsonType
}

func (t *JsonArray) Append(v ...JsonType) *JsonArray {
	t.Value = append(t.Value, v...)
	return t
}

func (t *JsonArray) Prepend(v ...JsonType) *JsonArray {
	t.Value = append(v, t.Value...)
	return t
}

func (t *JsonArray) Remove(i *JsonNumber) (JsonType, error) {
	if len(t.Value) == 0 {
		return Null, burrito.WrappedErrorf("Cannot remove from empty array")
	}
	index := int(i.IntValue())
	if index < 0 {
		index = len(t.Value) + index
	}
	if index >= 0 && index < len(t.Value) {
		t.Value = append(t.Value[:index], t.Value[index+1:]...)
	}
	return t, nil
}

func (t *JsonArray) RemoveIf(i *JsonLambda) (JsonType, error) {
	result := make([]JsonType, 0)
	for index, v := range t.Value {
		if b, err := i.Call(v, AsNumber(index)); err != nil {
			return Null, err
		} else if !b.BoolValue() {
			result = append(result, v)
		}
	}
	t.Value = result
	return t, nil
}

func (t *JsonArray) RemoveFront() (JsonType, error) {
	if len(t.Value) == 0 {
		return Null, burrito.WrappedErrorf("Cannot remove from empty array")
	}
	t.Value = t.Value[1:]
	return t, nil
}

func (t *JsonArray) RemoveBack() (JsonType, error) {
	if len(t.Value) == 0 {
		return Null, burrito.WrappedErrorf("Cannot remove from empty array")
	}
	t.Value = t.Value[:len(t.Value)-1]
	return t, nil
}

func (t *JsonArray) Parent() JsonType {
	return t.parent
}

func (t *JsonArray) ParentIndex() JsonType {
	return t.parentIndex
}

func (t *JsonArray) UpdateParent(parent JsonType, parentIndex JsonType) {
	t.parent = parent
	t.parentIndex = parentIndex
}

func (t *JsonArray) StringValue() string {
	return ToString(t.Unbox())
}

func (t *JsonArray) BoolValue() bool {
	return t.Value != nil && len(t.Value) > 0
}

func (t *JsonArray) Equals(value JsonType) bool {
	if IsNull(value) {
		return false
	}
	if b, ok := value.(*JsonArray); ok {
		return IsEqualArray(t.Value, b.Value)
	}
	return false
}

func (t *JsonArray) Unbox() interface{} {
	result := make([]interface{}, len(t.Value))
	for i, k := range t.Value {
		result[i] = k.Unbox()
	}
	return result
}

func (t *JsonArray) Negate() JsonType {
	// TODO: This should be removed, because `-array` and `array * -1` both should work and currently they don't.
	result := make([]JsonType, len(t.Value))
	for i, v := range t.Value {
		result[i] = v.Negate()
	}
	return &JsonArray{Value: result}
}

func (t *JsonArray) Index(i JsonType) (JsonType, error) {
	if b, ok := i.(*JsonNumber); ok {
		index := int(b.IntValue())
		if index < 0 {
			index = len(t.Value) + index
		}
		if index >= 0 && index < len(t.Value) {
			return t.Value[index], nil
		} else {
			return Null, burrito.WrappedErrorf("Index out of bounds: %d", index)
		}
	}
	if b, ok := i.(*JsonPath); ok {
		return b.Get(t)
	}
	return Null, burrito.WrappedErrorf("Index must be a number: %s", i.StringValue())
}

func (t *JsonArray) Add(i JsonType) JsonType {
	if IsArray(i) {
		return MergeArray(t, AsArray(i), false, "#")
	}
	if i == nil || IsNull(i) {
		return t
	}
	return NewString(t.StringValue() + i.StringValue())
}

func (t *JsonArray) LessThan(other JsonType) (bool, error) {
	return false, burrito.WrappedErrorf("Arrays cannot be compared")
}

// IsEqualArray returns true if the given JSON arrays are equal.
func IsEqualArray(a, b []JsonType) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if !b[k].Equals(v) {
			return false
		}
	}
	return true
}

// IsArray returns true if the given interface is an array.
func IsArray(obj interface{}) bool {
	if obj == nil {
		return false
	}
	if _, ok := obj.(*JsonArray); ok {
		return true
	}
	rt := reflect.TypeOf(obj)
	switch rt.Kind() {
	case reflect.Slice, reflect.Array:
		return true
	}
	return false
}

// AsArray returns the given interface as a JSON array.
func AsArray(obj interface{}) *JsonArray {
	if obj == nil {
		return NewJsonArray()
	}
	if b, ok := obj.(*JsonArray); ok {
		return b
	}
	rt := reflect.TypeOf(obj)
	switch rt.Kind() {
	case reflect.Slice, reflect.Array:
		rv := reflect.ValueOf(obj)
		result := make([]JsonType, rv.Len())
		for i := 0; i < rv.Len(); i++ {
			result[i] = Box(rv.Index(i).Interface())
		}
		return &JsonArray{Value: result}
	}
	return NewJsonArray()
}

// MergeArray merges two JSON arrays into a new JSON array.
func MergeArray(template, parent *JsonArray, keepOverrides bool, path string) *JsonArray {
	if template == nil {
		template = NewJsonArray()
	}
	if parent == nil {
		parent = NewJsonArray()
	}
	templateLen := len(template.Value)
	result := &JsonArray{Value: make([]JsonType, 0, templateLen+len(parent.Value))}
	for _, v := range template.Value {
		switch typed := v.(type) {
		case *JsonObject:
			result.Value = append(result.Value, DeepCopyObject(typed))
		case *JsonArray:
			result.Value = append(result.Value, DeepCopyArray(typed))
		default:
			result.Value = append(result.Value, v)
		}
	}
	for i, v := range parent.Value {
		childPath := joinArrayPath(path, i)
		switch typed := v.(type) {
		case *JsonObject:
			result.Value = append(result.Value, MergeObject(NewJsonObjectWithCapacity(typed.Size()), typed, keepOverrides, childPath))
		case *JsonArray:
			childCapacity := 0
			if typed != nil {
				childCapacity = len(typed.Value)
			}
			result.Value = append(result.Value, MergeArray(NewJsonArrayWithCapacity(childCapacity), typed, keepOverrides, childPath))
		default:
			result.Value = append(result.Value, v)
		}
	}
	return result
}

// DeepCopyArray creates a deep copy of the given JSON array.
func DeepCopyArray(object *JsonArray) *JsonArray {
	if object == nil {
		return NewJsonArray()
	}
	result := NewJsonArrayWithCapacity(len(object.Value))
	for _, v := range object.Value {
		switch typed := v.(type) {
		case *JsonObject:
			if typed == nil {
				result.Value = append(result.Value, Null)
				continue
			}
			result.Value = append(result.Value, DeepCopyObject(typed))
		case *JsonArray:
			if typed == nil {
				result.Value = append(result.Value, Null)
				continue
			}
			result.Value = append(result.Value, DeepCopyArray(typed))
		default:
			result.Value = append(result.Value, v)
		}
	}
	return result
}

func NewJsonArray() *JsonArray {
	return NewJsonArrayWithCapacity(0)
}

func NewJsonArrayWithCapacity(capacity int) *JsonArray {
	if capacity < 0 {
		capacity = 0
	}
	return &JsonArray{Value: make([]JsonType, 0, capacity)}
}
