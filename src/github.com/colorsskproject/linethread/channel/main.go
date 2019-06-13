package main

import (
	"fmt"
)

func main() {
	var c chan int //定义channel容器
	fmt.Printf("c=%v\n", c)

	c = make(chan int, 1) // 实例化（带缓存区的队列，因为指定了大小）
	fmt.Printf("c=%v\n", c)
	c <- 100 //入队

	/*
		data := <-c// 出队到data中
		fmt.Printf("data:%v\n", data)
	*/
	<-c // 仅仅是出队
}
