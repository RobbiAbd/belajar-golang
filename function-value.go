package main

import "fmt"

func main() {
	sayGoodBye := getGoodBye
	fmt.Println(sayGoodBye("robbi"))
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}
