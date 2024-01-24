package main
import "fmt"

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	Fibonacci(c,quit)
}

func Fibonacci(c , quit chan int) {
	x , y := 1 , 1
	for {
		select {
		case c <- x :
			//如果c可写，则会执行该case
			x = y
			y = x + y
		case <- quit :
			//如果quit可读，则会执行该case
			fmt.Println("quit")
			return 
		}
	}
} 