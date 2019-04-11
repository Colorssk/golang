package main

import (
	"fmt"
	"time"
)

// 时间戳转为时间
func stamptotime(timepara int64) time.Time {
	timeobj := time.Unix(timepara, 0)
	return timeobj
}

// 定时器
func timeTicker() {
	ticker := time.Tick(time.Second) // 返回一个管道的时间
	for i := range ticker {          // 根据循环时间来启动定时器
		fmt.Println("运行中", i)
	}
}
func main() {
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Printf("%02d", now.Month()) //直接转为数字而不是展示位英文
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second()) // printf(%02d)// 表示至少两位未满补零
	fmt.Println("当前时间是", stamptotime(now.Unix()))
	timeTicker()
}
