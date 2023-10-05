package packages_test

import (
	"fmt"
	p "github.com/wwfyde/go-examples/playground/packages"
)

func ExampleLength() {
	//x := Point{}
	x := p.Point{}
	y := p.Point{X: 1, Y: 1}
	fmt.Println("Hello, 世界!")
	fmt.Println(p.Length(x, y))
}
