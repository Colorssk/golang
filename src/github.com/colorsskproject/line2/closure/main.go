package main

import (
	"fmt"
	"time"
)

// 闭包（函数里面还有函数，函数实例化之后变量不会立即释放）
func add() func(int) int {
	var x int // 闭包所以这个变量不会被释放，所以作为全局可以一直使用
	return func(d int) int {
		x += d
		return x
	}
}

//闭包副作用
func testClosure() {
	for i := 0; i < 5; i++ {
		// go func() {
		// 	fmt.Println(i)// 输出的都是5 闭包的副作用 i 这个i可以立即始终指向一个i
		// }()
		go func(index int) {
			fmt.Println(index) // 输出的是正常的
		}(i)
	}
	time.Sleep(time.Second)
}
func main() {
	re_fn := add() // re_fn未释放，x也就不会释放
	result := re_fn(1)
	fmt.Printf("%d\n", result)
	result = re_fn(2)
	fmt.Printf("%d", result)
	//re_fn2 := add() // 此时x重新计算
	testClosure()
}
