package io

import (
	"errors"
	"os"
)

type IOResolver interface {
	Resolve(path string) ([]byte, error)
}

var Resolver = DefaultIOResolver{}

type DefaultIOResolver struct {
	IOResolver
}

func (r *DefaultIOResolver) Resolve(path string) (*os.File, error) {
	return os.Open(path)
}

type NoIOResolver struct {
	IOResolver
}

func (r *NoIOResolver) Resolve(path string) (*os.File, error) {
	return nil, errors.New("file loading has been disabled")
}
