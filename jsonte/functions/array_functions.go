package functions

import (
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"math"
	"sort"
)

func RegisterArrayFunctions() {
	const group = "array"
	RegisterGroup(Group{
		Name:    group,
		Title:   "Array functions",
		Summary: "Array functions are for getting information on and transforming arrays.",
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "asArray",
		Body:  asArray,
		Docs: Docs{
			Summary: "Converts an object to an array of objects with the given key and value",
			Arguments: []Argument{
				{
					Name:    "object",
					Summary: "The object to convert",
				},
				{
					Name:    "key",
					Summary: "The key to use for the new objects",
				},
				{
					Name:    "value",
					Summary: "The value to use for the new objects",
				},
			},
			Example: `
Given scope
<code>
{
  "testObject": {
    "test1": "someVal",
    "test2": "anotherVal"
  }
}
</code>
for query
<code>
{
  "$template": {
    "test": "{{asArray(testObject, 'key', 'value')}}"
  }
}
</code>
the result will be
<code>
[
  {
    "key": "test1",
    "value": "someVal"
  },
  {
    "key": "test2",
    "value": "anotherVal"
  }
]
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "keys",
		Body:  keys,
		Docs: Docs{
			Summary: "Returns an array of the keys of the given object",
			Arguments: []Argument{
				{
					Name:    "object",
					Summary: "The object to get the keys from",
				},
			},
			Example: `
Given scope
<code>
{
  "testObject": {
    "test1": "someVal",
    "test2": "anotherVal"
  }
}
</code>
for query
<code>
{
  "$template": {
    "test": "{{keys(testObject)}}"
  }
}
</code>
the result will be
<code>
[
  "test1", "test2"
]
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "values",
		Body:  values,
		Docs: Docs{
			Summary: "Returns an array of the values of the given object",
			Arguments: []Argument{
				{
					Name:    "object",
					Summary: "The object to get the values from",
				},
			},
			Example: `
Given scope
<code>
{
  "testObject": {
    "test1": "someVal",
    "test2": "anotherVal"
  }
}
</code>
for query
<code>
{
  "$template": {
    "test": "{{values(testObject)}}"
  }
}
</code>
the result will be
<code>
[
  "someVal",
  "anotherVal"
]
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "sort",
		Body:       sort_,
		IsInstance: true,
		Docs: Docs{
			Summary: "Sorts the given array",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The array to sort",
				},
				{
					Name:     "selector(element, index)",
					Summary:  "The selector to apply to each element before summing",
					Optional: true,
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be [1, 2, 3, 5, 8, 10]",
    "test": "{{[2, 3, 1, 5, 8, 10].sort()}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "sort",
		Body:       sortMap,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "reverse",
		Body:       reverse,
		IsInstance: true,
		Docs: Docs{
			Summary: "Reverses the given array",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The array to reverse",
				},
			},
			Example: `
<code>
{
  "$template": {
	"$comment": "The field below will be [10, 8, 5, 3, 2, 1]",
	"test": "{{[1, 2, 3, 5, 8, 10].reverse()}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "contains",
		Body:       arrayContains,
		IsInstance: true,
		Docs: Docs{
			Summary: "Checks if the given array contains the given value",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The array to check",
				},
				{
					Name:    "value",
					Summary: "The value to check for",
				},
			},
			Example: `
<code>
{
  "$template": {
	"$comment": "The field below will be true",
	"test": "{{[1, 2, 3, 5, 8, 10].contains(5)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "indexOf",
		Body:       arrayIndexOf,
		IsInstance: true,
		Docs: Docs{
			Summary: "Returns the first index of the given value in the given array",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The array to check",
				},
				{
					Name:    "value",
					Summary: "The value to check for",
				},
			},
			Example: `
<code>
{
  "$template": {
	"$comment": "The field below will be 3",
	"test": "{{[1, 2, 3, 5, 8, 10].indexOf(5)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "lastIndexOf",
		Body:       arrayLastIndexOf,
		IsInstance: true,
		Docs: Docs{
			Summary: "Returns the last index of the given value in the given array",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The array to check",
				},
				{
					Name:    "value",
					Summary: "The value to check for",
				},
			},
			Example: `
<code>
{
  "$template": {
	"$comment": "The field below will be 3",
	"test": "{{[1, 2, 3, 5, 8, 10].indexOf(5)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "any",
		Body:       any_,
		IsInstance: true,
		Docs: Docs{
			Summary: "Checks if any of the elements in the given array match the given predicate",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The array to check",
				},
				{
					Name:    "predicate(element, index)",
					Summary: "The predicate to check against",
				},
			},
			Example: `
<code>
{
  "$template": {
	"$comment": "The field below will be true",
	"test": "{{[1, 2, 3, 5, 8, 10].any(x => x > 5)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "all",
		Body:       all,
		IsInstance: true,
		Docs: Docs{
			Summary: "Checks if all of the elements in the given array match the given predicate",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The array to check",
				},
				{
					Name:    "predicate(element, index)",
					Summary: "The predicate to check against",
				},
			},
			Example: `
<code>
{
  "$template": {
	"$comment": "The field below will be false",
	"test": "{{[1, 2, 3, 5, 8, 10].all(x => x > 5)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "count",
		Body:       count,
		IsInstance: true,
		Docs: Docs{
			Summary: "Counts the number of elements in the given array that match the given predicate",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The array to check",
				},
				{
					Name:     "predicate(element, index)",
					Summary:  "The predicate to filter by",
					Optional: true,
				},
			},
			Example: `
<code>
{
  "$template": {
	"$comment": "The field below will be 3",
	"test": "{{[1, 2, 3, 5, 8, 10].count(x => x >= 5)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "count",
		Body:       countFilter,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "sum",
		Body:       sum,
		IsInstance: true,
		Docs: Docs{
			Summary: "Sums the elements in the given array",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The array to sum",
				},
				{
					Name:     "selector(element, index)",
					Summary:  "The selector to apply to each element before summing",
					Optional: true,
				},
			},
			Example: `
<code>
{
  "$template": {
	"$comment": "The field below will be 31",
	"test": "{{[1, 2, 3, 5, 8, 10].sum(x => x)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "sum",
		Body:       sumMap,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "none",
		Body:       none,
		IsInstance: true,
		Docs: Docs{
			Summary: "Checks if none of the elements in the given array match the given predicate",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The array to check",
				},
				{
					Name:    "predicate(element, index)",
					Summary: "The predicate to check against",
				},
			},
			Example: `
<code>
{
  "$template": {
	"$comment": "The field below will be false",
	"test": "{{[1, 2, 3, 5, 8, 10].none(x => x > 5)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "filter",
		Body:       filter,
		IsInstance: true,
		Docs: Docs{
			Summary: "Filters the elements in the given array based on the given predicate",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The array to filter",
				},
				{
					Name:    "predicate(element, index)",
					Summary: "The predicate to filter by",
				},
			},
			Example: `
<code>
{
  "$template": {
	"$comment": "The field below will be [5, 8, 10]",
	"test": "{{[1, 2, 3, 5, 8, 10].filter(x => x >= 5)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "map",
		Body:       map_,
		IsInstance: true,
		Docs: Docs{
			Summary: "Maps the elements in the given array based on the given selector",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The array to map",
				},
				{
					Name:    "selector(element, index)",
					Summary: "The selector to apply to each element",
				},
			},
			Example: `
<code>
{
  "$template": {
	"$comment": "The field below will be [2, 4, 6, 10, 16, 20]",
	"test": "{{[1, 2, 3, 5, 8, 10].map(x => x * 2)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "flatMap",
		Body:       flatMap,
		IsInstance: true,
		Docs: Docs{
			Summary: "Maps the elements in the given array based on the given selector, then flattens the result",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The array to map",
				},
				{
					Name:     "selector(element, index)",
					Summary:  "The selector to apply to each element",
					Optional: true,
				},
			},
			Example: `
<code>
{
  "$template": {
	"$comment": "The field below will be [1, 2, 2, 4, 3, 6, 5, 10, 8, 16, 10, 20]",
	"test": "{{[1, 2, 3, 5, 8, 10].flatMap(x => [x, x * 2])}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "flatMap",
		Body:       flatMapSimple,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "range",
		Body:       range_,
		IsInstance: true,
		Docs: Docs{
			Summary: "Creates an array of indices from the given array",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The array to create indices for",
				},
			},
			Example: `
<code>
{
  "$template": {
	"$comment": "The field below will be [0, 1, 2, 3, 4, 5]",
	"test": "{{[1, 2, 3, 5, 8, 10].range()}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "findFirst",
		Body:       findFirst,
		IsInstance: true,
		Docs: Docs{
			Summary: "Finds the first element in the given array that matches the given predicate",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The array to find in",
				},
				{
					Name:     "predicate(element, index)",
					Summary:  "The predicate to match by",
					Optional: true,
				},
			},
			Example: `
<code>
{
  "$template": {
	"$comment": "The field below will be 8",
	"test": "{{[1, 2, 3, 5, 8, 10].findFirst(x => x >= 5)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "findFirst",
		Body:       findFirstFilter,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "encode",
		Body:       encode,
		IsInstance: true,
		Docs: Docs{
			Summary: "Encodes the given array of numbers to a single number in a bit width calculated with the given maximum",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The array to encode",
				},
				{
					Name:    "maximum",
					Summary: "The maximum value in the array. Must be power of 2.",
				},
				{
					Name:     "selector(element, index)",
					Summary:  "The selector to apply to each element before encoding",
					Optional: true,
				},
			},
			Example: `
<code>
{
  "$template": {
	"$comment": "The field below will be -2023406815 (1000 0111 0110 0101 0100 0011 0010 0001)",
	"test": "{{(1..10).encode(16)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "encode",
		Body:       encodeSimple,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "sublist",
		Body:       sublist,
		IsInstance: true,
		Docs: Docs{
			Summary: "Returns a sublist of the given array",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The array to slice",
				},
				{
					Name:    "start",
					Summary: "The start index",
				},
				{
					Name:     "end",
					Summary:  "The end index",
					Optional: true,
				},
			},
			Example: `
<code>
{
  "$template": {
	"$comment": "The field below will be [1, 2, 3]",
	"test": "{{[1, 2, 3, 5, 8, 10].sublist(0, 3)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "sublist",
		Body:       sublistStart,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "max",
		Body:       maxArray,
		IsInstance: true,
		Docs: Docs{
			Summary: "Finds the maximum value in the given array",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The array to find the maximum value in",
				},
				{
					Name:     "selector(element, index)",
					Summary:  "The selector to apply to each element before finding the maximum value",
					Optional: true,
				},
			},
			Example: `
<code>
{
  "$template": {
	"$comment": "The field below will be 10",
	"test": "{{[1, 2, 3, 5, 8, 10].max()}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "max",
		Body:       maxArraySimple,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "min",
		Body:       minArray,
		IsInstance: true,
		Docs: Docs{
			Summary: "Finds the minimum value in the given array",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The array to find the minimum value in",
				},
				{
					Name:     "selector(element, index)",
					Summary:  "The selector to apply to each element before finding the minimum value",
					Optional: true,
				},
			},
			Example: `
<code>
{
  "$template": {
	"$comment": "The field below will be 1",
	"test": "{{[1, 2, 3, 5, 8, 10].min()}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "min",
		Body:       minArraySimple,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "reduce",
		Body:       reduce,
		IsInstance: true,
		Docs: Docs{
			Summary: "Reduces the given array to a single value using the given accumulator function",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The array to reduce",
				},
				{
					Name:    "accumulator(accumulator, element, index)",
					Summary: "The accumulator function",
				},
				{
					Name:     "initialValue",
					Summary:  "The initial value for the accumulator",
					Optional: true,
				},
			},
			Example: `
<code>
{
  "$template": {
	"$comment": "The field below will be 29",
	"test": "{{[1, 2, 3, 5, 8, 10].reduce((a, b) => a + b)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "reduce",
		Body:       reduceInit,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "findLast",
		Body:       findLast,
		IsInstance: true,
		Docs: Docs{
			Summary: "Finds the last element in the array that matches the given predicate",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The array to search",
				},
				{
					Name:     "predicate(element, index)",
					Summary:  "The predicate to test each element against",
					Optional: true,
				},
			},
			Example: `
<code>
{
  "$template": {
	"$comment": "The field below will be 10",
	"test": "{{[1, 2, 3, 5, 8, 10].findLast(e => e % 2 == 0)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "findLast",
		Body:       findLastFilter,
		IsInstance: true,
	})
}

func asArray(obj utils.NavigableMap[string, interface{}], key, value string) []interface{} {
	if obj.IsEmpty() {
		return nil
	}
	arr := make([]interface{}, obj.Size())
	i := 0
	for _, k := range obj.Keys() {
		u := utils.NewNavigableMap[string, interface{}]()
		u.Put(key, k)
		u.Put(value, obj.Get(k))
		arr[i] = u
		i++
	}
	return arr
}

func keys(obj utils.NavigableMap[string, interface{}]) []interface{} {
	if obj.IsEmpty() {
		return nil
	}
	arr := make([]interface{}, obj.Size())
	i := 0
	for _, k := range obj.Keys() {
		arr[i] = k
		i++
	}
	return arr
}

func values(obj utils.NavigableMap[string, interface{}]) []interface{} {
	if obj.IsEmpty() {
		return nil
	}
	arr := make([]interface{}, obj.Size())
	i := 0
	for _, v := range obj.Values() {
		arr[i] = v
		i++
	}
	return arr
}

func reverse(arr []interface{}) []interface{} {
	if arr == nil {
		return nil
	}
	rev := make([]interface{}, len(arr))
	i := 0
	for j := len(arr) - 1; j >= 0; j-- {
		rev[i] = arr[j]
		i++
	}
	return rev
}

func sort_(arr []interface{}) []interface{} {
	if arr == nil {
		return nil
	}
	result := make([]interface{}, len(arr))
	copy(result, arr)
	sort.SliceStable(result, func(i, j int) bool {
		return utils.Less(result[i], result[j])
	})
	return result
}

func sortMap(arr []interface{}, predicate utils.JsonLambda) ([]interface{}, error) {
	if arr == nil {
		return nil, nil
	}
	mapped, err := map_(arr, predicate)
	if err != nil {
		return nil, utils.WrapErrorf(err, "An error occurred while mapping the values for sorting the array")
	}
	result := make([]interface{}, len(arr))
	copy(result, arr)
	sort.SliceStable(result, func(i, j int) bool {
		return utils.Less(mapped[i], mapped[j])
	})
	return result, nil
}

func arrayContains(arr []interface{}, value interface{}) bool {
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

func arrayIndexOf(arr []interface{}, value interface{}) utils.JsonNumber {
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

func arrayLastIndexOf(arr []interface{}, value interface{}) utils.JsonNumber {
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

func any_(arr []interface{}, predicate utils.JsonLambda) (bool, error) {
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

func all(arr []interface{}, predicate utils.JsonLambda) (bool, error) {
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

func none(arr []interface{}, predicate utils.JsonLambda) (bool, error) {
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

func filter(arr []interface{}, predicate utils.JsonLambda) ([]interface{}, error) {
	if arr == nil {
		return nil, nil
	}
	result := make([]interface{}, 0)
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

func map_(arr []interface{}, predicate utils.JsonLambda) ([]interface{}, error) {
	if arr == nil {
		return nil, nil
	}
	result := make([]interface{}, len(arr))
	for i, v := range arr {
		b, err := predicate([]interface{}{v, i})
		if err != nil {
			return nil, utils.WrapErrorf(err, "An error occurred while evaluating the predicate for `map` at index %d", i)
		}
		result[i] = b
	}
	return result, nil
}

func flatMap(arr []interface{}, predicate utils.JsonLambda) ([]interface{}, error) {
	if arr == nil {
		return nil, nil
	}
	result := make([]interface{}, 0)
	for i, v := range arr {
		b, err := predicate([]interface{}{v, i})
		if err != nil {
			return nil, utils.WrapErrorf(err, "An error occurred while evaluating the predicate for `flatMap` at index %d", i)
		}
		if arr, ok := b.([]interface{}); ok {
			result = append(result, arr...)
		} else {
			result = append(result, b)
		}
	}
	return result, nil
}

func flatMapSimple(arr []interface{}) []interface{} {
	if arr == nil {
		return nil
	}
	result := make([]interface{}, 0)
	for _, v := range arr {
		if arr, ok := v.([]interface{}); ok {
			result = append(result, arr...)
		} else {
			result = append(result, v)
		}
	}
	return result
}

func countFilter(arr []interface{}, predicate utils.JsonLambda) (utils.JsonNumber, error) {
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

func count(arr []interface{}) utils.JsonNumber {
	if arr == nil {
		return utils.ToNumber(0)
	}
	return utils.ToNumber(len(arr))
}

func range_(arr []interface{}) []interface{} {
	if arr == nil {
		return nil
	}
	return utils.CreateRange(0, int32(len(arr))-1)
}

func findFirstFilter(arr []interface{}, predicate utils.JsonLambda) (interface{}, error) {
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

func findFirst(arr []interface{}) (interface{}, error) {
	if arr == nil {
		return nil, nil
	}
	if len(arr) == 0 {
		return nil, utils.WrappedError("No matching items found!")
	}
	return arr[0], nil
}

func encode(arr []interface{}, space utils.JsonNumber, predicate utils.JsonLambda) (utils.JsonNumber, error) {
	if space.IntValue() <= 0 || (space.IntValue()&(space.IntValue()-1)) != 0 {
		return utils.ToNumber(0), utils.WrappedError("Space must be a power of 2 and greater than 0!")
	}
	if arr == nil {
		return utils.ToNumber(0), nil
	}
	var result int32 = 0
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

func encodeSimple(arr []interface{}, space utils.JsonNumber) (utils.JsonNumber, error) {
	return encode(arr, space, func(args []interface{}) (interface{}, error) {
		return args[0], nil
	})
}

func sublist(arr []interface{}, start utils.JsonNumber, end utils.JsonNumber) []interface{} {
	if arr == nil {
		return nil
	}
	startIndex := start.IntValue()
	endIndex := end.IntValue()
	if startIndex < 0 {
		startIndex = 0
	}
	if endIndex > int32(len(arr)) {
		endIndex = int32(len(arr))
	}
	return arr[startIndex:endIndex]
}

func sublistStart(arr []interface{}, start utils.JsonNumber) []interface{} {
	if arr == nil {
		return nil
	}
	startIndex := start.IntValue()
	if startIndex < 0 {
		startIndex = 0
	}
	return arr[startIndex:]
}

func maxArray(arr []interface{}, predicate utils.JsonLambda) (interface{}, error) {
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

func maxArraySimple(arr []interface{}) interface{} {
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

func minArray(arr []interface{}, predicate utils.JsonLambda) (interface{}, error) {
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

func minArraySimple(arr []interface{}) interface{} {
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

func sumMap(arr []interface{}, predicate utils.JsonLambda) (utils.JsonNumber, error) {
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

func sum(arr []interface{}) utils.JsonNumber {
	if arr == nil {
		return utils.ToNumber(0)
	}
	s := 0.0
	for _, v := range arr {
		s = s + utils.ToNumber(v).FloatValue()
	}
	return utils.ToNumber(s)
}

func reduce(arr []interface{}, predicate utils.JsonLambda) (interface{}, error) {
	return reduceInit(arr, predicate, nil)
}

func reduceInit(arr []interface{}, predicate utils.JsonLambda, initialValue interface{}) (interface{}, error) {
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

func findLastFilter(arr []interface{}, predicate utils.JsonLambda) (interface{}, error) {
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

func findLast(arr []interface{}) (interface{}, error) {
	if arr == nil {
		return nil, nil
	}
	if len(arr) == 0 {
		return nil, utils.WrappedError("No matching items found!")
	}
	return arr[len(arr)-1], nil
}
