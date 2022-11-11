package types

import (
	"fmt"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/json"
	"reflect"
	"strconv"
)

type JsonType interface {
	// StringValue returns the value as a string.
	StringValue() string
	// BoolValue returns the number as a boolean.
	BoolValue() bool
	// Equals returns true if the value is equal to the given value.
	Equals(value JsonType) bool
	// IsNull returns true if the value is null.
	IsNull() bool
	// LessThan returns true if the value is less than the given value.
	LessThan(value JsonType) (bool, error)
	// Unbox returns the value without any containers.
	Unbox() interface{}
	// Negate returns the negated value.
	Negate() JsonType
	// Index returns the value at the given index.
	Index(index JsonType) (JsonType, error)
	// Add returns the sum of the value and the given value.
	Add(index JsonType) JsonType
}

type TypeDescriptor struct {
	Type   reflect.Type
	Name   string
	IsType func(interface{}) bool
	AsType func(interface{}) JsonType
}

var TypeDescriptors []TypeDescriptor

var NaN = JsonString{Value: "NaN"}

func Init() {
	TypeDescriptors = []TypeDescriptor{
		{
			Type:   reflect.TypeOf(Null),
			Name:   "null",
			IsType: func(i interface{}) bool { return i == nil || i == Null },
			AsType: func(i interface{}) JsonType { return Null },
		},
		{
			Type:   reflect.TypeOf(JsonBool{}),
			Name:   "boolean",
			IsType: IsBool,
			AsType: func(i interface{}) JsonType { return AsBool(i) },
		},
		{
			Type:   reflect.TypeOf(JsonString{}),
			Name:   "string",
			IsType: IsString,
			AsType: func(i interface{}) JsonType { return AsString(i) },
		},
		{
			Type:   reflect.TypeOf(JsonNumber{}),
			Name:   "number",
			IsType: IsNumber,
			AsType: func(i interface{}) JsonType { return AsNumber(i) },
		},
		{
			Type:   reflect.TypeOf(Semver{}),
			Name:   "semver",
			IsType: IsSemver,
			AsType: func(i interface{}) JsonType { return AsSemver(i) },
		},
		{
			Type:   reflect.TypeOf(JsonArray{}),
			Name:   "array",
			IsType: IsArray,
			AsType: func(i interface{}) JsonType { return AsArray(i) },
		},
		{
			Type:   reflect.TypeOf(JsonObject{}),
			Name:   "object",
			IsType: IsObject,
			AsType: func(i interface{}) JsonType { return AsObject(i) },
		},
		{
			Type:   reflect.TypeOf(JsonLambda{}),
			Name:   "lambda",
			IsType: func(i interface{}) bool { _, ok := i.(JsonLambda); return ok },
			AsType: func(i interface{}) JsonType {
				if b, ok := i.(JsonLambda); ok {
					return b
				}
				panic("Not a lambda")
			},
		},
	}
}

func TypeName(obj interface{}) string {
	if obj == nil || obj == Null {
		return "nil"
	}
	if IsArray(obj) {
		return "array"
	}
	if IsObject(obj) {
		return "object"
	}
	if IsBool(obj) {
		return "boolean"
	}
	if IsNumber(obj) {
		return "number"
	}
	if IsSemver(obj) {
		return "semver"
	}
	if IsString(obj) {
		return "string"
	}
	if _, ok := obj.(JsonLambda); ok {
		return "lambda"
	}
	return fmt.Sprintf("native<%s>", reflect.TypeOf(obj).String())
}

func Box(obj interface{}) JsonType {
	for _, descriptor := range TypeDescriptors {
		if descriptor.IsType(obj) {
			return descriptor.AsType(obj)
		}
	}
	panic("Unknown type: " + TypeName(obj))
}

func MergeJSON(template, parent JsonType, keepOverrides bool) (JsonType, error) {
	if template == nil || template == Null {
		return parent, nil
	}
	if parent == nil || parent == Null {
		return template, nil
	}
	if IsObject(template) && IsObject(parent) {
		templateMap := AsObject(template)
		parentMap := AsObject(parent)
		return MergeObject(templateMap, parentMap, keepOverrides), nil
	}
	if IsArray(template) && IsArray(parent) {
		templateArray := AsArray(template)
		parentArray := AsArray(parent)
		return MergeArray(templateArray, parentArray, keepOverrides), nil
	}
	if IsObject(template) != IsObject(parent) {
		return nil, burrito.WrappedErrorf("Cannot merge %s and %s", TypeName(template), TypeName(parent))
	}
	if IsArray(template) != IsArray(parent) {
		return nil, burrito.WrappedErrorf("Cannot merge %s and %s", TypeName(template), TypeName(parent))
	}

	return parent, nil
}

