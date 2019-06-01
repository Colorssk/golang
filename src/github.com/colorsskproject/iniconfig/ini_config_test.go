package iniconfig

import (
	"fmt"
	"io/ioutil" //用来读取文件
	"testing"
)

type Config struct {
	ServerConf ServerConfig `ini: "server"`
	MysqlConf  MysqlConfig  `ini: "mysql"`
}
type ServerConfig struct {
	Ip   string `ini: "ip"`
	Port int    `ini: "port"`
}
type MysqlConfig struct {
	Username string `ini: "username"`
	Passwd   string `ini: "passwd"`
	Database string `ini: "database"`
	Host     string `ini: "host"`
	Port     int    `ini: "port"`
}

//go test 之后会自动执行文件名是test结尾的方法，同时执行Test开头的方法
func TestIniConfig(t *testing.T) {
	fmt.Println("hello world")
	data, err := ioutil.ReadFile("./config.ini") // 读到的数据是[]byte
	if err != nil {
		t.Error("read file failed")
	}
	//t.Error("hello")   报错log
	var conf Config
	err = UnMarshal(data, &conf)
	if err != nil {
		t.Error("unmarshal failed", err)
	}
	t.Logf("unmarshal success,conf:%v", conf)
}
