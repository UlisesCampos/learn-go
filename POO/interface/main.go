package main

import "fmt"

type Persona struct {
	Nombre   string
	Apellido string
	Edad     uint8
}
type Programador struct {
	Persona
	Especialidad       string
	LenguajesFavoritos []string
}
type Vecino interface {
	Saludar()
}

func (p Persona) Saludar() {
	fmt.Printf("Hola, mi nombre es %s %s y tengo %d años\n",
		p.Nombre, p.Apellido, p.Edad)
}
func (p *Persona) Cumpleanios() {
	p.Edad += 1
}
func (p Programador) beberCafe() {
	fmt.Println("¡Me siento vivo!")
}
func main() {
	p1 := Programador{
		Persona{"Orlando", "Monteverde", 26},
		"Desarrollo Web",
		[]string{"Go", "Python", "JavaScript"},
	}
	p2 := Persona{"Daniel", "Herrera", 32}
	p3 := Persona{"Carmen", "Salazar", 20}

	var vecino Vecino
	vecino = p1
	fmt.Println(vecino) // {{Orlando Monteverde 26} Desarrollo Web [Go Python JavaScript]}
	vecino = p2
	fmt.Println(vecino) // {Daniel Herrera 32}
	vecino = p3
	fmt.Println(vecino) // {Carmen Salazar 20}
}
