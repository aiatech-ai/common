package logger

type Logger interface {
	Info(msg string)
	InfoF(msg string, params ...interface{})
	InfoS(msg string, keysAndValues ...interface{})

	Debug(msg string)
	DebugF(msg string, params ...interface{})
	DebugS(msg string, keysAndValues ...interface{})

	Error(msg string)
	ErrorF(msg string, params ...interface{})
	ErrorS(msg string, keysAndValues ...interface{})

	Fatal(msg string)
	FatalF(msg string, params ...interface{})
	FatalS(msg string, keysAndValues ...interface{})
}

var logger Logger

func Initialize(l Logger) {
	logger = l
}

func Info(msg string) {
	logger.Info(msg)
}
func InfoF(msg string, params ...interface{}) {
	logger.InfoF(msg, params...)
}
func InfoS(msg string, keysAndValues ...interface{}) {
	logger.InfoS(msg, keysAndValues...)
}

func Debug(msg string) {
	logger.Debug(msg)
}
func DebugF(msg string, params ...interface{}) {
	logger.DebugF(msg, params...)
}
func DebugS(msg string, keysAndValues ...interface{}) {
	logger.DebugS(msg, keysAndValues...)
}

func Error(msg string) {
	logger.Error(msg)
}
func ErrorF(msg string, params ...interface{}) {
	logger.ErrorF(msg, params...)
}
func ErrorS(msg string, keysAndValues ...interface{}) {
	logger.ErrorS(msg, keysAndValues...)
}

func Fatal(msg string) {
	logger.Fatal(msg)
}
func FatalF(msg string, params ...interface{}) {
	logger.FatalF(msg, params...)
}
func FatalS(msg string, keysAndValues ...interface{}) {
	logger.FatalS(msg, keysAndValues...)
}
