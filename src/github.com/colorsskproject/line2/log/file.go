package log

import "fmt"

type FileLog struct {
}

func NewFileLog(file string) LogInterface { //返回接口类型
	return &FileLog{} // 返回实例对象
}

func (f *FileLog) LogDebug(msg string) { // 这里的方法就代表实例对象继承了接口 
	fmt.Println(msg)

}
func (f *FileLog) LogWarn(msg string) {
	fmt.Println(msg)
}
