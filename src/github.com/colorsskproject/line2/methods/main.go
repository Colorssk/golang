package main

import (
	"fmt"
)

type People struct { // 结构体是值类型的
	Name   string
	Contry string
}

func (p People) Print() { // 这个的p是p1的拷贝所以只能读取，改变不会影响
	// 里面也能访问到p中的属性
	fmt.Println(p.Name, p.Contry)
}
func main() {
	p1 := People{
		Name:   "aa",
		Contry: "China",
	}
	// 此时Print方法就是p1内部的方法
	p1.Print()
}
