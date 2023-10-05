package main

import (
	"fmt"
	"gopl.io/ch4/treesort"
	"time"
)

// Employee 结构体声明
type Employee struct {
	ID            int
	Name, Address string // 相同且连续的成员变量可以的写法
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

func main() {
	var demo Employee
	demo.ID = 100
	type Point struct {
		X int
		Y int
	}
	pp := &Point{1, 2}
	fmt.Printf("%[1]T %[1]v %[2]T %[2]v \n", pp, *pp)
	id := &demo.ID
	mySlice := []int{2, 1, 3}
	treesort.Sort(mySlice)
	fmt.Println(mySlice)
	*id += 22
	fmt.Printf("%T %[1]v %[2]T %[2]v", id, *id)

}
