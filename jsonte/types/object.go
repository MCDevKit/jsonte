package types

import (
	"fmt"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/gammazero/deque"
	"reflect"
	"regexp"
	"strings"
)

type JsonObject struct {
	Value      *utils.NavigableMap[string, JsonType]
	StackValue *deque.Deque[*JsonObject]
}

func (o *JsonObject) Keys() []string {
	if o.Value != nil {
		return o.Value.Keys()
	}
	if o.StackValue != nil {
		keys := make([]string, 0)
		for i := 0; i < o.StackValue.Len(); i++ {
			keys = append(keys, o.StackValue.At(i).Keys()...)
		}
		return keys
	}
	return []string{}
}

func (o *JsonObject) Get(key string) JsonType {
	if o.Value != nil && o.ContainsKey(key) {
		return o.Value.Get(key)
	}
	if o.StackValue != nil {
		for i := o.StackValue.Len() - 1; i >= 0; i-- {
			if o.StackValue.At(i).ContainsKey(key) {
				return o.StackValue.At(i).Get(key)
			}
		}
	}
	return Null
}

func (o *JsonObject) Put(key string, value JsonType) {
	if o.Value != nil {
		o.Value.Put(key, value)
	} else {
		o.StackValue.At(0).Put(key, value)
	}
}

func (o *JsonObject) Remove(key string) {
	if o.Value != nil {
		o.Value.Remove(key)
	}
	if o.StackValue != nil {
		for i := o.StackValue.Len() - 1; i >= 0; i-- {
			if o.StackValue.At(i).ContainsKey(key) {
				o.StackValue.At(i).Remove(key)
			}
		}
	}
}

func (o *JsonObject) ContainsKey(key string) bool {
	if o.Value != nil {
		return o.Value.ContainsKey(key)
	}
	if o.StackValue != nil {
		for i := o.StackValue.Len() - 1; i >= 0; i-- {
			if o.StackValue.At(i).ContainsKey(key) {
				return true
			}
		}
	}
	return false
}

func (o *JsonObject) StringValue() string {
	return ToString(o.Unbox())
}

func (o *JsonObject) BoolValue() bool {
	return !o.IsEmpty()
}

func (o *JsonObject) LessThan(value JsonType) (bool, error) {
	return false, burrito.WrappedErrorf("Objects cannot be compared")
}

func (o *JsonObject) Unbox() interface{} {
	result := utils.NewNavigableMap[string, interface{}]()
	for _, k := range o.Keys() {
		result.Put(k, o.Get(k).Unbox())
	}
	return result
}

func (o *JsonObject) Equals(value JsonType) bool {
	if value == Null {
		return false
	}
	if b, ok := value.(*JsonObject); ok {
		return IsEqualObject(o.Unbox().(utils.NavigableMap[string, JsonType]), b.Unbox().(utils.NavigableMap[string, JsonType]))
	}
	return false
}

func (o *JsonObject) Negate() JsonType {
	return NaN()
}

func (o *JsonObject) Index(i JsonType) (JsonType, error) {
	if b, ok := i.(*JsonString); ok {
		if !o.ContainsKey(b.Value) {
			return Null, burrito.WrappedErrorf("Property '%s' not found in %s", b.StringValue(), o.StringValue())
		}
		return o.Get(b.Value), nil
	}
	if b, ok := i.(*JsonNumber); ok {
		if o.StackValue != nil {
			return Null, burrito.WrappedErrorf("Cannot index a combined object with a number")
		}
		if b.IntValue() < 0 || b.IntValue() >= int32(len(o.Value.Keys())) {
			return Null, burrito.WrappedErrorf("Index out of bounds: %d", b.IntValue())
		}
		return o.Value.Get(o.Value.Keys()[int(b.IntValue())]), nil
	}
	if b, ok := i.(*JsonPath); ok {
		return b.Get(o)
	}
	return Null, burrito.WrappedErrorf("Index must be a string or a number: %s", i.StringValue())
}

