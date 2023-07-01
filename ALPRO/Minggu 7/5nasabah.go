package main

import "fmt"

type nasabah struct {
	id, nama, namaBank, nomorRekening string
}

type tabNasabah [2022]nasabah

func addNasabah(t *tabNasabah, N int) {
	var (
		masukkan nasabah
	)

	for i := 0; i < N; i++ {
		fmt.Scan(&masukkan.id, &masukkan.nama, &masukkan.namaBank, &masukkan.nomorRekening)
		t[i] = masukkan
	}
}

func cetak(t tabNasabah, N int, X string) {
	for i := 0; i < N; i++ {
		if X == t[i].namaBank {
			fmt.Print("Kode : ", t[i].id, ", Nasabah : ", t[i].nama, ", Bank : ", t[i].namaBank, ", Rek : ", t[i].nomorRekening)
		}
	}
}

func main() {
	var (
		tab tabNasabah
		n   int
		x   string
	)

	fmt.Scan(&n)
	addNasabah(&tab, n)
	fmt.Scan(&x)
	cetak(tab, n, x)
}
