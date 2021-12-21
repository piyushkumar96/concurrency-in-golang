package main

import (
	"fmt"
	"runtime"
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

func print(in <-chan int) {
	for num := range in {
		fmt.Println(num, " ")
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	nums := []int{1, 5, 3, 7, 8}

	fmt.Println("The modified numbers ")
	print(square(square(generator(nums...))))
}
