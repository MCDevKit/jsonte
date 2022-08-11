package utils

import (
	"encoding/json"
	"math"
	"reflect"
	"strconv"
	"strings"
)

type JsonObject map[string]interface{}
type JsonArray []interface{}

type Number interface {
	IntValue() int
	FloatValue() float64
	BoolValue() bool
	StringValue() string
}

type JsonNumber struct {
	Number
	Value   float64
	Decimal bool
}

func (n JsonNumber) IntValue() int {
	return int(n.Value)
}

func (n JsonNumber) FloatValue() float64 {
	if n.Decimal {
		return n.Value
	} else {
		return math.Floor(n.Value)
	}
}

func (n JsonNumber) BoolValue() bool {
	if n.Value == 0 {
		return false
	}
	return true
}

func (n JsonNumber) StringValue() string {
	if n.Decimal {
		return strconv.FormatFloat(n.Value, 'f', -1, 64)
	}
	return strconv.FormatInt(int64(math.Floor(n.Value)), 10)
}

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
	return false
}

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
	return JsonNumber{
		Value:   0,
		Decimal: false,
	}
}

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
	if b, ok := obj.(bool); ok && b {
		return strconv.FormatBool(b)
	}
	if b, ok := obj.(string); ok {
		return b
	}
	marshal, err := json.Marshal(UnwrapContainers(obj))
	if err != nil {
		return "null"
	}
	return string(marshal)
}

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
	if _, ok := obj.(bool); ok {
		return true
	}
	return false
}

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

func IsObject(obj interface{}) bool {
	if obj == nil {
		return false
	}
	rt := reflect.TypeOf(obj)
	switch rt.Kind() {
	case reflect.Map:
		return true
	}
	return false
}

func MergeObject(template, parent JsonObject) (JsonObject, error) {
	result := JsonObject{}
	for k, v := range parent {
		if strings.HasPrefix(k, "$") && k != "$comment" {
			result[strings.TrimPrefix(k, "$")] = v
		} else if _, ok := template[k]; !ok {
			if IsObject(v) {
				merge, err := MergeObject(nil, v.(JsonObject))
				if err != nil {
					return nil, err
				}
				result[k] = merge
			} else if IsArray(v) {
				merge, err := MergeArray(nil, v.(JsonArray))
				if err != nil {
					return nil, err
				}
				result[k] = merge
			} else {
				result[k] = v
			}
		} else {
			if IsObject(v) {
				merge, err := MergeObject(template[k].(JsonObject), v.(JsonObject))
				if err != nil {
					return nil, err
				}
				result[k] = merge
			} else if IsArray(v) {
				merge, err := MergeArray(template[k].(JsonArray), v.(JsonArray))
				if err != nil {
					return nil, err
				}
				result[k] = merge
			} else {
				result[k] = v
			}
		}
	}
	return result, nil
}

func MergeArray(template, parent JsonArray) (JsonArray, error) {
	result := JsonArray{}
	for _, v := range template {
		if IsObject(v) {
			merge, err := MergeObject(nil, v.(JsonObject))
			if err != nil {
				return nil, err
			}
			result = append(result, merge)
		} else if IsArray(v) {
			merge, err := MergeArray(nil, v.(JsonArray))
			if err != nil {
				return nil, err
			}
			result = append(result, merge)
		} else {
			result = append(result, v)
		}
	}
	for _, v := range parent {
		if IsObject(v) {
			merge, err := MergeObject(nil, v.(JsonObject))
			if err != nil {
				return nil, err
			}
			result = append(result, merge)
		} else if IsArray(v) {
			merge, err := MergeArray(nil, v.(JsonArray))
			if err != nil {
				return nil, err
			}
			result = append(result, merge)
		} else {
			result = append(result, v)
		}
	}
	return result, nil
}

func IsEqual(a, b interface{}) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a == b {
		return true
	}
	if IsNumber(a) && IsNumber(b) {
		return ToNumber(a) == ToNumber(b)
	}
	if IsArray(a) && IsArray(b) {
		return IsEqualArray(a.(JsonArray), b.(JsonArray))
	}
	if IsObject(a) && IsObject(b) {
		return IsEqualObject(a.(JsonObject), b.(JsonObject))
	}
	return false
}

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
		if ToNumber(a).Value < ToNumber(b).Value {
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

func IsEqualObject(a, b JsonObject) bool {
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

func IsEqualArray(a, b JsonArray) bool {
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

func CreateRange(start, end int) JsonArray {
	result := JsonArray{}
	if start > end {
		return result
	}
	for i := start; i <= end; i++ {
		result = append(result, i)
	}
	return result
}

func UnescapeString(str string) string {
	if len(str) < 3 {
		return ""
	}
	str = strings.Trim(str, "\"'")
	str = strings.ReplaceAll(str, "\\\\\"", "\"")
	str = strings.ReplaceAll(str, "\\\\'", "'")
	str = strings.ReplaceAll(str, "\\\\n", "\n")
	str = strings.ReplaceAll(str, "\\\\\\\\", "\\\\")
	return str
}

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
	rt := reflect.TypeOf(obj)
	switch rt.Kind() {
	case reflect.Slice, reflect.Array:
		result := JsonArray{}
		for _, v := range obj.(JsonArray) {
			result = append(result, UnwrapContainers(v))
		}
		return result
	case reflect.Map:
		result := JsonObject{}
		for k, v := range obj.(JsonObject) {
			result[k] = UnwrapContainers(v)
		}
		return result
	}
	return obj

}

type EvaluationError struct {
	Message string
	Path    string
	Err     error
}

func (e *EvaluationError) Error() string {
	if e.Err != nil && e.Path != "" {
		return e.Message + ": " + e.Err.Error() + " at " + e.Path
	} else if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	} else if e.Path != "" {
		return e.Message + " at " + e.Path
	} else {
		return e.Message
	}
}

type JsonAction int
type JsonLambda func(args []interface{}) (interface{}, error)

const (
	Value JsonAction = iota
	Iteration
	Literal
	Predicate
)
