package main

import "fmt"

func logging() {
	fmt.Println("Loggin")
}

func runApp() {
	defer logging()
	fmt.Println("run app")
}

func main() {
	runApp()
}
