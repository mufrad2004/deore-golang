package main

import "fmt"

func main() {
	var (
		r, t float64
	)
	fmt.Scan(&r, &t)
	volume(r, t)
}

func volume(r, t float64) {
	var (
		pi     float64 = 3.14
		volume float64
	)

	volume = (pi * (r * r)) * t
	fmt.Print(volume)
}
