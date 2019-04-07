package main// main包用来生成可执行文件的，每个程序只有一个main包
import "fmt"
import "time"

func main () {
	fmt.Println("hello world")
	time.Sleep(time.Second*10)
}