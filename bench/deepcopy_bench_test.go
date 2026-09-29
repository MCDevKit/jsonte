package bench

import (
	"testing"

	"github.com/MCDevKit/jsonte/jsonte/types"
)

// BenchmarkDeepCopyScope deep-copies the compile benchmark's global scope,
// which is what module/template processing does repeatedly.
func BenchmarkDeepCopyScope(b *testing.B) {
	scope, err := types.ParseJsonObject([]byte(generateCompileScope()))
	if err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out := types.DeepCopyObject(scope)
		if out.Size() != scope.Size() {
			b.Fatal("size mismatch")
		}
	}
}
