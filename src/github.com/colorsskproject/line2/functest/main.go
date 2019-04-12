package main

import (
	"fmt"
)

//字数统计
func calc(str string) (charcount int, numcount int, spacecount int, othercount int) {
	utfchars := []rune(str)
	for i := 0; i < len(utfchars); i++ {
		if utfchars[i] >= 'a' && utfchars[i] <= 'z' || utfchars[i] >= 'A' && utfchars[i] <= 'Z' {
			charcount++ // 无需声明
			continue
		}
		if utfchars[i] >= '0' && utfchars[i] <= '9' {
			numcount++
			continue
		}
		if utfchars[i] == ' ' {
			spacecount++
			continue
		}
		othercount++

	}
	return // 此时无需指定返回的具体值，但一定要有返回标识
}
func main() {
	charcount, numcount, spacecount, othercount := calc("asdf啊实打实地方123123  sadf sdf 驱蚊器")
	fmt.Println(charcount, numcount, spacecount, othercount)
}
