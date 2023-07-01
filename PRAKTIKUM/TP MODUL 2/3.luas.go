package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		panjang, lebar, keliling, diagonal, luas int
	)

	fmt.Scan(&panjang, &lebar)
	luas = panjang * lebar
	diagonal = int(math.Sqrt((float64(panjang) * float64(panjang)) + (float64(lebar) * float64(lebar))))
	keliling = (2 * panjang) + (2 * lebar)

	fmt.Println("Luas : ", luas)
	fmt.Println("Keliling : ", keliling)
	fmt.Println("Panjang diagonal :", diagonal)
}
