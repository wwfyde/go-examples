// WaitGroup用于等待协程执行完毕
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Canceled!")
				return
			default:
				fmt.Println("Do Some Work!")
				time.Sleep(100 * time.Millisecond)
			}

		}
	}()
	wg.Wait()

}
