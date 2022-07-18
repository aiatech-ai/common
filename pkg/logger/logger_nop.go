package logger

import "fmt"

// func init() {
// 	Initialize(&NopLogger{})
// }

type NopLogger struct {
}

func (s NopLogger) Info(msg string) {
	fmt.Println("I'm Info")
}

func (s NopLogger) InfoF(msg string, params ...interface{}) {
	fmt.Println("I'm InfoF")
}

func (s NopLogger) InfoS(msg string, keysAndValues ...interface{}) {
	fmt.Println("I'm InfoS")
}

func (s NopLogger) Debug(msg string) {
	fmt.Println("I'm Debug")
}

func (s NopLogger) DebugF(msg string, params ...interface{}) {
	fmt.Println("I'm DebugF")
}

func (s NopLogger) DebugS(msg string, keysAndValues ...interface{}) {
	fmt.Println("I'm DebugS")
}

func (s NopLogger) Error(msg string) {
	fmt.Println("I'm Error")
}

func (s NopLogger) ErrorF(msg string, params ...interface{}) {
	fmt.Println("I'm ErrorF")
}

func (s NopLogger) ErrorS(msg string, keysAndValues ...interface{}) {
	fmt.Println("I'm ErrorS")
}

func (s NopLogger) Fatal(msg string) {
	fmt.Println("I'm Fatal")
}

func (s NopLogger) FatalF(msg string, params ...interface{}) {
	fmt.Println("I'm FatalF")
}

func (s NopLogger) FatalS(msg string, keysAndValues ...interface{}) {
	fmt.Println("I'm FatalS")
}
