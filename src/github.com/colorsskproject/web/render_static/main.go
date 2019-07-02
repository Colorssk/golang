package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static") // 后面一个参数是资源路径
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
