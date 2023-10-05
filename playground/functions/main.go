package main

import (
	"fmt"
	"github.com/wwfyde/go-examples/playground/functions/recursive"
)

func MultiReturn(a, b int) (c, d int) {
	c = a + b
	d = a * b
	return
}

func add(x int, y int) int   { return x + y }      // 返回表达式
func sub(x, y int) (z int)   { z = x - y; return } // 返回值命名
func first(x int, _ int) int { return x }          // 未使用参数
func zero(int, int) int      { return 0 }

func main() {
	c, d := MultiReturn(2, 5)
	println(c, d)
	fmt.Printf("%T\n", add)   // func(int, int) int
	fmt.Printf("%T\n", sub)   // func(int, int) int
	fmt.Printf("%T\n", first) // func(int, int) int
	fmt.Printf("%T\n", zero)  // func(int, int) int
	fmt.Println(recursive.Fib(6))
	f := recursive.Fib2()
	fmt.Println(f(), f(), f())
	fmt.Println(recursive.Fac(6))
	fmt.Println(recursive.Sum(5))

}
