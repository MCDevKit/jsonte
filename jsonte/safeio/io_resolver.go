package safeio

import (
	"errors"
	"io"
	"os"
)

type IOResolver func(path string) (io.ReadCloser, error)

var Resolver IOResolver = DefaultIOResolver

var DefaultIOResolver = func(path string) (io.ReadCloser, error) {
	return os.Open(path)
}

var NoIOResolver = func(path string) (io.ReadCloser, error) {
	return nil, errors.New("file loading has been disabled")
}
