package utils

import (
	"fmt"
	"github.com/Bedrock-OSS/go-burrito/burrito"
)

// WrappedJsonErrorf creates an error with a stack trace from formatted text and appends the path.
func WrappedJsonErrorf(path string, text string, args ...interface{}) error {
	text = fmt.Sprintf(text, args...)
	return burrito.WrappedErrorf("%s at %s", text, path)
}

// WrapJsonErrorf wraps an error with a stack trace and adds additional formatted
// text information.
func WrapJsonErrorf(path string, err error, text string, args ...interface{}) error {
	text = fmt.Sprintf(text, args...)
	return burrito.WrapErrorf(err, "%s at %s", text, path)
}

// BadDeveloperError panics with a message that indicates a bug in the code.
// It should only be used in places that should never be reached or can only be reached
// if there is a bug in the code.
func BadDeveloperError(msg string) {
	panic(fmt.Sprintf("%s. Please report this issue at https://github.com/MCDevKit/jsonte/issues/new", msg))
}
