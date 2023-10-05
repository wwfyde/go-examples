package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a string = "ww" // 声明字符串变量
	p5 := new(int)
	fmt.Println("p5 的类型", reflect.TypeOf(p5), p5, *p5)
	b := 2      // 声明整型变量
	p := &a     // 短声明指针
	var p2 = &b // 声明整型指正
	var p3 *int // 声明整型指正
	p4 := &b
	fmt.Println(p2 == p4, p3 == nil)
	if p3 == nil {
		fmt.Println("p3 的空值是nil")
	} else {
		fmt.Println("p3的空值不是nil")
	}
	*p = "吴方圆"
	*p2 = 123
	fmt.Println("变量a的类型和值", reflect.TypeOf(a), ":", a)
	fmt.Println("指针p", reflect.TypeOf(p), ":", p, *p)
	fmt.Println("指针p2", reflect.TypeOf(p2), ":", p2, *p2)
	fmt.Println("表达式*p", reflect.TypeOf(*p), ":", *p)
	q := 2
	demo(&q)
	fmt.Println("局部变量地址", f() == f())

	// param compare

	v1, v2 := 1, 1
	fmt.Println("指针传递v1:=1, 传递的是变量", PointerParam(&v1), v1)
	fmt.Println("变量传递v2:=1, 传递的是值", NormalParam(v2), v2)

}

func demo(p *int) int {

	//b := 2.3
	*p = 99
	fmt.Println("小数初始化类型:", reflect.TypeOf(p), p, *p)
	return 0
}

func f() *int {

	v := 12

	return &v
}

func PointerParam(p *int) int {
	*p++
	return *p
}

func NormalParam(p int) int {
	p++
	return p
}
