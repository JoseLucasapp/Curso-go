package main

import (
	"fmt"
	"time"
)

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base
	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	channel := make(chan int)

	go doisTresQuatroVezes(2, channel)

	a, b := <-channel, <-channel

	fmt.Println(a, b)
	fmt.Println(<-channel)
}
