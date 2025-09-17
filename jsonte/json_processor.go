package jsonte

import (
	"fmt"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/gammazero/deque"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// JsonModule represents a module that can be extended by a template
type JsonModule struct {
	Name     *types.JsonString
	Scope    *types.JsonObject
	Template *types.JsonObject
	Copy     *types.JsonString
}

type TemplateVisitor struct {
	scope       deque.Deque[*types.JsonObject]
	moduleScope deque.Deque[*types.JsonObject]
	globalScope *types.JsonObject
	deadline    int64
}

const MaxInt64 = int64(^uint64(0) >> 1)

// LoadModule loads a module from a file and returns a JsonModule
func LoadModule(input string, globalScope *types.JsonObject, timeout int64) (JsonModule, error) {
	// Set up the deadline
	deadline := time.Now().UnixMilli() + timeout
	if timeout <= 0 {
		deadline = MaxInt64
	}

	jsonObject, err := types.ParseJsonObject([]byte(input))
	if err != nil {
		return JsonModule{}, burrito.WrapErrorf(err, "Failed to parse JSON module")
	}
	moduleName, err := FindAnyCase[*types.JsonString](jsonObject, "$module")
	if err != nil {
		return JsonModule{}, utils.WrapJsonErrorf("$module", err, "Invalid $module")
	}
	s, err := FindAnyCase[*types.JsonObject](jsonObject, "$scope")
	if err != nil && burrito.HasTag(err, WrongTypeErrTag) {
		return JsonModule{}, utils.WrapJsonErrorf("$scope", err, "Invalid $scope")
	}
	scope := types.NewJsonObject()
	if err == nil {
		tempScope := deque.Deque[*types.JsonObject]{}
		tempScope.PushBack(globalScope)
		tempScope.PushBack(scope)
		scopeVisitor := TemplateVisitor{
			scope:       tempScope,
			globalScope: globalScope,
			deadline:    deadline,
		}
		object, err := scopeVisitor.visitObject(types.DeepCopyObject(*s), "$scope")
		if err != nil {
			return JsonModule{}, burrito.WrapErrorf(err, "Failed to template scope")
		}
		scope = types.MergeObject(object.(*types.JsonObject), scope, false, "#")
	}
	template, err := FindAnyCase[*types.JsonObject](jsonObject, "$template")
	if err != nil && burrito.HasTag(err, WrongTypeErrTag) {
		return JsonModule{}, utils.WrapJsonErrorf("$template", err, "Invalid $template")
	} else if err != nil {
		empty := types.NewJsonObject()
		template = &empty
	}
	c, err := FindAnyCase[*types.JsonString](jsonObject, "$copy")
	if err != nil && burrito.HasTag(err, WrongTypeErrTag) {
		return JsonModule{}, utils.WrapJsonErrorf("$copy", err, "Invalid $copy")
	}
	if err != nil {
		temp := types.NewString("")
		c = &temp
	}
	return JsonModule{
		Name:     *moduleName,
		Scope:    scope,
		Template: *template,
		Copy:     *c,
	}, nil
}

func ProcessAssertionsFile(name, input string, globalScope *types.JsonObject, timeout int64) error {
	// Set up the deadline
	deadline := time.Now().UnixMilli() + timeout
	if timeout <= 0 {
		deadline = MaxInt64
	}

	// Parse the input
	root, err := types.ParseJsonArray([]byte(input))
	if err != nil {
		return burrito.WrapErrorf(err, "Failed to parse JSON")
	}

	scope := deque.Deque[*types.JsonObject]{}
	scope.PushBack(globalScope)

	visitor := TemplateVisitor{
		scope:       scope,
		globalScope: globalScope,
		deadline:    deadline,
	}

	err = visitor.visitAssert(root, name+"#/")
	if err != nil {
		return burrito.PassError(err)
	}
	return nil
}

// Process processes a template and returns a map of the processed templates
func Process(name, input string, globalScope *types.JsonObject, modules map[string]JsonModule, timeout int64) (utils.NavigableMap[string, types.JsonType], error) {
	// Set up the deadline
	deadline := time.Now().UnixMilli() + timeout
	if timeout <= 0 {
		deadline = MaxInt64
	}

	// Parse the input
	result := utils.NewNavigableMap[string, types.JsonType]()
	root, err := types.ParseJsonObject([]byte(input))
	if err != nil {
		return result, burrito.WrapErrorf(err, "Failed to parse JSON")
	}

	// Define scope
	scope := types.DeepCopyObject(globalScope)
	s, err := FindAnyCase[*types.JsonObject](root, "$scope")
	if err != nil && burrito.HasTag(err, WrongTypeErrTag) {
		return result, utils.WrapJsonErrorf("$scope", err, "Invalid $scope")
	}
	if err == nil {
		tempScope := deque.Deque[*types.JsonObject]{}
		tempScope.PushBack(scope)
		scopeVisitor := TemplateVisitor{
			scope:       tempScope,
			globalScope: globalScope,
			deadline:    deadline,
		}
		object, err := scopeVisitor.visitObject(types.DeepCopyObject(*s), "$scope")
		if err != nil {
			return result, burrito.WrapErrorf(err, "Failed to template scope")
		}
		scope = types.MergeObject(object.(*types.JsonObject), scope, false, "#")
	}

	c, err := FindAnyCase[types.JsonType](root, "$copy")
	if err != nil && burrito.HasTag(err, WrongTypeErrTag) {
		return result, utils.WrapJsonErrorf("$copy", err, "Invalid $copy")
	}
	isCopy := err == nil
	tempExtend, err := FindAnyCase[types.JsonType](root, "$extend")
	isExtend := err == nil
	temp, err := FindAnyCase[types.JsonType](root, "$template")
	if err != nil && burrito.HasTag(err, WrongTypeErrTag) {
		return result, utils.WrapJsonErrorf("$template", err, "Invalid $template")
	}
	hasTemplate := err == nil

	// If none of the options are defined, return unmodified JSON
	if !hasTemplate && !isCopy && !isExtend {
		result.Put(name, root)
		return result, nil
	}

	visitor := TemplateVisitor{
		scope:       deque.Deque[*types.JsonObject]{},
		globalScope: globalScope,
		deadline:    deadline,
	}
	visitor.pushScope(types.DeepCopyObject(scope))
	visitor.pushScope(types.AsObject(map[string]interface{}{"$allModules": modules}))
	var template types.JsonType = types.Null

	// handle generating multiple files
	file, err := FindAnyCase[*types.JsonObject](root, "$files")
	if err != nil && burrito.HasTag(err, WrongTypeErrTag) {
		return result, utils.WrapJsonErrorf("$files", err, "Invalid $file field")
	}
	if err == nil {
		arrayExpression, err := FindAnyCase[*types.JsonString](*file, "array")
		if err != nil {
			return result, utils.WrapJsonErrorf("$files", err, "Invalid array expression")
		}
		array, err := Eval((*arrayExpression).StringValue(), visitor.getScope(), "$files.array")
		if err != nil {
			return result, burrito.WrapErrorf(err, "Failed to evaluate $files.array")
		}
		if array.Value == nil {
			return result, utils.WrappedJsonErrorf("$files.array", "The array evaluated to null")
		}
		if arr, ok := array.Value.(*types.JsonArray); ok {
			for i, item := range arr.Value {
				err := checkDeadline(deadline)
				if err != nil {
					return result, burrito.PassError(err)
				}
				extra := types.AsObject(map[string]interface{}{
					"index": types.AsNumber(i),
					"value": item,
				})
				visitor.pushScope(extra)
				if obj, ok := item.(*types.JsonObject); ok {
					visitor.pushScope(obj)
				}
				if isCopy {
					template, err = processCopy(*c, visitor, modules, "$files.array", timeout)
					if err != nil {
						return result, burrito.PassError(err)
					}
					visitor.pushScope(types.AsObject(map[string]interface{}{"$copy": template}))
				} else {
					template = *temp
				}
				mappedModules := map[string]JsonModule{}
				if isExtend {
					if types.IsNull(template) {
						template = types.NewJsonObject()
					} else if _, ok := template.(*types.JsonObject); !ok {
						return result, utils.WrappedJsonErrorf("$extend", "The extend option is not yet supported for types other than objects")
					}
					var resolvedModules []string
					template, resolvedModules, err = extendTemplate(*tempExtend, template.(*types.JsonObject), &visitor, modules)
					if err != nil {
						return result, burrito.PassError(err)
					}

					for _, module := range resolvedModules {
						mappedModules[module] = modules[module]
					}
					visitor.pushScope(types.AsObject(map[string]interface{}{"$modules": mappedModules}))
				}
				if isCopy && hasTemplate {
					template, err = types.MergeValues(template, *temp, true, "#")
					if err != nil {
						return result, burrito.PassError(err)
					}
				}
				fName, err := FindAnyCase[*types.JsonString](*file, "file", "name")
				if err != nil {
					return result, utils.WrapJsonErrorf("$files", err, "Invalid file name")
				}
				mFileName, err := visitor.visitString((*fName).StringValue(), "$files.filename")
				if err != nil {
					return result, burrito.WrapErrorf(err, "Failed to evaluate $files.filename")
				}
				var f types.JsonType
				if v, ok := template.(*types.JsonObject); ok {
					f, err = visitor.visitObject(types.DeepCopyObject(v), "$template")
					if err != nil {
						return result, burrito.PassError(err)
					}
				} else if v, ok := template.(*types.JsonArray); ok {
					f, err = visitor.visitArray(types.DeepCopyArray(v), "$template")
					if err != nil {
						return result, burrito.PassError(err)
					}
				} else {
					return result, utils.WrappedJsonErrorf("$template", "Attempted to template unsupported type. Currently only objects and arrays are supported")
				}
				if isExtend {
					for i := 0; i < len(mappedModules); i++ {
						visitor.moduleScope.PopBack()
					}
				}
				if isCopy {
					visitor.popScope()
				}
				visitor.popScope()
				if _, ok := item.(*types.JsonObject); ok {
					visitor.popScope()
				}
				result.Put(mFileName.StringValue(), types.MergeObject(types.NewJsonObject(), f.(*types.JsonObject), false, "#"))
				result.Put(mFileName.StringValue(), types.DeleteNulls(result.Get(mFileName.StringValue())))
			}
		} else {
			return result, utils.WrappedJsonErrorf("$files.array", "The array evaluated to a non-array")
		}
	} else {
		if isCopy {
			template, err = processCopy(*c, visitor, modules, "$copy", timeout)
			if err != nil {
				return result, burrito.PassError(err)
			}
			visitor.pushScope(types.AsObject(map[string]interface{}{"$copy": template}))
		}
		mappedModules := map[string]JsonModule{}
		if isExtend {
			if types.IsNull(template) {
				template = types.NewJsonObject()
			} else if _, ok := template.(*types.JsonObject); !ok {
				return result, utils.WrappedJsonErrorf("$extend", "The extend option is not yet supported for types other than objects")
			}
			var resolvedModules []string
			template, resolvedModules, err = extendTemplate(*tempExtend, template.(*types.JsonObject), &visitor, modules)
			if err != nil {
				return result, burrito.PassError(err)
			}

			for _, module := range resolvedModules {
				mappedModules[module] = modules[module]
			}
			visitor.pushScope(types.AsObject(map[string]interface{}{"$modules": mappedModules}))
		}
		if hasTemplate {
			template, err = types.MergeValues(template, *temp, true, "#")
			if err != nil {
				return result, burrito.PassError(err)
			}
		}
		var f types.JsonType
		if v, ok := template.(*types.JsonObject); ok {
			f, err = visitor.visitObject(types.DeepCopyObject(v), "$template")
			if err != nil {
				return result, burrito.PassError(err)
			}
			result.Put(name, types.MergeObject(types.NewJsonObject(), f.(*types.JsonObject), false, "#"))
		} else if v, ok := template.(*types.JsonArray); ok {
			f, err = visitor.visitArray(types.DeepCopyArray(v), "$template")
			if err != nil {
				return result, burrito.PassError(err)
			}
			result.Put(name, types.MergeArray(types.NewJsonArray(), f.(*types.JsonArray), false, "#"))
		} else {
			return result, utils.WrappedJsonErrorf("$template", "Attempted to template unsupported type. Currently only objects and arrays are supported")
		}
		if isExtend {
			for i := 0; i < len(mappedModules); i++ {
				visitor.moduleScope.PopFront()
			}
		}
		if isCopy {
			visitor.popScope()
		}
		result.Put(name, types.DeleteNulls(result.Get(name)))
	}

	return result, nil
}

func processCopy(c types.JsonType, visitor TemplateVisitor, modules map[string]JsonModule, path string, timeout int64) (types.JsonType, error) {
	var result types.JsonType = types.Null
	copies := make([]*types.JsonString, 0)
	if copyArray, ok := c.(*types.JsonArray); ok {
		for i, item := range copyArray.Value {
			if copyPath, ok := item.(*types.JsonString); ok {
				copies = append(copies, copyPath)
			} else {
				return types.Null, utils.WrappedJsonErrorf(fmt.Sprintf("%s[%d]", path, i), "The copy path is not a string")
			}
		}
	} else if copyPath, ok := c.(*types.JsonString); ok {
		copies = append(copies, copyPath)
	} else {
		return types.Null, utils.WrappedJsonErrorf(path, "The copy path is not a string")
	}
	for i, c := range copies {
		loopPath := fmt.Sprintf("%s[%d]", path, i)
		c, err := visitor.visitString(c.StringValue(), path)
		if err != nil {
			return types.Null, utils.WrapJsonErrorf(loopPath, err, "Failed to evaluate $copy")
		}
		if c == nil || types.IsNull(c) {
			utils.Logger.Debugf("Skipping null copy path at %s", loopPath)
			continue
		}
		insideCopies := make([]*types.JsonString, 0)
		if copyArray, ok := c.(*types.JsonArray); ok {
			for j, inside := range copyArray.Value {
				if insideCopyPath, ok := inside.(*types.JsonString); ok {
					insideCopies = append(insideCopies, insideCopyPath)
				} else {
					return types.Null, utils.WrappedJsonErrorf(fmt.Sprintf("%s[%d]", loopPath, j), "The copy path inside evaluated copy array is not a string")
				}
			}
		} else if copyPath, ok := c.(*types.JsonString); ok {
			insideCopies = append(insideCopies, copyPath)
		} else {
			return types.Null, utils.WrappedJsonErrorf(loopPath, "The copy path evaluated to a non-string")
		}
		for _, copyPath := range insideCopies {
			resolve, err := safeio.Resolver.Open(copyPath.StringValue())
			if err != nil {
				return types.Null, utils.WrapJsonErrorf(loopPath, err, "Failed to open %s", copyPath.StringValue())
			}
			all, err := io.ReadAll(resolve)
			if err != nil {
				return types.Null, utils.WrapJsonErrorf(loopPath, err, "Failed to read %s", copyPath.StringValue())
			}
			if strings.HasSuffix(copyPath.StringValue(), ".templ") {
				processedMap, err := Process("copy", string(all), visitor.globalScope, modules, timeout)
				if err != nil {
					return types.Null, utils.WrapJsonErrorf(loopPath, err, "Failed to process copy template %s", copyPath.StringValue())
				}
				if processedMap.Size() > 1 {
					return types.Null, utils.WrappedJsonErrorf(path, "The copy template must compile to a single object")
				}
				template := processedMap.Get("copy")
				result, err = types.MergeValues(result, template, false, "#")
				if err != nil {
					return types.Null, utils.WrapJsonErrorf(loopPath, err, "Failed to merge %s", copyPath.StringValue())
				}
				continue
			} else {
				template, err := types.ParseJsonValue(all)
				if err != nil {
					return types.Null, utils.WrapJsonErrorf(loopPath, err, "Failed to parse %s", copyPath.StringValue())
				}
				result, err = types.MergeValues(result, template, false, "#")
				if err != nil {
					return types.Null, utils.WrapJsonErrorf(loopPath, err, "Failed to merge %s", copyPath.StringValue())
				}
				continue
			}
		}
	}
	return result, nil
}

var templatePattern, _ = regexp.Compile("\\{\\{(?:\\\\.|[^{}])+}}")

func extendTemplate(extend types.JsonType, template *types.JsonObject, visitor *TemplateVisitor, modules map[string]JsonModule) (*types.JsonObject, []string, error) {
	resolvedModules := make([]string, 0)
	toResolve := make([]string, 0)
	isString := true

	if arr, ok := extend.(*types.JsonArray); ok {
		isString = false
		for _, mod := range arr.Value {
			if str, ok := mod.(*types.JsonString); ok {
				toResolve = append(toResolve, str.StringValue())
			} else {
				return types.NewJsonObject(), resolvedModules, utils.WrappedJsonErrorf("$extend", "The extend array must contain only strings")
			}
		}
	} else if str, ok := extend.(*types.JsonString); ok {
		toResolve = append(toResolve, str.StringValue())
	} else {
		return types.NewJsonObject(), resolvedModules, utils.WrappedJsonErrorf("$extend", "The extend value must be a string or array of strings")
	}
	for i, str := range toResolve {
		path := "$extend"
		if !isString {
			path += "[" + strconv.Itoa(i) + "]"
		}
		matches, err := FindTemplateMatches(str, "{", "}")
		if err != nil {
			return types.NewJsonObject(), resolvedModules, utils.WrapJsonErrorf(path, err, "Failed to parse template %s", str)
		}
		if len(matches) > 0 {
			eval, err := visitor.visitString(str, path)
			if err != nil {
				return types.NewJsonObject(), resolvedModules, utils.WrapJsonErrorf(path, err, "Failed to evaluate %s", path)
			}
			if eval == nil || types.IsNull(eval) {
				utils.Logger.Debugf("Skipping null module at %s", path)
				continue
			}
			if mods, ok := eval.(*types.JsonArray); ok {
				stringMods := make([]string, len(mods.Value))
				for i, mod := range mods.Value {
					stringMods[i] = types.ToString(mod)
				}
				resolvedModules = append(resolvedModules, stringMods...)
			} else if strMod, ok := eval.(*types.JsonString); ok {
				resolvedModules = append(resolvedModules, strMod.StringValue())
			} else {
				return types.NewJsonObject(), resolvedModules, utils.WrappedJsonErrorf(path, "The module name evaluated to a non-string")
			}
		} else {
			resolvedModules = append(resolvedModules, str)
		}
	}
	for _, module := range resolvedModules {
		if _, ok := modules[module]; !ok {
			return types.NewJsonObject(), resolvedModules, utils.WrappedJsonErrorf("$extend", "The module '%s' does not exist", module)
		}
		mod := modules[module]
		visitor.moduleScope.PushBack(mod.Scope)
		if mod.Copy.StringValue() != "" {
			object, err := processCopy(mod.Copy, *visitor, modules, "$copy", -1)
			if err != nil {
				return types.NewJsonObject(), resolvedModules, burrito.WrapErrorf(err, "Error processing $copy for module %s", mod.Name.StringValue())
			}
			template = types.MergeObject(object.(*types.JsonObject), template, true, "#")
		}
		if !mod.Template.IsEmpty() {
			parent, err := visitor.visitObject(mod.Template, "[Module "+module+"]")
			if err != nil {
				return types.NewJsonObject(), resolvedModules, burrito.WrapErrorf(err, "Error processing template for module %s", mod.Name.StringValue())
			}
			template = types.MergeObject(template, parent.(*types.JsonObject), true, "#")
		}
	}
	return template, resolvedModules, nil
}

func (v *TemplateVisitor) pushScope(obj *types.JsonObject) {
	if obj == nil {
		v.scope.PushBack(types.NewJsonObject())
		return
	}
	v.scope.PushBack(obj)
}

func (v *TemplateVisitor) popScope() {
	v.scope.PopBack()
}

func (v *TemplateVisitor) getScope() deque.Deque[*types.JsonObject] {
	// combine all scopes
	result := deque.Deque[*types.JsonObject]{}
	for i := 0; i < v.moduleScope.Len(); i++ {
		result.PushBack(v.moduleScope.At(i))
	}
	for i := 0; i < v.scope.Len(); i++ {
		result.PushBack(v.scope.At(i))
	}
	return result
}

func (v *TemplateVisitor) visit(obj types.JsonType, path string) (types.JsonType, error) {
	if obj == nil {
		return nil, nil
	}
	if str, ok := obj.(*types.JsonString); ok {
		return v.visitString(str.StringValue(), path)
	}
	if arr, ok := obj.(*types.JsonArray); ok {
		return v.visitArray(arr, path)
	}
	if objMap, ok := obj.(*types.JsonObject); ok {
		return v.visitObject(objMap, path)
	}
	return obj, nil
}

func (v *TemplateVisitor) visitObject(obj *types.JsonObject, path string) (types.JsonType, error) {
	var result = types.NewJsonObjectWithCapacity(obj.Size())
	for _, key := range obj.Keys() {
		err := checkDeadline(v.deadline)
		if err != nil {
			return result, err
		}
		value := obj.Get(key)
		// The only keys, that can be with $ prefix are ones specified in MergeObject function
		if strings.EqualFold(key, "$comment") {
			continue
		}
		if strings.EqualFold(key, "$assert") {
			if err = v.visitAssert(value, fmt.Sprintf("%s/%s", path, "$assert")); err != nil {
				return result, err
			}
			continue
		}
		if strings.HasPrefix(key, "{{") && strings.HasSuffix(key, "}}") && (key[2] == '#' || key[2] == '?') {
			eval, err := Eval(key, v.getScope(), fmt.Sprintf("%s/%s", path, key))
			if err != nil {
				return nil, utils.WrapJsonErrorf(path, err, "Failed to evaluate %s", key)
			}
			switch eval.Action {
			case types.Iteration:
				if _, ok := value.(*types.JsonObject); !ok {
					return nil, utils.WrappedJsonErrorf(path, "The value of the iteration key must be an object")
				}
				if arr, ok := eval.Value.(*types.JsonArray); ok {
					for i, val := range arr.Value {
						v.pushScope(types.AsObject(map[string]interface{}{
							eval.IndexName: types.AsNumber(i),
							eval.Name:      val,
						}))
						if obj, ok := val.(*types.JsonObject); ok {
							v.pushScope(obj)
						}
						o, err := v.visit(value, fmt.Sprintf("%s/%s[%d]", path, key, i))
						v.popScope()
						if _, ok := val.(*types.JsonObject); ok {
							v.popScope()
						}
						if err != nil {
							return nil, burrito.PassError(err)
						}
						u := o.(*types.JsonObject)
						for _, k := range u.Keys() {
							err = PutValue(result, k, u.Get(k), path)
							if err != nil {
								return nil, burrito.PassError(err)
							}
						}
					}
				} else {
					return nil, utils.WrappedJsonErrorf(path, "Iteration action must evaluate to an array")
				}
			case types.Predicate:
				if eval.Value.BoolValue() {
					o, err := v.visit(value, fmt.Sprintf("%s/%s", path, key))
					if err != nil {
						return nil, burrito.PassError(err)
					}
					if obj, ok := o.(*types.JsonObject); ok {
						for _, k := range obj.Keys() {
							err = PutValue(result, k, obj.Get(k), path)
							if err != nil {
								return nil, burrito.PassError(err)
							}
						}
					} else {
						return nil, utils.WrappedJsonErrorf(path, "The value of the predicate key must be an object")
					}
				}
			case types.Value:
				k, err := v.visitString(key, fmt.Sprintf("%s/%s", path, key))
				if err != nil {
					return nil, burrito.PassError(err)
				}
				r, err := v.visit(value, fmt.Sprintf("%s/%s", path, key))
				if err != nil {
					return nil, burrito.PassError(err)
				}
				err = PutValue(result, k.StringValue(), r, path)
				if err != nil {
					return nil, burrito.PassError(err)
				}
			default:
				return nil, utils.WrappedJsonErrorf(path, "Unsupported action %s", eval.Action.String())
			}
		} else {
			matches, err := FindTemplateMatches(key, "{", "}")
			if err != nil {
				return nil, utils.WrapJsonErrorf(path, err, "Failed to parse template %s", key)
			}
			if len(matches) > 0 {
				k, err := v.visitString(key, fmt.Sprintf("%s/%s", path, key))
				if err != nil {
					return nil, burrito.PassError(err)
				}
				r, err := v.visit(value, fmt.Sprintf("%s/%s", path, key))
				if err != nil {
					return nil, burrito.PassError(err)
				}
				err = PutValue(result, k.StringValue(), r, path)
				if err != nil {
					return nil, burrito.PassError(err)
				}
			} else {
				var err error
				r, err := v.visit(value, fmt.Sprintf("%s/%s", path, key))
				if err != nil {
					return nil, burrito.PassError(err)
				}
				err = PutValue(result, key, r, path)
				if err != nil {
					return nil, burrito.PassError(err)
				}
			}
		}
	}
	return result, nil
}

func PutValue(result *types.JsonObject, key string, r types.JsonType, path string) error {
	if r == nil || types.IsNull(r) {
		result.Put(key, types.Null)
	} else {
		if strings.HasPrefix(key, "$") && result.ContainsKey(key) && !types.IsReservedKey(key) {
			result.Put(key, r)
		} else {
			json, err := types.MergeJSON(result.Get(key), r, true)
			if err != nil {
				return utils.WrapJsonErrorf(path, err, "Failed to merge %s", key)
			}
			result.Put(key, json)
		}
	}
	return nil
}

func (v *TemplateVisitor) visitArray(arr *types.JsonArray, path string) (types.JsonType, error) {
	var result = make([]types.JsonType, 0, len(arr.Value))
	for i, value := range arr.Value {
		err := checkDeadline(v.deadline)
		if err != nil {
			return &types.JsonArray{Value: result}, err
		}
		a, err := v.visitArrayElement(result, value, fmt.Sprintf("%s[%d]", path, i))
		if err != nil {
			return nil, burrito.PassError(err)
		}
		result = a
	}
	return &types.JsonArray{Value: result}, nil
}

// Special visitor for cases when the array element is an object, that generates multiple values
func (v *TemplateVisitor) visitArrayElement(array []types.JsonType, element types.JsonType, path string) ([]types.JsonType, error) {
	if obj, ok := element.(*types.JsonObject); ok {
		filteredKeys := make([]string, 0, obj.Size())
		for _, key := range obj.Keys() {
			if key != "$assert" && key != "$comment" {
				filteredKeys = append(filteredKeys, key)
			}
		}
		if len(filteredKeys) == 1 || obj.Size() == 1 {
			for _, key := range obj.Keys() {
				value := obj.Get(key)
				if strings.HasPrefix(key, "{{") && strings.HasSuffix(key, "}}") && (key[2] == '#' || key[2] == '?') {
					eval, err := Eval(key, v.getScope(), path)
					if err != nil {
						return array, utils.WrapJsonErrorf(path, err, "Failed to evaluate %s", key)
					}
					switch eval.Action {
					case types.Iteration:
						if arr, ok := eval.Value.(*types.JsonArray); ok {
							for i, val := range arr.Value {
								v.pushScope(types.AsObject(map[string]interface{}{
									eval.IndexName: types.AsNumber(i),
									eval.Name:      val,
								}))
								if obj, ok := val.(*types.JsonObject); ok {
									v.pushScope(obj)
								}
								a, err := v.visitArrayElement(array, value, fmt.Sprintf("%s[%d]", path, i))
								array = a
								v.popScope()
								if _, ok := val.(*types.JsonObject); ok {
									v.popScope()
								}
								if err != nil {
									return array, burrito.PassError(err)
								}
							}
							return array, nil
						} else {
							return nil, utils.WrappedJsonErrorf(path, "Iteration action must evaluate to an array")
						}
					case types.Predicate:
						if eval.Value.BoolValue() {
							return v.visitArrayElement(array, value, path)
						}
						return array, nil
					case types.Value:
						visit, err := v.visit(element, path)
						if err != nil {
							return nil, burrito.PassError(err)
						}
						array = append(array, visit)
						return array, nil
					default:
						return nil, utils.WrappedJsonErrorf(path, "Unsupported action %s", eval.Action.String())
					}
				} else if key == "$assert" {
					err := v.visitAssert(value, path)
					if err != nil {
						return nil, burrito.PassError(err)
					}
					return nil, nil
				} else if key == "$comment" {
					return nil, nil
				}
			}
		}
	}
	visit, err := v.visit(element, path)
	if err != nil {
		return nil, burrito.PassError(err)
	}
	// If the result is an array, append all elements, but only if they are not equal to the original element
	if arr, ok := visit.(*types.JsonArray); ok && !visit.Equals(element) {
		array = append(array, arr.Value...)
	} else {
		array = append(array, visit)
	}
	return array, nil
}

func (v *TemplateVisitor) visitString(str string, path string) (types.JsonType, error) {
	err := checkDeadline(v.deadline)
	if err != nil {
		return nil, err
	}
	matches, err := FindTemplateMatches(str, "{", "}")
	if err != nil {
		return nil, burrito.WrapErrorf(err, "Failed to parse string '%s'", str)
	}
	var sb strings.Builder
	lastMatchEnd := 0
	strRunes := []rune(str)
	for _, match := range matches {
		if match.Start > lastMatchEnd {
			sb.WriteString(string(strRunes[lastMatchEnd:match.Start]))
		}
		result, err := Eval(match.EscapedMatch, v.getScope(), path)
		if err != nil {
			return nil, burrito.WrapErrorf(err, "Error evaluating '%s'", match.EscapedMatch)
		}
		if result.Value == nil {
			return nil, utils.WrappedJsonErrorf(path, "The expression '%s' evaluated to null", match.EscapedMatch)
		}
		if result.Action == types.Literal {
			return result.Value, nil
		} else if result.Action == types.Value {
			// Moved this check here to avoid not checking the action
			if _, ok := result.Value.(*types.JsonString); !ok && str == match.EscapedMatch {
				return result.Value, nil
			}
			sb.WriteString(types.ToString(result.Value))
			lastMatchEnd = match.Start + match.Length + 1
		} else {
			return nil, utils.WrappedJsonErrorf(path, "Unsupported action %s", result.Action.String())
		}
	}
	if lastMatchEnd < len(strRunes) {
		sb.WriteString(string(strRunes[lastMatchEnd:]))
	}
	return types.NewString(sb.String()), nil
}

func (v *TemplateVisitor) visitAssert(value interface{}, path string) error {
	if arr, ok := value.(*types.JsonArray); ok {
		for i, i2 := range arr.Value {
			err := v.visitAssert(i2, fmt.Sprintf("%s[%d]", path, i))
			if err != nil {
				return burrito.PassError(err)
			}
		}
	} else if str, ok := value.(*types.JsonString); ok {
		result, err := Eval(str.StringValue(), v.getScope(), path)
		if err != nil {
			return utils.WrapJsonErrorf(path, err, "Error evaluating '%s'", str.StringValue())
		}
		if result.Action != types.Value {
			return utils.WrappedJsonErrorf(path, "Unsupported action %s", result.Action.String())
		}
		if !result.Value.BoolValue() {
			return utils.WrappedJsonErrorf(path, "Assertion failed for '%s'", str.StringValue())
		}
	} else if obj, ok := value.(*types.JsonObject); ok {
		condition, err := FindAnyCase[*types.JsonString](obj, "condition")
		if err != nil {
			return utils.WrapJsonErrorf(path, err, "Invalid condition")
		}
		result, err := Eval((*condition).StringValue(), v.getScope(), path)
		if err != nil {
			return utils.WrapJsonErrorf(path, err, "Error evaluating '%s'", (*condition).StringValue())
		}
		if result.Action != types.Value {
			return utils.WrappedJsonErrorf(path, "Unsupported action %s", result.Action.String())
		}
		if !result.Value.BoolValue() {
			message, err := FindAnyCase[*types.JsonString](obj, "message")
			if err != nil {
				if burrito.HasTag(err, WrongTypeErrTag) {
					return utils.WrapJsonErrorf(path, err, "Invalid assertion message")
				}
				return utils.WrapJsonErrorf(path, err, "Assertion failed for '%s'", str.StringValue())
			}
			msg, err := v.visitString((*message).StringValue(), path)
			if err != nil {
				return utils.WrappedJsonErrorf(path, "Error evaluating message '%s'", (*message).StringValue())
			}
			return utils.WrappedJsonErrorf(path, types.ToPrettyString(msg))
		}
	}
	return nil
}

func checkDeadline(deadline int64) error {
	if time.Now().UnixMilli() > deadline {
		return burrito.WrappedErrorf("Template evaluation timed out")
	}
	return nil
}
