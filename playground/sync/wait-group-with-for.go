// WaitGroup用于等待协程执行完毕

//go:build ignore

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	//// Do Something
	//for {
	//	select {
	//	case <-ctx.Done():
	//		fmt.Println("Canceled!")
	//		return
	//	default:
	//		fmt.Println("Do Some Work!")
	//		time.Sleep(100 * time.Millisecond)
	//	}
	//
	//}
	fmt.Println("hello, 世界!")
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg)

	}

	wg.Wait()
	fmt.Println("main goroutine exit!")

}
