package types

import (
	"encoding/json"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"math"
	"strconv"
)

// Number is an interface that represents a number, that can be either integer or decimal.
type Number interface {
	JsonType
	// IntValue returns the number as an integer.
	IntValue() int
	// FloatValue returns the number as a float.
	FloatValue() float64
}

// JsonNumber is a struct that represents a number, that can be either integer or decimal.
type JsonNumber struct {
	Number
	Value   float64
	Decimal bool
}

func (n JsonNumber) IsNull() bool {
	return false
}

func (n JsonNumber) LessThan(other JsonType) (bool, error) {
	if other == nil || other.IsNull() {
		return n.FloatValue() < float64(0), nil
	}
	if IsNumber(other) {
		return n.FloatValue() < AsNumber(other).FloatValue(), nil
	}
	return false, incompatibleTypesError(n, other)
}

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
		return toFixed(n.Value, 6)
	} else {
		return math.Floor(n.Value)
	}
}

func (n JsonNumber) BoolValue() bool {
	if toFixed(n.Value, 6) == 0 {
		return false
	}
	return true
}

func (n JsonNumber) StringValue() string {
	if n.Decimal {
		return strconv.FormatFloat(n.FloatValue(), 'f', -1, 64)
	}
	return strconv.FormatInt(int64(n.IntValue()), 10)
}

func (n JsonNumber) Equals(value JsonType) bool {
	if value == Null {
		return false
	}
	if IsNumber(value) {
		return n.FloatValue() == AsNumber(value).FloatValue()
	}
	return false
}

func (n JsonNumber) Unbox() interface{} {
	if n.Decimal {
		return n.FloatValue()
	}
	return n.IntValue()
}

func (n JsonNumber) Negate() JsonType {
	return JsonNumber{
		Value:   -n.FloatValue(),
		Decimal: n.Decimal,
	}
}

func (n JsonNumber) Index(i JsonType) (JsonType, error) {
	return Null, burrito.WrappedErrorf("Cannot access %s from a number", i.StringValue())
}

func (n JsonNumber) Add(i JsonType) JsonType {
	if i.IsNull() {
		return JsonNumber{
			Value:   n.FloatValue(),
			Decimal: n.Decimal,
		}
	}
	if IsNumber(i) {
		return JsonNumber{
			Value:   n.FloatValue() + AsNumber(i).FloatValue(),
			Decimal: n.Decimal || AsNumber(i).Decimal,
		}
	}
	if IsBool(i) {
		if i.BoolValue() {
			return JsonNumber{
				Value:   n.FloatValue() + 1,
				Decimal: n.Decimal,
			}
		}
		return JsonNumber{
			Value:   n.FloatValue(),
			Decimal: n.Decimal || AsNumber(i).Decimal,
		}
	}
	return NewString(n.StringValue() + i.StringValue())
}

// AsNumber converts an interface to a JSON number.
func AsNumber(obj interface{}) JsonNumber {
	if obj == nil {
		return JsonNumber{
			Value:   0,
			Decimal: false,
		}
	}
	if b, ok := obj.(JsonNumber); ok {
		return b
	}
	if b, ok := obj.(JsonType); ok {
		obj = b.Unbox()
	}
	// Past this point, we are dealing with a primitive type.
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
	// TODO: Consider returning an error here.
	return JsonNumber{
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
