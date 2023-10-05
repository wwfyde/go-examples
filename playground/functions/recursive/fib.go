package recursive

func Fib(n int) int {

	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		fib := Fib(n-1) + Fib(n-2)
		return fib
	}

}
func Fib2() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
