package jsonte

import (
	"encoding/json"
	"github.com/gammazero/deque"
	"io/ioutil"
	"jsonte/jsonte/io"
	"jsonte/jsonte/utils"
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
	fullScope    utils.JsonObject
	currentScope deque.Deque[interface{}]
	deadline     int64
}

const MaxInt64 = int64(^uint64(0) >> 1)

func Process(name, input string, globalScope utils.JsonObject, modules map[string]JsonModule, timeout int64) (map[string]interface{}, error) {
	// Set up the deadline
	deadline := time.Now().UnixMilli() + timeout
	if timeout <= 0 {
		deadline = MaxInt64
	}

	// Parse the input
	result := make(utils.JsonObject)
	var root utils.JsonObject
	err := json.Unmarshal([]byte(input), &root)
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
		result[name] = input
		return result, nil
	}

	visitor := TemplateVisitor{
		fullScope:    utils.DeepCopyObject(scope),
		currentScope: deque.Deque[interface{}]{},
		deadline:     deadline,
	}
	var template utils.JsonObject

	if file, ok := root["$files"]; ok {
		//fileName := file.(utils.JsonObject)["fileName"]
		visitor.currentScope.PushBack(scope)
		array, err := Eval(file.(utils.JsonObject)["array"].(string), utils.JsonObject{}, visitor.fullScope, visitor.currentScope, "$files.array")
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
				if isCopy {
					c, err := visitor.visitString(c, extra, visitor.fullScope, "$files.array")
					if err != nil {
						return nil, err
					}
					if copyPath, ok := c.(string); ok {
						resolve, err := io.Resolver.Resolve(copyPath)
						if err != nil {
							return nil, err
						}
						all, err := ioutil.ReadAll(resolve)
						if err != nil {
							return nil, err
						}
						if strings.HasSuffix(copyPath, ".templ") {
							processedMap, err := Process("copy", string(all), scope, modules, timeout)
							if err != nil {
								return nil, err
							}
							if len(processedMap) > 1 {
								return nil, &utils.TemplatingError{
									Path:    name,
									Message: "The $copy template returned more than one result",
								}
							}
							template = processedMap["copy"].(utils.JsonObject)
						} else {
							err := json.Unmarshal(all, &template)
							if err != nil {
								return nil, err
							}
						}
					} else {
						return nil, &utils.TemplatingError{
							Path:    name,
							Message: "The $copy evaluated to a non-string",
						}
					}
				} else {
					template = root["$template"].(utils.JsonObject)
				}
				if isExtend {
					visitor.currentScope.PushBack(item)
					template, err = extendTemplate(root["$extend"], template, visitor, extra, modules)
					if err != nil {
						return nil, err
					}
					visitor.currentScope.PopBack()
				}
				// Ended at JsonProcessor:255
			}
		} else {
			return nil, &utils.TemplatingError{
				Path:    name,
				Message: "The array in the $files evaluated to a non-array",
			}
		}

		visitor.currentScope.PopBack()
	}

	return map[string]interface{}{}, nil
}

var templatePattern, _ = regexp.Compile("\\{\\{(?:\\\\.|[^{}])+}}")
var actionPattern, _ = regexp.Compile("^\\{\\{(?:\\\\.|[^{}])+}}$")

func extendTemplate(extend interface{}, template utils.JsonObject, visitor TemplateVisitor, extraScope utils.JsonObject, modules map[string]JsonModule) (utils.JsonObject, error) {
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
			eval, err := Eval(str, extraScope, visitor.fullScope, visitor.currentScope, path)
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
		moduleScope := utils.MergeObject(visitor.fullScope, mod.Scope)
		parent, err := visitor.visitObject(mod.Template, extraScope, moduleScope, "[Module "+module+"]")
		if err != nil {
			return nil, err
		}
		template = utils.MergeObject(template, parent.(utils.JsonObject))
	}
	return template, nil
}

func (v *TemplateVisitor) visitObject(obj utils.JsonObject, extraScope, fullScope utils.JsonObject, path string) (interface{}, error) {
	// TODO
	return obj, nil
}

func (v *TemplateVisitor) visitArray(arr utils.JsonArray, extraScope, fullScope utils.JsonObject, path string) (interface{}, error) {
	// TODO
	return arr, nil
}

// Special visitor for cases when the array element is an object, that generates multiple values
func (v *TemplateVisitor) visitArrayElement(arr utils.JsonArray, extraScope, fullScope utils.JsonObject, path string) (interface{}, error) {
	// TODO
	return arr, nil
}

func (v *TemplateVisitor) visitString(str string, extraScope, fullScope utils.JsonObject, path string) (interface{}, error) {
	matches := templatePattern.FindAllString(str, -1)
	replacements := make(map[string]string, len(matches))
	for _, match := range matches {
		result, err := Eval(match, extraScope, fullScope, v.currentScope, path)
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
