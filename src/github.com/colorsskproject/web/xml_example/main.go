package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Servers struct {
	Name    xml.Name `xml:"servers"` // 此处就和xml中的标签做好了映射，方便xml反序列化
	Version string   `xml:"version,attr"`
	Servers []Server `xml:"server"`
}

type Server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func main() {
	data, err := ioutil.ReadFile("./config.xml") //读取文件
	if err != nil {
		fmt.Printf("read config.xml failed, err:%v\n", err)
		return
	}
	var servers Servers
	err = xml.Unmarshal(data, &servers) // 使用xml的反序列化
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("xml: %#v\n", servers)
}
