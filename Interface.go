package main

import "fmt"

type HasName interface {
	GetName() string
}

type PersonData struct {
	First string
	Last  string
}

type Animal struct {
	sound string
}

func (person PersonData) GetName() string {
	fullName := "nama saya " + person.First + " " + person.Last
	return fullName
}

func (animal Animal) GetName() string {
	return animal.sound
}

func sayHello(name HasName) {

	fmt.Println("Hello", name.GetName())

}

func main() {

	var data PersonData
	data.First = "Jefri"
	data.Last = "Saputra"
	sayHello(data)

	kucing := Animal{
		sound: "meow",
	}

	sayHello(kucing)

}
