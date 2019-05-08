package main

import (
	"fmt"
)
//ipslice:=strings.Split(ips,";")
func main(){
	var(
		a int
		b bool
		c string = "hello"
		d float32
	)
	const e string = "world"
	const f  = 3/9
    // const(
	// 	a int = 100
	// 	b// b继承上面a的值100
	// 	c int = 200
	// 	d// d继承c的值200
	// )
	// const(
	// 	a = iota// a=0后面的数值递增 (每隔一行新增一，所以只要有iota就是看在第几行)
	// 	b// 等同于b=iota
	// 	c
	// )
	const(
		a1 = 3 << iota
		a2
		a3
		a4
		a5
		a6
	)
	var str [] string = [] string{"a","b","c"}
	result :=strings.Join(str,';')
	fmt.Println(a1,a2,a3,a4,a5,a6)//1<<iota:1 2 4 8 16 32    2<<iota:2 4 8 16 32 64  3<iota:3 6 12 24 48 96
	fmt.Printf("a=%d,b=%t,c=%s,d=%f\n",a,b,c,d)
}