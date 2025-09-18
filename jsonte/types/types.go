package types

import (
	"fmt"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/json"
	"github.com/MCDevKit/jsonte/jsonte/utils"
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
	// LessThan returns true if the value is less than the given value.
	LessThan(value JsonType) (bool, error)
	// Unbox returns the value without any containers.
	Unbox() interface{}
	// Negate returns the value negated with a minus operator.
	Negate() JsonType
	// Index returns the value at the given index.
	Index(index JsonType) (JsonType, error)
	// Add returns the sum of the value and the given value.
	Add(index JsonType) JsonType
	// Parent returns the parent of the value or nil.
	Parent() JsonType
	// ParentIndex returns the index of the value in the parent or nil.
	ParentIndex() JsonType
	// UpdateParent updates the parent of the value.
	UpdateParent(parent JsonType, parentIndex JsonType)
}

type TypeDescriptor struct {
	Type   reflect.Type
	Name   string
	IsType func(interface{}) bool
	AsType func(interface{}) JsonType
}

var TypeDescriptors []TypeDescriptor

func NaN() JsonType {
	return &JsonString{Value: "NaN"}
}

func Init() {
	TypeDescriptors = []TypeDescriptor{
		{
			Type:   reflect.TypeOf(Null),
			Name:   "null",
			IsType: func(i interface{}) bool { return i == nil || IsNull(i) },
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
			Type:   reflect.TypeOf(JsonPath{}),
			Name:   "path",
			IsType: IsJsonPath,
			AsType: func(i interface{}) JsonType { return AsJsonPath(i) },
		},
		{
			Type:   reflect.TypeOf(JsonLambda{}),
			Name:   "lambda",
			IsType: func(i interface{}) bool { _, ok := i.(*JsonLambda); return ok },
			AsType: func(i interface{}) JsonType {
				if b, ok := i.(*JsonLambda); ok {
					return b
				}
				utils.BadDeveloperError("Not a lambda")
				return nil
			},
		},
		{
			Type:   reflect.TypeOf(JsonSignal{}),
			Name:   "signal",
			IsType: IsSignal,
			AsType: func(i interface{}) JsonType {
				if b, ok := i.(*JsonSignal); ok {
					return b
				}
				utils.BadDeveloperError("Not a signal")
				return nil
			},
		},
	}
}

