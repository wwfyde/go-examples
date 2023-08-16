package hello

import "fmt"

func Hello(in string) string {
	c := fmt.Sprintf("Hello, %s!", in)
	return c
}
