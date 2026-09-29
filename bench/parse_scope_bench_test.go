package bench

import (
	"testing"

	"github.com/MCDevKit/jsonte/jsonte/types"
)

// BenchmarkParseScope measures JSONC parsing plus conversion into JsonType
// values for the generated compile scope.
func BenchmarkParseScope(b *testing.B) {
	data := []byte(generateCompileScope())
	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		obj, err := types.ParseJsonObject(data)
		if err != nil {
			b.Fatal(err)
		}
		compileSink += obj.Size()
	}
}
