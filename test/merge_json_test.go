package test

import (
	"testing"

	"github.com/MCDevKit/jsonte/jsonte/types"
)

func TestMergeJSON(t *testing.T) {
	cases := []struct {
		name     string
		template string
		parent   string
		keep     bool
		want     string
		wantErr  bool
	}{
		{"null_template", "null", "{\"a\":1}", false, "{\"a\":1}", false},
		{"null_parent", "{\"a\":1}", "null", false, "{\"a\":1}", false},
		{"object_union", "{\"a\":1}", "{\"b\":2}", false, "{\"a\":1,\"b\":2}", false},
		{"overwrite_value", "{\"a\":1}", "{\"a\":2}", false, "{\"a\":2}", false},
		{"merge_arrays", "{\"arr\":[1]}", "{\"arr\":[2]}", false, "{\"arr\":[1,2]}", false},
		{"merge_arrays_reverse", "{\"arr\":[1]}", "{\"^arr\":[2]}", false, "{\"arr\":[2,1]}", false},
		{"module_override_keep_false", "{\"groups\":{\"object_1\":{\"value\":1}}}", "{\"groups\":{\"$object_1\":{\"value\":0}}}", false, "{\"groups\":{\"object_1\":{\"value\":0}}}", false},
		{"module_override_keep_true", "{\"groups\":{\"object_1\":{\"value\":1}}}", "{\"groups\":{\"$object_1\":{\"value\":0}}}", true, "{\"groups\":{\"object_1\":{\"value\":1},\"$object_1\":{\"value\":0}}}", false},
		{"nested_object", "{\"root\":{\"a\":{\"x\":1}}}", "{\"root\":{\"a\":{\"y\":2},\"b\":3}}", false, "{\"root\":{\"a\":{\"x\":1,\"y\":2},\"b\":3}}", false},
		{"array_of_objects", "{\"arr\":[{\"id\":1}]}", "{\"arr\":[{\"id\":2}]}", false, "{\"arr\":[{\"id\":1},{\"id\":2}]}", false},
		{"mismatched_types", "{\"a\":1}", "[1,2,3]", false, "", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			template, err := parseJsonValue(tc.template)
			if err != nil {
				t.Fatalf("failed to parse template: %v", err)
			}
			parent, err := parseJsonValue(tc.parent)
			if err != nil {
				t.Fatalf("failed to parse parent: %v", err)
			}

			got, err := types.MergeJSON(template, parent, tc.keep)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got none")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			gotStr := types.ToString(got)
			if gotStr != tc.want {
				t.Fatalf("unexpected result\nwant: %s\ngot:  %s", tc.want, gotStr)
			}
		})
	}
}

func parseJsonValue(input string) (types.JsonType, error) {
	if input == "" {
		return nil, nil
	}
	if input == "null" {
		return types.Null, nil
	}
	return types.ParseJsonValue([]byte(input))
}

func BenchmarkMergeJSONObject(b *testing.B) {
	template, _ := types.ParseJsonValue([]byte(`{"root":{"a":{"x":1},"b":2}}`))
	parent, _ := types.ParseJsonValue([]byte(`{"root":{"a":{"y":3},"c":4}}`))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := types.MergeJSON(template, parent, false); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMergeJSONArray(b *testing.B) {
	template, _ := types.ParseJsonValue([]byte(`{"arr":[1,2,3,4,5]}`))
	parent, _ := types.ParseJsonValue([]byte(`{"arr":[6,7,8,9,10]}`))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := types.MergeJSON(template, parent, false); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMergeJSONModuleOverrides(b *testing.B) {
	template, _ := types.ParseJsonValue([]byte(`{"groups":{"object_1":{"value":1},"object_2":{"value":2}}}`))
	parent, _ := types.ParseJsonValue([]byte(`{"groups":{"$object_1":{"value":0},"^object_3":{"value":3}}}`))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := types.MergeJSON(template, parent, false); err != nil {
			b.Fatal(err)
		}
	}
}
