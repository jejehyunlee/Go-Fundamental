package main

import (
	"Belajar_Golang/helper"
	"fmt"
)

func main() {

	data := helper.SayHello("Jefri")

	fmt.Println(data)

}

// Acces Modifer : jika Func diawali huruf BESAR bisa diakses di luar package

// Acces Modifer : jika Func diawali huruf KECIL,,-> TIDAK bisa diakses di luar package
