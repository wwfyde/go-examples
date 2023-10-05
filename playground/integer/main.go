package main

import "fmt"

func main() {
	a := -1.5e-2
	mod1, mod2, mod3, mod4 := 3.5, -3.5, 0.5, -0.5
	fmt.Println(a, int(a))
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
	var i, j int8 = 2, 5
	fmt.Println(i, i+1, i*i)

	fmt.Println(i&^j, i&j)

	fmt.Println(int(mod1), int(mod2), int(mod3), int(mod4))
}
