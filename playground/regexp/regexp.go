package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "吴方圆你会不认识这几个字龦龯龿"
	bytes := []byte(s)
	matched, err := regexp.Match(`\p{Han}*?`, bytes)
	//matched, err := regexp.Match(`[a-z\x{300}-\x{302}\x{304}\x{308}\x{30C}]+?`, bytes)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if matched {
		print("matched!\n")
	} else {
		fmt.Println("not matched!")
	}

	re := regexp.MustCompile(`\p{Han}+`)
	str := "foo bar 吴方圆你会不认识这几个字龦龯龿鑨foo baz你? ,,,。"
	result := re.ReplaceAllString(str, "[-吴方圆-]")
	fmt.Println(result)
}
