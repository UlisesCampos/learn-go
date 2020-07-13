package main

import "fmt"

func main() {
	paises := make(map[string]string, 5)

	paises["Mexico"] = "D.F"
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises)

	campeonato := map[string]int{
		"Barcelona":   9,
		"Real madrid": 15,
		"Chivas":      37,
		"Pumas":       50,
	}

	campeonato["River Plate"] = 25
	campeonato["Chivas"] = 20
	delete(campeonato, "Real madrid")
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s, tiene un puntaje de : %d \n", equipo, puntaje)
	}

	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)
}
