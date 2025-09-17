package bench

import (
	"os"
	"testing"

	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/functions"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"go.uber.org/zap"
)

func TestMain(m *testing.M) {
	utils.InitLogging(zap.ErrorLevel)
	burrito.PrintStackTrace = false
	functions.SafeMode = true
	types.Init()
	functions.Init()
	os.Exit(m.Run())
}
