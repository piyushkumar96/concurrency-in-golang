package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg [2]sync.WaitGroup

	//
	// here the final value of i will be printed when go routine executes
	for i := 1; i <= 3; i++ {
		wg[0].Add(1)

		go func() {
			defer wg[0].Done()
			fmt.Println("The value of i: ", i)
		}()
	}
	wg[0].Wait()

	// here the current value of i will be printed when go routine executes
	for i := 1; i <= 3; i++ {
		wg[1].Add(1)

		go func(i int) {
			defer wg[1].Done()
			fmt.Println("The value of i: ", i)
		}(i)
	}
	wg[1].Wait()

	fmt.Println("Done")
}
