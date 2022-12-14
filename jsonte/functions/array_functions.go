package functions

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"math"
	"math/rand"
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
		Body:  objectAsArray,
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
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "random",
		Body:       randomElement,
		IsInstance: true,
		Docs: Docs{
			Summary: "Finds the random element in the array",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The array to get random element from",
				},
			},
			Example: `
<code>
{
  "$template": {
	"$comment": "The field below will be a randomly selected element from the array",
	"test": "{{[1, 2, 3, 5, 8, 10].random()}}"
  }
}
</code>`,
		},
	})
}

func randomElement(arr types.JsonArray) (types.JsonType, error) {
	if len(arr.Value) == 0 {
		return nil, burrito.WrappedError("Cannot get random element from an empty array")
	}
	return arr.Value[rand.Intn(len(arr.Value))], nil
}

func objectAsArray(obj types.JsonObject, key, value types.JsonString) types.JsonArray {
	if obj.IsEmpty() {
		return types.NewJsonArray()
	}
	arr := make([]types.JsonType, obj.Size())
	i := 0
	for _, k := range obj.Keys() {
		u := types.NewJsonObject()
		u.Put(key.StringValue(), types.AsString(k))
		u.Put(value.StringValue(), obj.Get(k))
		arr[i] = u
		i++
	}
	return types.JsonArray{Value: arr}
}

func keys(obj types.JsonObject) types.JsonArray {
	if obj.IsEmpty() {
		return types.NewJsonArray()
	}
	arr := make([]types.JsonType, obj.Size())
	i := 0
	for _, k := range obj.Keys() {
		arr[i] = types.AsString(k)
		i++
	}
	return types.JsonArray{Value: arr}
}

func values(obj types.JsonObject) types.JsonArray {
	if obj.IsEmpty() {
		return types.NewJsonArray()
	}
	arr := make([]types.JsonType, obj.Size())
	i := 0
	for _, v := range obj.Values() {
		arr[i] = v
		i++
	}
	return types.JsonArray{Value: arr}
}

func reverse(arr types.JsonArray) types.JsonArray {
	rev := make([]types.JsonType, len(arr.Value))
	i := 0
	for j := len(arr.Value) - 1; j >= 0; j-- {
		rev[i] = arr.Value[j]
		i++
	}
	return types.JsonArray{Value: rev}
}

func sort_(arr types.JsonArray) types.JsonArray {
	result := make([]types.JsonType, len(arr.Value))
	copy(result, arr.Value)
	sort.SliceStable(result, func(i, j int) bool {
		than, _ := result[i].LessThan(result[j])
		return than
	})
	return types.JsonArray{Value: result}
}

func sortMap(arr types.JsonArray, predicate types.JsonLambda) (types.JsonArray, error) {
	mapped, err := map_(arr, predicate)
	if err != nil {
		return types.NewJsonArray(), burrito.WrapErrorf(err, "An error occurred while mapping the values for sorting the array")
	}
	result := make([]types.JsonType, len(arr.Value))
	copy(result, arr.Value)
	sort.SliceStable(result, func(i, j int) bool {
		than, _ := mapped.Value[i].LessThan(mapped.Value[j])
		return than
	})
	return types.JsonArray{Value: result}, nil
}

func arrayContains(arr types.JsonArray, value types.JsonType) types.JsonBool {
	for _, v := range arr.Value {
		if v.Equals(value) {
			return types.True
		}
	}
	return types.False
}

func arrayIndexOf(arr types.JsonArray, value types.JsonType) types.JsonNumber {
	for i, v := range arr.Value {
		if v.Equals(value) {
			return types.AsNumber(i)
		}
	}
	return types.AsNumber(-1)
}

func arrayLastIndexOf(arr types.JsonArray, value types.JsonType) types.JsonNumber {
	for i := len(arr.Value) - 1; i >= 0; i-- {
		if arr.Value[i].Equals(value) {
			return types.AsNumber(i)
		}
	}
	return types.AsNumber(-1)
}

func any_(arr types.JsonArray, predicate types.JsonLambda) (types.JsonBool, error) {
	for i, v := range arr.Value {
		b, err := predicate.Value(paramsForLambda([]interface{}{v, i}))
		if err != nil {
			return types.False, burrito.WrapErrorf(err, "An error occurred while evaluating the predicate for `any` at index %d", i)
		}
		if types.AsBool(b).BoolValue() {
			return types.True, nil
		}
	}
	return types.False, nil
}

func all(arr types.JsonArray, predicate types.JsonLambda) (types.JsonBool, error) {
	for i, v := range arr.Value {
		b, err := predicate.Value(paramsForLambda([]interface{}{v, i}))
		if err != nil {
			return types.False, burrito.WrapErrorf(err, "An error occurred while evaluating the predicate for `all` at index %d", i)
		}
		if !types.AsBool(b).BoolValue() {
			return types.False, nil
		}
	}
	return types.True, nil
}

