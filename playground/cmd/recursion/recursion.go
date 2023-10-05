package main

import (
	"flag"
	"fmt"
	"github.com/wwfyde/go-examples/playground/recursion"
	"os"
	"strconv"
)

func main() {
	flag.Int("n", 0, "fib")
	n, _ := strconv.Atoi(os.Args[1])
	fmt.Println(recursion.Fib(n))
	fmt.Println(recursion.FibList(n))
	fmt.Println(recursion.FibList2(n))
}
