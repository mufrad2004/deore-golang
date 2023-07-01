package main

import "fmt"

const phi float64 = 3.14

func hitungVolume(r, t float64) float64 {
	return (phi * r * r) * t
}

func hitungMassa(r, t, p float64) float64 {
	return hitungVolume(r, t) * p
}

func display(m1, m2 float64) {
	if m1 == m2 {
		fmt.Println("BALANCE")
	} else if m1 < m2 {
		fmt.Println("Selisih : ", m2-m1)
	} else {
		fmt.Println("Selisih : ", m1-m2)
	}
}

func main() {
	var (
		r                     float64
		tKiri, tKanan         float64
		mjKiri, mjKanan       float64
		massaKiri, massaKanan float64
	)

	fmt.Scan(&r)
	fmt.Scanln(&tKiri, &mjKiri)
	fmt.Scanln(&tKanan, &mjKanan)

	massaKiri = hitungMassa(r, tKiri, mjKiri)
	massaKanan = hitungMassa(r, tKanan, mjKanan)
	display(massaKiri, massaKanan)
}
