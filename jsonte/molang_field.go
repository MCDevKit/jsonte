package jsonte

import (
	"strings"

	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	molang "github.com/stirante/molang-go"
	"github.com/stirante/molang-go/printer"
)

// molangFieldValidation reports whether invalid Molang in a field marked with types.MolangFieldPrefix
// is an error. When off, such values are minified token by token, like in .molang files.
var molangFieldValidation = true

// SetMolangValidation sets, whether invalid Molang in a field marked with types.MolangFieldPrefix is an error.
func SetMolangValidation(validate bool) {
	molangFieldValidation = validate
}

// minifyMolangValue validates and minifies every string in value, descending into objects and arrays.
// Values of other types are left alone, because Molang fields accept numbers and booleans as well.
func (v *TemplateVisitor) minifyMolangValue(value types.JsonType) (types.JsonType, error) {
	switch typed := value.(type) {
	case *types.JsonString:
		minified, err := v.minifyMolangString(typed.StringValue())
		if err != nil {
			return nil, err
		}
		return types.NewString(minified), nil
	case *types.JsonObject:
		var outerErr error
		typed.RangeEntries(func(key string, child types.JsonType) bool {
			n := v.pathPush(key)
			minified, err := v.minifyMolangValue(child)
			v.pathPop(n)
			if err != nil {
				outerErr = err
				return true
			}
			typed.Put(key, minified)
			return false
		})
		return typed, outerErr
	case *types.JsonArray:
		// An array of strings is a single Molang program, split over multiple entries
		if joined, ok := joinMolangArray(typed); ok {
			minified, err := v.minifyMolangString(joined)
			if err != nil {
				return nil, err
			}
			return types.NewString(minified), nil
		}
		for i, child := range typed.Value {
			n := v.pathPushIndex(i)
			minified, err := v.minifyMolangValue(child)
			v.pathPop(n)
			if err != nil {
				return nil, err
			}
			typed.Value[i] = minified
		}
		return typed, nil
	}
	return value, nil
}

// minifyMolangString minifies a single Molang value, reporting invalid Molang when validation is on.
func (v *TemplateVisitor) minifyMolangString(str string) (string, error) {
	if strings.TrimSpace(str) == "" {
		return "", nil
	}
	prog, err := molang.Parse(str)
	if err != nil {
		if !molangFieldValidation {
			return utils.MinifyWithShortAccessors(str), nil
		}
		msg, pos, ok := molangParseError(err)
		if !ok {
			return "", utils.WrapJsonErrorf(v.pathStr(), err, "Invalid Molang")
		}
		line, column, snippet := molangErrorLocation(str, pos)
		return "", utils.WrappedJsonErrorf(v.pathStr(), "Invalid Molang: %s (line %d, column %d: %s)", msg, line, column, snippet)
	}
	return printer.Minify(prog), nil
}

// joinMolangArray joins an array of strings into the single Molang program it represents.
// Entries are joined with a newline, so that a program split in the middle of an expression stays
// valid, without adding semicolons the author did not write. Any other array is not a program.
func joinMolangArray(arr *types.JsonArray) (string, bool) {
	if len(arr.Value) == 0 {
		return "", false
	}
	parts := make([]string, len(arr.Value))
	for i, child := range arr.Value {
		str, ok := child.(*types.JsonString)
		if !ok {
			return "", false
		}
		parts[i] = str.StringValue()
	}
	return strings.Join(parts, "\n"), true
}
