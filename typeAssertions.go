package main

import "fmt"

func random() interface{} {
	return true
}

func main() {

	var data interface{} = random()

	switch value := data.(type) {
	case string:
		fmt.Println(value, "ini string")
	case int:
		fmt.Println(value, "ini int")
	default:
		fmt.Println(value, "ini default")

	}

}