func none(arr types.JsonArray, predicate types.JsonLambda) (types.JsonBool, error) {
	for i, v := range arr.Value {
		b, err := predicate.Value(paramsForLambda([]interface{}{v, i}))
		if err != nil {
			return types.False, burrito.WrapErrorf(err, "An error occurred while evaluating the predicate for `none` at index %d", i)
		}
		if types.AsBool(b).BoolValue() {
			return types.False, nil
		}
	}
	return types.True, nil
}

func filter(arr types.JsonArray, predicate types.JsonLambda) (types.JsonArray, error) {
	result := make([]types.JsonType, 0)
	for i, v := range arr.Value {
		b, err := predicate.Value(paramsForLambda([]interface{}{v, i}))
		if err != nil {
			return types.NewJsonArray(), burrito.WrapErrorf(err, "An error occurred while evaluating the predicate for `filter` at index %d", i)
		}
		if types.AsBool(b).BoolValue() {
			result = append(result, v)
		}
	}
	return types.JsonArray{Value: result}, nil
}

func map_(arr types.JsonArray, predicate types.JsonLambda) (types.JsonArray, error) {
	result := make([]types.JsonType, len(arr.Value))
	for i, v := range arr.Value {
		b, err := predicate.Value(paramsForLambda([]interface{}{v, i}))
		if err != nil {
			return types.NewJsonArray(), burrito.WrapErrorf(err, "An error occurred while evaluating the predicate for `map` at index %d", i)
		}
		result[i] = b
	}
	return types.JsonArray{Value: result}, nil
}

func flatMap(arr types.JsonArray, predicate types.JsonLambda) (types.JsonArray, error) {
	result := make([]types.JsonType, 0)
	for i, v := range arr.Value {
		b, err := predicate.Value(paramsForLambda([]interface{}{v, i}))
		if err != nil {
			return types.NewJsonArray(), burrito.WrapErrorf(err, "An error occurred while evaluating the predicate for `flatMap` at index %d", i)
		}
		if arr, ok := b.(types.JsonArray); ok {
			result = append(result, arr.Value...)
		} else {
			result = append(result, b)
		}
	}
	return types.JsonArray{Value: result}, nil
}

func flatMapSimple(arr types.JsonArray) types.JsonArray {
	result := make([]types.JsonType, 0)
	for _, v := range arr.Value {
		if arr, ok := v.(types.JsonArray); ok {
			result = append(result, arr.Value...)
		} else {
			result = append(result, v)
		}
	}
	return types.JsonArray{Value: result}
}

func countFilter(arr types.JsonArray, predicate types.JsonLambda) (types.JsonNumber, error) {
	count := 0
	for i, v := range arr.Value {
		b, err := predicate.Value(paramsForLambda([]interface{}{v, i}))
		if err != nil {
			return types.AsNumber(0), burrito.WrapErrorf(err, "An error occurred while evaluating the predicate for `count` at index %d", i)
		}
		if types.AsBool(b).BoolValue() {
			count++
		}
	}
	return types.AsNumber(count), nil
}

func count(arr types.JsonArray) types.JsonNumber {
	return types.AsNumber(len(arr.Value))
}

func range_(arr types.JsonArray) types.JsonArray {
	return types.CreateRange(0, int32(len(arr.Value))-1)
}

func findFirstFilter(arr types.JsonArray, predicate types.JsonLambda) (types.JsonType, error) {
	for i, v := range arr.Value {
		b, err := predicate.Value(paramsForLambda([]interface{}{v, i}))
		if err != nil {
			return nil, burrito.WrapErrorf(err, "An error occurred while evaluating the predicate for `findFirst` at index %d", i)
		}
		if b.BoolValue() {
			return v, nil
		}
	}

	return nil, burrito.WrappedError("No matching items found!")
}

func findFirst(arr types.JsonArray) (types.JsonType, error) {
	if len(arr.Value) == 0 {
		return nil, burrito.WrappedError("No matching items found!")
	}
	return arr.Value[0], nil
}

func encode(arr types.JsonArray, space types.JsonNumber, predicate types.JsonLambda) (types.JsonNumber, error) {
	if space.IntValue() <= 0 || (space.IntValue()&(space.IntValue()-1)) != 0 {
		return types.AsNumber(0), burrito.WrappedError("Space must be a power of 2 and greater than 0!")
	}
	var result int32 = 0
	bitsPerElement := int(math.Log(float64(space.IntValue())) / math.Log(2))
	for i := 0; i < int(math.Min(float64(len(arr.Value)), float64(32/bitsPerElement))); i++ {
		i2, err := predicate.Value(paramsForLambda([]interface{}{arr.Value[i], i}))
		if err != nil {
			return types.AsNumber(0), burrito.WrapErrorf(err, "An error occurred while evaluating the predicate for `encode` at index %d", i)
		}
		if !types.IsNumber(i2) {
			return types.AsNumber(0), burrito.WrappedError("Predicate must return a number!")
		}
		number := types.AsNumber(i2)
		if number.IntValue() < 0 || number.IntValue() >= space.IntValue() {
			return types.AsNumber(0), burrito.WrappedErrorf("Number %d is out of range 0..%d", number.IntValue(), space.IntValue()-1)
		}
		result += number.IntValue() << (i * bitsPerElement)
	}
	return types.AsNumber(result), nil
}

