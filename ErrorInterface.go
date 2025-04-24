package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("tidak bisa dibagi 0")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {

	var data, err = pembagian(10, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Hasil", data)
	}
}
