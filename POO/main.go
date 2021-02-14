package main

import (
	"fmt"
	"time"

	us "github.com/Neil-uli/go-play/POO/user"
)

/*
type usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}*/

type pepe struct {
	us.Usuario
}

func main() {
	u := new(pepe)
	u.AltaUsuario(1, "Pablo", time.Now(), true)
	fmt.Println(u.Usuario)
	fmt.Println(u.Id)
}