func (o *JsonObject) Add(i JsonType) JsonType {
	if IsObject(i) {
		return MergeObject(AsObject(i), o, false, "#")
	}
	if i == nil || i == Null {
		return o
	}
	return NewString(o.StringValue() + i.StringValue())
}

func (o *JsonObject) IsEmpty() bool {
	if o.Value != nil {
		return o.Value.IsEmpty()
	}
	if o.StackValue != nil {
		for i := 0; i < o.StackValue.Len(); i++ {
			if !o.StackValue.At(i).IsEmpty() {
				return false
			}
		}
	}
	return true
}

func (o *JsonObject) Size() int {
	if o.Value != nil {
		return o.Value.Size()
	}
	if o.StackValue != nil {
		size := 0
		for i := 0; i < o.StackValue.Len(); i++ {
			size += o.StackValue.At(i).Size()
		}
		return size
	}
	return 0
}

func (o *JsonObject) Values() []JsonType {
	if o.Value != nil {
		return o.Value.Values()
	}
	if o.StackValue != nil {
		values := make([]JsonType, 0)
		for i := 0; i < o.StackValue.Len(); i++ {
			values = append(values, o.StackValue.At(i).Values()...)
		}
		return values
	}
	return []JsonType{}
}

// AsObject returns the given interface as a JSON object.
func AsObject(obj interface{}) *JsonObject {
	if obj == nil {
		return NewJsonObject()
	}
	if b, ok := obj.(*JsonObject); ok {
		return b
	}
	if _, ok := obj.(JsonType); ok {
		return NewJsonObject()
	}
	if b, ok := obj.(utils.NavigableMap[string, JsonType]); ok {
		return &JsonObject{&b, nil}
	}
	if b, ok := obj.(utils.NavigableMap[string, interface{}]); ok {
		result := NewJsonObject()
		for _, key := range b.Keys() {
			result.Put(key, Box(b.Get(key)))
		}
		return result
	}
	rt := reflect.TypeOf(obj)
	switch rt.Kind() {
	case reflect.Map:
		rv := reflect.ValueOf(obj)
		result := NewJsonObject()
		for _, key := range rv.MapKeys() {
			result.Put(key.String(), Box(rv.MapIndex(key).Interface()))
		}
		return result
	case reflect.Struct:
		rv := reflect.ValueOf(obj)
		result := NewJsonObject()
		for i := 0; i < rv.NumField(); i++ {
			result.Put(rv.Type().Field(i).Name, Box(rv.Field(i).Interface()))
		}
		return result
	}
	return NewJsonObject()
}

// IsObject returns true if the given interface is an object.
func IsObject(obj interface{}) bool {
	if obj == nil {
		return false
	}
	if _, ok := obj.(*JsonObject); ok {
		return true
	}
	if _, ok := obj.(JsonType); ok {
		return false
	}
	if _, ok := obj.(utils.NavigableMap[string, JsonType]); ok {
		return true
	}
	if _, ok := obj.(utils.NavigableMap[string, interface{}]); ok {
		return true
	}
	rt := reflect.TypeOf(obj)
	switch rt.Kind() {
	case reflect.Map, reflect.Struct:
		return true
	}
	return false
}

// IsEqualObject returns true if the given JSON objects are equal.
func IsEqualObject(a, b utils.NavigableMap[string, JsonType]) bool {
	if a.Size() != b.Size() {
		return false
	}
	for _, k := range a.Keys() {
		if !b.ContainsKey(k) {
			return false
		}
		if !b.Get(k).Equals(a.Get(k)) {
			return false
		}
	}
	for _, k := range b.Keys() {
		if !a.ContainsKey(k) {
			return false
		}
	}
	return true
}

// TODO: This should be moved to a shared package.
var actionPattern, _ = regexp.Compile("^\\{\\{(?:\\\\.|[^{}])+}}$")

