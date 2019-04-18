package main

import (
	"fmt"
)

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
	testPoint()
}
