package main

import "fmt"

func main() {
	total := sumAll(11, 10)
	fmt.Println(total)

	numbers := []int{10, 20}
	total = sumAll(numbers...)
	fmt.Println(total)
}

func sumAll(numbers ...int) int {
	total := 0

	for _, val := range numbers {
		total += val
	}

	return total
}
