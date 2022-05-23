package main

import "fmt"

func main() {
	const firstName = "Robbi"
	const lastName = "Abdul Rohman"

	fmt.Println(firstName)
	fmt.Println(lastName)

	const (
		age     int8 = 18
		address      = "Bandung"
	)

	fmt.Println(age)
	fmt.Println(address)
}
