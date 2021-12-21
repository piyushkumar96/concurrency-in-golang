package main

import (
	"fmt"
	"runtime"
	"sync"
)

func generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, v := range nums {
			out <- v
		}
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num * num
		}
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int)
	output := func(c <-chan int) {
		for num := range c {
			out <- num
		}
		defer wg.Done()
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

func print(in <-chan int) {
	for num := range in {
		fmt.Println(num, " ")
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	nums1 := []int{1, 5, 3, 7, 9}
	nums2 := []int{2, 6, 4, 10, 8}

	ch1 := square(generator(nums1...))
	ch2 := square(generator(nums2...))

	print(merge(ch1, ch2))
}
