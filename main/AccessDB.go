package main

import (
	_ "Belajar_Golang/helper" // BLANK IDENTIFIER
	"Belajar_Golang/package-init/database"
	"fmt"
)

func main() {

	getData := database.GetDB()
	fmt.Println(getData)

}
