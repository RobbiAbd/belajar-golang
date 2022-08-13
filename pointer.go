package main

import "fmt"

func changeCountyToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{
		City:     "Majalengka",
		Province: "Jabar",
		Country:  "ID",
	}

	address2 := &address1

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)

	// merubah semua data
	*address2 = Address{
		City:     "Surabaya",
		Province: "Jawa Timur",
		Country:  "ID",
	}

	fmt.Println(address1)
	fmt.Println(address2)

	address3 := new(Address)
	address3.Province = "Jawa Tengah"
	fmt.Println(address3)

	changeCountyToIndonesia(&address1)
	fmt.Println(address1)
}
