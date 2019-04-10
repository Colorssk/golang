package main

import "fmt"
// 翻转字符串(英文)
// func testreversestr(){
// 	var str = "hello"
// 	var bytes []byte = []byte(str)

// 	for i:=0; i<len(str)/2;i++{
// 		tmp:=bytes[len(str)-i-1]
// 		bytes[len(str)-i-1] = bytes[i]
// 		bytes[i] = tmp
// 	}
// 	str = string(bytes)
// 	fmt.Println(str)
// }
func testreversestr2(){
	var str = "hello中文"
	var runestr []rune = []rune(str)

	for i:=0; i<len(runestr)/2;i++{
		tmp:=runestr[len(runestr)-i-1]
		runestr[len(runestr)-i-1] = runestr[i]
		runestr[i] = tmp
	}
	str = string(runestr)
	fmt.Println(str)
}
// 判断是否是回文
func ishuiwen()bool{
	var str = "上海wp来水来自海上"
	var runstr []rune = []rune(str)
    var flag bool = true
	for i:=0;i<len(runstr);i++{
		if runstr[i]!= runstr[len(runstr)-i-1] {
			flag = false
		}
	}
	return flag
}
func main(){
	// testreversestr2()
	if ishuiwen() == true {
		fmt.Println("回文")
	}else{
		fmt.Println("不是回文")
	}
}