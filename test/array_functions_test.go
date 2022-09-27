package test

import (
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"testing"
)

func TestMap(t *testing.T) {
	eval := evaluate(t, `(1..10).map(x => x * 2)`)
	assertArray(t, eval, []interface{}{2, 4, 6, 8, 10, 12, 14, 16, 18, 20})
}

func TestFilter(t *testing.T) {
	eval := evaluate(t, `(1..10).filter(x => mod(x, 2) == 0)`)
	assertArray(t, eval, []interface{}{2, 4, 6, 8, 10})
}

func TestReduce(t *testing.T) {
	eval := evaluate(t, `(1..10).reduce((acc, x) => acc + x)`)
	assertNumber(t, eval, 55)
}

func TestReduceWithInitialValue(t *testing.T) {
	eval := evaluate(t, `(1..10).reduce((acc, x) => acc + x, 100)`)
	assertNumber(t, eval, 155)
}

func TestAny(t *testing.T) {
	eval := evaluate(t, `(1..10).any(x => x == 5)`)
	assertBool(t, eval, true)
}

func TestAll(t *testing.T) {
	eval := evaluate(t, `(1..10).all(x => x < 11)`)
	assertBool(t, eval, true)
}

func TestNone(t *testing.T) {
	eval := evaluate(t, `(1..10).none(x => x == 11)`)
	assertBool(t, eval, true)
}

func TestCount(t *testing.T) {
	eval := evaluate(t, `(1..10).count(x => x < 5)`)
	assertNumber(t, eval, 4)
}

func TestCountWithoutPredicate(t *testing.T) {
	eval := evaluate(t, `(1..10).count()`)
	assertNumber(t, eval, 10)
}

func TestFirst(t *testing.T) {
	eval := evaluate(t, `(1..10).findFirst()`)
	assertNumber(t, eval, 1)
}

func TestFirstWithPredicate(t *testing.T) {
	eval := evaluate(t, `(1..10).findFirst(x => x > 5)`)
	assertNumber(t, eval, 6)
}

func TestLast(t *testing.T) {
	eval := evaluate(t, `(1..10).findLast()`)
	assertNumber(t, eval, 10)
}

func TestLastWithPredicate(t *testing.T) {
	eval := evaluate(t, `(1..10).findLast(x => x < 5)`)
	assertNumber(t, eval, 4)
}

func TestMinArray(t *testing.T) {
	eval := evaluate(t, `(1..10).min()`)
	assertNumber(t, eval, 1)
}

func TestMinArrayWithSelector(t *testing.T) {
	eval := evaluate(t, `(1..10).min(x => x * 2)`)
	assertNumber(t, eval, 2)
}

func TestMaxArray(t *testing.T) {
	eval := evaluate(t, `(1..10).max()`)
	assertNumber(t, eval, 10)
}

func TestMaxArrayWithSelector(t *testing.T) {
	eval := evaluate(t, `(1..10).max(x => x * 2)`)
	assertNumber(t, eval, 20)
}

func TestSumArray(t *testing.T) {
	eval := evaluate(t, `(1..10).sum()`)
	assertNumber(t, eval, 55)
}

func TestSumArrayWithSelector(t *testing.T) {
	eval := evaluate(t, `(1..10).sum(x => x * 2)`)
	assertNumber(t, eval, 110)
}

func TestSort(t *testing.T) {
	eval := evaluate(t, `(1..10).sort()`)
	assertArray(t, eval, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}

func TestSortWithSelector(t *testing.T) {
	eval := evaluate(t, `(1..10).sort(x => x * -1)`)
	assertArray(t, eval, []interface{}{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
}

func TestReverse(t *testing.T) {
	eval := evaluate(t, `(1..10).reverse()`)
	assertArray(t, eval, []interface{}{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
}

func TestContains(t *testing.T) {
	eval := evaluate(t, `(1..10).contains(5)`)
	assertBool(t, eval, true)
}

func TestArrayLastIndexOf(t *testing.T) {
	eval := evaluate(t, `(1..10).lastIndexOf(5)`)
	assertNumber(t, eval, 4)
}

func TestArrayIndexOf(t *testing.T) {
	eval := evaluate(t, `(1..10).indexOf(5)`)
	assertNumber(t, eval, 4)
}

func TestKeys(t *testing.T) {
	// Needs to be sorted, because maps are unordered
	eval := evaluate(t, `keys({'a': 1, 'b': 2}).sort()`)
	assertArray(t, eval, []interface{}{"a", "b"})
}

func TestValues(t *testing.T) {
	// Needs to be sorted, because maps are unordered
	eval := evaluate(t, `values({'a': 1, 'b': 2}).sort()`)
	assertArray(t, eval, []interface{}{utils.ToNumber(1), utils.ToNumber(2)})
}

func TestArrayJoin(t *testing.T) {
	eval := evaluate(t, `['a', 'b', 'c'].join(', ')`)
	assertString(t, eval, "a, b, c")
}

func TestArrayJoinWithEmptyArray(t *testing.T) {
	eval := evaluate(t, `[].join(', ')`)
	assertString(t, eval, "")
}

func TestArrayJoinWithEmptySeparator(t *testing.T) {
	eval := evaluate(t, `['a', 'b', 'c'].join('')`)
	assertString(t, eval, "abc")
}

func TestAsArray(t *testing.T) {
	eval := evaluate(t, `asArray({'a': 1, 'b': 2}, 'key', 'value').sort(x => x.key)`)
	assertArray(t, eval, []interface{}{
		utils.ToNavigableMap("key", "a", "value", utils.ToNumber(1)),
		utils.ToNavigableMap("key", "b", "value", utils.ToNumber(2)),
	})
}

func TestFlatMap(t *testing.T) {
	eval := evaluate(t, `flatMap([1, 2, 3], x => [x, x * 2])`)
	assertArray(t, eval, []interface{}{1, 2, 2, 4, 3, 6})
}

func TestArrayRange(t *testing.T) {
	eval := evaluate(t, `(1..5).range()`)
	assertArray(t, eval, []interface{}{0, 1, 2, 3, 4})
}

func TestSublist(t *testing.T) {
	eval := evaluate(t, `(1..5).sublist(1, 3)`)
	assertArray(t, eval, []interface{}{2, 3})
}
