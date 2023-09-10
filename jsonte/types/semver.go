package types

import (
	"fmt"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"strings"
)

// Semver represents a semantic version
type Semver struct {
	Major       int
	Minor       int
	Patch       int
	parent      JsonType
	parentIndex JsonType
}

var majorAliases = []string{
	"major",
	"a",
	"x",
}
var minorAliases = []string{
	"minor",
	"b",
	"y",
}
var patchAliases = []string{
	"patch",
	"c",
	"z",
}

func (t *Semver) Parent() JsonType {
	return t.parent
}

func (t *Semver) ParentIndex() JsonType {
	return t.parentIndex
}

func (t *Semver) UpdateParent(parent JsonType, parentIndex JsonType) {
	t.parent = parent
	t.parentIndex = parentIndex
}

// CompareTo compares two semantic versions. Returns 0 if they are equal, -1 if
// the receiver is less than the argument, and 1 if the receiver is greater than
// the argument.
func (t *Semver) CompareTo(other *Semver) int {
	if t.Major > other.Major {
		return 1
	} else if t.Major < other.Major {
		return -1
	} else if t.Minor > other.Minor {
		return 1
	} else if t.Minor < other.Minor {
		return -1
	} else if t.Patch > other.Patch {
		return 1
	} else if t.Patch < other.Patch {
		return -1
	}
	return 0
}

func (t *Semver) LessThan(other JsonType) (bool, error) {
	if IsSemver(other) {
		return t.CompareTo(AsSemver(other)) == -1, nil
	}
	return false, burrito.WrapErrorf(nil, "Cannot compare semver to %s", other.StringValue())
}

// StringValue returns a string representation of the semantic version
func (t *Semver) StringValue() string {
	return fmt.Sprintf("%d.%d.%d", t.Major, t.Minor, t.Patch)
}

// BoolValue returns a string representation of the semantic version
func (t *Semver) BoolValue() bool {
	return true
}

// Equals returns true if the two semantic versions are equal
func (t *Semver) Equals(value JsonType) bool {
	if IsSemver(value) {
		return t.CompareTo(AsSemver(value)) == 0
	}
	return false
}

func (t *Semver) Unbox() interface{} {
	return t.StringValue()
}

func (t *Semver) Negate() JsonType {
	return NaN()
}

func (t *Semver) Index(i JsonType) (JsonType, error) {
	if value, ok := i.(*JsonString); ok {
		if IndexOf(majorAliases, value.StringValue()) != -1 {
			return AsNumber(t.Major), nil
		}
		if IndexOf(minorAliases, value.StringValue()) != -1 {
			return AsNumber(t.Minor), nil
		}
		if IndexOf(patchAliases, value.StringValue()) != -1 {
			return AsNumber(t.Patch), nil
		}
		return Null, burrito.WrappedErrorf("Cannot access %s because it is not a valid semver field", i.StringValue())
	}
	return Null, burrito.WrappedErrorf("Index must be a string: %s", i.StringValue())
}

func (t *Semver) Add(i JsonType) JsonType {
	return NewString(t.StringValue() + i.StringValue())
}

func (t *Semver) IsEmpty() bool {
	return t.Major == 0 && t.Minor == 0 && t.Patch == 0
}

// ParseSemverString parses a string representation of a semantic version
func ParseSemverString(version string) (*Semver, error) {
	var major, minor, patch int
	split := strings.Split(version, ".")
	if len(split) == 0 {
		return nil, burrito.WrapErrorf(nil, "Invalid semver string: %s", version)
	}
	_, err := fmt.Sscanf(split[0], "%d", &major)
	if err != nil {
		return nil, burrito.WrapErrorf(err, "Invalid semver string: %s", version)
	}
	if len(split) == 1 {
		return &Semver{major, 0, 0, nil, nil}, nil
	}
	_, err = fmt.Sscanf(split[1], "%d", &minor)
	if err != nil {
		return nil, burrito.WrapErrorf(err, "Invalid semver string: %s", version)
	}
	if len(split) == 2 {
		return &Semver{major, minor, 0, nil, nil}, nil
	}
	_, err = fmt.Sscanf(split[2], "%d", &patch)
	if err != nil {
		return nil, burrito.WrapErrorf(err, "Invalid semver string: %s", version)
	}
	return &Semver{major, minor, patch, nil, nil}, nil
}

// ParseSemverArray parses an array representation of a semantic version
func ParseSemverArray(version []interface{}) (*Semver, error) {
	size := len(version)
	if size == 0 {
		return nil, burrito.WrapErrorf(nil, "Invalid semver array: %v", version)
	}
	major, ok := version[0].(int32)
	if !ok {
		return nil, burrito.WrapErrorf(nil, "Invalid semver array: %v", version)
	}
	if size == 1 {
		return &Semver{int(major), 0, 0, nil, nil}, nil
	}
	minor, ok := version[1].(int32)
	if !ok {
		return nil, burrito.WrapErrorf(nil, "Invalid semver array: %v", version)
	}
	if size == 2 {
		return &Semver{int(major), int(minor), 0, nil, nil}, nil
	}
	patch, ok := version[2].(int32)
	if !ok {
		return nil, burrito.WrapErrorf(nil, "Invalid semver array: %v", version)
	}
	return &Semver{int(major), int(minor), int(patch), nil, nil}, nil
}

// IsSemver returns true if the given interface is a semver object.
func IsSemver(obj interface{}) bool {
	if obj == nil {
		return false
	}
	if _, ok := obj.(*Semver); ok {
		return true
	}
	return false
}

func AsSemver(obj interface{}) *Semver {
	if obj == nil {
		return nil
	}
	if b, ok := obj.(*Semver); ok {
		return b
	}
	return nil
}
