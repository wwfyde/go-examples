package struct_interface_constraint

import "fmt"

type Constraint interface {
	Hello()

	Print() string
}

type I struct {
	Constraint
	Name string
}

func (i *I) Hello() {
	fmt.Println(i.Name)
}

func (i *I) Print() string {
	return i.Name
}
