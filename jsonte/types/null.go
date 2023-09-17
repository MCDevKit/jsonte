package types

import "github.com/Bedrock-OSS/go-burrito/burrito"

// JsonNull is a struct that represents a null value.
type JsonNull struct {
	parent      JsonType
	parentIndex JsonType
}

var Null = &JsonNull{}

func NullWithParent(parent JsonType, parentIndex JsonType) *JsonNull {
	return &JsonNull{parent, parentIndex}
}

func (t *JsonNull) Parent() JsonType {
	return t.parent
}

func (t *JsonNull) ParentIndex() JsonType {
	return t.parentIndex
}

func (t *JsonNull) UpdateParent(parent JsonType, parentIndex JsonType) {
	t.parent = parent
	t.parentIndex = parentIndex
}

func (t *JsonNull) LessThan(other JsonType) (bool, error) {
	if other == nil {
		return true, nil
	}
	if IsNull(other) {
		return false, nil
	}
	if other.Equals(t) {
		return false, nil
	}
	result, err := other.LessThan(t)
	return !result, err
}

func (t *JsonNull) BoolValue() bool {
	return false
}

func (t *JsonNull) StringValue() string {
	return "null"
}

func (t *JsonNull) Equals(value JsonType) bool {
	if value == nil || IsNull(value) {
		return true
	}
	if b, ok := value.(JsonType); ok && IsNull(b) {
		return true
	}
	return false
}

func (t *JsonNull) Unbox() interface{} {
	return nil
}

func (t *JsonNull) Negate() JsonType {
	return True()
}

func (t *JsonNull) Index(i JsonType) (JsonType, error) {
	return Null, burrito.WrappedErrorf("Cannot access %s from a null", i.StringValue())
}

func (t *JsonNull) Add(i JsonType) JsonType {
	if IsNull(i) {
		return Null
	}
	return i.Add(t)
}

func IsNull(i interface{}) bool {
	_, ok := i.(*JsonNull)
	return ok
}
