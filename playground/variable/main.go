package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 变量初始化后的内存空间
	//var a int32 = 99
	b := 12
	c := "abcdefghdighh"
	e := map[string]string{"a": "1", "b": "2", "c": "3", "d": "4"}
	var g *string
	var h uintptr
	d := "a"
	f := struct{ a, b, c, d, e int }{1, 2, 4, 5, 18}
	fmt.Println(unsafe.Sizeof(f))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(new(int)))
	fmt.Println(unsafe.Sizeof(&c), len(c))
	fmt.Println(unsafe.Sizeof(""), len(d))

	fmt.Println("pointer size:", unsafe.Sizeof(g))
	fmt.Println("pointer size:", unsafe.Sizeof(h))
	fmt.Println("map size:", unsafe.Sizeof(e))
}
