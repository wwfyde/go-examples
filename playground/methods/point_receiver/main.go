package main

import (
	"fmt"
	"reflect"
)

type Point struct {
	X, Y float64
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// type P *int
// func (P nil) f() { /**/ }
func main() {
	p := Point{1.2, 1.5}
	p.ScaleBy(2.0)
	fmt.Printf("%+v %#[1]v %#.2[1]f \n", p)
	fmt.Println("%T", reflect.TypeOf(p))

}
