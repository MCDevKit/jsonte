package functions

import (
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"math"
	"sort"
)

func RegisterArrayFunctions() {
	RegisterFunction(JsonFunction{
		Name: "asArray",
		Body: asArray,
	})
	RegisterFunction(JsonFunction{
		Name: "keys",
		Body: keys,
	})
	RegisterFunction(JsonFunction{
		Name: "values",
		Body: values,
	})
	RegisterFunction(JsonFunction{
		Name:       "sort",
		Body:       sort_,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "sort",
		Body:       sortMap,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "reverse",
		Body:       reverse,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "contains",
		Body:       arrayContains,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "indexOf",
		Body:       arrayIndexOf,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "lastIndexOf",
		Body:       arrayLastIndexOf,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "any",
		Body:       any_,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "all",
		Body:       all,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "count",
		Body:       count,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "count",
		Body:       countFilter,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "sum",
		Body:       sum,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "sum",
		Body:       sumMap,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "none",
		Body:       none,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "filter",
		Body:       filter,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "map",
		Body:       map_,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "flatMap",
		Body:       flatMap,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "flatMap",
		Body:       flatMapSimple,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "range",
		Body:       range_,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "findFirst",
		Body:       findFirst,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "findFirst",
		Body:       findFirstFilter,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "encode",
		Body:       encode,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "encode",
		Body:       encodeSimple,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "sublist",
		Body:       sublist,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "sublist",
		Body:       sublistStart,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "max",
		Body:       maxArray,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "max",
		Body:       maxArraySimple,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "min",
		Body:       minArray,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "min",
		Body:       minArraySimple,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "reduce",
		Body:       reduce,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "reduce",
		Body:       reduceInit,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "findLast",
		Body:       findLast,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "findLast",
		Body:       findLastFilter,
		IsInstance: true,
	})
}

func asArray(obj utils.JsonObject, key, value string) utils.JsonArray {
	if obj == nil {
		return nil
	}
	arr := make(utils.JsonArray, len(obj))
	i := 0
	for k, v := range obj {
		arr[i] = utils.JsonObject{key: k, value: v}
		i++
	}
	return arr
}

func keys(obj utils.JsonObject) utils.JsonArray {
	if obj == nil {
		return nil
	}
	arr := make(utils.JsonArray, len(obj))
	i := 0
	for k := range obj {
		arr[i] = k
		i++
	}
	return arr
}

func values(obj utils.JsonObject) utils.JsonArray {
	if obj == nil {
		return nil
	}
	arr := make(utils.JsonArray, len(obj))
	i := 0
	for _, v := range obj {
		arr[i] = v
		i++
	}
	return arr
}

func reverse(arr utils.JsonArray) utils.JsonArray {
	if arr == nil {
		return nil
	}
	rev := make(utils.JsonArray, len(arr))
	i := 0
	for j := len(arr) - 1; j >= 0; j-- {
		rev[i] = arr[j]
		i++
	}
	return rev
}

func sort_(arr utils.JsonArray) utils.JsonArray {
	if arr == nil {
		return nil
	}
	result := make(utils.JsonArray, len(arr))
	copy(result, arr)
	sort.SliceStable(result, func(i, j int) bool {
		return utils.Less(result[i], result[j])
	})
	return result
}

func sortMap(arr utils.JsonArray, predicate utils.JsonLambda) (utils.JsonArray, error) {
	if arr == nil {
		return nil, nil
	}
	mapped, err := map_(arr, predicate)
	if err != nil {
		return nil, utils.WrapErrorf(err, "An error occurred while mapping the values for sorting the array")
	}
	result := make(utils.JsonArray, len(arr))
	copy(result, arr)
	sort.SliceStable(result, func(i, j int) bool {
		return utils.Less(mapped[i], mapped[j])
	})
	return result, nil
}

func arrayContains(arr utils.JsonArray, value interface{}) bool {
	if arr == nil {
		return false
	}
	for _, v := range arr {
		if utils.IsEqual(v, value) {
			return true
		}
	}
	return false
}

