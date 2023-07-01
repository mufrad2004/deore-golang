package main

import "fmt"

func faktor(a int) int {
	var hasil int = 1
	for i := 1; i <= a; i++ {
		hasil *= i
	}
	return hasil
}

func permutasi(a, b int) int {
	var selisih int
	var hasil int
	if a >= b {
		selisih = a - b
		hasil = faktor(a) / faktor(selisih)
	} else {
		selisih = b - a
		hasil = faktor(b) / faktor(selisih)
	}
	return hasil
}

func main() {
	var (
		a, b int
	)

	fmt.Scan(&a, &b)
	fmt.Print(faktor(a), faktor(b), permutasi(a, b))

}
