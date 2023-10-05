package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start count")
	tick := time.Tick(1 * time.Second)
	//var i = 0
	//for range time.Tick(1 * time.Second) {
	//	i++
	//	fmt.Printf("%v\n", i)
	//}
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	fmt.Println("countdown")

}
