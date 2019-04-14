package main

import (
	"fmt"
	"math/rand" // 随机函数的包 rand.intn(1000)产生0到999的随机数
	"time"
)

func main() {
	//  a:=[3]int{10,2,3}
	//  b:=[...]int{1,2,3}
	//c:=[5]int{2:10}//第2个元素是10
	// 数组长度是类型的一部分，比如b和a长度不同不能直接赋值

	// for index,val := range a{// 遍历数组

	// }

	// 二维数组
	var a [3][2]string = [3][2]string{
		{"a", "b"},
		{"a", "b"},
		{"a", "b"},
	}
	for index, val := range a { // 遍历数组 index是行号, val是对应的数组
		fmt.Println(index, val)
	}
	// 赋值拷贝
	var d [3]int = [3]int{1, 3, 4}
	b := d // 赋值之后是深拷贝，重新开辟了内存空间（说明这里的数组是个值类型不是引用类型）
	b[0] = 2
	fmt.Println(d)
	rand.Seed(time.Now().Unix()) // 初始化一个随机种子
	//此时使用产生的随机数不一样
	fmt.Println(rand.Intn(1000))

}
