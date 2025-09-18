package types

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/gammazero/deque"
	"reflect"
	"strings"
)

type JsonObject struct {
	Value       *utils.NavigableMap[string, JsonType]
	StackValue  *deque.Deque[*JsonObject]
	StackTarget *JsonObject
	parent      JsonType
	parentIndex JsonType
}

func (t *JsonObject) Keys() []string {
	if t.Value != nil {
		return t.Value.Keys()
	}
	if t.StackValue != nil {
		keysMap := make(map[string]bool)
		for _, k := range t.StackTarget.Keys() {
			keysMap[k] = true
		}
		for i := 0; i < t.StackValue.Len(); i++ {
			for _, k := range t.StackValue.At(i).Keys() {
				keysMap[k] = true
			}
		}
		keys := make([]string, 0)
		for k := range keysMap {
			keys = append(keys, k)
		}
		return keys
	}
	return []string{}
}

func (t *JsonObject) Get(key string) JsonType {
	if val, ok := t.TryGet(key); ok {
		return val
	}
	return Null
}

func (t *JsonObject) TryGet(key string) (JsonType, bool) {
	if t.Value != nil {
		return t.Value.TryGet(key)
	}
	if t.StackValue != nil {
		if t.StackTarget.ContainsKey(key) {
			return t.StackTarget.Get(key), true
		}
		for i := t.StackValue.Len() - 1; i >= 0; i-- {
			if t.StackValue.At(i).ContainsKey(key) {
				return t.StackValue.At(i).Get(key), true
			}
		}
	}
	return Null, false
}

func (t *JsonObject) Put(key string, value JsonType) {
	if t.Value != nil {
		t.Value.Put(key, value)
	} else {
		t.StackTarget.Put(key, value)
	}
}

func (t *JsonObject) Remove(key string) {
	if t.Value != nil {
		t.Value.Remove(key)
	}
	if t.StackValue != nil {
		if t.StackTarget.ContainsKey(key) {
			t.StackTarget.Remove(key)
		}
		for i := t.StackValue.Len() - 1; i >= 0; i-- {
			if t.StackValue.At(i).ContainsKey(key) {
				t.StackValue.At(i).Remove(key)
			}
		}
	}
}

func (t *JsonObject) ContainsKey(key string) bool {
	if t.Value != nil {
		return t.Value.ContainsKey(key)
	}
	if t.StackValue != nil {
		if t.StackTarget.ContainsKey(key) {
			return true
		}
		for i := t.StackValue.Len() - 1; i >= 0; i-- {
			if t.StackValue.At(i).ContainsKey(key) {
				return true
			}
		}
	}
	return false
}

func (t *JsonObject) Parent() JsonType {
	return t.parent
}

func (t *JsonObject) ParentIndex() JsonType {
	return t.parentIndex
}

func (t *JsonObject) UpdateParent(parent JsonType, parentIndex JsonType) {
	t.parent = parent
	t.parentIndex = parentIndex
}

func (t *JsonObject) StringValue() string {
	return ToString(t.Unbox())
}

func (t *JsonObject) BoolValue() bool {
	return !t.IsEmpty()
}

func (t *JsonObject) LessThan(value JsonType) (bool, error) {
	return false, burrito.WrappedErrorf("Objects cannot be compared")
}

func (t *JsonObject) Unbox() interface{} {
	result := utils.NewNavigableMap[string, interface{}]()
	for _, k := range t.Keys() {
		result.Put(k, t.Get(k).Unbox())
	}
	return result
}

func (t *JsonObject) Equals(value JsonType) bool {
	if IsNull(value) {
		return false
	}
	if b, ok := value.(*JsonObject); ok {
		return IsEqualObject(*AsObject(t.Unbox()).Value, *AsObject(b.Unbox()).Value)
	}
	return false
}

func (t *JsonObject) Negate() JsonType {
	return NaN()
}

func (t *JsonObject) Index(i JsonType) (JsonType, error) {
	if b, ok := i.(*JsonString); ok {
		if !t.ContainsKey(b.Value) {
			return Null, burrito.WrappedErrorf("Property '%s' not found in %s", b.StringValue(), t.StringValue())
		}
		return t.Get(b.Value), nil
	}
	if b, ok := i.(*JsonNumber); ok {
		if t.StackValue != nil {
			return Null, burrito.WrappedErrorf("Cannot index a stack based object with a number")
		}
		if b.IntValue() < 0 || b.IntValue() >= int32(len(t.Value.Keys())) {
			return Null, burrito.WrappedErrorf("Index out of bounds: %d", b.IntValue())
		}
		return t.Value.Get(t.Value.Keys()[int(b.IntValue())]), nil
	}
	if b, ok := i.(*JsonPath); ok {
		return b.Get(t)
	}
	return Null, burrito.WrappedErrorf("Index must be a string or a number: %s", i.StringValue())
}