func arrayIndexOf(arr utils.JsonArray, value interface{}) utils.JsonNumber {
	if arr == nil {
		return utils.ToNumber(-1)
	}
	for i, v := range arr {
		if utils.IsEqual(v, value) {
			return utils.ToNumber(i)
		}
	}
	return utils.ToNumber(-1)
}

func arrayLastIndexOf(arr utils.JsonArray, value interface{}) utils.JsonNumber {
	if arr == nil {
		return utils.ToNumber(-1)
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if utils.IsEqual(arr[i], value) {
			return utils.ToNumber(i)
		}
	}
	return utils.ToNumber(-1)
}

func any_(arr utils.JsonArray, predicate utils.JsonLambda) (bool, error) {
	if arr == nil {
		return false, nil
	}
	for i, v := range arr {
		b, err := predicate([]interface{}{v, i})
		if err != nil {
			return false, utils.WrapErrorf(err, "An error occurred while evaluating the predicate for `any` at index %d", i)
		}
		if utils.ToBoolean(b) {
			return true, nil
		}
	}
	return false, nil
}

func all(arr utils.JsonArray, predicate utils.JsonLambda) (bool, error) {
	if arr == nil {
		return false, nil
	}
	for i, v := range arr {
		b, err := predicate([]interface{}{v, i})
		if err != nil {
			return false, utils.WrapErrorf(err, "An error occurred while evaluating the predicate for `all` at index %d", i)
		}
		if !utils.ToBoolean(b) {
			return false, nil
		}
	}
	return true, nil
}

func none(arr utils.JsonArray, predicate utils.JsonLambda) (bool, error) {
	if arr == nil {
		return true, nil
	}
	for i, v := range arr {
		b, err := predicate([]interface{}{v, i})
		if err != nil {
			return false, utils.WrapErrorf(err, "An error occurred while evaluating the predicate for `none` at index %d", i)
		}
		if utils.ToBoolean(b) {
			return false, nil
		}
	}
	return true, nil
}

func filter(arr utils.JsonArray, predicate utils.JsonLambda) (utils.JsonArray, error) {
	if arr == nil {
		return nil, nil
	}
	result := make(utils.JsonArray, 0)
	for i, v := range arr {
		b, err := predicate([]interface{}{v, i})
		if err != nil {
			return nil, utils.WrapErrorf(err, "An error occurred while evaluating the predicate for `filter` at index %d", i)
		}
		if utils.ToBoolean(b) {
			result = append(result, v)
		}
	}
	return result, nil
}

func map_(arr utils.JsonArray, predicate utils.JsonLambda) (utils.JsonArray, error) {
	if arr == nil {
		return nil, nil
	}
	result := make(utils.JsonArray, len(arr))
	for i, v := range arr {
		b, err := predicate([]interface{}{v, i})
		if err != nil {
			return nil, utils.WrapErrorf(err, "An error occurred while evaluating the predicate for `map` at index %d", i)
		}
		result[i] = b
	}
	return result, nil
}

func flatMap(arr utils.JsonArray, predicate utils.JsonLambda) (utils.JsonArray, error) {
	if arr == nil {
		return nil, nil
	}
	result := make(utils.JsonArray, 0)
	for i, v := range arr {
		b, err := predicate([]interface{}{v, i})
		if err != nil {
			return nil, utils.WrapErrorf(err, "An error occurred while evaluating the predicate for `flatMap` at index %d", i)
		}
		if arr, ok := b.(utils.JsonArray); ok {
			result = append(result, arr...)
		} else {
			result = append(result, b)
		}
	}
	return result, nil
}

func flatMapSimple(arr utils.JsonArray) utils.JsonArray {
	if arr == nil {
		return nil
	}
	result := make(utils.JsonArray, 0)
	for _, v := range arr {
		if arr, ok := v.(utils.JsonArray); ok {
			result = append(result, arr...)
		} else {
			result = append(result, v)
		}
	}
	return result
}

