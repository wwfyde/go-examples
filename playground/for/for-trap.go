package main

import "fmt"

func main() {
	done := make(chan bool)
	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func(v string) {
			fmt.Println(v)
			// block goroutine
			done <- true
		}(v)
		//fmt.Println(v)
	}

	for _ = range values {
		<-done
	}

}
