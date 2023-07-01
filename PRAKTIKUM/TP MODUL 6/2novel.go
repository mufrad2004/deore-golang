package main

import "fmt"

func beliBuku_1301223029(N, M int) {
	if N-M == 1 {
		fmt.Print("beli 1 buku, rak buku penuh")
	} else {
		fmt.Println("beli 1 buku, sisa", (N-M)-1)
		beliBuku_1301223029(N, M+1)
	}
}

func main() {
	var (
		N, M int
	)

	fmt.Scan(&N, &M)
	fmt.Println("Sisa rak kosong ", N-M, " buku")
	beliBuku_1301223029(N, M)
}
