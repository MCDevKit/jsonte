package jsonte

import (
	"fmt"
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/gammazero/deque"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// JsonModule represents a module that can be extended by a template
type JsonModule struct {
	Name     string
	Scope    utils.NavigableMap[string, interface{}]
	Template utils.NavigableMap[string, interface{}]
	Copy     string
}

type TemplateVisitor struct {
	scope       deque.Deque[interface{}]
	globalScope utils.NavigableMap[string, interface{}]
	deadline    int64
}

const MaxInt64 = int64(^uint64(0) >> 1)

// LoadModule loads a module from a file and returns a JsonModule
func LoadModule(input string) (JsonModule, error) {
	json, err := utils.ParseJsonObject([]byte(input))
	if err != nil {
		return JsonModule{}, utils.WrapErrorf(err, "Failed to parse JSON module")
	}
	moduleName, ok := json.Get("$module").(string)
	if !ok {
		return JsonModule{}, utils.WrappedErrorf("$module", "The module does not have a name")
	}
	scope, ok := json.Get("$scope").(utils.NavigableMap[string, interface{}])
	if !ok {
		scope = utils.NewNavigableMap[string, interface{}]()
	}
	template, ok := json.Get("$template").(utils.NavigableMap[string, interface{}])
	if !ok {
		return JsonModule{}, utils.WrappedJsonErrorf("$template", "The module does not have a template")
	}
	c, isCopy := json.Get("$copy").(string)
	if !isCopy {
		c = ""
	}
	return JsonModule{
		Name:     moduleName,
		Scope:    scope,
		Template: template,
		Copy:     c,
	}, nil
}

// Process processes a template and returns a map of the processed templates
func Process(name, input string, globalScope utils.NavigableMap[string, interface{}], modules map[string]JsonModule, timeout int64) (utils.NavigableMap[string, interface{}], error) {
	// Set up the deadline
	deadline := time.Now().UnixMilli() + timeout
	if timeout <= 0 {
		deadline = MaxInt64
	}

	// Parse the input
	result := utils.NewNavigableMap[string, interface{}]()
	root, err := utils.ParseJsonObject([]byte(input))
	if err != nil {
		return utils.NewNavigableMap[string, interface{}](), utils.WrapErrorf(err, "Failed to parse JSON")
	}

	// Define scope
	scope := utils.DeepCopyObject(globalScope)
	if root.ContainsKey("$scope") {
		scope = utils.MergeObject(root.Get("$scope").(utils.NavigableMap[string, interface{}]), scope, false)
	}

	isCopy := root.ContainsKey("$copy")
	var c string
	if isCopy {
		c, isCopy = root.Get("$copy").(string)
	}
	isExtend := root.ContainsKey("$extend")
	hasTemplate := root.ContainsKey("$template")

	// If none of the options are defined, return unmodified JSON
	if !hasTemplate && !isCopy && !isExtend {
		result.Put(name, root)
		return result, nil
	}

	visitor := TemplateVisitor{
		scope:       deque.Deque[interface{}]{},
		globalScope: globalScope,
		deadline:    deadline,
	}
	visitor.pushScope(utils.DeepCopyObject(scope))
	var template utils.NavigableMap[string, interface{}]

	// handle generating multiple files
	if root.ContainsKey("$files") {
		file := root.Get("$files").(utils.NavigableMap[string, interface{}])
		array, err := Eval(file.Get("array").(string), visitor.scope, "$files.array")
		if err != nil {
			return utils.NewNavigableMap[string, interface{}](), utils.WrapErrorf(err, "Failed to evaluate $files.array")
		}
		if array.Value == nil {
			return utils.NewNavigableMap[string, interface{}](), utils.WrappedJsonErrorf("$files.array", "The array evaluated to null")
		}
		if arr, ok := array.Value.([]interface{}); ok {
			for i, item := range arr {
				err := checkDeadline(deadline)
				if err != nil {
					return result, err
				}
				extra := map[string]interface{}{
					"index": utils.ToNumber(i),
					"value": item,
				}
				visitor.pushScope(extra)
				visitor.pushScope(item)
				if isCopy {
					template, err = processCopy(c, visitor, modules, "$files.array", timeout)
					if err != nil {
						return utils.NewNavigableMap[string, interface{}](), utils.PassError(err)
					}
					visitor.pushScope(map[string]interface{}{"$copy": template})
				} else {
					template = root.Get("$template").(utils.NavigableMap[string, interface{}])
				}
				if isExtend {
					template, err = extendTemplate(root.Get("$extend"), template, visitor, modules)
					if err != nil {
						return utils.NewNavigableMap[string, interface{}](), utils.PassError(err)
					}
				}
				if isCopy && hasTemplate {
					if temp, ok := root.Get("$template").(utils.NavigableMap[string, interface{}]); ok {
						template = utils.MergeObject(template, temp, false)
					} else if temp1, ok := root.Get("$template").(utils.NavigableMap[string, interface{}]); ok {
						template = utils.MergeObject(template, temp1, false)
					}
				}
				mFileName, err := visitor.visitString(file.Get("fileName").(string), "$files.fileName")
				if err != nil {
					return utils.NewNavigableMap[string, interface{}](), utils.WrapErrorf(err, "Failed to evaluate $files.fileName")
				}
				f, err := visitor.visitObject(utils.DeepCopyObject(template), "$template")
				if err != nil {
					return utils.NewNavigableMap[string, interface{}](), utils.PassError(err)
				}
				if isCopy {
					visitor.popScope()
				}
				visitor.popScope()
				visitor.popScope()
				result.Put(mFileName.(string), utils.MergeObject(utils.NewNavigableMap[string, interface{}](), f.(utils.NavigableMap[string, interface{}]), false))
				result.Put(mFileName.(string), utils.DeleteNulls(result.Get(mFileName.(string)).(utils.NavigableMap[string, interface{}])))
			}
		} else {
			return utils.NewNavigableMap[string, interface{}](), utils.WrappedJsonErrorf("$files.array", "The array evaluated to a non-array")
		}
	} else {
		if isCopy {
			template, err = processCopy(c, visitor, modules, "$copy", timeout)
			if err != nil {
				return utils.NewNavigableMap[string, interface{}](), utils.PassError(err)
			}
			visitor.pushScope(map[string]interface{}{"$copy": template})
		}
		if isExtend {
			template, err = extendTemplate(root.Get("$extend"), template, visitor, modules)
			if err != nil {
				return utils.NewNavigableMap[string, interface{}](), utils.PassError(err)
			}
		}
		if hasTemplate {
			template = utils.MergeObject(template, root.Get("$template").(utils.NavigableMap[string, interface{}]), false)
		}
		f, err := visitor.visitObject(utils.DeepCopyObject(template), "$template")
		if err != nil {
			return utils.NewNavigableMap[string, interface{}](), utils.PassError(err)
		}
		if isCopy {
			visitor.popScope()
		}
		result.Put(name, utils.MergeObject(utils.NewNavigableMap[string, interface{}](), f.(utils.NavigableMap[string, interface{}]), false))
		result.Put(name, utils.DeleteNulls(result.Get(name).(utils.NavigableMap[string, interface{}])))
	}

	return utils.UnwrapContainers(result).(utils.NavigableMap[string, interface{}]), nil
}

func processCopy(c interface{}, visitor TemplateVisitor, modules map[string]JsonModule, path string, timeout int64) (utils.NavigableMap[string, interface{}], error) {
	c, err := visitor.visitString(c.(string), path)
	if err != nil {
		return utils.NewNavigableMap[string, interface{}](), utils.WrapErrorf(err, "Failed to evaluate $copy")
	}
	if copyPath, ok := c.(string); ok {
		resolve, err := safeio.Resolver.Open(copyPath)
		if err != nil {
			return utils.NewNavigableMap[string, interface{}](), utils.WrapErrorf(err, "Failed to open %s", copyPath)
		}
		all, err := ioutil.ReadAll(resolve)
		if err != nil {
			return utils.NewNavigableMap[string, interface{}](), utils.WrapErrorf(err, "Failed to read %s", copyPath)
		}
		if strings.HasSuffix(copyPath, ".templ") {
			processedMap, err := Process("copy", string(all), visitor.globalScope, modules, timeout)
			if err != nil {
				return utils.NewNavigableMap[string, interface{}](), utils.WrapErrorf(err, "Failed to process copy template %s", copyPath)
			}
			if processedMap.Size() > 1 {
				return utils.NewNavigableMap[string, interface{}](), utils.WrappedJsonErrorf(path, "The copy template must compile to a single object")
			}
			template := processedMap.Get("copy").(utils.NavigableMap[string, interface{}])
			return template, nil
		} else {
			template, err := utils.ParseJsonObject(all)
			if err != nil {
				return utils.NewNavigableMap[string, interface{}](), utils.WrapErrorf(err, "Failed to parse %s", copyPath)
			}
			return template, nil
		}
	} else {
		return utils.NewNavigableMap[string, interface{}](), utils.WrappedJsonErrorf(path, "The copy path evaluated to a non-string")
	}
}

var templatePattern, _ = regexp.Compile("\\{\\{(?:\\\\.|[^{}])+}}")
var actionPattern, _ = regexp.Compile("^\\{\\{(?:\\\\.|[^{}])+}}$")

func extendTemplate(extend interface{}, template utils.NavigableMap[string, interface{}], visitor TemplateVisitor, modules map[string]JsonModule) (utils.NavigableMap[string, interface{}], error) {
	resolvedModules := make([]string, 0)
	toResolve := make([]string, 0)
	isString := true

	if arr, ok := extend.([]interface{}); ok {
		isString = false
		for _, mod := range arr {
			if str, ok := mod.(string); ok {
				toResolve = append(toResolve, str)
			}
		}
	} else if str, ok := extend.(string); ok {
		toResolve = append(toResolve, str)
	}
	for i, str := range toResolve {
		path := "$extend"
		if !isString {
			path += "[" + strconv.Itoa(i) + "]"
		}
		if actionPattern.MatchString(str) {
			eval, err := Eval(str, visitor.scope, path)
			if err != nil {
				return utils.NewNavigableMap[string, interface{}](), utils.WrapJsonErrorf(path, err, "Failed to evaluate %s", path)
			}
			if mods, ok := eval.Value.([]interface{}); ok {
				stringMods := make([]string, len(mods))
				for i, mod := range mods {
					stringMods[i] = utils.ToString(mod)
				}
				resolvedModules = append(resolvedModules, stringMods...)
			} else if strMod, ok := eval.Value.(string); ok {
				resolvedModules = append(resolvedModules, strMod)
			} else {
				return utils.NewNavigableMap[string, interface{}](), utils.WrappedJsonErrorf(path, "The module name evaluated to a non-string")
			}
		} else {
			resolvedModules = append(resolvedModules, str)
		}
	}
	for _, module := range resolvedModules {
		if _, ok := modules[module]; !ok {
			return utils.NewNavigableMap[string, interface{}](), utils.WrappedJsonErrorf("$extend", "The module '%s' does not exist", module)
		}
		mod := modules[module]
		if mod.Template.IsEmpty() {
			return utils.NewNavigableMap[string, interface{}](), utils.WrappedJsonErrorf("$extend", "The module '%s' does not have a template", module)
		}
		visitor.scope.PushFront(mod.Scope)
		if mod.Copy != "" {
			object, err := processCopy(mod.Copy, visitor, modules, "$copy", -1)
			if err != nil {
				return utils.NewNavigableMap[string, interface{}](), utils.WrapErrorf(err, "Error processing $copy for module %s", mod.Name)
			}
			template = utils.MergeObject(object, template, true)
		}
		parent, err := visitor.visitObject(mod.Template, "[Module "+module+"]")
		visitor.scope.PopFront()
		if err != nil {
			return utils.NewNavigableMap[string, interface{}](), utils.WrapErrorf(err, "Error processing template for module %s", mod.Name)
		}
		template = utils.MergeObject(template, parent.(utils.NavigableMap[string, interface{}]), true)
	}
	return template, nil
}

func (v *TemplateVisitor) pushScope(obj interface{}) {
	v.scope.PushBack(obj)
}

func (v *TemplateVisitor) popScope() {
	v.scope.PopBack()
}

func (v *TemplateVisitor) visit(obj interface{}, path string) (interface{}, error) {
	if obj == nil {
		return nil, nil
	}
	if str, ok := obj.(string); ok {
		return v.visitString(str, path)
	}
	if arr, ok := obj.([]interface{}); ok {
		return v.visitArray(arr, path)
	}
	if objMap, ok := obj.(utils.NavigableMap[string, interface{}]); ok {
		return v.visitObject(objMap, path)
	}
	return obj, nil
}

func (v *TemplateVisitor) visitObject(obj utils.NavigableMap[string, interface{}], path string) (interface{}, error) {
	var result = utils.NewNavigableMap[string, interface{}]()
	for _, key := range obj.Keys() {
		err := checkDeadline(v.deadline)
		if err != nil {
			return result, err
		}
		value := obj.Get(key)
		if key == "$comment" {
			continue
		}
		if key == "$assert" {
			if err = v.visitAssert(value, fmt.Sprintf("%s/%s", path, "$assert")); err != nil {
				return result, err
			}
			continue
		}
		if actionPattern.MatchString(key) {
			eval, err := Eval(key, v.scope, fmt.Sprintf("%s/%s", path, key))
			if err != nil {
				return nil, utils.WrapJsonErrorf(path, err, "Failed to evaluate %s", key)
			}
			switch eval.Action {
			case utils.Iteration:
				if _, ok := value.(utils.NavigableMap[string, interface{}]); !ok {
					return nil, utils.WrappedJsonErrorf(path, "The value of the iteration key must be an object")
				}
				if arr, ok := eval.Value.([]interface{}); ok {
					for i := range arr {
						v.pushScope(map[string]interface{}{
							eval.IndexName: utils.ToNumber(i),
							eval.Name:      arr[i],
						})
						v.pushScope(arr[i])
						o, err := v.visit(value, fmt.Sprintf("%s/%s[%d]", path, key, i))
						v.popScope()
						v.popScope()
						if err != nil {
							return nil, utils.PassError(err)
						}
						u := o.(utils.NavigableMap[string, interface{}])
						for _, k := range u.Keys() {
							result.Put(k, u.Get(k))
						}
					}
				} else {
					return nil, utils.WrappedJsonErrorf(path, "Iteration action must evaluate to an array")
				}
			case utils.Predicate:
				if utils.ToBoolean(eval.Value) {
					o, err := v.visit(value, fmt.Sprintf("%s/%s", path, key))
					if err != nil {
						return nil, utils.PassError(err)
					}
					if obj, ok := o.(utils.NavigableMap[string, interface{}]); ok {
						for _, k := range obj.Keys() {
							result.Put(k, obj.Get(k))
						}
					} else {
						return nil, utils.WrappedJsonErrorf(path, "The value of the predicate key must be an object")
					}
				}
			case utils.Value:
				key, err := v.visitString(key, fmt.Sprintf("%s/%s", path, key))
				if err != nil {
					return nil, utils.PassError(err)
				}
				key = utils.ToString(key)
				r, err := v.visit(value, fmt.Sprintf("%s/%s", path, key))
				if err != nil {
					return nil, utils.PassError(err)
				}
				result.Put(key.(string), r)
			default:
				return nil, utils.WrappedJsonErrorf(path, "Unsupported action %s", eval.Action.String())
			}
		} else if templatePattern.MatchString(key) {
			key, err := v.visitString(key, fmt.Sprintf("%s/%s", path, key))
			if err != nil {
				return nil, utils.PassError(err)
			}
			key = utils.ToString(key)
			r, err := v.visit(value, fmt.Sprintf("%s/%s", path, key))
			if err != nil {
				return nil, utils.PassError(err)
			}
			result.Put(key.(string), r)
		} else {
			var err error
			r, err := v.visit(value, fmt.Sprintf("%s/%s", path, key))
			if err != nil {
				return nil, utils.PassError(err)
			}
			result.Put(key, r)
		}
	}
	return result, nil
}

func (v *TemplateVisitor) visitArray(arr []interface{}, path string) (interface{}, error) {
	var result = make([]interface{}, 0)
	for i, value := range arr {
		err := checkDeadline(v.deadline)
		if err != nil {
			return result, err
		}
		a, err := v.visitArrayElement(result, value, fmt.Sprintf("%s[%d]", path, i))
		if err != nil {
			return nil, utils.PassError(err)
		}
		result = a
	}
	return result, nil
}

// Special visitor for cases when the array element is an object, that generates multiple values
func (v *TemplateVisitor) visitArrayElement(array []interface{}, element interface{}, path string) ([]interface{}, error) {
	if obj, ok := element.(utils.NavigableMap[string, interface{}]); ok {
		if obj.Size() == 1 {
			for _, key := range obj.Keys() {
				value := obj.Get(key)
				if actionPattern.MatchString(key) {
					eval, err := Eval(key, v.scope, path)
					if err != nil {
						return array, utils.WrapJsonErrorf(path, err, "Failed to evaluate %s", key)
					}
					switch eval.Action {
					case utils.Iteration:
						if arr, ok := eval.Value.([]interface{}); ok {
							for i := range arr {
								v.pushScope(map[string]interface{}{
									eval.IndexName: utils.ToNumber(i),
									eval.Name:      arr[i],
								})
								v.pushScope(arr[i])
								a, err := v.visitArrayElement(array, value, fmt.Sprintf("%s[%d]", path, i))
								array = a
								v.popScope()
								v.popScope()
								if err != nil {
									return array, utils.PassError(err)
								}
							}
							return array, nil
						} else {
							return nil, utils.WrappedJsonErrorf(path, "Iteration action must evaluate to an array")
						}
					case utils.Predicate:
						if utils.ToBoolean(eval.Value) {
							return v.visitArrayElement(array, value, path)
						}
						return array, nil
					case utils.Value:
						visit, err := v.visit(element, path)
						if err != nil {
							return nil, utils.PassError(err)
						}
						array = append(array, visit)
						return array, nil
					default:
						return nil, utils.WrappedJsonErrorf(path, "Unsupported action %s", eval.Action.String())
					}
				}
			}
		}
	}
	visit, err := v.visit(element, path)
	if err != nil {
		return nil, utils.PassError(err)
	}
	if arr, ok := visit.([]interface{}); ok {
		array = append(array, arr...)
	} else {
		array = append(array, visit)
	}
	return array, nil
}

func (v *TemplateVisitor) visitString(str string, path string) (interface{}, error) {
	err := checkDeadline(v.deadline)
	if err != nil {
		return nil, err
	}
	matches := templatePattern.FindAllString(str, -1)
	replacements := make(map[string]string, len(matches))
	for _, match := range matches {
		result, err := Eval(match, v.scope, path)
		if err != nil {
			return nil, utils.WrapJsonErrorf(path, err, "Error evaluating '%s'", match)
		}
		if result.Value == nil {
			return nil, utils.WrappedJsonErrorf(path, "The expression '%s' evaluated to null", match)
		}
		if _, ok := result.Value.(string); !ok && str == match {
			return result.Value, nil
		}
		if result.Action == utils.Literal {
			return result.Value, nil
		} else if result.Action == utils.Value {
			replacements[match] = utils.ToString(result.Value)
		} else {
			return nil, utils.WrappedJsonErrorf(path, "Unsupported action %s", result.Action.String())
		}
	}
	result := templatePattern.ReplaceAllStringFunc(str, func(match string) string {
		return replacements[match]
	})
	return result, nil
}

func (v *TemplateVisitor) visitAssert(value interface{}, path string) error {
	if arr, ok := value.([]interface{}); ok {
		for i, i2 := range arr {
			err := v.visitAssert(i2, fmt.Sprintf("%s[%d]", path, i))
			if err != nil {
				return utils.PassError(err)
			}
		}
	} else if str, ok := value.(string); ok {
		result, err := Eval(str, v.scope, path)
		if err != nil {
			return utils.WrapJsonErrorf(path, err, "Error evaluating '%s'", str)
		}
		if result.Action != utils.Value {
			return utils.WrappedJsonErrorf(path, "Unsupported action %s", result.Action.String())
		}
		if !utils.ToBoolean(result.Value) {
			return utils.WrappedJsonErrorf(path, "Assertion failed for '%s'", str)
		}
	} else if obj, ok := value.(utils.NavigableMap[string, interface{}]); ok {
		condition, ok := obj.Get("condition").(string)
		if !ok {
			return utils.WrappedJsonErrorf(path, "Condition must be a string")
		}
		result, err := Eval(condition, v.scope, path)
		if err != nil {
			return utils.WrapJsonErrorf(path, err, "Error evaluating '%s'", condition)
		}
		if result.Action != utils.Value {
			return utils.WrappedJsonErrorf(path, "Unsupported action %s", result.Action.String())
		}
		if !utils.ToBoolean(result.Value) {
			message, ok := obj.Get("message").(string)
			if !ok {
				return utils.WrappedJsonErrorf(path, "Assertion failed for '%s'", str)
			}
			msg, err := v.visitString(message, path)
			if err != nil {
				return utils.WrappedJsonErrorf(path, "Error evaluating message '%s'", message)
			}
			return utils.WrappedJsonErrorf(path, utils.ToPrettyString(msg))
		}
	}
	return nil
}

func checkDeadline(deadline int64) error {
	if time.Now().UnixMilli() > deadline {
		return utils.WrappedErrorf("Template evaluation timed out")
	}
	return nil
}
