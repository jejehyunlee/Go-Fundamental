package main

import (
	"fmt"
	"reflect"
)

type dataPerson struct {
	Name string `required:"true"`
}

func isValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				fmt.Println(field.Name, "is required")
				return false
			}
		}
	}
	return true
}

func main() {

	dataField := dataPerson{
		Name: "",
	}
	fmt.Println(isValid(dataField))

}
