package main

import "fmt"

func konversi(a, b, c int) int {
	var (
		hasil int
	)
	hasil += a * 3600
	hasil += b * 60
	hasil += c

	return hasil
}

func main() {
	var (
		a, b, c int
	)

	fmt.Scan(&a, &b, &c)
	fmt.Print(konversi(a, b, c))

}
