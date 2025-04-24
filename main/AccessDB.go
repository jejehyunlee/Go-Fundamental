package main

import (
	//_ "Belajar_Golang/helper" // BLANK IDENTIFIER
	"github.com/jejehyunlee/Go-Fundamental/v2/package-init/database"

	"fmt"
)

func main() {

	getData := database.GetDB()
	fmt.Println(getData)

}
