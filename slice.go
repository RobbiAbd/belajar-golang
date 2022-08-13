package main

import "fmt"

func main() {
	var months = [...]string{
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"Mei",
		"Jun",
		"Jul",
		"Agu",
		"Sep",
		"Okt",
		"Nov",
		"Des",
	}

	fmt.Println(months)

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "Diubah"
	fmt.Println(slice1)

	slice1[1] = "Hello"
	fmt.Println(slice1)
	fmt.Println(months)

	var slice2 = months[2:4]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Robbi")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	// buat slice
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Robbi"
	newSlice[1] = "Abd"
	newSlice = append(newSlice, "Roh")

	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [...]string{"a", "b"}
	iniSlice := []string{"a", "b"}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
