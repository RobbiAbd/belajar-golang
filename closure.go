package main

import "fmt"

func main() {
	counter := 0
	name := "Robbi"

	increment := func() {
		fmt.Println("increment")
		name := "Abdul"
		fmt.Println(name)
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
