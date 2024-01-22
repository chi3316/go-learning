package main

import "fmt"

func main() {
	cityMap := make(map[string]string)
	fmt.Println(cityMap)

	//添加
	cityMap["China"] = "Bejing"
	cityMap["Japan"] = "Tokoy"
	fmt.Println(cityMap)

	//遍历
	for key, value := range cityMap {
		fmt.Println("key", key, ":", "value", value)
	}

	//删除 delete关键字  传入key
	delete(cityMap, "Japan")
	fmt.Println(cityMap)

	//修改
	cityMap["China"] = "HuiLai"
	printMap(cityMap)

	//拷贝map
	m4 := make(map[string]string)
	for key, value := range cityMap {
		m4[key] = value
	}
	fmt.Println("\n",m4)
}

//使用map进行传参，引用传递
func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Print("key = ", key, " value = ", value)
	}
}
