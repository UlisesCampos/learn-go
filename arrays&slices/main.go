package main

import "fmt"

var matriz [5][7]int

func main() {
	//Slice = vector dinamico
	matriz1 := []int{2, 4, 6}
	fmt.Println(matriz1)
	variant2()
	variant3()
	variant4()

	tabla := [10]int{5, 6, 8, 4, 97, 6, 2, 3}

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[3][5] = 1
	fmt.Println(matriz)
}

func variant2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[3:]

	fmt.Println(porcion)
}

func variant3() {
	elementos := make([]int, 10, 30)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}
func variant4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums))
}
