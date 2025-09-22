package types

import (
	"fmt"
	"testing"
)

var benchmarkNeedsMerge bool

func BenchmarkObjectNeedsMergeProcessingDeep(b *testing.B) {
	obj := buildObjectWithDeepOverride(6, 4)
	if !objectNeedsMergeProcessing(obj) {
		b.Fatalf("expected object to require merge processing")
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkNeedsMerge = objectNeedsMergeProcessing(obj)
	}
}

func BenchmarkArrayNeedsMergeProcessingDeep(b *testing.B) {
	arr := buildArrayWithDeepOverride(4, 3, 6)
	if !arrayNeedsMergeProcessing(arr) {
		b.Fatalf("expected array to require merge processing")
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkNeedsMerge = arrayNeedsMergeProcessing(arr)
	}
}

func buildObjectWithDeepOverride(depth, width int) *JsonObject {
	root := NewJsonObject()
	current := root

	for level := 0; level < depth; level++ {
		for i := 0; i < width; i++ {
			key := fmt.Sprintf("k_%d_%d", level, i)
			current.Put(key, AsNumber(level*width+i))
		}

		if level == depth-1 {
			override := NewJsonObject()
			override.Put("value", AsNumber(depth))
			current.Put("$override", override)
			break
		}

		childKey := fmt.Sprintf("child_%d", level)
		child := NewJsonObject()
		current.Put(childKey, child)
		current = child
	}

	return root
}

func buildArrayWithDeepOverride(depth, width, arrayLen int) *JsonArray {
	arr := NewJsonArrayWithCapacity(arrayLen)
	for i := 0; i < arrayLen; i++ {
		child := NewJsonObject()
		fillArrayOverride(child, depth, width, arrayLen, i == arrayLen-1)
		arr.Append(child)
	}
	return arr
}

func fillArrayOverride(node *JsonObject, depth, width, arrayLen int, addOverride bool) {
	for i := 0; i < width; i++ {
		key := fmt.Sprintf("v_%d", i)
		node.Put(key, AsNumber(i))
	}

	if depth == 0 {
		if addOverride {
			override := NewJsonObject()
			override.Put("flag", AsString("override"))
			node.Put("$override", override)
		}
		return
	}

	children := NewJsonArrayWithCapacity(arrayLen)
	for i := 0; i < arrayLen; i++ {
		child := NewJsonObject()
		fillArrayOverride(child, depth-1, width, arrayLen, addOverride && i == arrayLen-1)
		children.Append(child)
	}
	node.Put("children", children)
}
