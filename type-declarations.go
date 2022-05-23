package main

import "fmt"

func main() {
	type noKtp string
	type Married bool

	var ktp noKtp = "123"
	var marriedStatus Married = true

	fmt.Println(ktp)
	fmt.Println(marriedStatus)
}
