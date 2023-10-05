package main

import "fmt"

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)    // 00100010, 集合{1, 5}
	fmt.Printf("%08b\n", y)    // 00000110, 集合{1, 2}
	fmt.Printf("%08b\n", x&y)  // 00000010, 交集{1}
	fmt.Printf("%08b\n", x|y)  // 00100110, 并集{1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // 00100100, 对称差{2, 5}
	fmt.Printf("%08b\n", x&^y) // 00100000, 差集{5}

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i) // 1, 5
		}
	}
	fmt.Printf("%08b\n", x<<1) // 01000100  {2, 6}
	fmt.Printf("%08b\n", x>>1) // 00010001  {0, 4}
	fmt.Printf("%08b\n", x)    // 00100010  {1, 5}
	fmt.Printf("%08b\n", ^x)   // 11011101
}
