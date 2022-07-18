package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type ZapLogger struct {
	zaplogger     *zap.Logger
	sugaredLogger *zap.SugaredLogger
	NopLogger
}

func NewZapLogger() (Logger, error) {
	zapConfig := zap.NewDevelopmentConfig()
	zapLogger, err := zapConfig.Build(zap.AddCallerSkip(2))
	if err != nil {
		return nil, err
	}
	return &ZapLogger{
		zaplogger:     zapLogger,
		sugaredLogger: zapLogger.Sugar(),
	}, nil
}

func NewZapLoggerLevel(env Env, level Level) (Logger, error) {
	// Create Development or Production Zap
	var zapConfig zap.Config
	if env == DEVELOP {
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		zapConfig.DisableStacktrace = true
		// zapConfig.OutputPaths = []string{"sdtout"}
	} else {
		zapConfig = zap.NewProductionConfig()
	}

	// Set level logger
	var lv zapcore.Level
	switch level {
	case LEVEL_DEBUG:
		lv = zapcore.DebugLevel
	case LEVEL_INFO:
		lv = zapcore.InfoLevel
	case LEVEL_WARN:
		lv = zapcore.WarnLevel
	case LEVEL_ERROR:
		lv = zapcore.ErrorLevel
	case LEVEL_FATAL:
		lv = zapcore.FatalLevel
	default:
		return nil, fmt.Errorf("invalid log level %s", level.String())
	}
	zapConfig.Level = zap.NewAtomicLevelAt(lv)

	// Get line and file
	zapLogger, err := zapConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}
	zap.ReplaceGlobals(zapLogger)

	return &ZapLogger{
		zaplogger:     zapLogger,
		sugaredLogger: zapLogger.Sugar(),
	}, nil
}

func NewZapLoggerLevelFile(env Env, level Level) (Logger, error) {
	// Create Development or Production Zap
	var zapConfig zapcore.EncoderConfig
	if env == DEVELOP {
		zapConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		zapConfig = zap.NewProductionEncoderConfig()
	}
	zapConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	zapConfig.EncodeTime = syslogTimeEncoder
	zapConfig.StacktraceKey = ""
	jsonEncoder := zapcore.NewJSONEncoder(zapConfig)

	// Set level logger
	var lv zapcore.Level
	switch level {
	case LEVEL_DEBUG:
		lv = zapcore.DebugLevel
	case LEVEL_INFO:
		lv = zapcore.InfoLevel
	case LEVEL_WARN:
		lv = zapcore.WarnLevel
	case LEVEL_ERROR:
		lv = zapcore.ErrorLevel
	case LEVEL_FATAL:
		lv = zapcore.FatalLevel
	default:
		return nil, fmt.Errorf("invalid log level %s", level.String())
	}

	// Setup write to file
	writerSyncer := getLogWriter()
	core := zapcore.NewCore(jsonEncoder, writerSyncer, lv)

	// Get line and file
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	zap.ReplaceGlobals(zapLogger)

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

func getLogWriter() zapcore.WriteSyncer {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/app.log",
		MaxSize:    10, // megabytes
		MaxBackups: 5,
		MaxAge:     20, // days
	})
	return zapcore.AddSync(w)
}

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