func encodeSimple(arr types.JsonArray, space types.JsonNumber) (types.JsonNumber, error) {
	return encode(arr, space, types.IdentityLambda)
}

func sublist(arr types.JsonArray, start types.JsonNumber, end types.JsonNumber) types.JsonArray {
	startIndex := start.IntValue()
	endIndex := end.IntValue()
	if startIndex < 0 {
		startIndex = 0
	}
	if endIndex > int32(len(arr.Value)) {
		endIndex = int32(len(arr.Value))
	}
	return types.JsonArray{Value: arr.Value[startIndex:endIndex]}
}

func sublistStart(arr types.JsonArray, start types.JsonNumber) types.JsonArray {
	startIndex := start.IntValue()
	if startIndex < 0 {
		startIndex = 0
	}
	return types.JsonArray{Value: arr.Value[startIndex:]}
}

func maxArray(arr types.JsonArray, predicate types.JsonLambda) (types.JsonType, error) {
	max, err := predicate.Value(paramsForLambda([]interface{}{arr.Value[0], 0}))
	if err != nil {
		return nil, burrito.WrapErrorf(err, "An error occurred while evaluating the predicate for `max` at index %d", 0)
	}
	for i, v := range arr.Value {
		b, err := predicate.Value(paramsForLambda([]interface{}{v, i}))
		if err != nil {
			return nil, burrito.WrapErrorf(err, "An error occurred while evaluating the predicate for `max` at index %d", i)
		}
		if types.AsNumber(b).FloatValue() > types.AsNumber(max).FloatValue() {
			max = b
		}
	}
	return max, nil
}

func maxArraySimple(arr types.JsonArray) types.JsonType {
	max := arr.Value[0]
	for _, v := range arr.Value {
		if types.AsNumber(v).FloatValue() > types.AsNumber(max).FloatValue() {
			max = v
		}
	}
	return max
}

func minArray(arr types.JsonArray, predicate types.JsonLambda) (types.JsonType, error) {
	if len(arr.Value) == 0 {
		return types.Null, nil
	}
	min, err := predicate.Value(paramsForLambda([]interface{}{arr.Value[0], 0}))
	if err != nil {
		return nil, burrito.WrapErrorf(err, "An error occurred while evaluating the predicate for `min` at index %d", 0)
	}
	for i, v := range arr.Value {
		if i == 0 {
			continue
		}
		b, err := predicate.Value(paramsForLambda([]interface{}{v, i}))
		if err != nil {
			return nil, burrito.WrapErrorf(err, "An error occurred while evaluating the predicate for `min` at index %d", i)
		}
		if types.AsNumber(b).FloatValue() < types.AsNumber(min).FloatValue() {
			min = b
		}
	}
	return min, nil
}

func minArraySimple(arr types.JsonArray) types.JsonType {
	min := arr.Value[0]
	for _, v := range arr.Value {
		if types.AsNumber(v).FloatValue() < types.AsNumber(min).FloatValue() {
			min = v
		}
	}
	return min
}

func sumMap(arr types.JsonArray, predicate types.JsonLambda) (types.JsonNumber, error) {
	s := 0.0
	for i, v := range arr.Value {
		b, err := predicate.Value(paramsForLambda([]interface{}{v, i}))
		if err != nil {
			return types.AsNumber(0), burrito.WrapErrorf(err, "An error occurred while evaluating the predicate for `sum` at index %d", i)
		}
		s = s + types.AsNumber(b).FloatValue()
	}
	return types.AsNumber(s), nil
}

func sum(arr types.JsonArray) types.JsonNumber {
	s := 0.0
	for _, v := range arr.Value {
		s = s + types.AsNumber(v).FloatValue()
	}
	return types.AsNumber(s)
}

func reduce(arr types.JsonArray, predicate types.JsonLambda) (types.JsonType, error) {
	return reduceInit(arr, predicate, types.Null)
}

func reduceInit(arr types.JsonArray, predicate types.JsonLambda, initialValue types.JsonType) (types.JsonType, error) {
	var prev = initialValue
	for i, v := range arr.Value {
		b, err := predicate.Value(paramsForLambda([]interface{}{prev, v, i}))
		if err != nil {
			return types.AsNumber(0), burrito.WrapErrorf(err, "An error occurred while evaluating the predicate for `reduce` at index %d", i)
		}
		prev = b
	}
	return prev, nil
}

func findLastFilter(arr types.JsonArray, predicate types.JsonLambda) (types.JsonType, error) {
	for i := len(arr.Value) - 1; i >= 0; i-- {
		b, err := predicate.Value(paramsForLambda([]interface{}{arr.Value[i], i}))
		if err != nil {
			return nil, burrito.WrapErrorf(err, "An error occurred while evaluating the predicate for `findLast` at index %d", i)
		}
		if b.BoolValue() {
			return arr.Value[i], nil
		}
	}
	return nil, burrito.WrappedError("No matching items found!")
}

func findLast(arr types.JsonArray) (types.JsonType, error) {
	if len(arr.Value) == 0 {
		return nil, burrito.WrappedError("No matching items found!")
	}
	return arr.Value[len(arr.Value)-1], nil
}
