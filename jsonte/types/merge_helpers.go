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

func objectNeedsMergeProcessing(obj *JsonObject) bool {
	if obj == nil || obj.Size() == 0 {
		return false
	}
	stack := []*JsonObject{obj}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current == nil {
			continue
		}
		keys := current.Keys()
		for _, key := range keys {
			if key == "" {
				continue
			}
			if key[0] == '^' || (key[0] == '$' && !IsReservedKey(key)) {
				return true
			}
			switch typed := current.Get(key).(type) {
			case *JsonObject:
				stack = append(stack, typed)
			case *JsonArray:
				if arrayNeedsMergeProcessing(typed) {
					return true
				}
			}
		}
	}
	return false
}

func arrayNeedsMergeProcessing(arr *JsonArray) bool {
	if arr == nil || len(arr.Value) == 0 {
		return false
	}
	stack := []*JsonArray{arr}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current == nil {
			continue
		}
		for _, item := range current.Value {
			switch typed := item.(type) {
			case *JsonObject:
				if objectNeedsMergeProcessing(typed) {
					return true
				}
			case *JsonArray:
				stack = append(stack, typed)
			}
		}
	}
	return false
}
