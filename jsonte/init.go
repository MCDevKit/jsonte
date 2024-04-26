package jsonte

import (
	"github.com/MCDevKit/jsonte/jsonte/functions"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/fatih/color"
	"go.uber.org/zap/zapcore"
)

// LibraryInit initializes jsonte in library mode
// level is the logging level
// safeMode is whether to run in safe mode, which disables unsafe functions
func LibraryInit(level zapcore.Level, safeMode bool) {
	utils.InitLogging(level)
	color.NoColor = true
	types.Init()
	functions.Init()
	functions.SafeMode = safeMode
}

// FetchCache caches the vanilla packs
func FetchCache() error {
	return functions.FetchCache()
}
