package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           uint8
}

func main() {
	var customer Customer
	customer.Name = "Robbi"
	customer.Address = "Bandung"
	customer.Age = 18

	fmt.Println(customer)
	fmt.Println(customer.Name)

	customer1 := Customer{
		Name:    "Joko",
		Address: "Jkt",
		Age:     20,
	}

	fmt.Println(customer1)
}
