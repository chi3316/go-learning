package main

import (
	"fmt"
	"time"
)

//主goroutine
func main() {
	//主函数启动了 5 个 goroutine，每个 goroutine 打印一些数字。由于 goroutine 的轮询调度，它们会交替执行，产生交织输出。
	for i := 0; i < 5; i++ {
		go printNumbers(i)
	}

	time.Sleep(time.Second)
}


//子goroutine
func printNumbers(id int) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Goroutine %d: %d\n", id, i)
		time.Sleep(100 * time.Millisecond)
	}
}
