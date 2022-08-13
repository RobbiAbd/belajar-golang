package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Robbi"
	names[1] = "Abdul Rohman"

	fmt.Println(names)

	var values = [3]uint8{
		10, 20, 30,
	}

	fmt.Println(values)
	fmt.Println(len(values)) // panjang array
}
