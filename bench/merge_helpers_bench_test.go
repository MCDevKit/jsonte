package bench

import (
	"fmt"

	"github.com/MCDevKit/jsonte/jsonte/types"
)

var benchmarkNeedsMerge bool

func buildObjectWithDeepOverride(depth, width int) *types.JsonObject {
	root := types.NewJsonObject()
	current := root

	for level := 0; level < depth; level++ {
		for i := 0; i < width; i++ {
			key := fmt.Sprintf("k_%d_%d", level, i)
			current.Put(key, types.AsNumber(level*width+i))
		}

		if level == depth-1 {
			override := types.NewJsonObject()
			override.Put("value", types.AsNumber(depth))
			current.Put("$override", override)
			break
		}

		childKey := fmt.Sprintf("child_%d", level)
		child := types.NewJsonObject()
		current.Put(childKey, child)
		current = child
	}

	return root
}

func buildArrayWithDeepOverride(depth, width, arrayLen int) *types.JsonArray {
	arr := types.NewJsonArrayWithCapacity(arrayLen)
	for i := 0; i < arrayLen; i++ {
		child := types.NewJsonObject()
		fillArrayOverride(child, depth, width, arrayLen, i == arrayLen-1)
		arr.Append(child)
	}
	return arr
}

func fillArrayOverride(node *types.JsonObject, depth, width, arrayLen int, addOverride bool) {
	for i := 0; i < width; i++ {
		key := fmt.Sprintf("v_%d", i)
		node.Put(key, types.AsNumber(i))
	}

	if depth == 0 {
		if addOverride {
			override := types.NewJsonObject()
			override.Put("flag", types.AsString("override"))
			node.Put("$override", override)
		}
		return
	}

	children := types.NewJsonArrayWithCapacity(arrayLen)
	for i := 0; i < arrayLen; i++ {
		child := types.NewJsonObject()
		fillArrayOverride(child, depth-1, width, arrayLen, addOverride && i == arrayLen-1)
		children.Append(child)
	}
	node.Put("children", children)
}
