## 快速上手

### hello world

```go
package main //当前程序的包名
//导入多个包
import (
	"fmt" 
	"time"
)
//main函数
func main() {
	time.Sleep(1 * time.Second)
	fmt.Println("hello Go!")
}
```

- 程序入口 main包下的main函数

  - 第一行代码**package main**定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
  - **import "fmt"**告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。
  - **func main()**是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。

- 导入多个包的简洁语法

- 分号可加可不加 （最好别加）

- 花括号不换行

- 程序编译 ： `go build hello.go`

  运行 ： `./hello`

- go run 表示 直接编译go语言并执行应用程序，一步完成  `go run hello.go`  

### 变量声明

````go
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
	fmt.Println("aa =",aa,"bb =",bb) 

	var cc, dd = 300, "陈坚鹏"
	fmt.Println("cc =",cc,"dd =",dd) 

	//多行的多变量声明  一般用于声明全局变量
	var(
		ee int = 100;
		ff = "习近平"  // ff string = "习近平"
	)
	fmt.Println("ee =",ee,"ff =",ff) 
    
    //_表示变量7的赋值被废弃，变量 _  不具备读特性
	_, value := 7, 5
	fmt.Println(value) 
}

````

- 最常用的是  `name := "cjp"` ，其他的就是 `var + 变量名 + 变量类型  ` , 其中类型可以省略

- 全局变量要用`var`声明

- 不允许在同一个作用域内声明同名的局部变量。在同一个作用域内，每个变量名必须是唯一的，否则会导致编译错误。

  - c++ 是允许这么做的，而go和java一样，这么做编译不通过

    C++  正常

    ```c++
    #include<iostream>
    int main() {
        int a = 101;
        if(a > 100) {
            int a = 99;
            std::cout<<a<<std::endl;
        }
        
        // int a = 1009;  这样就肯定不行
    }
    ```

    java   编译出错

    ```java
    package variable;
    
    public class Test {
        public static void main(String[] args) {
            int a = 101;
            if(a > 100) {
                int a = 99; //Duplicate local variable a
                System.out.println(a);
            }
        }
    }
    ```

    go   正常

    ```Go
    package main
    
    import "fmt"
    
    func main() {
    	test := 101
    	if test > 100 {
    		//var test = 101
    		test := 100
    		fmt.Print(test)
    	}
        
        // test := 1009;  这样也肯定不行
    }
    ```

    

### 常量

```go
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
```

- `iota` 关键字的使用，只能在const里出现

### 流程控制语句

- if 语句

  - 普通的`if - else if - else` 不用加括号就是

  - 带有初始化语句的 if 语句,在条件判断之前，可以执行一个初始化语句。

    ```go
    package main
    
    import "fmt"
    
    func main() {
        // 基本的 if 语句
        x := 10
        if x > 5 {
            fmt.Println("x 大于 5")
        }
    
        // 带有 else 分支的 if 语句
        y := 3
        if y > 5 {
            fmt.Println("y 大于 5")
        } else {
            fmt.Println("y 不大于 5")
        }
    
        // 带有初始化语句的 if 语句
        if z := 7; z > 5 {
            fmt.Println("z 大于 5")
        }
    }
    ```

- 循环控制语句

  在 Go 语言中，有两种主要的循环控制语句：`for` 循环和 `range` 循环。

  - 普通的for循环语句

    `for` 循环是 Go 语言中基本的循环结构，它有三种形式：

    **a. 基本形式：**

    ```go
    for 初始化语句; 条件表达式; 后置语句 {
        // 循环体
    }
    ```

    ```go
    for i := 0; i < 5; i++ {
        // 循环体
    }
    ```

    **b. 类似 while 循环：**

    ```go
    for 条件表达式 {
        // 循环体
    }
    ```

    ```go
    count := 0
    for count < 5 {
        // 循环体
        count++
    }
    ```

    **c. 无限循环：**

    ```go
    for {
        // 无限循环体
    }
    ```

    ```go
    for {
        fmt.Println("无限循环")
    }
    ```

  - `range`关键字

    `range` 循环用于迭代数组、切片、字符串、map 等集合类型的元素。它返回两个值，第一个是索引（或键），第二个是元素的值。

    ```go
    for 索引, 值 := range 集合 {
        // 循环体
    }
    ```

    ```go
    numbers := []int{1, 2, 3, 4, 5}
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
    ```

