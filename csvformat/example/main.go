package main

import (
	"fmt"

	"github.com/Neil-uli/go-play/csvformat"
)

func main() {
	if err := csvformat.AddMoviesFromText(); err != nil {
		panic(err)
	}
	if err := csvformat.WriteCSVOutput(); err != nil {
		panic(err)
	}
	buffer, err := csvformat.WiteCSVBuffer()
	if err != nil {
		panic(err)
	}
	fmt.Println("Buffer = ", buffer.String())
}
