package main

import "fmt"

func nilai(value int) int {

	hasil := 1
	for i := value; i > 0; i-- {
		hasil = hasil * i
	}
	return hasil
}

func factorialNew(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialNew(n-1)
}

func main() {

	factorial := nilai(3)

	fmt.Println(factorial)

	fmt.Println(3 * 2 * 1)

	fmt.Println(factorialNew(3))
}
