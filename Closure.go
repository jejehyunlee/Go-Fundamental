package main

import "fmt"

//func clsr() func(int) int {
//
//
//
//
//}

func main() {

	counter := 0
	increment := func() {
		fmt.Println("incremment")
		counter++

	}

	increment()
	fmt.Println(counter)
}
