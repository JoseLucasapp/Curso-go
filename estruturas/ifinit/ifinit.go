package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numaleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10)
}

func main() {
	if i := numaleatorio(); i > 5 {
		fmt.Println("Ganhou!!!", i)
	} else {
		fmt.Println("Perdeu!!!", i)
	}
}
