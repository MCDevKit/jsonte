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

func forEachEntry(obj *JsonObject, fn func(string, JsonType)) {
	if obj == nil || fn == nil {
		return
	}
	if obj.Value != nil && obj.StackValue == nil {
		obj.Value.ForEach(fn)
		return
	}
	for _, key := range obj.Keys() {
		fn(key, obj.Get(key))
	}
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
	if template == nil {
		template = NewJsonObject()
	}
	if parent == nil {
		parent = NewJsonObject()
	}
	templateCapacity := 0
	if template.Value != nil {
		templateCapacity = template.Value.Size()
	}
	parentCapacity := 0
	if parent.Value != nil {
		parentCapacity = parent.Value.Size()
	}
	result := NewJsonObjectWithCapacity(templateCapacity + parentCapacity)
	forEachEntry(template, func(key string, value JsonType) {
		switch typed := value.(type) {
		case *JsonObject:
			result.Put(key, DeepCopyObject(typed))
		case *JsonArray:
			result.Put(key, DeepCopyArray(typed))
		default:
			result.Put(key, value)
		}
	})
	var skipKeys map[string]struct{}
	forEachEntry(parent, func(rawKey string, value JsonType) {
		if rawKey == "" {
			return
		}
		if skipKeys != nil {
			if _, ok := skipKeys[rawKey]; ok {
				return
			}
		}
		isReversedMerge := rawKey[0] == '^'
		key := rawKey
		if isReversedMerge {
			if len(key) == 1 {
				return
			}
			key = rawKey[1:]
		}
		if len(key) == 0 {
			return
		}
		if key[0] == '$' && !IsReservedKey(key) {
			trimmedKey := key[1:]
			if keepOverrides {
				result.Put(key, value)
			} else {
				childPath := ""
				switch typed := value.(type) {
				case *JsonObject:
					childCapacity := 0
					if typed != nil {
						childCapacity = typed.Size()
					}
					if typed != nil {
						childPath = joinObjectPath(path, key)
					}
					result.Put(trimmedKey, MergeObject(NewJsonObjectWithCapacity(childCapacity), typed, keepOverrides, childPath))
				case *JsonArray:
					childCapacity := 0
					if typed != nil {
						childCapacity = len(typed.Value)
					}
					if typed != nil {
						childPath = joinObjectPath(path, key)
					}
					result.Put(trimmedKey, MergeArray(NewJsonArrayWithCapacity(childCapacity), typed, keepOverrides, childPath))
				default:
					result.Put(trimmedKey, value)
				}
				if result.ContainsKey(key) {
					result.Remove(key)
				}
			}
			if skipKeys == nil {
				skipKeys = make(map[string]struct{}, parentCapacity)
			}
			skipKeys[trimmedKey] = struct{}{}
			return
		}
		baseValue, exists := result.TryGet(key)
		if !exists {
			childPath := ""
			switch typed := value.(type) {
			case *JsonObject:
				childCapacity := 0
				if typed != nil {
					childCapacity = typed.Size()
				}
				if typed != nil {
					childPath = joinObjectPath(path, key)
				}
				result.Put(key, MergeObject(NewJsonObjectWithCapacity(childCapacity), typed, keepOverrides, childPath))
			case *JsonArray:
				childCapacity := 0
				if typed != nil {
					childCapacity = len(typed.Value)
				}
				if typed != nil {
					childPath = joinObjectPath(path, key)
				}
				result.Put(key, MergeArray(NewJsonArrayWithCapacity(childCapacity), typed, keepOverrides, childPath))
			default:
				result.Put(key, value)
			}
			return
		}
		switch typed := value.(type) {
		case *JsonObject:
			if templateObj, ok := baseValue.(*JsonObject); ok {
				childPath := joinObjectPath(path, key)
				result.Put(key, MergeObject(templateObj, typed, keepOverrides, childPath))
			} else {
				result.Put(key, value)
			}
		case *JsonArray:
			if _, ok := baseValue.(*JsonArray); ok {
				base := AsArray(result.Get(key))
				childPath := joinObjectPath(path, key)
				if isReversedMerge {
					result.Put(key, MergeArray(typed, base, keepOverrides, childPath))
				} else {
					result.Put(key, MergeArray(base, typed, keepOverrides, childPath))
				}
			} else {
				result.Put(key, value)
			}
		default:
			result.Put(key, value)
		}
	})
	return result
}

// DeepCopyObject creates a deep copy of the given JSON object.
func DeepCopyObject(object *JsonObject) *JsonObject {
	if object == nil {
		return NewJsonObject()
	}
	result := NewJsonObjectWithCapacity(object.Size())
	for _, k := range object.Keys() {
		v := object.Get(k)
		switch typed := v.(type) {
		case *JsonObject:
			result.Put(k, DeepCopyObject(typed))
		case *JsonArray:
			result.Put(k, DeepCopyArray(typed))
		default:
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
