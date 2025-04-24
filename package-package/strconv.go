package main

import (
	"fmt"
	"strconv"
)

func main() {

	str := "1000000000"
	//Dari String Ke Integer
	number, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(number)
	}

	angka, _ := strconv.Atoi("true")
	fmt.Println(angka)

	value := strconv.FormatInt(1000, 10)

	fmt.Println(value)
}
