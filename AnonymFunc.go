package main

import "fmt"

type blacklist func(string) bool

func registerUser(name string, blacklist blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {

	registerUser("admin", func(name string) bool {

		return name == "admin"
	})

	blacklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", blacklist)
	registerUser("Jeje", blacklist)
}
