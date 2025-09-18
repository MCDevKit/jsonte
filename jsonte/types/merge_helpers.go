package types

import "strconv"

// joinObjectPath mirrors fmt.Sprintf("%s/%s", path, key) without fmt allocation.
func joinObjectPath(path, key string) string {
	return path + "/" + key
}

// joinArrayPath mirrors fmt.Sprintf("%s[%d]", path, idx) without fmt allocation.
func joinArrayPath(path string, idx int) string {
	return path + "[" + strconv.Itoa(idx) + "]"
}
