//go:build ignore

package main

import (
	"fmt"
)

func fib(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
func main() {
	n := 90
	//fmt.Println(fib(n))
	ch := make(chan int, n)
	fib(n, ch)
	for i := range ch {
		fmt.Println(i)
	}

}
