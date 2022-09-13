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

type JsonModule struct {
	Name     string
	Scope    utils.JsonObject
	Template utils.JsonObject
}

type TemplateVisitor struct {
	scope       deque.Deque[interface{}]
	globalScope utils.JsonObject
	deadline    int64
}

const MaxInt64 = int64(^uint64(0) >> 1)

func LoadModule(input string) (JsonModule, error) {
	json, err := utils.ParseJson([]byte(input))
	if err != nil {
		return JsonModule{}, err
	}
	moduleName, ok := json["$module"].(string)
	if !ok {
		return JsonModule{}, &utils.TemplatingError{
			Path:    "$module",
			Message: "The $module field is missing or not a string",
		}
	}
	scope, ok := json["$scope"].(utils.JsonObject)
	if !ok {
		scope = utils.JsonObject{}
	}
	template, ok := json["$template"].(utils.JsonObject)
	if !ok {
		return JsonModule{}, &utils.TemplatingError{
			Path:    "$template",
			Message: "The $template field is missing or not an object",
		}
	}
	return JsonModule{
		Name:     moduleName,
		Scope:    scope,
		Template: template,
	}, nil
}

func Process(name, input string, globalScope utils.JsonObject, modules map[string]JsonModule, timeout int64) (map[string]interface{}, error) {
	// Set up the deadline
	deadline := time.Now().UnixMilli() + timeout
	if timeout <= 0 {
		deadline = MaxInt64
	}

	// Parse the input
	result := make(utils.JsonObject)
	root, err := utils.ParseJson([]byte(input))
	if err != nil {
		return nil, err
	}

	// Define scope
	scope := utils.DeepCopyObject(globalScope)
	if s, ok := root["$scope"]; ok {
		scope = utils.MergeObject(s.(utils.JsonObject), scope)
	}

	c, isCopy := root["$copy"].(string)
	_, isExtend := root["$extend"]
	_, hasTemplate := root["$template"]

	// If none of the options are defined, return unmodified JSON
	if !hasTemplate && !isCopy && !isExtend {
		result[name] = root
		return result, nil
	}

	visitor := TemplateVisitor{
		scope:       deque.Deque[interface{}]{},
		globalScope: globalScope,
		deadline:    deadline,
	}
	visitor.pushScope(utils.DeepCopyObject(scope))
	var template utils.JsonObject

	if file, ok := root["$files"]; ok {
		//fileName := file.(utils.JsonObject)["fileName"]
		array, err := Eval(file.(utils.JsonObject)["array"].(string), visitor.scope, "$files.array")
		if err != nil {
			return nil, err
		}
		if array.Value == nil {
			return nil, &utils.TemplatingError{
				Path:    name,
				Message: "The array in the $files evaluated to null",
			}
		}
		if arr, ok := array.Value.(utils.JsonArray); ok {
			for i, item := range arr {
				checkDeadline(deadline)
				extra := utils.JsonObject{
					"index": i,
					"value": item,
				}
				visitor.pushScope(extra)
				visitor.pushScope(item)
				if isCopy {
					template, err = processCopy(name, c, visitor, modules, "$files.array", timeout)
					if err != nil {
						return nil, err
					}
				} else {
					template = root["$template"].(utils.JsonObject)
				}
				if isExtend {
					template, err = extendTemplate(root["$extend"], template, visitor, modules)
					if err != nil {
						return nil, err
					}
				}
				if isCopy && hasTemplate {
					if temp, ok := root["$template"].(utils.JsonObject); ok {
						template = utils.MergeObject(template, temp)
					} else if temp1, ok := root["$template"].(map[string]interface{}); ok {
						template = utils.MergeObject(template, temp1)
					}
				}
				utils.DeleteNulls(template)
				mFileName, err := visitor.visitString(file.(utils.JsonObject)["fileName"].(string), "$files.fileName")
				if err != nil {
					return nil, err
				}
				result[mFileName.(string)], err = visitor.visitObject(utils.DeepCopyObject(template), "$files.template")
				if err != nil {
					return nil, err
				}
				visitor.popScope()
				visitor.popScope()
			}
		} else {
			return nil, &utils.TemplatingError{
				Path:    name,
				Message: "The array in the $files evaluated to a non-array",
			}
		}
	} else {
		if isCopy {
			template, err = processCopy(name, c, visitor, modules, "$copy", timeout)
			if err != nil {
				return nil, err
			}
		}
		if isExtend {
			template, err = extendTemplate(root["$extend"], template, visitor, modules)
			if err != nil {
				return nil, err
			}
		}
		if hasTemplate {
			template = utils.MergeObject(template, root["$template"].(utils.JsonObject))
		}
		utils.DeleteNulls(template)
		result[name], err = visitor.visitObject(utils.DeepCopyObject(template), "$template")
		if err != nil {
			return nil, err
		}
	}

	return utils.UnwrapContainers(result).(utils.JsonObject), nil
}

