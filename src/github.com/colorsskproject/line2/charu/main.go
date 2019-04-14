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

// 冒泡排序
func bop_sort(a [8]int) [8]int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

// 选择排序
func choose_sort(a [8]int) [8]int { // 这里的a仅仅是个拷贝所以需要返回
	for i := 1; i < len(a); i++ {
		for j := i + 1; j < len(a); j-- {
			if a[j] < a[i] {
				a[i], a[j] = a[j], a[i]
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
