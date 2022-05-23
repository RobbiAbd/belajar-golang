package main

import "fmt"

func main() {
	var name string

	name = "Robbi"
	fmt.Println(name)

	name = "Robbi Abdul"
	fmt.Println(name)

	var friendName = "Agus"
	fmt.Println(friendName)

	var age int8 = 20
	fmt.Println(age)

	address := "Bandung"
	fmt.Println(address)

	var (
		firstName      = "Robbi"
		lastName       = "Abdul"
		noHp      int8 = 123
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(noHp)
}
