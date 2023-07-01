package main

import "fmt"

func jumlahBus_1301223029(penumpang, kapasistas int) int {
	// Mengembalikan jumlah Bus yang diperlukan, apabila diketahui jumlah penumpang dan kapasitas dari Bus
	var (
		bus int
	)
	if kapasistas >= penumpang {
		bus++
	} else {
		if penumpang%kapasistas == 0 {
			bus = penumpang / kapasistas
		} else {
			bus++
			for penumpang > kapasistas {
				kapasistas += kapasistas
				bus++
			}
		}

	}

	return bus
}

func main() {
	var (
		penumpang, kapasistas int
	)

	fmt.Scan(&penumpang, &kapasistas)

	fmt.Print(jumlahBus_1301223029(penumpang, kapasistas), " bus")
}
