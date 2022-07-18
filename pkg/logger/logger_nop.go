package logger

func init() {
	Initialize(&NopLogger{})
}

type NopLogger struct {
}

func (s NopLogger) Info(msg string) {
}

func (s NopLogger) InfoF(msg string, params ...interface{}) {
}

func (s NopLogger) InfoS(msg string, keysAndValues ...interface{}) {
}

func (s NopLogger) Debug(msg string) {
}

func (s NopLogger) DebugF(msg string, params ...interface{}) {
}

func (s NopLogger) DebugS(msg string, keysAndValues ...interface{}) {
}

func (s NopLogger) Error(msg string) {
}

func (s NopLogger) ErrorF(msg string, params ...interface{}) {
}

func (s NopLogger) ErrorS(msg string, keysAndValues ...interface{}) {
}

func (s NopLogger) Fatal(msg string) {
}

func (s NopLogger) FatalF(msg string, params ...interface{}) {
}

func (s NopLogger) FatalS(msg string, keysAndValues ...interface{}) {
}
