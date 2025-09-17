package bench

import (
	"testing"

	"github.com/MCDevKit/jsonte/jsonte/functions"
	"github.com/MCDevKit/jsonte/jsonte/types"
)

var (
	arrayBenchmarkResolve = func(string) types.JsonType { return types.Null }

	arrayNumbers = buildSequentialArray(4096)
	arrayNested  = buildSequentialArray(2048)

	lambdaEven = types.NewLambda(func(this *types.JsonLambda, args []types.JsonType) (types.JsonType, error) {
		value := types.AsNumber(args[0]).IntValue()
		if value%2 == 0 {
			return types.True(), nil
		}
		return types.False(), nil
	}, "value => value % 2 == 0", nil, []string{"value"})

	lambdaEvenIndexed = types.NewLambda(func(this *types.JsonLambda, args []types.JsonType) (types.JsonType, error) {
		value := types.AsNumber(args[0]).IntValue()
		if value%2 == 0 {
			return types.True(), nil
		}
		return types.False(), nil
	}, "(value, index) => value % 2 == 0", nil, []string{"value", "index"})

	lambdaDouble = types.NewLambda(func(this *types.JsonLambda, args []types.JsonType) (types.JsonType, error) {
		value := types.AsNumber(args[0]).IntValue()
		return types.AsNumber(int(value) * 2), nil
	}, "value => value * 2", nil, []string{"value"})

	lambdaDoubleIndexed = types.NewLambda(func(this *types.JsonLambda, args []types.JsonType) (types.JsonType, error) {
		value := types.AsNumber(args[0]).IntValue()
		return types.AsNumber(int(value) * 2), nil
	}, "(value, index) => value * 2", nil, []string{"value", "index"})

	lambdaAccumulator = types.NewLambda(func(this *types.JsonLambda, args []types.JsonType) (types.JsonType, error) {
		prev := args[0]
		if types.IsNull(prev) {
			return args[1], nil
		}
		sum := types.AsNumber(prev).IntValue() + types.AsNumber(args[1]).IntValue()
		return types.AsNumber(int(sum)), nil
	}, "(prev, value) => prev + value", nil, []string{"prev", "value"})

	lambdaAccumulatorInit = types.NewLambda(func(this *types.JsonLambda, args []types.JsonType) (types.JsonType, error) {
		sum := types.AsNumber(args[0]).IntValue() + types.AsNumber(args[1]).IntValue()
		return types.AsNumber(int(sum)), nil
	}, "(prev, value) => prev + value", nil, []string{"prev", "value"})

	lambdaSquareKey = types.NewLambda(func(this *types.JsonLambda, args []types.JsonType) (types.JsonType, error) {
		value := types.AsNumber(args[0]).IntValue()
		return types.AsNumber(int(value * value)), nil
	}, "value => value * value", nil, []string{"value"})

	lambdaDescendingKey = types.NewLambda(func(this *types.JsonLambda, args []types.JsonType) (types.JsonType, error) {
		value := types.AsNumber(args[0]).FloatValue()
		return types.AsNumber(-value), nil
	}, "value => -value", nil, []string{"value"})

	lambdaFlatMap = types.NewLambda(func(this *types.JsonLambda, args []types.JsonType) (types.JsonType, error) {
		value := int(types.AsNumber(args[0]).IntValue())
		inner := &types.JsonArray{Value: []types.JsonType{types.AsNumber(value), types.AsNumber(value + 1)}}
		return inner, nil
	}, "value => [value, value + 1]", nil, []string{"value"})

	lambdaGreaterThan = types.NewLambda(func(this *types.JsonLambda, args []types.JsonType) (types.JsonType, error) {
		if types.AsNumber(args[0]).IntValue() > 1024 {
			return types.True(), nil
		}
		return types.False(), nil
	}, "value => value > 1024", nil, []string{"value"})

	lambdaGreaterThanIndexed = types.NewLambda(func(this *types.JsonLambda, args []types.JsonType) (types.JsonType, error) {
		if types.AsNumber(args[0]).IntValue() > 1024 {
			return types.True(), nil
		}
		return types.False(), nil
	}, "(value, index) => value > 1024", nil, []string{"value", "index"})
)

