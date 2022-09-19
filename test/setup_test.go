package test

import (
	"github.com/MCDevKit/jsonte/jsonte/functions"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"go.uber.org/zap"
	"testing"
)

func TestMain(m *testing.M) {
	utils.InitLogging(zap.DebugLevel)
	functions.Init()
	m.Run()
}