func TypeName(obj interface{}) string {
	if obj == nil || IsNull(obj) {
		return "null"
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
	if IsJsonPath(obj) {
		return "path"
	}
	if _, ok := obj.(*JsonLambda); ok {
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
	utils.BadDeveloperError("Unknown type: " + TypeName(obj))
	return nil
}

func MergeJSON(template, parent JsonType, keepOverrides bool) (JsonType, error) {
	if template == nil || IsNull(template) {
		return parent, nil
	}
	if parent == nil || IsNull(parent) {
		return template, nil
	}

	templateObj, templateIsJsonObject := template.(*JsonObject)
	parentObj, parentIsJsonObject := parent.(*JsonObject)
	if templateIsJsonObject && parentIsJsonObject {
		return MergeObject(templateObj, parentObj, keepOverrides, "#"), nil
	}

	templateArray, templateIsJsonArray := template.(*JsonArray)
	parentArray, parentIsJsonArray := parent.(*JsonArray)
	if templateIsJsonArray && parentIsJsonArray {
		return MergeArray(templateArray, parentArray, keepOverrides, "#"), nil
	}

	templateIsObject := templateIsJsonObject
	if !templateIsObject {
		templateIsObject = IsObject(template)
	}
	parentIsObject := parentIsJsonObject
	if !parentIsObject {
		parentIsObject = IsObject(parent)
	}
	if templateIsObject || parentIsObject {
		if templateIsObject && parentIsObject {
			return MergeObject(AsObject(template), AsObject(parent), keepOverrides, "#"), nil
		}
		templateType := TypeName(template)
		parentType := TypeName(parent)
		return nil, burrito.WrappedErrorf("Cannot merge %s and %s", templateType, parentType)
	}

	templateIsArray := templateIsJsonArray
	if !templateIsArray {
		templateIsArray = IsArray(template)
	}
	parentIsArray := parentIsJsonArray
	if !parentIsArray {
		parentIsArray = IsArray(parent)
	}
	if templateIsArray || parentIsArray {
		if templateIsArray && parentIsArray {
			return MergeArray(AsArray(template), AsArray(parent), keepOverrides, "#"), nil
		}
		templateType := TypeName(template)
		parentType := TypeName(parent)
		return nil, burrito.WrappedErrorf("Cannot merge %s and %s", templateType, parentType)
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
		utils.BadDeveloperError(err.Error())
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
		utils.BadDeveloperError(err.Error())
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
func CreateRange(start, end int32) *JsonArray {
	var result []JsonType
	if start > end {
		return NewJsonArray()
	}
	for i := start; i <= end; i++ {
		result = append(result, AsNumber(i))
	}
	return &JsonArray{Value: result}
}

// DeleteNulls removes all keys with null values from the given JSON value.
func DeleteNulls(object JsonType) JsonType {
	if IsObject(object) {
		return DeleteNullsFromObject(AsObject(object))
	} else if IsArray(object) {
		return DeleteNullsFromArray(AsArray(object))
	}
	return object
}

// DeleteNullsFromObject removes all keys with null values from the given JSON object.
func DeleteNullsFromObject(object *JsonObject) *JsonObject {
	keys := object.Keys()
	toRemove := make([]int, 0)
	for i, k := range keys {
		v := object.Get(k)
		if IsObject(v) {
			object.Put(k, DeleteNulls(AsObject(v)))
		} else if IsArray(v) {
			object.Put(k, DeleteNullsFromArray(AsArray(v)))
		} else if v == nil || IsNull(v) {
			toRemove = append(toRemove, i)
		}
	}
	for _, idx := range toRemove {
		object.Remove(keys[idx])
	}
	return object
}

// DeleteNullsFromArray removes all keys inside elements of JSON object type with null values from the given JSON array.
func DeleteNullsFromArray(array *JsonArray) *JsonArray {
	for i, v := range array.Value {
		if IsObject(v) {
			array.Value[i] = DeleteNulls(AsObject(v))
		} else if IsArray(v) {
			array.Value[i] = DeleteNullsFromArray(AsArray(v))
		}
	}
	return array
}

// ParseJsonValue parses a JSON string into a JSON object. It includes support for comments and detects common syntax errors.
func ParseJsonValue(str []byte) (JsonType, error) {
	dat, err := json.UnmarshallJSONC(str)
	if err != nil {
		return NewJsonObject(), err
	}
	if IsObject(dat) {
		return AsObject(dat), nil
	} else if IsArray(dat) {
		return AsArray(dat), nil
	} else if IsString(dat) {
		return AsString(dat), nil
	} else if IsNumber(dat) {
		return AsNumber(dat), nil
	} else if IsBool(dat) {
		return AsBool(dat), nil
	}
	return Null, burrito.WrappedErrorf("JSON must be an object, array, string, number or boolean")
}

// ParseJsonObject parses a JSON string into a JSON object. It includes support for comments and detects common syntax errors.
func ParseJsonObject(str []byte) (*JsonObject, error) {
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
func ParseJsonArray(str []byte) (*JsonArray, error) {
	dat, err := json.UnmarshallJSONC(str)
	if err != nil {
		return NewJsonArray(), err
	}
	if !IsArray(dat) {
		return NewJsonArray(), burrito.WrappedErrorf("JSON must be an array")
	}
	return AsArray(dat), nil
}

// MergeValues merges two JSON values into a new JSON value.
// If the same value, that is not an object or an array exists in both objects, the value from the second object will be used.
func MergeValues(template, parent JsonType, keepOverrides bool, path string) (JsonType, error) {
	if IsNull(template) {
		return parent, nil
	}
	if IsNull(parent) {
		return template, nil
	}
	if IsObject(template) {
		if !IsObject(parent) {
			return template, burrito.WrappedErrorf("Cannot merge object with non-object")
		}
		return MergeObject(AsObject(template), AsObject(parent), keepOverrides, path), nil
	} else if IsArray(template) {
		if !IsArray(parent) {
			return template, burrito.WrappedErrorf("Cannot merge array with non-array")
		}
		return MergeArray(AsArray(template), AsArray(parent), keepOverrides, path), nil
	}
	return template, nil
}

// DeepCopyValue creates a deep copy of the given JSON value.
func DeepCopyValue(object JsonType) JsonType {
	if IsObject(object) {
		return DeepCopyObject(AsObject(object))
	} else if IsArray(object) {
		return DeepCopyArray(AsArray(object))
	}
	return Box(object.Unbox())
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
