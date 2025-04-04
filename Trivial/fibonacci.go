package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x,y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	start := time.Now()
	c := make(chan int, 90)
	go fibonacci(cap(c), c)
	var count int = 0
	for i := range c {
		// dont need to add spaces between, does fmt print does it for you
		fmt.Println("fib#", count, "=", i)
		count += 1
	}
	elapsed := time.Since(start)
	fmt.Println("Execution time:", elapsed)
}