package main

import (
	"fmt"
)

func getMsg(ch1 chan<- string) {
	ch1 <- "Hello, I'm Piyush Kumar"
}

func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	data := <-ch1
	ch2 <- data
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go getMsg(ch1)
	go relayMsg(ch1, ch2)

	fmt.Println("Getting message from channel 2:- ", <-ch2)
}
