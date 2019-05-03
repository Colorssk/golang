package main

import (
	"github.com/colorsskproject/line2/log"
)

func main() {
	// file := log.NewFileLog("c:/a.log")
	// file.LogDebug("this is a debug log")
	// file.LogWarn("this is a warn log")
	// 定义接口以后切换实例
	log := log.NewFileLog("aaa")
	//或者 log：=log.NewConsoleLog("aaa")
	log.LogDebug("this is a debug log")
	log.LogWarn("this is a warn log")
}
