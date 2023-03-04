package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro
	turboLigado bool
}

func main() {
	f := ferrari{}

	f.nome = "f40"
	f.velocidadeAtual = 122
	f.turboLigado = true

	if f.turboLigado {
		fmt.Printf("A ferrari %s está a %d KM/h com o turbo ligado.", f.nome, f.velocidadeAtual)
	} else {
		fmt.Printf("A ferrari %s está a %d KM/h com o turbo desligado.", f.nome, f.velocidadeAtual)
	}
}
