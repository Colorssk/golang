package main

import (
	"fmt"
)

func main() {
	// 申明的时候指明内存分配
	var a *int = new(int) // 直接给指针分配新的内存
	*a = 100
	fmt.Println(*a)
	// new 给切片分配内存，但是还是要用make才能够初始化
	var b *[]int = new([]int)
	fmt.Println(*b)
	(*b) = make([]int, 5, 100)
	(*b)[0] = 100
	fmt.Println(*b)

}
