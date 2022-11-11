package test

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
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
	burrito.Debug = false
	types.Init()
	functions.Init()
	m.Run()
}
