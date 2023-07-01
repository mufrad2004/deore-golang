package main

import "fmt"

func main() {
	var (
		n        int
		tabAngka = []int{2, 5, 4, 3, 2, 1, 1, 2, 3, 4, 5, 6, 2, 2, 2, 2, 2, 1, 7, 8, 9, 6, 9}
		baskom   int
		buku     int
	)

	// descending
	n = len(tabAngka)

	for pass := 1; pass < n; pass++ {
		// orang yang mengang buku
		buku = pass - 1
		// cari max, mulai bandingin dari orang ke 2
		for i := pass; i < n; i++ {
			if tabAngka[buku] < tabAngka[i] {
				buku = i
			}
		}
		// tuker orang max dengan orang yang megang buku
		baskom = tabAngka[buku]
		tabAngka[buku] = tabAngka[pass-1]
		tabAngka[pass-1] = baskom
	}

	fmt.Print(tabAngka)
}
