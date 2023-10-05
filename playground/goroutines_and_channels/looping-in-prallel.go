package main

import (
	"fmt"
)

func readFiles(fileNames []string) (newFileNames []string) {
	for _, filename := range fileNames {
		fmt.Println(filename)
		newFileNames = append(newFileNames, filename)
	}
	return newFileNames
}

func main() {
	filenames := []string{"a", "b", "c"}
	ch := make(chan string)
	for _, f := range filenames {
		go func(f string) {
			fmt.Println(f)
			ch <- f

		}(f)
	}
	println("hello")
	println("hello")
	println("hello")
	println("hello")
	println("hello")
	println("hello")
	println("hello")
	println("hello")
	println("hello")
	println("hello")
	for range filenames {
		<-ch
	}
	//time.Sleep(5 * time.Second)
}
