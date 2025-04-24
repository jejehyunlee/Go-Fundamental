package main

import (
	"fmt"
	"os"
)

func main() {

	hostname := os.Getenv("HOSTNAME")
	fmt.Println(hostname)

	name, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(name)

	}
	
}
