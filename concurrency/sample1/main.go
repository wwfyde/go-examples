package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			fmt.Println("hello, ", i)
		}(i)
	}
	wg.Wait()
}
