package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesi")
	message := recover()

	if message != nil {
		fmt.Println("Error :", message)
	}
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("aplikasi error")
	}

	fmt.Println("run app")
}

func main() {
	runApp(false)
}
