package main

import "fmt"

func main() {
	var s []int    // len(s) == 0, s == nil
	s = nil        // len(s) == 0, s == nil
	s = []int(nil) // len(s) == 0, s == nil
	s = []int{}    // len(s) == 0, s != nil
	v := []int(nil)
	//println(s, v)
	s2 := make([]int, 3)
	fmt.Printf("%T %[1]v \n", v)
	fmt.Println(s, s == nil, cap(s), len(s))
	fmt.Println(s2, s2 == nil, cap(s2), len(s2))
	raw := [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(raw[:])
	reverse(raw[:3])
	reverse(raw[3:6])
	fmt.Println(raw[:])

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
