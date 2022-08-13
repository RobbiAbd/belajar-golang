package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	res := Man{
		Name: "Eko",
	}
	res.Married()

	fmt.Println(res.Name)
}
