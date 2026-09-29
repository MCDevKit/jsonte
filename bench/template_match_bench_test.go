package bench

import (
	"testing"

	"github.com/MCDevKit/jsonte/jsonte"
)

// BenchmarkFindTemplateMatchesMCFunction scans one generated 90-line mcfunction
// for `#{...}` expressions. Non-standard tokens are never cached, so every
// call pays the full scan like a real compile does.
func BenchmarkFindTemplateMatchesMCFunction(b *testing.B) {
	p := ensureCompilePack()
	input := p.functions[p.funcNames[0]]
	b.SetBytes(int64(len(input)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matches, err := jsonte.FindTemplateMatches(input, "#", "")
		if err != nil || len(matches) == 0 {
			b.Fatal(err, len(matches))
		}
	}
}
