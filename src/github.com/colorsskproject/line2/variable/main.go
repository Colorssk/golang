package main

// import (
// 	"fmt"
// )

var a int = 100 // 全局变量
func main() {
	//var a int = 200 // 局部变量优先级高于全局变量
	var b int = 200 // 变量作用域：局部变量
	if b == 200 {
		//var c int ==300// 只在语句块中有效
	}
}
