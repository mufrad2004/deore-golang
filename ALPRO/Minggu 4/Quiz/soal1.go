package main

import "fmt"

func proses(h1 string, h2 *string) {
	var (
		hari            []string
		hitung, putaran int
		day             string
	)

	hari = append(hari, "senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu")
	for i := 0; i >= 0; i++ {
		if i == 6 {
			putaran++
			i = 0
		}
		if day == hari[i] {
			hitung++
		}

		if hitung == 2 {
			break
		}
	}

	*h2 = day
}

func main() {
	var (
		hari string
	)

	fmt.Scan(&hari)
	proses(hari, &hari)
	fmt.Print(hari)
}
