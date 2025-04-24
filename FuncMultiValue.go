package main

import "fmt"

func dataSaya(nama string, umur int) (string, int) {
	return nama, umur

}

func main() {

	nama, umur := dataSaya("Budi", 20)

	fmt.Println("Nama saya", nama, "Umur saya", umur)

}
