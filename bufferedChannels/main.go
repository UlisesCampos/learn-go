package main

import "fmt"

func main() {
	c := make(chan int, 2) // if change 2 to 1 you will have a runtime error, but 3 will be fine
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)

	cha := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-cha)
		}
		quit <- 0
	}()
	fibonacci(cha, quit)
}

// We can use range to make the 'buffer channels' work as a list and a map
// Select is for listen more channels
func fibonacci(cha, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case cha <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