func processCopy(name string, c interface{}, visitor TemplateVisitor, modules map[string]JsonModule, path string, timeout int64) (utils.JsonObject, error) {
	c, err := visitor.visitString(c.(string), path)
	if err != nil {
		return nil, err
	}
	if copyPath, ok := c.(string); ok {
		resolve, err := safeio.Resolver.Open(copyPath)
		if err != nil {
			return nil, err
		}
		all, err := ioutil.ReadAll(resolve)
		if err != nil {
			return nil, err
		}
		if strings.HasSuffix(copyPath, ".templ") {
			processedMap, err := Process("copy", string(all), visitor.globalScope, modules, timeout)
			if err != nil {
				return nil, err
			}
			if len(processedMap) > 1 {
				return nil, &utils.TemplatingError{
					Path:    name,
					Message: "The $copy template returned more than one result",
				}
			}
			template := processedMap["copy"].(utils.JsonObject)
			return template, nil
		} else {
			template, err := utils.ParseJson(all)
			if err != nil {
				return nil, err
			}
			return template, nil
		}
	} else {
		return nil, &utils.TemplatingError{
			Path:    name,
			Message: "The $copy evaluated to a non-string",
		}
	}
}

var templatePattern, _ = regexp.Compile("\\{\\{(?:\\\\.|[^{}])+}}")
var actionPattern, _ = regexp.Compile("^\\{\\{(?:\\\\.|[^{}])+}}$")

func extendTemplate(extend interface{}, template utils.JsonObject, visitor TemplateVisitor, modules map[string]JsonModule) (utils.JsonObject, error) {
	resolvedModules := make([]string, 0)
	toResolve := make([]string, 0)
	isString := true

	if arr, ok := extend.(utils.JsonArray); ok {
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
				return nil, err
			}
			if mods, ok := eval.Value.(utils.JsonArray); ok {
				stringMods := make([]string, len(mods))
				for i, mod := range mods {
					stringMods[i] = utils.ToString(mod)
				}
				resolvedModules = append(resolvedModules, stringMods...)
			} else if strMod, ok := eval.Value.(string); ok {
				resolvedModules = append(resolvedModules, strMod)
			} else {
				return nil, &utils.TemplatingError{
					Path:    path,
					Message: "The $extend action returned a non-string",
				}
			}
		} else {
			resolvedModules = append(resolvedModules, str)
		}
	}
	for _, module := range resolvedModules {
		if _, ok := modules[module]; !ok {
			return nil, &utils.TemplatingError{
				Path:    "$extend",
				Message: "The module " + module + " is not defined",
			}
		}
		mod := modules[module]
		if mod.Template == nil {
			return nil, &utils.TemplatingError{
				Path:    "$extend",
				Message: "The module " + module + " is missing a template",
			}
		}
		visitor.scope.PushFront(mod.Scope)
		parent, err := visitor.visitObject(mod.Template, "[Module "+module+"]")
		visitor.scope.PopFront()
		if err != nil {
			return nil, err
		}
		template = utils.MergeObject(template, parent.(utils.JsonObject))
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
	if arr, ok := obj.(utils.JsonArray); ok {
		return v.visitArray(arr, path)
	}
	if objMap, ok := obj.(utils.JsonObject); ok {
		return v.visitObject(objMap, path)
	}
	return obj, nil
}

