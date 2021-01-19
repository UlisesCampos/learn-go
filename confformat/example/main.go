package main

import "github.com/Neil-uli/go-play/confformat"

func main() {
	if err := confformat.MarshalAll(); err != nil {
		panic(err)
	}
	if err := confformat.UnmarshalAll(); err != nil {
		panic(err)
	}
	if err := confformat.OtherJSONExamples(); err != nil {
		panic(err)
	}
}
