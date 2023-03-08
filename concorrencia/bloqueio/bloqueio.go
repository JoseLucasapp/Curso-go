package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1
	fmt.Println("Só dps q foi lido")
}

func main() {
	c := make(chan int)

	go rotina(c)
	fmt.Println(<-c)
	fmt.Println("Lido")
	fmt.Println(<-c)
	fmt.Println("Fim")
}
