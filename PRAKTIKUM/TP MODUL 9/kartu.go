package main

import "fmt"

func main() {
	var (
		arrKartu      []int
		kartu         int
		hitung_mundur int
	)

	fmt.Scan(&kartu)
	for kartu != 0 && len(arrKartu) < 52 {
		arrKartu = append(arrKartu, kartu)
		fmt.Scan(&kartu)
	}

	hitung_mundur = len(arrKartu) - 1
	for hitung_mundur >= 0 {
		fmt.Print(arrKartu[hitung_mundur], " ")
		hitung_mundur--
	}
}
