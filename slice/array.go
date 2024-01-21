package main

import "fmt"

func main() {
	//定义数组

	//固定数组长度
	var myArray [4]int
	for i, v := range myArray {
		v = i + 1
		fmt.Println("index:",i,"value:",v)
	}

	fmt.Println("========================")

	myArray2  := [5]int {1,2,3}
	var myArray3 = [5]int {0,1,2,3,4}; //还是不要这样写
	printArray(myArray2)
	fmt.Println("========================")
	printArray(myArray3)
}

//将固定数组作为参数传递，类型要匹配，这里就接收不了[4]int
//而且传递固定数组是值拷贝
func printArray(myArray [5]int) {
	for i, v := range myArray {
		fmt.Println("index:",i,"value:",v)
	}
}