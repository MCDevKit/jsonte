package bench

import (
	"fmt"
	"runtime"
	"sort"
	"strings"
	"sync"
	"testing"

	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
)

// Compile benchmarks simulate a full `jsonte compile` run over a generated,
// realistic add-on: a global scope with many entities, a chain of modules,
// one template per entity, a $files template, mcfunctions, a lang file and
// molang files. Everything is held in memory so the numbers reflect the
// engine rather than disk I/O.
//
// A real compile is a one-shot process, so the "Cold" variants clear the
// expression cache on every iteration; the "Warm" variants keep it.
//
// Run: go test ./bench/ -run='^$' -bench=BenchmarkCompile -benchmem

const (
	compileEntityCount   = 48
	compileVariantCount  = 6
	compileFunctionCount = 24
)

type compilePack struct {
	scope      *types.JsonObject
	modules    []string
	templates  map[string]string
	functions  map[string]string
	molang     map[string]string
	lang       string
	fileNames  []string // sorted template names for deterministic order
	funcNames  []string
	molangName []string
}

var (
	compilePackOnce sync.Once
	compilePackData *compilePack
)

func generateCompileScope() string {
	families := []string{"monster", "undead", "arthropod", "animal", "aquatic", "illager"}
	var sb strings.Builder
	sb.WriteString(`{"namespace": "bench", "difficulty": 2, "debug": false, "entities": [`)
	for i := 0; i < compileEntityCount; i++ {
		if i > 0 {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, `{"id": "mob_%d", "name": "Mob Number %d", "health": %d, "speed": %g, "damage": %d, "hostile": %t, "family": ["%s", "%s"], "loot": [`,
			i, i, 10+i%30, 0.2+float64(i%7)*0.05, 2+i%9, i%3 != 0, families[i%len(families)], families[(i+2)%len(families)])
		for j := 0; j < 4; j++ {
			if j > 0 {
				sb.WriteString(",")
			}
			fmt.Fprintf(&sb, `{"item": "minecraft:item_%d", "min": %d, "max": %d, "chance": %g}`, (i+j)%17, j, j+2, 0.1*float64(j+1))
		}
		sb.WriteString(`], "variants": [`)
		for v := 0; v < compileVariantCount; v++ {
			if v > 0 {
				sb.WriteString(",")
			}
			fmt.Fprintf(&sb, `{"name": "Variant_%d", "texture": "textures/entity/mob_%d/v%d", "scale": %g, "weight": %d}`, v, i, v, 0.8+float64(v)*0.1, 1+v*3)
		}
		sb.WriteString(`]}`)
	}
	sb.WriteString(`]}`)
	return sb.String()
}

const compileBaseModule = `{
	"$module": "base_entity",
	"$scope": {
		"formatVersion": "1.20.80",
		"defaultHealth": 20
	},
	"$template": {
		"format_version": "{{=formatVersion}}",
		"minecraft:entity": {
			"description": {
				"identifier": "{{namespace}}:{{entity.id}}",
				"is_spawnable": true,
				"is_summonable": true,
				"properties": {
					"{{namespace}}:variant": {
						"type": "int",
						"range": [0, "{{=count(entity.variants) - 1}}"],
						"default": 0,
						"client_sync": true
					}
				}
			},
			"components": {
				"minecraft:health": {
					"value": "{{=entity.health ?? defaultHealth}}",
					"max": "{{=entity.health ?? defaultHealth}}"
				},
				"minecraft:type_family": {
					"family": "{{=entity.family}}"
				},
				"minecraft:collision_box": {
					"width": 0.6,
					"height": 1.8
				},
				"minecraft:physics": {},
				"minecraft:pushable": {
					"is_pushable": true,
					"is_pushable_by_piston": true
				}
			},
			"component_groups": {},
			"events": {}
		}
	}
}`

const compileMobModule = `{
	"$module": "mob",
	"$scope": {
		"baseSpeed": 0.25
	},
	"$template": {
		"minecraft:entity": {
			"components": {
				"minecraft:movement": {
					"value": "{{=entity.speed ?? baseSpeed}}"
				},
				"minecraft:navigation.walk": {
					"can_path_over_water": true,
					"avoid_damage_blocks": true
				},
				"minecraft:behavior.random_stroll": {
					"priority": 6,
					"speed_multiplier": "{{=round((entity.speed ?? baseSpeed) * 4, 2)}}"
				},
				"minecraft:behavior.look_at_player": {
					"priority": 7,
					"look_distance": 6
				}
			},
			"component_groups": {
				"{{#entity.variants}}": {
					"{{namespace}}:{{toLowerCase(value.name)}}": {
						"minecraft:scale": {
							"value": "{{=value.scale}}"
						},
						"minecraft:variant": {
							"value": "{{=index}}"
						}
					}
				}
			},
			"events": {
				"minecraft:entity_spawned": {
					"randomize": [
						{
							"{{#entity.variants}}": {
								"weight": "{{=value.weight}}",
								"add": {
									"component_groups": ["{{namespace}}:{{toLowerCase(value.name)}}"]
								},
								"set_property": {
									"{{namespace}}:variant": "{{=index}}"
								}
							}
						}
					]
				}
			}
		}
	}
}`

