package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func pong(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func imprimir(c1, c2 chan string) {
	for {
		ping := <-c1
		fmt.Println(ping)
		time.Sleep(time.Second * 1)

		pong := <-c2
		fmt.Println(pong)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go ping(c1)
	go pong(c2)
	go imprimir(c1, c2)
	var inside string
	fmt.Scanln(&inside)
}
