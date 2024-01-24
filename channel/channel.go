package main

import "fmt"

func main() {
	c := make(chan int)
	//如何得到这个go程执行结束后的返回值val ,无法直接定义变量去接收，而是需要通过channel
	go func(a int, b int) int {
		fmt.Println("a = ", a, "b = ", b)
		val := a + b
		c <- val
		return val
	}(10, 20)
	
	num := <- c
	fmt.Println(num)
}
