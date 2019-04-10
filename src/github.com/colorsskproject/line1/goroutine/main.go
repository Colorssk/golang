package main

import "fmt"
import "time"

func calc() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("run",i,"times")
	}
	fmt.Println("calc finished")
}
func main() {
	//calc()
	go calc()
	fmt.Println("main exit")
	time.Sleep(11*time.Second)// 目的是为了让子线成能够执行完，否则主线程一结束退出之后，子线程也会被关闭
}
