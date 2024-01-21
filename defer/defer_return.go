package main

import "fmt"

func return_func() int{
	fmt.Println("return end.")
	return 1
}

func defer_func() {
	fmt.Println("defer end.")
}

func test() int {
	defer defer_func()
	return return_func() 
}
func main()  {
	test()
}