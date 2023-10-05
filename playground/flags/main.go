package main

import (
	"flag"
	"fmt"
)

func main() {
	//	flag.Args()
	flag.String("s", "", "print strings to stdout")
	flag.Parse()
	fmt.Println("hello, 世界!")
}
