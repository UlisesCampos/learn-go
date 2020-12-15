package main

import "fmt"

func one(c1 chan string) {
	for i := 0; i < 5; i++ {
		c1 <- "Channel One"
	}
	close(c1)
}
func two(c2 chan string) {
	for i := 0; i < 5; i++ {
		c2 <- "Channel two"
	}
	close(c2)
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go one(c1)
	go two(c2)
	for {
		select {
		case msg, ok := <-c1:
			fmt.Println(msg)
			if !ok {
				c1 = nil
			}
		case msg, ok := <-c2:
			fmt.Println(msg)
			if !ok {
				c2 = nil
			}
		}
		if c1 == nil && c2 == nil {
			break
		}
	}
}
