package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:16379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	defer c.Close()
	_, err = c.Do("expire", "abc", 10) // 设置缓存的超时时间，10秒之后过期
	if err != nil {
		fmt.Println(err)
		return
	}
}