// ToString converts an interface to a string.
func ToString(obj interface{}) string {
	if b, ok := obj.(JsonType); ok {
		obj = b.Unbox()
	}
	if obj == nil {
		return "null"
	}
	if b, ok := obj.(float64); ok {
		return strconv.FormatFloat(b, 'f', -1, 64)
	}
	if b, ok := obj.(float32); ok {
		return strconv.FormatFloat(float64(b), 'f', -1, 64)
	}
	if b, ok := obj.(int); ok {
		return strconv.FormatInt(int64(b), 10)
	}
	if b, ok := obj.(int32); ok {
		return strconv.FormatInt(int64(b), 10)
	}
	if b, ok := obj.(bool); ok && b {
		return strconv.FormatBool(b)
	}
	if b, ok := obj.(string); ok {
		return b
	}
	jsonc, err := json.MarshalJSONC(obj, false)
	if err != nil {
		panic(err)
	}
	return string(jsonc)
}

// ToPrettyString converts an interface to a string. In case of an object or array, it will be pretty printed.
func ToPrettyString(obj interface{}) string {
	if b, ok := obj.(JsonType); ok {
		obj = b.Unbox()
	}
	if obj == nil {
		return "null"
	}
	if b, ok := obj.(float64); ok {
		return strconv.FormatFloat(b, 'f', -1, 64)
	}
	if b, ok := obj.(float32); ok {
		return strconv.FormatFloat(float64(b), 'f', -1, 64)
	}
	if b, ok := obj.(int); ok {
		return strconv.FormatInt(int64(b), 10)
	}
	if b, ok := obj.(int32); ok {
		return strconv.FormatInt(int64(b), 10)
	}
	if b, ok := obj.(bool); ok && b {
		return strconv.FormatBool(b)
	}
	if b, ok := obj.(string); ok {
		return b
	}
	jsonc, err := json.MarshalJSONC(obj, true)
	if err != nil {
		panic(err)
	}
	return string(jsonc)
}

func IndexOf(array interface{}, value interface{}) int {
	if array == nil {
		return -1
	}
	if IsArray(array) {
		val := Box(value)
		for i, v := range AsArray(array).Value {
			if v.Equals(val) {
				return i
			}
		}
	}
	if arr, ok := array.([]interface{}); ok {
		for i, v := range arr {
			if v == value {
				return i
			}
		}
	}
	return -1
}

// CreateRange creates a range of numbers from start to end as a JSON array.
func CreateRange(start, end int32) JsonArray {
	var result []JsonType
	if start > end {
		return NewJsonArray()
	}
	for i := start; i <= end; i++ {
		result = append(result, AsNumber(i))
	}
	return JsonArray{Value: result}
}

// DeleteNulls removes all keys with null values from the given JSON object.
func DeleteNulls(object JsonObject) JsonObject {
	for _, k := range object.Keys() {
		v := object.Get(k)
		if IsObject(v) {
			object.Put(k, DeleteNulls(AsObject(v)))
		} else if IsArray(v) {
			object.Put(k, DeleteNullsFromArray(AsArray(v)))
		} else if v == nil || v == Null {
			object.Remove(k)
		}
	}
	return object
}

// DeleteNullsFromArray removes all keys inside elements of JSON object type with null values from the given JSON array.
func DeleteNullsFromArray(array JsonArray) JsonArray {
	for i, v := range array.Value {
		if IsObject(v) {
			array.Value[i] = DeleteNulls(AsObject(v))
		} else if IsArray(v) {
			array.Value[i] = DeleteNullsFromArray(AsArray(v))
		}
	}
	return array
}

// ParseJsonObject parses a JSON string into a JSON object. It includes support for comments and detects common syntax errors.
func ParseJsonObject(str []byte) (JsonObject, error) {
	dat, err := json.UnmarshallJSONC(str)
	if err != nil {
		return NewJsonObject(), err
	}
	if !IsObject(dat) {
		return NewJsonObject(), burrito.WrappedErrorf("JSON must be an object")
	}
	return AsObject(dat), nil
}

// ParseJsonArray parses a JSON string into a JSON array. It includes support for comments and detects common syntax errors.
func ParseJsonArray(str []byte) (JsonArray, error) {
	dat, err := json.UnmarshallJSONC(str)
	if err != nil {
		return NewJsonArray(), err
	}
	if !IsArray(dat) {
		return NewJsonArray(), burrito.WrappedErrorf("JSON must be an array")
	}
	return AsArray(dat), nil
}

// JsonAction is an enum for the different actions that can be performed via jsonte.
type JsonAction int

const (
	Value JsonAction = iota
	Iteration
	Literal
	Predicate
)

// String returns a string representation of the given JsonAction.
func (a JsonAction) String() string {
	switch a {
	case Value:
		return "Value"
	case Iteration:
		return "Iteration"
	case Literal:
		return "Literal"
	case Predicate:
		return "Predicate"
	}
	return string(rune(a))
}

func incompatibleTypesError(a, b interface{}) error {
	return burrito.WrappedErrorf("Cannot compare %s and %s", TypeName(a), TypeName(b))
}
