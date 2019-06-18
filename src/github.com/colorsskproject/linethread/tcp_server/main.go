package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	for {
		conn, err := listen.Accept() // 服务端接受到客户端的确认连接
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}

		go process(conn)
	}
}

func process(conn net.Conn) {

	defer conn.Close()
	for {
		var buf [128]byte           // 传输字节流
		n, err := conn.Read(buf[:]) // 读取类型是切片，存储到切片中，n是长度
		if err != nil {
			fmt.Printf("read from conn failed, err:%v", err)
			break
		}

		str := string(buf[:n])
		fmt.Printf("recv from client, data:%v\n", str)
	}
}
