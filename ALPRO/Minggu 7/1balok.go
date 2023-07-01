package main

import "fmt"

func main() {
	var (
		balok                         []int
		masukkan, luasSelimut, volume int
	)

	for i := 0; i <= 1; i++ {
		fmt.Scan(&masukkan)
		balok = append(balok, masukkan)
	}

	fmt.Print(balok)
	luasSelimut = ((balok[0] * balok[1]) * 4) + ((balok[1] * balok[1]) * 2)
	volume = 
}
