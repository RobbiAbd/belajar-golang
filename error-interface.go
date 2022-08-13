package main

import (
	"errors"
	"fmt"
)

func pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Tidak boleh nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	var cthError = errors.New("error")

	fmt.Println(cthError.Error())

	hasil, err := pembagi(10, 10)

	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println(err.Error())
	}
}
