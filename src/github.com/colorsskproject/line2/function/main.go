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
func main() {
	sum := functest(10, 20, 30)
	fmt.Println(sum)
}
