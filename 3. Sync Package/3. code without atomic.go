package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var counter int32
	var wg sync.WaitGroup

	wg.Add(50)
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&counter, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("Counter :- ", counter)
}
