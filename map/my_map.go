package main

import "fmt"

func main() {
	//var myMap map[string]string = make(map[string]string, 10)
	var myMap map[string]string
	fmt.Print(myMap == nil,"\n")

	//var m = make(map[string]string, 10)
	
	m := make(map[string]string, 10) 
	//也可以不指定大小   m := make(map[string]string) 
	m["java"] = "wife"
	m["go"] = "lover"
	fmt.Println(m)

	//声明的同时顺便初始化值，就不需要make分配空间了
	m3 := map[string]string {
		"wife: " : "java",
		"lover" : "go",
		"anotherLover" : "python",
	}

	fmt.Println(m3)
}