func (v *TemplateVisitor) visitObject(obj utils.JsonObject, path string) (interface{}, error) {
	var result = make(utils.JsonObject)
	for key, value := range obj {
		if key == "$comment" {
			continue
		}
		if actionPattern.MatchString(key) {
			eval, err := Eval(key, v.scope, fmt.Sprintf("%s/%s", path, key))
			if err != nil {
				return nil, err
			}
			switch eval.Action {
			case utils.Iteration:
				if _, ok := value.(utils.JsonObject); !ok {
					return nil, &utils.TemplatingError{
						Path:    path,
						Message: "The value of an iteration action must be an object",
					}
				}
				if arr, ok := eval.Value.(utils.JsonArray); ok {
					for i := range arr {
						v.pushScope(utils.JsonObject{
							"index":   i,
							eval.Name: arr[i],
						})
						v.pushScope(arr[i])
						o, err := v.visit(value, fmt.Sprintf("%s/%s[%d]", path, key, i))
						v.popScope()
						v.popScope()
						if err != nil {
							return nil, err
						}
						for k, v := range o.(utils.JsonObject) {
							result[k] = v
						}
					}
				} else {
					return nil, &utils.TemplatingError{
						Path:    path,
						Message: "The iteration action returned a non-array",
					}
				}
			case utils.Predicate:
				if utils.ToBoolean(eval.Value) {
					o, err := v.visit(value, fmt.Sprintf("%s/%s", path, key))
					if err != nil {
						return nil, err
					}
					if obj, ok := o.(utils.JsonObject); ok {
						for k, v := range obj {
							result[k] = v
						}
					} else {
						return nil, &utils.TemplatingError{
							Path:    path,
							Message: "The predicate action requires an object value",
						}
					}
				}
			case utils.Value:
				key, err := v.visitString(key, fmt.Sprintf("%s/%s", path, key))
				if err != nil {
					return nil, err
				}
				key = utils.ToString(key)
				result[key.(string)], err = v.visit(value, fmt.Sprintf("%s/%s", path, key))
				if err != nil {
					return nil, err
				}
			default:
				return nil, &utils.TemplatingError{
					Path:    path,
					Message: "Unsupported action",
				}
			}
		} else if templatePattern.MatchString(key) {
			key, err := v.visitString(key, fmt.Sprintf("%s/%s", path, key))
			if err != nil {
				return nil, err
			}
			key = utils.ToString(key)
			result[key.(string)], err = v.visit(value, fmt.Sprintf("%s/%s", path, key))
			if err != nil {
				return nil, err
			}
		} else {
			var err error
			result[key], err = v.visit(value, fmt.Sprintf("%s/%s", path, key))
			if err != nil {
				return nil, err
			}
		}
	}
	return result, nil
}

func (v *TemplateVisitor) visitArray(arr utils.JsonArray, path string) (interface{}, error) {
	var result = make(utils.JsonArray, 0)
	for i, value := range arr {
		a, err := v.visitArrayElement(result, value, fmt.Sprintf("%s[%d]", path, i))
		if err != nil {
			return nil, err
		}
		result = a
	}
	return result, nil
}

// Special visitor for cases when the array element is an object, that generates multiple values
func (v *TemplateVisitor) visitArrayElement(array utils.JsonArray, element interface{}, path string) (utils.JsonArray, error) {
	if obj, ok := element.(utils.JsonObject); ok {
		if len(obj) == 1 {
			for key, value := range obj {
				if actionPattern.MatchString(key) {
					eval, err := Eval(key, v.scope, path)
					if err != nil {
						return array, err
					}
					switch eval.Action {
					case utils.Iteration:
						if arr, ok := eval.Value.(utils.JsonArray); ok {
							for i := range arr {
								v.pushScope(utils.JsonObject{
									"index":   i,
									eval.Name: arr[i],
								})
								v.pushScope(arr[i])
								a, err := v.visitArrayElement(array, value, fmt.Sprintf("%s[%d]", path, i))
								array = a
								v.popScope()
								v.popScope()
								if err != nil {
									return array, err
								}
							}
							return array, nil
						} else {
							return nil, &utils.TemplatingError{
								Path:    path,
								Message: "The $iteration action returned a non-array",
							}
						}
					case utils.Predicate:
						if utils.ToBoolean(eval.Value) {
							return v.visitArrayElement(array, value, path)
						}
						return array, nil
					case utils.Value:
						visit, err := v.visit(element, path)
						if err != nil {
							return nil, err
						}
						array = append(array, visit)
						return array, nil
					default:
						return nil, &utils.TemplatingError{
							Path:    path,
							Message: "Unsupported action",
						}
					}
				}
			}
		}
	}
	visit, err := v.visit(element, path)
	if err != nil {
		return nil, err
	}
	array = append(array, visit)
	return array, nil
}

func (v *TemplateVisitor) visitString(str string, path string) (interface{}, error) {
	matches := templatePattern.FindAllString(str, -1)
	replacements := make(map[string]string, len(matches))
	for _, match := range matches {
		result, err := Eval(match, v.scope, path)
		if err != nil {
			return nil, err
		}
		if result.Value == nil {
			return nil, &utils.TemplatingError{
				Path:    path,
				Message: "The expression evaluated to null",
			}
		}
		if result.Action == utils.Literal {
			return result.Value, nil
		} else if result.Action == utils.Value {
			replacements[match] = utils.ToString(result.Value)
		} else {
			return nil, &utils.TemplatingError{
				Path:    path,
				Message: "Cannot execute action here",
			}
		}
	}
	result := templatePattern.ReplaceAllStringFunc(str, func(match string) string {
		return replacements[match]
	})
	return result, nil
}

func checkDeadline(deadline int64) {
	if time.Now().UnixMilli() > deadline {
		panic("Deadline exceeded")
	}
}
