package main

import "fmt"


type animal struct {
	Name string
	canFly bool
}

func main() {
	anAnimal := animal{Name: "Lion", canFly: false}
	fmt.Println(anAnimal)
	fmt.Println(anAnimal.Name)
	// Pointers 
    aLionptr := &anAnimal
    fmt.Println(aLionptr.canFly)
}
