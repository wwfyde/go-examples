package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 4
	ch <- 5
	ch <- 6
	//ch <- 7
	a := len(ch)
	for i := 0; i < a; i++ {
		result := <-ch
		fmt.Println(result)
	}
}
