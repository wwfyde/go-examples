package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)

	go func() {
		msg <- "ping"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println(<-msg)
	}()

	time.Sleep(3 * time.Second)
}
