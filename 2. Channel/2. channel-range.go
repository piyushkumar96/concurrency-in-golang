package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Sending i: ", i)
			ch <- i
		}
		close(ch)
	}()

	fmt.Println("The values read from channel ")
	for v := range ch {
		fmt.Println("Receiving i: ", v)
	}
}
