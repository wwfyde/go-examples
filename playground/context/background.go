package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	go func(ctx context.Context) {
		// The returned context is done and closed,
		// if the parent context (in this case, Background) is closed.
		fmt.Println("start a goroutine")
		<-ctx.Done()
		fmt.Println("Function execution terminated.")
	}(ctx)

	// Now, let's simulate some work:
	fmt.Println("Doing some work...")
	// Time to sleep to simulate work
	time.Sleep(2 * time.Second)

	fmt.Println("Work is done.")
}
