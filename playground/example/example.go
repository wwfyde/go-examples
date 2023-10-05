package example

import "fmt"

func Hello() {
	world := "世界!"
	fmt.Println("hello,", world)
}

func Hello2() {
	world := "吴方圆!"
	fmt.Printf("hello, %v", world)
}
