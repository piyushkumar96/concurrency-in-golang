package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 5)

	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			fmt.Println("Sending i: ", i)
			ch <- i
		}
	}()

	fmt.Println("The values read from buffered channel ")
	for v := range ch {
		fmt.Println("Receiving i: ", v)
	}
}
