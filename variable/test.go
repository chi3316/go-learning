package main

import "fmt"

func main() {
	test := 101
	if test > 100 {
		//var test = 101
		test := 100
		fmt.Print(test)
	}
}