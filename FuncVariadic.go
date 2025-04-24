package main

import "fmt"

func data(nama string, nums ...int) (string, int) {

	angka := 0
	for _, value := range nums {
		angka += value
	}
	return nama, angka

}

func main() {

	nama, angka := data("Budi")
	fmt.Println("nama saya", nama, "usia", angka, "Tahun")

	nama, angka = data("Budi")
	fmt.Println("nama saya", nama, "usia", angka, "Tahun")

}
