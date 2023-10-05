package main

import (
	"fmt"
	"log"
	"time"
)

func demo() {
	defer trace("demo")()
	log.Println("surrounding function")
	time.Sleep(2 * time.Second)

}
func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}
func triple(x int) (result int) {

	defer func() { result += x }()
	fmt.Println("执行defer后的result:", result)
	return double(x)
}

func main() {
	demo()
	fmt.Println(triple(4))
}

func trace(msg string) func() {
	start := time.Now()
	log.Println("开始执行trace")
	return func() { log.Printf("结束执行trace %s (%s)", msg, time.Since(start)) }
}
