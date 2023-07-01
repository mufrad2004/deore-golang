package main

import "fmt"

func pecahUang(uang int, k10, k2, k1 *int) {
	/*
	   I.S terdefinisi bilangan bulat yang menyatakan nominala uang
	   F.S variabel k10,k2,k1 berisi denom 10000,2000, 1000 dari nominal uang
	*/

	for uang != 0 {
		if uang%10000 == 0 {
			*k10++
			uang -= 10000
		} else if uang%2000 == 0 {
			*k2++
			uang -= 2000
		} else if uang%1000 == 0 {
			*k1++
			uang -= 1000
		} else {
			fmt.Println("Uang bukan kelipatan 10000, 2000, dan 1000")
			uang = 0
		}
	}

}

func main() {
	var (
		uang int
	)
	k10 := 0
	k2 := 0
	k1 := 0
	fmt.Scan(&uang)
	pecahUang(uang, &k10, &k2, &k1)
	fmt.Println(k10, " lembar 10000")
	fmt.Println(k2, " lembar 2000")
	fmt.Println(k1, " lembar 1000")

}