### 函数

#### 普通函数

```go
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
	fmt.Printf("myfunc3 形参列表：%s %T,%s %T ",a,a,b,b)
	//r1 r2 作为两个局部变量，有默认值
	r1 = 666
	r2 = 777
	return 
}

//返回多个值，并给返回值命名，并且多个返回值的类型相同
func myfunc4(a string, b string) (r1 ,r2 int) {
	fmt.Printf("myfunc4 形参列表：%s %T,%s %T ",a,a,b,b)
	//r1 r2 作为两个局部变量，有默认值，在这里会返回0
	return 
}
```

- 可以给返回值指定名字，方便进行一些复杂操作后得到返回值
- 允许给返回值变量赋值后还是返回了其他值

#### init函数

```go
package main
/*搞不明白  它会从goroot去找，所以我得把go_code 复制到std( (D:\sdk_golang\go1.21.5\src\)里，否则报错：
PS D:\go-project\src\go_code\go-learning\init_> go run main.go
main.go:5:2: package go_code/go-learning/init_/lib1 is not in std (D:\sdk_golang\go1.21.5\src\go_code\go-learning\init_\lib1)
main.go:6:2: package go_code/go-learning/init_/lib2 is not in std (D:\sdk_golang\go1.21.5\src\go_code\go-learning\init_\lib2)
*/
import (
	"fmt"
	"go_code/go-learning/init_/lib1"
	"go_code/go-learning/init_/lib2"
)

func main() {
	lib1.Lab1Test()
	lib2.Lab2Test()
}

func init() {
	fmt.Println("main init...")
}

```

golang里面有两个保留的函数：init函数（能够应用于所有的package）和main函数（只能应用于package main）。这两个函数在定义时不能有任何的参数和返回值。

- init 函数可在package main中，可在其他package中，可在同一个package中出现多次。

- main 函数只能在package main中。

- 虽然一个package里面可以写任意多个init函数，但这无论是对于可读性还是以后的可维护性来说，我们都强烈建议用户在一个package中每个文件只写一个init函数。

  go程序会自动调用init()和main()，所以你不需要在任何地方调用这两个函数。每个package中的init函数都是可选的，但package main就必须包含一个main函数。

  程序的初始化和执行都起始于main包。

  如果main包还导入了其它的包，那么就会在编译时将它们依次导入。有时一个包会被多个包同时导入，那么它只会被导入一次（例如很多包可能都会用到fmt包，但它只会被导入一次，因为没有必要导入多次）。

  当一个包被导入时，如果该包还导入了其它的包，那么会先将其它包导入进来，然后再对这些包中的包级常量和变量进行初始化，接着执行init函数（如果有的话），依次类推。

  等所有被导入的包都加载完毕了，就会开始对main包中的包级常量和变量进行初始化，然后执行main包中的init函数（如果存在的话），最后执行main函数。下图详细地解释了整个执行过程：

  ![image-20240121153033787](C:\Users\chchi\OneDrive\文档\Markdown\编程学习笔记\Golang\Golang.assets\image-20240121153033787.png)

#### import 导包

- 给包取一个别名，而且是匿名，无法使用当前包的方法，但是会执行当前包内部的`init()`方法

```go
import (
    "fmt"
    _ "GolangTraining/InitLib1"
    _ "GolangTraining/InitLib2"
)
```

- 给包取别名 可以通过`myLib2.Println()`来直接调用

```go
import (
    myLib2 "fmt"
     "GolangTraining/InitLib2"
)
```

- 导入包中的全部方法，导入到当前本包的作用域中，fmt包中的全部方法可以直接使用API来调用,不需要通过`fmt.Println()`

```go
import (
    . "fmt"
)
```

#### defer关键字

defer语句被用于预定对一个函数的调用。可以把这类被defer语句调用的函数称为延迟函数。

defer作用：

- 释放占用的资源
- 捕捉处理异常
- 输出日志

```go
package main

import "fmt"

func main() {
	defer fmt.Println("main end1") 
    //在当前函数生命周期结束之后会被调用，有点类似析构函数,defer后面跟的就是让这个函数（defer所在的那里）
    //结束之前做的事，有多个defer以栈的形式，先进后出
	defer fmt.Println("main end2") 

	fmt.Println("main::hello go1")
	fmt.Println("main::hello go2")
}
```

