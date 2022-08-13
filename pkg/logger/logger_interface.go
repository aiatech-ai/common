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

var mylog Logger

func Initialize(l Logger) {
	mylog = l
}

func Info(msg string) {
	mylog.Info(msg)
}
func InfoF(msg string, params ...interface{}) {
	mylog.InfoF(msg, params...)
}
func InfoS(msg string, keysAndValues ...interface{}) {
	mylog.InfoS(msg, keysAndValues...)
}

func Debug(msg string) {
	mylog.Debug(msg)
}
func DebugF(msg string, params ...interface{}) {
	mylog.DebugF(msg, params...)
}
func DebugS(msg string, keysAndValues ...interface{}) {
	mylog.DebugS(msg, keysAndValues...)
}

func Error(msg string) {
	mylog.Error(msg)
}
func ErrorF(msg string, params ...interface{}) {
	mylog.ErrorF(msg, params...)
}
func ErrorS(msg string, keysAndValues ...interface{}) {
	mylog.ErrorS(msg, keysAndValues...)
}

func Fatal(msg string) {
	mylog.Fatal(msg)
}
func FatalF(msg string, params ...interface{}) {
	mylog.FatalF(msg, params...)
}
func FatalS(msg string, keysAndValues ...interface{}) {
	mylog.FatalS(msg, keysAndValues...)
}
