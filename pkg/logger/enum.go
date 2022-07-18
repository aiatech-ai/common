package logger

import "strings"

type Env int

const (
	unknow Env = iota
	LOCAL
	DEVELOP
	PRODUCTION
)

func ParseEnv(env string) Env {
	envLower := strings.ToLower(env)
	switch envLower {
	case "local":
		return LOCAL
	case "dev":
		return DEVELOP
	case "develop":
		return DEVELOP
	case "prod":
		return PRODUCTION
	case "production":
		return PRODUCTION
	default:
		return unknow
	}
}

type Level int

const (
	UNKNOW Level = iota
	LEVEL_DEBUG
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERROR
	LEVEL_FATAL
)

func ParseLevel(level string) Level {
	levelLower := strings.ToLower(level)
	switch levelLower {
	case "debug":
		return LEVEL_DEBUG
	case "info":
		return LEVEL_INFO
	case "warn":
		return LEVEL_WARN
	case "error":
		return LEVEL_ERROR
	case "fatal":
		return LEVEL_FATAL
	default:
		return UNKNOW
	}
}

func (s Level) String() string {
	return [...]string{
		"unknow",
		"debug",
		"info",
		"warn",
		"error",
		"fatal",
	}[s]
}
