package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	Balance should be always be zero due to safe concurrency
*/

func main() {
	var balance int
	var wg sync.WaitGroup
	var mutex sync.Mutex

	deposit := func(amount int) {
		mutex.Lock()
		balance += amount
		mutex.Unlock()
	}

	withdraw := func(amount int) {
		mutex.Lock()
		defer mutex.Unlock()
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
