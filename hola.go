package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello ▼")
	a := naiveSum(10)
	fmt.Println(a)
	x := map[int]int{1:10,2:20}
	sw(x)
	message := "Hello world"
	go sayHello(message)
	time.Sleep(5*time.Second)
}

func naiveSum(n int) int {
	sum := 0
	for i := 0; i<n; i++ {
		sum += i
	}
	return sum
}
func sw(n map[int]int) {
	for k, v := range n {
		fmt.Println("key", k)
		fmt.Println("value", v)
		
	}
}
func sayHello(what string) {
	fmt.Println(what)
	
}
