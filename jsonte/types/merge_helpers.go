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

		needsProcessing := false
		forEachObjectEntryUntil(current, func(key string, value JsonType) bool {
			if key == "" {
				return false
			}
			if key[0] == '^' || (key[0] == '$' && !IsReservedKey(key)) {
				needsProcessing = true
				return true
			}
			switch typed := value.(type) {
			case *JsonObject:
				if typed != nil {
					stack = append(stack, typed)
				}
			case *JsonArray:
				if arrayNeedsMergeProcessing(typed) {
					needsProcessing = true
					return true
				}
			}
			return false
		})
		if needsProcessing {
			return true
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
				if typed != nil && len(typed.Value) > 0 {
					stack = append(stack, typed)
				}
			}
		}
	}
	return false
}

func forEachObjectEntryUntil(obj *JsonObject, fn func(string, JsonType) bool) {
	if obj == nil || fn == nil {
		return
	}
	if obj.Value != nil && obj.StackValue == nil {
		obj.Value.ForEachUntil(func(key string, value JsonType) bool {
			return fn(key, value)
		})
		return
	}
	for _, key := range obj.Keys() {
		if fn(key, obj.Get(key)) {
			return
		}
	}
}
