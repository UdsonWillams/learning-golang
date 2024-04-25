package main

import (
	"fmt"
	"time"
)

func pingar(c chan string) {
	for {
		c <- "ping"
		time.Sleep(1000000000)
	}
}

func pongar(c chan string) {
	for {
		c <- "pong"
		time.Sleep(time.Second)
	}
}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
	}
}

func main() {
	var c chan string = make(chan string)
	go pingar(c)
	go pongar(c)
	go imprimir(c)

	var fim string
	fmt.Scanln(&fim)
}
