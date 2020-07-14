package main

import "fmt"

/*Son interceptores que permiten ejecutar instrucciones comunes a varias funciones que reciben
y devuelven los mismos tipos de variables*/

var result int

func main() {
	fmt.Println("Inicio")

	result = operacion(sumar)(2, 9)
	fmt.Println(result)
	result = operacion(restar)(7, 9)
	fmt.Println(result)
	result = operacion(multiplicar)(10, 9)
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}
func restar(a, b int) int {
	return a - b
}
func multiplicar(a, b int) int {
	return a * b
}

func operacion(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("inicio de operacion")
		return f(a, b)
	}
}
