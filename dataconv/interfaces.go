package dataconv

import "fmt"

// CheckType will print based on the interface type
func CheckType(s interface{}) {
	switch s.(type) {
	case string:
		fmt.Println("It's a string!")
	case int:
		fmt.Println("it's a number!")
	default:
		fmt.Println("not sure what it is..")
	}
}

// Interfaces demonstrates cating from anonymous interfaces to types
func Interfaces() {
	CheckType("test")
	CheckType(1)
	CheckType(false)

	var i interface{}
	i = "test"
	if val, ok := i.(int); ok {
		fmt.Println("val is", val)
	}
	// this one should fail
	if _, ok := i.(int); !ok {
		fmt.Println("glad we handled this")
	}
}
