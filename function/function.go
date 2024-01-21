package main

import "fmt"

func main() {
	fmt.Println("return ",myfunc1("p1","p2"))
	// fmt.Println("return ",myfunc2("p1","p2")) multiple-value myfunc2("p1", "p2") (value of type (int, int)) in single-value context
	result1,result2 := myfunc2("p1","p2")
	fmt.Println("return ",result1,result2)

	result3,result4 := myfunc3("p1","p2")
	fmt.Println("return ",result3,result4)

	result5,result6 := myfunc4("p1","p2")
	fmt.Println("return ",result5,result6)
}

// 定义函数 func关键字
func myfunc1(a string, b string) int {
	fmt.Printf("myfunc1 形参列表：%s %T,%s %T ",a,a,b,b)  
	return 77
}

//返回多个值
func myfunc2(a string, b string) (int,int) {
	fmt.Printf("myfunc2 形参列表：%s %T,%s %T ",a,a,b,b)
	return 66,77
}

//返回多个值，并给返回值命名
func myfunc3(a string, b string) (r1 int,r2 int) {
	fmt.Printf("myfunc2 形参列表：%s %T,%s %T ",a,a,b,b)
	//r1 r2 作为两个局部变量，有默认值
	r1 = 666
	r2 = 777
	return 
}

//返回多个值，并给返回值命名，并且多个返回值的类型相同
func myfunc4(a string, b string) (r1 ,r2 int) {
	fmt.Printf("myfunc2 形参列表：%s %T,%s %T ",a,a,b,b)
	//r1 r2 作为两个局部变量，有默认值，在这里会返回0
	return 
}