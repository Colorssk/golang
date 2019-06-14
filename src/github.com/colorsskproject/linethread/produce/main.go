package main

import (
	"fmt"
)

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func main() {
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch //  读取队列数据，如果关闭了管道，OK==false
		if ok == false {
			fmt.Println("chan is closed")
			break // 管道关闭就终止输出
		}
		fmt.Println("Received ", v)
	}
}
