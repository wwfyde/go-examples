package main

func main() {
	println(MajorityElements([]int{2, 2, 1, 1, 1, 3, 4}))
}

func MajorityElements[T ~int | ~float64](nums []T) (res T) {
	i := 0
	for _, n := range nums {
		//
		if i == 0 {
			res = n
			i = 1
		} else if n == res {
			i += 1
		} else {
			i -= 1
		}
		println(n, i, res)
	}
	return
}
