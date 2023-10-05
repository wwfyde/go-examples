package main

import (
	"log"
	"time"
)

// var c = make(chan int, 10)
var a string

func f(s string) {
	log.Println(s, "运行子协程")
	a = "hello, world"
	log.Println(s, "发送之前")
	time.Sleep(time.Second)
	log.Println(s, "发送之后")

}

func main() {
	go f("go a:")
	log.Println("接收之前")
	//<-c
	log.Println("接收之后")
	log.Println(a)
	f("b:")
	log.Println("主协程")
	time.Sleep(5 * time.Second)
	log.Fatal()
}
