package main

import "fmt"

type User struct {
	Name string
}

func (u *User) SetName() {

	u.Name = "Mr." + u.Name
}

func main() {

	u := User{"Jefri"}
	u.SetName()
	fmt.Println(u)
}