// MergeObject merges two JSON objects into a new JSON object.
// If the same value, that is not an object or an array exists in both objects, the value from the second object will be used.
func MergeObject(template, parent *JsonObject, keepOverrides bool, path string) *JsonObject {
	result := NewJsonObject()
	for _, k := range template.Keys() {
		v := template.Get(k)
		if IsObject(v) {
			result.Put(k, DeepCopyObject(AsObject(v)))
		} else if IsArray(v) {
			result.Put(k, DeepCopyArray(AsArray(v)))
		} else {
			result.Put(k, v)
		}
	}
	skipKeys := make([]string, 0)
out:
	for _, k := range parent.Keys() {
		v := parent.Get(k)
		for _, key := range skipKeys {
			if key == k {
				continue out
			}
		}
		isReversedMerge := strings.HasPrefix(k, "^")
		k = strings.TrimPrefix(k, "^")
		if strings.HasPrefix(k, "$") && !IsReservedKey(k) {
			if keepOverrides {
				result.Put(k, v)
			} else {
				if IsObject(v) {
					result.Put(strings.TrimPrefix(k, "$"), MergeObject(NewJsonObject(), AsObject(v), keepOverrides, fmt.Sprintf("%s/%s", path, k)))
				} else if IsArray(v) {
					result.Put(strings.TrimPrefix(k, "$"), MergeArray(NewJsonArray(), AsArray(v), keepOverrides, fmt.Sprintf("%s/%s", path, k)))
				} else {
					result.Put(strings.TrimPrefix(k, "$"), v)
				}
			}
			skipKeys = append(skipKeys, strings.TrimPrefix(k, "$"))
		} else if !template.ContainsKey(k) {
			if IsObject(v) {
				merge := MergeObject(NewJsonObject(), AsObject(v), keepOverrides, fmt.Sprintf("%s/%s", path, k))
				result.Put(k, merge)
			} else if IsArray(v) {
				merge := MergeArray(NewJsonArray(), AsArray(v), keepOverrides, fmt.Sprintf("%s/%s", path, k))
				result.Put(k, merge)
			} else {
				result.Put(k, v)
			}
		} else {
			if IsObject(v) && IsObject(result.Get(k)) {
				merge := MergeObject(AsObject(template.Get(k)), AsObject(v), keepOverrides, fmt.Sprintf("%s/%s", path, k))
				result.Put(k, merge)
			} else if IsArray(v) && IsArray(template.Get(k)) {
				var merge, v1 *JsonArray
				if result.ContainsKey(k) {
					v1 = AsArray(result.Get(k))
				} else {
					v1 = AsArray(template.Get(k))
				}
				if isReversedMerge {
					merge = MergeArray(AsArray(v), v1, keepOverrides, fmt.Sprintf("%s/%s", path, k))
				} else {
					merge = MergeArray(v1, AsArray(v), keepOverrides, fmt.Sprintf("%s/%s", path, k))
				}
				result.Put(k, merge)
			} else {
				result.Put(k, v)
			}
		}
	}
	return result
}

// DeepCopyObject creates a deep copy of the given JSON object.
func DeepCopyObject(object *JsonObject) *JsonObject {
	result := NewJsonObject()
	for _, k := range object.Keys() {
		v := object.Get(k)
		if IsObject(v) {
			result.Put(k, DeepCopyObject(AsObject(v)))
		} else if IsArray(v) {
			result.Put(k, DeepCopyArray(AsArray(v)))
		} else {
			result.Put(k, v)
		}
	}
	return result
}

func NewJsonObject() *JsonObject {
	navigableMap := utils.NewNavigableMap[string, JsonType]()
	return &JsonObject{&navigableMap, nil}
}

func IsReservedKey(k string) bool {
	return strings.EqualFold(k, "$comment") || strings.EqualFold(k, "$assert")
}
