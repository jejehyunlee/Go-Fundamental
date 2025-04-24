package main

import "fmt"

func sayHelloNew(name string) interface{} {
	if name == "jeje" {
		return "hello jeje"
	} else if name == "budi" {
		return true
	} else {
		return 100
	}
}

func main() {

	var data interface{} = sayHelloNew("budi")
	fmt.Println(data)

	result := sayHelloNew("aman")
	fmt.Println(result)

}
