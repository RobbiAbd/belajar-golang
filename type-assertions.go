package main

import (
	"fmt"
)

func random() interface{} {
	return "OK"
}

func main() {
	result := random()

	switch resultString := result.(type) {
	case string:
		fmt.Println("string =", resultString)
	case int:
		fmt.Println("int =", resultString)
	default:
		fmt.Println("unknown")
	}

}
