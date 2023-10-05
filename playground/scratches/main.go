package main

import (
	"time"
)

var a string

func hello() {
	ch := make(chan struct{})
	go func() { a = "hello"; ch <- struct{}{} }()
	<-ch
	print(a)
	time.Sleep(time.Second)
	print(a)
}
func hello2() {
	go func() { a = "hello"; time.Sleep(time.Second) }()
	time.Sleep(time.Second)
	print(a)
	time.Sleep(time.Second)
	print(a)
}
func hello3() {
	go func() { a = "hello" }()
	print(a)
}
func main() {
	//hello()
	//print(100)
	//hello2()
	hello3()
}
