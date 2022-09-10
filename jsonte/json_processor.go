package jsonte

import (
	"fmt"
	"github.com/gammazero/deque"
	"io/ioutil"
	"jsonte/jsonte/safeio"
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
				if isCopy {
					template, err = processCopy(name, c, visitor, modules, "$files.array", timeout)
					if err != nil {
						return nil, err
					}
				} else {
					template = root["$template"].(utils.JsonObject)
				}
				if isExtend {
					visitor.pushScope(item)
					template, err = extendTemplate(root["$extend"], template, visitor, modules)
					if err != nil {
						return nil, err
					}
					visitor.popScope()
				}
				if isCopy && hasTemplate {
					template = utils.MergeObject(template, root["$template"].(map[string]interface{}))
				}
				visitor.popScope()
				utils.DeleteNulls(template)
				visitor.pushScope(item)
				mFileName, err := visitor.visitString(file.(map[string]interface{})["fileName"].(string), "$files.fileName")
				if err != nil {
					return nil, err
				}
				result[mFileName.(string)], err = visitor.visitObject(utils.DeepCopyObject(template), "$files.template")
				if err != nil {
					return nil, err
				}
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
		result[name], err = visitor.visitObject(utils.DeepCopyObject(template), name)
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
		resolve, err := safeio.Resolver(copyPath)
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
	// TODO
	var result = make(utils.JsonObject)
	for key, value := range obj {
		if key == "$comment" {
			continue
		}
		if actionPattern.MatchString(key) {
			if _, ok := value.(utils.JsonObject); !ok {
				return nil, &utils.TemplatingError{
					Path:    path,
					Message: "The value of an iteration action must be an object",
				}
			}
			eval, err := Eval(key, v.scope, path)
			if err != nil {
				return nil, err
			}
			switch eval.Action {
			case utils.Iteration:
				if arr, ok := eval.Value.(utils.JsonArray); ok {
					for i := range arr {
						v.pushScope(utils.JsonObject{
							"index":   i,
							eval.Name: arr[i],
						})
						o, err := v.visit(value, fmt.Sprintf("%s[%d]", path, i))
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
						Message: "The $iteration action returned a non-array",
					}
				}
			}
		} else if templatePattern.MatchString(key) {
			key, err := v.visitString(key, path)
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
	// TODO
	return arr, nil
}

// Special visitor for cases when the array element is an object, that generates multiple values
func (v *TemplateVisitor) visitArrayElement(arr utils.JsonArray, path string) (interface{}, error) {
	// TODO
	return arr, nil
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
