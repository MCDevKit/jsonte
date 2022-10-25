package utils

import (
	"encoding/json"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// Number is an interface that represents a number, that can be either integer or decimal.
type Number interface {
	// IntValue returns the number as an integer.
	IntValue() int
	// FloatValue returns the number as a float.
	FloatValue() float64
	// BoolValue returns the number as a boolean.
	BoolValue() bool
	// StringValue returns the number as a string.
	StringValue() string
}

// JsonNumber is a struct that represents a number, that can be either integer or decimal.
type JsonNumber struct {
	Number
	Value   float64
	Decimal bool
}

// CacheDir is a directory used for cache
var CacheDir string

func (n JsonNumber) IntValue() int32 {
	return int32(n.Value)
}

// toFixed rounds a float to a given precision.
func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return math.Round(num*output) / output
}

func (n JsonNumber) FloatValue() float64 {
	if n.Decimal {
		return toFixed(n.Value, 10)
	} else {
		return math.Floor(n.Value)
	}
}

func (n JsonNumber) BoolValue() bool {
	if toFixed(n.Value, 10) == 0 {
		return false
	}
	return true
}

func (n JsonNumber) StringValue() string {
	if n.Decimal {
		return strconv.FormatFloat(n.Value, 'f', -1, 64)
	}
	return strconv.FormatInt(int64(n.IntValue()), 10)
}

// ToBoolean converts an interface to a boolean.
func ToBoolean(obj interface{}) bool {
	if obj == nil {
		return false
	}
	if b, ok := obj.(bool); ok {
		return b
	}
	if b, ok := obj.(int); ok && b != 0 {
		return true
	}
	if b, ok := obj.(float64); ok && b != 0 {
		return true
	}
	if b, ok := obj.(float32); ok && b != 0 {
		return true
	}
	if b, ok := obj.(string); ok && strings.Trim(b, "\n\r") != "" {
		return true
	}
	if b, ok := obj.(JsonNumber); ok {
		return b.BoolValue()
	}
	return obj != nil
}

// ToNumber converts an interface to a JSON number.
func ToNumber(obj interface{}) JsonNumber {
	if obj == nil {
		return JsonNumber{
			Value:   0,
			Decimal: false,
		}
	}
	if b, ok := obj.(JsonNumber); ok {
		return b
	}
	if b, ok := obj.(float64); ok {
		return JsonNumber{
			Value:   b,
			Decimal: true,
		}
	}
	if b, ok := obj.(float32); ok {
		return JsonNumber{
			Value:   float64(b),
			Decimal: true,
		}
	}
	if b, ok := obj.(int); ok {
		return JsonNumber{
			Value:   float64(b),
			Decimal: false,
		}
	}
	if b, ok := obj.(int32); ok {
		return JsonNumber{
			Value:   float64(b),
			Decimal: false,
		}
	}
	if b, ok := obj.(uint32); ok {
		return JsonNumber{
			Value:   float64(b),
			Decimal: false,
		}
	}
	if b, ok := obj.(int64); ok {
		return JsonNumber{
			Value:   float64(b),
			Decimal: false,
		}
	}
	if b, ok := obj.(bool); ok && b {
		return JsonNumber{
			Value:   1,
			Decimal: false,
		}
	}
	if b, ok := obj.(string); ok {
		result, err := strconv.ParseInt(b, 10, 64)
		if err != nil {
			result1, err := strconv.ParseFloat(b, 64)
			if err != nil {
				return JsonNumber{
					Value:   0,
					Decimal: false,
				}
			}
			return JsonNumber{
				Value:   result1,
				Decimal: true,
			}
		}
		return JsonNumber{
			Value:   float64(result),
			Decimal: false,
		}
	}
	if b, ok := obj.(json.Number); ok {
		result, err := strconv.ParseInt(string(b), 10, 64)
		if err != nil {
			result1, err := strconv.ParseFloat(string(b), 64)
			if err != nil {
				return JsonNumber{
					Value:   0,
					Decimal: false,
				}
			}
			return JsonNumber{
				Value:   result1,
				Decimal: true,
			}
		}
		return JsonNumber{
			Value:   float64(result),
			Decimal: false,
		}
	}
	return JsonNumber{
		Value:   0,
		Decimal: false,
	}
}

