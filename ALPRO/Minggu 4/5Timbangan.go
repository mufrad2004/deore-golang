package main

import "fmt"

const pi float64 = 3.14

func volume(r, t float64) float64 {
	return (pi * (r * r) * t)
}

func massa(r, t, p float64) float64 {
	return volume(r, t) * p
}

func display(m1, m2 float64, pesan1 *string, pesan2 *float64) {
	if m1 == m2 {
		*pesan1 = "BELANCE"
	} else if m1 > m2 {
		*pesan2 = m1 - m2
	} else {
		*pesan2 = m2 - m1
	}

}

func main() {
	var (
		r, p1, t1, p2, t2 float64
		massa1, massa2    float64
	)

	fmt.Scan(&r)
	fmt.Scan(&p1, &t1)
	fmt.Scan(&p2, &t2)
	massa1 = massa(r, t1, p1)
	massa2 = massa(r, t2, p2)
	pesan1 := ""
	pesan2 := 0.0
	display(massa1, massa2, &pesan1, &pesan2)

	if pesan1 == "BELANCE" {
		fmt.Print(pesan1)
	} else {
		fmt.Print(pesan2)
	}

}
