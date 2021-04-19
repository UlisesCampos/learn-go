package math

import (
	"fmt"
	"math"
)

// Examples demonstrates some of the functions in the math package
func Examples() {
	// sqrt examples
	i := 25

	result := math.Sqrt(float64(i))
	fmt.Println(result) // 5

	// ceil rounds up
	result = math.Ceil(9.5)
	fmt.Println(result) //10

	// floor rounds down
	result = math.Floor(9.5)
	fmt.Println(result)

	fmt.Println("Pi: ", math.Pi, "E: ", math.E)
}
