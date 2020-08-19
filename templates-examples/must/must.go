package main

import (
	"fmt"
	"text/template"
)

func main() {
	tOk := template.New("first")
	template.Must(tOk.Parse(" algún texto estático  /* y un comentario */"))
	fmt.Println("Analizis de la primera plantilla: OK.")

	template.Must(template.New("second").Parse("algún texto estático {{ .Name }}"))
	fmt.Println("Análisis de la segunda plantilla OK.")

	fmt.Println("La siguiente plantilla va a fallar.")
	tErr := template.New("Verificar el error de análisi con Must")
	template.Must(tErr.Parse(" some static text {{ .Name }"))
}
