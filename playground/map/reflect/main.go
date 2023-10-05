// check map with reflect package

package main

import "fmt"

func main() {
	type MyMap map[string]any
	m := MyMap{"a": 1, "b": 2, "c": "d"}
	if _, ok := m["c"]; ok == true {
		fmt.Printf("%v\n", m["a"])
	} else {
		fmt.Println("key demo no exists in m")
	}
}
