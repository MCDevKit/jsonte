package types

import "github.com/Bedrock-OSS/go-burrito/burrito"

// JsonNull is a struct that represents a null value.
type JsonNull struct {
}

var Null = &JsonNull{}

func (n *JsonNull) LessThan(other JsonType) (bool, error) {
	if other == nil {
		return true, nil
	}
	if other == Null {
		return false, nil
	}
	if other.Equals(n) {
		return false, nil
	}
	result, err := other.LessThan(n)
	return !result, err
}

func (n *JsonNull) BoolValue() bool {
	return false
}

func (n *JsonNull) StringValue() string {
	return "null"
}

func (n *JsonNull) Equals(value JsonType) bool {
	if value == Null || value == nil {
		return true
	}
	if b, ok := value.(JsonType); ok && b == Null {
		return true
	}
	return false
}

func (n *JsonNull) Unbox() interface{} {
	return nil
}

func (n *JsonNull) Negate() JsonType {
	return True()
}

func (n *JsonNull) Index(i JsonType) (JsonType, error) {
	return Null, burrito.WrappedErrorf("Cannot access %s from a null", i.StringValue())
}

func (n *JsonNull) Add(i JsonType) JsonType {
	if i == Null {
		return Null
	}
	return i.Add(n)
}
