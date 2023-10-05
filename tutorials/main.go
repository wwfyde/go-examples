package main

import (
	"fmt"
	"github.com/wwfyde/go-examples/tutorials/demo_package"
)

func main() {
	fmt.Println("Hello, 世界!")
	hello := demo_package.Hello(12)
	fmt.Println("hello:", hello)
}
