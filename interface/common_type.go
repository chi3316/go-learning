package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("myFunc arg:",arg)
	value , ok := arg.(string)
	if ok {
		fmt.Println("arg:",value,"的类型是string")
	}
}

func main() {
	myFunc(666)
	myFunc("cjp")
}