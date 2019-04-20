package main

func studentinterface() {

}
func student() {
	var stuMap map[int]map[string]interface{} = make(map[int]map[string]interface{}, 128) // 最后一个interface{}表示可以是任何数据类型
	id := 1
	name := "clolorssk"
	age := 18
	value, ok := stuMap[id]
	if !ok {
		value = make(map[string]interface{}, 128)
	}

	value["name"] = name
	value["age"] = age
	stuMap[id] = value

	stuMap[id] = value
}
func main() {

}
