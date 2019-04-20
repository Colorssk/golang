package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := make(map[string]int, 1024)
	for i := 0; i < 100; i++ {
		key := string(i)
		value := rand.Intn(1000)
		a[key] = value
	}
	// 设置key的切片
	keys := make([]string, 0, 100)
	for key, _ := range a {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 实际这里并没有对无序的map进行排序而是有序了对应的key值
	for _, value := range keys {
		fmt.Println(a[value])
	}
}
