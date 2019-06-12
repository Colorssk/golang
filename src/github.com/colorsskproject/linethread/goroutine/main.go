package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello goroutine")
}

func main() {
	go hello() // 加了go以后就让这个函数另起线程执行，不会等待这个函数执行完
	fmt.Println("main thread terminate")
	time.Sleep(time.Second)
}
