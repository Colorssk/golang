package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) { // 形参类型这里也要注意
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done() // 计数减一
}
func main() {
	no := 3
	var wg sync.WaitGroup // 相对于计数器（设定为一种数据类型）
	wg.Wait()
	fmt.Println("wait return")
	for i := 0; i < no; i++ {
		wg.Add(1)          // 计数加一
		go process(i, &wg) //参数作为指针类型
	}
	wg.Wait() // 计数为零这里才会有返回
	fmt.Println("All go routines finished executing")
}
