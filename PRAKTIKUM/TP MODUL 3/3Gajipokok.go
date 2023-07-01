package main

import "fmt"

func main() {
	var (
		golongan            string
		jr, jl              int
		gajiA, gajiB, total int
	)

	for golongan != "Z" {
		fmt.Scan(&golongan, &jr, &jl)
		if golongan == "A" {
			gajiA = (jr * 75000) + (jl * 90000)
		} else if golongan == "B" {
			gajiB = (jr * 125000) + (jl * 70000)
		} else {
			fmt.Println("Golongan tidak dikenali")
		}
	}
	total = gajiA + gajiB

	fmt.Print("Biaya yang dikeluarkan perusahaan untuk gaji karyawan ", total)
}
