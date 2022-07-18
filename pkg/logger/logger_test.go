package logger_test

import (
	"fmt"
	"testing"

	"github/davidtrse/shinycommon/pkg/logger"
)

func TestRunDefaultLogger(t *testing.T) {
	log, err := logger.NewZapLoggerLevel(logger.DEVELOP, logger.LEVEL_DEBUG)
	if err != nil {
		fmt.Printf("NewZapLogger error: %s\n", err.Error())
		return
	}
	logger.Initialize(log)

	// logger
	log.InfoF("new log with message: %s", "Minh")

	// Logger with key value
	obj := make(map[string]interface{}, 0)
	obj["name"] = "Minh"
	obj["age"] = 24
	logger.DebugS("DebugS test....", "name", "Hoang", "age", 12, "obj", obj)
	logger.ErrorS("ErrorS test....", "name", "Hoang", "age", 12, "obj", obj)
	logger.ErrorF("ErrorF test....", "name", "Hoang", "age", 12, "obj", obj)
}