const compileHostileModule = `{
	"$module": "hostile",
	"$template": {
		"minecraft:entity": {
			"components": {
				"minecraft:attack": {
					"damage": "{{=entity.damage * (difficulty + 1) / 2}}"
				},
				"minecraft:behavior.melee_attack": {
					"priority": 3,
					"track_target": true
				},
				"minecraft:behavior.nearest_attackable_target": {
					"priority": 2,
					"entity_types": [
						{
							"filters": {
								"any_of": [
									{"test": "is_family", "subject": "other", "value": "player"},
									{"{{#entity.family.filter(f => f != 'monster')}}": {"test": "is_family", "subject": "other", "value": "{{=value}}_hunter"}}
								]
							},
							"max_dist": "{{=16 + entity.damage}}"
						}
					]
				}
			}
		}
	}
}`

// Every entity gets its own template file, as is typical in real packs.
func generateCompileEntityTemplate(i int) string {
	module := `["base_entity", "mob"]`
	if i%3 != 0 {
		module = `["base_entity", "mob", "hostile"]`
	}
	return fmt.Sprintf(`{
	"$extend": %s,
	"$scope": {
		"entity": "{{=entities[%d]}}"
	},
	"$template": {
		"minecraft:entity": {
			"description": {
				"runtime_identifier": "{{=entity.hostile ? 'minecraft:zombie' : 'minecraft:pig'}}"
			},
			"components": {
				"minecraft:loot": {
					"table": "loot_tables/entities/{{entity.id}}.json"
				},
				"minecraft:nameable": {},
				"{{?entity.hostile}}": {
					"minecraft:burns_in_daylight": {}
				},
				"{{?!entity.hostile}}": {
					"minecraft:breedable": {
						"require_tame": false,
						"breed_items": "{{=entity.loot.map(x => x.item)}}"
					}
				}
			},
			"events": {
				"{{#1..3 as level}}": {
					"{{namespace}}:upgrade_{{level}}": {
						"sequence": [
							{"{{#entity.variants.filter(v => v.weight > level)}}": {
								"filters": {"test": "int_property", "domain": "{{namespace}}:variant", "value": "{{=index}}"},
								"add": {"component_groups": ["{{namespace}}:{{toLowerCase(value.name)}}"]}
							}}
						]
					}
				}
			}
		}
	}
}`, module, i)
}

// A single $files template generating a loot table per entity.
const compileLootTemplate = `{
	"$files": {
		"array": "{{entities}}",
		"fileName": "loot_{{value.id}}"
	},
	"$template": {
		"pools": [
			{
				"rolls": 1,
				"entries": [
					{
						"{{#value.loot}}": {
							"type": "item",
							"name": "{{=value.item}}",
							"weight": "{{=floor(value.chance * 100)}}",
							"functions": [
								{"function": "set_count", "count": {"min": "{{=value.min}}", "max": "{{=value.max}}"}}
							]
						}
					}
				]
			},
			{
				"{{?value.hostile}}": {
					"rolls": "{{=value.damage}}",
					"conditions": [{"condition": "killed_by_player"}],
					"entries": [{"type": "item", "name": "minecraft:emerald", "weight": 1}]
				}
			}
		]
	}
}`

// A spawn rules / client entity style template that leans on array functions.
const compileSummaryTemplate = `{
	"$template": {
		"totalHealth": "{{=entities.map(e => e.health).sum()}}",
		"hostileIds": "{{=entities.filter(e => e.hostile).map(e => e.id)}}",
		"byFamily": {
			"{{#['monster', 'undead', 'arthropod', 'animal', 'aquatic', 'illager'] as fam}}": {
				"{{fam}}": "{{=entities.filter(e => e.family.contains(fam)).map(e => e.id).sort()}}"
			}
		},
		"textures": {
			"{{#entities as e}}": {
				"{{#e.variants as v}}": {
					"{{e.id}}_{{toLowerCase(v.name)}}": "{{=replace(v.texture, 'textures/', '')}}"
				}
			}
		},
		"strongest": "{{=entities.sort(e => -e.damage).map(e => e.id).sublist(0, 5)}}"
	}
}`

