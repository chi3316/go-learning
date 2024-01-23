package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func testReflect(arg interface{}) {
	fmt.Println("type: ",reflect.TypeOf(arg))
	fmt.Println("value:",reflect.ValueOf(arg))
}

func main() {
	testReflect("echo")
}