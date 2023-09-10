package types

import (
	"fmt"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"math"
	"strconv"
	"strings"
)

// JsonPath represents a simplified JSONPath
type JsonPath struct {
	Path        []JsonType
	parent      JsonType
	parentIndex JsonType
}

func (t *JsonPath) Parent() JsonType {
	return t.parent
}

func (t *JsonPath) ParentIndex() JsonType {
	return t.parentIndex
}

func (t *JsonPath) UpdateParent(parent JsonType, parentIndex JsonType) {
	t.parent = parent
	t.parentIndex = parentIndex
}

func (t *JsonPath) LessThan(other JsonType) (bool, error) {
	return false, burrito.WrappedErrorf("JsonPaths cannot be compared")
}

// StringValue returns a string representation of the semantic version
func (t *JsonPath) StringValue() string {
	sb := "#"
	for _, p := range t.Path {
		if i, ok := p.(*JsonNumber); ok {
			sb += fmt.Sprintf("[%d]", i.IntValue())
		} else if s, ok := p.(*JsonString); ok {
			sb += fmt.Sprintf("/%s", s.StringValue())
		} else {
			sb += fmt.Sprintf("[Unsupported value %s]", p.StringValue())
		}
	}
	return sb
}

// BoolValue returns a string representation of the semantic version
func (t *JsonPath) BoolValue() bool {
	return true
}

// Equals returns true if the two semantic versions are equal
func (t *JsonPath) Equals(value JsonType) bool {
	if IsJsonPath(value) {
		return t.StringValue() == value.StringValue()
	}
	return false
}

func (t *JsonPath) Unbox() interface{} {
	return t.StringValue()
}

func (t *JsonPath) Negate() JsonType {
	return NaN()
}

func (t *JsonPath) Index(i JsonType) (JsonType, error) {
	if b, ok := i.(*JsonNumber); ok {
		index := int(b.IntValue())
		if index < 0 {
			index = len(t.Path) + index
		}
		if index >= 0 && index < len(t.Path) {
			return t.Path[index], nil
		} else {
			return Null, burrito.WrappedErrorf("Index out of bounds: %d", index)
		}
	}
	return Null, burrito.WrappedErrorf("Index must be a number: %s", i.StringValue())
}

func (t *JsonPath) SetIndex(i, value JsonType) error {
	if b, ok := i.(*JsonNumber); ok {
		index := int(b.IntValue())
		if index < 0 {
			index = len(t.Path) + index
		}
		if index >= 0 && index < len(t.Path) {
			t.Path[index] = value
			return nil
		} else {
			return burrito.WrappedErrorf("Index out of bounds: %d", index)
		}
	}
	return burrito.WrappedErrorf("Index must be a number: %s", i.StringValue())
}

func (t *JsonPath) Add(i JsonType) JsonType {
	if b, ok := i.(*JsonNumber); ok {
		p := make([]JsonType, len(t.Path)+1)
		copy(p, t.Path)
		p[len(t.Path)] = b
		return &JsonPath{Path: p}
	} else if b, ok := i.(*JsonString); ok {
		p := make([]JsonType, len(t.Path)+1)
		copy(p, t.Path)
		p[len(t.Path)] = b
		return &JsonPath{Path: p}
	} else if b, ok := i.(*JsonPath); ok {
		p := make([]JsonType, len(t.Path)+len(b.Path))
		copy(p, t.Path)
		copy(p[len(t.Path):], b.Path)
		return &JsonPath{Path: p}
	} else {
		return NewString(t.StringValue() + i.StringValue())
	}
}

func (t *JsonPath) IsEmpty() bool {
	return t.Path == nil || len(t.Path) == 0
}

func (t *JsonPath) ParentPath() *JsonPath {
	if len(t.Path) == 0 {
		return t
	}
	return &JsonPath{Path: t.Path[:len(t.Path)-1]}
}

func (t *JsonPath) Get(x JsonType) (JsonType, error) {
	if _, ok := x.(*JsonObject); !ok {
		if _, ok := x.(*JsonArray); !ok {
			return nil, burrito.WrappedErrorf("Cannot get %s from %s", t.StringValue(), x.StringValue())
		}
	}
	var err error
	for i := 0; i < len(t.Path); i++ {
		//parent := x
		x, err = x.Index(t.Path[i])
		//x.UpdateParent(parent, t.Path[i])
		if err != nil {
			return nil, burrito.WrapErrorf(err, "Cannot get %s from %s", t.StringValue(), x.StringValue())
		}
	}
	return x, nil
}

