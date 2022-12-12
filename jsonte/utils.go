package jsonte

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

const WrongTypeErrTag = "wrong_type"

// FindAnyCase returns the first key that matches the given key, ignoring case.
func FindAnyCase[T types.JsonType](o types.JsonObject, key ...string) (*T, error) {
	if key == nil || len(key) == 0 {
		return nil, burrito.WrappedError("The key is nil or empty")
	}
	if len(key) > 1 {
		snakeCase := strings.Join(key, "_")
		for _, k := range o.Keys() {
			if strings.EqualFold(k, snakeCase) {
				t, ok := o.Get(k).(T)
				if !ok {
					err := burrito.WrappedErrorf("The value of %s is not a %s", k, types.TypeName(*new(T)))
					burrito.AsBurritoError(err).AddTag(WrongTypeErrTag)
					return nil, err
				}
				return &t, nil
			}
		}
	}
	camelCase := toCamelCase(key...)
	for _, k := range o.Keys() {
		if strings.EqualFold(k, camelCase) {
			t, ok := o.Get(k).(T)
			if !ok {
				err := burrito.WrappedErrorf("The value of %s is not a %s", k, types.TypeName(*new(T)))
				burrito.AsBurritoError(err).AddTag(WrongTypeErrTag)
				return nil, err
			}
			return &t, nil
		}
	}
	return nil, burrito.WrappedErrorf("The key %s is not found", key)
}

func toCamelCase(s ...string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s[0]
	}
	var sb strings.Builder
	for _, v := range s {
		sb.WriteString(cases.Title(language.Und).String(v))
	}
	return sb.String()
}
