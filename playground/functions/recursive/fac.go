package recursive

func Fac(n int) (fac int) {
	if n <= 1 {
		return 1
	} else {
		return n * Fac(n-1)
	}
}
