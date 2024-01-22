package main

import "fmt"

//定义一个父类
type Person struct {
	name string
	age int
}

func (this *Person) Walk() {
	fmt.Println("Person walk...")
}

func (this *Person) Talk() {
	fmt.Println("Person talk...")
}

//定义子类
type Student struct {
	Person
	id int
}

//子类可以重写父类方法
func (this *Student) Walk() {
	fmt.Println("Student Walk...")
}

//扩展新的方法
func (this *Student) Eat() {
	fmt.Println("Student Eat...")
}

func main() {
	p := Person{name: "人",age: 100};
	fmt.Println("父类：",p)
	p.Walk()
	p.Talk()
	fmt.Println("========================")

	//s := Student{Person{name:"cjp",age : 20},1001}
	//也可以这么声明子类
	var s Student
	s.name = "cjp"
	s.age = 20
	s.id = 1001
	fmt.Println("子类：",s)
	s.Walk()
	s.Talk()
	s.Eat()
}