func countFilter(arr utils.JsonArray, predicate utils.JsonLambda) (utils.JsonNumber, error) {
	if arr == nil {
		return utils.ToNumber(0), nil
	}
	count := 0
	for i, v := range arr {
		b, err := predicate([]interface{}{v, i})
		if err != nil {
			return utils.ToNumber(0), utils.WrapErrorf(err, "An error occurred while evaluating the predicate for `count` at index %d", i)
		}
		if utils.ToBoolean(b) {
			count++
		}
	}
	return utils.ToNumber(count), nil
}

func count(arr utils.JsonArray) utils.JsonNumber {
	if arr == nil {
		return utils.ToNumber(0)
	}
	return utils.ToNumber(len(arr))
}

func range_(arr utils.JsonArray) utils.JsonArray {
	if arr == nil {
		return nil
	}
	return utils.CreateRange(0, len(arr)-1)
}

func findFirstFilter(arr utils.JsonArray, predicate utils.JsonLambda) (interface{}, error) {
	if arr == nil {
		return nil, nil
	}
	for i, v := range arr {
		b, err := predicate([]interface{}{v, i})
		if err != nil {
			return nil, utils.WrapErrorf(err, "An error occurred while evaluating the predicate for `findFirst` at index %d", i)
		}
		if utils.ToBoolean(b) {
			return v, nil
		}
	}

	return nil, utils.WrappedError("No matching items found!")
}

func findFirst(arr utils.JsonArray) (interface{}, error) {
	if arr == nil {
		return nil, nil
	}
	if len(arr) == 0 {
		return nil, utils.WrappedError("No matching items found!")
	}
	return arr[0], nil
}

func encode(arr utils.JsonArray, space utils.JsonNumber, predicate utils.JsonLambda) (utils.JsonNumber, error) {
	if space.IntValue() <= 0 || (space.IntValue()&(space.IntValue()-1)) != 0 {
		return utils.ToNumber(0), utils.WrappedError("Space must be a power of 2 and greater than 0!")
	}
	if arr == nil {
		return utils.ToNumber(0), nil
	}
	result := 0
	bitsPerElement := int(math.Log(float64(space.IntValue())) / math.Log(2))
	for i := 0; i < int(math.Min(float64(len(arr)), float64(32/bitsPerElement))); i++ {
		i2, err := predicate([]interface{}{arr[i], i})
		if err != nil {
			return utils.ToNumber(0), utils.WrapErrorf(err, "An error occurred while evaluating the predicate for `encode` at index %d", i)
		}
		if !utils.IsNumber(i2) {
			return utils.ToNumber(0), utils.WrappedError("Predicate must return a number!")
		}
		number := utils.ToNumber(i2)
		if number.IntValue() < 0 || number.IntValue() >= space.IntValue() {
			return utils.ToNumber(0), utils.WrappedErrorf("Number %d is out of range 0..%d", number.IntValue(), space.IntValue()-1)
		}
		result += number.IntValue() << (i * bitsPerElement)
	}
	return utils.ToNumber(result), nil
}

func encodeSimple(arr utils.JsonArray, space utils.JsonNumber) (utils.JsonNumber, error) {
	return encode(arr, space, func(args []interface{}) (interface{}, error) {
		return args[0], nil
	})
}

func sublist(arr utils.JsonArray, start utils.JsonNumber, end utils.JsonNumber) utils.JsonArray {
	if arr == nil {
		return nil
	}
	startIndex := start.IntValue()
	endIndex := end.IntValue()
	if startIndex < 0 {
		startIndex = 0
	}
	if endIndex > len(arr) {
		endIndex = len(arr)
	}
	return arr[startIndex:endIndex]
}

func sublistStart(arr utils.JsonArray, start utils.JsonNumber) utils.JsonArray {
	if arr == nil {
		return nil
	}
	startIndex := start.IntValue()
	if startIndex < 0 {
		startIndex = 0
	}
	return arr[startIndex:]
}