func (t *JsonObject) Add(i JsonType) JsonType {
	if IsObject(i) {
		return MergeObject(AsObject(i), t, false, "#")
	}
	if i == nil || IsNull(i) {
		return t
	}
	return NewString(t.StringValue() + i.StringValue())
}

func (t *JsonObject) IsEmpty() bool {
	if t.Value != nil {
		return t.Value.IsEmpty()
	}
	if t.StackValue != nil {
		if !t.StackTarget.IsEmpty() {
			return false
		}
		for i := 0; i < t.StackValue.Len(); i++ {
			if !t.StackValue.At(i).IsEmpty() {
				return false
			}
		}
	}
	return true
}

func (t *JsonObject) Size() int {
	if t.Value != nil {
		return t.Value.Size()
	}
	if t.StackValue != nil {
		return len(t.Keys())
	}
	return 0
}

func (t *JsonObject) Values() []JsonType {
	if t.Value != nil {
		return t.Value.Values()
	}
	if t.StackValue != nil {
		values := make([]JsonType, 0)
		keys := t.Keys()
		for _, k := range keys {
			values = append(values, t.Get(k))
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
		return &JsonObject{&b, nil, nil, nil, nil}
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
	skipKeys := make(map[string]struct{})
out:
	for _, k := range parent.Keys() {
		v := parent.Get(k)
		if _, ok := skipKeys[k]; ok {
			continue out
		}
		isReversedMerge := strings.HasPrefix(k, "^")
		k = strings.TrimPrefix(k, "^")
		if strings.HasPrefix(k, "$") && !IsReservedKey(k) {
			trimmedKey := strings.TrimPrefix(k, "$")
			if keepOverrides {
				result.Put(k, v)
			} else {
				if IsObject(v) {
					result.Put(trimmedKey, MergeObject(NewJsonObject(), AsObject(v), keepOverrides, joinObjectPath(path, k)))
				} else if IsArray(v) {
					result.Put(trimmedKey, MergeArray(NewJsonArray(), AsArray(v), keepOverrides, joinObjectPath(path, k)))
				} else {
					result.Put(trimmedKey, v)
				}
			}
			skipKeys[trimmedKey] = struct{}{}
		} else if !template.ContainsKey(k) {
			if IsObject(v) {
				merge := MergeObject(NewJsonObject(), AsObject(v), keepOverrides, joinObjectPath(path, k))
				result.Put(k, merge)
			} else if IsArray(v) {
				merge := MergeArray(NewJsonArray(), AsArray(v), keepOverrides, joinObjectPath(path, k))
				result.Put(k, merge)
			} else {
				result.Put(k, v)
			}
		} else {
			if IsObject(v) && IsObject(result.Get(k)) {
				merge := MergeObject(AsObject(template.Get(k)), AsObject(v), keepOverrides, joinObjectPath(path, k))
				result.Put(k, merge)
			} else if IsArray(v) && IsArray(template.Get(k)) {
				var merge, v1 *JsonArray
				if result.ContainsKey(k) {
					v1 = AsArray(result.Get(k))
				} else {
					v1 = AsArray(template.Get(k))
				}
				if isReversedMerge {
					merge = MergeArray(AsArray(v), v1, keepOverrides, joinObjectPath(path, k))
				} else {
					merge = MergeArray(v1, AsArray(v), keepOverrides, joinObjectPath(path, k))
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
	return &JsonObject{&navigableMap, nil, nil, nil, nil}
}

// NewJsonObjectWithCapacity creates a JsonObject preallocating internal storage.
func NewJsonObjectWithCapacity(capacity int) *JsonObject {
	navigableMap := utils.NewNavigableMapWithCapacity[string, JsonType](capacity)
	return &JsonObject{&navigableMap, nil, nil, nil, nil}
}

func IsReservedKey(k string) bool {
	return strings.EqualFold(k, "$comment") || strings.EqualFold(k, "$assert")
}
