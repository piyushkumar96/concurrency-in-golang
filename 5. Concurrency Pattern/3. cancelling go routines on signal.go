package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func generator(done chan struct{}, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, v := range nums {
			select {
			case out <- v:
			case <-done:
				return
			}

		}
	}()

	return out
}

func square(done chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			select {
			case out <- num * num:
			case <-done:
				return
			}

		}
	}()

	return out
}

func merge(done chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int)
	output := func(c <-chan int) {
		defer wg.Done()
		for num := range c {
			select {
			case out <- num:
			case <-done:
				return
			}
		}

	}

	for _, c := range cs {
		wg.Add(1)
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func print(done chan struct{}, in <-chan int) {
	for num := range in {
		fmt.Println(num, " ")
		close(done)
		return
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	nums1 := []int{1, 5, 3, 7, 9}
	nums2 := []int{2, 6, 4, 10, 8}

	done := make(chan struct{})

	ch1 := square(done, generator(done, nums1...))
	ch2 := square(done, generator(done, nums2...))

	print(done, merge(done, ch1, ch2))

	time.Sleep(2 * time.Millisecond)

	fmt.Printf("The number of go routines running %d\n", runtime.NumGoroutine())
}
