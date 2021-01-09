package main

import templates "github.com/Neil-uli/go-play/templates-examples/example2"

func main() {
	if err := templates.RunTemplate(); err != nil {
		panic(err)
	}
	if err := templates.InitTemplates(); err != nil {
		panic(err)
	}
	if err := templates.HTMLDifferences(); err != nil {
		panic(err)
	}
}
