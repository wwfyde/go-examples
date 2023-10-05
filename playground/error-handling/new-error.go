package main

import (
	"errors"
	"fmt"
	"reflect"
)

type MyError struct {
	Context string
}

func (e MyError) Error() string {
	return fmt.Sprintf("My first Error: %s", e.Context)
}

type NotFoundError struct {
	s string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("Not found error: %s", e.s)
}

func doSomething(n int) (bool, error) {
	if n == 1 {
		return false, MyError{"thrown error"}
	} else if n == 2 {
		return false, NotFoundError{"/a/b/c"}
	}

	return true, nil
}
func main() {
	ok, err := doSomething(3)
	if !ok {
		fmt.Println(err.Error())

		switch e := err.(type) {
		case MyError:
			fmt.Println(e.Error())
		case NotFoundError:
			fmt.Println(e.Error())
		default:
			fmt.Println("unknown error type")
		}
	} else {
		fmt.Println("no error")
	}
	e := errors.New("demo")
	fmt.Println("error type: ", reflect.TypeOf(e))
}
