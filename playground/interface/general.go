// 通用结构通常适用于类型参数
package main

import "fmt"

type Float interface {
	~float32 | ~float64
	error
}

func Print[T Float](s T) {
	print(s)

}

type demo float64

func (demo) Error() string {
	return "demo error"
}
func main() {
	var b demo = 1.2
	Print(b)
	fmt.Println()
	println(b.Error())
}
