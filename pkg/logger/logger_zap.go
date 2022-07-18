package logger

import (
	"fmt"

	"go.uber.org/zap"
)

type ZapLogger struct {
	zaplogger     *zap.Logger
	sugaredLogger *zap.SugaredLogger
	NopLogger
}

func NewZapLogger() (Logger, error) {
	zapConfig := zap.NewProductionConfig()
	zapConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	zapLogger, err := zapConfig.Build(zap.AddCallerSkip(2))
	if err != nil {
		return nil, err
	}
	return &ZapLogger{
		zaplogger:     zapLogger,
		sugaredLogger: zapLogger.Sugar(),
	}, nil
}

func (s ZapLogger) Info(msg string) {
	s.zaplogger.Info(msg)
}

func (s ZapLogger) InfoF(msg string, params ...interface{}) {
	s.zaplogger.Info(fmt.Sprintf(msg, params...))
}

func (s ZapLogger) InfoS(msg string, keysAndValues ...interface{}) {
	s.sugaredLogger.Infow(msg, keysAndValues...)
}

func (s ZapLogger) Debug(msg string) {
	s.zaplogger.Debug(msg)
}

func (s ZapLogger) DebugF(msg string, params ...interface{}) {
	s.zaplogger.Debug(fmt.Sprintf(msg, params...))
}

// func (s ZapLogger) DebugS(msg string, keysAndValues ...interface{}) {
// 	s.sugaredLogger.Debugw(msg, keysAndValues...)
// }

func (s ZapLogger) Error(msg string) {
	s.zaplogger.Error(msg)
}

func (s ZapLogger) ErrorF(msg string, params ...interface{}) {
	s.zaplogger.Error(fmt.Sprintf(msg, params...))
}

func (s ZapLogger) ErrorS(msg string, keysAndValues ...interface{}) {
	s.sugaredLogger.Errorw(msg, keysAndValues...)
}

func (s ZapLogger) Fatal(msg string) {
	s.zaplogger.Fatal(msg)
}

func (s ZapLogger) FatalF(msg string, params ...interface{}) {
	s.zaplogger.Fatal(fmt.Sprintf(msg, params...))
}

func (s ZapLogger) FatalS(msg string, keysAndValues ...interface{}) {
	s.sugaredLogger.Fatalw(msg, keysAndValues...)
}
