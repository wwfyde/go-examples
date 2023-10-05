package recursive

func Sum(n int) int {
	if n == 0 {
		return 0
	} else {
		return n + Sum(n-1)
	}
}

// 中文字体
