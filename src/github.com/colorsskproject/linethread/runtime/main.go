package main

import (
	"time"
	//"runtime"
)

var i int

func calc() {
	for {
		i++
	}
}

func main() {
	/*
		cpu := runtime.NumCPU()
		fmt.Println("cpu:", cpu)

		runtime.GOMAXPROCS(cpu)// 设置这个程序占用的核数
	*/
	for i := 0; i < 10; i++ {
		go calc()
	}

	time.Sleep(time.Hour)
}
