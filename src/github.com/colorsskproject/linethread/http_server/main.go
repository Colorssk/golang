package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                           // 解析参数，默认是不会解析的
	fmt.Fprintf(w, "%v\n", r.Form)          // Fprintf的第一个参数是io.writer，正好w也是，所以这里可以直接写到页面
	fmt.Fprintf(w, "path:%s\n", r.URL.Path) // 获取表单数据和URL
	fmt.Fprintf(w, "schema:%s\n", r.URL.Scheme)
	fmt.Fprintf(w, "hello world\n")
}

func main() {
	http.HandleFunc("/", sayHello)           // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口（底层还是tcp还是需要监听端口）
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}

}