func maxArray(arr utils.JsonArray, predicate utils.JsonLambda) (interface{}, error) {
	if arr == nil {
		return nil, nil
	}
	max, err := predicate([]interface{}{arr[0], 0})
	if err != nil {
		return nil, utils.WrapErrorf(err, "An error occurred while evaluating the predicate for `max` at index %d", 0)
	}
	for i, v := range arr {
		b, err := predicate([]interface{}{v, i})
		if err != nil {
			return nil, utils.WrapErrorf(err, "An error occurred while evaluating the predicate for `max` at index %d", i)
		}
		if utils.ToNumber(b).FloatValue() > utils.ToNumber(max).FloatValue() {
			max = b
		}
	}
	return max, nil
}

func maxArraySimple(arr utils.JsonArray) interface{} {
	if arr == nil {
		return nil
	}
	max := arr[0]
	for _, v := range arr {
		if utils.ToNumber(v).FloatValue() > utils.ToNumber(max).FloatValue() {
			max = v
		}
	}
	return max
}

func minArray(arr utils.JsonArray, predicate utils.JsonLambda) (interface{}, error) {
	if arr == nil || len(arr) == 0 {
		return nil, nil
	}
	min, err := predicate([]interface{}{arr[0], 0})
	if err != nil {
		return nil, utils.WrapErrorf(err, "An error occurred while evaluating the predicate for `min` at index %d", 0)
	}
	for i, v := range arr {
		if i == 0 {
			continue
		}
		b, err := predicate([]interface{}{v, i})
		if err != nil {
			return nil, utils.WrapErrorf(err, "An error occurred while evaluating the predicate for `min` at index %d", i)
		}
		if utils.ToNumber(b).FloatValue() < utils.ToNumber(min).FloatValue() {
			min = b
		}
	}
	return min, nil
}

func minArraySimple(arr utils.JsonArray) interface{} {
	if arr == nil {
		return nil
	}
	min := arr[0]
	for _, v := range arr {
		if utils.ToNumber(v).FloatValue() < utils.ToNumber(min).FloatValue() {
			min = v
		}
	}
	return min
}

func sumMap(arr utils.JsonArray, predicate utils.JsonLambda) (utils.JsonNumber, error) {
	if arr == nil {
		return utils.ToNumber(0), nil
	}
	s := 0.0
	for i, v := range arr {
		b, err := predicate([]interface{}{v, i})
		if err != nil {
			return utils.ToNumber(0), utils.WrapErrorf(err, "An error occurred while evaluating the predicate for `sum` at index %d", i)
		}
		s = s + utils.ToNumber(b).FloatValue()
	}
	return utils.ToNumber(s), nil
}

func sum(arr utils.JsonArray) utils.JsonNumber {
	if arr == nil {
		return utils.ToNumber(0)
	}
	s := 0.0
	for _, v := range arr {
		s = s + utils.ToNumber(v).FloatValue()
	}
	return utils.ToNumber(s)
}

func reduce(arr utils.JsonArray, predicate utils.JsonLambda) (interface{}, error) {
	return reduceInit(arr, predicate, nil)
}

func reduceInit(arr utils.JsonArray, predicate utils.JsonLambda, initialValue interface{}) (interface{}, error) {
	if arr == nil {
		return nil, nil
	}
	var prev = initialValue
	for i, v := range arr {
		b, err := predicate([]interface{}{prev, v, i})
		if err != nil {
			return utils.ToNumber(0), utils.WrapErrorf(err, "An error occurred while evaluating the predicate for `reduce` at index %d", i)
		}
		prev = b
	}
	return prev, nil
}

func findLastFilter(arr utils.JsonArray, predicate utils.JsonLambda) (interface{}, error) {
	if arr == nil {
		return nil, nil
	}
	for i := len(arr) - 1; i >= 0; i-- {
		b, err := predicate([]interface{}{arr[i], i})
		if err != nil {
			return nil, utils.WrapErrorf(err, "An error occurred while evaluating the predicate for `findLast` at index %d", i)
		}
		if utils.ToBoolean(b) {
			return arr[i], nil
		}
	}
	return nil, utils.WrappedError("No matching items found!")
}

func findLast(arr utils.JsonArray) (interface{}, error) {
	if arr == nil {
		return nil, nil
	}
	if len(arr) == 0 {
		return nil, utils.WrappedError("No matching items found!")
	}
	return arr[len(arr)-1], nil
}
