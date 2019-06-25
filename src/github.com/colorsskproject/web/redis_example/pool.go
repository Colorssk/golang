package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool

//初始化一个pool
func newPool(server, password string) *redis.Pool { // 连接池的函数格式
	return &redis.Pool{
		MaxIdle:     64,                //空闲连接数量
		MaxActive:   1000,              // 最大活跃数量
		IdleTimeout: 240 * time.Second, // 空闲超时时间
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			/* //如果有密码
			   if _, err := c.Do("AUTH", password); err != nil {
			   c.Close()
			   return nil, err
			   }*/
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute { // 看此链接是否有效
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func main() {
	pool = newPool("localhost:16379", "")
	for {
		time.Sleep(time.Second)
		conn := pool.Get() // 获取连接
		conn.Do("set", "abc", 100)

		r, err := redis.Int(conn.Do("get", "abc"))
		if err != nil {
			fmt.Printf("do failed, err:%v\n", err)
			continue
		}

		fmt.Printf("get from redis, result:%v\n", r)
	}

	/*
		c, err := redis.Dial("tcp", "localhost:16379")
		if err != nil {
			fmt.Println("conn redis failed,", err)
			return
		}
		defer c.Close()
		_, err = c.Do("Set", "abc", 100)
		if err != nil {
			fmt.Println(err)
			return
		}
		r, err := redis.Int(c.Do("Get", "abc"))
		if err != nil {
			fmt.Println("get abc failed,", err)
			return
		}
		fmt.Println("get from redis", r)
	*/
}
