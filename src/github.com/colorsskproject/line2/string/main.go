package main

import(
	"fmt"
)
func teststring(){//字符串以数组方式访问只可读不接写
	var str = "hello"

	for index,val:=range str{//循环字符串 index是key val是value
		fmt.Printf("%v:%v",index,val)// 输出的时候可以直接当做数组输出str[0]
	}
	// str[0] = "0"// 字符串复制以后不能够局部使用数组方式改变值，要通过的定义切片改变
	var bytestr []byte = []byte(str)// 此时字符串转化成为字节切片
	bytestr[0]= 'o'//这里修改字符串的时候需要用单引号
	//最后还要把切片转回字符类型
	str =  string(bytestr)
	// len(str)// 一个英文占一个字节，但是一个中文占三个字节
	var b rune = '中'// rune也要用单引号
	fmt.Println(b)
	// 转化为rune切片再求长度
	var runslice [] rune
	runslice = [] rune(str)// 字符串转化为rune切片，此时可以len获取长度
	fmt.Println(len(runslice))

}
func main(){
	teststring()
}