func generateCompileFunction(i int) string {
	var sb strings.Builder
	for j := 0; j < 30; j++ {
		fmt.Fprintf(&sb, "execute as @e[type=#{namespace}:#{entities[%d].id}] at @s run summon #{namespace}:#{entities[%d].id} ~ ~#{%d + 1} ~\n", (i+j)%compileEntityCount, (i*3+j)%compileEntityCount, j)
		fmt.Fprintf(&sb, "scoreboard players set @s health_%d #{entities[%d].health * difficulty}\n", j, (i+j)%compileEntityCount)
		sb.WriteString("say plain line without any expression at all\n")
	}
	return sb.String()
}

func generateCompileLang() string {
	var sb strings.Builder
	for i := 0; i < compileEntityCount; i++ {
		fmt.Fprintf(&sb, "entity.bench:mob_%d.name=##{entities[%d].name}\n", i, i)
		fmt.Fprintf(&sb, "item.spawn_egg.entity.bench:mob_%d.name=Spawn ##{entities[%d].name}\n", i, i)
		for v := 0; v < compileVariantCount; v++ {
			fmt.Fprintf(&sb, "variant.bench.mob_%d.%d=##{entities[%d].variants[%d].name} of ##{toUpperCase(entities[%d].id)}\n", i, v, i, v, i)
		}
		sb.WriteString("## plain comment line\n")
	}
	return sb.String()
}

func generateCompileMolang(i int) string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "variable.speed = #{entities[%d].speed};\n", i%compileEntityCount)
	fmt.Fprintf(&sb, "variable.health = #{entities[%d].health};\n", i%compileEntityCount)
	sb.WriteString("variable.is_moving = query.modified_move_speed > 0.1;\n")
	sb.WriteString("variable.bob = math.sin(query.anim_time * 180.0) * (variable.is_moving ? 1.0 : 0.25);\n")
	fmt.Fprintf(&sb, "variable.damage = #{entities[%d].damage * difficulty} * variable.speed;\n", i%compileEntityCount)
	sb.WriteString("return variable.bob + variable.damage;\n")
	return sb.String()
}

func ensureCompilePack() *compilePack {
	compilePackOnce.Do(func() {
		scope, err := types.ParseJsonObject([]byte(generateCompileScope()))
		if err != nil {
			panic(err)
		}
		p := &compilePack{
			scope:     scope,
			modules:   []string{compileBaseModule, compileMobModule, compileHostileModule},
			templates: map[string]string{},
			functions: map[string]string{},
			molang:    map[string]string{},
			lang:      generateCompileLang(),
		}
		for i := 0; i < compileEntityCount; i++ {
			p.templates[fmt.Sprintf("mob_%d", i)] = generateCompileEntityTemplate(i)
			p.molang[fmt.Sprintf("mob_%d", i)] = generateCompileMolang(i)
		}
		p.templates["loot"] = compileLootTemplate
		p.templates["summary"] = compileSummaryTemplate
		for i := 0; i < compileFunctionCount; i++ {
			p.functions[fmt.Sprintf("func_%d", i)] = generateCompileFunction(i)
		}
		for k := range p.templates {
			p.fileNames = append(p.fileNames, k)
		}
		for k := range p.functions {
			p.funcNames = append(p.funcNames, k)
		}
		for k := range p.molang {
			p.molangName = append(p.molangName, k)
		}
		sort.Strings(p.fileNames)
		sort.Strings(p.funcNames)
		sort.Strings(p.molangName)
		compilePackData = p
	})
	return compilePackData
}

// loadCompileModules loads the modules; templates extend the whole chain
// explicitly (modules cannot extend other modules).
func loadCompileModules(b *testing.B, p *compilePack) map[string]jsonte.JsonModule {
	modules := make(map[string]jsonte.JsonModule, len(p.modules))
	for _, src := range p.modules {
		m, err := jsonte.LoadModule(src, p.scope, -1)
		if err != nil {
			b.Fatal(err)
		}
		modules[m.Name.StringValue()] = m
	}
	return modules
}

