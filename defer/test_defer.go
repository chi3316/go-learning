package main

import "fmt"

func main() {
	defer fmt.Println("main end1") //在这个函数里面执行结束之前会被调用，类似析构函数,defer后面跟的就是让这个函数（defer所在的那里）结束之前做的事，有多个defer以栈的形式，先进后出
	defer fmt.Println("main end2") 

	fmt.Println("main::hello go1")
	fmt.Println("main::hello go2")
}