package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var i int
	i, _ = strconv.Atoi(os.Args[1])
	fmt.Println("command line args:", i)
	switch {
	case i < 10:
		fmt.Println("i is less than 10")
		fallthrough
	case i < 50:
		fmt.Println("i is less than 50")
		fallthrough
	case i < 100:
		fmt.Println("i is less than 100")
	}

}
