package main

import "fmt"

type a int

type b a

type Name struct {
	ID   int
	Name string
}
type 名称 = Name
type 标识符 名称

func main() {
	a := 名称{
		ID:   12,
		Name: "name",
	}
	b := 标识符{
		ID:   13,
		Name: "b",
	}
	c := Name{
		ID:   13,
		Name: "b",
	}
	d := b
	var f int
	var g = f
	g = 12
	var h any = 33

	fmt.Println(a.ID, a.Name)
	fmt.Printf("%T, %T, %v, %+v\n", a, a, a, a)
	fmt.Printf("%T, %T, %T, %T\n", b, a, c, d)
	s := fmt.Sprintf("%[3]*.[2]*[1]f\n", 12.0, 2, 7)
	s2 := fmt.Sprintf("%[3]*.[3]*[1]f\n", 12.0, 2, 6)
	s3 := fmt.Sprintf("%.[3]*[1]f\n", 12.0, 2, 6)
	s4 := fmt.Sprintf("%6.2[1]f\n", 12.0, 2, 6)
	fmt.Println(s, s2, s3, s4)
	fmt.Println(12, 12)
	if res, ok := h.(int); ok {
		fmt.Println(res)
	} else {
		fmt.Println("res is not int type")
	}
	fmt.Println(g)

}
