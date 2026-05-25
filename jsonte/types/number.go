package types

import (
	"encoding/json"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"math"
	"strconv"
)

const smallIntMax = 4096
const smallIntNegMin = -128

var smallIntCache [smallIntMax]*JsonNumber
var smallIntNegCache [128]*JsonNumber // indices for [-128, -1]
var smallFloatCache [smallIntMax]*JsonNumber
var smallIntStrings [smallIntMax]string
var smallIntNegStrings [128]string // strings for [-128, -1]

func init() {
	for i := 0; i < smallIntMax; i++ {
		smallIntCache[i] = &JsonNumber{Value: float64(i), Decimal: false}
		smallFloatCache[i] = &JsonNumber{Value: float64(i), Decimal: true}
		smallIntStrings[i] = strconv.FormatInt(int64(i), 10)
	}
	for i := 0; i < 128; i++ {
		v := -(i + 1)
		smallIntNegCache[i] = &JsonNumber{Value: float64(v), Decimal: false}
		smallIntNegStrings[i] = strconv.FormatInt(int64(v), 10)
	}
}

func cachedInt(n int64) *JsonNumber {
	if n >= 0 && n < smallIntMax {
		return smallIntCache[n]
	}
	if n >= smallIntNegMin && n < 0 {
		return smallIntNegCache[-(n + 1)]
	}
	return nil
}

func cachedFloat(f float64) *JsonNumber {
	i := int64(f)
	if f == float64(i) && i >= 0 && i < smallIntMax {
		return smallFloatCache[i]
	}
	return nil
}

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
	if other == nil || IsNull(other) {
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
	iv := t.IntValue()
	if iv >= 0 && int(iv) < smallIntMax {
		return smallIntStrings[iv]
	}
	if iv < 0 && iv >= smallIntNegMin {
		return smallIntNegStrings[-(iv + 1)]
	}
	return strconv.FormatInt(int64(iv), 10)
}

func (t *JsonNumber) Equals(value JsonType) bool {
	if IsNull(value) {
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
	if IsNull(i) {
		if !t.Decimal {
			if c := cachedInt(int64(t.IntValue())); c != nil {
				return c
			}
		}
		return &JsonNumber{Value: t.FloatValue(), Decimal: t.Decimal}
	}
	if IsNumber(i) {
		n := AsNumber(i)
		decimal := t.Decimal || n.Decimal
		sum := t.FloatValue() + n.FloatValue()
		if !decimal {
			iv := int64(sum)
			if float64(iv) == sum {
				if c := cachedInt(iv); c != nil {
					return c
				}
			}
		} else {
			if c := cachedFloat(sum); c != nil {
				return c
			}
		}
		return &JsonNumber{Value: sum, Decimal: decimal}
	}
	if IsBool(i) {
		if i.BoolValue() {
			sum := t.FloatValue() + 1
			if !t.Decimal {
				iv := int64(sum)
				if c := cachedInt(iv); c != nil {
					return c
				}
			}
			return &JsonNumber{Value: sum, Decimal: t.Decimal}
		}
		if c := cachedInt(int64(t.IntValue())); !t.Decimal && c != nil {
			return c
		}
		return &JsonNumber{Value: t.FloatValue(), Decimal: t.Decimal}
	}
	return NewString(t.StringValue() + i.StringValue())
}

// NewIntNumber returns a *JsonNumber for the given integer, using the cache for small non-negative values.
func NewIntNumber(n int64) *JsonNumber {
	if c := cachedInt(n); c != nil {
		return c
	}
	return &JsonNumber{Value: float64(n), Decimal: false}
}

// NewFloatNumber returns a *JsonNumber for the given float, using the integer cache when the value is a whole number.
func NewFloatNumber(f float64) *JsonNumber {
	if c := cachedFloat(f); c != nil {
		return c
	}
	return &JsonNumber{Value: f, Decimal: true}
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
		if c := cachedFloat(b); c != nil {
			return c
		}
		return &JsonNumber{Value: b, Decimal: true}
	}
	if b, ok := obj.(float32); ok {
		if c := cachedFloat(float64(b)); c != nil {
			return c
		}
		return &JsonNumber{Value: float64(b), Decimal: true}
	}
	if b, ok := obj.(int); ok {
		if c := cachedInt(int64(b)); c != nil {
			return c
		}
		return &JsonNumber{Value: float64(b), Decimal: false}
	}
	if b, ok := obj.(int32); ok {
		if c := cachedInt(int64(b)); c != nil {
			return c
		}
		return &JsonNumber{Value: float64(b), Decimal: false}
	}
	if b, ok := obj.(uint32); ok {
		if c := cachedInt(int64(b)); c != nil {
			return c
		}
		return &JsonNumber{Value: float64(b), Decimal: false}
	}
	if b, ok := obj.(int64); ok {
		if c := cachedInt(b); c != nil {
			return c
		}
		return &JsonNumber{Value: float64(b), Decimal: false}
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
