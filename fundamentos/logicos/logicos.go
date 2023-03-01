package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprartv50 := trab1 && trab2
	comprartv32 := trab1 != trab2
	comprarsorvete := trab1 || trab2

	return comprartv50, comprartv32, comprarsorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)

	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}
