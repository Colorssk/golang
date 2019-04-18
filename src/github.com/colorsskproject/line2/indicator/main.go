package main

import (
	"fmt"
)

func arrmodify(arr *[3]int) {
	(*arr)[0] = 2
}
func modify(arr []int) {
	arr[1] = 1
}
func testPoint() {
	var a int = 1
	var b *int = &a
	c := &a
	fmt.Println(*b)
	fmt.Println(*c)
	//也可以直接修改
	*c = 2
	fmt.Println(a)
}
func main() {
	var a [3]int = [3]int{1, 2, 3}
	arrmodify(&a)
	fmt.Println(a)
	// 等同于用引用类型的切片
	modify(a[:])
	fmt.Println(a)
	testPoint()
}
