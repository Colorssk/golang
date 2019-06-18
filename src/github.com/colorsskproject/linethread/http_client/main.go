package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80") //
	if err != nil {
		fmt.Printf("dial failed, err:%v\n", err)
		return
	}

	data := "GET / HTTP/1.1\r\n" // 数据包  需要/r/n
	data += "HOST: www.baidu.com\r\n"
	data += "connection: close\r\n"
	data += "\r\n\r\n"

	_, err = io.WriteString(conn, data) // 往连接中写入数据
	if err != nil {
		fmt.Printf("wirte string failed, err:%v\n", err)
		return
	}

	var buf [1024]byte
	for {
		n, err := conn.Read(buf[:]) // 读取请求之后接受到的数据
		if err != nil || n == 0 {
			break
		}

		fmt.Println(string(buf[:n]))
	}
}
