package types

import (
	"fmt"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"reflect"
	"regexp"
	"strings"
)

type JsonObject struct {
	Value *utils.NavigableMap[string, JsonType]
}

func (o JsonObject) Keys() []string {
	if o.Value == nil {
		return []string{}
	}
	return o.Value.Keys()
}

func (o JsonObject) Get(key string) JsonType {
	if o.Value == nil || !o.ContainsKey(key) {
		return Null
	}
	return o.Value.Get(key)
}

func (o JsonObject) Put(key string, value JsonType) {
	o.Value.Put(key, value)
}

func (o JsonObject) Remove(key string) {
	o.Value.Remove(key)
}

func (o JsonObject) ContainsKey(key string) bool {
	if o.Value == nil {
		return false
	}
	return o.Value.ContainsKey(key)
}

func (o JsonObject) StringValue() string {
	return ToString(o.Unbox())
}

func (o JsonObject) BoolValue() bool {
	return o.Value != nil && !o.Value.IsEmpty()
}

func (o JsonObject) IsNull() bool {
	return false
}

func (o JsonObject) LessThan(value JsonType) (bool, error) {
	return false, burrito.WrappedErrorf("Objects cannot be compared")
}

func (o JsonObject) Unbox() interface{} {
	result := utils.NewNavigableMap[string, interface{}]()
	for _, k := range o.Keys() {
		result.Put(k, o.Get(k).Unbox())
	}
	return result
}

func (o JsonObject) Equals(value JsonType) bool {
	if value == Null {
		return false
	}
	if b, ok := value.(JsonType); ok && b.IsNull() {
		return false
	}
	if b, ok := value.(JsonObject); ok {
		return IsEqualObject(*o.Value, *b.Value)
	}
	return false
}

func (o JsonObject) Negate() JsonType {
	return NaN
}

func (o JsonObject) Index(i JsonType) (JsonType, error) {
	if b, ok := i.(JsonString); ok {
		if !o.ContainsKey(b.Value) {
			return Null, burrito.WrappedErrorf("Property '%s' not found in %s", b.StringValue(), o.StringValue())
		}
		return o.Get(b.Value), nil
	}
	if b, ok := i.(JsonNumber); ok {
		if b.IntValue() < 0 || b.IntValue() >= int32(len(o.Value.Keys())) {
			return Null, burrito.WrappedErrorf("Index out of bounds: %d", b.IntValue())
		}
		return o.Value.Get(o.Value.Keys()[int(b.IntValue())]), nil
	}
	return Null, burrito.WrappedErrorf("Index must be a string or a number: %s", i.StringValue())
}

func (o JsonObject) Add(i JsonType) JsonType {
	if IsObject(i) {
		return MergeObject(AsObject(i), o, false, false, "#")
	}
	if i == nil || i == Null {
		return o
	}
	return NewString(o.StringValue() + i.StringValue())
}

func (o JsonObject) IsEmpty() bool {
	return o.Value.IsEmpty()
}

func (o JsonObject) Size() int {
	return o.Value.Size()
}

func (o JsonObject) Values() []JsonType {
	return o.Value.Values()
}

// AsObject returns the given interface as a JSON object.
func AsObject(obj interface{}) JsonObject {
	if obj == nil {
		return NewJsonObject()
	}
	if b, ok := obj.(JsonObject); ok {
		return b
	}
	if _, ok := obj.(JsonType); ok {
		return NewJsonObject()
	}
	if b, ok := obj.(utils.NavigableMap[string, JsonType]); ok {
		return JsonObject{&b}
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
	if _, ok := obj.(JsonObject); ok {
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
		if !b.Get(k).Equals(a.Get(k)) {
			return false
		}
	}
	return true
}

// TODO: This should be moved to a shared package.
var actionPattern, _ = regexp.Compile("^\\{\\{(?:\\\\.|[^{}])+}}$")

// MergeObject merges two JSON objects into a new JSON object.
// If the same value, that is not an object or an array exists in both objects, the value from the second object will be used.
func MergeObject(template, parent JsonObject, keepOverrides, insideTemplate bool, path string) JsonObject {
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
		if strings.HasPrefix(k, "$") && !isReservedKey(k) {
			if insideTemplate {
				utils.Logger.Warnf("Overriding inside templated object is not supported. Unexpected behavior may occur at %s/%s", path, k)
			}
			if keepOverrides {
				result.Put(k, v)
			} else {
				result.Put(strings.TrimPrefix(k, "$"), v)
				skipKeys = append(skipKeys, strings.TrimPrefix(k, "$"))
			}
		} else if !template.ContainsKey(k) {
			if IsObject(v) {
				merge := MergeObject(NewJsonObject(), AsObject(v), keepOverrides, insideTemplate || actionPattern.MatchString(k), fmt.Sprintf("%s/%s", path, k))
				result.Put(k, merge)
			} else if IsArray(v) {
				merge := MergeArray(NewJsonArray(), AsArray(v), keepOverrides, insideTemplate || actionPattern.MatchString(k), fmt.Sprintf("%s/%s", path, k))
				result.Put(k, merge)
			} else {
				result.Put(k, v)
			}
		} else {
			if IsObject(v) && IsObject(result.Get(k)) {
				merge := MergeObject(AsObject(template.Get(k)), AsObject(v), keepOverrides, insideTemplate || actionPattern.MatchString(k), fmt.Sprintf("%s/%s", path, k))
				result.Put(k, merge)
			} else if IsArray(v) && IsArray(template.Get(k)) {
				var merge, v1 JsonArray
				if result.ContainsKey(k) {
					v1 = AsArray(result.Get(k))
				} else {
					v1 = AsArray(template.Get(k))
				}
				if isReversedMerge {
					merge = MergeArray(AsArray(v), v1, keepOverrides, insideTemplate || actionPattern.MatchString(k), fmt.Sprintf("%s/%s", path, k))
				} else {
					merge = MergeArray(v1, AsArray(v), keepOverrides, insideTemplate || actionPattern.MatchString(k), fmt.Sprintf("%s/%s", path, k))
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
func DeepCopyObject(object JsonObject) JsonObject {
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

func NewJsonObject() JsonObject {
	navigableMap := utils.NewNavigableMap[string, JsonType]()
	return JsonObject{&navigableMap}
}

func isReservedKey(k string) bool {
	return strings.EqualFold(k, "$comment") || strings.EqualFold(k, "$assert")
}
