package main

import "sync"

func worker(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case res := <-ch:
		println(res)
		return
	}
}
func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	ch := make(chan int)
	go worker(ch, &wg)
	ch <- 1
	wg.Wait()

}
