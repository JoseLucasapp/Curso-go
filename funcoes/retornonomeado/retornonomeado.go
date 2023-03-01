package main

import "fmt"

func trocar(p1, p2 int) (seg int, prim int) {
	seg = p2
	prim = p1

	return
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)

	segundo, primeiro := trocar(7, 1)

	fmt.Println(segundo, primeiro)
}
