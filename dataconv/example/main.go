package main

import (
	"github.com/Neil-uli/go-play/dataconv"
)

func main() {
	dataconv.ShowConv()
	if err := dataconv.Strconv(); err != nil {
		panic(err)
	}
	dataconv.Interfaces()
}
