package utils

import (
	"fmt"
	"strings"
)

// Semver represents a semantic version
type Semver struct {
	Major int
	Minor int
	Patch int
}

// CompareTo compares two semantic versions. Returns 0 if they are equal, -1 if
// the receiver is less than the argument, and 1 if the receiver is greater than
// the argument.
func (s Semver) CompareTo(other Semver) int {
	if s.Major > other.Major {
		return 1
	} else if s.Major < other.Major {
		return -1
	} else if s.Minor > other.Minor {
		return 1
	} else if s.Minor < other.Minor {
		return -1
	} else if s.Patch > other.Patch {
		return 1
	} else if s.Patch < other.Patch {
		return -1
	}
	return 0
}

// String returns a string representation of the semantic version
func (s Semver) String() string {
	return fmt.Sprintf("%d.%d.%d", s.Major, s.Minor, s.Patch)
}

// Equals returns true if the two semantic versions are equal
func (s Semver) Equals(other Semver) bool {
	return s.CompareTo(other) == 0
}

// ParseSemverString parses a string representation of a semantic version
func ParseSemverString(version string) (Semver, error) {
	var major, minor, patch int
	split := strings.Split(version, ".")
	if len(split) == 0 {
		return Semver{}, WrapErrorf(nil, "Invalid semver string: %s", version)
	}
	_, err := fmt.Sscanf(split[0], "%d", &major)
	if err != nil {
		return Semver{}, WrapErrorf(err, "Invalid semver string: %s", version)
	}
	if len(split) == 1 {
		return Semver{major, 0, 0}, nil
	}
	_, err = fmt.Sscanf(split[1], "%d", &minor)
	if err != nil {
		return Semver{}, WrapErrorf(err, "Invalid semver string: %s", version)
	}
	if len(split) == 2 {
		return Semver{major, minor, 0}, nil
	}
	_, err = fmt.Sscanf(split[2], "%d", &patch)
	if err != nil {
		return Semver{}, WrapErrorf(err, "Invalid semver string: %s", version)
	}
	return Semver{major, minor, patch}, nil
}

// ParseSemverArray parses an array representation of a semantic version
func ParseSemverArray(version []interface{}) (Semver, error) {
	version = UnwrapContainers(version).([]interface{})
	size := len(version)
	if size == 0 {
		return Semver{}, WrapErrorf(nil, "Invalid semver array: %v", version)
	}
	major, ok := version[0].(int32)
	if !ok {
		return Semver{}, WrapErrorf(nil, "Invalid semver array: %v", version)
	}
	if size == 1 {
		return Semver{int(major), 0, 0}, nil
	}
	minor, ok := version[1].(int32)
	if !ok {
		return Semver{}, WrapErrorf(nil, "Invalid semver array: %v", version)
	}
	if size == 2 {
		return Semver{int(major), int(minor), 0}, nil
	}
	patch, ok := version[2].(int32)
	if !ok {
		return Semver{}, WrapErrorf(nil, "Invalid semver array: %v", version)
	}
	return Semver{int(major), int(minor), int(patch)}, nil
}
