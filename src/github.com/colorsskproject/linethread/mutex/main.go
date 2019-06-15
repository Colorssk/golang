package main

import (
	"fmt"
	"sync"
)

var x int
var wg sync.WaitGroup
var mutex sync.Mutex // 互斥锁

func add() {
	for i := 0; i < 5000; i++ {
		mutex.Lock() // 加锁(同步操作中就会只被一个线程调用，除非解锁才能被其他线程使用)
		x = x + 1
		mutex.Unlock() // 解锁
	}
	wg.Done()
}

func main() {

	wg.Add(2)
	go add()
	go add()

	wg.Wait()
	fmt.Println("x:", x)
}
