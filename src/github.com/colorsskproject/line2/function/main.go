package main

import (
	"fmt"
)

func functest(b ...int) int {
	sum := 0
	for i := 0; i < len(b); i++ {
		sum = sum + b[i]
	}
	return sum
}
func functest2(a int, b ...int) int {
	sum := a
	for i := 0; i < len(b); i++ {
		sum = sum + b[i]
	}
	return sum
}
func add(a, b int) int {
	return a + b
}

// 参数是函数
func calc(a, b int, op func(int, int) int) int {
	return op(a, b)
}
func main() {
	//f1 := functest2 // 函数赋值 打印类型 %T
	// f2 := func(a int) { // 匿名函数
	// 	fmt.Println(a)
	// }
	sum := functest(10, 20, 30)
	result := calc(10, 20, add)
	fmt.Println(sum, result)
}
