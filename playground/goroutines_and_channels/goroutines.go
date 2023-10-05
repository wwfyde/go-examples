package main

import (
	"fmt"
	"time"
)

func say(s string, done chan struct{}) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})
	go say("<world", done)
	//say(">hello")
	// 开启一个新的线程, 与主线程并行运行,
	//time.Sleep(5 * time.Second)
	<-done // 阻塞主线程直到协程执行完毕
	print("end...")
}
