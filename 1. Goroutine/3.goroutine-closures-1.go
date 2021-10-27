package main

import (
	"fmt"
	"sync"
)

func main() {
	var i int
	var wg sync.WaitGroup

	// here in go even this below function returns, it has access to i variable as before
	// exiting the function scope, go runtime copies this i variable to heap
	increment := func(w *sync.WaitGroup) {
		w.Add(1)

		go func() {
			defer w.Done()
			i++
			fmt.Println("The value of i: ", i)
		}()
		fmt.Println("return from function")
		return
	}

	increment(&wg)
	wg.Wait()

	fmt.Println("Done")
}