func (t *JsonPath) Set(x, value JsonType) (JsonType, error) {
	original := x
	if _, ok := x.(*JsonObject); !ok {
		if _, ok := x.(*JsonArray); !ok {
			return original, burrito.WrappedErrorf("Cannot set %s in %s", t.StringValue(), x.StringValue())
		}
	}
	var err error
	for i := 0; i < len(t.Path)-1; i++ {
		x, err = x.Index(t.Path[i])
		if err != nil {
			return original, burrito.WrapErrorf(err, "Cannot get %s from %s", t.StringValue(), x.StringValue())
		}
	}
	if b, ok := x.(*JsonObject); ok {
		if k, ok := t.Path[len(t.Path)-1].(*JsonString); ok {
			b.Value.Put(k.StringValue(), value)
			return original, nil
		} else {
			return original, burrito.WrappedErrorf("Cannot set %s in %s", t.StringValue(), x.StringValue())
		}
	} else if b, ok := x.(*JsonArray); ok {
		if k, ok := t.Path[len(t.Path)-1].(*JsonNumber); ok {
			if k.IntValue() < 0 {
				k = AsNumber(int32(len(b.Value)) + k.IntValue())
			}
			if k.IntValue()-int32(len(b.Value)) == 1 {
				b.Value = append(b.Value, value)
				return original, nil
			} else if k.IntValue() > int32(len(b.Value)) {
				return original, burrito.WrappedErrorf("Cannot set %s in %s", t.StringValue(), x.StringValue())
			}
			b.Value[k.IntValue()] = value
			return original, nil
		} else {
			return original, burrito.WrappedErrorf("Cannot set %s in %s", t.StringValue(), x.StringValue())
		}
	} else {
		return original, burrito.WrappedErrorf("Cannot set %s in %s", t.StringValue(), x.StringValue())
	}
}

// IsJsonPath returns true if the given interface is a semver object.
func IsJsonPath(obj interface{}) bool {
	if obj == nil {
		return false
	}
	if _, ok := obj.(*JsonPath); ok {
		return true
	}
	return false
}

func AsJsonPath(obj interface{}) *JsonPath {
	if obj == nil {
		return nil
	}
	if b, ok := obj.(*JsonPath); ok {
		return b
	}
	if b, ok := obj.(*JsonString); ok {
		path, err := ParseJsonPath(b.StringValue())
		if err != nil {
			return nil
		}
		return path
	}
	return nil
}

func ParseJsonPath(path string) (*JsonPath, error) {
	path = strings.TrimPrefix(path, "#")
	if !strings.HasPrefix(path, "/") && !strings.HasPrefix(path, "[") {
		path = "/" + path
	}
	parts := make([]JsonType, 0)
	runes := []rune(path)
	for i := 0; i < len(runes)-1; i++ {
		if runes[i] == '[' {
			end := strings.IndexRune(string(runes[i+1:]), ']')
			if end == -1 {
				return nil, burrito.WrappedErrorf("Unclosed index notation: %s", path)
			}
			atoi, err := strconv.Atoi(string(runes[i+1 : i+1+end]))
			if err != nil {
				return nil, burrito.WrappedErrorf("Index is not a number: %s", path)
			}
			parts = append(parts, AsNumber(atoi))
			i += end
		} else if runes[i] == '/' {
			end := strings.IndexRune(string(runes[i+1:]), '/')
			end1 := strings.IndexRune(string(runes[i+1:]), '[')
			if end == -1 && end1 == -1 {
				end = len(runes) - i - 1
			}
			if end == -1 {
				end = math.MaxInt32
			}
			if end1 == -1 {
				end1 = math.MaxInt32
			}
			if i+1+int(math.Min(float64(end), float64(end1))) > len(runes) {
				break
			}
			parts = append(parts, NewString(string(runes[i+1:i+1+int(math.Min(float64(end), float64(end1)))])))
			i += int(math.Min(float64(end), float64(end1)))
		} else {
			return nil, burrito.WrappedErrorf("Invalid path: %s", path)
		}
	}
	return &JsonPath{Path: parts}, nil
}
