package main

import (
	"fmt"
	"sync"
)

var sharedResc map[string]string

func main() {
	sharedResc = make(map[string]string)
	var wg sync.WaitGroup

	mu := sync.Mutex{}
	cond := sync.NewCond(&mu)

	wg.Add(1)

	go func() {
		defer wg.Done()

		cond.L.Lock()
		for len(sharedResc) == 0 {
			cond.Wait()
		}
		fmt.Println("Name is ", sharedResc["Name"])
		cond.L.Unlock()
	}()

	cond.L.Lock()
	sharedResc["Name"] = "Piyush"
	cond.Signal()
	cond.L.Unlock()

	wg.Wait()
}
