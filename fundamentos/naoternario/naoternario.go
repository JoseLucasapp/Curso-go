package main

import "fmt"

//n tem ternario

func obterresultado(nota float64) string {
	if nota >= 6 {
		return "aprovado"
	}

	return "Reprovado"
}

func main() {
	fmt.Println(obterresultado(6.5))
}
