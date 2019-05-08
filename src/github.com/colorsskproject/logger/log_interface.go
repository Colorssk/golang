package logger

type LogInterface interface {
	// 设置日志级别
	SetLevel(level int)
	// 打印debug
	Debug(format string, args ...interface{}) //  拓展符号...接受可变参数
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}
