package functions

import (
	"fmt"
	"jsonte/jsonte/utils"
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

func sort_(arr utils.JsonArray) (utils.JsonArray, error) {
	if arr == nil {
		return nil, nil
	}
	result := make(utils.JsonArray, len(arr))
	copy(result, arr)
	sort.SliceStable(result, func(i, j int) bool {
		return utils.Less(result[i], result[j])
	})
	return result, nil
}

func sortMap(arr utils.JsonArray, predicate utils.JsonLambda) (utils.JsonArray, error) {
	if arr == nil {
		return nil, nil
	}
	mapped, err := map_(arr, predicate)
	if err != nil {
		return nil, err
	}
	result := make(utils.JsonArray, len(arr))
	copy(result, arr)
	sort.SliceStable(result, func(i, j int) bool {
		return utils.Less(mapped[i], mapped[j])
	})
	return result, nil
}

func arrayContains(arr utils.JsonArray, value interface{}) (bool, error) {
	if arr == nil {
		return false, nil
	}
	for _, v := range arr {
		if utils.IsEqual(v, value) {
			return true, nil
		}
	}
	return false, nil
}

func arrayIndexOf(arr utils.JsonArray, value interface{}) (utils.JsonNumber, error) {
	if arr == nil {
		return utils.ToNumber(-1), nil
	}
	for i, v := range arr {
		if utils.IsEqual(v, value) {
			return utils.ToNumber(i), nil
		}
	}
	return utils.ToNumber(-1), nil
}

func arrayLastIndexOf(arr utils.JsonArray, value interface{}) (utils.JsonNumber, error) {
	if arr == nil {
		return utils.ToNumber(-1), nil
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if utils.IsEqual(arr[i], value) {
			return utils.ToNumber(i), nil
		}
	}
	return utils.ToNumber(-1), nil
}

func any_(arr utils.JsonArray, predicate utils.JsonLambda) (bool, error) {
	if arr == nil {
		return false, nil
	}
	for i, v := range arr {
		b, err := predicate([]interface{}{v, i})
		if err != nil {
			return false, err
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
			return false, err
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
			return false, err
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
			return nil, err
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
			return nil, err
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
			return nil, err
		}
		if arr, ok := b.(utils.JsonArray); ok {
			result = append(result, arr...)
		} else {
			result = append(result, b)
		}
	}
	return result, nil
}

func flatMapSimple(arr utils.JsonArray) (utils.JsonArray, error) {
	if arr == nil {
		return nil, nil
	}
	result := make(utils.JsonArray, 0)
	for _, v := range arr {
		if arr, ok := v.(utils.JsonArray); ok {
			result = append(result, arr...)
		} else {
			result = append(result, v)
		}
	}
	return result, nil
}

func countFilter(arr utils.JsonArray, predicate utils.JsonLambda) (utils.JsonNumber, error) {
	if arr == nil {
		return utils.ToNumber(0), nil
	}
	count := 0
	for i, v := range arr {
		b, err := predicate([]interface{}{v, i})
		if err != nil {
			return utils.ToNumber(0), err
		}
		if utils.ToBoolean(b) {
			count++
		}
	}
	return utils.ToNumber(count), nil
}

func count(arr utils.JsonArray) (utils.JsonNumber, error) {
	if arr == nil {
		return utils.ToNumber(0), nil
	}
	return utils.ToNumber(len(arr)), nil
}

func range_(arr utils.JsonArray) (utils.JsonArray, error) {
	if arr == nil {
		return nil, nil
	}
	return utils.CreateRange(0, len(arr)-1), nil
}

func findFirstFilter(arr utils.JsonArray, predicate utils.JsonLambda) (interface{}, error) {
	if arr == nil {
		return nil, nil
	}
	for i, v := range arr {
		b, err := predicate([]interface{}{v, i})
		if err != nil {
			return nil, err
		}
		if utils.ToBoolean(b) {
			return v, nil
		}
	}
	return nil, &utils.EvaluationError{
		Message: "No matching items found!",
	}
}

func findFirst(arr utils.JsonArray) (interface{}, error) {
	if arr == nil {
		return nil, nil
	}
	if len(arr) == 0 {
		return nil, &utils.EvaluationError{
			Message: "No matching items found!",
		}
	}
	return arr[0], nil
}

func encode(arr utils.JsonArray, space utils.JsonNumber, predicate utils.JsonLambda) (utils.JsonNumber, error) {
	if space.IntValue() <= 0 || (space.IntValue()&(space.IntValue()-1)) != 0 {
		return utils.ToNumber(0), &utils.EvaluationError{
			Message: "Space must be a power of 2 and greater than 0!",
		}
	}
	if arr == nil {
		return utils.ToNumber(0), nil
	}
	result := 0
	bitsPerElement := int(math.Log(float64(space.IntValue())) / math.Log(2))
	for i := 0; i < int(math.Min(float64(len(arr)), float64(32/bitsPerElement))); i++ {
		i2, err := predicate([]interface{}{arr[i], i})
		if err != nil {
			return utils.ToNumber(0), err
		}
		if !utils.IsNumber(i2) {
			return utils.ToNumber(0), &utils.EvaluationError{
				Message: "Predicate must return a number!",
			}
		}
		number := utils.ToNumber(i2)
		if number.IntValue() < 0 || number.IntValue() >= space.IntValue() {
			return utils.ToNumber(0), &utils.EvaluationError{
				Message: fmt.Sprintf("Number %s is out of range 0..%d", number.StringValue(), space.IntValue()-1),
			}
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

func sublist(arr utils.JsonArray, start utils.JsonNumber, end utils.JsonNumber) (utils.JsonArray, error) {
	if arr == nil {
		return nil, nil
	}
	startIndex := start.IntValue()
	endIndex := end.IntValue()
	if startIndex < 0 {
		startIndex = 0
	}
	if endIndex > len(arr) {
		endIndex = len(arr)
	}
	return arr[startIndex:endIndex], nil
}

func sublistStart(arr utils.JsonArray, start utils.JsonNumber) (utils.JsonArray, error) {
	if arr == nil {
		return nil, nil
	}
	startIndex := start.IntValue()
	if startIndex < 0 {
		startIndex = 0
	}
	return arr[startIndex:], nil
}

func maxArray(arr utils.JsonArray, predicate utils.JsonLambda) (interface{}, error) {
	if arr == nil {
		return nil, nil
	}
	max := arr[0]
	for i, v := range arr {
		b, err := predicate([]interface{}{v, i})
		if err != nil {
			return nil, err
		}
		if utils.ToNumber(b).FloatValue() > utils.ToNumber(max).FloatValue() {
			max = b
		}
	}
	return max, nil
}

func maxArraySimple(arr utils.JsonArray) (interface{}, error) {
	if arr == nil {
		return nil, nil
	}
	max := arr[0]
	for _, v := range arr {
		if utils.ToNumber(v).FloatValue() > utils.ToNumber(max).FloatValue() {
			max = v
		}
	}
	return max, nil
}

func minArray(arr utils.JsonArray, predicate utils.JsonLambda) (interface{}, error) {
	if arr == nil || len(arr) == 0 {
		return nil, nil
	}
	min, err := predicate([]interface{}{arr[0], 0})
	if err != nil {
		return nil, err
	}
	for i, v := range arr {
		if i == 0 {
			continue
		}
		b, err := predicate([]interface{}{v, i})
		if err != nil {
			return nil, err
		}
		if utils.ToNumber(b).FloatValue() < utils.ToNumber(min).FloatValue() {
			min = b
		}
	}
	return min, nil
}

func minArraySimple(arr utils.JsonArray) (interface{}, error) {
	if arr == nil {
		return nil, nil
	}
	min := arr[0]
	for _, v := range arr {
		if utils.ToNumber(v).FloatValue() < utils.ToNumber(min).FloatValue() {
			min = v
		}
	}
	return min, nil
}

func sumMap(arr utils.JsonArray, predicate utils.JsonLambda) (utils.JsonNumber, error) {
	if arr == nil {
		return utils.ToNumber(0), nil
	}
	s := 0.0
	for i, v := range arr {
		b, err := predicate([]interface{}{v, i})
		if err != nil {
			return utils.ToNumber(0), err
		}
		s = s + utils.ToNumber(b).FloatValue()
	}
	return utils.ToNumber(s), nil
}

func sum(arr utils.JsonArray) (utils.JsonNumber, error) {
	if arr == nil {
		return utils.ToNumber(0), nil
	}
	s := 0.0
	for _, v := range arr {
		s = s + utils.ToNumber(v).FloatValue()
	}
	return utils.ToNumber(s), nil
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
			return utils.ToNumber(0), err
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
			return nil, err
		}
		if utils.ToBoolean(b) {
			return arr[i], nil
		}
	}
	return nil, &utils.EvaluationError{
		Message: "No matching items found!",
	}
}

func findLast(arr utils.JsonArray) (interface{}, error) {
	if arr == nil {
		return nil, nil
	}
	if len(arr) == 0 {
		return nil, &utils.EvaluationError{
			Message: "No matching items found!",
		}
	}
	return arr[len(arr)-1], nil
}
