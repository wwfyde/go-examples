package main

import (
	"fmt"
	"reflect"
)

func readData(datas ...any) {
	fmt.Println(datas, reflect.TypeOf(datas))
	fmt.Println(reflect.TypeOf(datas))
	fmt.Println(reflect.TypeOf(datas))
	fmt.Println(reflect.TypeOf([]any{"a", "b"}))
	for _, data := range datas {
		fmt.Println(data)
	}
}
func main() {
	readData("a", 12, map[string]int{"a": 1})
}
