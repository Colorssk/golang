package main

import (
	"fmt"
	"time"
)

func process(i int, ch chan bool) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	ch <- true
}
func main() {
	no := 3
	exitChan := make(chan bool, no)
	for i := 0; i < no; i++ {
		go process(i, exitChan)
	}
	for i := 0; i < no; i++ { // 所以这里的线程数和goroutine组的数目相同
		<-exitChan //当只有所有线程都完成,此时主线程中的出队操作才不会死锁，最后执行完
	}
	fmt.Println("All go routines finished executing")
}
