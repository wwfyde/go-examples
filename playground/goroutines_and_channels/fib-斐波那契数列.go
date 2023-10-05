package main

import (
	"fmt"
)

func fibonacci(n uint, c chan uint) {
	x, y := uint(0), uint(1)
	for i := uint(0); i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan uint, 100)
	go fibonacci(uint(cap(c)), c)
	for i := range c {
		fmt.Println(i)
	}
}
