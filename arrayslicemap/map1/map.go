package main

import (
	"fmt"
)

func main() {
	//var aprovados map[int]string
	aprovados := make(map[int]string)

	aprovados[123456789] = "Maria"
	aprovados[111111111] = "Ana"
	aprovados[222222222] = "Pedro"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[123456789])
	delete(aprovados, 123456789)
	fmt.Println(aprovados[123456789])
}
