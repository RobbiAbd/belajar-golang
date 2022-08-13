package main

import "fmt"

func main() {
	var nilaiAkhir uint8 = 90
	var absensi uint8 = 80

	lulusNilaiAkhir := nilaiAkhir > 80
	lulusAbsensi := absensi > 80

	lulus := lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulusNilaiAkhir)
	fmt.Println(lulusAbsensi)
	fmt.Println(lulus)
}