return会先与defer执行：

```go
package main

import "fmt"

func return_func() int{
	fmt.Println("return end.")
	return 1
}

func defer_func() {
	fmt.Println("defer end.")
}

func test() int {
	defer defer_func()
	return return_func() 
}
func main()  {
	test()
}

//out:
return end.
defer end.
```

### slice 切片

#### 数组

```go
package main

import "fmt"

func main() {
	//定义数组

	//固定数组长度
	var myArray [4]int
	for i, v := range myArray {
		v = i + 1
		fmt.Println("index:",i,"value:",v)
	}

	fmt.Println("========================")

	myArray2  := [5]int {1,2,3}
	var myArray3 = [5]int {0,1,2,3,4}; //最好不要这样写
	printArray(myArray2)
	fmt.Println("========================")
	printArray(myArray3)
}

//将固定数组作为参数传递，类型要匹配，这里就接收不了[4]int
//而且传递固定数组是值拷贝
func printArray(myArray [5]int) {
	for i, v := range myArray {
		fmt.Println("index:",i,"value:",v)
	}
}
```

- 固定⻓度的数组在传参的时候， 是严格匹配数组类型的
- 固定数组作为形参传递是值拷贝

#### 动态数组 (slice)

```go
package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4}
	printSlice(array)
}

func printSlice(array []int) {
	for _, value := range array {
		fmt.Println(value)
	}
}
```

- 动态数组在传参上是引⽤传递，⽽且不同元素⻓度的动态数组他们的形参是⼀致

**声明切片的方式：**

```go
package main

import "fmt"

func main() {
	//声明slice是一个切片，并且初始化
	slice1 := []int{1, 2, 3, 4}
	printSlice(slice1)

	//声明切片但没有分配空间
	var slice2 []int
	//fmt.Println(slice2[0]) //数组越界错误

	var slice3 []int = make([]int, 3)
	fmt.Println(slice3[0]) //分配空间
 
	//省略var  最常用
	slice4 := make([]int,3)
	fmt.Println(slice4[0])

	//判断空切片
	if slice2 == nil {
		fmt.Println("slice2是一个空切片")
	} else {
		fmt.Println("slice2有空间")
	}
}

func printSlice(array []int) {
	for _, value := range array {
		fmt.Print(value," ")
	}
	fmt.Print("length :",len(array))
}
```
**切片的基本操作**
- 扩容 append
  ![slice示意图](image.png)
  - 利用append()函数在slice末尾追加元素。append会返回一个新的slice
  ```go
  func main() {
	slice := make([]int,3, 3)
	fmt.Println("slice长度：",len(slice),"slice容量：",cap(slice))

	//在slice末尾追加一个元素
	slice = append(slice, 1)
	//fmt.Println(append(slice, 1))
	fmt.Println("slice长度：",len(slice),"slice容量：",cap(slice))

	//out；
	// slice长度： 3 slice容量： 3
	// slice长度： 4 slice容量： 6
	}
	```
	- 当slice的容量满的时候，底层会自动开辟空间，新开辟的大小就是上次的容量，相当于扩容一倍
- 截取
  ```go
  func main() {
	slice := []int{0, 1, 2, 3, 4}
	//[lower-bound:upper-bound)   这样新切片的长度 =  upper-bound - lower-bound
	part1 := slice[1:3]
	fmt.Println(part1)

	//[0,2)
	part2 := slice[:3]
	fmt.Println(part2)

	//[1,len(slice))
	part3 := slice[1:]
	fmt.Println(part3)

	//[0,len(slice)) 截取全部
	part4 := slice[:]
	fmt.Println(part4)
	
	//这样对slice进行截取后，part和原来的slice还是指向同一块内存空间，只是它们维护的头尾指针不同
	//对slice[0]进行修改 , part[0]也会跟着改变
	slice[0] = 999
	fmt.Println(slice)
	fmt.Println(part2)

	//要实现slice的深拷贝，使用copy(new,old)
	newSlice := make([]int,len(slice))
	copy(newSlice,part3)
	fmt.Println(slice)
	fmt.Println(newSlice)
	}
	```
	- 对slice进行截取后，part和原来的slice还是指向同一块内存空间，只是它们维护的头尾指针不同
	- 要实现slice的深拷贝，使用copy(new,old)

### map

