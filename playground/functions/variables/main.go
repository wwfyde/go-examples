package main

import (
	"fmt"
	"strings"
)

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func main() {
	f := square
	fmt.Println(f(3)) // "9"
	f = negative
	fmt.Println(f(3))     // "-3
	fmt.Printf("%T\n", f) // "func(int) int"

	//f = product // "Cannot use 'product' (type func(m int, n int) int) as the type func(n int) int"
	{

		//var f func(int) int

		//f(3) // 调用空函数 引发宕机
	}
	{
		var f func(int) int

		if f != nil {
			f(3)

		}
	}

	fmt.Println(strings.Map(add1, "HAL-9000")) // IBM.:111

	fmt.Println(strings.Map(add1, "VMS")) // 	WNT
}
func add1(r rune) rune { return r + 1 }
