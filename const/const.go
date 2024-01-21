package main

import "fmt"

func main() {
	//常量的定义
	const length int = 100
	const another = 9
	fmt.Println("length =",length,"another =",another)
	fmt.Println("BEIJING =",BEIJING,"SHANGHAI =",SHANGHAI)
	fmt.Println("SUZHOU =",SUZHOU,"NANJING =",NANJING)
	fmt.Println("JIEYANG =",JIEYANG,"QINGDAO =",QINGDAO)
	fmt.Println("a =",a,"b =",b,"c =",c,"d =",d)
	fmt.Println("e =",e,"f =",f,"g =",g,"h =",h)
}

//const来定义枚举类型
const (
	BEIJING = 0
	SHANGHAI = 1
	SHENZHEN = 2
	GUANGZHOU = 3
)

//关键字 iota 每行的iota都会累加1，第一行的iota默认值是 0 
const (
	// 等价于 0 1 2 3
	SUZHOU = iota
	NANJING
	TAIPEI
	CHENGDU
)

const (
	// 等价于 0 10 20 30
	JIEYANG = iota * 10
	QINGDAO
	SHANTOU
	FUJIAN
)

const (
	//a = 1,b = 2,c = 2,d = 3
	a,b = iota + 1,iota + 2
	c,d 

	//e = 20,f = 40,g = 30,h = 60
	e,f = iota * 10,iota * 20
	g,h
)