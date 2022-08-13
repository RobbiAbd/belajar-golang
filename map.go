package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Robbi",
		"address": "Bandung",
	}

	person["age"] = "19"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
}
