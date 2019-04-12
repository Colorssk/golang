package main

import (
	"fmt"
)

func testDefer() {
	defer fmt.Println("hello")
	defer fmt.Println("hello2")
	defer fmt.Println("hello3") //执行顺序是倒着执行的
	fmt.Println("aaa")          // 先执行
}
func testdefer02() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d", i) // 输出4,3,2,1,0因此defer是吧执行结果存储起来(所以保存的是结果)，函数返回时候再逆序执行
	}
}
func main() {
	testDefer()
	testdefer02()
}
