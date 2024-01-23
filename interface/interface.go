package main

import "fmt"

//接口，本质是一个指针
type Animal interface{
	Sleep()
	GetColor() string
	GetType() string
}

//只要重写了这个接口的所有类，就可以用接口类型的指针指向该类
//猫类
type Cat struct {
	color string
	the_type string
}
func (this *Cat) Sleep() {
	fmt.Println("小猫老弟在睡觉...")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return this.the_type
}

//狗类
type Dog struct {
	color string
	the_type string
}
func (this *Dog) Sleep() {
	fmt.Println("小狗老弟在睡觉...")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return this.the_type
}

//使用接口类型作为形参
func showAnimal(animal Animal) {
	fmt.Println(animal.GetColor(),animal.GetType())
	animal.Sleep()
}

func main() {
	//var animal Animal //接口类型的指针
	var cat Animal = &Cat{"黄色","小猫"}
	showAnimal(cat)

	var dog Animal = &Dog{"快乐","小狗"}
	showAnimal(dog)
}