func compileTemplate(b *testing.B, p *compilePack, name string, modules map[string]jsonte.JsonModule, minify bool) int {
	out, err := jsonte.Process(name, p.templates[name], p.scope, modules, -1)
	if err != nil {
		b.Fatalf("%s: %v", name, err)
	}
	size := 0
	for _, key := range out.Keys() {
		if minify {
			size += len(types.ToString(out.Get(key)))
		} else {
			size += len(types.ToPrettyString(out.Get(key)))
		}
	}
	return size
}

// compilePackOnce runs the full pipeline sequentially, in the same order as
// the compile command: molang, modules, templates (+serialization),
// mcfunctions, lang.
func compileWholePack(b *testing.B, p *compilePack, minify bool) int {
	size := 0
	for _, name := range p.molangName {
		out, err := jsonte.ProcessMolangFile(p.molang[name], p.scope)
		if err != nil {
			b.Fatalf("molang %s: %v", name, err)
		}
		size += len(out)
	}
	modules := loadCompileModules(b, p)
	for _, name := range p.fileNames {
		size += compileTemplate(b, p, name, modules, minify)
	}
	for _, name := range p.funcNames {
		out, err := jsonte.ProcessMCFunction(p.functions[name], p.scope)
		if err != nil {
			b.Fatalf("mcfunction %s: %v", name, err)
		}
		size += len(out)
	}
	out, err := jsonte.ProcessLangFile(p.lang, p.scope)
	if err != nil {
		b.Fatal(err)
	}
	return size + len(out)
}

var compileSink int

func runCompileBench(b *testing.B, cold bool, fn func() int) {
	b.Helper()
	jsonte.ClearExpressionCache()
	compileSink = fn() // warm-up and sanity check
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if cold {
			b.StopTimer()
			jsonte.ClearExpressionCache()
			b.StartTimer()
		}
		compileSink = fn()
	}
}

// Full pack, sequential, pretty-printed output.
func BenchmarkCompilePackCold(b *testing.B) {
	p := ensureCompilePack()
	runCompileBench(b, true, func() int { return compileWholePack(b, p, false) })
}

func BenchmarkCompilePackWarm(b *testing.B) {
	p := ensureCompilePack()
	runCompileBench(b, false, func() int { return compileWholePack(b, p, false) })
}

func BenchmarkCompilePackMinify(b *testing.B) {
	p := ensureCompilePack()
	runCompileBench(b, true, func() int { return compileWholePack(b, p, true) })
}

// Full pack with templates processed concurrently, like `compile --workers 0`.
func BenchmarkCompilePackParallel(b *testing.B) {
	p := ensureCompilePack()
	workers := runtime.NumCPU()
	runCompileBench(b, true, func() int {
		modules := loadCompileModules(b, p)
		var mu sync.Mutex
		var wg sync.WaitGroup
		sem := make(chan struct{}, workers)
		total := 0
		for _, name := range p.fileNames {
			wg.Add(1)
			go func(name string) {
				defer wg.Done()
				sem <- struct{}{}
				defer func() { <-sem }()
				n := compileTemplate(b, p, name, modules, false)
				mu.Lock()
				total += n
				mu.Unlock()
			}(name)
		}
		for _, name := range p.funcNames {
			wg.Add(1)
			go func(name string) {
				defer wg.Done()
				sem <- struct{}{}
				defer func() { <-sem }()
				out, err := jsonte.ProcessMCFunction(p.functions[name], p.scope)
				if err != nil {
					b.Error(err)
					return
				}
				mu.Lock()
				total += len(out)
				mu.Unlock()
			}(name)
		}
		wg.Wait()
		return total
	})
}

// Stage benchmarks, to attribute time within the pipeline.

func BenchmarkCompileEntityTemplates(b *testing.B) {
	p := ensureCompilePack()
	runCompileBench(b, true, func() int {
		modules := loadCompileModules(b, p)
		n := 0
		for i := 0; i < compileEntityCount; i++ {
			name := fmt.Sprintf("mob_%d", i)
			out, err := jsonte.Process(name, p.templates[name], p.scope, modules, -1)
			if err != nil {
				b.Fatal(err)
			}
			n += out.Size()
		}
		return n
	})
}

func BenchmarkCompileFilesTemplate(b *testing.B) {
	p := ensureCompilePack()
	modules := map[string]jsonte.JsonModule{}
	runCompileBench(b, true, func() int { return compileTemplate(b, p, "loot", modules, false) })
}

func BenchmarkCompileSummaryTemplate(b *testing.B) {
	p := ensureCompilePack()
	modules := map[string]jsonte.JsonModule{}
	runCompileBench(b, true, func() int { return compileTemplate(b, p, "summary", modules, false) })
}

