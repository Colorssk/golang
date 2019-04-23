package main

import (
	"encoding/json"
	"fmt"
)

//结构体序列化
type Student struct {
	Id   int
	Name string
}
type Class struct {
	Id       int
	Name     string
	Students []*Student
}

var class *Class

func init() {
	class = &Class{
		Id:   1,
		Name: "一班",
	}
	arr := []string{"stu01", "stu02", "stu03"}
	for key, stu := range arr {
		s := new(Student)
		s.Id = key + 1
		s.Name = fmt.Sprintf("%s", stu)
		class.Students = append(class.Students, s)
	}

}

// 反序列化 成为结构体
func retranfs(jsonstr string) (c1 *Class) {
	c1 = &Class{}
	err := json.Unmarshal([]byte(jsonstr), c1) //第一个参数一定要是赐福数组,转化完之后存储在c1中
	if err != nil {
		fmt.Println("wrong")
		return
	}
	//fmt.Println(c1) // 此时输出的时候student就是地址所以遍历输出的时候要变
	return
}
func main() {
	fmt.Println(class) // 把结构体没序列化之前切片student输出的是地址
	data, err := json.Marshal(class)
	if err != nil {
		fmt.Println("failed")
		return
	}
	fmt.Println(string(data)) // json序列化之后输出的正常
	result := retranfs(string(data))
	fmt.Println(result)
	for _, v := range result.Students {
		fmt.Println((*v)) // 因为指针类型的students中存储的还是地址
	}

}
