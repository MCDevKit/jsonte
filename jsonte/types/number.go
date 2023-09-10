package types

import (
	"encoding/json"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"math"
	"strconv"
)

// Number is an interface that represents a number, that can be either integer or decimal.
type Number interface {
	// IntValue returns the number as an integer.
	IntValue() int
	// FloatValue returns the number as a float.
	FloatValue() float64
}

// JsonNumber is a struct that represents a number, that can be either integer or decimal.
type JsonNumber struct {
	Value       float64
	Decimal     bool
	parent      JsonType
	parentIndex JsonType
}

func (t *JsonNumber) Parent() JsonType {
	return t.parent
}

func (t *JsonNumber) ParentIndex() JsonType {
	return t.parentIndex
}

func (t *JsonNumber) UpdateParent(parent JsonType, parentIndex JsonType) {
	t.parent = parent
	t.parentIndex = parentIndex
}

func (t *JsonNumber) LessThan(other JsonType) (bool, error) {
	if other == nil || other == Null {
		return t.FloatValue() < float64(0), nil
	}
	if IsNumber(other) {
		return t.FloatValue() < AsNumber(other).FloatValue(), nil
	}
	return false, incompatibleTypesError(t, other)
}

func (t *JsonNumber) IntValue() int32 {
	return int32(t.Value)
}

// toFixed rounds a float to a given precision.
func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return math.Round(num*output) / output
}

func (t *JsonNumber) FloatValue() float64 {
	if t.Decimal {
		return toFixed(t.Value, 6)
	} else {
		return math.Floor(t.Value)
	}
}

func (t *JsonNumber) BoolValue() bool {
	if toFixed(t.Value, 6) == 0 {
		return false
	}
	return true
}

func (t *JsonNumber) StringValue() string {
	if t.Decimal {
		return strconv.FormatFloat(t.FloatValue(), 'f', -1, 64)
	}
	return strconv.FormatInt(int64(t.IntValue()), 10)
}

func (t *JsonNumber) Equals(value JsonType) bool {
	if value == Null {
		return false
	}
	if IsNumber(value) {
		return t.FloatValue() == AsNumber(value).FloatValue()
	}
	return false
}

func (t *JsonNumber) Unbox() interface{} {
	if t.Decimal {
		return t.FloatValue()
	}
	return t.IntValue()
}

func (t *JsonNumber) Negate() JsonType {
	return &JsonNumber{
		Value:   -t.FloatValue(),
		Decimal: t.Decimal,
	}
}

func (t *JsonNumber) Index(i JsonType) (JsonType, error) {
	return Null, burrito.WrappedErrorf("Cannot access %s from a number", i.StringValue())
}

func (t *JsonNumber) Add(i JsonType) JsonType {
	if i == Null {
		return &JsonNumber{
			Value:   t.FloatValue(),
			Decimal: t.Decimal,
		}
	}
	if IsNumber(i) {
		return &JsonNumber{
			Value:   t.FloatValue() + AsNumber(i).FloatValue(),
			Decimal: t.Decimal || AsNumber(i).Decimal,
		}
	}
	if IsBool(i) {
		if i.BoolValue() {
			return &JsonNumber{
				Value:   t.FloatValue() + 1,
				Decimal: t.Decimal,
			}
		}
		return &JsonNumber{
			Value:   t.FloatValue(),
			Decimal: t.Decimal || AsNumber(i).Decimal,
		}
	}
	return NewString(t.StringValue() + i.StringValue())
}

// AsNumber converts an interface to a JSON number.
func AsNumber(obj interface{}) *JsonNumber {
	if obj == nil {
		return &JsonNumber{
			Value:   0,
			Decimal: false,
		}
	}
	if b, ok := obj.(*JsonNumber); ok {
		return b
	}
	if b, ok := obj.(JsonType); ok {
		obj = b.Unbox()
	}
	// Past this point, we are dealing with a primitive type.
	if b, ok := obj.(float64); ok {
		return &JsonNumber{
			Value:   b,
			Decimal: true,
		}
	}
	if b, ok := obj.(float32); ok {
		return &JsonNumber{
			Value:   float64(b),
			Decimal: true,
		}
	}
	if b, ok := obj.(int); ok {
		return &JsonNumber{
			Value:   float64(b),
			Decimal: false,
		}
	}
	if b, ok := obj.(int32); ok {
		return &JsonNumber{
			Value:   float64(b),
			Decimal: false,
		}
	}
	if b, ok := obj.(uint32); ok {
		return &JsonNumber{
			Value:   float64(b),
			Decimal: false,
		}
	}
	if b, ok := obj.(int64); ok {
		return &JsonNumber{
			Value:   float64(b),
			Decimal: false,
		}
	}
	if b, ok := obj.(bool); ok && b {
		return &JsonNumber{
			Value:   1,
			Decimal: false,
		}
	}
	if b, ok := obj.(string); ok {
		result, err := strconv.ParseInt(b, 10, 64)
		if err != nil {
			result1, err := strconv.ParseFloat(b, 64)
			if err != nil {
				return &JsonNumber{
					Value:   0,
					Decimal: false,
				}
			}
			return &JsonNumber{
				Value:   result1,
				Decimal: true,
			}
		}
		return &JsonNumber{
			Value:   float64(result),
			Decimal: false,
		}
	}
	if b, ok := obj.(json.Number); ok {
		result, err := strconv.ParseInt(string(b), 10, 64)
		if err != nil {
			result1, err := strconv.ParseFloat(string(b), 64)
			if err != nil {
				return &JsonNumber{
					Value:   0,
					Decimal: false,
				}
			}
			return &JsonNumber{
				Value:   result1,
				Decimal: true,
			}
		}
		return &JsonNumber{
			Value:   float64(result),
			Decimal: false,
		}
	}
	// TODO: Consider returning an error here.
	return &JsonNumber{
		Value:   0,
		Decimal: false,
	}
}

// IsNumber returns true if the given interface is a number.
func IsNumber(obj interface{}) bool {
	if obj == nil {
		return false
	}
	if _, ok := obj.(json.Number); ok {
		return true
	}
	if _, ok := obj.(*JsonNumber); ok {
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
