package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const (
	NumStr  = "0123456789"
	CharSet = "abcdefghijklmnopqrstuvwxyz"
	SpexStr = "!@#$%^&*"
)

var (
	length  int
	charset string
)

func parseArgs() {
	//用户可以通过-l指定生成密码的长度
	flag.IntVar(&length, "l", 16, "-l生成密码的长度") // 第一个参数取得是整数的地址，第二个参数是名称，第三个值是值
	flag.StringVar(&charset, "t", "num",
		`-t 制定生成密码的字符集,
	num:只使用数字[0-9],
	char只使用英文字母[a-zA-Z],
	mix使用数字和字符串,
	advance:使用数字字母和特殊字符`)
	flag.Parse()
}
func generate() string {
	var password []byte = make([]byte, length, length) // 注意：如果本身的长度不够(第一个参数)那么直接password[]访问就会越界
	var SourceStr string
	if charset == "num" {
		SourceStr = "NumStr"
	} else if charset == "char" {
		SourceStr = CharSet
	} else if charset == "mix" {
		SourceStr = fmt.Sprintf("%s%s", NumStr, CharSet)
	} else if charset == "advance" {
		SourceStr = fmt.Sprintf("%s%s%s", NumStr, CharSet, SpexStr)
	} else {
		SourceStr = "NumStr"
	}
	//fmt.Println(SourceStr)
	for i := 0; i < length; i++ {
		index := rand.Intn(len(SourceStr)) // 生成下标
		password[i] = SourceStr[index]
	}
	return string(password)
}
func main() {
	rand.Seed(time.Now().UnixNano())
	parseArgs()
	// fmt.Println(length)
	//fmt.Println(charset)
	password := generate()
	fmt.Println(password)
}

// 通过命令行参数指定值
// 运行脚本的时候指定脚本 password.exe -l 32 (指定长度)
// 命令执行: password.exe -l 32 -t mix
