package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:20000") // 拨号连接
	if err != nil {
		fmt.Printf("dial failed, err:%v\n", err)
		return
	}

	reader := bufio.NewReader(os.Stdin) // 从终端输入测试数据
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed, err:%v\n", err)
			break
		}

		data = strings.TrimSpace(data)
		_, err = conn.Write([]byte(data)) // write发送给服务端
		if err != nil {
			fmt.Printf("write failed, err:%v\n", err)
			break
		}
	}
}
