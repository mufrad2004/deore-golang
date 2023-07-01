package main

import (
	"fmt"
	"math"
)

func z1_1301223029(x, y int) float64 {
	// Mengembalikan nilai z sesuai persamaan di atas dan bilangan bulat x dan y
	var (
		hasil float64
	)

	hasil = math.Sqrt((3 * float64(x)) * (6 * float64(y)) / (4 * float64(y)))

	return hasil

}

// func z2_1301223029(x, y int) float64 {
// 	// Mengembalikan nilai z sesuai persamaan di atas dan bilangan bulat x dan y
// 	var (
// 		hasil float64
// 	)

// 	hasil = math.Sqrt((3 * float64(y)) * (6 * float64(x)) / (4 * float64(x)))

// 	return hasil

// }

func main() {
	var (
		a, b int
	)

	fmt.Scan(&a, &b)
	fmt.Print(z1_1301223029(a, b), z1_1301223029(b, a))
}
