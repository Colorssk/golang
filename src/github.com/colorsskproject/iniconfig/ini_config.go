package iniconfig

import (
	//"fmt"
	"errors"
	"reflect"
	"strings"
)

// 序列化
func Marshal(data interface{}) (result []byte, err error) {
	return
}

//反序列化
func UnMarshal(data []byte, result interface{}) (err error) {
	lineArr := strings.Split(string(data), "\n")
	_ = lineArr
	typeInfo := reflect.TypeOf(result)
	// 因为传递过来的是指针
	if typeInfo.Kind() != reflect.Ptr {
		err = errors.New("please pass address")
		return
	}
	// for _, v := range lineArr {
	// 	fmt.Printf("%v\n", v)
	// }
	var m map[string]string
	_ = m
	return
}
