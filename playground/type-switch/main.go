package main

import "fmt"

func typeSwitch(i any) {
	switch v := i.(type) {
	case int:
		fmt.Printf("twice of %v is %v\n", i, v*2)
	case string:
		fmt.Printf("len of %v is %v\n", i, len(v))
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	typeSwitch(12)
	typeSwitch("hello")
	typeSwitch([]any{1, "12"})
}
