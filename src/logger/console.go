package logger

import (
	"fmt"
	"time"
)

// ConsoleLogger ...
type ConsoleLogger struct {
	Level Level
}

// NewConsoleLogger ...
func NewConsoleLogger(levelStr string) *ConsoleLogger {
	level := strToLevel(levelStr)
	return &ConsoleLogger{
		Level: level,
	}
}

func (c *ConsoleLogger) logOrNot(level Level) bool {
	return level >= c.Level
}

// 输出操作
func (c *ConsoleLogger) log(level Level, format string, a ...interface{}) {
	// 验证等级 符合的输出
	if c.logOrNot(level) {
		msg := fmt.Sprintf(format, a...)
		levelStr := levelToStr(level)
		funcname, filename, line := getFileInfo(3)
		timeStr := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] [%s|%s|%d] [%s] %s\n", timeStr, funcname, filename, line, levelStr, msg)
	}
}

// Debug ...
func (c *ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

// Info ...
func (c *ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}

// Warning ...
func (c *ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}

// Error ...
func (c *ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}

// Fatal ...
func (c *ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
