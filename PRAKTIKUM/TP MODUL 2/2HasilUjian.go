package main

import "fmt"

func main() {
	var (
		i, n, n_passed, n_failed int
		n1, n2, n3, avg          float64
	)
	fmt.Print("Berapa jumlah siswa yang nilainya akan diproses? ")
	fmt.Scan(&n)
	n_passed = 0
	n_failed = 0
	for i = 1; i <= n; i++ {
		fmt.Scan(&n1, &n2, &n3)
		avg = (n1 + n2 + n3) / 3

		if avg > 80.0 {
			fmt.Println("Memenuhi syarat administratif")
			n_passed++
		} else {
			fmt.Println("Tidak memenuhi syarat administratif")
			n_failed++
		}
	}

	fmt.Println("Jumlah siswa lolos seleksi admistrasi", n_passed)
	fmt.Println("Jumlah siswa tidak lolos seleksi admistrasi", n_failed)
}
