package dataconv

import "fmt"

// ShowConv demonstrates some type conversion
func ShowConv() {
	var a = 24
	// float
	var b = 2.54

	c := float64(a) * b
	fmt.Println(c)

	// fmt.Sprintf is a good way to convert to strings
	presision := fmt.Sprintf("%.2f", b)
	fmt.Printf("%s - %T\n", presision, presision)
}
