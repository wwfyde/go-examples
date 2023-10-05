package main

import (
	"fmt"
)

func main() {
	var z float64
	v := "获取字符\n串的字符长度\n"
	w := []rune(v)
	s := `wwfyde\n吴方圆`
	const GoUsage = `Go is a tool for managing Go source code.

Usage:
	go command [arguments]
`
	fmt.Println(v, w, len(w), s)
	fmt.Print(GoUsage)
	// 获取字符串的字符长度 [33719 21462 23383 31526 20018 30340 23383 31526 38271 24230] 10
	//for index, value := range w {
	//	fmt.Printf("%d, %c\n", index, value)
	//}
	fmt.Println(z, -z, 1/z, -1/z, z/z) // +1.797693e+308+3.402823e+0380 -0 +Inf -Inf NaN
	//nan := math.NaN()
	//fmt.Println(nan == nan, nan < nan, nan > nan)
}
