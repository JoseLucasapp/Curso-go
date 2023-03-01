package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela": 15,
			"Guga":     84,
		},
		"J": {
			"José": 11,
		},
		"P": {
			"Pedro": 12,
		},
	}

	for letra, funcs := range funcsPorLetra {
		fmt.Printf(letra, funcs)
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Printf(letra, funcs)
	}
}
