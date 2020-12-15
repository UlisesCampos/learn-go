package main

import (
	"fmt"
	"time"
)

func main() {
	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("Llegue hasta aqui")

	msg := <-canal1
	fmt.Println(msg)

	// Second channel
	a := []int{7,2,8,-9,4,0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
    	x, y := <-c, <-c  // receive c

    	fmt.Println(x, y, x + y)
	}

func bucle(canal1 chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 100000; i++ {

	}
	final := time.Now()
	canal1 <- final.Sub(inicio)
}
// Second channel
func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
	total += v
	}
	c <- total // send total to c
}
