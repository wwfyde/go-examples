package main

import (
	"fmt"
	"time"
)

var (
	co = 0
	ma = 0
)

func newTask() {
	for {
		co += 1
		fmt.Printf("---concurrent task: %d\n", co)
		time.Sleep(time.Second)
	}
}
func main() {
	go newTask()

	for {
		ma += 1
		fmt.Printf("---main task: %d\n", ma)
		time.Sleep(time.Second)
	}

}
