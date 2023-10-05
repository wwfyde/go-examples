package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(time.Since(now).Seconds())

}
