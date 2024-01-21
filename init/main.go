package main

/*搞不明白  它会从goroot去找，所以我得把go_code 复制到std( (D:\sdk_golang\go1.21.5\src\)里，否则报错：
PS D:\go-project\src\go_code\go-learning\init_> go run main.go
main.go:5:2: package go_code/go-learning/init_/lib1 is not in std (D:\sdk_golang\go1.21.5\src\go_code\go-learning\init_\lib1)
main.go:6:2: package go_code/go-learning/init_/lib2 is not in std (D:\sdk_golang\go1.21.5\src\go_code\go-learning\init_\lib2)
*/
import (
	"fmt"
	// "go_code/go-learning/init_/lib1"
	// "go_code/go-learning/init_/lib2"
)

func main() {
	// lib1.Lab1Test()
	// lib2.Lab2Test()
}

func init() {
	fmt.Println("main init...")
}
