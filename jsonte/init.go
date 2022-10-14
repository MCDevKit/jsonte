package jsonte

import (
	"github.com/MCDevKit/jsonte/jsonte/functions"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/fatih/color"
	"go.uber.org/zap/zapcore"
)

// LibraryInit initializes jsonte in library mode
func LibraryInit(level zapcore.Level) {
	utils.InitLogging(level)
	color.NoColor = true
	functions.Init()
}

// FetchCache caches the vanilla packs
func FetchCache() error {
	return functions.FetchCache()
}
