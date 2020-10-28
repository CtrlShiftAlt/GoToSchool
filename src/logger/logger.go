package logger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

// Level type
type Level uint16

const (
	// DEBUG level
	DEBUG Level = iota
	// INFO level
	INFO
	// WARNING level
	WARNING
	// ERROR level
	ERROR
	// FATAL level
	FATAL
)

func strToLevel(levelStr string) (level Level) {
	levelStr = strings.ToUpper(levelStr)
	switch levelStr {
	case "DEBUG":
		level = DEBUG
	case "INFO":
		level = INFO
	case "WARNING":
		level = WARNING
	case "ERROR":
		level = ERROR
	case "FATAL":
		level = FATAL
	default:
		level = DEBUG
	}
	return
}

func levelToStr(level Level) (levelStr string) {
	switch level {
	case DEBUG:
		levelStr = "DEBUG"
	case INFO:
		levelStr = "INFO"
	case WARNING:
		levelStr = "WARNING"
	case ERROR:
		levelStr = "ERROR"
	case FATAL:
		levelStr = "FATAL"
	default:
		levelStr = "DEBUG"
	}
	return
}

func getFileInfo(skip int) (funcname, filename string, line int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("get fileInfo fail, err\n")
	}
	funcname = runtime.FuncForPC(pc).Name()
	filename = path.Base(file)
	return
}

// Logger ...
type Logger interface {
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}

// NewLogger ...
func NewLogger(t string, levelStr string, a ...interface{}) Logger {
	if t == "c" {
		return NewConsoleLogger(levelStr)
	} else if t == "f" {
		var fieldPath = "./"
		var fieldName = "log.log"
		for index, arg := range a {
			if value, ok := arg.(string); ok && index == 0 {
				fieldPath = value
			}
			if value, ok := arg.(string); ok && index == 1 {
				fieldName = value
			}
		}
		FileLogger := NewFileLogger(levelStr, fieldPath, fieldName)
		return FileLogger
	}
	return NewConsoleLogger(levelStr)
}
