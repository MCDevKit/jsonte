package bench

import (
	"sync"
	"testing"

	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/types"
)

// Feature-focused micro-benchmarks isolating individual engine hot paths.
// Run: go test ./bench/ -run='^$' -bench=BenchmarkFeature -benchmem

var (
	featureOnce  sync.Once
	featureScope *types.JsonObject
)

const featureScopeJSON = `{
	"items": [
		{"name": "sword", "damage": 7, "tags": ["melee", "iron"]},
		{"name": "bow", "damage": 5, "tags": ["ranged", "wood"]},
		{"name": "axe", "damage": 9, "tags": ["melee", "iron"]},
		{"name": "wand", "damage": 3, "tags": ["magic", "gold"]},
		{"name": "spear", "damage": 6, "tags": ["melee", "wood"]},
		{"name": "dagger", "damage": 4, "tags": ["melee", "iron"]},
		{"name": "staff", "damage": 8, "tags": ["magic", "wood"]},
		{"name": "mace", "damage": 10, "tags": ["melee", "gold"]}
	],
	"count": 64,
	"prefix": "entity"
}`

func ensureFeatureScope() {
	featureOnce.Do(func() {
		scope, err := types.ParseJsonObject([]byte(featureScopeJSON))
		if err != nil {
			panic(err)
		}
		featureScope = scope
	})
}

func benchmarkTemplate(b *testing.B, template string) {
	b.Helper()
	ensureFeatureScope()
	modules := map[string]jsonte.JsonModule{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := jsonte.Process("feature", template, featureScope, modules, -1)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// String interpolation: many template strings mixing literal text and expressions.
func BenchmarkFeatureStringInterpolation(b *testing.B) {
	benchmarkTemplate(b, `{
		"$template": {
			"{{#(0..count)}}": {
				"{{prefix}}_{{index}}": "state_{{prefix}}_{{index}}_{{index * 2}}"
			}
		}
	}`)
}

// Array iteration: {{#array}} loop generating one object per element.
func BenchmarkFeatureArrayIteration(b *testing.B) {
	benchmarkTemplate(b, `{
		"$template": {
			"{{#(0..count)}}": {
				"entry_{{index}}": {
					"id": "{{index}}",
					"double": "{{index * 2}}",
					"flag": "{{mod(index, 2) == 0}}"
				}
			}
		}
	}`)
}

// Nested iteration: outer loop over items, inner loop over their tags.
func BenchmarkFeatureNestedIteration(b *testing.B) {
	benchmarkTemplate(b, `{
		"$template": {
			"{{#items}}": {
				"{{name}}": {
					"damage": "{{damage}}",
					"{{#tags}}": {
						"tag_{{value}}": true
					}
				}
			}
		}
	}`)
}

// Function pipeline: map/filter/sort chains over arrays.
func BenchmarkFeatureMapFilterPipeline(b *testing.B) {
	benchmarkTemplate(b, `{
		"$template": {
			"melee": "{{items.filter(x => x.tags.contains('melee')).map(x => x.name).sort()}}",
			"strong": "{{items.filter(x => x.damage > 5).map(x => x.name)}}",
			"total": "{{items.map(x => x.damage).sum()}}",
			"flat": "{{items.flatMap(x => x.tags)}}"
		}
	}`)
}

// Expression evaluation: arithmetic and string operations, no iteration.
func BenchmarkFeatureExpressionEval(b *testing.B) {
	benchmarkTemplate(b, `{
		"$template": {
			"a": "{{1 + 2 * 3 - 4 / 2}}",
			"b": "{{(count * 3 + 7) / 2}}",
			"c": "{{prefix + '_suffix'}}",
			"d": "{{min(count, 100) + max(count, 10)}}",
			"e": "{{count > 32 ? 'big' : 'small'}}"
		}
	}`)
}

// Object merge: template merged into an existing JSON structure via $copy-like nesting.
func BenchmarkFeatureDeepObjectGeneration(b *testing.B) {
	benchmarkTemplate(b, `{
		"$template": {
			"level1": {
				"level2": {
					"level3": {
						"{{#items}}": {
							"{{name}}": {
								"nested": {
									"damage": "{{damage}}",
									"name_upper": "{{name.toUpperCase()}}"
								}
							}
						}
					}
				}
			}
		}
	}`)
}
