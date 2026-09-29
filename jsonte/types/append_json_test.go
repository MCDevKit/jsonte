package types

import (
	"math"
	"testing"

	"github.com/MCDevKit/jsonte/jsonte/json"
)

// TestAppendJSONMatchesMarshalJSONC checks that the direct serializer used by
// ToString/ToPrettyString produces exactly what json.MarshalJSONC produced for
// the unboxed value before the fast path was added.
func TestAppendJSONMatchesMarshalJSONC(t *testing.T) {
	Init()
	strs := []string{"", "plain", "quote\"back\\slash", "tab\tnl\nret\r", "\x00\x01\x1f\x7f",
		"<html>&amp;</html>", "zażółć gęślą jaźń", "emoji 😀", "  ", "inv\xffalid"}
	nums := []*JsonNumber{
		AsNumber(0), AsNumber(1), AsNumber(-1), AsNumber(123456789), AsNumber(1 << 53),
		AsNumber(0.5), AsNumber(-0.25), AsNumber(1e21), AsNumber(1e-7), AsNumber(math.Copysign(0, -1)),
		{Value: 2, Decimal: true}, {Value: 1.5, Decimal: true}, {Value: 1e20, Decimal: true},
	}
	build := func(i int) JsonType {
		obj := NewJsonObject()
		obj.Put("s", AsString(strs[i%len(strs)]))
		obj.Put("n", nums[i%len(nums)])
		obj.Put("b", AsBool(i%2 == 0))
		obj.Put("null", Null)
		obj.Put(strs[(i+3)%len(strs)], nums[(i+5)%len(nums)])
		arr := NewJsonArray()
		for j := 0; j < i%4; j++ {
			arr.Append(AsString(strs[(i+j)%len(strs)]))
			arr.Append(nums[(i+j)%len(nums)])
		}
		obj.Put("arr", arr)
		obj.Put("emptyObj", NewJsonObject())
		obj.Put("emptyArr", NewJsonArray())
		nested := NewJsonObject()
		nested.Put("inner", obj.Get("arr"))
		nested.Put("deep", AsArray([]interface{}{map[string]interface{}{"x": 1}, []interface{}{}}))
		obj.Put("nested", nested)
		if i%3 == 0 {
			return arr
		}
		return obj
	}
	for i := 0; i < len(strs)*len(nums); i++ {
		v := build(i)
		for _, pretty := range []bool{false, true} {
			expected, err := json.MarshalJSONC(v.Unbox(), pretty)
			if err != nil {
				t.Fatal(err)
			}
			var got string
			if pretty {
				got = ToPrettyString(v)
			} else {
				got = ToString(v)
			}
			if got != string(expected) {
				t.Fatalf("case %d pretty=%v mismatch:\nexpected: %s\ngot:      %s", i, pretty, string(expected), got)
			}
		}
	}
}
