package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串和数字相互转换

	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))             // 123 123
	fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"
}
