package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name    string
	Age     int
	Address string
}

func main() {
	stud := Student{Name: "John Doe", Age: 30, Address: "123 Main St"}

	fieldName := "Age"
	value := reflect.ValueOf(stud)
	field := value.FieldByName(fieldName)

	if field.IsValid() {
		fmt.Printf("The struct has the field '%s'.\n", fieldName)
	} else {
		fmt.Printf("The struct does not have the field '%s'.\n", fieldName)
	}
}
