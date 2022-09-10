package test

import (
	"jsonte/jsonte/functions"
	"testing"
)

func TestMain(m *testing.M) {
	functions.Init()
	m.Run()
}
