package types

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"reflect"
)

type JsonArray struct {
	JsonType
	Value []JsonType
}

func (o JsonArray) StringValue() string {
	return ToString(o.Unbox())
}

func (o JsonArray) BoolValue() bool {
	return o.Value != nil && len(o.Value) > 0
}

func (o JsonArray) Equals(value JsonType) bool {
	if value.IsNull() {
		return false
	}
	if b, ok := value.(JsonArray); ok {
		return IsEqualArray(o.Value, b.Value)
	}
	return false
}

func (o JsonArray) Unbox() interface{} {
	result := make([]interface{}, len(o.Value))
	for i, k := range o.Value {
		result[i] = k.Unbox()
	}
	return result
}

func (o JsonArray) Negate() JsonType {
	result := make([]JsonType, len(o.Value))
	for i, v := range o.Value {
		result[i] = v.Negate()
	}
	return JsonArray{Value: result}
}

func (o JsonArray) Index(i JsonType) (JsonType, error) {
	if b, ok := i.(JsonNumber); ok {
		index := int(b.IntValue())
		if index < 0 {
			index = len(o.Value) + index
		}
		if index >= 0 && index < len(o.Value) {
			return o.Value[index], nil
		} else {
			return Null, burrito.WrappedErrorf("Index out of bounds: %d", index)
		}
	}
	return Null, burrito.WrappedErrorf("Index must be a number: %s", i.StringValue())
}

func (o JsonArray) Add(i JsonType) JsonType {
	if IsArray(i) {
		return MergeArray(o, AsArray(i), false)
	}
	if i == nil || i == Null {
		return o
	}
	return NewString(o.StringValue() + i.StringValue())
}

func (o JsonArray) IsNull() bool {
	return false
}

func (o JsonArray) LessThan(other JsonType) (bool, error) {
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
	if _, ok := obj.(JsonArray); ok {
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
func AsArray(obj interface{}) JsonArray {
	if obj == nil {
		return NewJsonArray()
	}
	if b, ok := obj.(JsonArray); ok {
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
		return JsonArray{Value: result}
	}
	return NewJsonArray()
}

// MergeArray merges two JSON arrays into a new JSON array.
func MergeArray(template, parent JsonArray, keepOverrides bool) JsonArray {
	var result = NewJsonArray()
	for _, v := range template.Value {
		if IsObject(v) {
			merge := MergeObject(NewJsonObject(), AsObject(v), keepOverrides)
			result.Value = append(result.Value, merge)
		} else if IsArray(v) {
			merge := MergeArray(NewJsonArray(), AsArray(v), keepOverrides)
			result.Value = append(result.Value, merge)
		} else {
			result.Value = append(result.Value, v)
		}
	}
	for _, v := range parent.Value {
		if IsObject(v) {
			merge := MergeObject(NewJsonObject(), AsObject(v), keepOverrides)
			result.Value = append(result.Value, merge)
		} else if IsArray(v) {
			merge := MergeArray(NewJsonArray(), AsArray(v), keepOverrides)
			result.Value = append(result.Value, merge)
		} else {
			result.Value = append(result.Value, v)
		}
	}
	return result
}

// DeepCopyArray creates a deep copy of the given JSON array.
func DeepCopyArray(object JsonArray) JsonArray {
	var result = NewJsonArray()
	for _, v := range object.Value {
		if IsObject(v) {
			result.Value = append(result.Value, DeepCopyObject(AsObject(v)))
		} else if IsArray(v) {
			result.Value = append(result.Value, DeepCopyArray(AsArray(v)))
		} else {
			result.Value = append(result.Value, v)
		}
	}
	return result
}

func NewJsonArray() JsonArray {
	return JsonArray{Value: make([]JsonType, 0)}
}
