package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"你的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ",tagInfo,"doc: ",tagDoc)
	}
}

func main() {
	var r resume
	findTag(&r)
}