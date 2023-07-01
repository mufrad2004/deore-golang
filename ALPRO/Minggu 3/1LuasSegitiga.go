package main

import "fmt"

func hitungLuas(alas, tinggi int, luas *float64) {
	/*
	   Initial State bilangan bulat yang menyatakan alas dan tinggi dari segitiga
	   Final State luas berisi luas dari segitiga, yaitu setengah alas kali tinggi
	*/

	*luas = (float64(alas) * float64(tinggi)) / 2

}

func main() {
	var (
		alas, tinggi, n int
		hasil           float64
	)

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&alas, &tinggi)
		hitungLuas(alas, tinggi, &hasil)
		fmt.Println("Hasil ke ", i, " : ", hasil)
	}
}
