package main

import "fmt"

func main() {
	var (
		total1, total2, total3 int
		predicate              string
	)

	for predicate != "STOP" {
		fmt.Scan(&predicate)
		if predicate == "Sufficient" {
			total1 += 1
		} else if predicate == "Good" {
			total2 += 1
		} else if predicate == "Excellent" {
			total3 += 1
		}
	}

	fmt.Println("Siswa predikat Sufficient = ", total1)
	fmt.Println("Siswa predikat Good = ", total2)
	fmt.Println("Siswa predikat Excelent = ", total3)
}
