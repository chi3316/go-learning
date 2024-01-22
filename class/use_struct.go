package main

import "fmt"

//type 关键字  给数据类型取一个别名
type myInt int

//定义结构体
type Book struct {
	title string
	price myInt
}

func main() {
	var book Book;
	book.price = 10
	book.title = "Java"
	fmt.Println(book)

	changeBook1(book)
	fmt.Println("book1 :",book)

	changeBook2(&book)
	fmt.Println("book2 :",book)
}

//结构体当形参
//值传递
func changeBook1(book Book) {
	book.title = "Golang"
}
//引用传递
func changeBook2(book *Book) {
	book.title = "Golang"
}