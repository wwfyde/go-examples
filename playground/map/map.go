package main

import "fmt"

func main() {
	args := make(map[string]int)
	args2 := map[string]int{
		"a": 1,
		"b": 2,
	}

	args["a"] = 1
	args["b"] = 2

	args2["c"] = 3
	_, _ = &args, args2["a"]
	names := make([]string, 0, 5)
	names = append(names, "a")
	names = append(names, "b")
	for index, name := range names {
		println(index, name)
	}

	fmt.Printf("%T %[1]v\n", args)
	fmt.Printf("%T %[1]v\n", args2)
	fmt.Printf("%T %[1]v\n", args2["d"])

	result := equal(map[string]int{"a": 1}, map[string]int{"a": 2})
	fmt.Println(result)
}

// 判断两个字典是否相等, 基本思路, 首先判断键数是否相等,
//
// 如果相等, 逐个判断x中的键在y中是否存在且值相等
// 如果每个值均存在且相等则说明x和y拥有相同的键值对
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
