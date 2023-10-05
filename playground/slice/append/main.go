package main

import "fmt"

func main() {
	var runes []rune
	for _, r := range "Hello, 世界!" {
		runes = append(runes, r)
	}
	fmt.Printf("%T %[1]q\n", runes)
	rune2 := []rune("Hello, 世界!")
	fmt.Printf("%T %[1]q\n", rune2)
	rawArray := [3]int{1, 2, 3}
	x := rawArray[:]
	var y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i, i, i+1)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
	fmt.Printf("%v\n", y)
}

// 这是函数注解
//
// 没什么好事
func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	fmt.Printf("...类型: %T %[1]v\n", y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z[len(x):], y)
	}
	return z
}
