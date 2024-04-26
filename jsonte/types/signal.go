package types

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
)

type JsonSignal struct {
	Value JsonType
	Type  SignalType
}

type SignalType int

const (
	SignalReturn SignalType = iota
	SignalBreak
	SignalContinue
)

var Break = &JsonSignal{Value: Null, Type: SignalBreak}
var Continue = &JsonSignal{Value: Null, Type: SignalContinue}

func (t *JsonSignal) Parent() JsonType {
	return t.Value.Parent()
}

func (t *JsonSignal) ParentIndex() JsonType {
	return t.Value.ParentIndex()
}

func (t *JsonSignal) UpdateParent(parent JsonType, parentIndex JsonType) {
	t.Value.UpdateParent(parent, parentIndex)
}

func (t *JsonSignal) StringValue() string {
	if t.Type == SignalBreak {
		return "break"
	} else if t.Type == SignalContinue {
		return "continue"
	}
	return t.Value.StringValue()
}

func (t *JsonSignal) BoolValue() bool {
	return t.Value.BoolValue()
}

func (t *JsonSignal) Equals(value JsonType) bool {
	return t.Value.Equals(value)
}

func (t *JsonSignal) Unbox() interface{} {
	return t.Value.Unbox()
}

func (t *JsonSignal) Negate() JsonType {
	return t.Value.Negate()
}

func (t *JsonSignal) Index(i JsonType) (JsonType, error) {
	return t.Value.Index(i)
}

func (t *JsonSignal) Add(i JsonType) JsonType {
	return t.Value.Add(i)
}

func (t *JsonSignal) LessThan(other JsonType) (bool, error) {
	return false, burrito.WrappedErrorf("Arrays cannot be compared")
}

func NewReturn(value JsonType) *JsonSignal {
	return &JsonSignal{Value: value, Type: SignalReturn}
}

func IsReturn(value JsonType) bool {
	if signal, ok := value.(*JsonSignal); ok {
		return signal.Type == SignalReturn
	}
	return false
}

func IsBreak(value JsonType) bool {
	if signal, ok := value.(*JsonSignal); ok {
		return signal.Type == SignalBreak
	}
	return false
}

func IsContinue(value JsonType) bool {
	if signal, ok := value.(*JsonSignal); ok {
		return signal.Type == SignalContinue
	}
	return false
}

func IsSignal(value interface{}) bool {
	_, ok := value.(*JsonSignal)
	return ok
}
