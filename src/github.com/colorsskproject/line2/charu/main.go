package main

import (
	"fmt"
)

func insert_sort(a [8]int) [8]int { // 这里的a仅仅是个拷贝所以需要返回
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j] // 无需temp
			} else {
				break
			}
		}
	}
	return a
}

// 插入排序
func main() {
	var i [8]int = [8]int{8, 1, 2, 3, 7, 2, 6, 3} // 数组定义的形式
	j := insert_sort(i)
	fmt.Println(j)
}
