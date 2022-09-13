package utils

import "fmt"

type Semver struct {
	Major int
	Minor int
	Patch int
}

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

func (s Semver) String() string {
	return fmt.Sprintf("%d.%d.%d", s.Major, s.Minor, s.Patch)
}

func (s Semver) Equals(other Semver) bool {
	return s.CompareTo(other) == 0
}

func ParseSemverString(version string) (Semver, error) {
	var major, minor, patch int
	_, err := fmt.Sscanf(version, "%d.%d.%d", &major, &minor, &patch)
	if err != nil {
		return Semver{}, WrapErrorf(err, "Failed to parse semver string: %s", version)
	}
	return Semver{major, minor, patch}, nil
}

func ParseSemverArray(version []interface{}) Semver {
	return Semver{int(version[0].(float64)), int(version[1].(float64)), int(version[2].(float64))}
}
