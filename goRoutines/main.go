package main

import (
	"fmt"
	"strings"
	"time"
	"runtime"
)

func main() {
	go myName("Neil Ulises")
	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x)

	go say("world") //new goroutine
	say("hello") // current goroutine

}
func myName(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
func say(s string) {
for i := 0; i < 5; i++ {
	//tells the CPU to run other goroutines, and at some point to come back
	runtime.Gosched()
	fmt.Println(s)
	}
}

