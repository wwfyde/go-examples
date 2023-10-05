package read_file

import (
	"os"
)

func Read() string {

	//data, _ := os.ReadFile("/Users/wwfyde/Desktop/Go/go-playground/files/demo.txt")
	data, err := os.ReadFile("files/read_file/demo.txt")

	if err != nil {
		panic(err)
	}

	return string(data)
}
