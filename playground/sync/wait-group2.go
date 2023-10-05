package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(i int, wg *sync.WaitGroup) {
	fmt.Println(rand.Intn(2))
	time.Sleep(time.Second)
	defer wg.Done()
	if i%2 == 0 {
		fmt.Println(fmt.Sprintf("i:%d", i))
	} else {
		fmt.Println("execute normally: ", i)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	//os.Exit(0)
	fmt.Println(" main exit!")
}
