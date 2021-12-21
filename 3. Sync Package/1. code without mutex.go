package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	Balance should be always be zero but due to unsafe concurrency it's not
*/

func main() {
	var balance int
	var wg sync.WaitGroup

	deposit := func(amount int) {
		balance += amount
	}

	withdraw := func(amount int) {
		balance -= amount
	}
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			withdraw(1)
		}()
	}
	wg.Wait()

	fmt.Println("Balance is ", balance)
}
