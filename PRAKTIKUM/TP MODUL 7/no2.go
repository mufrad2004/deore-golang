package main

import "fmt"

func average_1301223029(bil, i int, hasil *float64) {

	if bil <= 0 {
		*hasil = *hasil / float64(i)
	} else {
		digit := bil % 10
		*hasil += float64(digit)
		average_1301223029(bil/10, i+1, hasil)
	}

}

func main() {
	var (
		angka int
		hasil float64
	)

	fmt.Scan(&angka)
	if angka == 0 {
		fmt.Print(0)
	} else {
		average_1301223029(angka, 0, &hasil)
		fmt.Print(hasil)
	}
}
