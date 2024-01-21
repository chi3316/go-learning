package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4}
	//[lower-bound:upper-bound)   这样新切片的长度 =  upper-bound - lower-bound
	part1 := slice[1:3]
	fmt.Println(part1)

	//[0,2)
	part2 := slice[:3]
	fmt.Println(part2)

	//[1,len(slice))
	part3 := slice[1:]
	fmt.Println(part3)

	//[0,len(slice)) 截取全部
	part4 := slice[:]
	fmt.Println(part4)
	
	//这样对slice进行截取后，part和原来的slice还是指向同一块内存空间，只是它们维护的头尾指针不同
	//对slice[0]进行修改 , part[0]也会跟着改变
	slice[0] = 999
	fmt.Println(slice)
	fmt.Println(part2)

	//要实现slice的深拷贝，使用copy(new,old)
	newSlice := make([]int,len(slice))
	copy(newSlice,part3)
	fmt.Println(slice)
	fmt.Println(newSlice)
}
