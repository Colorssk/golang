package main

import (
	"fmt"
	"io/ioutil"

	"github.com/colorsskproject/web/protobuf/address" // 代码已经转化，可以直接调用
	"github.com/golang/protobuf/proto"
)

func writeProto(filename string) (err error) {
	var contactBook address.ContactBook //对应的接结构体
	for i := 0; i < 64; i++ {
		p := &address.Person{ // 这里都一律用指针赋值
			Id:   int32(i),
			Name: fmt.Sprintf("陈%d", i),
		}

		phone := &address.Phone{
			Type:   address.PhoneType_HOME,
			Number: "15910624165",
		}

		p.Phones = append(p.Phones, phone)
		contactBook.Persons = append(contactBook.Persons, p)
	}

	data, err := proto.Marshal(&contactBook) // protobuf 序列化
	if err != nil {
		fmt.Printf("marshal proto buf failed, err:%v\n", err)
		return
	}

	err = ioutil.WriteFile(filename, data, 0755)
	if err != nil {
		fmt.Printf("write file failed, err:%v\n", err)
		return
	}
	return
}

func readProto(filename string) (err error) {
	var contactBook address.ContactBook
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = proto.Unmarshal(data, &contactBook) // 反序列化
	if err != nil {
		return
	}

	fmt.Printf("proto:%#v\n", contactBook)
	return
}

func main() {
	filename := "c:/tmp/contactbook.dat"
	err := writeProto(filename)
	if err != nil {
		fmt.Printf("write proto failed, err:%v\n", err)
		return
	}
	err = readProto(filename)
	if err != nil {
		fmt.Printf("read proto failed, err:%v\n", err)
		return
	}
}
