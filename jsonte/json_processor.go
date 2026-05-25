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
	"sync"
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
	// keysBuffer is a reusable scratch buffer for visitArrayElement.
	// Safe to share because the function always returns before re-entering the range loop.
	keysBuffer []string
	// pathBuf holds the current JSON path as a growing byte buffer.
	// Callers set it via pathSet/pathPush/pathPop to avoid per-key string allocations.
	pathBuf []byte
}

const MaxInt64 = int64(^uint64(0) >> 1)

// strBuf wraps a []byte so the pool can recycle the backing array across calls.
// strings.Builder.Reset() sets buf=nil, destroying capacity; this avoids that.
type strBuf struct{ b []byte }

var stringBuilderPool = sync.Pool{
	New: func() interface{} { return &strBuf{b: make([]byte, 0, 128)} },
}

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
		scopeVisitor.pathSet("$scope")
		object, err := scopeVisitor.visitObject(types.DeepCopyObject(*s))
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

	visitor.pathSet(name + "#/")
	err = visitor.visitAssert(root)
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
		scopeVisitor.pathSet("$scope")
		object, err := scopeVisitor.visitObject(types.DeepCopyObject(*s))
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
		array, err := evalWithExtraScope((*arrayExpression).StringValue(), &visitor.scope, &visitor.moduleScope)
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
				visitor.pathSet("$files.filename")
				mFileName, err := visitor.visitString((*fName).StringValue())
				if err != nil {
					return result, burrito.WrapErrorf(err, "Failed to evaluate $files.filename")
				}
				var f types.JsonType
				if v, ok := template.(*types.JsonObject); ok {
					visitor.pathSet("$template")
					f, err = visitor.visitObject(types.DeepCopyObject(v))
					if err != nil {
						return result, burrito.PassError(err)
					}
				} else if v, ok := template.(*types.JsonArray); ok {
					visitor.pathSet("$template")
					f, err = visitor.visitArray(types.DeepCopyArray(v))
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
			visitor.pathSet("$template")
			f, err = visitor.visitObject(types.DeepCopyObject(v))
			if err != nil {
				return result, burrito.PassError(err)
			}
			result.Put(name, types.MergeObject(types.NewJsonObject(), f.(*types.JsonObject), false, "#"))
		} else if v, ok := template.(*types.JsonArray); ok {
			visitor.pathSet("$template")
			f, err = visitor.visitArray(types.DeepCopyArray(v))
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
		visitor.pathBuf = []byte(path)
		c, err := visitor.visitString(c.StringValue())
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
			visitor.pathBuf = []byte(path)
			eval, err := visitor.visitString(str)
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
			visitor.pathBuf = []byte("[Module " + module + "]")
			parent, err := visitor.visitObject(mod.Template)
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

// pathSet initialises pathBuf to path, using a fresh allocation to avoid sharing backing arrays with callers that pass visitor by value.
func (v *TemplateVisitor) pathSet(path string) {
	v.pathBuf = []byte(path)
}

// pathPush appends "/" + seg to pathBuf and returns the old length for later restoration.
func (v *TemplateVisitor) pathPush(seg string) int {
	n := len(v.pathBuf)
	v.pathBuf = append(v.pathBuf, '/')
	v.pathBuf = append(v.pathBuf, seg...)
	return n
}

// pathPushIndex appends "[i]" to pathBuf and returns the old length.
func (v *TemplateVisitor) pathPushIndex(i int) int {
	n := len(v.pathBuf)
	v.pathBuf = append(v.pathBuf, '[')
	v.pathBuf = strconv.AppendInt(v.pathBuf, int64(i), 10)
	v.pathBuf = append(v.pathBuf, ']')
	return n
}

// pathPop restores pathBuf to the length saved by a prior pathPush/pathPushIndex.
func (v *TemplateVisitor) pathPop(n int) {
	v.pathBuf = v.pathBuf[:n]
}

// pathStr returns the current path as a string. Only call in error paths to avoid allocs.
func (v *TemplateVisitor) pathStr() string {
	return string(v.pathBuf)
}

// pathStrAt returns pathBuf[:n] as a string (the parent path). Only call in error paths.
func (v *TemplateVisitor) pathStrAt(n int) string {
	return string(v.pathBuf[:n])
}

func (v *TemplateVisitor) visit(obj types.JsonType) (types.JsonType, error) {
	if obj == nil {
		return nil, nil
	}
	if str, ok := obj.(*types.JsonString); ok {
		return v.visitString(str.StringValue())
	}
	if arr, ok := obj.(*types.JsonArray); ok {
		return v.visitArray(arr)
	}
	if objMap, ok := obj.(*types.JsonObject); ok {
		return v.visitObject(objMap)
	}
	return obj, nil
}

func (v *TemplateVisitor) visitObject(obj *types.JsonObject) (types.JsonType, error) {
	var result = types.NewJsonObjectWithCapacity(obj.Size())
	var outerErr error
	nilResult := false
	parentPathLen := len(v.pathBuf)

	obj.RangeEntries(func(key string, value types.JsonType) bool {
		if err := checkDeadline(v.deadline); err != nil {
			outerErr = err
			return true
		}
		// The only keys, that can be with $ prefix are ones specified in MergeObject function
		if strings.EqualFold(key, "$comment") {
			return false
		}
		if strings.EqualFold(key, "$assert") {
			n := v.pathPush("$assert")
			err := v.visitAssert(value)
			v.pathPop(n)
			if err != nil {
				outerErr = err
				return true
			}
			return false
		}
		if strings.HasPrefix(key, "{{") && strings.HasSuffix(key, "}}") && len(key) > 2 && (key[2] == '#' || key[2] == '?') {
			nk := v.pathPush(key)
			eval, err := evalWithExtraScope(key, &v.scope, &v.moduleScope)
			if err != nil {
				v.pathPop(nk)
				nilResult = true
				outerErr = utils.WrapJsonErrorf(v.pathStrAt(parentPathLen), err, "Failed to evaluate %s", key)
				return true
			}
			switch eval.Action {
			case types.Iteration:
				if _, ok := value.(*types.JsonObject); !ok {
					v.pathPop(nk)
					nilResult = true
					outerErr = utils.WrappedJsonErrorf(v.pathStrAt(parentPathLen), "The value of the iteration key must be an object")
					return true
				}
				if arr, ok := eval.Value.(*types.JsonArray); ok {
					for i, val := range arr.Value {
						scopeObj := getScopeObject()
						scopeObj.Put(eval.IndexName, types.NewIntNumber(int64(i)))
						scopeObj.Put(eval.Name, val)
						v.pushScope(scopeObj)
						if innerObj, ok := val.(*types.JsonObject); ok {
							v.pushScope(innerObj)
						}
						ni := v.pathPushIndex(i)
						o, iterErr := v.visit(value)
						v.pathPop(ni)
						if _, ok := val.(*types.JsonObject); ok {
							v.popScope()
						}
						v.popScope()
						putScopeObject(scopeObj)
						if iterErr != nil {
							v.pathPop(nk)
							nilResult = true
							outerErr = burrito.PassError(iterErr)
							return true
						}
						u := o.(*types.JsonObject)
						iterKeys := 0
						u.RangeEntries(func(k string, v2 types.JsonType) bool {
							iterKeys++
							iterErr = PutValue(result, k, v2, v, parentPathLen)
							return iterErr != nil
						})
						if iterErr != nil {
							v.pathPop(nk)
							nilResult = true
							outerErr = burrito.PassError(iterErr)
							return true
						}
						if i == 0 && len(arr.Value) > 1 && iterKeys > 0 {
							result.Value.Reserve((len(arr.Value) - 1) * iterKeys)
						}
					}
				} else {
					v.pathPop(nk)
					nilResult = true
					outerErr = utils.WrappedJsonErrorf(v.pathStrAt(parentPathLen), "Iteration action must evaluate to an array")
					return true
				}
			case types.Predicate:
				if eval.Value.BoolValue() {
					o, err := v.visit(value)
					if err != nil {
						v.pathPop(nk)
						nilResult = true
						outerErr = burrito.PassError(err)
						return true
					}
					if predObj, ok := o.(*types.JsonObject); ok {
						var predErr error
						predObj.RangeEntries(func(k string, v2 types.JsonType) bool {
							predErr = PutValue(result, k, v2, v, parentPathLen)
							return predErr != nil
						})
						if predErr != nil {
							v.pathPop(nk)
							nilResult = true
							outerErr = burrito.PassError(predErr)
							return true
						}
					} else {
						v.pathPop(nk)
						nilResult = true
						outerErr = utils.WrappedJsonErrorf(v.pathStrAt(parentPathLen), "The value of the predicate key must be an object")
						return true
					}
				}
			case types.Value:
				k, err := v.visitString(key)
				if err != nil {
					v.pathPop(nk)
					nilResult = true
					outerErr = burrito.PassError(err)
					return true
				}
				r, err := v.visit(value)
				if err != nil {
					v.pathPop(nk)
					nilResult = true
					outerErr = burrito.PassError(err)
					return true
				}
				if err = PutValue(result, k.StringValue(), r, v, parentPathLen); err != nil {
					v.pathPop(nk)
					nilResult = true
					outerErr = burrito.PassError(err)
					return true
				}
			default:
				v.pathPop(nk)
				nilResult = true
				outerErr = utils.WrappedJsonErrorf(v.pathStrAt(parentPathLen), "Unsupported action %s", eval.Action.String())
				return true
			}
			v.pathPop(nk)
		} else {
			matches, err := FindTemplateMatches(key, "{", "}")
			if err != nil {
				nilResult = true
				outerErr = utils.WrapJsonErrorf(v.pathStrAt(parentPathLen), err, "Failed to parse template %s", key)
				return true
			}
			if len(matches) > 0 {
				nk := v.pathPush(key)
				k, err := v.visitString(key)
				if err != nil {
					v.pathPop(nk)
					nilResult = true
					outerErr = burrito.PassError(err)
					return true
				}
				r, err := v.visit(value)
				if err != nil {
					v.pathPop(nk)
					nilResult = true
					outerErr = burrito.PassError(err)
					return true
				}
				v.pathPop(nk)
				if err = PutValue(result, k.StringValue(), r, v, parentPathLen); err != nil {
					nilResult = true
					outerErr = burrito.PassError(err)
					return true
				}
			} else {
				// HOT PATH: plain key, no template substitution — push/pop avoids string allocation
				nk := v.pathPush(key)
				r, err := v.visit(value)
				v.pathPop(nk)
				if err != nil {
					nilResult = true
					outerErr = burrito.PassError(err)
					return true
				}
				if err = PutValue(result, key, r, v, parentPathLen); err != nil {
					nilResult = true
					outerErr = burrito.PassError(err)
					return true
				}
			}
		}
		return false
	})

	if outerErr != nil {
		if nilResult {
			return nil, outerErr
		}
		return result, outerErr
	}
	return result, nil
}

func PutValue(result *types.JsonObject, key string, r types.JsonType, v *TemplateVisitor, parentPathLen int) error {
	if r == nil || types.IsNull(r) {
		result.Put(key, types.Null)
	} else {
		existing, hasExisting := result.TryGet(key)
		if !hasExisting || existing == nil || types.IsNull(existing) {
			result.Put(key, r)
			return nil
		}

		if strings.HasPrefix(key, "$") && !types.IsReservedKey(key) {
			result.Put(key, r)
			return nil
		}

		if types.IsObject(existing) && types.IsObject(r) {
			merged, err := types.MergeJSON(existing, r, true)
			if err != nil {
				return utils.WrapJsonErrorf(v.pathStrAt(parentPathLen), err, "Failed to merge %s", key)
			}
			result.Put(key, merged)
			return nil
		}

		if types.IsArray(existing) && types.IsArray(r) {
			merged, err := types.MergeJSON(existing, r, true)
			if err != nil {
				return utils.WrapJsonErrorf(v.pathStrAt(parentPathLen), err, "Failed to merge %s", key)
			}
			result.Put(key, merged)
			return nil
		}

		result.Put(key, r)
	}
	return nil
}

func (v *TemplateVisitor) visitArray(arr *types.JsonArray) (types.JsonType, error) {
	var result = make([]types.JsonType, 0, len(arr.Value))
	for i, value := range arr.Value {
		err := checkDeadline(v.deadline)
		if err != nil {
			return &types.JsonArray{Value: result}, err
		}
		n := v.pathPushIndex(i)
		a, err := v.visitArrayElement(result, value)
		v.pathPop(n)
		if err != nil {
			return nil, burrito.PassError(err)
		}
		result = a
	}
	return &types.JsonArray{Value: result}, nil
}

// Special visitor for cases when the array element is an object, that generates multiple values
func (v *TemplateVisitor) visitArrayElement(array []types.JsonType, element types.JsonType) ([]types.JsonType, error) {
	if obj, ok := element.(*types.JsonObject); ok {
		// Reuse keysBuffer to avoid allocating a new []string each call.
		// Safe: the loop always returns before v.keysBuffer could be re-overwritten by recursion.
		v.keysBuffer = obj.AppendKeys(v.keysBuffer[:0])
		keys := v.keysBuffer
		filteredCount := 0
		for _, key := range keys {
			if key != "$assert" && key != "$comment" {
				filteredCount++
			}
		}
		if filteredCount == 1 || obj.Size() == 1 {
			for _, key := range keys {
				value := obj.Get(key)
				if strings.HasPrefix(key, "{{") && strings.HasSuffix(key, "}}") && len(key) > 2 && (key[2] == '#' || key[2] == '?') {
					eval, err := evalWithExtraScope(key, &v.scope, &v.moduleScope)
					if err != nil {
						return array, utils.WrapJsonErrorf(v.pathStr(), err, "Failed to evaluate %s", key)
					}
					switch eval.Action {
					case types.Iteration:
						if arr, ok := eval.Value.(*types.JsonArray); ok {
							for i, val := range arr.Value {
								scopeObj := getScopeObject()
								scopeObj.Put(eval.IndexName, types.NewIntNumber(int64(i)))
								scopeObj.Put(eval.Name, val)
								v.pushScope(scopeObj)
								if innerObj, ok := val.(*types.JsonObject); ok {
									v.pushScope(innerObj)
								}
								n := v.pathPushIndex(i)
								a, err := v.visitArrayElement(array, value)
								v.pathPop(n)
								array = a
								if _, ok := val.(*types.JsonObject); ok {
									v.popScope()
								}
								v.popScope()
								putScopeObject(scopeObj)
								if err != nil {
									return array, burrito.PassError(err)
								}
							}
							return array, nil
						} else {
							return nil, utils.WrappedJsonErrorf(v.pathStr(), "Iteration action must evaluate to an array")
						}
					case types.Predicate:
						if eval.Value.BoolValue() {
							return v.visitArrayElement(array, value)
						}
						return array, nil
					case types.Value:
						visit, err := v.visit(element)
						if err != nil {
							return nil, burrito.PassError(err)
						}
						array = append(array, visit)
						return array, nil
					default:
						return nil, utils.WrappedJsonErrorf(v.pathStr(), "Unsupported action %s", eval.Action.String())
					}
				} else if key == "$assert" {
					err := v.visitAssert(value)
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
	visit, err := v.visit(element)
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

func (v *TemplateVisitor) visitString(str string) (types.JsonType, error) {
	err := checkDeadline(v.deadline)
	if err != nil {
		return nil, err
	}
	matches, err := FindTemplateMatches(str, "{", "}")
	if err != nil {
		return nil, burrito.WrapErrorf(err, "Failed to parse string '%s'", str)
	}
	if len(matches) == 0 {
		return types.NewString(str), nil
	}
	sbp := stringBuilderPool.Get().(*strBuf)
	buf := sbp.b[:0]
	if cap(buf) < len(str) {
		buf = make([]byte, 0, len(str))
	}
	lastMatchEnd := 0
	for _, match := range matches {
		if match.StartByte > lastMatchEnd {
			buf = append(buf, str[lastMatchEnd:match.StartByte]...)
		}
		result, err := evalWithExtraScope(match.EscapedMatch, &v.scope, &v.moduleScope)
		if err != nil {
			sbp.b = buf
			stringBuilderPool.Put(sbp)
			return nil, burrito.WrapErrorf(err, "Error evaluating '%s'", match.EscapedMatch)
		}
		if result.Value == nil {
			sbp.b = buf
			stringBuilderPool.Put(sbp)
			return nil, utils.WrappedJsonErrorf(v.pathStr(), "The expression '%s' evaluated to null", match.EscapedMatch)
		}
		if result.Action == types.Literal {
			sbp.b = buf
			stringBuilderPool.Put(sbp)
			return result.Value, nil
		} else if result.Action == types.Value {
			// Moved this check here to avoid not checking the action
			if _, ok := result.Value.(*types.JsonString); !ok && str == match.EscapedMatch {
				sbp.b = buf
				stringBuilderPool.Put(sbp)
				return result.Value, nil
			}
			buf = append(buf, types.ToString(result.Value)...)
			lastMatchEnd = match.EndByte
		} else {
			sbp.b = buf
			stringBuilderPool.Put(sbp)
			return nil, utils.WrappedJsonErrorf(v.pathStr(), "Unsupported action %s", result.Action.String())
		}
	}
	if lastMatchEnd < len(str) {
		buf = append(buf, str[lastMatchEnd:]...)
	}
	s := string(buf)
	sbp.b = buf
	stringBuilderPool.Put(sbp)
	return types.NewString(s), nil
}

func (v *TemplateVisitor) visitAssert(value interface{}) error {
	if arr, ok := value.(*types.JsonArray); ok {
		for i, i2 := range arr.Value {
			n := v.pathPushIndex(i)
			err := v.visitAssert(i2)
			v.pathPop(n)
			if err != nil {
				return burrito.PassError(err)
			}
		}
	} else if str, ok := value.(*types.JsonString); ok {
		result, err := evalWithExtraScope(str.StringValue(), &v.scope, &v.moduleScope)
		if err != nil {
			return utils.WrapJsonErrorf(v.pathStr(), err, "Error evaluating '%s'", str.StringValue())
		}
		if result.Action != types.Value {
			return utils.WrappedJsonErrorf(v.pathStr(), "Unsupported action %s", result.Action.String())
		}
		if !result.Value.BoolValue() {
			return utils.WrappedJsonErrorf(v.pathStr(), "Assertion failed for '%s'", str.StringValue())
		}
	} else if obj, ok := value.(*types.JsonObject); ok {
		condition, err := FindAnyCase[*types.JsonString](obj, "condition")
		if err != nil {
			return utils.WrapJsonErrorf(v.pathStr(), err, "Invalid condition")
		}
		result, err := evalWithExtraScope((*condition).StringValue(), &v.scope, &v.moduleScope)
		if err != nil {
			return utils.WrapJsonErrorf(v.pathStr(), err, "Error evaluating '%s'", (*condition).StringValue())
		}
		if result.Action != types.Value {
			return utils.WrappedJsonErrorf(v.pathStr(), "Unsupported action %s", result.Action.String())
		}
		if !result.Value.BoolValue() {
			message, err := FindAnyCase[*types.JsonString](obj, "message")
			if err != nil {
				if burrito.HasTag(err, WrongTypeErrTag) {
					return utils.WrapJsonErrorf(v.pathStr(), err, "Invalid assertion message")
				}
				return utils.WrapJsonErrorf(v.pathStr(), err, "Assertion failed for '%s'", str.StringValue())
			}
			msg, err := v.visitString((*message).StringValue())
			if err != nil {
				return utils.WrappedJsonErrorf(v.pathStr(), "Error evaluating message '%s'", (*message).StringValue())
			}
			return utils.WrappedJsonErrorf(v.pathStr(), types.ToPrettyString(msg))
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
