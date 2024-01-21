package main

import "fmt"

func main() {
	slice := make([]int,3, 3)
	fmt.Println("slice长度：",len(slice),"slice容量：",cap(slice))

	//在slice末尾追加一个元素
	//当slice的容量满的时候，底层会自动开辟空间，新开辟的大小就是上次的容量，相当于扩容一倍
	slice = append(slice, 1)
	//fmt.Println(append(slice, 1))
	fmt.Println("slice长度：",len(slice),"slice容量：",cap(slice))

	//out；
	// slice长度： 3 slice容量： 3
	// slice长度： 4 slice容量： 6
}