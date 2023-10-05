package main

import (
	"fmt"
	"github.com/wwfyde/go-examples/playground/example"
	"github.com/wwfyde/go-examples/playground/files/read_file"
)

func main() {
	example.Hello()
	fmt.Println("hello, gin!")
	fmt.Println(read_file.Read())
}
