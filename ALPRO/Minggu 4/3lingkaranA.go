package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(((a - c) * (a - c)) + ((b - d) * (b - d)))

}

func didalam(x, y int, a, b, r int) bool {
	var (
		jaraknya float64
	)

	jaraknya = jarak(float64(x), float64(y), float64(a), float64(b))

	if jaraknya < float64(r) {
		return true
	} else {
		return false
	}

}

func main() {
	var (
		// jaraknya       float64
		hasil1, hasil2 bool
		x1, y1, x2, y2 int
		r1, r2         int
		x, y           int
	)

	fmt.Scan(&x1, &y1, &r1)
	fmt.Scan(&x2, &y2, &r2)
	fmt.Scan(&x, &y)
	hasil1 = didalam(x, y, x1, y1, r1)
	hasil2 = didalam(x, y, x2, y2, r2)
	fmt.Print(hasil1, hasil2)

	if hasil1 && hasil2 {
		fmt.Print("Titik di dalam lingkaran ", 1, " dan ", 2)
	} else if !hasil1 && !hasil2 {
		fmt.Print("Titik di luar lingkaran ", 1, " dan ", 2)
	} else if hasil1 {
		fmt.Print("Titik di luar lingkaran ", 1)
	} else if hasil2 {
		fmt.Print("Titik di luar lingkaran ", 2)
	}

}
