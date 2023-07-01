package main

import "fmt"

func inputBilangan_1301223029(bil *int) {
	fmt.Scan(&*bil)
	for *bil < 0 {
		fmt.Scan(&*bil)
	}

}

func stop_1301223029(bil int) bool {
	var (
		hasil bool
	)
	if bil == 0 {
		hasil = true
	} else {
		hasil = false
	}

	return hasil
}

func hitung_1301223029(mean *float64, min, max, n *int) {
	var (
		angka, jumlah int
	)

	inputBilangan_1301223029(&angka)
	*min = angka
	*max = angka
	for !stop_1301223029(angka) {
		jumlah += angka
		*n++
		if *min > angka {
			*min = angka
		}
		if *max < angka {
			*max = angka
		}
		inputBilangan_1301223029(&angka)

	}
	*mean = float64(jumlah) / float64(*n)
}

func main() {
	var (
		min, max, n int
		mean        float64
	)

	hitung_1301223029(&mean, &min, &max, &n)
	if n == 0 {
		fmt.Print("- - - - ")
	} else {
		fmt.Print(mean, max, min, n)
	}

}
