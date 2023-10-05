package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}

	}()
	select {
	case <-time.After(5 * time.Second):
		//no
	case <-abort:
		fmt.Println("launch aborted")
		return
	}
	launch()
}

func launch() {
	fmt.Println("launch success")
}
