package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [...]int{1, 3}
	a[0] = 11
	var d []int
	//append(d, 12)
	fmt.Println(cap(d), len(d), d)
	fmt.Println(a, b, c) //true false false
	//d := [3]int{1, 2}
	//fmt.Println(a == d) // mismatched types [2]int and [3]int
}
