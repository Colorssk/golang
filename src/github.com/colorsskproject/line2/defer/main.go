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
		//defer fmt.Printf("%d", i) // 输出4,3,2,1,0因此defer是吧执行结果存储起来(所以保存的是结果)，函数返回时候再逆序执行
		defer func() { // 匿名函数加上defer之后保存的是个函数，那么i的值在后续改变会影响到这里的执行效果的,所以输出后是55555
			fmt.Println() // 原因匿名函数这里没有传参数进去，所以执行的时候(函数返回的时候)会在作用域内搜索变量
		}()
		// defer func(a int) {
		// 	fmt.Println(a) //此时就是输出4,3,2,1,0因为传参了
		// }(i)
	}
	return
}
func main() {
	// testDefer()
	testdefer02()
}
