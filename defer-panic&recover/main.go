package main

import (
	"log"
)

func main() {
	/*archivo := "prueba.txt"
	f, err := os.Open(archivo)

	defer f.Close()

	if err != nil {
		fmt.Println("error abriendo el archivo")
		os.Exit(1)
	}*/

	ejemploPanic()
}
func ejemploPanic() {
	defer func() {
		recov := recover()
		if recov != nil {
			log.Fatalf("Ocurrio un error que genero panic \n %v", recov)
		}
	}()

	a := 1
	if a == 1 {
		panic("Se encontro el valor de 1")
	}
}
