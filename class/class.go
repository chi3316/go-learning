// 与结构体绑定，添加方法
package main

import "fmt"

type Book struct {
	title string
	price int

	Publish_year string
}

//将当前方法与结构体绑定
/*
func (this Book) Show() {
	 fmt.Println(this)
}

func (this Book) GetTitile() string {
	return this.title
}

func (this Book) SetTitile(title string) {
	
	this.title = title
}
*/

//像上面那样写this指向的是调用该方法的对象的一个副本
//一般是要加个*,才是指向该对象
func (this *Book) Show() {
	fmt.Println(this)
}

func (this *Book) GetTitile() string {
   return this.title
}

func (this *Book) SetTitile(title string) {
   this.title = title
}

func main() {
	//创建一个对象
	book := Book{title: "java",price: 100}
	fmt.Println("title:",book.GetTitile())
	book.Show()

	book2 := Book{}
	book2.price = 100
	book2.title = "C++"
	book2.Show()
}