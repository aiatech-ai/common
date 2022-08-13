package logger_test

import (
	"fmt"
	"testing"

	"github.com/shinystarvn/common/pkg/logger"
)

type School struct {
	*Student
}
type Student struct {
	Name string
	Age  int
}

func TestRunDefaultLogger(t *testing.T) {
	// logger
	// log, err := logger.NewZapLogger()
	log, err := logger.NewZapLoggerLevel(logger.DEVELOP, logger.LEVEL_DEBUG)
	if err != nil {
		fmt.Printf("NewZapLogger error: %s\n", err.Error())
		return
	}
	logger.Initialize(log)

	// SU DUNG LOG
	log.InfoF("new log with message: %s", "Minh")
	WriteToLog(log, "Toi thu")
	// Logger with key value
	obj := make(map[string]interface{}, 0)
	obj["name"] = "Minh"
	obj["age"] = 24

	logger.InfoF("DebugS test: value=%s", obj)
	logger.ErrorS("ErrorS test....", "35", "Hoang", "age", 12, "obj", obj)
	logger.ErrorF("ErrorF test....", "name", "Hoang", "age", 12, "obj", obj)
}

func WriteToLog(l logger.Logger, s string) {
	l.Debug(s)
	logger.Debug(s)
}
