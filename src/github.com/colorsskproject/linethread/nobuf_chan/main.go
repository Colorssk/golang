package main

import (
	"fmt"
	"time"
)

func produce(c chan int) {
	c <- 1000                       //不带缓存区所以这里无法压入队列
	fmt.Println("produce finished") // consume线程中没取数据，上一步channel会阻塞
}

func consume(c chan int) {
	data := <-c // 这里有取得操作，上面的压入操作就可以执行，相对的这里也会有数据
	fmt.Println(data)
}

func main() {
	var c chan int
	fmt.Printf("c=%v\n", c)

	c = make(chan int) // 不带缓存区的队列
	go produce(c)      //  启动线程
	go consume(c)
	time.Sleep(time.Second * 5)
}
