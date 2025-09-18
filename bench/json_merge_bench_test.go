package bench

import (
	"github.com/MCDevKit/jsonte/jsonte/types"
	"testing"
)

var (
	mergeJSONTemplateObj    *types.JsonObject
	mergeJSONParentObj      *types.JsonObject
	mergeJSONArrayTemplate  *types.JsonObject
	mergeJSONArrayParent    *types.JsonObject
	mergeJSONModuleTemplate *types.JsonObject
	mergeJSONModuleParent   *types.JsonObject
)

func init() {
	types.Init()
	mergeJSONTemplateObj, _ = types.ParseJsonObject([]byte(`{"root":{"a":{"x":1},"b":2}}`))
	mergeJSONParentObj, _ = types.ParseJsonObject([]byte(`{"root":{"a":{"y":3},"c":4}}`))
	mergeJSONArrayTemplate, _ = types.ParseJsonObject([]byte(`{"arr":[1,2,3,4,5]}`))
	mergeJSONArrayParent, _ = types.ParseJsonObject([]byte(`{"arr":[6,7,8,9,10]}`))
	mergeJSONModuleTemplate, _ = types.ParseJsonObject([]byte(`{"groups":{"object_1":{"value":1},"object_2":{"value":2}}}`))
	mergeJSONModuleParent, _ = types.ParseJsonObject([]byte(`{"groups":{"$object_1":{"value":0},"^object_3":{"value":3}}}`))
}

func BenchmarkMergeJSONObject(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := types.MergeJSON(mergeJSONTemplateObj, mergeJSONParentObj, false); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMergeJSONArray(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := types.MergeJSON(mergeJSONArrayTemplate, mergeJSONArrayParent, false); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMergeJSONModuleOverrides(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := types.MergeJSON(mergeJSONModuleTemplate, mergeJSONModuleParent, false); err != nil {
			b.Fatal(err)
		}
	}
}
