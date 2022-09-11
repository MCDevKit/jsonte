package test

import (
	"github.com/MCDevKit/jsonte/jsonte/functions"
	"testing"
)

func TestMain(m *testing.M) {
	functions.Init()
	m.Run()
}
