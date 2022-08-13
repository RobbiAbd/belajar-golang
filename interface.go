package main

import "fmt"

type HashName interface {
	GetName() string
}

func sayHello(hashName HashName) {
	fmt.Println("Hello", hashName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	var eko Person
	eko.Name = "Eko"

	sayHello(eko)
}
