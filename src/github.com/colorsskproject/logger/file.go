package logger

import (
	"fmt"
	"os"
	"time"
)

type FileLogger struct {
	level   int
	logPath string
	logName string
	// 把debug和warn放在一个文件error和fatal放在一个文件
	file     *os.File
	warnFile *os.File
}

func NewFileLogger(level int, logPath, logName string) LogInterface { // 使用的时候是一个接口，所以返回的是一个接口
	logger := &FileLogger{
		level:   level,
		logPath: logPath,
		logName: logName,
	}
	logger.init()
	return logger
}

func (f *FileLogger) init() {

	filename := fmt.Sprintf("%s%s.log", f.logPath, f.logName)
	// init中打开文件
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed,err:%v", filename, err)) // panic相当于return了
	}
	f.file = file
	//写错误日志和fatal日志的文件
	filename = fmt.Sprintf("%s%s.log.wf", f.logPath, f.logName)
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed,err:%v", filename, err)) // panic相当于return了
	}
	f.warnFile = file
}

// 然后实现这个接口

func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	now := time.Now()
	nowstr := now.Format("2006-01-02 15:04:05.999")
	fmt.Fprintf(f.file, nowstr)
	fmt.Fprintf(f.file, format, args...)
	fmt.Fprintln(f.file)
}
func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	fmt.Fprintf(f.file, format, args...)
}
func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	fmt.Fprintf(f.file, format, args...)
}
func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	fmt.Fprintf(f.file, format, args...)
}
func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	fmt.Fprintf(f.warnFile, format, args...)
}
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	fmt.Fprintf(f.warnFile, format, args...)
}
func (f *FileLogger) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	f.level = level
}
func (f *FileLogger) Close() {
	f.file.Close()
	f.warnFile.Close()
}
