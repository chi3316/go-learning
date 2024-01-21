package main

import "fmt"

func main() {
	fmt.Println("============声明单个变量===========")
	//1. 声明不赋值
	var a int
	fmt.Printf("type of a = %T:,value of b = %d\n", a, a)

	//2. 声明并赋值
	var b int = 10
	fmt.Printf("type of b = %T:,value of b = %d\n", b, b)

	//3. 类型推断
	var c = "Go"
	fmt.Printf("type of c = %T:,value of c = %s\n", c, c)

	//4. 省去var，自动匹配  最常用  := 只能声明在局部变量，不能声明全局变量
	d := "cjp"
	//d := 100  type of c = int:,value of c = %!s(int=100) ???
	fmt.Printf("type of d = %T:,value of c = %s\n", d, d)

	fmt.Println("============声明多个变量===========")
	var aa, bb int = 100, 200
	fmt.Println("aa =", aa, "bb =", bb)

	var cc, dd = 300, "陈坚鹏"
	fmt.Println("cc =", cc, "dd =", dd)

	//多行的多变量声明
	var (
		ee int = 100
		ff     = "习近平"
	)
	fmt.Println("ee =", ee, "ff =", ff)

	//_表示变量7的赋值被废弃，变量 _  不具备读特性
	_, value := 7, 5
	fmt.Println(value)
}
