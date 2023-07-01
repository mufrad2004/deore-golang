package main

import "fmt"

func konversiCelcius(c int, r, f, k *float64) {
	*r = (float64(c) * 4.0) / 5.0
	*f = (float64(c)*9.0)/5.0 + 32
	*k = (float64(c)) + 273.15
}

func main() {
	var (
		c       int
		r, f, k float64 = 0, 0, 0
	)

	fmt.Scan(&c)
	konversiCelcius(c, &r, &f, &k)
	fmt.Println(int(r), "R")
	fmt.Println(int(f), "F")
	fmt.Println(k, "K")
}
