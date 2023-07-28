package helloworld

import "fmt"

func hello(in string) string {
	c := fmt.Sprintf("hello, %s!", in)
	return c

}
