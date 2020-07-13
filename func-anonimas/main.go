package main

import "fmt"

var calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo 5 + 7 = %d \n ", calculo(5, 7))

	calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("Restamos 6 - 4 = %d \n ", calculo(6, 4))

	operaciones()

	//Closures
	tablede1 := 2
	mitabla := tabla(tablede1)
	for i := 1; i < 11; i++ {
		fmt.Println(mitabla())
	}

}

func operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Println(resultado())
}

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}
