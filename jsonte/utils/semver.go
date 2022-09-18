package utils

import "fmt"

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
	_, err := fmt.Sscanf(version, "%d.%d.%d", &major, &minor, &patch)
	if err != nil {
		return Semver{}, WrapErrorf(err, "Failed to parse semver string: %s", version)
	}
	return Semver{major, minor, patch}, nil
}

// ParseSemverArray parses an array representation of a semantic version
func ParseSemverArray(version []interface{}) Semver {
	return Semver{int(version[0].(float64)), int(version[1].(float64)), int(version[2].(float64))}
}
