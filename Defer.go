package main

import "fmt"

func deferFunc() {
	fmt.Println("ini defer")
}

func runapp(value int) {
	defer deferFunc()
	fmt.Println("ini runapp")
	hasil := 10 / value
	fmt.Println(hasil)
}

func main() {
	runapp(10)
}
