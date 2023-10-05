package main

import "fmt"

func main() {
	s := "世界和平"
	fmt.Printf("% x\n", s) // e4 b8 96 e7 95 8c e5 92 8c e5 b9 b3
	r := []rune(s)
	fmt.Printf("%x\n", r)        // [4e16 754c 548c 5e73]
	fmt.Println(string(r))       // 世界和平
	fmt.Println(string(65))      // A 而不是65
	fmt.Println(string(0x4e16))  // 世
	fmt.Println(string(1234567)) // 非法符号: � = \uFFFD
}