// ToString converts an interface to a string.
func ToString(obj interface{}) string {
	if obj == nil {
		return "null"
	}
	if b, ok := obj.(JsonNumber); ok {
		return b.StringValue()
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
	str, err := MarshalJSONC(UnwrapContainers(obj), false)
	if err != nil {
		return "null"
	}
	return strings.ReplaceAll(string(str), "\n", "")
}

// ToPrettyString converts an interface to a string. In case of an object or array, it will be pretty printed.
func ToPrettyString(obj interface{}) string {
	if obj == nil {
		return "null"
	}
	if b, ok := obj.(JsonNumber); ok {
		return b.StringValue()
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
	str, err := MarshalJSONC(UnwrapContainers(obj), true)
	if err != nil {
		return "null"
	}
	return string(str)
}

// IsNumber returns true if the given interface is a number.
func IsNumber(obj interface{}) bool {
	if obj == nil {
		return false
	}
	if _, ok := obj.(JsonNumber); ok {
		return true
	}
	if _, ok := obj.(float64); ok {
		return true
	}
	if _, ok := obj.(float32); ok {
		return true
	}
	if _, ok := obj.(int); ok {
		return true
	}
	if _, ok := obj.(int32); ok {
		return true
	}
	if _, ok := obj.(bool); ok {
		return true
	}
	return false
}

// IsArray returns true if the given interface is an array.
func IsArray(obj interface{}) bool {
	if obj == nil {
		return false
	}
	rt := reflect.TypeOf(obj)
	switch rt.Kind() {
	case reflect.Slice, reflect.Array:
		return true
	}
	return false
}

// AsArray returns the given interface as a JSON array.
func AsArray(obj interface{}) []interface{} {
	if obj == nil {
		return nil
	}
	rt := reflect.TypeOf(obj)
	switch rt.Kind() {
	case reflect.Slice, reflect.Array:
		rv := reflect.ValueOf(obj)
		result := make([]interface{}, rv.Len())
		for i := 0; i < rv.Len(); i++ {
			result[i] = rv.Index(i).Interface()
		}
		return result
	}
	return nil
}

// AsObject returns the given interface as a JSON object.
func AsObject(obj interface{}) NavigableMap[string, interface{}] {
	if obj == nil {
		return NewNavigableMap[string, interface{}]()
	}
	if b, ok := obj.(NavigableMap[string, interface{}]); ok {
		return b
	}
	rt := reflect.TypeOf(obj)
	switch rt.Kind() {
	case reflect.Map:
		rv := reflect.ValueOf(obj)
		result := NewNavigableMap[string, interface{}]()
		for _, key := range rv.MapKeys() {
			result.Put(key.String(), rv.MapIndex(key).Interface())
		}
		return result
	}
	return NewNavigableMap[string, interface{}]()
}

// IsObject returns true if the given interface is an object.
func IsObject(obj interface{}) bool {
	if obj == nil {
		return false
	}
	if _, ok := obj.(NavigableMap[string, interface{}]); ok {
		return true
	}
	rt := reflect.TypeOf(obj)
	switch rt.Kind() {
	case reflect.Map:
		return true
	}
	return false
}

// MergeObject merges two JSON objects into a new JSON object.
// If the same value, that is not an object or an array exists in both objects, the value from the second object will be used.
func MergeObject(template, parent NavigableMap[string, interface{}], keepOverrides bool) NavigableMap[string, interface{}] {
	result := NewNavigableMap[string, interface{}]()
	for _, k := range template.Keys() {
		v := template.Get(k)
		if IsObject(v) {
			if val, ok := v.(NavigableMap[string, interface{}]); ok {
				result.Put(k, DeepCopyObject(val))
			}
		} else if IsArray(v) {
			if val, ok := v.([]interface{}); ok {
				result.Put(k, DeepCopyArray(val))
			} else if val, ok := v.([]interface{}); ok {
				result.Put(k, DeepCopyArray(val))
			}
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
		if strings.HasPrefix(k, "$") && k != "$comment" {
			if keepOverrides {
				result.Put(k, v)
			} else {
				result.Put(strings.TrimPrefix(k, "$"), v)
				skipKeys = append(skipKeys, strings.TrimPrefix(k, "$"))
			}
		} else if !template.ContainsKey(k) {
			if IsObject(v) {
				merge := MergeObject(NewNavigableMap[string, interface{}](), v.(NavigableMap[string, interface{}]), keepOverrides)
				result.Put(k, merge)
			} else if IsArray(v) {
				merge := MergeArray(nil, v.([]interface{}), keepOverrides)
				result.Put(k, merge)
			} else {
				result.Put(k, v)
			}
		} else {
			if IsObject(v) && IsObject(result.Get(k)) {
				merge := MergeObject(template.Get(k).(NavigableMap[string, interface{}]), v.(NavigableMap[string, interface{}]), keepOverrides)
				result.Put(k, merge)
			} else if IsArray(v) && IsArray(template.Get(k)) {
				merge := MergeArray(template.Get(k).([]interface{}), v.([]interface{}), keepOverrides)
				result.Put(k, merge)
			} else {
				result.Put(k, v)
			}
		}
	}
	return result
}

// MergeArray merges two JSON arrays into a new JSON array.
func MergeArray(template, parent []interface{}, keepOverrides bool) []interface{} {
	var result []interface{}
	for _, v := range template {
		if IsObject(v) {
			merge := MergeObject(NewNavigableMap[string, interface{}](), v.(NavigableMap[string, interface{}]), keepOverrides)
			result = append(result, merge)
		} else if IsArray(v) {
			merge := MergeArray(nil, v.([]interface{}), keepOverrides)
			result = append(result, merge)
		} else {
			result = append(result, v)
		}
	}
	for _, v := range parent {
		if IsObject(v) {
			merge := MergeObject(NewNavigableMap[string, interface{}](), v.(NavigableMap[string, interface{}]), keepOverrides)
			result = append(result, merge)
		} else if IsArray(v) {
			merge := MergeArray(nil, v.([]interface{}), keepOverrides)
			result = append(result, merge)
		} else {
			result = append(result, v)
		}
	}
	return result
}

// IsEqual returns true if the given interfaces are equal.
func IsEqual(a, b interface{}) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if (IsObject(a) != IsObject(b)) || (IsArray(a) != IsArray(b)) || (IsNumber(a) != IsNumber(b)) {
		return false
	}
	if IsNumber(a) && IsNumber(b) {
		return ToNumber(a).FloatValue() == ToNumber(b).FloatValue()
	}
	if IsArray(a) && IsArray(b) {
		return IsEqualArray(a.([]interface{}), b.([]interface{}))
	}
	if IsObject(a) && IsObject(b) {
		return IsEqualObject(a.(NavigableMap[string, interface{}]), b.(NavigableMap[string, interface{}]))
	}
	if a == b {
		return true
	}
	return false
}

// Less returns true if the first interface is less than the second interface.
func Less(a, b interface{}) bool {
	if a == nil && b == nil {
		return false
	}
	if a == nil {
		return false
	}
	if b == nil {
		return true
	}
	if IsNumber(a) && IsNumber(b) {
		if ToNumber(a).FloatValue() < ToNumber(b).FloatValue() {
			return true
		} else {
			return false
		}
	}
	if s1, ok1 := a.(string); ok1 {
		if s2, ok2 := b.(string); ok2 {
			return strings.Compare(s1, s2) == -1
		}
	}
	return false
}

// IsEqualObject returns true if the given JSON objects are equal.
func IsEqualObject(a, b NavigableMap[string, interface{}]) bool {
	if a.Size() != b.Size() {
		return false
	}
	for _, k := range a.Keys() {
		if !IsEqual(b.Get(k), a.Get(k)) {
			return false
		}
	}
	return true
}

// IsEqualArray returns true if the given JSON arrays are equal.
func IsEqualArray(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if !IsEqual(b[k], v) {
			return false
		}
	}
	return true
}

// CreateRange creates a range of numbers from start to end as a JSON array.
func CreateRange(start, end int32) []interface{} {
	var result []interface{}
	if start > end {
		return result
	}
	for i := start; i <= end; i++ {
		result = append(result, ToNumber(i))
	}
	return result
}

// UnescapeString removes quotes and unescapes a string.
func UnescapeString(str string) string {
	if len(str) < 3 {
		return ""
	}
	str = str[1 : len(str)-1]
	str = strings.ReplaceAll(str, "\\\\\"", "\"")
	str = strings.ReplaceAll(str, "\\\\'", "'")
	str = strings.ReplaceAll(str, "\\\\n", "\n")
	str = strings.ReplaceAll(str, "\\\\\\\\", "\\\\")
	return str
}

// UnwrapContainers removes all containers from the given interface.
// Currently only unpacks JsonNumber into an actual number with correct type.
func UnwrapContainers(obj interface{}) interface{} {
	if obj == nil {
		return nil
	}
	if b, ok := obj.(JsonNumber); ok {
		if b.Decimal {
			return b.FloatValue()
		} else {
			return b.IntValue()
		}
	}
	if IsObject(obj) {
		object := AsObject(obj)
		for _, k := range object.Keys() {
			object.Put(k, UnwrapContainers(object.Get(k)))
		}
		return object
	} else if IsArray(obj) {
		array := AsArray(obj)
		for k, v := range array {
			array[k] = UnwrapContainers(v)
		}
		return array
	}
	return obj
}

// DeepCopyObject creates a deep copy of the given JSON object.
func DeepCopyObject(object NavigableMap[string, interface{}]) NavigableMap[string, interface{}] {
	result := NewNavigableMap[string, interface{}]()
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

// DeepCopyArray creates a deep copy of the given JSON array.
func DeepCopyArray(object []interface{}) []interface{} {
	var result []interface{}
	for _, v := range object {
		if IsObject(v) {
			result = append(result, DeepCopyObject(AsObject(v)))
		} else if IsArray(v) {
			result = append(result, DeepCopyArray(AsArray(v)))
		} else {
			result = append(result, v)
		}
	}
	return result
}

// DeleteNulls removes all keys with null values from the given JSON object.
func DeleteNulls(object NavigableMap[string, interface{}]) NavigableMap[string, interface{}] {
	for _, k := range object.Keys() {
		v := object.Get(k)
		if IsObject(v) {
			object.Put(k, DeleteNulls(AsObject(v)))
		} else if IsArray(v) {
			object.Put(k, DeleteNullsFromArray(AsArray(v)))
		} else if v == nil {
			object.Remove(k)
		}
	}
	return object
}

// DeleteNullsFromArray removes all keys inside elements of JSON object type with null values from the given JSON array.
func DeleteNullsFromArray(array []interface{}) []interface{} {
	for i, v := range array {
		if IsObject(v) {
			array[i] = DeleteNulls(AsObject(v))
		} else if IsArray(v) {
			array[i] = DeleteNullsFromArray(AsArray(v))
		}
	}
	return array
}

// ParseJsonObject parses a JSON string into a JSON object. It includes support for comments and detects common syntax errors.
func ParseJsonObject(str []byte) (NavigableMap[string, interface{}], error) {
	dat, err := UnmarshallJSONC(str)
	if err != nil {
		return NewNavigableMap[string, interface{}](), err
	}
	// Convert all numbers to JsonNumber
	if !IsObject(dat) {
		return NewNavigableMap[string, interface{}](), WrappedErrorf("JSON must be an object")
	}
	return convertNumbersObject(AsObject(dat)), nil
}

// ParseJsonArray parses a JSON string into a JSON array. It includes support for comments and detects common syntax errors.
func ParseJsonArray(str []byte) ([]interface{}, error) {
	dat, err := UnmarshallJSONC(str)
	if err != nil {
		return []interface{}{}, err
	}
	// Convert all numbers to JsonNumber
	if !IsArray(dat) {
		return []interface{}{}, WrappedErrorf("JSON must be an array")
	}
	return convertNumbersArray(AsArray(dat)), nil
}

func convertNumbersObject(object NavigableMap[string, interface{}]) NavigableMap[string, interface{}] {
	result := NewNavigableMap[string, interface{}]()
	for _, k := range object.Keys() {
		v := object.Get(k)
		if IsObject(v) {
			result.Put(k, convertNumbersObject(AsObject(v)))
		} else if IsArray(v) {
			result.Put(k, convertNumbersArray(AsArray(v)))
		} else if _, ok := v.(json.Number); ok {
			result.Put(k, ToNumber(v))
		} else {
			result.Put(k, v)
		}
	}
	return result
}

func convertNumbersArray(object []interface{}) []interface{} {
	var result []interface{}
	for _, v := range object {
		if IsObject(v) {
			result = append(result, convertNumbersObject(AsObject(v)))
		} else if IsArray(v) {
			result = append(result, convertNumbersArray(AsArray(v)))
		} else if _, ok := v.(json.Number); ok {
			result = append(result, ToNumber(v))
		} else {
			result = append(result, v)
		}
	}
	return result
}

// JsonAction is an enum for the different actions that can be performed via jsonte.
type JsonAction int

// JsonLambda is a function that can be executed.
type JsonLambda func(args []interface{}) (interface{}, error)

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
