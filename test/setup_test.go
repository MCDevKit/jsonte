package test

import (
	"github.com/MCDevKit/jsonte/jsonte/functions"
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"go.uber.org/zap"
	"testing"
)

var CacheFS = safeio.CreateFakeFS(map[string]interface{}{}, true)

func TestMain(m *testing.M) {
	utils.InitLogging(zap.DebugLevel)
	types.Init()
	functions.Init()
	m.Run()
}
