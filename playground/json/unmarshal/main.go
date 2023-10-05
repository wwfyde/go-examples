package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func Unmarhshal(data []byte) {
	type Title struct {
		Title string
	}
	var titles *[]Title // 声明结构体类型
	if err := json.Unmarshal(data, titles); err != nil {
		log.Fatalf("JSON Unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}

func main() {

}
