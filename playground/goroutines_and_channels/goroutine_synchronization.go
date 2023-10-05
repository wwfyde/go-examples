package main

import (
	"log"
	"time"
)

var c = make(chan int, 10)
var a string

func f() {
	log.Println("goroutine started")
	a = "hello, world"
	time.Sleep(time.Second)
	log.Println("async task finished and send result")
	c <- 0
	log.Println("send successful, and goroutine continuing")
	time.Sleep(time.Second)
	log.Println("goroutine end")

}

func main() {
	log.Println("main stated")
	go f()
	log.Println("goroutine created")

	time.Sleep(2 * time.Second)

	<-c
	log.Println("main receive from goroutine: ", a)
	time.Sleep(5 * time.Second)
	log.Println("main exited")
}