#### 定义map
```go
func main() {
	//var myMap map[string]string = make(map[string]string, 10)
	var myMap map[string]string
	fmt.Print(myMap == nil,"\n")

	//var m = make(map[string]string, 10)
	//也可以不指定大小   
	m := make(map[string]string) 
	m["java"] = "wife"
	m["go"] = "lover"
	fmt.Println(m)

	//声明的同时顺便初始化值，就不需要make分配空间了
	m3 := map[string]string {
		"wife: " : "java",
		"lover" : "go",
		"anotherLover" : "python",
	}

	fmt.Println(m3)
}
```
#### 使用map进行基础的增删改查
```go
func main() {
	cityMap := make(map[string]string)
	fmt.Println(cityMap)

	//添加
	cityMap["China"] = "Bejing"
	cityMap["Japan"] = "Tokoy"
	fmt.Println(cityMap)

	//遍历
	for key , value := range cityMap {
		fmt.Println("key",key,":","value",value)
	}

	//删除 delete关键字  传入key
	delete(cityMap,"Japan")
	fmt.Println(cityMap)

	//修改
	cityMap["China"] = "HuiLai"
	printMap(cityMap)

	//拷贝map
	m4 := make(map[string]string)
	for key, value := range cityMap {
		m4[key] = value
	}
	fmt.Println("\n",m4)
}

//使用map进行传参，引用传递
func printMap(cityMap map[string]string) {
	for key , value := range cityMap {
		fmt.Print("key = ",key," value = ",value)
	}
}
```
- 使用map进行传参，是引用传递
- 拷贝map可以通过遍历逐一赋值到新的map当中

### 面向对象
#### 类
- 定义类：通过struct + func相互绑定来定义类
  - struct的定义：
	```go
	//type 关键字  给数据类型取一个别名
	type myInt int

	//定义结构体
	type Book struct {
		title string
		price myInt 
		}
	```
  	将结构体当作形参进行传递时，有值传递和引用传递两种方式
	```go
	//结构体作形参
	//值传递
	func changeBook1(book Book) {
		book.title = "Golang"
	}
	//引用传递
	func changeBook2(book *Book) {
		book.title = "Golang"
	}   
	```
  - 与func绑定：
	```go
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
	//要加个*,才是指向该对象
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
	}
	```
  - go中的对象封装是基于包进行封装的。而类的访问权限是通过结构体，字段或者方法的首字母是否大小写来决定的。
	- 大写：对包外部可见，其它包也可以访问
	- 小写：对包外部不可见，只能在本包访问
#### 继承
定义父类和子类的基本语法：
```go
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
```
声明子类：
1. `s := Student{Person{name:"cjp",age : 20},1001}`
2. 也可以这样写:
 ```go
	var s Student
	s.name = "cjp"
	s.age = 20
	s.id = 1001
```
- go当中的继承没有那么多权限

#### 多态 interface
**go语言通过接口实现多态的特性。**
- 定义一个接口（本质是一个指针）
	```go
	//接口，本质是一个指针
	type Animal interface{
		Sleep()
		GetColor() string
		GetType() string
	}
	```

- 定义类实现该接口（重写接口的所有方法）
	```go
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
	```
- 这样就完成了多态的设计，可以使用接口类型的指针指向实现改接口的类的对象,通过动态绑定机制完成不同的功能
	```go
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
	```

**通用万能类型 : 空接口 interface{}**
- 有点像java里的Object类。可以使用`inteface{}`类型接收任意类型的变量
	```go
	package main

	import "fmt"

	func myFunc(arg interface{}) {
		fmt.Println("myFunc arg:",arg)
	}

	func main() {
		myFunc(666)
		myFunc("cjp")
	}
	```
- go提供了**类型断言机制** 用来判断引用的底层数据类型（arg的具体类型）
  ```go
	func myFunc(arg interface{}) {
		fmt.Println("myFunc arg:",arg)
		value , ok := arg.(string)
		if ok {
			fmt.Println("arg:",value,"的类型是string")
		}
	}
  ```


**go语言当中变量组成**
![Alt text](image-1.png)
每个变量都由这两部分组成，可以表示为一个pair数据结构:`pair<type,value>`。将变量赋值给另一个变量(可以相互兼容转换吗，比如string -> interface{})的过程中，这个pari是不可变的。
```go
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
```
### 反射