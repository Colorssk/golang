package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {
		t, err := template.ParseFiles("./login.html") // 解析返回模板
		if err != nil {
			fmt.Fprintf(w, "load login.html failed")
			return
		}

		t.Execute(w, nil) // 模板调用函数输出到浏览器(w)
	} else if method == "POST" {
		r.ParseForm()
		username := r.FormValue("username") // 获取参数
		passwd := r.FormValue("password")
		fmt.Printf("username:%s\n", username)
		fmt.Printf("passwod:%s\n", passwd)
		if username == "admin" && passwd == "admin123" {
			fmt.Fprintf(w, "username:%s login success\n", r.FormValue("username"))
		} else {
			fmt.Fprintf(w, "username:%s login failed\n", r.FormValue("username"))
		}
	}
}

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("listen server failed, err:%v\n", err)
		return
	}
}
