package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read a book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a book")
}

func main() {
	//book : pair<type : Book , value : Book{}地址>
	book := &Book{}

	//r : pair<type:,value:>  未确定
	var r Reader
	r = book  //r : pair<type : Book , value : Book{}地址>
	r.ReadBook()

	//w : pair<type:,value:>  未确定
	var w Writer
	w = book //w : pair<type : Book , value : Book{}地址>
	w.WriteBook()

	w = r.(Writer) //此处的断言能成功，因为 w , r这两个变量的pair中的type是一致的
}