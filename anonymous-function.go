package main

import "fmt"

type Blacklist func(string) bool

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("robbi", blacklist)
	registerUser("root", func(name string) bool {
		return name == "root"
	})
}

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
