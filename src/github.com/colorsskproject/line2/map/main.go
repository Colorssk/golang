package main

import (
	"fmt"
)

func main() {
	var a map[string]int
	if a == nil { //同理其他引用类型都是用的nil来判断,比如切片，指针类型
		a = make(map[string]int, 16) // 第二个参数表示map的容量，也可以不指定
		//此时初始化之后才可以赋值
		a["colorssk"] = 11
	}
	fmt.Printf("值是:%v", a)
	// 申明时候初始化
	key := "aa"
	b := map[string]int{
		"aa": 11,
		"bb": 22,
	}
	fmt.Printf("b的值是:%v\n", b)
	fmt.Println(b[key])
}
