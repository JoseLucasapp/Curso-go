package main

import (
	"fmt"
)

func display(array map[string]float64) {
	for sal, nome := range array {
		fmt.Printf("%v: $%v \n", sal, nome)
	}
	fmt.Println("")
}

func main() {
	funcsESalarios := map[string]float64{
		"José":  1200,
		"João":  2500,
		"Maria": 1000,
	}

	display(funcsESalarios)

	funcsESalarios["Paulo"] = 1350

	display(funcsESalarios)

}