func BenchmarkCompileLoadModules(b *testing.B) {
	p := ensureCompilePack()
	runCompileBench(b, true, func() int { return len(loadCompileModules(b, p)) })
}

func BenchmarkCompileMCFunctions(b *testing.B) {
	p := ensureCompilePack()
	runCompileBench(b, true, func() int {
		n := 0
		for _, name := range p.funcNames {
			out, err := jsonte.ProcessMCFunction(p.functions[name], p.scope)
			if err != nil {
				b.Fatal(err)
			}
			n += len(out)
		}
		return n
	})
}

func BenchmarkCompileLang(b *testing.B) {
	p := ensureCompilePack()
	runCompileBench(b, true, func() int {
		out, err := jsonte.ProcessLangFile(p.lang, p.scope)
		if err != nil {
			b.Fatal(err)
		}
		return len(out)
	})
}

func BenchmarkCompileMolang(b *testing.B) {
	p := ensureCompilePack()
	runCompileBench(b, true, func() int {
		n := 0
		for _, name := range p.molangName {
			out, err := jsonte.ProcessMolangFile(p.molang[name], p.scope)
			if err != nil {
				b.Fatal(err)
			}
			n += len(out)
		}
		return n
	})
}

// Serialization of the already-processed output of the whole pack.
func BenchmarkCompileSerialize(b *testing.B) {
	p := ensureCompilePack()
	modules := loadCompileModules(b, p)
	outputs := make([]utils.NavigableMap[string, types.JsonType], 0, len(p.fileNames))
	for _, name := range p.fileNames {
		out, err := jsonte.Process(name, p.templates[name], p.scope, modules, -1)
		if err != nil {
			b.Fatal(err)
		}
		outputs = append(outputs, out)
	}
	b.Run("pretty", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			n := 0
			for _, out := range outputs {
				for _, key := range out.Keys() {
					n += len(types.ToPrettyString(out.Get(key)))
				}
			}
			compileSink = n
		}
	})
	b.Run("minify", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			n := 0
			for _, out := range outputs {
				for _, key := range out.Keys() {
					n += len(types.ToString(out.Get(key)))
				}
			}
			compileSink = n
		}
	})
}

// Sanity check that the generated pack actually compiles to what we expect,
// so the benchmarks can't silently measure an error path.
func TestCompilePackSanity(t *testing.T) {
	p := ensureCompilePack()
	modules := map[string]jsonte.JsonModule{}
	for _, src := range p.modules {
		m, err := jsonte.LoadModule(src, p.scope, -1)
		if err != nil {
			t.Fatal(err)
		}
		modules[m.Name.StringValue()] = m
	}
	out, err := jsonte.Process("mob_1", p.templates["mob_1"], p.scope, modules, -1)
	if err != nil {
		t.Fatal(err)
	}
	s := types.ToString(out.Get("mob_1"))
	for _, want := range []string{`"bench:mob_1"`, `"minecraft:attack"`, `"bench:variant_5"`, `"minecraft:burns_in_daylight"`, `"randomize"`, `"bench:upgrade_3"`} {
		if !strings.Contains(s, want) {
			t.Errorf("mob_1 output missing %s:\n%s", want, s)
		}
	}
	loot, err := jsonte.Process("loot", p.templates["loot"], p.scope, modules, -1)
	if err != nil {
		t.Fatal(err)
	}
	if loot.Size() != compileEntityCount {
		t.Errorf("expected %d loot files, got %d", compileEntityCount, loot.Size())
	}
	summary, err := jsonte.Process("summary", p.templates["summary"], p.scope, modules, -1)
	if err != nil {
		t.Fatal(err)
	}
	if s := types.ToString(summary.Get("summary")); !strings.Contains(s, `"mob_0_variant_0"`) {
		t.Errorf("summary missing textures:\n%s", s)
	}
	fn, err := jsonte.ProcessMCFunction(p.functions["func_0"], p.scope)
	if err != nil || strings.Contains(fn, "#{") {
		t.Errorf("mcfunction not processed: %v\n%s", err, fn)
	}
	lang, err := jsonte.ProcessLangFile(p.lang, p.scope)
	if err != nil || !strings.Contains(lang, "Mob Number 3") {
		t.Errorf("lang not processed: %v", err)
	}
	ml, err := jsonte.ProcessMolangFile(p.molang["mob_0"], p.scope)
	if err != nil || strings.Contains(ml, "#{") {
		t.Errorf("molang not processed: %v\n%s", err, ml)
	}
}
