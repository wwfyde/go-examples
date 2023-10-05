package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	time.Sleep(time.Second)
	//select {
	//case res := <-ch:
	//	println(res)
	//default:
	//	println("finished!")
	//}
	//wg.Add(1)

	go func() {
		wg.Add(1)
		defer wg.Done()
		select {
		case <-ch:
			fmt.Println("Received!")
		}
	}()
	ch <- 2

	wg.Wait()

}
