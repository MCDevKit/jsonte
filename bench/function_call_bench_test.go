package bench

import (
	"fmt"
	"testing"

	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/functions"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/gammazero/deque"
)

const benchmarkLambdaChainExpression = "entities.filter(e => e.family.contains(fam)).map(e => e.id)"

// BenchmarkEvalLambdaChain measures a warm evaluation (parse tree cached) of an
// instance-function call chain with lambdas, exercising the function-call path
// and name resolution in VisitField.
func BenchmarkEvalLambdaChain(b *testing.B) {
	entities := make([]interface{}, 48)
	for i := range entities {
		family := []interface{}{"mob"}
		if i%2 == 0 {
			family = append(family, "monster")
		}
		entities[i] = map[string]interface{}{
			"id":     fmt.Sprintf("bench:entity_%d", i),
			"family": family,
		}
	}
	scopeObject := types.AsObject(map[string]interface{}{
		"entities": entities,
		"fam":      "monster",
	})
	scope := deque.Deque[*types.JsonObject]{}
	scope.PushBack(scopeObject)

	r, err := jsonte.Eval(benchmarkLambdaChainExpression, scope, "#")
	if err != nil {
		b.Fatal(err)
	}
	if arr, ok := r.Value.(*types.JsonArray); !ok || len(arr.Value) != 24 {
		b.Fatalf("unexpected result: %s", r.Value.StringValue())
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result jsonte.Result
	for i := 0; i < b.N; i++ {
		r, err := jsonte.Eval(benchmarkLambdaChainExpression, scope, "#")
		if err != nil {
			b.Fatal(err)
		}
		result = r
	}
	benchmarkEvalResult = result
}

func benchCallFunction(b *testing.B, name string, args []types.JsonType) {
	b.Helper()
	if _, err := functions.CallFunction(name, args, nil); err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := functions.CallFunction(name, args, nil); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCallFunctionCount(b *testing.B) {
	arr := types.NewJsonArray()
	for i := 0; i < 8; i++ {
		arr.Append(types.AsNumber(i))
	}
	benchCallFunction(b, "count", []types.JsonType{arr})
}

func BenchmarkCallFunctionRoundPrecision(b *testing.B) {
	benchCallFunction(b, "round", []types.JsonType{types.AsNumber(3.14159), types.AsNumber(2)})
}

func BenchmarkCallFunctionToLowerCase(b *testing.B) {
	benchCallFunction(b, "toLowerCase", []types.JsonType{types.AsString("Hello World")})
}
