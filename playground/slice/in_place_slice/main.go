package main

import "fmt"

func main() {
	s := []string{"abcd", "", "ef"}
	s = noempty(s)
	fmt.Printf("%v", s)
}

func noempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// 前面的值不变, 只需考虑被修改部分的值, 并返回新的 slice
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])

	return slice[:len(slice)-1]

}

// 如果不考虑顺序, 直接将末尾的值插入到要被移除的元素上, 并返回新的 slice
func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
