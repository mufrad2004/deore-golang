package main

import "fmt"

func main() {
	var (
		uang, uangTerendah, uangTertinggi, total, i int
		rata_rata                                   float64
	)

	fmt.Scan(&uang)

	uangTerendah = uang
	uangTertinggi = uang
	for uang >= 0 {
		total += uang
		i++

		if uang < uangTerendah {
			uangTerendah = uang
		}

		if uang > uangTertinggi {
			uangTertinggi = uang
		}
		fmt.Scan(&uang)
	}

	rata_rata = float64(total) / float64(i)

	fmt.Println("Jumalah = ", total)
	fmt.Println("Rata - rata = ", rata_rata)
	fmt.Println("Uang tertinggi = ", uangTertinggi)
	fmt.Println("Uang terendah = ", uangTerendah)

}
