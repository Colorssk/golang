package logger

import (
	"testing" // 单元测试
)

func TestFileLogger(t *testing.T) {
	logger := NewFileLogger(LogLevelDebug, "c:/logs/", "test")
	logger.Debug("%d 测试 debug logger", 123)
	logger.Warn("%d 测试 warn logger")
	logger.Close() //continue------
}
