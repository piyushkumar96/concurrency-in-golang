package main

import (
	"fmt"
	"sync"
)

var sharedResc1 map[string]string

func main() {
	sharedResc1 = make(map[string]string)
	var wg sync.WaitGroup

	mu := sync.Mutex{}
	cond := sync.NewCond(&mu)

	wg.Add(1)

	go func() {
		defer wg.Done()

		cond.L.Lock()
		for len(sharedResc1) < 1 {
			cond.Wait()
		}
		fmt.Println("First name is ", sharedResc1["Firstname"])
		cond.L.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		cond.L.Lock()
		for len(sharedResc1) < 2 {
			cond.Wait()
		}
		fmt.Println("Last name is ", sharedResc1["Lastname"])
		cond.L.Unlock()
	}()

	cond.L.Lock()
	sharedResc1["Firstname"] = "Piyush"
	sharedResc1["Lastname"] = "Kumar"
	cond.Broadcast()
	cond.L.Unlock()

	wg.Wait()
}
