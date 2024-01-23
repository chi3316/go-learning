package main

import (
	"encoding/json"
	"fmt"
)

//给字段添加元数据，以便在运行时进行解释和处理。很像java的注解
type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"果宝特攻", 2002, 20, []string{"甘蔗精", "苹果精", "菠萝精"}}

	//编码，将struct -> json
	jsonStr , err := json.Marshal(movie)
	if err != nil {
		fmt.Println("Json marshal error",err)
		return
	}

	fmt.Printf("jsonStr = %s\n",jsonStr)
}