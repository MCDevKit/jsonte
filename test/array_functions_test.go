package test

import (
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"math/rand"
	"testing"
)

func TestMap(t *testing.T) {
	eval := evaluate(t, `(1..10).map(x => x * 2)`)
	assertArray(t, eval, types.AsArray([]interface{}{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}))
}

func TestFilter(t *testing.T) {
	eval := evaluate(t, `(1..10).filter(x => mod(x, 2) == 0)`)
	assertArray(t, eval, types.Box([]interface{}{2, 4, 6, 8, 10}).(*types.JsonArray))
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
	assertArray(t, eval, types.Box([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).(*types.JsonArray))
}

func TestSortWithSelector(t *testing.T) {
	eval := evaluate(t, `(1..10).sort(x => x * -1)`)
	assertArray(t, eval, types.Box([]interface{}{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}).(*types.JsonArray))
}

func TestReverse(t *testing.T) {
	eval := evaluate(t, `(1..10).reverse()`)
	assertArray(t, eval, types.Box([]interface{}{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}).(*types.JsonArray))
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
	assertArray(t, eval, types.Box([]interface{}{"a", "b"}).(*types.JsonArray))
}

func TestValues(t *testing.T) {
	// Needs to be sorted, because maps are unordered
	eval := evaluate(t, `values({'a': 1, 'b': 2}).sort()`)
	assertArray(t, eval, types.Box([]interface{}{types.AsNumber(1), types.AsNumber(2)}).(*types.JsonArray))
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
	assertArray(t, eval, types.Box([]interface{}{
		utils.ToNavigableMap("key", "a", "value", types.AsNumber(1)),
		utils.ToNavigableMap("key", "b", "value", types.AsNumber(2)),
	}).(*types.JsonArray))
}

func TestFlatMap(t *testing.T) {
	eval := evaluate(t, `flatMap([1, 2, 3], x => [x, x * 2])`)
	assertArray(t, eval, types.Box([]interface{}{1, 2, 2, 4, 3, 6}).(*types.JsonArray))
}

func TestArrayRange(t *testing.T) {
	eval := evaluate(t, `(1..5).range()`)
	assertArray(t, eval, types.Box([]interface{}{0, 1, 2, 3, 4}).(*types.JsonArray))
}

func TestSublist(t *testing.T) {
	eval := evaluate(t, `(1..5).sublist(1, 3)`)
	assertArray(t, eval, types.Box([]interface{}{2, 3}).(*types.JsonArray))
}

func TestRandomElement(t *testing.T) {
	rand.Seed(0)
	eval := evaluate(t, `(1..5).random()`)
	assertNumber(t, eval, 5)
}

func TestAppendElement(t *testing.T) {
	eval := evaluate(t, `[1, 2, 3].append(4)`)
	assertArray(t, eval, types.Box([]interface{}{1, 2, 3, 4}).(*types.JsonArray))
}

func TestAppendElements(t *testing.T) {
	eval := evaluate(t, `[1, 2, 3].append(4, 5)`)
	assertArray(t, eval, types.Box([]interface{}{1, 2, 3, 4, 5}).(*types.JsonArray))
}

func TestAppendArray(t *testing.T) {
	eval := evaluate(t, `[1, 2, 3].append([4, 5])`)
	assertArray(t, eval, types.Box([]interface{}{1, 2, 3, []interface{}{4, 5}}).(*types.JsonArray))
}

func TestAppendSpreadArray(t *testing.T) {
	eval := evaluate(t, `[1, 2, 3].append(...[4, 5])`)
	assertArray(t, eval, types.Box([]interface{}{1, 2, 3, 4, 5}).(*types.JsonArray))
}

func TestPrependElement(t *testing.T) {
	eval := evaluate(t, `[1, 2, 3].prepend(4)`)
	assertArray(t, eval, types.Box([]interface{}{4, 1, 2, 3}).(*types.JsonArray))
}

func TestPrependElements(t *testing.T) {
	eval := evaluate(t, `[1, 2, 3].prepend(4, 5)`)
	assertArray(t, eval, types.Box([]interface{}{4, 5, 1, 2, 3}).(*types.JsonArray))
}

func TestPrependArray(t *testing.T) {
	eval := evaluate(t, `[1, 2, 3].prepend([4, 5])`)
	assertArray(t, eval, types.Box([]interface{}{[]interface{}{4, 5}, 1, 2, 3}).(*types.JsonArray))
}

func TestPrependSpreadArray(t *testing.T) {
	eval := evaluate(t, `[1, 2, 3].prepend(...[4, 5])`)
	assertArray(t, eval, types.Box([]interface{}{4, 5, 1, 2, 3}).(*types.JsonArray))
}

func TestRemoveElement(t *testing.T) {
	eval := evaluate(t, `[1, 2, 3].remove(1)`)
	assertArray(t, eval, types.Box([]interface{}{1, 3}).(*types.JsonArray))
}

func TestRemoveFront(t *testing.T) {
	eval := evaluate(t, `[1, 2, 3].removeFront()`)
	assertArray(t, eval, types.Box([]interface{}{2, 3}).(*types.JsonArray))
}

func TestRemoveBack(t *testing.T) {
	eval := evaluate(t, `[1, 2, 3].removeBack()`)
	assertArray(t, eval, types.Box([]interface{}{1, 2}).(*types.JsonArray))
}

func TestRemoveElementWithPredicate(t *testing.T) {
	eval := evaluate(t, `[1, 2, 3].remove(x => x == 1)`)
	assertArray(t, eval, types.Box([]interface{}{2, 3}).(*types.JsonArray))
}
