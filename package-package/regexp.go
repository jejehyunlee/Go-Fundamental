package main

import (
	"fmt"
	"regexp"
)

type DataEmail struct {
	Email string
}

func regex(email DataEmail) bool {
	dataRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

	if !dataRegex.MatchString(email.Email) {
		fmt.Println("Email Tidak Valid")
		return false
	}
	fmt.Println("Email Benar")
	fmt.Println(dataRegex.MatchString(email.Email))
	return true
}

func main() {

	email := DataEmail{
		Email: "jefri@gmail.com",
	}
	regex(email)

}