var arrayBenchResult types.JsonType

func buildSequentialArray(size int) *types.JsonArray {
	values := make([]types.JsonType, size)
	for i := 0; i < size; i++ {
		values[i] = types.AsNumber(i)
	}
	return &types.JsonArray{Value: values}
}

func runArrayInstanceBenchmark(b *testing.B, name string, array *types.JsonArray, args ...types.JsonType) {
	b.Helper()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result, err := functions.CallInstanceFunction(name, array, args, arrayBenchmarkResolve)
		if err != nil {
			b.Fatal(err)
		}
		arrayBenchResult = result
	}
}

func BenchmarkArrayFilter(b *testing.B) {
	runArrayInstanceBenchmark(b, "filter", arrayNumbers, lambdaEven)
}

func BenchmarkArrayFilterIndexed(b *testing.B) {
	runArrayInstanceBenchmark(b, "filter", arrayNumbers, lambdaEvenIndexed)
}

func BenchmarkArrayMap(b *testing.B) {
	runArrayInstanceBenchmark(b, "map", arrayNumbers, lambdaDouble)
}

func BenchmarkArrayMapIndexed(b *testing.B) {
	runArrayInstanceBenchmark(b, "map", arrayNumbers, lambdaDoubleIndexed)
}

func BenchmarkArrayFlatMap(b *testing.B) {
	runArrayInstanceBenchmark(b, "flatMap", arrayNumbers, lambdaFlatMap)
}

func BenchmarkArrayReduce(b *testing.B) {
	runArrayInstanceBenchmark(b, "reduce", arrayNumbers, lambdaAccumulator)
}

func BenchmarkArrayReduceInit(b *testing.B) {
	runArrayInstanceBenchmark(b, "reduce", arrayNumbers, lambdaAccumulatorInit, types.AsNumber(0))
}

func BenchmarkArraySortBy(b *testing.B) {
	runArrayInstanceBenchmark(b, "sort", arrayNumbers, lambdaDescendingKey)
}

func BenchmarkArrayCount(b *testing.B) {
	runArrayInstanceBenchmark(b, "count", arrayNumbers, lambdaGreaterThan)
}

func BenchmarkArrayCountIndexed(b *testing.B) {
	runArrayInstanceBenchmark(b, "count", arrayNumbers, lambdaGreaterThanIndexed)
}

func BenchmarkArrayAny(b *testing.B) {
	runArrayInstanceBenchmark(b, "any", arrayNumbers, lambdaGreaterThan)
}

func BenchmarkArrayAll(b *testing.B) {
	runArrayInstanceBenchmark(b, "all", arrayNumbers, lambdaGreaterThan)
}

func BenchmarkArrayNone(b *testing.B) {
	runArrayInstanceBenchmark(b, "none", arrayNumbers, lambdaGreaterThan)
}

func BenchmarkArrayFindFirst(b *testing.B) {
	runArrayInstanceBenchmark(b, "findFirst", arrayNumbers, lambdaGreaterThan)
}

func BenchmarkArrayFindLast(b *testing.B) {
	runArrayInstanceBenchmark(b, "findLast", arrayNumbers, lambdaGreaterThan)
}

func BenchmarkArrayMaxBy(b *testing.B) {
	runArrayInstanceBenchmark(b, "max", arrayNumbers, lambdaSquareKey)
}

func BenchmarkArrayMinBy(b *testing.B) {
	runArrayInstanceBenchmark(b, "min", arrayNumbers, lambdaSquareKey)
}

func BenchmarkArraySumBy(b *testing.B) {
	runArrayInstanceBenchmark(b, "sum", arrayNumbers, lambdaDouble)
}

func BenchmarkArraySortMap(b *testing.B) {
	runArrayInstanceBenchmark(b, "sort", arrayNumbers, lambdaSquareKey)
}

func BenchmarkArrayFlatMapNested(b *testing.B) {
	runArrayInstanceBenchmark(b, "flatMap", arrayNested, lambdaFlatMap)
}
