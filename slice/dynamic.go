package main

import "fmt"

func main() {
	//声明slice是一个切片，并且初始化
	slice1 := []int{1, 2, 3, 4}
	printSlice(slice1)

	//声明切片但没有分配空间
	var slice2 []int
	//fmt.Println(slice2[0]) //数组越界错误

	var slice3 []int = make([]int, 3)
	fmt.Println(slice3[0]) //分配空间
 
	//省略var  最常用
	slice4 := make([]int,3)
	fmt.Println(slice4[0])

	//判断空切片
	if slice2 == nil {
		fmt.Println("slice2是一个空切片")
	} else {
		fmt.Println("slice2有空间")
	}
}

func printSlice(array []int) {
	for _, value := range array {
		fmt.Print(value," ")
	}
	fmt.Print("length :",len(array))
}