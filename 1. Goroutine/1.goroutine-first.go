package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("No of goroutines before:- ", runtime.NumGoroutine())
	var data int

	go func() {
		data++
	}()
	go func() {
		data++
	}()

	fmt.Println("No of goroutines after:- ", runtime.NumGoroutine())
	if data == 0 {
		fmt.Println("The value of data is ", data)
	}
}
