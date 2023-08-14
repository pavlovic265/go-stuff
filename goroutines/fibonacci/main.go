package main

import "fmt"

func fibonacci(max int, ch chan int) {
	fib := make([]int, max)
	fib[0] = 0
	fib[1] = 1
	ch <- fib[0]
	ch <- fib[1]

	for i := 2; i < max; i++ {
		fib[i] = fib[i-1] + fib[i-2]
		ch <- fib[i]
	}

	close(ch)

	// return fib
}

func main() {

	// ch <- x - sending value to channel
	// x := <- ch - receiving value from channel
	// <- ch - discarding value from channel

	ch := make(chan int)
	go fibonacci(20, ch)
	for {
		// msg := <-ch
		// fmt.Printf("%d\n", msg)
		for msg := range ch {
			fmt.Printf("%d\n", msg)
		}
	}
}
