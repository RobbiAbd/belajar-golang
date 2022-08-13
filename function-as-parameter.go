package main

import "fmt"

type Filter func(string) string

func main() {
	filter := sayHelloWithFilter("anjing", spamFilter)
	fmt.Println(filter)
	filter = sayHelloWithFilter("robbi", spamFilter)
	fmt.Println(filter)
}

func sayHelloWithFilter(name string, filter Filter) string {
	nameFiltered := filter(name)
	return nameFiltered
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	}

	return name
}
