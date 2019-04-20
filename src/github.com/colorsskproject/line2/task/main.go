package main

import (
	"fmt"
	"strings"
)

// 计算字符串中单词出现的个数
func calc(str string) map[string]int {
	result := make(map[string]int, 100)
	words := strings.Split(str, " ") //返回的是一个切片
	for _, value := range words {
		_, ok := result[value]
		if !ok {
			result[value] = 1
		} else {
			result[value]++
		}
	}
	return result
}
func main() {
	result := calc("hello world yy hh yy")
	fmt.Println(result)
}
