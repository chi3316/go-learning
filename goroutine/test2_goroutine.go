package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//创建一个go程承载一个形参为空，返回值为空的匿名函数
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")

			//退出当前的goroutine
			runtime.Goexit()

			fmt.Println("B")
		}() //()表示调用一个匿名函数

		fmt.Println("A")
	}()

	//如何得到这个go程执行结束后的返回值true ,无法直接定义变量去接收，而是需要通过channel
	go func(a int , b int) bool{
		fmt.Println("a = ",a,"b = ",b)
		return true
	}(10,20)

	//死循环保证main程不死亡
	for {
		time.Sleep(1 * time.Second)
	}
}