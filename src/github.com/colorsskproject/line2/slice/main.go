package main

import (
	"fmt"
)

func sumarr(a []int) int {
	var sum int = 0
	for _, v := range a {
		sum = sum + v
	}
	return sum
}
func testaumarr() {
	var a [10]int = [10]int{1, 2, 3, 4, 5, 5, 6}
	sum := sumarr(a[:])
	fmt.Println(sum)
}

func main() {

}
