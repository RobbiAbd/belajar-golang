package main

import "fmt"

func main() {
	loop := factorialLoop(5)
	loopRecursive := factorialRecursive(5)
	fmt.Println(loop)
	fmt.Println(loopRecursive)
}

func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
