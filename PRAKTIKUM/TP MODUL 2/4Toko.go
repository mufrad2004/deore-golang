package main

import "fmt"

func main() {
	var (
		tahunLahir, totalBelanja, diskon int
		digit                            [4]int
		jmlHarga                         float64
	)

	fmt.Scan(&tahunLahir, &totalBelanja)

	for i := 0; i <= 3; i++ {
		digit[i] = (tahunLahir % 10)
		tahunLahir /= 10
	}
	diskon = digit[3] * ((digit[1] * 10) + digit[0])
	fmt.Println("besar diskon : ", diskon, "%")
	jmlHarga = float64(totalBelanja) - (float64(totalBelanja) * (float64(diskon) / 100))
	fmt.Print("Jumlah yang dibayar : ", int(jmlHarga))

	// var (
	// 	tahunLahir, totalBelanja, digit1, digit2, digit3, diskon int
	// )
	// fmt.Scan(&tahunLahir, &totalBelanja)

	// digit3 = tahunLahir % 10
	// tahunLahir = tahunLahir / 10
	// digit2 = tahunLahir % 10
	// tahunLahir = tahunLahir / 100
	// digit1 = tahunLahir % 10

	// diskon = (digit2 * 10) + digit3
	// fmt.Println(diskon, "%")
	// fmt.Print(digit1, digit2, digit3)

}
